package main

/* Packages

Every Go program is made up of packages.

Programs start running in package main.


		Go Syntax (It is more readable)

			x: int
			p: pointer to int
			a: array[3] of int

			For the the sake of Brevity
			x int
			p *int
			a [3]int

		Go's basic types are

			bool  (true, false)

			string ("Hello World")

			int  int8  int16  int32  int64 (Whole numbers)

			uint uint8 uint16 uint32 uint64 uintptr (+ve whole numbers including zero)

			float32 float64 (Real numbers)

			Other Data types:
				byte

				complex64 complex128
*/

// importing packages
import (
	"fmt"  // for printing
	"math" // for maths operations
	"time" // for date & time
)

func main() {
	fmt.Println("************** Welcome to Distributed Systems ********************")

	fmt.Printf("\n\n> ========== Unit 1 : Different ways to declare & initialize variables ========= \n\n")
	var a0 int                                                               // declaring a variable
	a0 = 100                                                                 //initializing a variable
	var a1, a2, a3, a4, a5 = 1, 2, 3.0, "(String) := Short Assignment", true // compiler will infer the type of each variable
	a1 = a0
	// or
	var b1, b2 int
	var b3 float64
	b1, b2 = a1, a2
	b3 = a3
	b4 := a4 // Short way of variable declartion (& initializing) if we use ":=" then we can skip "var"
	// c1 := 1.1
	// c2 := "Hello"
	const g = 9.89
	var b5 = a5 // Again the compiler will infer the type

	fmt.Println("Different ways of variable declaration and initialization:", b1, b2, b3, b4, b5)

	var (
		test1 bool
		test2 string
		test3 int
		test4 byte
		test5 float32
		test6 float64
		test7     = "This is String"
		test8 int = 64
	)

	fmt.Println(">Default values of basic data types: bool, string, int, byte, float32, and float64, respectively", test1, test2, test3, test4, test5, test6)
	fmt.Println("> ", test7, test8)

	var a float64 = 3.7656
	var b int = int(a) // type conversion -> convert float to integer

	fmt.Println(">Type converstion test:", b)

	fmt.Println("\n>Convert float or int to string", fmt.Sprint(3.14), fmt.Sprint(3))

	// ******************************* Unit 2: ***************************************************
	fmt.Printf("\n\n>========== Unit 2: Loops, if-else, Switch Statments ========= \n\n")

	for i := 0; i < 5; i++ { // for loop
		fmt.Println(i)
	}
	fmt.Println("> i cannot be accessed outside of for loop")

	// var i = -1
	// for i = 0; i < 10; i++ {
	// 	fmt.Println("2nd way of for loop", i)
	// }

	sum := 1 // for is also while in Go
	for sum < 5 {
		sum += sum
		fmt.Println(sum)
	}

	if sum > 5 { // if condition
		fmt.Println(">Sum is greater than 5")
	}

	sum = 100

	// fmt.Println(math.Sqrt(sum))

	if temp := math.Sqrt(float64(sum)); temp > float64(sum) { // if-else
		fmt.Println(">Scope: temp cannot be accessed outside of if {}")
	} else if temp == 0 {
		fmt.Println(">Scope: temp can also be accessed in else if block {}", temp)
	} else {
		fmt.Println(">Scope: temp can also be accessed in else block {}", temp)
	}

	// swtich statement
	fmt.Println("When is Saturday?")

	today := time.Now().Weekday()
	fmt.Printf("Today is %s\n", today)

	switch time.Saturday {
	case today + 0: // doesnt has to be a constant over here i.e., you can evaluate an expression infront of "case"
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In 2 days")
	case today + 3:
		fmt.Println("In 3 days")
	default:
		fmt.Println("Do not know.")
	}

	// Switch without a condition => more cleaner way to write longer if else statmetns
	switch {
	case time.Now().Hour() < 12:
		fmt.Println("Good Morning")
	case time.Now().Hour() < 17:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good evening")
	}

	// ******************************* Unit 3: ***************************************************
	fmt.Printf("\n\n>==========Unit 3: Pointers & Structs ========= \n\n")
	a = 3.14 // float
	b = 3    // int

	// pointers
	var p1 *float64 // declaring a float pointer. its zero value is nil.
	p1 = &a         // point to a
	p2 := &b        // point to b i.e., combined the declartion and initializing step

	*p1 = 99.0909
	*p2 = 888

	fmt.Println("Pointers:", a, b, *p1, *p2, p1, p2)

	// Structs
	point1 := Point2D{0.0, 2.2}

	point1.x = 100.788 // use dot operator to access the field
	fmt.Println(">Structs:", point1)

	p3 := &point1
	(*p3).y = 111.11
	p3.x = 200.22 // go allow to skip explicit defrence just to make notation clean

	fmt.Println(">Structs", point1)

	// other ways to declare struct
	var (
		point2 = Point2D{1.1, 3.3}
		p4     = &Point2D{2.2, 4.4} // same as p3 but here it is a refrence to struct
	)

	fmt.Println("Ways to initialize structs:", point1, point2, p3, p4)

	// ******************************* Unit 4: ***************************************************
	fmt.Printf("\n\n> ==========Unit 4: Array ========= \n\n")
	var array1 [2]string // you cannot resize an array when it is declared
	array1[0] = "Distributed"
	array1[1] = "Systems"

	array2 := [4]float64{0.0, 1.1, 2.2, 3.3} // another way of declartion
	array2[0] = 88

	fmt.Println(">Array ", array1, array2)

	// ******************************* Unit 5: ***************************************************
	fmt.Printf("\n\n> ==========Unit 5: Slices ========= \n\n")
	/*
		A dynamically Sized array.
	*/
	slice1 := []int{1, 2, 3} // slices are declared like arrays but without length
	l := len(slice1)         // number of elems contain in the slice
	c := cap(slice1)         // capacity of the array. capacity >= length
	fmt.Printf(">Slice 1: %v, lenght = %d, capacity = %d \n", slice1, l, c)

	length := 0
	capacity := 20
	slice2 := make([]string, length, capacity) // another way to initialize slices
	fmt.Printf("Slice2 = %v, len = %d, c= %d\n", slice2, len(slice2), cap(slice2))
	slice2 = append(slice2, "First elem")                // return a slice which contain all the elems
	slice2 = append(slice2, "Second elem", "Third elem") // insert multiples entries
	fmt.Printf("Slice2 (inserting elements) = %v, len = %d, c= %d\n", slice2, len(slice2), cap(slice2))

	var slice3 []int // another way of declaring slice
	fmt.Printf("Slice3 = %v, len = %d, c= %d\n", slice3, len(slice3), cap(slice3))
	slice3 = append(slice3, 2, 3, 5, 7) // append will increase the capacity of the array automatically
	fmt.Printf("Slice3 = %v, len = %d, c= %d\n", slice3, len(slice3), cap(slice3))

	// ******************************* Unit 6: ***************************************************
	fmt.Printf("\n\n> ========== Unit 6: Maps (Dictionary or key-value data structure) ========= \n\n")

	/*
		To create an empty map, use the builtin make: make(map[key-type]val-type).
	*/

	var student2courses map[string][]string     // nil map i.e. contain nothing
	student2courses = make(map[string][]string) // initialize the map
	//prof2courses := make(map[string][]string)

	// inserting keys
	student2courses["student1"] = []string{"Data-structure", "computer-networks", "Security"}
	student2courses["student2"] = []string{"Algorithms", "Distributed-Systems", "Machine-Learning"}
	student2courses["student3"] = append(student2courses["student3"], "Computer-Vision", "Deep-Learning")

	fmt.Printf("Student2Courses %v", student2courses)

	// delete a key
	delete(student2courses, "student1") // to delete a key and its value from map
	fmt.Println("After deletion of key", student2courses)
	value1, ok1 := student2courses["student1"] // to check wehter key exists or not
	_, ok2 := student2courses["student2"]
	fmt.Println(ok1, value1, ok2)

	// ******************************* Unit 7: ***************************************************
	fmt.Printf("\n\n> ========== Unit 7: Functions ========= \n\n")

	fmt.Println("Addition Function: add1() & add2 () : ", add1(3, 4), add2(3, 4))
	v1, _, v3 := returnMultipleValues1(4, 6, "hello")
	fmt.Println("Returning Multiple Values 1", v1, v3)

	// ************************* Advance Level *****************************************
	// Function can also be used as values (e.g., arguments to a function, return values)
	hypot := func(x, y float64) float64 { // in python or java such functions are called lambda functions
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(3.0, 2.0))
	fmt.Println(computeFun(3.0, 2.0, hypot))

	// //Advance: clousure: Go functions may be closures. A clousre is a function value.
	// adder1 := closureExampleFunction() // adder1 is clousre function
	// adder2 := closureExampleFunction() // adder2 is now a different clousre function. i.e., adder1 and adder2 functions have their own separate independent variable
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(adder1(i), adder2(-i))
	// }

	// ******************************* Unit 8: ***************************************************
	fmt.Printf("\n\n> ========== Unit 8: Methods ========= \n\n")
	/*
		Go doesnt have classes but you can link methods to types.
		A method is a function with a special receiver argument.
		The receiver appears in its own argument list between the func keyword and the method name.
		In this example, the computeDistance method has a receiver of type Point2D named p1.
	*/
	p5 := Point2D{20, 20}
	p6 := Point2D{10, 10}
	f := MyFloat(3.9999)

	fmt.Printf("Calling Methods:  %g, %g, %d \n", p5.computeDistance(p6), f, f.convertToInteger())

	p5.scaleByTen()

	fmt.Println(p5)

	// ******************************* Unit 9: ***************************************************
	fmt.Printf("\n\n> ========== Unit 9: Interfaces ========= \n\n")

	/*
		In Go, an interface is a type which only contain "methods signatures".
	*/

	rec := MyRect{width: 3, height: 4}
	cir := MyCircle{radius: 5}

	measure(rec)
	measure(cir)

}

