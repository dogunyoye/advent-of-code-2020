package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/scylladb/go-set/strset"
)

func findIngredientsWithNoAllergens(ingredientsMap map[string]int, allergensMap map[string]*strset.Set, possibleAllergens *strset.Set) int {

	for possibleAllergens.Size() != 0 {
		for k, v := range allergensMap {
			if v.Size() == 1 {
				for kk, vv := range allergensMap {
					if kk == k {
						continue
					}
					vv.Remove(v.List()[0])
				}
				possibleAllergens.Remove(k)
			}
		}
	}

	var result = 0
	for k, v := range ingredientsMap {
		var found = false
		for _, vv := range allergensMap {
			if vv.List()[0] == k {
				found = true
				break
			}
		}

		if !found {
			result += v
		}
	}

	return result
}

func main() {
	file, err := os.Open("../../data/day21.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var ingredientsMap = make(map[string]int)
	var allergensMap = make(map[string]*strset.Set)

	var possibleAllergens = strset.New()

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " (contains ")
		ingredients := strings.Split(split[0], " ")
		var allergenSet = strset.New()

		for _, ingredient := range ingredients {
			ingredientsMap[ingredient]++
			allergenSet.Add(ingredient)
		}

		allergens := strings.ReplaceAll(split[1], ")", "")

		split2 := strings.Split(allergens, ", ")
		for _, s := range split2 {
			_, exists := allergensMap[s]
			possibleAllergens.Add(s)

			if !exists {
				allergensMap[s] = allergenSet
			} else {
				allergensMap[s] = strset.Intersection(allergensMap[s], allergenSet)
			}
		}
	}

	file.Close()

	var part1 = findIngredientsWithNoAllergens(ingredientsMap, allergensMap, possibleAllergens)

	var allergensList []string

	for k := range allergensMap {
		allergensList = append(allergensList, k)
	}

	var dangerList = ""
	sort.Strings(allergensList)

	for _, k := range allergensList {
		dangerList += allergensMap[k].List()[0] + ","
	}

	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", dangerList[:len(dangerList)-1])
}
