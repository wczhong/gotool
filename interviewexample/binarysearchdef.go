package main

import (
	"fmt"
)

// 一分为2， 肯定有一半是有序的
func BinarySearchDef(data []int, key int, low int, high int) int {
	if len(data) == 0 {
		return -1
	}
	// 注意临界值
	for low <= high {
		mid := low + (high-low)/2

		if key == data[mid] {
			return mid
		}

		// 如果 key 比中间值小
		if key < data[mid] {
			if key == data[low] {
				return low
			}
			//如果key 比中间值小，但是比 low 位置上的值大
			if key > data[low] {
				// 如果 key比中间值小， 比low位置上的值大， 同时 low 到 mid 之间的元素是无序的
				if mid-1 >= low && data[low] >= data[mid-1] {
					return BinarySearchDef(data, key, low, mid-1)
				}
				// 有序则反查另一半
				high = mid - 1
			}
			if key < data[low] {
				if mid-1 >= low && data[low] >= data[mid-1] {
					return BinarySearchDef(data, key, low, mid-1)
				}
				low = mid + 1
			}
		}

		if key > data[mid] {
			if key == data[high] {
				return high
			}
			if key > data[high] {
				if mid+1 <= high && data[mid+1] >= data[high] {
					return BinarySearchDef(data, key, mid+1, high)
				}
				// 有序则反查另一半, 因为该部分中肯定无对应的值
				high = mid - 1
			}
			if key < data[high] {
				if mid+1 <= high && data[mid+1] >= data[high] {
					return BinarySearchDef(data, key, mid+1, high)
				}
				low = mid + 1
			}
		}
	}
	return -1
}

func main() {
	data := []int{100, 102, 104, 106, 108, 110, 112, 114, 116, 118, 120, 121, 122, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	i := BinarySearchDef(data, 3, 0, len(data)-1)
	fmt.Println(i)
}
