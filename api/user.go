package mode

import log "github.com/sirupsen/logrus"

// AuthGet - GET /api/account - get authorizing user
func AuthGet() {
	log.Trace("Modeanalytics get authorizing user")
	url := getEndpoint("/api/account")
	resp := Request("GET", url, []byte{})

	if resp.StatusCode != 200 {
		showFullResponseBody(resp)
		log.Fatalf("RequestError: %v", resp.Status)
	}
}
