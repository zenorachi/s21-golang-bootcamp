package reader

import (
	"encoding/xml"
)

type XMLReader struct {
	DB DataBase
}

func (xr *XMLReader) ReadDB(fileName string) (error, DataBase) {
	fileData := getFileData(fileName)
	if err := xml.Unmarshal(fileData, &xr.DB); err != nil {
		return err, xr.DB
	}
	return nil, xr.DB
}
