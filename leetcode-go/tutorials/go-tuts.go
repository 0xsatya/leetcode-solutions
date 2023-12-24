package tutorials

import (
	"fmt"
	"math"
	"time"
)

// `const` declares a constant value.
const s string = "constant"

func Main_tuts() {

	///-------------------------------------------------------------///
	/// 	        		USING PRINTS 	         				///
	///-------------------------------------------------------------///

	fmt.Println("go" + "lang")

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	fmt.Println(s)

	// A `const` statement can appear anywhere a `var` statement can.
	const n = 500000000

	// Constant expressions perform arithmetic with arbitrary precision.
	const d = 3e20 / n
	fmt.Println(d)

	// A numeric constant has no type until it's given
	// one, such as by an explicit conversion.
	fmt.Println(int64(d))

	// A number can be given a type by using it in a
	// context that requires one, such as a variable
	// assignment or function call. For example, here
	// `math.Sin` expects a `float64`.
	fmt.Println(math.Sin(n))
	///-------------------------------------------------------------///
	/// 	        		IF ELSE       	         				///
	///-------------------------------------------------------------///
	// Here's a basic example.
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	//Note that you don’t need parentheses around conditions in Go, but that the braces are required.
	// You can have an `if` statement without an else.
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// A statement can precede conditionals; any variables
	// declared in this statement are available in the current
	// and all subsequent branches.
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	///-------------------------------------------------------------///
	///      	        		For LOOP 	         				///
	///-------------------------------------------------------------///
	// FOR LOOP
	// The most basic type, with a single condition.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// A classic initial/condition/after `for` loop.
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// `for` without a condition will loop repeatedly until you `break` out of the loop or `return` from
	// the enclosing function.
	for {
		fmt.Println("loop")
		break
	}

	// You can also `continue` to the next iteration of	the loop.
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	///-------------------------------------------------------------///
	///      	        		SWITCH STATEMENTS      				///
	///-------------------------------------------------------------///
	// Here's a basic `switch`.
	{
		i := 2
		fmt.Print("Write ", i, " as ")
		switch i {
		case 1:
			fmt.Println("one")
		case 2:
			fmt.Println("two")
		case 3:
			fmt.Println("three")
		}

		// You can use commas to separate multiple expressions
		// in the same `case` statement. We use the optional
		// `default` case in this example as well.
		switch time.Now().Weekday() {
		case time.Saturday, time.Sunday:
			fmt.Println("It's the weekend")
		default:
			fmt.Println("It's a weekday")
		}

		// `switch` without an expression is an alternate way
		// to express if/else logic. Here we also show how the
		// `case` expressions can be non-constants.
		t := time.Now()
		switch {
		case t.Hour() < 12:
			fmt.Println("It's before noon")
		default:
			fmt.Println("It's after noon")
		}

		// A type `switch` compares types instead of values.  You
		// can use this to discover the type of an interface
		// value.  In this example, the variable `t` will have the
		// type corresponding to its clause.
		whatAmI := func(i interface{}) {
			switch t := i.(type) {
			case bool:
				fmt.Println("I'm a bool")
			case int:
				fmt.Println("I'm an int")
			default:
				fmt.Printf("Don't know type %T\n", t)
			}
		}
		whatAmI(true)
		whatAmI(1)
		whatAmI("hey")
	}

	///-------------------------------------------------------------///
	///      	        		ARRAYS			      				///
	///-------------------------------------------------------------///
	{
		// Here we create an array `a` that will hold exactly
		// 5 `int`s. The type of elements and length are both
		// part of the array's type. By default an array is
		// zero-valued, which for `int`s means `0`s.
		var a [5]int
		fmt.Println("emp:", a)

		// We can set a value at an index using the
		// `array[index] = value` syntax, and get a value with
		// `array[index]`.
		a[4] = 100
		fmt.Println("set:", a)
		fmt.Println("get:", a[4])

		// The builtin `len` returns the length of an array.
		fmt.Println("len:", len(a))

		// Use this syntax to declare and initialize an array
		// in one line.
		b := [5]int{1, 2, 3, 4, 5}
		fmt.Println("dcl:", b)

		// Array types are one-dimensional, but you can
		// compose types to build multi-dimensional data
		// structures.
		var twoD [2][3]int
		for i := 0; i < 2; i++ {
			for j := 0; j < 3; j++ {
				twoD[i][j] = i + j
			}
		}
		fmt.Println("2d: ", twoD)
	}
	///-------------------------------------------------------------///
	///      	        		SLICES			      				///
	///-------------------------------------------------------------///
	{
		// Unlike arrays, slices are typed only by the
		// elements they contain (not the number of elements).
		// An uninitialized slice equals to nil and has
		// length 0.
		var s []string
		fmt.Println("uninit:", s, s == nil, len(s) == 0)

		// To create an empty slice with non-zero length, use
		// the builtin `make`. Here we make a slice of
		// `string`s of length `3` (initially zero-valued).
		// By default a new slice's capacity is equal to its
		// length; if we know the slice is going to grow ahead
		// of time, it's possible to pass a capacity explicitly
		// as an additional parameter to `make`.
		s = make([]string, 3)
		fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

		// We can set and get just like with arrays.
		s[0] = "a"
		s[1] = "b"
		s[2] = "c"
		fmt.Println("set:", s)
		fmt.Println("get:", s[2])

		// `len` returns the length of the slice as expected.
		fmt.Println("len:", len(s))

		// In addition to these basic operations, slices
		// support several more that make them richer than
		// arrays. One is the builtin `append`, which
		// returns a slice containing one or more new values.
		// Note that we need to accept a return value from
		// `append` as we may get a new slice value.
		s = append(s, "d")
		s = append(s, "e", "f")
		fmt.Println("apd:", s)

		// Slices can also be `copy`'d. Here we create an
		// empty slice `c` of the same length as `s` and copy
		// into `c` from `s`.
		c := make([]string, len(s))
		copy(c, s)
		fmt.Println("cpy:", c)

		// Slices support a "slice" operator with the syntax
		// `slice[low:high]`. For example, this gets a slice
		// of the elements `s[2]`, `s[3]`, and `s[4]`.
		l := s[2:5]
		fmt.Println("sl1:", l)

		// This slices up to (but excluding) `s[5]`.
		l = s[:5]
		fmt.Println("sl2:", l)

		// And this slices up from (and including) `s[2]`.
		l = s[2:]
		fmt.Println("sl3:", l)

		// We can declare and initialize a variable for slice
		// in a single line as well.
		t := []string{"g", "h", "i"}
		fmt.Println("dcl:", t)

		// Slices can be composed into multi-dimensional data
		// structures. The length of the inner slices can
		// vary, unlike with multi-dimensional arrays.
		twoD := make([][]int, 3)
		for i := 0; i < 3; i++ {
			innerLen := i + 1
			twoD[i] = make([]int, innerLen)
			for j := 0; j < innerLen; j++ {
				twoD[i][j] = i + j
			}
		}
		fmt.Println("2d: ", twoD)
	}
	///-------------------------------------------------------------///
	///      	        		MAPS			      				///
	///-------------------------------------------------------------///
	{
		// To create an empty map, use the builtin `make`:
		// `make(map[key-type]val-type)`.
		m := make(map[string]int)

		// Set key/value pairs using typical `name[key] = val`
		// syntax.
		m["k1"] = 7
		m["k2"] = 13

		// Printing a map with e.g. `fmt.Println` will show all of
		// its key/value pairs.
		fmt.Println("map:", m)

		// Get a value for a key with `name[key]`.
		v1 := m["k1"]
		fmt.Println("v1:", v1)

		// If the key doesn't exist, the
		// [zero value](https://go.dev/ref/spec#The_zero_value) of the
		// value type is returned.
		v3 := m["k3"]
		fmt.Println("v3:", v3)

		// The builtin `len` returns the number of key/value
		// pairs when called on a map.
		fmt.Println("len:", len(m))

		// The builtin `delete` removes key/value pairs from
		// a map.
		delete(m, "k2")
		fmt.Println("map:", m)

		// The optional second return value when getting a
		// value from a map indicates if the key was present
		// in the map. This can be used to disambiguate
		// between missing keys and keys with zero values
		// like `0` or `""`. Here we didn't need the value
		// itself, so we ignored it with the _blank identifier_
		// `_`.
		_, prs := m["k2"]
		fmt.Println("prs:", prs)

		// You can also declare and initialize a new map in
		// the same line with this syntax.
		n := map[string]int{"foo": 1, "bar": 2}
		fmt.Println("map:", n)
	}
	///-------------------------------------------------------------///
	///      	        		RANGE			      				///
	///-------------------------------------------------------------///
	{
		// Here we use `range` to sum the numbers in a slice.
		// Arrays work like this too.
		nums := []int{2, 3, 4}
		sum := 0
		for _, num := range nums {
			sum += num
		}
		fmt.Println("sum:", sum)

		// `range` on arrays and slices provides both the
		// index and value for each entry. Above we didn't
		// need the index, so we ignored it with the
		// blank identifier `_`. Sometimes we actually want
		// the indexes though.
		for i, num := range nums {
			if num == 3 {
				fmt.Println("index:", i)
			}
		}

		// `range` on map iterates over key/value pairs.
		kvs := map[string]string{"a": "apple", "b": "banana"}
		for k, v := range kvs {
			fmt.Printf("%s -> %s\n", k, v)
		}

		// `range` can also iterate over just the keys of a map.
		for k := range kvs {
			fmt.Println("key:", k)
		}

		// `range` on strings iterates over Unicode code
		// points. The first value is the starting byte index
		// of the `rune` and the second the `rune` itself.
		// See [Strings and Runes](strings-and-runes) for more
		// details.
		for i, c := range "go" {
			fmt.Println(i, c)
		}
	}
	///-------------------------------------------------------------///
	///      	        		RANGE			      				///
	///-------------------------------------------------------------///
	{
		// Here we use `range` to sum the numbers in a slice.
		// Arrays work like this too.
		nums := []int{2, 3, 4}
		sum := 0
		for _, num := range nums {
			sum += num
		}
		fmt.Println("sum:", sum)

		// `range` on arrays and slices provides both the
		// index and value for each entry. Above we didn't
		// need the index, so we ignored it with the
		// blank identifier `_`. Sometimes we actually want
		// the indexes though.
		for i, num := range nums {
			if num == 3 {
				fmt.Println("index:", i)
			}
		}

		// `range` on map iterates over key/value pairs.
		kvs := map[string]string{"a": "apple", "b": "banana"}
		for k, v := range kvs {
			fmt.Printf("%s -> %s\n", k, v)
		}

		// `range` can also iterate over just the keys of a map.
		for k := range kvs {
			fmt.Println("key:", k)
		}

		// `range` on strings iterates over Unicode code
		// points. The first value is the starting byte index
		// of the `rune` and the second the `rune` itself.
		// See [Strings and Runes](strings-and-runes) for more
		// details.
		for i, c := range "go" {
			fmt.Println(i, c)
		}
	}
	///-------------------------------------------------------------///
	///      	        		FUNCTIONS			      				///
	///-------------------------------------------------------------///
	{
		// Call a function just as you'd expect, with
		// `name(args)`.
		res := plus(1, 2)
		fmt.Println("1+2 =", res)

		res = plusPlus(1, 2, 3)
		fmt.Println("1+2+3 =", res)

		// Here we use the 2 different return values from the
		// call with _multiple assignment_.
		a, b := vals()
		fmt.Println(a)
		fmt.Println(b)

		// If you only want a subset of the returned values,
		// use the blank identifier `_`.
		_, c := vals()
		fmt.Println(c)

		// Variadic functions can be called in the usual way
		// with individual arguments.
		sum(1, 2)
		sum(1, 2, 3)

		// If you already have multiple args in a slice,
		// apply them to a variadic function using
		// `func(slice...)` like this.
		nums := []int{1, 2, 3, 4}
		sum(nums...)
	}
	///-------------------------------------------------------------///
	///      	        		CLOSURES     	      				///
	///-------------------------------------------------------------///
	{
		// We call `intSeq`, assigning the result (a function)
		// to `nextInt`. This function value captures its
		// own `i` value, which will be updated each time
		// we call `nextInt`.
		nextInt := intSeq()

		// See the effect of the closure by calling `nextInt`
		// a few times.
		fmt.Println(nextInt())
		fmt.Println(nextInt())
		fmt.Println(nextInt())

		// To confirm that the state is unique to that
		// particular function, create and test a new one.
		newInts := intSeq()
		fmt.Println(newInts())
	}
	///-------------------------------------------------------------///
	///      	        		RECURSION			      				///
	///-------------------------------------------------------------///
	{
		fmt.Println(fact(7))

		// Closures can also be recursive, but this requires the
		// closure to be declared with a typed `var` explicitly
		// before it's defined.
		var fib func(n int) int

		fib = func(n int) int {
			if n < 2 {
				return n
			}

			// Since `fib` was previously declared in `main`, Go
			// knows which function to call with `fib` here.
			return fib(n-1) + fib(n-2)
		}

		fmt.Println(fib(7))
	}
}

