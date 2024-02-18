package main

import (
	"fmt"
	"strings"
)

func main() {
	data := NewMap("")
	if data != nil {
		fmt.Println(data)
	} else {
		fmt.Println("Data is null/nil")
	}
}

// NewMap -> the data types those can be null/nil
// are interface, function, map, slice, pointer and channel
func NewMap(name string) (newMap map[string]string) {
	if v := strings.Trim(name, ""); len(v) > 0 {
		newMap = map[string]string{
			"name": name,
		}
	}
	return newMap
}
