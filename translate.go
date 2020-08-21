package go_trans

import (
	"io/ioutil"
	"log"
	"strings"
)

var descriptions map[string]map[string]map[string]string

type R map[string]string

type unmarshaler interface {
	Unmarshal([]byte, *map[string]map[string]string) error
}

func Initialize(path string, unmarshaler unmarshaler) error {
	descriptions := make(map[string]map[string]map[string]string)

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for _, f := range files {
		input, err := ioutil.ReadFile(path + "/" + f.Name())
		if err != nil {
			log.Println(err)
			continue
		}

		v := map[string]map[string]string{}
		err = unmarshaler.Unmarshal(input,&v)
		if err != nil {
		    log.Println(err)
		    continue
		}

		keyName := strings.Split(f.Name(), ".")
		descriptions[keyName[0]] = v
	}
	return nil
}

func Trans(key string, local string, replace ...R) string {
	keys := strings.Split(key, ".")
	txt, isSet := descriptions[keys[0]][keys[1]][local]

	if isSet == false {
		return key
	}

	if len(replace) < 1 {
		return txt
	}

	for key, value := range replace[0] {
		txt = strings.ReplaceAll(txt, ":"+key, value)
	}

	return txt
}
