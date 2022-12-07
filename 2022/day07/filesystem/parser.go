package filesystem

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func DirsFromSlice(s []string) (Dir, error) {
	var err error
	dirs := make(Dir)
	dirs["/"] = make(Dir)

	dirPointer := dirs["/"].(Dir)
	dirPointerPath := []string{"/"}

	for i := 0; i < len(s); i++ {
		l := s[i]
		parts := strings.Split(l, " ")
		if isCommand(l) {
			cmd := parts[1]
			switch cmd {
			case "cd":
				if len(parts) < 3 {
					return nil, errors.New("cd command expects two args")
				}
				arg := parts[2]
				switch arg {
				case "/":
					dirPointer = dirs["/"].(Dir)
					dirPointerPath = []string{"/"}
				case "..":
					dirPointerPath = dirPointerPath[:len(dirPointerPath)-1]
					dirPointer, err = getDirByPath(dirs, dirPointerPath)
					if err != nil {
						return nil, err
					}
				default:
					dirPointerPath = append(dirPointerPath, arg)
					dirPointer, err = getDirByPath(dirs, dirPointerPath)
					if err != nil {
						return nil, err
					}
				}
			case "ls":
				for i+1 < len(s) && !isCommand(s[i+1]) {
					i++
					l = s[i]
					parts := strings.Split(l, " ")
					switch parts[0] {
					case "dir":
						if _, ok := dirPointer[parts[1]]; !ok {
							dirPointer[parts[1]] = make(Dir)
						}
					default:
						size, err := strconv.Atoi(parts[0])
						if err != nil {
							return nil, err
						}
						dirPointer[parts[1]] = file(size)
					}
				}
			}
		}
	}

	return dirs, nil
}

func isCommand(s string) bool {
	rs := []rune(s)
	return rs[0] == '$'
}

func getDirByPath(dir Dir, p []string) (Dir, error) {
	if _, ok := dir[p[0]].(Dir); !ok {
		return nil, errors.New(fmt.Sprintf("corrupt dir pointer path: %v is not a directory", p[0]))
	}

	ptr := dir[p[0]].(Dir)
	path := p[1:]
	for _, v := range path {
		if _, ok := ptr[v].(Dir); !ok {
			return nil, errors.New(fmt.Sprintf("corrupt dir pointer path: %v is not a directory", v))
		}
		ptr = ptr[v].(Dir)
	}
	return ptr, nil
}
