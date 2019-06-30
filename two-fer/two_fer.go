/*Package twofer contains the source code of the program that prints "One for X, one for me" ,
where X is a given input name ,or in case of no input is "you" */
package twofer

import (
	"fmt"
)

// ShareWith takes as input a name  of type string and returns a sentence of type
//string the "One for X, one for me."
// if no name is given , adds "you"  .
func ShareWith(name string) (sentence string) {
	if name == "" {
		name = "you"

	}
	sentence = fmt.Sprintf("One for %s, one for me.", name)

	return

}
