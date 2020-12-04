package main

/*
 Import format package for printing to console.
 Import os console to exit the program.
 Import eiannonne keyboard to capture keyboard input
*/
import (
	"fmt"
	"os"

	"github.com/eiannone/keyboard"
)

// The size of the map in x and y dimensions.
var MAPX int
var MAPY int

// The player's position where 0,0 is the top left corner.
var playerX int
var playerY int

// The enemy's position where 0,0 is the top left corner.
var enemyX int
var enemyY int

// Renders the play screen of the game with the background, player and enemy.
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
		fmt.Println()
	}
}

// Really simple enemy AI that moves away from the player.
func enemyAI() {
	if enemyX < playerX {
		if enemyX > 0 {
			enemyX--
		} else {
			enemyX++
		}
	} else {
		if enemyX < MAPX {
			enemyX++
		} else {
			enemyX--
		}
	}

	if enemyY < playerY {
		if enemyY > 0 {
			enemyY--
		} else {
			enemyY++
		}
	} else {
		if enemyY < MAPY {
			enemyY++
		} else {
			enemyY--
		}
	}
}

// Checks if the player has caught the enemy.
func checkWin() {
	if playerX == enemyX && playerY == enemyY {
		renderGame()
		fmt.Println("\n\n\nWINNER! YOU CAUGHT THE '@'!")
		os.Exit(0)
	}
}

// The main loop of the game, this loop pauses until it sees keyboard input.
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

		if err != nil {
			panic(err)
		} else if key == keyboard.KeyEsc {
			break
		}

		// Switch handles the user's input for movement.
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
		checkWin()
	}
}

// Main function and start of execution.
func main() {

	MAPX = 16
	MAPY = 8

	playerX = 1
	playerY = 1
	enemyX = 8
	enemyY = 4

	gameLoop()
}
