package faker

import "net/http"

const (
	headerUpgrade        string = "Upgrade"
	headerValueWebsocket string = "websocket"
)

// isWebsocketUpgradeRequest checks if the given http.Request is a websocket upgrade request
// by checking if the Upgrade header is set to "websocket"
func isWebsocketUpgradeRequest(r *http.Request) bool {
	for _, header := range r.Header[headerUpgrade] {
		if header == headerValueWebsocket {
			return true
		}
	}
	return false
}
