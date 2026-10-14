package main

import (
	"fmt"
	"os"
)

func main() {
	var n, m int
	fmt.Fscan(os.Stdin, &n, &m)

	player := make([][3]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(os.Stdin, &player[i][0], &player[i][1], &player[i][2])
	}

	for i := 0; i < m; i++ {
		var order string
		deck := [3]int{}
		fmt.Fscan(os.Stdin, &order, &deck[0], &deck[1], &deck[2])

		switch order {
		case "a":
			for i := range player {
				player[i][0] = player[i][0] | deck[0]
				player[i][1] = player[i][1] | deck[1]
				player[i][2] = player[i][2] | deck[2]
			}
		case "b":
			for i := range player {
				player[i][0] = player[i][0] & deck[0]
				player[i][1] = player[i][1] & deck[1]
				player[i][2] = player[i][2] & deck[2]
			}
		case "c":
			for i := range player {
				player[i][0] = map[bool]int{true: 1, false: 0}[player[i][0] != deck[0]]
				player[i][1] = map[bool]int{true: 1, false: 0}[player[i][1] != deck[1]]
				player[i][2] = map[bool]int{true: 1, false: 0}[player[i][2] != deck[2]]
			}
		}

		// fmt.Fprint(os.Stdout, "第", i+1, "回/ order: ", order, "/ deck: ", deck, "\n")
		// fmt.Fprint(os.Stdout, "player1: ", player[0], "\n")
		// fmt.Fprint(os.Stdout, "player2: ", player[1], "\n")
		// fmt.Fprint(os.Stdout, "player3: ", player[2], "\n")
	}

	winner := []int{0}
	maxValue := 0

	for i := 0; i < n; i++ {
		value := 0
		for _, bit := range player[i] {
			value = value*2 + bit
		}
		if value > maxValue {
			maxValue = value
			winner = []int{i}
		} else if value == maxValue {
			if i == 0 {
				continue
			}
			winner = append(winner, i)
		}
	}

	for _, w := range winner {
		fmt.Fprintln(os.Stdout, w+1)
	}
}
