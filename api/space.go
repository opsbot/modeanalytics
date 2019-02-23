package mode

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
)

// SpaceList - GET /api/{org}/spaces - get list of spaces
func SpaceList() SpaceCollection {
	log.Trace("Mode API SpaceList")
	url := getEndpoint(fmt.Sprintf("/api/%v/spaces", Config.Org))
	resp := Request("GET", url, []byte{})

	if resp.StatusCode != 200 {
		showFullResponseBody(resp)
		log.Fatalf("RequestError: %v", resp.Status)
	}

	var collection map[string]SpaceCollection
	json.Unmarshal([]byte(resp.Body), &collection)
	// utils.PrettyPrint(collection)

	spaces := collection["_embedded"]
	// utils.PrettyPrint(spaces)

	// for _, v := range spaces.Spaces {
	// 	fmt.Printf("(%v) %v\n", v.Token, v.Name)
	// }
	return spaces
}

// SpaceCreate - POST /api/{org}/spaces/ - get list of spaces
func SpaceCreate(space []byte) APIResponse {
	log.Trace("Mode API SpaceCreate")
	url := getEndpoint(fmt.Sprintf("/api/%v/spaces/", Config.Org))
	resp := Request("POST", url, space)

	if resp.StatusCode != 200 {
		showFullResponseBody(resp)
		log.Fatalf("RequestError: %v", resp.Status)
	}
	return resp
}

// SpaceGet - GET /api/{org}/spaces/{space} - get list of spaces
func SpaceGet(token string) Space {
	log.Trace("Mode API SpaceGet")
	url := getEndpoint(fmt.Sprintf("/api/%v/spaces/%v", Config.Org, token))
	resp := Request("GET", url, []byte{})

	var space Space
	json.Unmarshal([]byte(resp.Body), &space)

	if resp.StatusCode != 200 {
		showFullResponseBody(resp)
		log.Fatalf("RequestError: %v", resp.Status)
	}
	return space
}

// SpaceDelete - DELETE /api/{org}/spaces/{space} - get list of spaces
func SpaceDelete(space string) APIResponse {
	log.Trace("Mode API SpaceDelete")
	url := getEndpoint(fmt.Sprintf("/api/%v/spaces/%v", Config.Org, space))
	resp := Request("DELETE", url, []byte{})

	if resp.StatusCode != 200 {
		showFullResponseBody(resp)
		log.Fatalf("RequestError: %v", resp.Status)
	}
	return resp
}
