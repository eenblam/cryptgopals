package util

import (
	"os"
	"path"
)

// DataPath will try to provide an absolute path to the given
// data filename based on the current $GOPATH.
func DataPath(filename string) string {
	goPath := os.Getenv("GOPATH")
	return path.Join(goPath,
		"src/github.com/eenblam/cryptgopals/data",
		filename)
}
