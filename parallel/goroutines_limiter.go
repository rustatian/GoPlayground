package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	tokens := make(chan struct{}, 10)

	coffees, err := sweetmarias.AllCoffees()
	if err != nil {
		//return errors.Wrap(err, "sweetmarias.AllCoffees")
	}

	e := json.NewEncoder(os.Stdout)

	for _, url := range coffees {
		wg.Add(1)
		tokens <- struct{}{}
		url := url
		go func() {
			defer func() {
				<-tokens
				wg.Done()
			}()
			c, err := sweetmarias.LoadCoffee(url)
			if err != nil {
				fmt.Fprintln(os.Stderr, errors.Wrap(err, "sweetmarias.LoadCoffee"))
				return
			}
			err = e.Encode(c)
			if err != nil {
				fmt.Fprintln(os.Stderr, errors.Wrap(err, "json.Encode"))
			}
		}()
	}

	wg.Wait()
}
