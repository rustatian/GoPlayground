package main

import (
	"net/url"
)

func main() {
	var urls []url.URL
	urls = append(urls, url.URL{
		Host: "foo",
	})
	urls = append(urls, url.URL{
		Host: "bar",
	})
	urls = append(urls, url.URL{
		Host: "baz",
	})

	var urls2 []*url.URL

	for _, u := range urls {
		ur := u
		urls2 = append(urls2, &ur)
	}

	println("foo")
}
