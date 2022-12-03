package rucksack

import "log"

type Group [][]ItemType

func GroupFromStrings(sl ...string) Group {
	var g Group
	for _, s := range sl {
		g = append(g, []ItemType(s))
	}
	return g
}

func (g Group) FindUnifyingItemType() ItemType {
	for _, it := range g[0] {
		for _, it2 := range g[1] {
			for _, it3 := range g[2] {
				if it == it2 && it == it3 {
					return it
				}
			}
		}
	}

	log.Fatalln("no unifying item type was found")
	return 0
}

type ItemType rune

func (it ItemType) Priority() int {
	if it >= 97 && it <= 122 {
		return int(it) - 96
	}
	if it >= 65 && it <= 90 {
		return int(it) - 38
	}

	log.Fatalln("unsupported item type")
	return -1
}

type Compartments [2][]ItemType

func CompartmentsFromString(s string) Compartments {
	types := []ItemType(s)
	cLen := len(types) / 2
	first := types[:cLen]
	second := types[cLen:]
	return Compartments{first, second}
}

func (c Compartments) FindProblematicItemType() ItemType {
	for _, it := range c[0] {
		for _, it2 := range c[1] {
			if it == it2 {
				return it
			}
		}
	}

	log.Fatalln("no problematic item type was found")
	return 0
}
