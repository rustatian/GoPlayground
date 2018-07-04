package main

import (
	"bitbucket.org/inturnco/go-sdk/helpers"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Metrics interface {
	StartCollectingMetrics()
}

type metrics struct {
	sync.Mutex
	s chan os.Signal

	t  time.Duration
	dt time.Duration

	p string
}

func NewMetricsExporter(s chan os.Signal, scrapeInterval time.Duration, pathClearingInterval time.Duration, path string) Metrics {
	CreateDirIfNotExist(path)

	return &metrics{
		s:  s,
		t:  scrapeInterval,
		dt: pathClearingInterval,
		p:  path,
	}
}

func (m *metrics) StartCollectingMetrics() {
	t := time.NewTimer(m.t)
	dt := time.NewTimer(m.dt)

	for {
		select {
		case <-t.C:
			m.Lock()

			response, err := http.Get("http://localhost:5685/metrics")
			if err != nil {
				m.Unlock()
				fmt.Printf("%s", err)
				os.Exit(1)
			} else {
				contents, err := ioutil.ReadAll(response.Body)
				if err != nil {
					m.Unlock()
					fmt.Printf("%s", err)
					os.Exit(1)
				}
				fmt.Printf("%s\n", string(contents))

				uuid, _ := helpers.GenerateUUID()
				filename := uuid + "_metrics.prom"

				err = ioutil.WriteFile(m.p+filename, contents, 0644)
				response.Body.Close()

				if err != nil {
					m.Unlock()
					panic(err)
				}

				t.Reset(time.Duration(time.Second * 1))
				m.Unlock()
			}
		case <-dt.C:
			m.Lock()
			err := os.RemoveAll(m.p)
			if err != nil {
				m.Unlock()
				panic(err)
			}
			err = os.MkdirAll(m.p, 0755)
			if err != nil {
				m.Unlock()
				panic(err)
			}
			dt.Reset(time.Duration(time.Second * 3))
			m.Unlock()

		case sig := <-m.s:
			print(sig.String())
			return
		}
	}
}

func main() {
	fs := flag.NewFlagSet("metrics grabber", flag.ExitOnError)

	scrapeInterval := fs.Int("interval", 1, "Scrape interval in seconds")
	metricsPath := fs.String("path", "./metrics/", "Path to store metrics information")

	helpers.UsageFor(fs, os.Args[0]+" [flags]")
	fs.Parse(os.Args[1:])

	s := make(chan os.Signal)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)

	a := NewMetricsExporter(s, time.Duration(*scrapeInterval)*time.Second, time.Duration(3)*time.Second, *metricsPath)
	a.StartCollectingMetrics()
}

func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}
