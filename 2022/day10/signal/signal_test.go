package signal

import (
	"fmt"
	"strings"
	"testing"
)

func TestRegister_SignalStrengthAtCycle(t *testing.T) {
	reg := NewRegister(1)
	reg.Noop()
	reg.Addx(3)
	reg.Addx(-5)
	reg.Noop()

	if reg.SignalStrengthAtCycle(1) != 1 {
		t.Log(fmt.Sprintf("signal strength after cycle 1 should be 1, was %d instead", reg.SignalStrengthAtCycle(1)))
		t.Fail()
	}

	if reg.SignalStrengthAtCycle(3) != 3 {
		t.Log(fmt.Sprintf("signal strength after cycle 3 should be 3, was %d instead", reg.SignalStrengthAtCycle(3)))
		t.Fail()
	}

	if reg.SignalStrengthAtCycle(4) != 16 {
		t.Log(fmt.Sprintf("signal strength after cycle 4 should be 16, was %d instead", reg.SignalStrengthAtCycle(4)))
		t.Fail()
	}

	if reg.SignalStrengthAtCycle(6) != -6 {
		t.Log(fmt.Sprintf("signal strength after cycle 6 should be -6, was %d instead", reg.SignalStrengthAtCycle(6)))
		t.Fail()
	}
}

func TestRegister_ExecuteProgram(t *testing.T) {
	i := `addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop`
	reg := NewRegister(1)
	_ = reg.ExecuteProgram(strings.NewReader(i))

	exp := map[int]int{
		20:  420,
		60:  1140,
		100: 1800,
		140: 2940,
		180: 2880,
		220: 3960,
	}

	for i, v := range exp {
		if reg.SignalStrengthAtCycle(i) != v {
			t.Log(fmt.Sprintf("signal strength after cycle %d should be %d, was %d instead", i, v, reg.SignalStrengthAtCycle(i)))
			t.Fail()
		}
	}
}

func TestRegister_DrawCRT(t *testing.T) {
	i := `addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop`
	reg := NewRegister(1)
	_ = reg.ExecuteProgram(strings.NewReader(i))
	crt := reg.DrawCRT()
	exp := "##..##..##..##..##..##..##..##..##..##..\n" +
		"###...###...###...###...###...###...###.\n" +
		"####....####....####....####....####....\n" +
		"#####.....#####.....#####.....#####.....\n" +
		"######......######......######......####\n" +
		"#######.......#######.......#######....."
	if crt != exp {
		t.Log("Expected\n" + exp + "\nGot\n" + crt)
		t.Fail()
	}
}
