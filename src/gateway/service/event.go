package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/sef-comp/Hangover/gateway/models"
)


func CheckEventServiceHealth(serviceAddress string) (any, error){
	checkHealthURL := fmt.Sprintf("%s/manage/health", serviceAddress)
	log.Println(checkHealthURL)
	req, err := http.NewRequest(http.MethodGet, checkHealthURL, nil)

	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Timeout: 10 * time.Minute,
	}

	resp, err := client.Do(req)
	return resp, err
}


func GetEvents(serviceAddress string) (*[]models.Event, error){
	if _, err := CheckEventServiceHealth(serviceAddress); err != nil {
		return &[]models.Event{}, err
	}

	requestURL := fmt.Sprintf("%s/api/v1/events", serviceAddress)

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Println("Failed to create an http request")
		return nil, err
	}

	client := &http.Client{
		Timeout: 10 * time.Minute,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed request to event service: %w", err)
	}

	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			fmt.Println("Failed to close response body")
		}
	}(res.Body)

	events := &[]models.Event{}
	if err = json.NewDecoder(res.Body).Decode(events); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return events, nil
}
