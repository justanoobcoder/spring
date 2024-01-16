package springboot

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
)

// request body for the download request
type Request struct {
	Dependencies string `json:"dependencies"`
	JavaVersion  string `json:"javaVersion"`
	Type         string `json:"type"`
	Version      string `json:"version"`
	Packaging    string `json:"packaging"`
	Language     string `json:"language"`
	BootVersion  string `json:"platformVersion"`
	GroupId      string `json:"groupId"`
	ArtifactId   string `json:"artifactId"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	PackageName  string `json:"packageName"`
}

// encode the request as a form url encoded string,
// the result will look something like this:
// type=maven-project&language=java&groupId=com.example
// the key names are the same as the `json` tag name of the fields
// the values are the string values of the fields
func urlEncode(reqBody Request) string {
	data := url.Values{}
	val := reflect.ValueOf(reqBody)
	for i := 0; i < val.Type().NumField(); i++ {
		k := val.Type().Field(i).Tag.Get("json")
		v := val.Field(i).String()
		data.Add(k, v)
	}

	return data.Encode()
}

func Download(reqBody Request, filename string) (int, error) {
	// the request must be a form url encoded POST request to work
	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", initializrUrl, filename),
		strings.NewReader(urlEncode(reqBody)),
	)
	if err != nil {
		return 0, fmt.Errorf("error creating download request\n%v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// set request timeout otherwise it will hang forever
	client := http.Client{Timeout: timeout}
	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("error sending download request\n%v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return resp.StatusCode,
			fmt.Errorf("error downloading file\nstatus code: %s\nmessage: %v",
				resp.Status, resp.Body,
			)
	}

	// golang doesn't have a built-in way to download a file from a request,
	// so we have to create a file and copy the response body to it
	file, err := os.Create(filename)
	if err != nil {
		return 0, fmt.Errorf("error creating download file\n%v", err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return 0, fmt.Errorf("error copying download file\n%v", err)
	}

	return resp.StatusCode, nil
}
