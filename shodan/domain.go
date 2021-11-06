package shodan

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Domain struct {
	Domain     string       `json:"domain"`
	Tags       []string     `json:"tags"`
	Subdomains []string     `json:"subdomains"`
	Data       []DomainData `json:"data"`
	More       bool         `json:"more"`
}

type DomainData struct {
	Subdomain string `json:"subdomain"`
	Type      string `json:"type"`
	Value     string `json:"value"`
	LastSeen  string `json:"last_seen"`
}

func (s *ShodanClient) DomainInfo(domain string) (*Domain, error) {
	res, err := http.Get(fmt.Sprintf("%s/dns/domain/%s?key=%s", BASE_URL, domain, s.apiKey))

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var domains Domain

	if err := json.NewDecoder(res.Body).Decode(&domains); err != nil {
		return nil, err
	}

	return &domains, nil
}
