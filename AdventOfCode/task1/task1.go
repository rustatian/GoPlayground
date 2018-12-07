package task1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	a, _ := os.Open("/Users/0xdev/Projects/repo/GoPlayground/AdventOfCode/task1.txt")
	var num int

	scanner := bufio.NewScanner(a)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.ContainsAny("+", text) {
			trim := strings.Trim(text, "+")
			n, _ := strconv.Atoi(trim)
			num += n
		} else {
			trim := strings.Trim(text, "-")
			n, _ := strconv.Atoi(trim)
			num -= n
		}
	}
	fmt.Println(num)
}



