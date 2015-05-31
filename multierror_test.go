package multierror // import "github.com/chronos-tachyon/go-multierror"

import "fmt"

type exampleError struct {
	Name string
}

func (err exampleError) Error() string {
	return err.Name
}

var (
	errorA = exampleError{"A"}
	errorB = exampleError{"B"}
	errorC = exampleError{"C"}
)

func ExampleMultipleErrors_empty() {
	fmt.Printf("%T: %[1]v", Of())
	// Output:
	// <nil>: <nil>
}

func ExampleMultipleErrors_nil() {
	fmt.Printf("%T: %[1]v", Of(nil, nil, nil))
	// Output:
	// <nil>: <nil>
}

func ExampleMultipleErrors_one() {
	fmt.Printf("%T: %[1]v", Of(errorA))
	// Output:
	// multierror.exampleError: A
}

func ExampleMultipleErrors_many() {
	fmt.Printf("%T: %[1]v", Of(errorA, errorB, errorC))
	// Output:
	// multierror.MultipleErrors: encountered multiple errors:
	//	... A
	//	... B
	//	... C
}
