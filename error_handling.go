package main

import "fmt"

func main() {
	message, e := runApp(false)
	println()
	fmt.Println("error: ", e)
	fmt.Println("message: ", message)

	println()
	message, e = runApp(true)
	println()
	fmt.Println("error: ", e)
	fmt.Println("message: ", message)
}

func logging() {
	println("Complete")
}

func runApp(isError bool) (message any, error bool) {
	error = false
	message = "Successfully run app"
	println("Run Application")
	defer func() {
		var e = recover() // catch error
		if e != nil {
			error = true
			message = e
		} // this equal to catch error
	}()
	defer logging() // this equal to finally
	if isError {
		panic("there's something is missing out") // this equal to throw
	}
	return
}
