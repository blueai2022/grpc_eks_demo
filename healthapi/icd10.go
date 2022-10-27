package healthapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/blueai2022/appsubmission/config"
)

const (
	backChanProtocol = "http"
	jsonContentType  = "application/json"
)

var (
	ErrMedicalEntityNotFound = errors.New("medical entity not found")
)

type icd10Request struct {
	MedicalText string `json:"medical_text"`
}

func ICD10(config *config.Config, medicalText string, icdOut interface{}) error {
	req := &icd10Request{
		MedicalText: medicalText,
	}
	reqBody, err := json.Marshal(req)
	if err != nil {
		verr := fmt.Errorf("failed to marshal request: %s", err)
		return verr
	}

	url := fmt.Sprintf("%s://%s%s", backChanProtocol, config.HealthApiServerAddress, config.HealthApiUrlPath)
	rsp, err := post(url, reqBody, jsonContentType)
	if err != nil {
		return err
	}

	if rsp.StatusCode != http.StatusOK && rsp.StatusCode != http.StatusNoContent {
		err = fmt.Errorf("status code %s", rsp.Status)
		return err
	}

	if rsp.StatusCode == http.StatusNoContent {
		return ErrMedicalEntityNotFound
	}

	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, icdOut)
	if err != nil {
		log.Printf("unexpected response body from health api: %s", body)
		return err
	}

	return nil
}
