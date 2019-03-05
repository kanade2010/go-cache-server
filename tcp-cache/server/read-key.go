package tcp

import (
  "bufio"
  "io"
)

func (s *Server) readKey(r *bufio.Reader) (string, error) {

  kl, e := readLen(r)
  if e != nil {
    return "", e
  }

  k := make([]byte, kl)

}