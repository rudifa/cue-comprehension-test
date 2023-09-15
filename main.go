/*
Copyright © 2023 Rudolf Farkas @rudifa rudi.farkas@gmail.com
*/
package main

import (
	"github.com/rudifa/cue-comprehension-test/cmd"
)

func main() {

	cmd.Execute()
}

// package main

// import (
//     "flag"
//     "fmt"
// )

// func main() {
//     // Define a flag called "size" with a default value of 10 and a short description
//     sizePtr := flag.Int("size", 10, "Size of test data to generate")

//     // Parse the command-line arguments to set the value of the "size" flag
//     flag.Parse()

//     // Get the value of the "size" flag
//     size := *sizePtr

//     fmt.Printf("Generating test data of size %d\n", size)
// }
