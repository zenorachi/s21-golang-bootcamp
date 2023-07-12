package walkfn

import (
	"fmt"
	"os"
	"path/filepath"
)

func WalkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if s, e := filepath.EvalSymlinks(path); e == nil && s != path {
		fmt.Println(path, "->", s)
	} else {
		fmt.Println(path)
	}

	return nil
}
