package k8s

import (
	"encoding/json"
	"fmt"
	"go.etcd.io/bbolt"
	"log"
	"os"
	"testing"
	"unicode"
)

// /var/lib/docker/buildkit下面有containerdmeta.db、metadata_v2.db等db文件

func TestCatDbInfo(t *testing.T) {
	fileName := "local-kv.db"
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		// check the existence of the file, so it won't be created automatically
		fmt.Println(err)
		return
	}

	db, err := bbolt.Open(fileName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer func(db *bbolt.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	fmt.Printf("%s\n", getContentAsJson(db))
}

func getContentAsJson(db *bbolt.DB) string {
	content := make(map[string]map[string]string)
	err := db.View(func(tx *bbolt.Tx) error {
		return tx.ForEach(func(nameAsBytes []byte, _ *bbolt.Bucket) error {
			name := string(nameAsBytes)
			content[name] = getKeyValuePairs(db, name)
			return nil
		})
	})
	if err != nil {
		return ""
	}

	contentAsJson, _ := json.MarshalIndent(content, "", "  ")
	return string(contentAsJson)
}

func getKeyValuePairs(db *bbolt.DB, bucketName string) map[string]string {
	keyValuePairs := make(map[string]string)
	err := db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		cursor := bucket.Cursor()

		for key, value := cursor.First(); key != nil; key, value = cursor.Next() {
			dbValue := string(removeNonPrintableLeadingBytes(value))
			keyValuePairs[string(key)] = dbValue
		}

		return nil
	})
	if err != nil {
		return nil
	}
	return keyValuePairs
}

func removeNonPrintableLeadingBytes(inputBytes []byte) []byte {
	var newBytes []byte
	var i int
	for ; i < len(inputBytes); i++ {
		if unicode.IsPrint(rune(inputBytes[i])) {
			break
		}
	}
	jsonStartIndex := i
	if jsonStartIndex >= 0 {
		newBytes = inputBytes[jsonStartIndex:]
	}
	return newBytes
}
