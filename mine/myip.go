package main

import (
	"io"
	"net/http"
)

func GetIp() (string, error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", "http://metadata/computeMetadata/v1/instance/network-interfaces/0/access-configs/0/external-ip", nil)
	if err != nil {
		return "", err
	}
	request.Header.Set("Metadata-Flavor", "Google")
	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(ip), err
}
