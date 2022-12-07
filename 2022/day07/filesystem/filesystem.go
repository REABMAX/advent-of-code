package filesystem

type sizer interface {
	Size() int
}

type Dir map[string]sizer

func (d Dir) Size() int {
	s := 0
	for _, v := range d {
		s += v.Size()
	}
	return s
}

func (d Dir) Dirs(f func(d Dir) bool) DirList {
	var ds []Dir
	for _, v := range d {
		if _, ok := v.(Dir); !ok {
			continue
		}

		if f(v.(Dir)) {
			ds = append(ds, v.(Dir))
		}

		ds = append(ds, v.(Dir).Dirs(f)...)
	}
	return ds
}

func (d Dir) ChooseDirToFreeSpace(maxSpace, neededSpace int) Dir {
	neededSpace -= maxSpace - d.Size()
	return d.smallestDirWithAtLeastSize(d, neededSpace)
}

func (d Dir) smallestDirWithAtLeastSize(dir Dir, n int) Dir {
	cmp := func(d Dir) bool {
		return d.Size() > n && d.Size() < dir.Size()
	}

	for _, v := range d.Dirs(cmp) {
		if cmp(v) {
			dir = v
		}

		smDir := v.smallestDirWithAtLeastSize(dir, n)
		if cmp(smDir) {
			dir = smDir
		}
	}

	return dir
}

type file int

func (f file) Size() int {
	return int(f)
}

type DirList []Dir

func (dl DirList) Size() int {
	s := 0
	for _, d := range dl {
		s += d.Size()
	}
	return s
}
