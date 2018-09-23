package http

import (
	"io/ioutil"
	"testing"
)

func readJsonFile(t *testing.T, filePath string) []byte {
	objectBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Error on read file [%s]. %s", filePath, err)
	}

	return objectBytes
}
