package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func DownloadFile(url, destDir string) error {
	filename := filepath.Base(url)
	fmt.Println(">>filename", filename)
	filePath := filepath.Join(destDir, filename)
	fmt.Println(">>filename2", filePath)
	out, err := os.Create(filePath)
	fmt.Println(">>filename3", out)

	if err != nil {
		return err
	}

	defer out.Close()

	fmt.Println("downloading", url)

	start := time.Now()

	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Bad Status: %s", resp.Status)
	}

	_, err = io.Copy(out, resp.Body)

	if err != nil {
		return err
	}

	fmt.Printf("Download %s took %s \n", url, time.Since(start))
	return nil
}

type Result struct {
	URL string
	Error error
	Size  int64
}

func downloadURLs(wg *sync.WaitGroup, limiter chan struct{}, result chan Result, url, destDir string, errors []error)  {

	defer wg.Done()
	//each time the goroutine runs, we write to the limiter channel
	limiter <- struct{}{}
	//we release the slot after the return of the surrounding function
	defer func(){ <-limiter }()

	start := time.Now()
	//the base file name
	filename := filepath.Base(url)
	
	filePath := filepath.Join(destDir, filename)
	out, err := os.Create(filePath)

	resp, err := http.Get(url)

	if err != nil {
		os.Remove(filePath)
		result <- Result{
			URL: "",
			Error: err,
		}
		errors = append(errors, err)
		return
	}

	defer out.Close()

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		result <- Result{
			URL: "",
			Error: fmt.Errorf("Bad Status: %s", resp.Status),
		}
		errors = append(errors, fmt.Errorf("Bad Status: %s", resp.Status))
		return
	}

	size, err := io.Copy(out, resp.Body)

	if err != nil {
		result <- Result{
			URL: "",
			Error: err,
		}
		errors = append(errors, err)
		return
	}

	result <- Result{
		URL: url,
		Error: nil,
		Size: size,
	}
	fmt.Printf("the url(%v) has finished downloading at %s\n", url, time.Since(start))
}

func readResults(result chan Result, wg *sync.WaitGroup, totalSize *int64) {
	defer wg.Done()

	for msg := range(result) {
		*totalSize += msg.Size
		fmt.Println("here is the result of the operation for this url", msg.Size)
		return
	}
}

func concurrentDowloader(urls []string, destDir string, maxCurrent int) error {
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return err
	}

	//channel for our results
	results := make(chan Result)

	//waitgroup
	var wg sync.WaitGroup

	//this is the limiter channel that rate limites how many goroutines excecutes simulataneously at once
	limiter := make(chan struct{}, maxCurrent)

	start := time.Now()

	errors := make([]error, 0)


	var total_size int64 = 0


	for _, url := range(urls) {
		wg.Add(2)
		go downloadURLs(&wg, limiter, results, url, destDir, errors)
		go readResults(results, &wg, &total_size)
	}
	wg.Wait()
	if len(errors) > 0 {
		for _, err := range(errors) {
			fmt.Println("err", err)
		}
	}
	close(results)
	fmt.Printf("these urls finished downloading at an average of %d with a total size of %v\n", time.Since(start), total_size)
	return nil
}

func main() {
	urls := []string{"http://www1.lasalle.edu/~zelenenkic1/red_eurpean_dragon.jpg", "https://contentful.harrypotter.com/usf1vwtuqyxm/5twS0kKj0AcwMeIia4YecE/9f0ffa8b8d2503690f07e21112cd3e40/Dragon_WB_F4_HeadDetailHorntail_Illust_100615_Land.jpg"}
	
	//err := DownloadFile(url, "./projects/file-downloader")

	err := concurrentDowloader(urls, "./projects/file-downloader", 4)

	if err != nil {
		fmt.Println("err", err)
	}

	fmt.Println("Done>>")
}