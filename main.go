package main

import "fmt"

func main() {

	dataToBeSorted := []int{20, 35, -15, -15, 7, 55, 1, -22}
	fmt.Println(dataToBeSorted)

	for lastUnSortedIndex := len(dataToBeSorted) - 1; lastUnSortedIndex > 0; lastUnSortedIndex-- {
		for i := 0; i < lastUnSortedIndex; i++ {
			if dataToBeSorted[i] > dataToBeSorted[i+1] {
				swap(dataToBeSorted, i, i+1)
			}
		}
	}
	fmt.Println(dataToBeSorted)

}

//swap function to swap the array values if greater than next
func swap(sortData []int, num int, numk int) {

	if num == numk {
		return
	}
	temp := sortData[num]
	sortData[num] = sortData[numk]
	sortData[numk] = temp

}
