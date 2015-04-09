package box

import (
	"golang.org/x/oauth2"
)

// Endpoint is Box's OAuth 2.0 endpoint.
var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://app.box.com/api/oauth2/authorize",
	TokenURL: "https://app.box.com/api/oauth2/token",
}
