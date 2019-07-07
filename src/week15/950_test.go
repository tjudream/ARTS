package week15

import (
	"github.com/golang-collections/collections/queue"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func deckRevealedIncreasing(deck []int) []int {
	N := len(deck)
	if N < 2 {
		return deck
	}
	sort.Ints(deck)
	res := []int{deck[N - 1]}
	for i := N - 2; i >= 0; i-- {
		tmp := res[len(res) - 1]
		res = res[:len(res) - 1]
		res = append([]int{deck[i], tmp}, res...)
	}
	return res
}

func deckRevealedIncreasingS1(deck []int) []int {
	N := len(deck)
	index := queue.New()
	for i := 0; i < N; i++ {
		index.Enqueue(i)
	}
	sort.Ints(deck)
	ans := make([]int, N)
	for i := 0; i < N; i++ {
		ans[index.Dequeue().(int)] = deck[i]
		if (index.Len() > 0) {
			index.Enqueue(index.Dequeue())
		}
	}
	return ans
}

func TestDeckRevealedIncreasing(t *testing.T) {
	input := []int{17,13,11,2,3,5,7}
	output := []int{2,13,3,11,5,17,7}
	assert.Equal(t, output, deckRevealedIncreasing(input))
}

func TestDeckRevealedIncreasingS1(t *testing.T) {
	input := []int{17,13,11,2,3,5,7}
	output := []int{2,13,3,11,5,17,7}
	assert.Equal(t, output, deckRevealedIncreasingS1(input))
}
