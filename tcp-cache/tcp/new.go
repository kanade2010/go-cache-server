package tcp

import (
  "../cache"
  "net"
  "log"
)

type Server struct {
  cache.Cache
}

func (s *Server) Listen() {

  ln, e := net.Listen("tcp", ":12346")
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
    log.Println("new connection go handler:")
    go s.handleConnection(conn)
  }

}

func New(c cache.Cache) *Server {
  return &Server{c}
}