// Here's a function that takes two `int`s and returns
// their sum as an `int`.
func plus(a int, b int) int {

	// Go requires explicit returns, i.e. it won't
	// automatically return the value of the last
	// expression.
	return a + b
}

// When you have multiple consecutive parameters of
// the same type, you may omit the type name for the
// like-typed parameters up to the final parameter that
// declares the type.
func plusPlus(a, b, c int) int {
	return a + b + c
}

// The `(int, int)` in this function signature shows that
// the function returns 2 `int`s.
func vals() (int, int) {
	return 3, 7
}

// Here's a function that will take an arbitrary number
// of `int`s as arguments.
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	// Within the function, the type of `nums` is
	// equivalent to `[]int`. We can call `len(nums)`,
	// iterate over it with `range`, etc.
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// This function `intSeq` returns another function, which
// we define anonymously in the body of `intSeq`. The
// returned function _closes over_ the variable `i` to
// form a closure.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// This `fact` function calls itself until it reaches the
// base case of `fact(0)`.
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

///-------------------------------------------------------------///
///      	        		POINTERS		      				///
///-------------------------------------------------------------///

// We'll show how pointers work in contrast to values with
// 2 functions: `zeroval` and `zeroptr`. `zeroval` has an
// `int` parameter, so arguments will be passed to it by
// value. `zeroval` will get a copy of `ival` distinct
// from the one in the calling function.
func zeroval(ival int) {
	ival = 0
}

