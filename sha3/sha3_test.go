package sha3

import (
	stdsha3 "crypto/sha3"
	"math/rand"
	"testing"

	"github.com/jnsoft/jngo/hex"
	. "github.com/jnsoft/jngo/testhelper"
)

func TestSha3_256(t *testing.T) {
	str := "Hello, World!"

	t.Run("SHA3_256_HelloWorld", func(t *testing.T) {
		bytes := []byte(str)
		hash := Hash256(bytes)
		hashStr := hex.ToHexString(hash, false)

		ref := stdsha3.Sum256(bytes)
		refStr := hex.ToHexString(ref[:], false)

		AssertEqual(t, hashStr, refStr)
	})

	t.Run("SHA3_256_nist_test_vector_abc", func(t *testing.T) {
		abc := "abc"
		// NIST SHA3-256("abc")
		expected := "3a985da74fe225b2045c172d6bd390bd855f086e3e9d525b46bfe24511431532"

		got := Hash256([]byte(abc))
		gotStr := hex.ToHexString(got, false)

		AssertEqual(t, gotStr, expected)
	})

	t.Run("SHA3_256_long_string_matches_stdlib", func(t *testing.T) {
		s := "abcdbcdecdefdefgefghfghighijhijkijkljklmklmnlmnomnopnopq"
		data := []byte(s)

		got := Hash256(data)
		gotStr := hex.ToHexString(got, true)

		ref := stdsha3.Sum256(data)
		refStr := hex.ToHexString(ref[:], true)

		AssertEqual(t, gotStr, refStr)
	})
}

func TestSha3_512(t *testing.T) {
	msgs := []string{
		"",
		"abc",
		"Hello, World!",
	}

	for _, s := range msgs {
		data := []byte(s)

		got := Hash512(data)
		gotStr := hex.ToHexString(got, false)

		ref := stdsha3.Sum512(data)
		refStr := hex.ToHexString(ref[:], false)

		AssertEqual(t, gotStr, refStr)
	}
}

func TestShake256(t *testing.T) {
	msg := []byte("Hello, World!")
	outLen := 1024

	got := Shake256(msg, outLen)
	gotStr := hex.ToHexString(got, false)

	ref := make([]byte, outLen)
	shake := stdsha3.NewSHAKE256()
	shake.Write(msg)
	shake.Read(ref)
	refStr := hex.ToHexString(ref, false)

	AssertEqual(t, gotStr, refStr)
}

func TestSha3_256_LargeInput(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping large-input SHA3 test in short mode")
	}

	const size = 256 * 1024 * 1024 // 256 MB
	data := make([]byte, size)

	// Deterministic pseudo-random data
	r := rand.New(rand.NewSource(42))
	if _, err := r.Read(data); err != nil {
		t.Fatalf("failed to generate random data: %v", err)
	}

	got := Hash256(data)
	ref := stdsha3.Sum256(data)

	// Compare raw bytes to avoid extra allocations
	if len(got) != len(ref) {
		t.Fatalf("length mismatch: got %d, want %d", len(got), len(ref))
	}
	for i := range got {
		if got[i] != ref[i] {
			t.Fatalf("hash mismatch at byte %d", i)
		}
	}
}
