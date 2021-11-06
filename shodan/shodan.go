package shodan

const BASE_URL = "https://api.shodan.io"

type ShodanClient struct {
	apiKey string
}

func New(apikey string) *ShodanClient {
	return &ShodanClient{apiKey: apikey}
}
