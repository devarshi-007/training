package filehandling

import (
	"io/ioutil"
	"path/filepath"
	"time"
)

var ReceiptDir string = filepath.Join("uploads")

type Receipt struct {
	ReceiptName string    `json:"name"`
	UploadData  time.Time `json:"uploadDate"`
}

func GetReceipts() ([]Receipt, error) {
	receipts := make([]Receipt, 0)
	files, err := ioutil.ReadDir(ReceiptDir)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		receipts = append(receipts, Receipt{f.Name(), f.ModTime()})
	}

	return receipts, nil
}
