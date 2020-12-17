package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/wpgitupdater/wpgitupdater/internal/constants"
	"github.com/wpgitupdater/wpgitupdater/internal/git"
	"github.com/wpgitupdater/wpgitupdater/internal/interfaces"
	"github.com/wpgitupdater/wpgitupdater/internal/utils"
	"io/ioutil"
	"net/http"
)

func UpdateUsage(usageType string, slug string, stats bool) error {
	var provider string
	var repository string
	if stats {
		provider = git.GetProvider()
		repository = git.GetRepository()
	} else {
		slug = ""
		provider = "*"
		repository = "*/*"
	}
	body := interfaces.UpdateUsage{Type: usageType, Provider: provider, Repository: repository, Slug: slug}
	body.Meta = interfaces.UpdateUsageMeta{Build: constants.Build, Version: constants.Version, ConfigVersion: constants.ConfigVersion}
	data, err := json.Marshal(body)
	if err != nil {
		return err
	}
	url := constants.ApiUrl + "/" + utils.GetUpdaterToken() + "/usage"
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Add("User-Agent", constants.UserAgent)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != 201 {
		return errors.New(string(responseBody))
	}

	return nil
}
