package flock

// Identity function I := λx.x
func I[T any](f T) T {
	return f
}

// The self-application combinator M (for “Mockingbird”), aka ω, λf.ff
//func M[T any](f T) {
//	return f(f)
//}

// B combinator: function composition
func B(f, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}

// Cardinal := C := flip := λfab.fba
func C[A, B, C any](f func(A) func(B) C) func(B) func(A) C {
	return func(b B) func(A) C {
		return func(a A) C {
			return f(a)(b)
		}
	}
}

// Conditional combinator
func Cond(p func(int) bool, f, g func(int) int) func(int) int {
	return func(x int) int {
		if p(x) {
			return f(x)
		}
		return g(x)
	}
}

// Y combinator for recursion
func Y(f func(func(int) int) func(int) int) func(int) int {
	return func(x int) int {
		return f(Y(f))(x)
	}
}

// K combinator: constant function
func K(x int) func(int) int {
	return func(_ int) int {
		return x
	}
}

// S combinator: applies f to x and g(x)
func S(f func(int, int) int, g func(int) int) func(int) int {
	return func(x int) int {
		return f(x, g(x))
	}
}

func True[T any](x T) func(T) T {
	return func(y T) T {
		return x
	}
}

func False[T any](x T) func(T) T {
	return func(y T) T {
		return y
	}
}

// facCps calculates the factorial in continuation-passing style
func facCps(n int, k func(int) int) int {
	if n == 0 {
		return k(1) // Base case: Apply the continuation to 1
	}
	return facCps(n-1, func(result int) int {
		return k(n * result) // Recursive case: Compose the continuation
	})
}

// fac is a simple wrapper for facCps with the identity continuation
func fac(n int) int {
	return facCps(n, func(x int) int { return x }) // `k` starts as the identity function
}
