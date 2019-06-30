//Package strand contains the function for RNA composition
package strand

// ToRNA composes RNA strand based on the DNA input
func ToRNA(dna string) string {

	rna := []byte{}

	for _, v := range dna {
		switch v {
		case 'A':
			rna = append(rna, 'U')
		case 'C':
			rna = append(rna, 'G')
		case 'G':
			rna = append(rna, 'C')
		case 'T':
			rna = append(rna, 'A')
		default:
			return ""
		}

	}
	return string(rna)
}
