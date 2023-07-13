package archiver

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

func ArchiveFile(dir, fileName string, wg *sync.WaitGroup) {
	defer wg.Done()
	var err error

	dir = strings.TrimSuffix(dir, "/")
	path := fmt.Sprintf("%s/%s.tar.gz", dir, fileName)
	out, err := os.Create(path)
	if err != nil {
		log.Fatalln(err)
	}

	defer func(f *os.File) {
		if err = f.Close(); err != nil {
			log.Fatalln(err)
		}
	}(out)

	err = createArchive(path, out)
	if err != nil {
		log.Fatalln(err)
	}
}
