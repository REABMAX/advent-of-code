package cleaning

import (
	"errors"
	"strconv"
	"strings"
)

type Pair [2]sections

func PairFromString(s string) (Pair, error) {
	pairs := strings.Split(s, ",")
	if len(pairs) != 2 {
		return Pair{}, errors.New("wrong input format")
	}

	sects1, err := sectionsFromString(pairs[0])
	if err != nil {
		return Pair{}, err
	}
	sects2, err := sectionsFromString(pairs[1])
	if err != nil {
		return Pair{}, err
	}

	return Pair{sects1, sects2}, nil
}

func (p Pair) ContainEachOther() bool {
	return p[0].contains(p[1]) || p[1].contains(p[0])
}

func (p Pair) OverlapEachOther() bool {
	return p[0].overlaps(p[1]) || p[1].overlaps(p[0])
}

// first is min, second is max
type sections [2]int

func sectionsFromString(s string) (sections, error) {
	sects := strings.Split(s, "-")
	if len(sects) != 2 {
		return sections{}, errors.New("wrong input format")
	}

	min, err := strconv.Atoi(sects[0])
	if err != nil {
		return sections{}, err
	}
	max, err := strconv.Atoi(sects[1])
	if err != nil {
		return sections{}, err
	}

	if min > max {
		return sections{}, errors.New("first number in sections range should be less or equal than the second number")
	}

	return sections{min, max}, nil
}

func (s sections) min() int {
	return s[0]
}

func (s sections) max() int {
	return s[1]
}

func (s sections) contains(os sections) bool {
	return s.min() <= os.min() && s.max() >= os.max()
}

func (s sections) overlaps(os sections) bool {
	if s.min() >= os.min() && s.min() <= os.max() {
		return true
	}
	if s.max() >= os.min() && s.max() <= os.max() {
		return true
	}
	return false
}
