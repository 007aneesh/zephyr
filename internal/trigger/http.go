package trigger

import (
    "net/http"
)

func HTTPTrigger(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("HTTP Trigger Invoked"))
}
