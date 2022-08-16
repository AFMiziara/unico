package utils

import (
	"io/ioutil"
	"os"
)

func init() {

}

func ReadJsonFile(path string) ([]byte, error) {
	jsonFile, err := os.Open(path)
	// if we os.Open returns an error then handle it
	if err != nil {
		return []byte(""), err
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue, nil
}
