package filesystem

import (
	"fmt"
	"log"
	"testing"
)

var expDirs = Dir{
	"/": Dir{
		"a": Dir{
			"e": Dir{
				"i": file(584),
			},
			"f":     file(29116),
			"g":     file(2557),
			"h.lst": file(62596),
		},
		"b.txt": file(14848514),
		"c.dat": file(8504156),
		"d": Dir{
			"j":     file(4060174),
			"d.log": file(8033020),
			"d.ext": file(5626152),
			"k":     file(7214296),
		},
	},
}

func TestDir_Size(t *testing.T) {
	d := Dir{
		"a": Dir{
			"e": Dir{
				"i": file(584),
			},
			"f":     file(29116),
			"g":     file(2557),
			"h.lst": file(62596),
		},
	}
	exp := 94853
	s := d.Size()
	if s != exp {
		t.Log(fmt.Sprintf("expected size to be %d, was %d instead", exp, s))
		t.Fail()
	}
}

func TestDir_DirsWithTotalSizeOfMost(t *testing.T) {
	d := expDirs
	exp := DirList{d["/"].(Dir)["a"].(Dir), d["/"].(Dir)["a"].(Dir)["e"].(Dir)}
	res := d.Dirs(func(d Dir) bool {
		return d.Size() <= 100_000
	})
	for i, v := range exp {
		log.Println(res)
		if !assertDirs(v, res[i]) {
			t.Fail()
		}
	}
}

func TestDir_ChooseDirToFreeSpace(t *testing.T) {
	d := expDirs
	exp := d["/"].(Dir)["d"].(Dir)
	res := d.ChooseDirToFreeSpace(70_000_000, 30_000_000)
	log.Println(res)
	if !assertDirs(exp, res) {
		t.Fail()
	}
}

func TestDirList_Size(t *testing.T) {
	dl := DirList{expDirs["/"].(Dir)["a"].(Dir), expDirs["/"].(Dir)["a"].(Dir)["e"].(Dir)}
	exp := 95437
	if dl.Size() != exp {
		t.Log(fmt.Sprintf("expected dirlist size to be %d, was %d instead", exp, dl.Size()))
		t.Fail()
	}
}

func assertDirs(dirs Dir, exp Dir) bool {
	for k, d := range exp {
		if _, ok := d.(file); ok {
			log.Println(fmt.Sprintf("assert file %v == %v", dirs[k], d))
			if dirs[k] != d {
				return false
			}
		} else {
			log.Println(fmt.Sprintf("assert Dir %v", d))
			return assertDirs(d.(Dir), dirs[k].(Dir))
		}
	}
	return true
}
