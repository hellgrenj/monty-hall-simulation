package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Monty Hall Problem Simulation")
	numberOfIterations := 100_000
	var playerChangesDoor bool

	playerChangesDoor = false
	wins := simulate(playerChangesDoor, numberOfIterations)
	fmt.Printf("player won %d out of %d games when staying with original pick\n", wins, numberOfIterations)

	playerChangesDoor = true
	wins = simulate(playerChangesDoor, numberOfIterations)
	fmt.Printf("player won %d out of %d games when changing door\n", wins, numberOfIterations)
}

func simulate(change bool, iterations int) int {
	wins := 0
	for i := 0; i < iterations; i++ {
		doors := []int{1, 2, 3}
		priceDoor := doors[getRandomDoorIndex(len(doors))]
		playerSelectedDoor := doors[getRandomDoorIndex(len(doors))]
		doors = showHostRemovesDoor(doors, priceDoor, playerSelectedDoor)
		if change {
			playerSelectedDoor = changeDoor(playerSelectedDoor, doors)
		}
		if playerSelectedDoor == priceDoor {
			wins++
		}

	}
	return wins
}
func showHostRemovesDoor(doors []int, priceDoor int, playerSelectedDoor int) []int {
	var rmIndex int
	for i, v := range doors {
		if v != priceDoor && v != playerSelectedDoor {
			rmIndex = i
			break
		}
	}
	doors = append(doors[:rmIndex], doors[rmIndex+1:]...)
	return doors
}
func changeDoor(currentDoor int, doors []int) int {
	for _, v := range doors {
		if currentDoor != v {
			currentDoor = v
			break
		}
	}
	return currentDoor
}
func getRandomDoorIndex(numberOfDoors int) int {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := numberOfDoors
	return rand.Intn(max-min) + min
}
