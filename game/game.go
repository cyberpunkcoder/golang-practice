package main

import "fmt"

func updateScreen(playerX, playerY, enemyX, enemyY int) {
	for y := 0; y <= 8; y++ {
                for x := 0; x <= 16; x++ {
                        fmt.Print("░")
                }
                fmt.Println();
        }
}

func main() {
	updateScreen(1, 1, 2, 2)
}

