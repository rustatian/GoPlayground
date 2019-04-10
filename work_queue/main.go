package main

import (
	"fmt"
	"github.com/ValeryPiashchynski/Worker"
	"io/ioutil"
	"net/http"
	"runtime"
	"time"
)

func main() {
	// create a variable
	var w Worker.Work

	// example of sites to get info from in parallel
	sites := []string{"http://google.com", "http://amazon.com", "http://spiralscout.com", "https://inturn.co"}

	// add this sites to worker
	for _, v := range sites {
		w.Add(v)
	}
	time.Now()

	// Run the work in 10 goroutines (for example)
	// So, we know, that we working with strings, and we need to make type assertion
	// Each task will be handled by separate goroutine
	fmt.Println(runtime.NumGoroutine())
	w.Run(100, func(item interface{}) {
		str := item.(string)

		// we also could add work during the process of running
		w.Add("https://tut.by")
		r, err := http.Get(str)
		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
	})
	fmt.Println(runtime.NumGoroutine())

	fmt.Print("FINISH")
}
