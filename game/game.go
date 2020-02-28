package main

import "fmt"

var playerX int
var playerY int

var enemyX int
var enemyY int

func updateScreen() {
	for y := 0; y <= 8; y++ {
                for x := 0; x <= 16; x++ {
                        if x == playerX && y == playerY {
				fmt.Print("╬")
			} else if x == enemyX && y == enemyY {
				fmt.Print("@")
			} else {
				fmt.Print("░")
			}
                }
                fmt.Println();
        }
}

func main() {

	playerX = 1
	playerY = 1
	enemyX = 8
	enemyY = 4

	updateScreen()
}

