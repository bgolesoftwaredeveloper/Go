package main

import (
	"fmt"

	boyer_moore "github.com/bgolsoftwaredeveloper/boyer_moore/BoyerMooreImplementation"
)

func main() {
	var text string = "XYZXYXZYXYZXYYYXYZXYZZYZX"
	var pattern string = "XYZ"

	var indices []int = boyer_moore.BoyerMooreSearch(text, pattern)

	fmt.Printf("Pattern '%s' found at indices: %v\n", pattern, indices)
}
