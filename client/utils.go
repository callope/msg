package client

import (
  "time"
  "math/rand"
)

// newHashString generates a hash string.
func newHashString(n int) string {
  rand.Seed(time.Now().UnixNano())
  chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890"
  hashstr := make([]byte, n)

  for i := 0; i < n; i++ {
    hashstr[i] = chars[rand.Intn(60)]
  }

  return string(hashstr)
}
