package tcp

import (
  "bufio"
  "io"
  "log"
  "net"
)


func (s *Server) get(conn net.Conn, r *bufio.Reader) error {
  
  k, e := s.readKey(r)
  if e != nil {
    return e
  }

  v, e := s.Get(k)

  return sendResponse(v, e, conn)

}