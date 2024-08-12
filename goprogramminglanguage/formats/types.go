package formats

type Celcius float64
type Farenheight float64

// Type coersion is must
func CtoF(c Celcius) Farenheight {
	return Farenheight(c * 9/5 + 32)
}

func FToC(f Farenheight) Celcius {
	return Celcius((f - 32) * 5 / 9)
}
