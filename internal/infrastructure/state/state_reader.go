package state

import (
	"github.com/MRushtonScratch/terrahawk/internal/domain/entity"
	"io/ioutil"
	"encoding/json"
	"errors"
)

func check (e error) {
	if e != nil{
		panic(e)
	}
}

func Detach (filename string) (entity.State, error) {

	buf, err := ioutil.ReadFile(filename)

	check(err)

	var state entity.State
	err = json.Unmarshal(buf, &state)


	if err != nil {
		return state, errors.New("There was an error serialising Json from file")
	}

	return state, nil
}
