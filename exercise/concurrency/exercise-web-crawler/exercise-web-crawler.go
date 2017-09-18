/* ****************************************************************************
 * Exercise: Web Crawler
 * Link: https://tour.golang.org/concurrency/10
 * ----------------------------------------------------------------------------
 * In this exercise you'll use Go's concurrency features to parallelize a web
 * crawler.
 *
 * Modify the Crawl function to fetch URLs in parallel without fetching the
 * same URL twice.
 *
 * Hint: you can keep a cache of the URLs that have been fetched on a map, but
 * maps alone are not safe for concurrent use!
 * ***************************************************************************/

package main

import (
	"fmt"
	"sync"
	"time"
)

// Fetcher interface
type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type cacheResult struct {
	body string
	urls []string
	err  error
}

type cache struct {
	results map[string]*cacheResult
	mux     sync.Mutex
}

func (c *cache) set(url string, body string, urls []string, err error) {
	cacher.mux.Lock()
	cacher.results[url] = &cacheResult{body, urls, err}
	cacher.mux.Unlock()
}

func (c *cache) get(url string) (*cacheResult, bool) {
	cacher.mux.Lock()
	defer cacher.mux.Unlock()

	cacheValue, ok := cacher.results[url]
	return cacheValue, ok
}

var cacher = cache{results: make(map[string]*cacheResult)}

// Crawl func
func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}

	var body string
	var urls []string
	var err error

	cacheValue, ok := cacher.get(url)

	if ok {
		body, urls, err = cacheValue.body, cacheValue.urls, cacheValue.err
	} else {
		body, urls, err = fetcher.Fetch(url)
		cacher.set(url, body, urls, err)
	}

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher)
	}
	return
}

type fakeResult struct {
	body string
	urls []string
}

type fakeFetcher map[string]*fakeResult

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Package",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}

func main() {
	go Crawl("http://golang.org/", 4, fetcher)

	time.Sleep(time.Second)
}
