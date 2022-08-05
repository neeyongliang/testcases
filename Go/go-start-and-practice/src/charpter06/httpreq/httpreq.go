package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	client := &http.Client{}

	req, err := http.NewRequest("POST",
		"https://www.163.com", strings.NewReader("key=value"))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	req.Header.Add("User-Agent", "chrome")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

	defer resp.Body.Close()
}
