package rucksack

import (
	"fmt"
	"testing"
)

func TestCompartmentsFromString(t *testing.T) {
	s := "vJrwpWtwJgWrhcsFMMfFFhFp"
	c := CompartmentsFromString(s)

	expectFirst := "vJrwpWtwJgWr"
	expectSecond := "hcsFMMfFFhFp"
	if string(c[0]) != expectFirst {
		t.Log(fmt.Sprintf("first compartment has to be %s, is %s instead", expectFirst, string(c[0])))
		t.Fail()
	}
	if string(c[1]) != expectSecond {
		t.Log(fmt.Sprintf("first compartment has to be %s, is %s instead", expectSecond, string(c[0])))
		t.Fail()
	}
}

func TestCompartments_FindProblematicItemType(t *testing.T) {
	c := Compartments{
		[]ItemType("vJrwpWtwJgWr"),
		[]ItemType("hcsFMMfFFhFp"),
	}
	it := c.FindProblematicItemType()
	exp := ItemType('p')
	if it != exp {
		t.Log(fmt.Sprintf("problematic item type should be %v, is %v instead", exp, it))
		t.Fail()
	}
}

func TestItemType_Priority(t *testing.T) {
	types := []ItemType("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	for i := 0; i < 52; i++ {
		p := i + 1
		it := types[i]
		if it.Priority() != p {
			t.Log(fmt.Sprintf("priority of item type %v should be %d, is %d instead", it, p, it.Priority()))
			t.Fail()
		}
	}
}

func TestGroupFromStrings(t *testing.T) {
	g := GroupFromStrings("vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg")
	if string(g[0]) != "vJrwpWtwJgWrhcsFMMfFFhFp" {
		t.Fail()
	}
	if string(g[1]) != "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL" {
		t.Fail()
	}
	if string(g[2]) != "PmmdzqPrVvPwwTWBwg" {
		t.Fail()
	}
}

func TestGroup_FindUnifyingItemType(t *testing.T) {
	g := GroupFromStrings("vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg")
	exp := ItemType('r')
	it := g.FindUnifyingItemType()
	if it != exp {
		t.Log(fmt.Sprintf("unifying item type should be %v, is %v instead", exp, it))
		t.Fail()
	}

	g = GroupFromStrings("wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw")
	exp = ItemType('Z')
	it = g.FindUnifyingItemType()
	if it != exp {
		t.Log(fmt.Sprintf("unifying item type should be %v, is %v instead", exp, it))
		t.Fail()
	}
}
