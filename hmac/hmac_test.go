package hmac

import (
	"crypto/hmac"
	"crypto/sha256"
	"testing"

	"github.com/jnsoft/jngo/hex"
	. "github.com/jnsoft/jngo/testhelper"
)

func TestSha2(t *testing.T) {

	tests := []struct {
		key, msg, expected string
	}{
		{
			key:      "key1",
			msg:      "message1",
			expected: "1a3388d6717d3599d5fa0d5a3207c31287cd478d1d69f04cc3328ba474101032",
		},
		{
			key:      "key2",
			msg:      "message2",
			expected: "3f899d6b2a33b1a202344f54d27337e8ab144a9ed3ef83f4b2a96545544c7778",
		},
	}

	t.Run("HMAC Compute", func(t *testing.T) {
		for _, tt := range tests {
			key := []byte(tt.key)
			msg := []byte(tt.msg)
			expected, _ := hex.FromHexString(tt.expected)
			res := Compute(key, msg)
			CollectionAssertEqual(t, res, expected)

			h := hmac.New(sha256.New, key)
			h.Write(msg)
			hmac := h.Sum(nil)
			CollectionAssertEqual(t, res, hmac)
		}
	})

	t.Run("HMAC Verify", func(t *testing.T) {
		for _, tt := range tests {
			key := []byte(tt.key)
			msg := []byte(tt.msg)
			res := Compute(key, msg)
			isOk := Verify(key, msg, res)
			AssertTrue(t, isOk)
		}
	})

}
