package signal

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Register struct {
	x      int
	cycles []int
}

func NewRegister(x int) Register {
	return Register{x, []int{}}
}

var DefaultRegister = NewRegister(1)

func (r *Register) ExecuteProgram(f io.Reader) error {
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		if line == "noop" {
			r.Noop()
		} else {
			parts := strings.Split(line, " ")
			num, err := strconv.Atoi(parts[1])
			if err != nil {
				return err
			}
			r.Addx(num)
		}
	}

	return nil
}

func (r *Register) DrawCRT() string {
	crt := ""
	for c, x := range r.cycles {
		if c%40 == 0 && c != 0 {
			crt += "\n"
		}
		if c%40 == x-1 || c%40 == x || c%40 == x+1 {
			crt += "#"
		} else {
			crt += "."
		}
	}
	return crt
}

func (r *Register) Noop() {
	r.cycles = append(r.cycles, r.x)
}

func (r *Register) Addx(n int) {
	r.cycles = append(r.cycles, r.x)
	r.cycles = append(r.cycles, r.x)
	r.x += n
}

func (r Register) X() int                          { return r.x }
func (r Register) Cycles() int                     { return len(r.cycles) }
func (r Register) SignalStrengthAtCycle(c int) int { return c * r.cycles[c-1] }
