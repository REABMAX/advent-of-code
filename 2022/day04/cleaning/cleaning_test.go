package cleaning

import (
	"fmt"
	"testing"
)

func TestPairFromString(t *testing.T) {
	s := "12-24,13-13"
	_, err := PairFromString(s)
	if err != nil {
		t.Log(fmt.Sprintf("err: %v", err))
		t.Fail()
	}

	s = "12-13,13-14,12"
	_, err = PairFromString(s)
	if err == nil {
		t.Log(fmt.Sprintf("err expected, for %s no err given", s))
		t.Fail()
	}

	s = "12-14"
	_, err = PairFromString(s)
	if err == nil {
		t.Log(fmt.Sprintf("err expected, for %s no err given", s))
		t.Fail()
	}

	s = "12-4,13-14"
	_, err = PairFromString(s)
	if err == nil {
		t.Log(fmt.Sprintf("err expected, for %s no err given", s))
		t.Fail()
	}
}

func TestPair_ContainEachOther(t *testing.T) {
	s := "12-24,13-13"
	p, _ := PairFromString(s)
	if !p.ContainEachOther() {
		t.Log(fmt.Sprintf("expected pairs to contain each other"))
		t.Fail()
	}

	s = "12-15,13-14"
	p, _ = PairFromString(s)
	if !p.ContainEachOther() {
		t.Log(fmt.Sprintf("expected pairs to contain each other"))
		t.Fail()
	}

	s = "12-14,24-25"
	p, _ = PairFromString(s)
	if p.ContainEachOther() {
		t.Log(fmt.Sprintf("didnt expect pairs to contain each other"))
		t.Fail()
	}
}

func TestPair_OverlapEachOther(t *testing.T) {
	s := "2-4,6-8"
	p, _ := PairFromString(s)
	if p.OverlapEachOther() {
		t.Log(fmt.Sprintf("%s is not expected to overlap", s))
		t.Fail()
	}

	s = "2-3,4-5"
	p, _ = PairFromString(s)
	if p.OverlapEachOther() {
		t.Log(fmt.Sprintf("%s is not expected to overlap", s))
		t.Fail()
	}

	s = "5-7,7-9"
	p, _ = PairFromString(s)
	if !p.OverlapEachOther() {
		t.Log(fmt.Sprintf("%s is expected to overlap", s))
		t.Fail()
	}

	s = "2-8,3-7"
	p, _ = PairFromString(s)
	if !p.OverlapEachOther() {
		t.Log(fmt.Sprintf("%s is expected to overlap", s))
		t.Fail()
	}

	s = "6-6,4-6"
	p, _ = PairFromString(s)
	if !p.OverlapEachOther() {
		t.Log(fmt.Sprintf("%s is expected to overlap", s))
		t.Fail()
	}

	s = "2-6,4-8"
	p, _ = PairFromString(s)
	if !p.OverlapEachOther() {
		t.Log(fmt.Sprintf("%s is expected to overlap", s))
		t.Fail()
	}
}
