package server

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"

	"github.com/minhluantran017/jenkins-cli/internal/client"
	"github.com/minhluantran017/jenkins-cli/internal/helper"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Validate Jenkinsfile at specified path
func ValidateJenkinsfile(path string) {
	profile := client.GetProfile(viper.GetString("profile"))
	crumb := GetServerCrumb(&profile)

	log.Debug("Opening Jenkinsfile at ", path)
	file, err := os.Open(path)
	helper.HandleError(err)
	defer file.Close()

	log.Debug("Creating multi-part form field")
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormField("jenkinsfile")
	helper.HandleError(err)
	_, err = io.Copy(part, file)
	helper.HandleError(err)
	err = writer.Close()
	helper.HandleError(err)
	
	log.Debug("Validating Jenkinsfile")
	log.Info(SendPostRequest(&profile, crumb, "pipeline-model-converter/validate", body, writer.FormDataContentType()))
}