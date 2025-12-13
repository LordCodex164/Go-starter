package main

func main () {
	truckID := 42

	anotherTruckID := &truckID

	println(&truckID)
	println(*anotherTruckID)
}
