package springboot

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
	"time"
)

func urlEncode(req Request) string {
	data := url.Values{}
	val := reflect.ValueOf(req)
	for i := 0; i < val.Type().NumField(); i++ {
		k := val.Type().Field(i).Tag.Get("json")
		v := val.Field(i).String()
		data.Add(k, v)
	}

	return data.Encode()
}

func Download(body Request, filename string) {
	req, err := http.NewRequest(
		"POST",
		"https://start.spring.io/"+filename,
		strings.NewReader(urlEncode(body)),
	)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		log.Fatal("error creating request", err)
	}

	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("error sending request", err)
	}

	defer resp.Body.Close()
	out, err := os.Create(filename)
	if err != nil {
		log.Fatal("error creating file", err)
	}

	defer out.Close()
	_, err = io.Copy(out, resp.Body)

	if err != nil {
		log.Fatal("error copying response body to file", err)
	}
}
