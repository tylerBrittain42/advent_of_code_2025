package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(p1("01_pt1.txt"))
	fmt.Println(p2("01_pt1.txt"))
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
		pos = normalizePos(pos)
		if pos == 0 {
			count += 1
		}

	}

	return count, nil
}

func normalizePos(num int) int {
	return ((num % 100) + 100) % 100
}

func p2(fileName string) (int, error) {
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
		dir := scanner.Text()[0]

		fmt.Println(num)
		for num > 0 {
			pos = tick(pos, dir)
			if pos == 100 {
				pos = 0
			}
			if pos == 0 {
				count++
			}
			if pos == -1 {
				pos = 99
			}
			num--
		}

	}

	return count, nil
}

func tick(start int, dir byte) int {
	if dir == 'L' {
		start--
	} else {
		start += 1
	}
	return start

}

// func normalizePos2(num int) (int, bool) {
// return ((num % 100) + 100) % 100, false
// }
