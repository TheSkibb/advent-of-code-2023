package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main(){
    //readlines("./test-input.txt")
    readlines("./input.txt")
}

//returns gameId if the game is possible, 0 if not (game id cannot be 0)
func gamePossible(game string) int{
    fmt.Println(game)

    // define restrictions
    rest_red := 12
    rest_green := 13
    rest_blue := 14

    _ = rest_red
    _ = rest_blue
    _ = rest_green

    //get gameId
    gameId, _ := strconv.Atoi(game[5:strings.Index(game, ":")])

    //fmt.Println(gameId)

    //split into sets
    sets :=  strings.Split(game, ";")

    //fmt.Println("sets", len(sets))
    //iterate sets and check if possible
    for i := 0; i<len(sets); i++{
        //fmt.Println("set", sets[i])
        red := checkColor("red", sets[i])
        blue := checkColor("blue", sets[i])
        green := checkColor("green", sets[i])

        if red > rest_red || blue > rest_blue || green > rest_green{
            return 0
        }
    }
    return gameId
}

func checkColor(color string, set string) int{

    colorIndex := strings.Index(set, color)

    if colorIndex == -1{
        return 0
    }

    var colorNumber int

    firstDigit := string(set[colorIndex-3])
    secondDigit := string(set[colorIndex-2])
    if firstDigit == " "{
        colorNumber, _ = strconv.Atoi(secondDigit)
    } else{
        colorNumber, _ = strconv.Atoi(firstDigit + secondDigit)
    }

    return colorNumber
}

func readlines(filename string){
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

    sum := 0

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
        result := gamePossible(scanner.Text())
        sum += result
	}

    fmt.Println("end result:", sum)

	file.Close()
}
