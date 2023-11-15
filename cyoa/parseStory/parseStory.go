package parseStory

import (
	"encoding/json"
	"io/ioutil"
)

type OptionType struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

type Chapter struct {
	Title      string       `json:"title"`
	Paragraphs []string     `json:"story"`
	Options    []OptionType `json:"options"`
}

type Story map[string]Chapter

func JSON(fileName string) (Story, error) {
	var fileData Story
	dataFile, err := ioutil.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(dataFile, &fileData); err != nil {
		return nil, err
	}

	return fileData, nil
}
