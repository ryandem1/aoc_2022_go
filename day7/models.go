package day7

// comFile represents a single file stored on an elf's communication device
type comFile struct {
	name string
	size int
}

// comDirectory represents a single directory on an elf's communication device
type comDirectory struct {
	name        string
	parent      *comDirectory
	directories []*comDirectory
	files       []*comFile
}

// totalSize will return the size of all comFile objections in a comDirectory
func (dir *comDirectory) totalSize() (totalSize int) {
	for _, file := range dir.files {
		totalSize += file.size
	}
	return totalSize
}

// comTerminal represents a single terminal session on an elf's communication device
type comTerminal struct {
	curDir *comDirectory
}

func newTerminal() *comTerminal {
	term := &comTerminal{curDir: nil}
	return term
}
