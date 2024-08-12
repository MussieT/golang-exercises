package interfaces

// Note: Good demo - how authorizer uses an interface for different providers, you just call provider.Users (the provider can be postgres, mongodb, etc..)

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
    return 2*r.Width + 2*r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * 3.14 * c.Radius
}

func measure(s Shape) {
    fmt.Println(s)
    fmt.Println(s.Area())
    fmt.Println(s.Perimeter())
}

func InterfaceFunctions () {
    rectangle := Rectangle{Width: 4, Height: 5}
    circle := Circle{Radius: 3}

	measure(rectangle)
	measure(circle)
}

func TypeCorecions(any interface{}) {
	fmt.Printf("%T\n", any)
}