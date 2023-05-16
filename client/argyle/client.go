package argyle

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ArgyleClient struct {
	Client *http.Client
}

type CreateUserTokenResp struct {
	UserToken string `json:"user_token"`
	Id        string `json:"id"`
}

func NewArgyleClient() *ArgyleClient {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	return &ArgyleClient{
		Client: client,
	}
}

func (a *ArgyleClient) GenerateNewUserToken(ctx context.Context) (string, error) {
	req, err := http.NewRequest("POST", "https://api-sandbox.argyle.com/v2/users", nil)
	if err != nil {
		fmt.Println("Error creating POST request:", err)
		return "", err
	}
	authString := "1b2d278e-6159-44e4-bef0-92f24c30cd43" + ":" + "aVGhUP7anZUeXDkR"
	authEncoded := base64.StdEncoding.EncodeToString([]byte(authString))
	req.Header.Add("Authorization", "Basic "+authEncoded)

	resp, err := a.Client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return "", err
	}

	defer resp.Body.Close()

	var tokenResp CreateUserTokenResp
	err = json.NewDecoder(resp.Body).Decode(&tokenResp)
	if err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return "", nil
	}

	return tokenResp.UserToken, nil

}
