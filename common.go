package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func JsonDecode(data string, result interface{}) (err error) {
	r := json.NewDecoder(strings.NewReader(data))
	err = r.Decode(result)

	if err != nil {
		return err
	}
	return
}

func SendHttpRequest(url string, jsonDecode bool, result interface{}) (err error) {
	res, err := http.Get(url)

	fmt.Println("Attempting connection to: " + url)

	if err != nil {
		fmt.Println(err)
		return err
	}

	if res.StatusCode != 200 {
		fmt.Printf("HTTP status code: %d", res.StatusCode)
		return errors.New("Status code was not 200.")
	}

	contents, _ := ioutil.ReadAll(res.Body)

	fmt.Printf("Recieved raw: %s\n", string(contents))

	if jsonDecode {
		err = JsonDecode(string(contents), result)

		if err != nil {
			return errors.New("Unable to JSON decode body.")
		}
	} else {
		result = contents
	}

	return
}
