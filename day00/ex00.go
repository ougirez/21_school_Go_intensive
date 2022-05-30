package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	// "os"
)

func fillArray() []int {
	var temp int
	var nums []int

	for {
		_, err := fmt.Scan(&temp)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatalf("%s\n", "Wrong input")
		}
		if temp < -100000 || temp > 100000 {
			log.Fatalf("%s\n", "Out of range")
		}
		nums = append(nums, temp)
	}
	return nums
}

func mean(nums []int) float64 {
	var mean float64

	for _, n := range nums {
		mean += float64(n)
	}
	return mean / float64(len(nums))
}

func median(nums []int) float64 {
	temp := append([]int{}, nums...)
	sort.Ints(temp)
	if len(temp)%2 == 1 {
		return float64(temp[len(temp)/2])
	} else {
		return float64(temp[len(temp)/2-1]+temp[len(temp)/2]) / 2
	}
}

func mode(nums []int) int {
	var freqMap = make(map[int]int)
	var mode = []int{0, 0}

	for _, n := range nums {
		freqMap[n] += 1
		if freqMap[n] > mode[1] ||
			freqMap[n] == mode[1] && n < mode[0] {
			mode[0] = n
			mode[1] = freqMap[n]
		}
	}

	return mode[0]
}

func SD(nums []int) float64 {
	var sd float64
	var mean = mean(nums)

	if len(nums) < 2 {
		return 0
	}
	for _, n := range nums {
		sd += math.Pow(float64(n)-mean, 2)
	}
	sd /= float64(len(nums))
	return math.Sqrt(sd)

}

func main() {
	var nums = fillArray()
	if len(nums) == 0 {
		log.Fatalf("%s\n", "input data is empty")
	}
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "mean" {
			fmt.Printf("Mean: %.2f\n", mean(nums))
		} else if os.Args[i] == "median" {
			fmt.Printf("Median: %.2f\n", median(nums))
		} else if os.Args[i] == "mode" {
			fmt.Printf("Mode: %d\n", mode(nums))
		} else if os.Args[i] == "sd" {
			fmt.Printf("SD: %.2f\n", SD(nums))
		}
	}
	if len(os.Args) == 1 {
		fmt.Printf("Mean: %.2f\n", mean(nums))
		fmt.Printf("Median: %.2f\n", median(nums))
		fmt.Printf("Mode: %d\n", mode(nums))
		fmt.Printf("SD: %.2f\n", SD(nums))
	}
}
