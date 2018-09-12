package cos418_hw1_1

import (
	"fmt"
	"testing"
)

func test(t *testing.T, fileName string, num int, expected int) {
	result := sum(num, fileName)
	if result != expected {
		t.Fatal(fmt.Sprintf(
			"Sum of %s failed: got %d, expected %d\n", fileName, result, expected))
	}
}

func Test1(t *testing.T) {
	test(t, "q2_test1.txt", 1, 499500)
}

func Test2(t *testing.T) {
	test(t, "q2_test1.txt", 10, 499500)
}

func Test3(t *testing.T) {
	test(t, "q2_test2.txt", 1, 117652)
}

func Test4(t *testing.T) {
	test(t, "q2_test2.txt", 10, 117652)
}
