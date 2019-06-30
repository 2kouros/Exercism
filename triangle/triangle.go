//Package triangle contains the function responsible for distinguishing the kinds of triangles
package triangle

import (
	"math"
)

// The Kind of triangle
type Kind int

//Constants
const (
	NaT = iota //not a triangle
	Equ        //Equilateral
	Iso        //Isosceles
	Sca        //Scalene
)

// KindFromSides analyzes the size of the triangle's side and decides its kind
func KindFromSides(a, b, c float64) Kind {
	if !isTriangle(a, b, c) {
		return NaT
	} else if isEquilateral(a, b, c) {
		return Equ
	} else if isIsosceles(a, b, c) {
		return Iso
	} else if isScalene(a, b, c) {
		return Sca
	}
	return NaT
}

func isTriangle(a, b, c float64) bool {
	return a+b >= c && a+c >= b && c+b >= a && isLegit(a) && isLegit(b) && isLegit(c)

}
func isLegit(s float64) bool {
	return s > 0 && !math.IsNaN(s) && !math.IsInf(s, 1)
}
func isEquilateral(a, b, c float64) bool {
	return a == b && a == c && b == c

}
func isIsosceles(a, b, c float64) bool {
	return a == b || a == c || b == c

}
func isScalene(a, b, c float64) bool {
	return a != b && a != c && b != c

}
