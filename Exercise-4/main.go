// Exercise: Type Inference Gone Wrong

package main

import "math/rand"

func main() {
	//var seed = 1234456789       // what is wrong here
	var seed int64 = 1234456789 // rand.Seed requires a variable of the int64 type.
	rand.Seed(seed)
}
