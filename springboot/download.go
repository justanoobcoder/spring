package springboot

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// create a function that send a post request to https://start.spring.io/starter.zip
// request body should be a json with the following structure:
// {
//     "dependencies": "web,lombok",
//     "javaVersion": "17",
//     "type": "maven-project",
//     "applicationName": "MyApp"
// }
//

func CreateProject(body Request) {
	mashal, err := json.Marshal(body)
	if err != nil {
		log.Fatal("error marshaling json", err)
	}

	req, err := http.NewRequest("POST", "https://start.spring.io/starter.zip", bytes.NewBuffer(mashal))
	if err != nil {
		log.Fatal("error creating request", err)
	}

	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("error sending request", err)
	}
	defer resp.Body.Close()
	out, err := os.Create(body.Name + ".zip")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatal("error copying response body to file", err)
	}
}
