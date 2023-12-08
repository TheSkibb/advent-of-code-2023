// https://adventofcode.com/2023/day/1
package main

import (
	"bufio"
	"fmt"
	"strconv"
	"unicode"

	//"bufio"
	"log"
	"os"
)

func main(){
    fmt.Println("....")

    calibrationVal := findCalibrationVal("../test-input.txt")

    fmt.Println(calibrationVal == 142)

    fmt.Println(findCalibrationVal("../input.txt"))
}

func findCalibrationVal(file string) int{
    fmt.Println(file)

    f, err := os.Open(file)

    if err != nil{
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)

    calVal := 0

    for scanner.Scan() {
        calVal+=findNumInLine(scanner.Text())
    }

    return calVal
}

func findNumInLine(line string) int{
    num := ""
    //fund first num
    for _, c := range line {
        if unicode.IsDigit(c) {
            num += string(c)
            break
        }
    }

    runes := []rune(line)

    for i := len(runes) - 1; i >= 0; i--{
        if unicode.IsDigit(runes[i]) {
            num += string(runes[i])
            break
        }
    }

    //fmt.Println(num)
    result, _ := strconv.Atoi(num)
    return result
}
