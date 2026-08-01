package complexnumbers

import "math"

// Define the Number type here.
type Number struct {
    a float64
    b float64 
}

func (n Number) Real() float64 {
	return n.a
}

func (n Number) Imaginary() float64 {
	return n.b
}

func (n1 Number) Add(n2 Number) Number {
	newA := n1.Real() + n2.Real()
    newB := n1.Imaginary() + n2.Imaginary()
    return Number{a: newA, b: newB}
}

func (n1 Number) Subtract(n2 Number) Number {
	newA := n1.Real() - n2.Real()
    newB := n1.Imaginary() - n2.Imaginary()
    return Number{a: newA, b: newB}
}

func (n1 Number) Multiply(n2 Number) Number {
    newA := n1.Real() * n2.Real() - n1.Imaginary() * n2.Imaginary()
    newB := n1.Imaginary() * n2.Real() + n1.Real() * n2.Imaginary()
	return Number{a: newA, b: newB}
}

func (n Number) Times(factor float64) Number {
	newA := n.Real() * factor
    newB := n.Imaginary() * factor
    return Number{a: newA, b: newB}
}

func (n1 Number) Divide(n2 Number) Number {
	newA := (n1.Real() * n2.Real() + n1.Imaginary() * n2.Imaginary()) / (n2.Real() * n2.Real() + n2.Imaginary() * n2.Imaginary())
    newB := (n1.Imaginary() * n2.Real() - n1.Real() * n2.Imaginary()) / (n2.Real() * n2.Real() + n2.Imaginary() * n2.Imaginary())
    return Number{a: newA, b: newB}
}

func (n Number) Conjugate() Number {
	return Number{a: n.Real(), b: n.Imaginary() * (-1.0)}
}

func (n Number) Abs() float64 {
    return math.Sqrt(math.Pow(n.Real(), 2) + math.Pow(n.Imaginary(), 2))
	
}

func (n Number) Exp() Number {
	newA := math.Exp(n.Real()) * math.Cos(n.Imaginary())
    newB := math.Exp(n.Real()) * math.Sin(n.Imaginary())
    return Number{a: newA, b: newB}
}
