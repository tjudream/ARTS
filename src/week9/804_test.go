package week9

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var dict = []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
func uniqueMorseRepresentations(words []string) int {
	morseMap := make(map[string]string)
	for _,word := range words {
		morse := ""
		for i := 0; i < len(word); i++ {
			morse = morse + dict[word[i] - 'a']
		}
		morseMap[morse] = word
	}
	return len(morseMap)
}

func TestUniqueMorseRepresentations(t *testing.T) {
	words := []string{"gin", "zen", "gig", "msg"}
	assert.Equal(t, 2, uniqueMorseRepresentations(words))
}
