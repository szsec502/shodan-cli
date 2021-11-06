package shodan

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type APIInfo struct {
	QueryCredits int    `json:"query_credits"`
	ScanCredits  int    `json:"scan_credits"`
	Telnet       bool   `json:"telnet"`
	Plan         string `json:"plan"`
	HTTPS        bool   `json:"https"`
	Unlocked     bool   `json:"unlocked"`
}

func (s *ShodanClient) APIInfo() (*APIInfo, error) {
	res, err := http.Get(fmt.Sprintf("%s/api-info?key=%s", BASE_URL, s.apiKey))

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var apiInfo APIInfo

	if err := json.NewDecoder(res.Body).Decode(&apiInfo); err != nil {
		return nil, err
	}

	return &apiInfo, nil
}
