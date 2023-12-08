// https://adventofcode.com/2023/day/1
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type number_value struct {
    name string
    value int
    index int
}

func main(){
    fmt.Println("-----")

    //calibrationVal := findCalibrationVal("../test-input.txt")

    //fmt.Println(calibrationVal == 281)

    fmt.Println(findCalibrationVal("../input.txt"))
}

func findCalibrationVal(file string) int{

    f, err := os.Open(file)

    if err != nil{
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)

    calVal := 0
    lines := 0

    for scanner.Scan() {
        calVal+=findNumInLine(scanner.Text())
        lines += 1
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    return calVal
}

func findNumInLine(line string) int{

    stringNumbers := []number_value{
        {"one", 1, 0},
        {"two", 2, 0},
        {"three", 3, 0},
        {"four", 4, 0},
        {"five", 5, 0},
        {"six", 6, 0},
        {"seven", 7, 0},
        {"eight", 8, 0},
        {"nine", 9, 0},
        {"1", 1, 0},
        {"2", 2, 0},
        {"3", 3, 0},
        {"4", 4, 0},
        {"5", 5, 0},
        {"6", 6, 0},
        {"7", 7, 0},
        {"8", 8, 0},
        {"9", 9, 0},
    }

    var num []number_value

    for _, number := range stringNumbers {
        index := strings.Index(line, number.name)
        if index != -1 {
            number.index = index
            num = append(num, number)
        }
    }


    sort.Slice(num, func(i, j int) bool {
      return num[i].index < num[j].index
    })

    fmt.Println(num)

    firstnum := strconv.Itoa(num[0].value)
    lastnum := strconv.Itoa(num[len(num) - 1].value)


    fmt.Println(firstnum + lastnum)


    result, _ := strconv.Atoi(firstnum + lastnum)

    return result
}