func add1(x int, y int) int {
	return x + y
}

// if consective variables share the same data type, you can skip all but the last
func add2(x, y int) int {
	return x + y
}

// ************************* Returning Values from function ***************************
func returnMultipleValues1(x int, y int, s string) (int, int, string) {

	return x, y, s
}

func returnMultipleValues2(x int, y int, s string) (v1 int, v2 int, v3 string) {

	v1 = x
	v2 = y
	v3 = s

	return // This return is called "naked" return.
}

func computeFun(x, y float64, f func(float64, float64) float64) float64 {
	return f(x, y)
}

func closureExampleFunction() func(int) int { // compute the sum
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// struct declartion
type Point2D struct {
	x float64
	y float64
}

// declaring a method on a data type

func (p1 Point2D) computeDistance(p2 Point2D) float64 { // p1 is the reciever

	return math.Sqrt((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y))
}

// declaring a method on built in data type
type MyFloat float64

func (f MyFloat) convertToInteger() int {
	return int(f)
}

// remove pointer sign and see its effect
func (p *Point2D) scaleByTen() {
	p.x = p.x + 10
	p.y = p.y + 10
}

// interfaces
type Geometry interface {
	area() float64
	perim() float64
}

type MyRect struct {
	width, height float64
}

type MyCircle struct {
	radius float64
}

func (r MyRect) area() float64 {
	return r.width * r.height
}
func (r MyRect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c MyCircle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c MyCircle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g Geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
