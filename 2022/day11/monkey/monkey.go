package monkey

func PlayRound(divideWorryLevel bool, monkeys ...*Monkey) {
	generator := uint64(1)
	for _, m := range monkeys {
		generator *= m.test.v
	}

	for _, m := range monkeys {
		for range m.items {
			m.Inspect(divideWorryLevel, generator)
			m.Throw()
		}
	}
}

type Item uint64

type operation func(old Item) Item

type test struct {
	v           uint64
	trueMonkey  *Monkey
	falseMonkey *Monkey
}

func (t test) test(i Item) bool {
	return uint64(i)%t.v == 0
}

type Monkey struct {
	items        []Item
	operation    operation
	test         test
	inspectCount int
}

func (m Monkey) InspectCount() int { return m.inspectCount }

func (m *Monkey) Inspect(divideWorryLevel bool, generator uint64) {
	i := m.items[0]
	m.inspectCount++
	i = m.operation(i) % Item(generator)
	if divideWorryLevel {
		i = i / 3
	}
	m.items[0] = i
}

func (m *Monkey) Throw() {
	i := m.items[0]
	m.items = m.items[1:]
	if m.test.test(i) {
		m.test.trueMonkey.catch(i)
	} else {
		m.test.falseMonkey.catch(i)
	}
}

func (m *Monkey) catch(i Item) {
	m.items = append(m.items, i)
}
