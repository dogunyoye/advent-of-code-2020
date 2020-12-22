package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type combatPlayer struct {
	cards []int
}

type recursiveCombatPlayer struct {
	playerName string
	game       int
	round      int
	cards      []int
	order      map[string]struct{}
}

func copyArray(array []int) []int {
	arrNew := make([]int, 0)
	arrNew = append(arrNew, array...)
	return arrNew
}

func cardOrder(cards []int) string {
	var order = ""
	for _, card := range cards {
		order += strconv.Itoa(card) + ","
	}

	return order
}

func playRecursiveCombat(player1 recursiveCombatPlayer, player2 recursiveCombatPlayer) (int, recursiveCombatPlayer) {

	var result = 0
	var round = 1

	for {

		if len(player1.cards) == 0 || len(player2.cards) == 0 {
			break
		}

		p1Order := cardOrder(player1.cards)
		p2Order := cardOrder(player2.cards)

		_, exists1 := player1.order[p1Order]
		_, exists2 := player2.order[p2Order]

		if exists1 || exists2 {
			// player 1 wins - don't care about result
			return 0, player1
		}

		player1.order[p1Order] = struct{}{}
		player2.order[p2Order] = struct{}{}

		player1Card := player1.cards[0]
		player2Card := player2.cards[0]

		if player1Card <= len(player1.cards)-1 && player2Card <= len(player2.cards)-1 {
			// recurse into sub-game

			var game = player1.game + 1
			var _, winningPlayer = playRecursiveCombat(
				recursiveCombatPlayer{"player1", game, round, copyArray(player1.cards[1 : player1Card+1]), make(map[string]struct{})},
				recursiveCombatPlayer{"player2", game, round, copyArray(player2.cards[1 : player2Card+1]), make(map[string]struct{})})

			if winningPlayer.playerName == "player1" {
				player1.cards = player1.cards[1:]
				player1.cards = append(player1.cards, player1Card)
				player1.cards = append(player1.cards, player2Card)

				player2.cards = player2.cards[1:]
			} else {
				player2.cards = player2.cards[1:]
				player2.cards = append(player2.cards, player2Card)
				player2.cards = append(player2.cards, player1Card)

				player1.cards = player1.cards[1:]
			}
		} else {
			// play game normally
			if player1Card > player2Card {
				player1.cards = player1.cards[1:]
				player1.cards = append(player1.cards, player1Card)
				player1.cards = append(player1.cards, player2Card)

				player2.cards = player2.cards[1:]
			} else {
				player2.cards = player2.cards[1:]
				player2.cards = append(player2.cards, player2Card)
				player2.cards = append(player2.cards, player1Card)

				player1.cards = player1.cards[1:]
			}
		}

		round++
	}

	var winner recursiveCombatPlayer

	if len(player1.cards) > len(player2.cards) {
		winner = player1
	} else {
		winner = player2
	}

	for i := 0; i < len(winner.cards); i++ {
		result += (len(winner.cards) - i) * winner.cards[i]
	}

	return result, winner
}

func playCombat(player1 combatPlayer, player2 combatPlayer) int {

	for {

		if len(player1.cards) == 0 || len(player2.cards) == 0 {
			break
		}

		player1Card := player1.cards[0]
		player2Card := player2.cards[0]

		if player1Card > player2Card {
			player1.cards = player1.cards[1:]
			player1.cards = append(player1.cards, player1Card)
			player1.cards = append(player1.cards, player2Card)

			player2.cards = player2.cards[1:]
		} else {
			player2.cards = player2.cards[1:]
			player2.cards = append(player2.cards, player2Card)
			player2.cards = append(player2.cards, player1Card)

			player1.cards = player1.cards[1:]
		}
	}

	var winner combatPlayer

	if len(player1.cards) > len(player2.cards) {
		winner = player1
	} else {
		winner = player2
	}

	var result = 0
	for i := 0; i < len(winner.cards); i++ {
		result += (len(winner.cards) - i) * winner.cards[i]
	}

	return result
}

func main() {
	file, err := os.Open("../../data/day22.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var player1Combat []int
	var player2Combat []int

	var player1RecursiveCombat []int
	var player2RecursiveCombat []int

	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) != 0 {
			card, _ := strconv.Atoi(line)
			player1Combat = append(player1Combat, card)
			player1RecursiveCombat = append(player1RecursiveCombat, card)
			continue
		}

		break
	}

	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) != 0 {
			card, _ := strconv.Atoi(line)
			player2Combat = append(player2Combat, card)
			player2RecursiveCombat = append(player2RecursiveCombat, card)
			continue
		}

		break
	}

	var p1 = combatPlayer{player1Combat}
	var p2 = combatPlayer{player2Combat}

	var part1 = playCombat(p1, p2)

	var rp1 = recursiveCombatPlayer{"player1", 1, 1, player1RecursiveCombat, make(map[string]struct{})}
	var rp2 = recursiveCombatPlayer{"player2", 1, 1, player2RecursiveCombat, make(map[string]struct{})}

	var part2, _ = playRecursiveCombat(rp1, rp2)

	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)

	file.Close()

}
