package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Color struct {
    name string
    number int
}

type Set struct {
    colors []Color
}

type Game struct {
    gameId int
    sets []Set
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

    color := Color{
        name: colorName,
        number: 0,
    }

    return color
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



