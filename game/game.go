package main

import (
	"fmt"
	"time"
	"github.com/eiannone/keyboard"
)

var MAPX int
var MAPY int

var playerX int
var playerY int

var enemyX int
var enemyY int

func renderGame() {
	fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")

	for y := 0; y <= MAPY; y++ {
                for x := 0; x <= MAPX; x++ {
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

func enemyAI() {
	if enemyX < playerX {

	} else {

	}

	if enemyY < playerY {

	} else
	{

	}
}

func gameLoop() {
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	fmt.Println("OBJECTIVE: Catch the '@' sign!\n")

	fmt.Println("CONTROLS:")
	fmt.Println("Move: w,a,s,d")
	fmt.Println("Quit: ESC\n")

	fmt.Println("Press any key to begin...\n")

	for {
		char, key, err := keyboard.GetKey()

		if (err != nil) {
			panic(err)
		} else if (key == keyboard.KeyEsc) {
			break
		}

		switch char {
			case 'w':
				if playerY > 0 {
					playerY--
				}
			case 'a':
				if playerX > 0 {
					playerX--
				}
			case 's':
				if playerY < MAPY {
					playerY++
				}
			case 'd':
				if playerX < MAPX {
					playerX++
				}
		}

		enemyAI()
		renderGame()

		time.Sleep(500)
	}
}

func main() {

	MAPX = 16
	MAPY = 8

	playerX = 1
	playerY = 1
	enemyX = 8
	enemyY = 4

	gameLoop()
}

