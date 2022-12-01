package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
	"strconv"
	"math"
	"sort"
)

func main() {
    file, err := os.Open("day1/data_day1.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
	var res float64
	var cur float64
	var values []float64

    for scanner.Scan() {
        line := scanner.Text()

		// fmt.Print(line)
		// fmt.Print("test")
		// fmt.Println(len(line))

		if len(line) > 1 {
			intVar, _ := strconv.ParseFloat(line, 64)
			cur = cur + intVar
			// fmt.Println(cur)
			// fmt.Print(err)
		} else {
			values = append(values, cur)
			res = math.Max(res, cur)
			cur = 0
		}
    }

	if cur != 0 {
		values = append(values, cur)
		res = math.Max(res, cur)
	}

	fmt.Println("\nResult: ")
	fmt.Print(res)

	sort.Sort(sort.Reverse(sort.Float64Slice(values)))
	// fmt.Print("test", values)

	var res2 float64
	for i:= 0; i < 3; i++ {
		res2 += values[i]
	}

	fmt.Println("\n\nResult2: ")
	fmt.Print(res2)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}