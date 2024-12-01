package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal("Couldn't open input file", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var l1 []int
	var l2 []int

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 2 {
			log.Fatalf("Invalid line: %s,", line)
		}
		num1, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("Invalid number %s: %v", parts[0], err)
		}

		num2, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("Invalid number %s: %v", parts[1], err)
		}

		l1 = append(l1, int(num1))
		l2 = append(l2, int(num2))
	}

	sort.Ints(l1)
	sort.Ints(l2)

	var ans = 0

	for i := 0; i < len(l1); i++ {
		ans += int(math.Abs(float64(l1[i] - l2[i])))
	}

	fmt.Printf("ans %d", ans)
}
