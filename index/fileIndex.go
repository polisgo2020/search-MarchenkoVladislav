package index

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"polisgomarchenko/utils"
	"regexp"
	"strings"
)

func GetInvertedIndexMap(dir string) (map[string][]string, error)  {
	indexMap := make(map[string][]string)
	var re = regexp.MustCompile(`[[:punct:]]`)

	inputFiles, err := ioutil.ReadDir(dir)

	if err != nil {
		return nil, err
	}

	for _, fileInfo := range inputFiles {
		currentFile, err := ioutil.ReadFile(filepath.Join(dir, fileInfo.Name()))

		if err != nil {
			return nil,err
		}

		words := strings.Fields(re.ReplaceAllString(string(currentFile), ""))

		for _, word := range words {

			indexMap[strings.ToLower(word)] = utils.AppendIfMissing(indexMap[strings.ToLower(word)], fileInfo.Name())
		}
	}

	return indexMap, nil
}

func WriteInvertedIndexIntoFile(dir string, outputFileName string) error {

	indexMap, err := GetInvertedIndexMap(dir)

	if err != nil {
		return err
	}

	file, err := os.Create(outputFileName)

	if err != nil{
		return err
	}

	defer file.Close()

	res, err := json.Marshal(indexMap)

	if err != nil{
		return err
	}

	if _, err = file.Write(res); err != nil {
		return err
	}

	return nil
}