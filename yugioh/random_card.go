package yugioh

import "net/http"

// RandomCardService ...
type RandomCardService struct {
	client *Client
}

// One ...
func (s *RandomCardService) One() (*Card, *http.Response, error) {
	u := defaultBaseURL + "randomcard.php"
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	cards := new([]Card)
	resp, err := s.client.Do(req, cards)
	if err != nil {
		return nil, resp, err
	}
	return &(*cards)[0], resp, err
}