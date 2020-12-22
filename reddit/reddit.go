package reddit

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const AccessTokenUrl = "https://www.reddit.com/api/v1/access_token"
const UserAgent = "playlist reddit-bot"

type Credentials struct {
	Username     string
	Password     string
	ClientId     string
	ClientSecret string
}

type Client struct {
	Initialized bool
	accessToken string
	credentials *Credentials
	httpClient  *http.Client
}

func InitClient(credentials *Credentials) (*Client, error) {
	client := &Client{credentials: credentials, httpClient: &http.Client{}}
	form := url.Values{
		"grant_type": {"password"},
		"username":   {credentials.Username},
		"password":   {credentials.Password},
	}
	log.Printf("Initializing Reddit Client 2")

	req, reqError := http.NewRequest("POST", AccessTokenUrl, strings.NewReader(form.Encode()))
	if reqError != nil {
		return nil, fmt.Errorf("error during Access token request creation: %w", reqError)
	}
	req.Header.Add("User-agent", UserAgent)
	req.SetBasicAuth(credentials.ClientId, credentials.ClientSecret)
	resp, respError := client.httpClient.Do(req)
	if respError != nil {
		return nil, fmt.Errorf("error during fetch access token:  %w", respError)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}
	var responseData map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&responseData)
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error response code from auth api: %d", resp.StatusCode)
	}
	for k, v := range responseData {
		fmt.Printf("Key: %s, Value: %s\n", k, v)
	}
	client.Initialized = true
	client.accessToken = fmt.Sprintf("%v", responseData["access_token"])
	return client, nil
}
