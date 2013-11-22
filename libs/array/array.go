package array

import (
	"fmt"
	"math/rand"
	"time"
)

var seed int64

//打乱一维切片
func StringShuffle(array *[]string) {

	seed++

	rand.Seed(time.Now().UnixNano() + seed)

	length := len(*array)

	res := make([]string, length)
	lock := make([]string, length)

	for _, val := range *array {
		num := rand.Intn(length - 1)
		for {
			if lock[num] != "yes" {
				fmt.Println(num)
				res[num] = val
				lock[num] = "yes"
				break
			} else if num < length-1 {
				num++
				continue
			} else {
				num = 0
			}
		}
	}

	*array = res

	fmt.Println(array)
}

//打乱一维切片
func IntShuffle(array *[]int) {

	seed++

	rand.Seed(time.Now().UnixNano() + seed)

	length := len(*array)

	res := make([]int, length)
	lock := make([]string, length)

	for _, val := range *array {
		num := rand.Intn(length - 1)
		for {
			if lock[num] != "yes" {
				fmt.Println(num)
				res[num] = val
				lock[num] = "yes"
				break
			} else if num < length-1 {
				num++
				continue
			} else {
				num = 0
			}
		}
	}

	*array = res

	fmt.Println(array)
}

//从小到大排序
func Sort(array *[]int) {

	data2 := *array

	for i := 0; i < len(data2); i++ {
		for j := 0; j < len(data2); j++ {
			if j > i && data2[j] < data2[i] {
				temp := data2[i]
				data2[i] = data2[j]
				data2[j] = temp
			}
		}
	}

	*array = data2
}

//从大到小排序
func RSort(array *[]int) {

	data2 := *array

	for i := 0; i < len(data2); i++ {
		for j := 0; j < len(data2); j++ {
			if j > i && data2[j] > data2[i] {
				temp := data2[i]
				data2[i] = data2[j]
				data2[j] = temp
			}
		}
	}

	*array = data2
}
