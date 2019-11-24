package week35

import (
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"
)
func isScramble(s1 string, s2 string) bool {
    if s1 == s2 {
        return true
    }
    l := len(s1)
    var dic [26]int
    for i := 0; i < l; i++ {
      dic[s1[i]-'a']++
      dic[s2[i]-'a']--
    }

    for i := 0; i < 26; i++ {
      if dic[i] != 0 {
        return false
      }
    }

    for i := 1; i < l; i++ {
      if isScramble(s1[0:i],s2[0:i]) && isScramble(s1[i:l],s2[i:l]) {
        return true
      }
      if isScramble(s1[0:i],s2[l-i:l]) && isScramble(s1[i:l],s2[0:l-i]) {
        return true
      }
    }
    return false
}

func TestIsScramble(t *testing.T) {
  s1,s2 := "great","rgeat"
  assert.True(t, isScramble(s1,s2))
  s1,s2 = "abcde","caebd"
  assert.False(t, isScramble(s1,s2))
  s1,s2 = "abc","bca"
  assert.True(t,isScramble(s1,s2))
}
