package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
)

var one = big.NewInt(1)

// taken from Rosetta code
// https://rosettacode.org/wiki/Chinese_remainder_theorem#Go
func crt(a, n []*big.Int) (*big.Int, error) {
	p := new(big.Int).Set(n[0])
	for _, n1 := range n[1:] {
		p.Mul(p, n1)
	}

	var x, q, s, z big.Int
	for i, n1 := range n {
		q.Div(p, n1)
		z.GCD(nil, &s, n1, &q)
		if z.Cmp(one) != 0 {
			return nil, fmt.Errorf("%d not coprime", n1)
		}
		x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
	}

	return x.Mod(&x, p), nil
}

func main() {
	file, err := os.Open("../../data/day13.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var busIDs []int

	scanner.Scan()
	var timeStamp, _ = strconv.Atoi(scanner.Text())

	scanner.Scan()
	buses := strings.Split(scanner.Text(), ",")

	for _, id := range buses {
		busID, err := strconv.Atoi(id)
		if err == nil {
			busIDs = append(busIDs, busID)
		}
	}

	file.Close()

	var busFound = false
	var timeElapsed = 0
	var id = -1

	for {
		var currentTime = timeStamp + timeElapsed
		for _, b := range busIDs {
			if currentTime%b == 0 {
				busFound = true
				id = b
				break
			}
		}

		if busFound {
			break
		}

		timeElapsed++
	}

	var part1 = timeElapsed * id

	var remainders []*big.Int
	var mods []*big.Int

	for i, id := range buses {
		busID, err := strconv.Atoi(id)
		if err == nil {
			remainders = append(remainders, big.NewInt(int64(-i)))
			mods = append(mods, big.NewInt(int64(busID)))
		}
	}

	fmt.Println("Part1:", part1)

	// Part 2 is related to CRT (Chinese Remainder Theorem)
	// Can use an online calculator https://www.dcode.fr/chinese-remainder
	// with the following values specific to my input:
	// remainder	modulo
	// ---------	------
	//  0			29
	// -19			41
	// -29			601
	// -37			23
	// -42			13
	// -46			17
	// -48			19
	// -60			463
	// -97			37

	var part2, _ = crt(remainders, mods)
	fmt.Println("Part2:", part2)

	var time = 0
	var step = 1

	for offset, busID := range buses {
		if busID != "x" {
			b, _ := strconv.Atoi(busID)
			for (time+offset)%b != 0 {
				time += step
			}
			step *= b
		}
	}

	fmt.Println("Part2 (Alternative):", time)
}
