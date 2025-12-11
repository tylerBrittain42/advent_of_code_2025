package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(p1("01_pt1.txt"))
}

func p1(fileName string) (int, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return 0, fmt.Errorf("unable to open file: %v", err)
	}

	scanner := bufio.NewScanner(f)

	count := 0
	pos := 50
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text()[1:])
		if err != nil {
			return 0, fmt.Errorf("unable to parse int: %v", err)
		}
		if scanner.Text()[0] == 'L' {
			pos -= num
		} else if scanner.Text()[0] == 'R' {
			pos += num
		} else {
			return 0, fmt.Errorf("unexpected starting char")
		}
		fmt.Printf("prenorm: %d\n", pos)
		pos = normalizePos(pos)
		fmt.Printf("post: %d\n", pos)
		if pos == 0 {
			count += 1
		}

		fmt.Println(pos)

	}

	return count, nil
}

func normalizePos(num int) int {
	return ((num % 100) + 100) % 100
}
