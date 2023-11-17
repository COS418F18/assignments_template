package cos418_hw1_1

import (
	"os"
	"strconv"
	"strings"
)

// Sum numbers from channel `nums` and output sum to `out`.
// You should only output to `out` once.
// Do NOT modify function signature.
func sumWorker(nums chan int, out chan int) {
	// TODO: implement me
	// HINT: use for loop over `nums`
	ans := 0
	for value := range nums {
		ans += value
	}
	out <- ans
	close(out)
}

// Read integers from the file `fileName` and return sum of all values.
// This function must launch `num` go routines running
// `sumWorker` to find the sum of the values concurrently.
// You should use `checkError` to handle potential errors.
// Do NOT modify function signature.
func sum(num int, fileName string) int {
	// TODO: implement me
	// HINT: use `readInts` and `sumWorkers`
	// HINT: used buffered channels for splitting numbers between workers
	numbers, err := readInts(fileName)
	if err != nil {
		checkError(err)
	}
	subArraySize := len(numbers) / num
	ans := 0
	for i := 0; i < num; i++ {
		res := make(chan int, 1)
		from := i * subArraySize
		to := (i + 1) * subArraySize
		numsChan := make(chan int, subArraySize)
		for _, ele := range numbers[from:to] {
			numsChan <- ele
		}
		close(numsChan)
		go sumWorker(numsChan, res)
		ans += <-res
	}
	// to get rid of the warning
	return ans
}

// Read a list of integers separated by whitespace from `r`.
// Return the integers successfully read with no error, or
// an empty slice of integers and the error that occurred.
// Do NOT modify this function.
func readInts(path string) ([]int, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		checkError(err)
	}
	elems := strings.Fields(string(content))
	intElements := []int{}
	for _, ele := range elems {
		num, err := strconv.Atoi(ele)
		if err != nil {
			checkError(err)
		}

		intElements = append(intElements, num)
	}
	return intElements, nil
}
