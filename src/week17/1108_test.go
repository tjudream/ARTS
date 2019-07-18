package week17

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func defangIPaddr(address string) string {
	return strings.Replace(address, ".", "[.]", -1)
}

func defangIPaddr2(address string) string {
	res := ""
	for _,s := range strings.Split(address, "") {
		if s == "." {
			res += "[.]"
		} else {
			res += s
		}
	}
	return res
}

func TestDefangIPaddr(t *testing.T)  {
	assert.Equal(t, "1[.]1[.]1[.]1", defangIPaddr("1.1.1.1"))
	assert.Equal(t, "1[.]1[.]1[.]1", defangIPaddr2("1.1.1.1"))
}
