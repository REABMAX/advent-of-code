package signal

import (
	"errors"
)

func GetStartOfPacket(sig string) (int, error) {
	return detectMarker(sig, 4)
}

func GetStartOfMessage(sig string) (int, error) {
	return detectMarker(sig, 14)
}

func detectMarker(sig string, distinctChars int) (int, error) {
	signal := []rune(sig)

	for i := distinctChars; i < len(signal); i++ {
		if len(setOf(signal[i-distinctChars:i])) == distinctChars {
			return i, nil
		}
	}

	return 0, errors.New("could not detect marker")
}

func setOf[T comparable](rs []T) map[T]struct{} {
	set := make(map[T]struct{})
	for _, r := range rs {
		set[r] = struct{}{}
	}
	return set
}
