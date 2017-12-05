package main

import (
	"fmt"
	"sync"
)

func main() {

	var mutex = &sync.Mutex{}

	d := make(chan float64)
	var wg sync.WaitGroup
	var result float64
	numberOfPage := 100

	wg.Add(numberOfPage)
	for activity := 1295185809; activity < 1295185809+numberOfPage; activity++ {
		go func(index int) {
			url := fmt.Sprintf("https://www.strava.com/activities/%d", index)
			err := Crawl(url, d)
			if err != nil {
				fmt.Println(err)
				wg.Done()
			}
			mutex.Lock()
			result = result + <-d
			fmt.Println("here: ", result)
			mutex.Unlock()
			wg.Done()
		}(activity)
	}
	wg.Wait()
	fmt.Println("end: ", result)
	return
}
