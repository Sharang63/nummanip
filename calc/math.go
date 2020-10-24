package calc

import "errors"

//Add two numbers
func Add(numbers ...int) (error, int) {
	sum := 0
	if len(numbers) < 2 {
		return errors.New("Provide more than two numbers"), sum
	} else {
		for _, num := range numbers {
			sum = sum + num
		}
		return nil, sum
	}
}
