package geometry

type Rectangle struct {
	Height float64
	Width  float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width //Here r is similar to this keyword
}

func (c Circle) Area() float64 {
	return 314.15
}
