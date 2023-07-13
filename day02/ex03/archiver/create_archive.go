package archiver

import (
	"archive/tar"
	"compress/gzip"
	"log"
	"os"
)

func createArchive(path string, out *os.File) error {
	free := func(g *gzip.Writer, t *tar.Writer) {
		errG := g.Close()
		errT := t.Close()
		if errG != nil {
			log.Fatal(errG)
		}
		if errT != nil {
			log.Fatal(errT)
		}
	}
	gw := gzip.NewWriter(out)
	tw := tar.NewWriter(gw)
	defer free(gw, tw)

}
