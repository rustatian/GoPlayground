package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	type Queue struct {
		Name string `json:name`
		VHost string `json:vhost`
	}

	manager := "http://127.0.0.1:15672/api/queues/"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", manager, nil)
	req.SetBasicAuth("guest", "guest")
	resp, _ := client.Do(req)

	value := make([]Queue, 0)
	json.NewDecoder(resp.Body).Decode(&value)
	fmt.Println(value)

}
