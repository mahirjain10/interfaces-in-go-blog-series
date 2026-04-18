package main

import "fmt"

// =============================================================================
// INTERFACES
// =============================================================================

// Shape interface has a "verify" method
// Different shapes can define how they want to verify if an object satisfies the condition to be of that particular shape
type Shape interface {
	verify() string
	// PLEASE KEEP IT COMMENTED (REFERENCED GIVEN AT BOTTOM OF THE CODE)
	// exampleFunc(param1 int, param2 string) bool
}

// =============================================================================
// TYPES
// =============================================================================

// Circle with radius
// It defines the property of a circle, that it should have a radius
type circle struct {
	radius *float32
}

// Unknown shape
type unknown struct {
	value *float32
}

// Triangle with sides
// It defines the properties of a triangle, i.e. three sides
type triangle struct {
	sideA, sideB, sideC *float32
}

// =============================================================================
// IMPLEMENTATIONS
// =============================================================================

// Triangle wants an object to have three sides to let it pass through the cutout
// "An object should have three sides if it wants to be called a triangular object"
func (t triangle) verify() string {
	if t.sideA == nil || t.sideB == nil || t.sideC == nil {
		return "It is not a triangle"
	}
	return "It is a triangle"
}

// Circle wants an object to have a radius to let it pass through the cutout
// "An object should have a radius if it wants to be called a circular object"
func (c circle) verify() string {
	if c.radius == nil {
		return "It is not a circle"
	}
	return "It is a circle"
}

// =============================================================================
// FUNCTIONS
// =============================================================================

// Result takes Shape as a parameter
// Those structs that have implemented the methods defined by the interface can satisfy this parameter requirement
// If "exampleFunc" is uncommented and not implemented by our concrete structs (circle and triangle),
// those variables would face a compile-time error
func result(s Shape) {
	fmt.Println(s.verify())
}

// This is called type switch assertion
func assert(i any) {
	switch v := i.(type) {
	case string:
		updatedString := fmt.Sprintf("%s %s", v, "updated")
		fmt.Printf("%v is a string\n", updatedString)

	// You can only access fields of circle if the underlying type is a circle
	case circle:
		fmt.Printf("%v underlying concrete type is circle\n", *v.radius)

	case Shape:
		fmt.Printf("%T implements the Shape interface\n", v)

	default:
		fmt.Println("unknown type")
	}
}

// =============================================================================
// MAIN
// =============================================================================

func main() {
	r := float32(10.0)
	a := float32(9.5)
	b := float32(9.5)
	cVal := float32(8.0)

	tennisBall := circle{radius: &r}

	// Calling it an "icecreamCone" object just for the sake of this example
	// Please ignore a few properties of geometrical shapes
	icecreamCone := triangle{sideA: &a, sideB: &b, sideC: &cVal}

	// randomShape := unknown{value: &cVal}

	// ---------------------------------------------------------------------
	// USING INTERFACE
	// ---------------------------------------------------------------------

	result(tennisBall)
	result(icecreamCone)

	// ---------------------------------------------------------------------
	// COMPILE-TIME CHECKS
	// ---------------------------------------------------------------------

	// No compile-time error as the concrete type has implemented the Shape interface
	var _ Shape = tennisBall
	var _ Shape = icecreamCone

	// Uncommenting randomShape would throw a compile-time error
	// cannot use randomShape (variable of struct type unknown) as Shape value: unknown does not implement Shape (missing method verify)
	// var _ Shape = randomShape
	// result(randomShape)

	// ---------------------------------------------------------------------
	// TYPE ASSERTION
	// ---------------------------------------------------------------------

	var i interface{} = tennisBall
	// var i any = tennisBall // you can also use "any"

	// UNSAFE - will panic at runtime as i holds a circle, not a triangle
	// panic: interface conversion: interface {} is main.circle, not main.triangle
	// tp := i.(triangle)
	// fmt.Println("triangle side is ", tp.sideA)

	// SAFE - won't panic as we check ok before accessing
	if c, ok := i.(circle); ok {
		fmt.Printf("circle: %T\n", c)
	}

	// ---------------------------------------------------------------------
	// COMPILE-TIME CONCRETE TYPE CHECK
	// ---------------------------------------------------------------------

	var _ circle = tennisBall
	// var _ circle = icecreamCone // ERROR: cannot use icecreamCone (variable of struct type triangle) as circle value (compiler IncompatibleAssign)

	// ---------------------------------------------------------------------
	// INVALID ASSERTION EXAMPLE
	// ---------------------------------------------------------------------

	// Uncommenting below would throw a compile-time error
	// invalid operation: tennisBall (variable of struct type circle) is not an interface
	// if str, ok := tennisBall.(*circle); ok {
	// }

	// ---------------------------------------------------------------------
	// RUNNING ASSERTIONS
	// ---------------------------------------------------------------------

	assert("hi")
	assert(tennisBall)
	assert(icecreamCone)
}
