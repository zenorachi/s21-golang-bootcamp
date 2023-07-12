package reader

import (
	"encoding/json"
)

type JSONReader struct {
	DB DataBase
}

func (jr *JSONReader) ReadDB(fileName string) (error, DataBase) {
	fileData := getFileData(fileName)
	if err := json.Unmarshal(fileData, &jr.DB); err != nil {
		return err, jr.DB
	}
	return nil, jr.DB
}
