package main

import (
	"fmt"
)

// func main() {
// main_task1()
// main_task2()
// main_task3()

// nums1 := []int{7, 7, 7}
// nums2 := []int{7, 7, 7}

// result := merge(nums1, nums2)
// fmt.Println(result)

// case1 := []int{1, 2, 2, 3}
// case2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

// result1 := removeDuplicates(case1)
// fmt.Println(result1)

// result2 := removeDuplicates(case2)
// fmt.Println(result2)

// removeDuplicatesInPlace(case1)

// mzCase1 := []int{4, 0, 1, 1, 0, 2, 2, 0, 3, 4}
// moveZeroes(mzCase1)

// }

/*
[1, 2, 3, 4, 5]
[1]
    [2, 3, 4, 5] ...
*/

// func main() {

// Это код показывает как в в слайсах меняется каписити и что добавление нового элемента в сред слайса , перезатирает значение
// slice := []int{1, 2, 3, 4, 5} // [1, 2, 3, 4, 5], 5, 5
// slice2 := slice[:1] // [1] 1, 5
// slice3 := slice[1:len(slice):cap(slice)] // [2, 3, 4, 5] 4, 4

// slice2[0] = 10 //
// slice2 = append(slice2, 100)

// fmt.Println(slice, len(slice), cap(slice)) // [10, 100, 3, 4, 5], 5, 5
// fmt.Println(slice2, len(slice2), cap(slice2)) // [10, 100] 2, 5
// fmt.Println(slice3, len(slice3), cap(slice3)) // [100, 3, 4, 5] 4, 4

// // Этот код показывает как меняется каписити слайса, когда происходит реалокация памяти
// t := []int{1, 2, 3, 4, 5}
// for i := 0; i < 20; i++ {
// 	fmt.Println(len(t), cap(t)) // ?
// 	t = append(t, i)
// }

// // cap_new - start_index

// slice4 := slice[1:3:4]
// fmt.Println(slice4, len(slice4), cap(slice4)) // ?
// fmt.Println(slice)
// slice4 = append(slice4, 100)
// fmt.Println(slice)
// fmt.Println(slice4)

// }

// func main() {
// 	lst := []string{"a", "b", "c", "d"}
// 	for i, v := range lst {
// 		if i == 0 {
// 			lst = []string{"1", "2", "3", "4"}
// 		}

// 		fmt.Println(v)
// 	}
// }

func main() {
	arr := make([]*int, 0, 7)
	var q int // 0xc000012080
	for range 5 {
		q++
		// if q == 4 {
		// 	b := 70
		// 	arr = append(arr, &b)
		// 	continue
		// }

		arr = append(arr, &q)
	}

	fmt.Println(arr)
	print(arr) // [5, 5, 5, 5, 5]

	mult(arr, 3)

	print(arr) // Ожидаем: [15, 15, 15, 15, 15]

	l := len(arr)
	arr = append(arr, &l)

	print(arr) // Ожидаем: [15 15 15 15 15 5]

	mult(arr, 3)

	print(arr) // Ожидаем: [45, 45, 45, 45, 45, 15]

	// TODO: исправить, тесты
}
