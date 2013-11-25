package models

import (
	"github.com/fhbzyc/poker/libs/array"
)

func GetPoker() []int {
	return []int{
		2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, //红桃
		102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, //方快
		1002, 1003, 1004, 1005, 1006, 1007, 1008, 1009, 1010, 1011, 1012, 1013, 1014, //黑桃
		10002, 10003, 10004, 10005, 10006, 10007, 10008, 10009, 10010, 10011, 10012, 10013, 10014, //梅花
	}
}

//皇家同花顺
func IsRoyalFlush(pokers []int) bool {

	array.RSort(&pokers)

	for i := range pokers {
		if i <= len(pokers)-5 && pokers[i]%100 == 14 {
			if pokers[i]-1 == pokers[i+1] &&
				pokers[i]-2 == pokers[i+2] &&
				pokers[i]-3 == pokers[i+3] &&
				pokers[i]-4 == pokers[i+4] {

				return true
			}
		}
	}

	return false
}

//同花顺
func IsStraightFlush(pokers []int) (bool, []int) {

	for i := range pokers {
		if pokers[i]%100 == 14 {
			pokers = append(pokers, pokers[i]-13)
		}
	}

	array.RSort(&pokers)

	for i := range pokers {
		if i <= len(pokers)-5 &&
			pokers[i]-1 == pokers[i+1] &&
			pokers[i]-2 == pokers[i+2] &&
			pokers[i]-3 == pokers[i+3] &&
			pokers[i]-4 == pokers[i+4] {

			if pokers[i+4]%100 == 1 {
				pokers[i+4] += 13
			}

			return true, []int{pokers[i], pokers[i+1], pokers[i+2], pokers[i+3], pokers[i+4]}

		}
	}

	return false, []int{}
}

//4条
func IsFourOfAKind(pokers []int) (bool, []int) {

	putOffColor(&pokers)
	array.RSort(&pokers)

	for i := range pokers {
		if i <= len(pokers)-4 &&
			pokers[i] == pokers[i+1] &&
			pokers[i] == pokers[i+2] &&
			pokers[i] == pokers[i+3] {

			for _, val := range pokers {
				if val != pokers[i] {
					return true, []int{pokers[i], pokers[i+1], pokers[i+2], pokers[i+3], val}
				}
			}
		}
	}

	return false, []int{}
}

//葫芦(满堂彩)
func IsFullHouse(pokers []int) (bool, []int) {

	putOffColor(&pokers)
	array.RSort(&pokers)

	for i := range pokers {
		if i <= len(pokers)-3 &&
			pokers[i] == pokers[i+1] &&
			pokers[i] == pokers[i+2] {

			for j := range pokers {
				if j <= len(pokers)-2 && pokers[j] != pokers[i] && pokers[j+1] == pokers[j] {
					return true, []int{pokers[i], pokers[i+1], pokers[i+2], pokers[j], pokers[j+1]}
				}
			}
		}
	}

	return false, []int{}
}

//同花
func IsFlush(pokers []int) (bool, []int) {

	array.RSort(&pokers)

	var base int
	for i := range pokers {
		if i <= len(pokers)-5 {
			base = pokers[i] - (pokers[i] % 100)
			if base == pokers[i+1]-(pokers[i+1]%100) &&
				base == pokers[i+2]-(pokers[i+2]%100) &&
				base == pokers[i+3]-(pokers[i+3]%100) &&
				base == pokers[i+4]-(pokers[i+4]%100) {

				return true, []int{pokers[i], pokers[i+1], pokers[i+2], pokers[i+3], pokers[i+4]}
			}
		}

	}

	return false, []int{}
}

//顺子
func IsStraight(pokers []int) (bool, []int) {

	for i := range pokers {
		if pokers[i]%100 == 14 {
			pokers = append(pokers, pokers[i]-13)
		}
	}

	putOffColor(&pokers)
	array.RSort(&pokers)

	var i int
	var find int
	for _, val := range pokers {

		find = 1
		for i = 1; i < 5; i++ {

			for _, v := range pokers {
				if v == val-i {

					find++
					if find == 5 {
						if val-4 == 1 {
							return true, []int{val, val - 1, val - 2, val - 3, 14}
						}
						return true, []int{val, val - 1, val - 2, val - 3, val - 4}
					}
				}
			}
		}
	}

	return false, []int{}
}

//3条
func IsThreeOfAKind(pokers []int) (bool, []int) {

	putOffColor(&pokers)
	array.RSort(&pokers)

	for i := range pokers {
		if i <= len(pokers)-3 &&
			pokers[i] == pokers[i+1] &&
			pokers[i] == pokers[i+2] {

			for j := range pokers {
				if pokers[j] != pokers[i] && pokers[j+1] != pokers[i] {
					return true, []int{pokers[i], pokers[i+1], pokers[i+2], pokers[j], pokers[j+1]}
				}
			}
		}
	}

	return false, []int{}
}

//两对
func IsTowPair(pokers []int) (bool, []int) {

	putOffColor(&pokers)
	array.RSort(&pokers)

	for i := range pokers {
		if i <= len(pokers)-4 && pokers[i] == pokers[i+1] {
			for j := range pokers {
				if j <= len(pokers)-2 && pokers[j] != pokers[i] && pokers[j+1] == pokers[j] {
					for _, val := range pokers {
						if val != pokers[i] && val != pokers[j] {
							return true, []int{pokers[i], pokers[i+1], pokers[j], pokers[j+1], val}
						}
					}
				}
			}
		}
	}

	return false, []int{}
}

//一对
func IsOnePair(pokers []int) (bool, []int) {

	putOffColor(&pokers)
	array.RSort(&pokers)

	var res []int
	for i := range pokers {
		if i <= len(pokers)-2 && pokers[i] == pokers[i+1] {

			res = append(res, pokers[i])
			res = append(res, pokers[i+1])

			for j := range pokers {
				if j != i && j != i+1 {
					res = append(res, pokers[j])
					if len(res) == 5 {
						return true, res
					}
				}
			}
		}
	}

	return false, []int{}
}

//去掉花色
func putOffColor(pokers *[]int) {

	data := *pokers

	for index, val := range data {
		if val > 10000 {
			data[index] -= 10000
		} else if val > 1000 {
			data[index] -= 1000
		} else if val > 100 {
			data[index] -= 100
		}
	}

	*pokers = data
}
