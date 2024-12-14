package examples

import (
	"fmt"
	// "time"
	"math"
	"math/cmplx"
	"math/rand"
)

// The var statement declares a list of variables; as in function argument lists, the type is last.
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// Constants declared with const keyword
const Pi = 3.14


// Functions can take 0 or more args. Types come after variable name
func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

// Go's return values can be named and a return statement without arguments returns the named valued
func split(sum int) (x, y int){
	x = sum * 4 / 9
	y = sum - x
	return
}

func loop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// init and post statements are optional
	newSum := 1
	for ; newSum < 1000; {
		newSum += newSum
	}
	fmt.Println(newSum)

	// While loop is just "For"
	newSum = 1
	for newSum < 1000 {
		newSum += newSum
	}
	fmt.Println(newSum)

}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	// Variables declared by statement are only in scope until end of the if
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func Sqrt(x float64) float64{
	z := 1.0

	for i := 0; i <= 30; i++{
		if z == x {
			break;
		} 
		z -= (z*z - x) / (2*z)
	}

	return z
}

func PartOne() {
	// A var statement can be at package or function level. We see both in this example.
	// If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
	
	// var i, j int = 1, 2
	
	// Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
    // Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
	
	// var i, j = 1, 2
	// i, j := 1, 2

	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	// fmt.Println(math.pi) # Exported names have a capital letter so this doesn't work
	fmt.Println(math.Pi)
	fmt.Println(add(42, 13))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// Type Conversion
	// Unlike in C, in Go assignment between items of different type requires an explicit conversion.
	i := 42
	f := float64(i)
	u := uint(f)
	fmt.Println(i, f, u)

	// Type Inference
	// When declaring a variable without specifying an explicit type 
	// (either by using the := syntax or var = expression syntax), 
	// the variable's type is inferred from the value on the right hand side.
	v := 0.867 + 0.5i // change me!
	fmt.Printf("v is of type %T\n", v)
	loop()

	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	
	// A defer statement defers the execution of a function until the surrounding function returns.
 	// The deferred call's arguments are evaluated immediately, but the function call is not executed until 
	// the surrounding function returns.
	// Deferred function calls are pushed onto a stack. 
	// When a function returns, its deferred calls are executed in last-in-first-out order.
	defer fmt.Println("world")
	fmt.Println(Sqrt(9))
}