package archiver

import (
	"archive/tar"
	"compress/gzip"
	"io"
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

	err := addData(tw, path)
	if err != nil {
		return err
	}
	return nil
}

func addData(tw *tar.Writer, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		if e := f.Close(); e != nil {
			log.Fatalln(e)
		}
	}(file)

	info, err := os.Stat(path)
	if err != nil {
		return err
	}

	header, err := tar.FileInfoHeader(info, info.Name())
	if err != nil {
		return err
	}

	header.Name = path
	err = tw.WriteHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(tw, file)
	if err != nil {
		return err
	}

	return nil
}
