package xjson

import (
	"encoding/json"
)

// StructToJson Convert the structure to a JSON string
// If it is nil, return 'null'; if it is an object, return {}; if it is an array, return []
func StructToJson(v any) string {
	if v == nil {
		return "null"
	}
	content, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(content)
}

// StructToJsonIndent Convert the structure to a JSON string with indentation
// If it is nil, return 'null'; if it is an object, return {}; if it is an array, return []
func StructToJsonIndent(v any) string {
	if v == nil {
		return "null"
	}
	content, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return ""
	}
	return string(content)
}
