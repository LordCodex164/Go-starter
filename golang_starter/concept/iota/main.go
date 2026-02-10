package main

import "fmt"

//it allows automatically numbering
const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thurday
	Friday
	Saturday
	Sunday
)

//we can use iota to represent bitflags
//bitflags are a compact way of representing multiple boolean options using individual bits within a single numeric value
//allowng efficient storage and manipulation of multiple flags or settings.
const (
	Readable = 1 << iota
)

func main () {
	fmt.Println(">>Readable", Readable)
}
