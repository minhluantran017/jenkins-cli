package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/minhluantran017/jenkins-cli/internal/client"
	"github.com/minhluantran017/jenkins-cli/internal/helper"
	log "github.com/sirupsen/logrus"
)

func SendHttpRequest(client *http.Client, req *http.Request) ([]byte, error) {
	log.Debug(" >> Request URL: ", req.URL)
	log.Debug(" >> Request method: ", req.Method)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	log.Debug(" << Response status code: ", res.StatusCode)
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetServerCrumb(profile *client.Profile) string {
	log.Debug("Getting Jenkins crumb")
	client := &http.Client{}
	url := fmt.Sprintf("%s/crumbIssuer/api/json", profile.Url)
	req, err := http.NewRequest("GET", url, nil)
	helper.HandleError(err)
	req.SetBasicAuth(profile.UserName, profile.Password)
	data, err := SendHttpRequest(client, req)
	helper.HandleError(err)

	var result map[string]string
	helper.HandleError(json.Unmarshal(data, &result))
	return result["crumb"]
}

func SendGetRequest(profile *client.Profile, crumb string, resourceUri string) string {
	client := &http.Client{}
	url := fmt.Sprintf("%s/%s", profile.Url, resourceUri)

	log.Debug("Sending GET request to ", url)
	req, err := http.NewRequest("GET", url, nil)
	helper.HandleError(err)
	req.SetBasicAuth(profile.UserName, profile.Password)
	req.Header.Set("Jenkins-Crumb", crumb)

	data, err := SendHttpRequest(client, req)
	helper.HandleError(err)
	return string(data)
}

func SendPostRequest(profile *client.Profile, crumb string, resourceUri string, requestBody *bytes.Buffer, contentType string) string {
	client := &http.Client{}
	url := fmt.Sprintf("%s/%s", profile.Url, resourceUri)

	log.Debug("Sending POST request to ", url)
	req, err := http.NewRequest("POST", url, requestBody)
	helper.HandleError(err)
	req.SetBasicAuth(profile.UserName, profile.Password)
	req.Header.Set("Jenkins-Crumb", crumb)
	req.Header.Set("Content-Type", contentType)

	data, err := SendHttpRequest(client, req)
	helper.HandleError(err)
	return string(data)
}