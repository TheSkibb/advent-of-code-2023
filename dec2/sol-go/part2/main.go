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

    // initialize
    fewest_red := 0
    fewest_green := 0
    fewest_blue := 0

    _ = fewest_red
    _ = fewest_blue
    _ = fewest_green

    //get gameId
    //gameId, _ := strconv.Atoi(game[5:strings.Index(game, ":")])

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

        if red != 0 && red > fewest_red {
            fewest_red = red
        }

        if blue != 0 && blue > fewest_blue {
            fewest_blue = blue
        }

        if green != 0 && green > fewest_green{
            fewest_green = green
        }
    }

    fmt.Println("fewest red:", fewest_red)
    fmt.Println("fewest green :", fewest_green)
    fmt.Println("fewest blue:", fewest_blue)

    power := fewest_red * fewest_green * fewest_blue
    return power
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
