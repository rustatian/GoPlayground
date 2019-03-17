package main

import (
	"encoding/json"
)

var aa string = `
{
  "http://xxx.xxx.xxx.xxx/v1/document": {
    "file_ext": {
      "TXT": "5-50, Mb, 60",
      "JPEG": "10, Kb, 10",
      "XML": "100, Kb, 5"
    },
    "duration": "20s",
    "request_type": "multipart/form-data"
  },
  "http://xxx.xxx.xxx.xxx/v1/document2": {
    "file_ext": {
      "TXT": "5-50, Mb, 60",
      "JPEG": "10, Kb, 10",
      "XML": "100, Kb, 5"
    },
    "duration": "20s",
    "request_type": "multipart/form-data"
  }
}
`

func main() {
	//var buf bytes.Buffer
	//enc := gob.NewEncoder(&buf)
	//err := enc.Encode(aa)
	//if err != nil {
	//	panic(err)
	//}

	var ii map[string]interface{}
	err := json.Unmarshal([]byte(aa), &ii)
	if err != nil {
		panic(err)
	}

	for k, _ := range ii {
		print(k)
	}

}
