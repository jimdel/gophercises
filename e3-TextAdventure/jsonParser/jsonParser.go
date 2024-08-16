package jsonParser

import (
	"cyoa/types"
	"encoding/json"
	"log"
	"os"
)

func ReadAndParse(fp string) types.CYOAGameConfig {
	data, err := os.ReadFile(fp)
	if err != nil {
		log.Fatal(err)
	}
	cyoa := parseJson(data)
	return cyoa
}

func parseJson(jsonBytes []byte) types.CYOAGameConfig {
	var target types.CYOAGameConfig
	json.Unmarshal(jsonBytes, &target)
	return target
}
