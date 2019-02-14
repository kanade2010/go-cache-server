package cache-server

import (
  "io/ioutil"
  "net/http"
  "log"
  "strings"
)



func (h *cacheHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  log.Println("url ", r.URL, " Method ", r.Method)

  //Split Get Key
  key := strings.Split(r.URL.EscapedPath(), "/")[2]

  if len(key) == 0 {
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  m = r.Method
  
}