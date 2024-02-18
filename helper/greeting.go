package helper

import "fmt"

func SayHello(name string) (value string) {
	value = fmt.Sprintf("Hello %s", name)
	return
}
