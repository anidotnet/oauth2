package box

import (
    "golang.org/x/oauth2"
    "net/url"
    "net/http"
    "strings"
)

// Endpoint is Box's OAuth 2.0 endpoint.
var Endpoint = oauth2.Endpoint{
    AuthURL:  "https://app.box.com/api/oauth2/authorize",
    TokenURL: "https://app.box.com/api/oauth2/token",
    RevokeURL:"https://www.box.com/api/oauth2/revoke",
}

// Destroys both access_token and refresh_token for Box
func Revoke(config *oauth2.Config, token *oauth2.Token) error {
    v := url.Values {
        "client_id":		{config.ClientID},
        "client_secret": 	{config.ClientSecret},
        "token":			{token.RefreshToken},
    }
    client := config.Client(oauth2.NoContext, token)
    request, err := http.NewRequest("POST", config.Endpoint.RevokeURL, strings.NewReader(v.Encode()))
    if err != nil {
        return err
    }

    _, err = client.Do(request)
    if err != nil {
        return err
    }
    return nil
}
