package springboot

import (
	"errors"
	"fmt"
	"io"
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

func Download(body Request, filename string) (int, error) {
	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", springUrl, filename),
		strings.NewReader(urlEncode(body)),
	)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		return 0, errors.New("error creating request")
	}

	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return 0, errors.New("error sending request")
	}

	defer resp.Body.Close()
	out, err := os.Create(filename)
	if err != nil {
		return 0, errors.New("error creating file")
	}

	defer out.Close()
	_, err = io.Copy(out, resp.Body)

	if err != nil {
		return 0, errors.New("error copying file")
	}

	return resp.StatusCode, nil
}
