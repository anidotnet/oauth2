package oauth2

import (
    "net/http"
    "net/http/httputil"
    "fmt"
    "strings"
)

func debugRequest(request *http.Request) {
    data, _ := httputil.DumpRequestOut(request, false)
    fmt.Printf(strings.Repeat("-", 10) + "\nRequest:\n%s" + strings.Repeat("-", 10) + "\n\n", data)
}

func debugResponse(response *http.Response) {
   var data []byte
    if response != nil && response.StatusCode >= 200 && response.StatusCode < 300 {
        data, _ = httputil.DumpResponse(response, false)
    } else {
        data, _ = httputil.DumpResponse(response, true)
    }
    fmt.Printf(strings.Repeat("-", 10) + "\nResponse:\n%s\n" + strings.Repeat("-", 10) + "\n\n", data)
}
