package archiver

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func ArchiveFile(dir, fileName string, wg *sync.WaitGroup) {
	defer wg.Done()
	var err error

	dir = strings.TrimSuffix(dir, "/")
	name := strings.Split(fileName, filepath.Ext(fileName))
	path := fmt.Sprintf("%s/%s_%d.tar.gz", dir, name[0], time.Now().Unix())
	out, err := os.Create(path)
	if err != nil {
		log.Fatalln(err)
	}

	defer func(f *os.File) {
		if err = f.Close(); err != nil {
			log.Fatalln(err)
		}
	}(out)

	err = createArchive(fileName, out)
	if err != nil {
		log.Fatalln(err)
	}
}
