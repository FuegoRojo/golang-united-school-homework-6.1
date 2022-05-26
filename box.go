package golang_united_school_homework

import (
	"errors"
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

	quantity := len(b.shapes)

	if quantity < b.shapesCapacity {
		b.shapes = append(b.shapes, shape)
		return nil
	} else {
		return errors.New("too much arguments")
	}

}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if b.shapesCapacity < i {
		return nil, errors.New("there are fewer figures in the box than the index")
	}
	if i >= len(b.shapes) || i < 0 {
		return nil, errors.New("the element does not exist")
	} else {
		return b.shapes[i], nil
	}

}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if b.shapesCapacity < i {
		return nil, errors.New("there are fewer figures in the box than the index")
	}
	if i >= len(b.shapes) || i < 0 {
		return nil, errors.New("the element does not exist")
	} else {
		result := b.shapes[i]
		b.shapes = append(b.shapes[0:i], b.shapes[i+1:]...)
		return result, nil
	}

}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if b.shapesCapacity < i {
		return nil, errors.New("there are fewer figures in the box than the index")
	}
	if i >= len(b.shapes) || i < 0 {
		return nil, errors.New("the element does not exist")
	} else {
		result := b.shapes[i]
		b.shapes[i] = shape
		return result, nil
	}

}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sumPerimeter float64
	for _, value := range b.shapes {
		sumPerimeter += value.CalcPerimeter()
	}
	return sumPerimeter
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sumArea float64
	for _, value := range b.shapes {
		sumArea += value.CalcArea()
	}
	return sumArea
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	var existCirkle bool
	var newList []Shape
	for _, value := range b.shapes {
		a, ok := value.(*Circle)
		if ok {
			existCirkle = true
		} else {
			newList = append(newList, a)
		}

	}
	if existCirkle {
		b.shapes = newList
		return nil
	} else {
		return errors.New(" circles are not exist in the list")
	}
}
