package main

import (
	"fmt"
	"reflect"
)

func main() {
	names := [...]string{
		"Fajra",
		"Risqulla",
		"Cisnux",
		"Yagami",
		"Nanami",
		"Jack",
	}
	fmt.Println(names)
	slice := names[4:6]
	//slice[len(slice)-1] = "10"
	fmt.Println(slice) // [Nanami 10]
	fmt.Println(names) // [Fajra Risqulla Cisnux Yagami Nanami 10]

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jum'at", "Sabtu", "Minggu"}
	daySlice1 := days[5:]
	daySlice1[0] = days[5] + " Baru"
	daySlice1[1] = days[6] + " Baru"
	fmt.Println(daySlice1)
	fmt.Println(days) // [Senin Selasa Rabu Kamis Jum'at Sabtu Baru Minggu Baru]

	daySlice2 := append(daySlice1, names[:]...)
	daySlice2[0] = "Ups"
	fmt.Println(daySlice2) // [Ups Minggu Baru Fajra Risqulla Cisnux Yagami Nanami Jack]
	fmt.Println(days)      // [Senin Selasa Rabu Kamis Jum'at Sabtu Baru Minggu Baru]

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Fajra"
	newSlice[1] = "Risqulla"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)
	fmt.Println(toSlice)

	array1 := [...]int{1, 2, 3, 4, 5}
	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Println(array1)
	fmt.Println(reflect.TypeOf(array1))
	fmt.Println(slice2)
	fmt.Println(reflect.TypeOf(slice2))
}
