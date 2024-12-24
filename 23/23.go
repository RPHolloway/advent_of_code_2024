package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

const FILE_NAME = "input.txt"

type Pair struct {
	C1 string
	C2 string
}

var Pairs []Pair

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func run() {
	connectionMap := make(map[string][]string)
	for _, pair := range Pairs {
		connectionMap[pair.C1] = append(connectionMap[pair.C1], pair.C2)
		connectionMap[pair.C2] = append(connectionMap[pair.C2], pair.C1)
	}

	groups := make(map[[3]string]struct{})
	for c1, peers := range connectionMap {
		for _, c2 := range peers {
			for _, c3 := range connectionMap[c2] {
				if slices.Contains(connectionMap[c3], c1) {
					if c1[0] == 't' || c2[0] == 't' || c3[0] == 't' {
						set := [3]string{c1, c2, c3}
						// Sort for uniqueness
						slices.Sort(set[:])
						groups[set] = struct{}{}
					}
				}
			}
		}
	}

	//for k := range groups {
	//	fmt.Println(k)
	//}

	fmt.Println(len(groups))

	lanParty := make(map[string]struct{})
	for c1, peers := range connectionMap {
		for i := range len(peers) {
			party := make(map[string]struct{})
			party[c1] = struct{}{}
			party[peers[i]] = struct{}{}

			for _, peer := range peers {
				lan := true
				for c := range party {
					if !slices.Contains(connectionMap[peer], c) {
						lan = false
						break
					}
				}

				if lan {
					party[peer] = struct{}{}
				}
			}

			//fmt.Println(party)
			if len(party) > len(lanParty) {
				lanParty = party
			}
		}
	}

	//fmt.Println(lanParty)
	var lanPartySet []string
	for k := range lanParty {
		lanPartySet = append(lanPartySet, k)
		slices.Sort(lanPartySet[:])
	}

	for _, c := range lanPartySet {
		fmt.Printf("%s,", c)
	}
	fmt.Println()
}

func main() {
	// Read input
	data, _ := os.ReadFile(FILE_NAME)
	input := string(data)

	// Parse input
	lines := strings.Split(input, "\r\n")
	for _, line := range lines {
		x := strings.Split(line, "-")
		Pairs = append(Pairs, Pair{C1: x[0], C2: x[1]})
	}

	defer timeTrack(time.Now())

	run()
}
