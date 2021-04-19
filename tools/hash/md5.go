package hash

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"github.com/keithzetterstrom/araneus/internal/repository/models"
)

func MD5(item models.Item) string {
	data := prepareData(item)
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}

func prepareData(item models.Item) []byte {
	data := bytes.Join(
		[][]byte{
			[]byte(item.Title),
			[]byte(item.Author),
			[]byte(item.Description),
			[]byte(item.PubDate),
			[]byte(item.Link),
		},
		[]byte{},
	)

	return data
}
