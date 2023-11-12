package gogeta_json

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

func makePublicMember(key string) string {
	if key == "" {
		return key
	} // Since strings are immutable in Go we cannot directly
	// modify the characters of a string. Instead we first convert
	// them to a slice of runes ([]int32) do the modification and
	// then convert back to a string
	runes := []rune(key)
	// Capitalize the first character of the key
	// to make it a public member.
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func generate(jsonData map[string]interface{}, sb *strings.Builder, tabSpaces int) {
	sb.WriteString("struct {")
	sb.WriteByte('\n')
	for key, val := range jsonData {
		// Every new line is indented with tab space(s).
		for i := 0; i < tabSpaces; i++ {
			sb.WriteByte('\t')
		}
		sb.WriteString(makePublicMember(key))
		sb.WriteByte('\t')
		var valType string
		// This check is for nested JSON objects.
		if nestedJSON, ok := val.(map[string]interface{}); ok {
			generate(nestedJSON, sb, tabSpaces+1)
			sb.WriteString(fmt.Sprintf("\t`json:\"%s\"`\n", key))
			continue
		}
		valType = fmt.Sprintf("%T", val)
		// TODO: Appropriate handling of "null" from JSON is required.
		if valType == "<nil>" {
			valType = "bool"
		}
		// It is for now expected that the JSON contains only
		// homogenous arrays i.e. [1, 2, 3] or ["a", "b", "c"]
		// and not ["a", 1.32, "b"].
		if array, ok := val.([]interface{}); ok {
			// This check expects the array to contain homogenous nested
			// JSON objects.
			if nestedJSON, ok := array[0].(map[string]interface{}); ok {
				sb.WriteString("[]")
				generate(nestedJSON, sb, tabSpaces+1)
				sb.WriteString(fmt.Sprintf("\t`json:\"%s\"`\n", key))
				continue
			}
			valType = fmt.Sprintf("[]%T", array[0])
		}
		sb.WriteString(valType)
		sb.WriteString(fmt.Sprintf("\t`json:\"%s\"`", key))
		sb.WriteByte('\n')
	}
	for i := 0; i < tabSpaces-1; i++ {
		sb.WriteByte('\t')
	}
	sb.WriteByte('}')
}

func GenerateGoStruct(jsonFile string) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	file, err := os.Open(strings.Join([]string{cwd, jsonFile}, "/"))
	if err != nil {
		return "", err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	var jsonData map[string]interface{}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		return "", err
	}

	sb := strings.Builder{}
	sb.WriteString("type GogetaGenerated ")
	generate(jsonData, &sb, 1)
	return sb.String(), nil
}