// `zeroptr` in contrast has an `*int` parameter, meaning
// that it takes an `int` pointer. The `*iptr` code in the
// function body then _dereferences_ the pointer from its
// memory address to the current value at that address.
// Assigning a value to a dereferenced pointer changes the
// value at the referenced address.
func zeroptr(iptr *int) {
	*iptr = 0
}

func Pointer_Example() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// The `&i` syntax gives the memory address of `i`,
	// i.e. a pointer to `i`.
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// Pointers can be printed too.
	fmt.Println("pointer:", &i)
}

///-------------------------------------------------------------///
///      	        		STRUCTS			      				///
///-------------------------------------------------------------///
// This `person` struct type has `name` and `age` fields.
type person struct {
	name string
	age  int
}

// `newPerson` constructs a new person struct with the given name.
func newPerson(name string) *person {
	// You can safely return a pointer to local variable
	// as a local variable will survive the scope of the function.
	p := person{name: name}
	p.age = 42
	return &p
}

func Structs_Examples() {
	// This syntax creates a new struct.
	fmt.Println(person{"Bob", 20})

	// You can name the fields when initializing a struct.
	fmt.Println(person{name: "Alice", age: 30})

	// Omitted fields will be zero-valued.
	fmt.Println(person{name: "Fred"})

	// An `&` prefix yields a pointer to the struct.
	fmt.Println(&person{name: "Ann", age: 40})

	// It's idiomatic to encapsulate new struct creation in constructor functions
	fmt.Println(newPerson("Jon"))

	// Access struct fields with a dot.
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// You can also use dots with struct pointers - the
	// pointers are automatically dereferenced.
	sp := &s
	fmt.Println(sp.age)

	// Structs are mutable.
	sp.age = 51
	fmt.Println(sp.age)

	// If a struct type is only used for a single value, we don't
	// have to give it a name. The value can have an anonymous
	// struct type. This technique is commonly used for
	// [table-driven tests](testing-and-benchmarking).
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)

}

type rect struct {
	width, height int
}

// This `area` method has a _receiver type_ of `*rect`.
func (r *rect) area() int {
	return r.width * r.height
}

// Methods can be defined for either pointer or value
// receiver types. Here's an example of a value receiver.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func Struct_Methods_Example() {
	r := rect{width: 10, height: 5}

	// Here we call the 2 methods defined for our struct.
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// Go automatically handles conversion between values
	// and pointers for method calls. You may want to use
	// a pointer receiver type to avoid copying on method
	// calls or to allow the method to mutate the
	// receiving struct.
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}

///-------------------------------------------------------------///
///      	        		ARRAYS			      				///
///-------------------------------------------------------------///
