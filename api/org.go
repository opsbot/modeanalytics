package mode

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
)

// OrgGet - GET /api/account - get authorizing user
func OrgGet() Organization {
	log.Trace("Modeanalytics get organization info")
	url := getEndpoint(fmt.Sprintf("/api/%v/", Config.Org))
	resp := Request("GET", url, []byte{})

	var org Organization
	json.Unmarshal([]byte(resp.Body), &org)

	if resp.StatusCode != 200 {
		showFullResponseBody(resp)
		log.Fatalf("RequestError: %v", resp.Status)
	}
	return org
}
