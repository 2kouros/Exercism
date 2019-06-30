//Package hamming contains the hamming distance calculator function
package hamming

import "errors"

//Distance compares two strings and calculates their hamming distance
func Distance(a, b string) (int, error) {

	hamming := 0

	if len(a) != len(b) {
		return hamming, errors.New("not same length of DNA genomes")
	}
	for i := range a {

		if a[i] != b[i] {
			hamming++
		}
	}

	return hamming, nil

}
