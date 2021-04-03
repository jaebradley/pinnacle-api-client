package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Client is object that contains resources to call Pinnacle's API
type Client struct {
	authorizationToken string
	Client             *http.Client
}

type HaywireDetails struct {
	ApiKey string `json:"apiKey"`
}

type ApiDetails struct {
	HaywireDetails HaywireDetails `json:"haywire"`
}

type ApplicationDetails struct {
	ApiDetails ApiDetails `json:"api"`
}

type ServiceStatus struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type UpstreamServiceStatus struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Health string `json:"health"`
}

type StatusDetails struct {
	Code             string                  `json:"code"`
	Description      string                  `json:"description"`
	Services         []ServiceStatus         `json:"services"`
	UpstreamServices []UpstreamServiceStatus `json:"upstream"`
}

func New(authorizationToken *string, httpClient *http.Client) (client Client) {
	var token string
	if nil != authorizationToken {
		token = *authorizationToken
	}

	var httpClientInstance *http.Client = http.DefaultClient
	if nil != httpClient {
		httpClientInstance = httpClient
	}

	return Client{
		authorizationToken: token,
		Client:             httpClientInstance,
	}
}

func fetchApplicationDetails(client *Client) (details ApplicationDetails, err error) {
	details = ApplicationDetails{}

	u := url.URL{
		Scheme: "https",
		Host:   "www.pinnacle.com",
		Path:   "config/app.json",
	}

	response, err := http.Get(u.String())
	if err != nil {
		return details, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return details, err
	}

	err = json.Unmarshal(body, &details)
	if err != nil {
		return details, err
	}

	return details, nil
}

func FetchStatus(client *Client) (details StatusDetails, err error) {
	details = StatusDetails{}

	u := url.URL{
		Scheme: "https",
		Host:   "guest.api.arcadia.pinnacle.com",
		Path:   "0.1/status",
	}

	request, err := http.NewRequest(
		"GET",
		u.String(),
		nil,
	)

	if err != nil {
		return details, err
	}

	request.Header.Add("X-API-KEY", client.authorizationToken)

	response, err := client.Client.Do(request)
	if err != nil {
		return details, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return details, err
	}

	err = json.Unmarshal(body, &details)
	if err != nil {
		return details, err
	}

	return details, nil
}
