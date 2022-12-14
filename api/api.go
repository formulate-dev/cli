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

func CreateForm(title, template string) (form model.Form, err error) {
	requestBody, err := json.Marshal(map[string]string{"title": title, "template": template})
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

func PublishForm(form model.Form) error {
	requestBody, err := json.Marshal(form)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/publish?id=%s&secret=%s", BASE_URL, form.Id, form.Secret)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
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
		return fmt.Errorf("failed to publish form – invalid status code %d", r.StatusCode)
	}

	return nil
}

type setCustomIdResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func SetCustomId(form *model.Form, id string) error {
	requestBody, err := json.Marshal(map[string]string{"customId": id})
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/set_custom_id?id=%s&secret=%s", BASE_URL, form.Id, form.Secret)

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
		return fmt.Errorf("failed to set custom form ID – invalid status code %d", r.StatusCode)
	}

	var responseData setCustomIdResponse
	err = json.NewDecoder(r.Body).Decode(&responseData)
	if err != nil {
		return err
	}

	if !responseData.Success {
		return fmt.Errorf("failed to set custom form ID: %s", responseData.Message)
	}

	return nil
}

func VerifyEmail(form *model.Form, email string) error {
	requestBody, err := json.Marshal(map[string]string{"email": email})
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/verify_email?id=%s&secret=%s", BASE_URL, form.Id, form.Secret)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
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
		return fmt.Errorf("failed to send verification email – invalid status code %d", r.StatusCode)
	}

	return nil
}

type setEmailResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func SetEmail(form *model.Form, code string) error {
	requestBody, err := json.Marshal(map[string]string{"verifyCode": code})
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/set_email?id=%s&secret=%s", BASE_URL, form.Id, form.Secret)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
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
		return fmt.Errorf("failed to verify email address – invalid status code %d", r.StatusCode)
	}

	var responseData setEmailResponse
	err = json.NewDecoder(r.Body).Decode(&responseData)
	if err != nil {
		return err
	}

	if !responseData.Success {
		return fmt.Errorf("failed to verify email address: %s", responseData.Message)
	}

	return nil
}
