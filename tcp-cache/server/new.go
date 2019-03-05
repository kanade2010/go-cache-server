package tcp

import (
  "../cache"
  "net"
)

type Server struct {
  cache.Cache
}

func (s *Server) Listen() {

  ln, e := net.Listen("tcp", ":12345")
  if e != nil {
    panic(e)
  }

  for {
    conn, err := ln.Accept()
    if err != nil {
      // handle error
      panic(e)
      continue
    }
    go handleConnection(conn)
  }

}

func New(c cache.Cache) *Server {
  return &Server{c}
}