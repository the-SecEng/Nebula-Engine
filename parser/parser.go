package main

import (
	"fmt"

	"github.com/dop251/goja"
)

func main() {
	// Define the input program
	src := `
		console.log("The parser works correctly.");
		let x = 5 + 3;
		let y = x * 2;
		console.log(x, y)
	`

	// Create the VM and add the console object
	vm := goja.New()
	console := func(call goja.FunctionCall) goja.Value {
		for _, arg := range call.Arguments {
			fmt.Println(arg.String())
		}
		return goja.Undefined()
	}
	consoleObj := map[string]interface{}{
		"log": console,
	}
	vm.Set("console", consoleObj)

	// Parse the program and execute it
	val, err := vm.RunString(src)
	if err != nil {
		panic(err)
	}

	// Print out the result for debugging purposes
	fmt.Println(val)
}
