package whois

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Result response alternatives
const (
	PromptAPIResultError      = "error"
	PromptAPIResultRegistered = "registered"
	PromptAPIResultAvailable  = "available"
)

// Message response alternatives
const (
	PromptAPIMessageInvalidDomain = "Not a valid domain name"
)

// PromptAPIWhois PromptAPIWhois representation model
type PromptAPIWhois struct {
	APIKey string
}

type result struct {
	Result string `json:"result"`
}

// ExistsDomain PromptAPI ExistsDomain implementation.
func (p PromptAPIWhois) ExistsDomain(domain string) (exists bool, err error) {
	domain = cutHostname(domain)

	if p.APIKey == "" {
		err = errors.New("promptapiwhois: Api key not found")
		return
	}

	url := fmt.Sprintf(
		"https://api.promptapi.com/whois/check?domain=%s",
		domain,
	)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		err = fmt.Errorf("promptapiwhois: Failed to get request:\n%s", err)
		return
	}

	req.Header.Set("apikey", p.APIKey)

	res, err := client.Do(req)
	if res.Body != nil {
		defer res.Body.Close()
	}
	if err != nil {
		err = fmt.Errorf("promptapiwhois: Failed to do request:\n%s", err)
		return
	}
	var data result

	if err = json.NewDecoder(res.Body).Decode(&data); err != nil {
		err = fmt.Errorf("promptapiwhois: Failed when decoding response:\n%s", err)
	}

	switch data.Result {
	case PromptAPIResultAvailable:
	case PromptAPIResultError:
		exists = false
	case PromptAPIResultRegistered:
		exists = true
	}
	return
}
