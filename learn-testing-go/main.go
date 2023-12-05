package main

//func main() {
//	fmt.Println(GenerateArray())
//}

func Max(nums []int) int {
	var max int
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

//func GenerateArray() []int {
//	rand.Seed(time.Now().UnixNano())
//	arraySize := 15
//	randomArray := make([]int, arraySize)
//	for i := 0; i < arraySize; i++ {
//		randomArray[i] = rand.Intn(100)
//	}
//	return randomArray
//}
