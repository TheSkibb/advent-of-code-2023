package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Color struct {
    name string
    number int
}

type Game struct {
    gameId int
    colors []Color
}

func main(){
    fmt.Println("december 2. solution")

    getGames("./test-input.txt")
}

func checkGame(line string) Game {

    //get gameId
    gameId, _ := strconv.Atoi(string((line[5])))

    //get number red
    red := getColor("red", line)
    blue := getColor("blue", line)
    green := getColor("green", line)


    //initialize color array
    var colors []Color

    colors = append(colors, red)
    colors = append(colors, blue)
    colors = append(colors, green)

    game := Game{
        gameId: gameId,
        colors: colors,
    }

    fmt.Println(game)

    return game
}

func getColor(colorName string, line string) Color{

    indexes := findAllOccurrencesIndexes(line, colorName)
    var colorNumber int

    for i := 0; i<len(indexes); i++{
        firstDigit := string(line[indexes[i]-3])
        lastDigit := string(line[indexes[i]-2])
        if firstDigit == " "{
            colorNumber, _ = strconv.Atoi(lastDigit)
        }else{
            colorNumber, _ = strconv.Atoi(firstDigit + lastDigit)
        }
    }

    color := Color{
        name: colorName,
        number: colorNumber,
    }

    return color
}

//gpt generated, need to revisit this
func findAllOccurrencesIndexes(s, substr string) []int {
	indexes := make([]int, 0)
	idx := 0
	for {
		i := strings.Index(s[idx:], substr)
		if i == -1 {
			break
		}
		idx += i
		indexes = append(indexes, idx)
		idx++
	}
	return indexes
}

//reads file and returns all the games
func getGames(fileName string) []Game {
    fmt.Println("getting games from file")

    file, err := os.Open(fileName)

    if err != nil {
        log.Fatal("failed opening file")
    }

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    var txtlines []string

    for scanner.Scan(){
        txtlines = append(txtlines, scanner.Text())
    }

    file.Close()

    for _, eachline := range txtlines {
        checkGame(eachline)
    }
    return nil
}



