package main

import "fmt"

func main() {
	var nums []int
	fmt.Println(nums, len(nums), cap(nums))

	nums = make([]int, 5)
	fmt.Println(nums, len(nums), cap(nums))

	citys := []string{"São Paulo"}
	fmt.Println(citys, len(citys), cap(citys))
	citys = append(citys, "Sweto")
	fmt.Println(citys, len(citys), cap(citys))

	singers := make([]string, 5)
	singers[0] = "Drake"
	singers[1] = "Emicida"
	singers[2] = "Mano Brown"
	singers[3] = "Ice Blue"
	singers[4] = "Michael Jackson"
	fmt.Println(singers, len(singers), cap(singers))
	singers[2] = "Criolo"
	fmt.Println(singers, len(singers), cap(singers))
	for _, singer := range singers {
		fmt.Printf("This singer is %s\r\n", singer)
	}

	// Print the second and third itens
	singersBr := singers[1:3]
	fmt.Println(singersBr, len(singersBr), cap(singersBr))

	// First two itens from slice
	firstCitys := citys[:2]
	fmt.Println(firstCitys)
	// Last item from slice
	lastCitys := citys[len(citys)-1]
	fmt.Println(lastCitys)

	// Deleting an item from slice
	indexDeleteItem := 3
	deleteItem := singers[:indexDeleteItem]
	deleteItem = append(deleteItem, singers[indexDeleteItem+1:]...)
	copy(singers, deleteItem)
	fmt.Println(singers)
}
