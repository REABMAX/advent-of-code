package monkey

import (
	"errors"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func Parse(s string) ([]*Monkey, error) {
	parts := strings.Split(s, "\n\n")
	var monkeys []*Monkey
	for _, p := range parts {
		lines := strings.Split(p, "\n")
		m, err := parseMonkey(lines)
		if err != nil {
			return nil, err
		}
		monkeys = append(monkeys, &m)
	}

	for i, p := range parts {
		lines := strings.Split(p, "\n")
		test, err := parseTest(lines[3], lines[4], lines[5], monkeys)
		if err != nil {
			return nil, err
		}
		monkeys[i].test = test
	}

	return monkeys, nil
}

func parseMonkey(lines []string) (Monkey, error) {
	items := parseItems(lines[1])
	operation, err := parseOperation(lines[2])
	if err != nil {
		return Monkey{}, err
	}
	return Monkey{
		items:        items,
		operation:    operation,
		inspectCount: 0,
	}, nil
}

func parseItems(s string) []Item {
	reg := regexp.MustCompile("\\d+")
	parsed := reg.FindAllString(s, -1)
	var items []Item
	for _, i := range parsed {
		it, _ := strconv.Atoi(i)
		items = append(items, Item(it))
	}
	return items
}

func parseOperation(s string) (operation, error) {
	reg := regexp.MustCompile("([\\*|\\+]) (\\d+|old)")
	parsed := reg.FindStringSubmatch(s)
	if len(parsed) < 3 {
		return nil, errors.New("could not parse operation")
	}

	var num Item
	if parsed[2] != "old" {
		n, err := strconv.Atoi(parsed[2])
		if err != nil {
			return nil, err
		}
		num = Item(n)
	}

	return func(old Item) Item {
		if parsed[2] == "old" {
			num = old
		}

		switch parsed[1] {
		case "*":
			return old * num
		case "+":
			return old + num
		default:
			log.Fatalln("no valid operator")
			return 0
		}
	}, nil
}

func parseTest(tests, trues, falses string, monkeys []*Monkey) (test, error) {
	reg := regexp.MustCompile("\\d+")
	conds := reg.FindString(tests)
	cond, err := strconv.Atoi(conds)
	if err != nil {
		return test{}, err
	}

	s := reg.FindString(trues)
	monkeyNum, err := strconv.Atoi(s)
	if err != nil {
		return test{}, err
	}
	trueMonkey := monkeys[monkeyNum]

	s = reg.FindString(falses)
	monkeyNum, err = strconv.Atoi(s)
	if err != nil {
		return test{}, err
	}
	falseMonkey := monkeys[monkeyNum]

	return test{
		v:           uint64(cond),
		trueMonkey:  trueMonkey,
		falseMonkey: falseMonkey,
	}, nil
}
