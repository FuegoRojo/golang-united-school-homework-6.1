package golang_united_school_homework

import (
	"fmt"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if b.shapesCapacity != len(b.shapes) {
		b.shapes = append(b.shapes, shape)
		return nil
	} else {
		return fmt.Errorf("it goes out of the shapesCapacity range")
	}
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i > len(b.shapes)-1 || b.shapes == nil {
		return nil, fmt.Errorf("shape by index doesn't exist")
	} else {
		return b.shapes[i], nil
	}

}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if i > len(b.shapes)-1 || b.shapes[i] == nil {
		return nil, fmt.Errorf("shape by index doesn't exist or index went out of the range")
	} else {
		shapesToReturn := b.shapes[i]
		b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
		return shapesToReturn, nil
	}

}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if i > len(b.shapes)-1 || b.shapes[i] == nil {
		return nil, fmt.Errorf("index went out of the range")
	} else {
		shapesToReturn := b.shapes[i]
		b.shapes[i] = shape
		return shapesToReturn, nil
	}

}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sum float64 = 0
	for _, v := range b.shapes {
		sum += v.CalcPerimeter()
	}
	return sum

}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sum float64 = 0
	for _, v := range b.shapes {
		sum += v.CalcArea()
	}
	return sum

}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	flagExists := false

	for i := 0; i < len(b.shapes); i++ {
		if IsShapeCircle(b.shapes[i]) {
			flagExists = true
			b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
			i--

		}
	}

	if !flagExists {
		return fmt.Errorf("error: there is no circle in the box")
	} else {
		return nil
	}

}
func IsShapeCircle(shape interface{}) bool {
	switch shape.(type) {
	case *Circle:
		return true
	default:
		return false
	}
}
