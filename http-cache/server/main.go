/*package main

import (
	"./cache"
  "fmt"
)

func main() {
	c := cache.New("inmemory")

  //k, v := "sola", []byte{'a','i','l','u','m','i','y','a'}

  k, v := "sola", []byte("ailumiyana")


  c.Set(k, v)

  tmp, _ := c.Get(k)
  fmt.Println("key: ", k, " value: ", tmp)

  c.Del(k)

  tmp, _ = c.Get(k)
  fmt.Println("key: ", k, " value: ", tmp)

}*/


package main

import (
  "./cache"
  "io/ioutil"
  "net/http"
  "log"
  "strings"
)

type cacheHandler struct {
  cache.Cache
}

func (h *cacheHandler) CacheHandler(w http.ResponseWriter, r *http.Request) {
  log.Println("url ", r.URL, " Method ", r.Method)

  //Split Get Key
  key := strings.Split(r.URL.EscapedPath(), "/")[2]

  if len(key) == 0 {
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  m := r.Method

  if m == http.MethodPut {
    h.HandlePut(key, w, r)
    return
  } else if m == http.MethodGet {
    h.HandleGet(key, w, r)
    return
  } else if m == http.MethodDelete {
    h.HandleDelete(key, w, r)
    return
  }

  w.WriteHeader(http.StatusMethodNotAllowed)
}


func (h *cacheHandler) HandlePut(k string, w http.ResponseWriter, r *http.Request){
  b, _ := ioutil.ReadAll(r.Body)

  if len(b) != 0 {
    e := h.Set(k, b)
    if e != nil {
      log.Println(e)
      w.WriteHeader(http.StatusInternalServerError)
    } else {
      w.Write([]byte("successful"))
    }
  }
}

func (h *cacheHandler) HandleGet(k string, w http.ResponseWriter, r *http.Request){
  b, e := h.Get(k)
  
  if e != nil {
    log.Println(e)
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  if len(b) == 0 {
    w.WriteHeader(http.StatusNotFound)
    return
  }

  w.Write(b)

}

func (h *cacheHandler) HandleDelete(k string, w http.ResponseWriter, r *http.Request){
  e := h.Del(k)

  if e != nil {
    log.Println(e)
    w.WriteHeader(http.StatusInternalServerError)
  } else {
    w.Write([]byte("successful"))
  }

}

func main() {
  c := cache.New("inmemory")
  h := cacheHandler{c}
  http.HandleFunc("/cache/", h.CacheHandler)
  http.ListenAndServe(":26316", nil)

}