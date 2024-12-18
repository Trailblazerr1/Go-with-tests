package geometry

type Rectangle struct {
	Height float64
	Width  float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	base   float64
	height float64
}

// declare an interface shape which has a method Area,
// which is similar to Rect, Circle and Triangle, since all of them have area
// method. But none other have area method
type Shape interface {
	Area() float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width //Here r is similar to this keyword
}

func (c Circle) Area() float64 {
	return 314.15
}

func (t Triangle) Area() float64 {
	return (t.base * t.height) * 0.5
}
