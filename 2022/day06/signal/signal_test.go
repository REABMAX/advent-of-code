package signal

import (
	"fmt"
	"testing"
)

func TestGetStartOfPacket(t *testing.T) {
	sig := "abcdefgh"
	exp := 4
	sop, err := GetStartOfPacket(sig)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if sop != exp {
		t.Log(fmt.Sprintf("start-of-packet was expected to be %d, was %d instead", exp, sop))
		t.Fail()
	}

	sig = "bvwbjplbgvbhsrlpgdmjqwftvncz"
	exp = 5
	sop, err = GetStartOfPacket(sig)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if sop != exp {
		t.Log(fmt.Sprintf("start-of-packet was expected to be %d, was %d instead", exp, sop))
		t.Fail()
	}

	sig = "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"
	exp = 10
	sop, err = GetStartOfPacket(sig)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if sop != exp {
		t.Log(fmt.Sprintf("start-of-packet was expected to be %d, was %d instead", exp, sop))
		t.Fail()
	}
}

func TestGetStartOfMessage(t *testing.T) {
	sig := "mjqjpqmgbljsphdztnvjfqwrcgsmlb"
	exp := 19
	sop, err := GetStartOfMessage(sig)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if sop != exp {
		t.Log(fmt.Sprintf("start-of-message was expected to be %d, was %d instead", exp, sop))
		t.Fail()
	}

	sig = "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"
	exp = 29
	sop, err = GetStartOfMessage(sig)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if sop != exp {
		t.Log(fmt.Sprintf("start-of-message was expected to be %d, was %d instead", exp, sop))
		t.Fail()
	}
}
