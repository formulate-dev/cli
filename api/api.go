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
