package jsonParser

import (
	"cyoa/helpers"
	"cyoa/types"
	"encoding/json"
	"os"
)

func ReadAndParse(fp string) types.CYOAGameConfig {
	data, err := os.ReadFile(fp)
	helpers.CheckError(err)
	cyoa := parseJson(data)
	return cyoa
}

func parseJson(jsonBytes []byte) types.CYOAGameConfig {
	var target types.CYOAGameConfig
	err := json.Unmarshal(jsonBytes, &target)
	helpers.CheckError(err)
	return target
}
