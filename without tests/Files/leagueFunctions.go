package files

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func PlayerDataFromFiles() (*PlayerDatabase, func(), error) {
	db, err := os.OpenFile("PlayerData.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, nil, err
	}
	closeFile := func() {
		db.Close()
	}
	store, err := NewPlayerData(db)
	if err != nil {
		return nil, nil, fmt.Errorf("problem creating file system player store, %v", err)
	}

	return store,
		closeFile,
		nil
}
func NewPlayerData(file *os.File) (*PlayerDatabase, error) {
	err := isFileNew(file)
	if err != nil {
		return nil, fmt.Errorf("problem loading players info %v", err)
	}
	league, err := JsonFileToLeague(file)
	if err != nil {
		log.Fatal(err)
	}
	return &PlayerDatabase{
		jsonFile: json.NewEncoder(&Tape{File: file}),
		league:   league,
	}, nil
}

func isFileNew(file *os.File) error {
	file.Seek(0, io.SeekStart)

	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("problem getting file info %v", err)
	}

	if fileInfo.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, io.SeekStart)
	}
	return nil
}

func JsonFileToLeague(file io.Reader) (League, error) {
	var league League
	err := json.NewDecoder(file).Decode(&league)
	if err != nil {
		err = fmt.Errorf("error decoding file %v", err)
	}
	return league, err
}
