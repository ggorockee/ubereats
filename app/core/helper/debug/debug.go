package debug

import (
	"encoding/json"
	"log"
)

func PrintStruct(s any) {
	data, _ := json.MarshalIndent(s, "", "  ")
	log.Println(string(data))
}
