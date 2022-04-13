package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func main() {
	endpoint := "http://fgo.vn/tai-hinh-anh.html"
	fileDir := "download/"
	prefix := "E"
	start := 40000
	end := 100

	// Clean up download folder
	Cleanup(fileDir)

	for i := start; i <= start+end; i++ {
		fileName := fmt.Sprintf("%s%05d", prefix, i)
		fmt.Print("Fetching " + fileName)
		if body, ok := FetchFile(endpoint, fileName); ok {
			// Write byte to file
			err := ioutil.WriteFile(fileDir+fileName+".JPG", body, 0644)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(" - Done")
		} else {
			fmt.Println(" - Failed")
		}
	}

}

func FetchFile(endpoint string, fileName string) ([]byte, bool) {
	data := url.Values{}
	data.Set("taianh", "1")
	data.Set("code", fileName)

	client := &http.Client{}
	r, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode())) // URL-encoded payload
	if err != nil {
		return nil, false
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	res, err := client.Do(r)
	if err != nil {
		return nil, false
	}

	// Check header for content type
	//	Content-Disposition: attachment; filename=D40029.JPG
	if res.Header.Get("Content-Disposition") != "attachment; filename="+fileName+".JPG" {
		return nil, false
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, false
	}
	return body, true
}

func Cleanup(fileDir string) {
	files, err := ioutil.ReadDir(fileDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		err = os.Remove(fileDir + f.Name())
		if err != nil {
			log.Fatal(err)
		}
	}
}
