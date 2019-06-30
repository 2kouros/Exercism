package space

type Planet string

//Age calculates the age in different planets for a given input in Earth seconds.
func Age(seconds float64, planet Planet) float64 {

	const (
		earthYearSec = 31557600
		MercuryT     = 0.2408467
		VenusT       = 0.61519726
		MarsT        = 1.8808158
		JupiterT     = 11.862615
		SaturnT      = 29.447498
		UranusT      = 84.016846
		NeptuneT     = 164.79132
	)
	var actual float64
	switch planet {

	case "Earth":
		actual = seconds / earthYearSec
	case "Mercury":
		actual = seconds / (earthYearSec * MercuryT)
	case "Venus":
		actual = seconds / (earthYearSec * VenusT)
	case "Mars":
		actual = seconds / (earthYearSec * MarsT)
	case "Jupiter":
		actual = seconds / (earthYearSec * JupiterT)
	case "Saturn":
		actual = seconds / (earthYearSec * SaturnT)
	case "Uranus":
		actual = seconds / (earthYearSec * UranusT)
	case "Neptune":
		actual = seconds / (earthYearSec * NeptuneT)

	}
	return actual
}
