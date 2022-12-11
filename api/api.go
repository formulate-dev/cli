package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/formulate-dev/cli/model"
)

const (
	BASE_URL = "https://formulate-form.0xfff.workers.dev"
)

func CreateForm(title string) (form model.Form, err error) {
	requestBody, err := json.Marshal(map[string]string{"title": title})
	if err != nil {
		return
	}

	r, err := http.Post(BASE_URL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return
	}
	defer r.Body.Close()

	if r.StatusCode != 201 {
		return form, fmt.Errorf("failed to create form – invalid status code %d", r.StatusCode)
	}

	err = json.NewDecoder(r.Body).Decode(&form)
	return
}

func UpdateForm(form model.Form) error {
	requestBody, err := json.Marshal(form)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s?id=%s&secret=%s", BASE_URL, form.Id, form.Secret)

	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	r, err := client.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if r.StatusCode != 200 {
		return fmt.Errorf("failed to update form – invalid status code %d", r.StatusCode)
	}

	return nil
}
