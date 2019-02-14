package main

import (
	"./cache"
  "fmt"
)

func main() {
	c := cache.New("inmemory")

  k, v := "sola", []byte{'a','i','l','u','m','i','y','a'}

  c.Set(k, v)

  tmp, _ := c.Get(k)
  fmt.Println("key: ", k, " value: ", tmp)

  c.Del(k)

  tmp, _ = c.Get(k)
  fmt.Println("key: ", k, " value: ", tmp)

}
