package microsoft

import (
	"golang.org/x/oauth2"
)

var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://login.live.com/oauth20_authorize.srf",
	TokenURL: "https://login.live.com/oauth20_token.srf",
}
