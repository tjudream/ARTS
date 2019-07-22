package week18

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func judgeCircle(moves string) bool {
	x,y := 0,0
	for _, s := range moves {
		if s == 'U' {
			y++
		} else if s == 'D' {
			y--
		} else if s == 'L' {
			x--
		} else if s == 'R' {
			x++
		}
	}
	if x == 0 && y == 0 {
		return true
	}
	return false
}

func judgeCircle2(moves string) bool  {
	x,y := 0,0
	l := len(moves)
	for i := 0; i < l; i++ {
		if moves[i] == 'U' {
			y++
		} else if moves[i] == 'D' {
			y--
		} else if moves[i] == 'L' {
			x--
		} else if moves[i] == 'R' {
			x++
		}
	}
	if x == 0 && y == 0 {
		return true
	}
	return false
}

func judgeCircle3(moves string) bool {
	if (((strings.Count(moves, "U")*1) + (strings.Count(moves, "D")* -1)) == 0 &&
		((strings.Count(moves, "R")*1) + (strings.Count(moves, "L")* -1)) == 0) {
		return true
	}
	return false
}

func TestJudgeCircle(t *testing.T) {
	assert.Equal(t, true, judgeCircle("UD"))
	assert.Equal(t, false, judgeCircle("LL"))
}