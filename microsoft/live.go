package microsoft

import (
    "golang.org/x/oauth2"
    "errors"
)

var Endpoint = oauth2.Endpoint {
    AuthURL:  "https://login.live.com/oauth20_authorize.srf",
    TokenURL: "https://login.live.com/oauth20_token.srf",
    RevokeURL:"https://login.live.com/oauth20_logout.srf",
}

// Destroys both access_token and refresh_token of Live
func Revoke(config *oauth2.Config, token *oauth2.Token) error {
    /*var buf bytes.Buffer
    buf.WriteString(config.Endpoint.RevokeURL)
    buf.WriteByte('?')
    v := url.Values {
        "client_id":		{config.ClientID},
        "redirect_uri": 	{config.RedirectURL},
    }
    buf.WriteString(v.Encode())
    client := config.Client(oauth2.NoContext, token)
    request, err := http.NewRequest("GET", buf.String(), nil)
    if err != nil {
        return err
    }
    data, err := httputil.DumpRequestOut(request, true)
    if err != nil {
        return err
    }
    fmt.Printf("%s", data)
    response, err := client.Do(request)
    if err != nil {
        return err
    }
    data, err = httputil.DumpResponse(response, true)
    if err != nil {
        return err
    }
    fmt.Printf("%s", data)
    return nil*/
    return errors.New("Not implemented, due to https://github.com/OneDrive/onedrive-api-docs/issues/58")
}