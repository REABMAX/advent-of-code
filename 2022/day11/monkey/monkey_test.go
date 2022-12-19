package monkey

import (
	"fmt"
	"testing"
)

func TestMonkey_Inspect_DivideWorryLevel(t *testing.T) {
	m := Monkey{
		items:        []Item{22, 33},
		inspectCount: 0,
		test: test{
			2,
			nil,
			nil,
		},
		operation: func(old Item) Item {
			return old * 2
		},
	}

	m.Inspect(true, 1)

	if m.items[0] != 14 {
		t.Log(fmt.Sprintf("expected first monkey item to be 14 after inspection, was %d instead", m.items[0]))
		t.Fail()
	}

	if m.inspectCount != 1 {
		t.Log("didnt increase inspect count")
		t.Fail()
	}
}

func TestMonkey_Inspect_NotDivideWorryLevel(t *testing.T) {
	m := Monkey{
		items:        []Item{22, 33},
		inspectCount: 0,
		test: test{
			2,
			nil,
			nil,
		},
		operation: func(old Item) Item {
			return old * 2
		},
	}

	m.Inspect(false, 1)

	if m.items[0] != 44 {
		t.Log(fmt.Sprintf("expected first monkey item to be 44 after inspection, was %d instead", m.items[0]))
		t.Fail()
	}

	if m.inspectCount != 1 {
		t.Log("didnt increase inspect count")
		t.Fail()
	}
}

func TestMonkey_Throw(t *testing.T) {
	var m1, m2, m3 Monkey
	m1 = Monkey{
		items:        []Item{5, 33},
		inspectCount: 0,
		test: test{
			5,
			&m3,
			&m2,
		},
		operation: func(old Item) Item {
			return old * 5
		},
	}
	m2 = Monkey{
		items:        []Item{10, 23},
		inspectCount: 0,
		test: test{
			10,
			&m1,
			&m3,
		},
		operation: func(old Item) Item {
			return old + 1
		},
	}
	m3 = Monkey{
		items:        []Item{20, 54},
		inspectCount: 0,
		test: test{
			20,
			&m2,
			&m1,
		},
		operation: func(old Item) Item {
			return old * 2
		},
	}

	m3.Throw()

	if m3.items[0] == 20 {
		t.Log("didnt throw away the item")
		t.Fail()
	}

	if m2.items[len(m2.items)-1] != 20 {
		t.Log("didnt throw item to the correct monkey")
		t.Fail()
	}
}

func TestPlayRound(t *testing.T) {
	var m0, m1, m2, m3 Monkey
	m0 = Monkey{
		items: []Item{79, 98},
		operation: func(old Item) Item {
			return old * 19
		},
		test: test{
			v:           23,
			trueMonkey:  &m2,
			falseMonkey: &m3,
		},
	}
	m1 = Monkey{
		items: []Item{54, 65, 75, 74},
		operation: func(old Item) Item {
			return old + 6
		},
		test: test{
			v:           19,
			trueMonkey:  &m2,
			falseMonkey: &m0,
		},
	}
	m2 = Monkey{
		items: []Item{79, 60, 97},
		operation: func(old Item) Item {
			return old * old
		},
		test: test{
			v:           13,
			trueMonkey:  &m1,
			falseMonkey: &m3,
		},
	}
	m3 = Monkey{
		items: []Item{74},
		operation: func(old Item) Item {
			return old + 3
		},
		test: test{
			v:           17,
			trueMonkey:  &m0,
			falseMonkey: &m1,
		},
	}

	PlayRound(true, &m0, &m1, &m2, &m3)
	if len(m2.items) > 0 || len(m3.items) > 0 {
		t.Log("didnt expect m2 and m3 to have items")
		t.Fail()
	}
	if len(m0.items) != 4 {
		t.Log(fmt.Sprintf("expected m0 to have exactly 4 items, was %d instead", len(m0.items)))
		t.Fail()
	}
	if len(m1.items) != 6 {
		t.Log(fmt.Sprintf("expected m1 to have exactly 6 items, was %d instead", len(m1.items)))
		t.Fail()
	}
}
