package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var intVar int = 10
	// Star (*) is used to declare a pointer
	// Dereference operator (&) is used to get the address of a variable
	var addressOfIntVar *int = &intVar
	valueOfAddressOfIntVar := *addressOfIntVar

	fmt.Printf("intVar: %d\naddressOfIntVar: %p\nvalueOfAddressOfIntVar: %d\n", intVar, addressOfIntVar, valueOfAddressOfIntVar)

	// Change the value of intVar
	*addressOfIntVar = 20
	fmt.Printf("Changed intVar: %d\n", intVar)

	// Perform pointer arithmetic using unsafe package
	unsafeAddress := unsafe.Pointer(addressOfIntVar)
	unsafeAddress = unsafe.Pointer(uintptr(unsafeAddress) + unsafe.Sizeof(intVar))
	newAddressOfIntVar := (*int)(unsafeAddress)
	fmt.Printf("New addressOfIntVar: %p\n", newAddressOfIntVar)
	fmt.Printf("Value at new addressOfIntVar: %d\n", *newAddressOfIntVar)
}