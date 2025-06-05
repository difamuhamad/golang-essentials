package main

import "fmt"

func main() {
	// normal array :
	var months = [...]string{
		"Januari", "Februari", "Maret", "April", "Mei", "Juni",
		"Juli", "Agustus", "September", "Oktober", "November", "Desember",
	}

	// slices :
	var slice []string = months[:]

	var slice1 = months[9:]
	var slice2 = months[:3]

	var spring = months[2:5]
	var summer = months[5:8]
	var autumn = months[8:11]
	var winter = months[11:12]
	winter = append(winter, months[0:2]...)

	fmt.Println("Slice :", slice)
	fmt.Println("Slice 1 :", slice1)
	fmt.Println("Slice 2 :", slice2)
	fmt.Println("Winter:", winter)
	fmt.Println("Spring:", spring)
	fmt.Println("Summer:", summer)
	fmt.Println("Autumn:", autumn)

	// ----------------------------  Pointer
	var weeks = [...]string{"Monday", "Tuesday", "Wednesday", "Thursday",
		"Friday", "Saturday", "Sunday"}

	var weekSlice []string = weeks[:]
	fmt.Println("Weeks:", weeks)

	weekSlice[4] = "Liburrr"
	weekSlice[6] = "Liburrr"
	fmt.Println("Updated Weeks:", weeks)

	var newWeek = append(weekSlice, "Weekend Off", "Weekend Off")
	fmt.Println("New Weeks:", newWeek)
	fmt.Println("Weeks:", weeks)

	copy(weeks[:], newWeek[:])

	weeks[0] = "Start Week"
	weeks[3] = "Almost Weekend"

	fmt.Println("Updated Weeks:", weeks)

	// ---------------------------- Making Slice
	// make(slice type, length, capacity)
	var newSlice []string = make([]string, 2, 5)

	newSlice[0] = "First"
	newSlice[1] = "Second"
	// newSlice[2] = "Third"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Third")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	copiedSlice := newSlice[:]
	pastedSlice := make([]string, len(copiedSlice), cap(copiedSlice))

	copy(pastedSlice, copiedSlice)
	fmt.Println(copiedSlice)
	fmt.Println(pastedSlice)

	// will affect the array
	newSlice2[0] = "Empty"
	fmt.Println(newSlice)

	// ---------------------------- Different between array & slice

	arrayVar := [...]int{1, 2, 3, 4, 5}
	sliceVar := []int{1, 2, 3, 4, 5}

	fmt.Println("This is array :", arrayVar)
	fmt.Println("This is slice :", sliceVar)
}
