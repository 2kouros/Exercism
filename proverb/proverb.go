//Package proverb contains the function for composing a proverb
package proverb

import (
	"fmt"
)

// Proverb composes the proverb based on the number of inputs it gets
func Proverb(rhyme []string) []string {

	proverb := []string{}

	if len(rhyme) == 0 {
		return proverb
	}

	for i, v := range rhyme {
		if i < len(rhyme)-1 {
			line := fmt.Sprintf("For want of a %s the %s was lost.", v, rhyme[i+1])
			proverb = append(proverb, line)
		}

	}
	proverb = append(proverb, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))

	return proverb
}
