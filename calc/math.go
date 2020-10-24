package calc

//Add two numbers
func Add(numbers ...int) int {
	sum := 0
	for _, num := range numbers {
		sum = sum + num
	}
	return sum
}
