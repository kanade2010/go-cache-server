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
  _, e = io.ReadFull(r, k)
  if e != nil {
    return "", e
  }

  return string(k), nil

}

func (s *Server) readKeyAndValue(r *bufio.Reader) (string, []byte, error) {
  kl, e := readLen(r)
  if e != nil {
    return "", nil, e
  }

  vl, e := readLen(r)
  if e != nil {
    return "", nil, e
  }

  k := make([]byte, kl)
  _, e = io.ReadFull(r, k)
  if e != nil {
    return "", nil, e
  }

  v := make([]byte, vl)
  _, e = io.ReadFull(r, v)
  if e != nil {
    return "", nil, e
  }
  
  return string(k), v, nil
}