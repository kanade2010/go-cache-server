package main

import (
  "bufio"
  "fmt"
  "net"
  "strconv"
  "strings"
)

func readLen(r *bufio.Reader) (int, error) {

  tmp, e := r.ReadString(' ')
  if e != nil {
    return 0, e
  }
  fmt.Println("ReadString : ", tmp)

  l, e := strconv.Atoi(strings.TrimSpace(tmp))
  if e!= nil {
    return 0, e
  }
  fmt.Println("Atoi : ", l)

  return l, nil

}


func sendResponse(value []byte, err error, conn net.Conn) error {

  if err != nil {
    errString := err.Error()
    tmp := fmt.Sprintf("-%d ", len(errString)) + errString
    _, e := conn.Write([]byte(tmp))
    return e
  }

  vlen := fmt.Sprintf("%d ", len(value))
  _, e := conn.Write(append([]byte(vlen), value...))
  return e

}


func main() {

  fmt.Println("main start")

  s := strings.NewReader("-100 58")
  br := bufio.NewReader(s)


  t, _ := readLen(br);

  fmt.Println("test 1 : ", t)

  b, _ := br.Peek(5)
  fmt.Printf("%s\n", b)
  // ABCDE

  b[0] = 'a'
  b, _ = br.Peek(5)
  fmt.Printf("%s\n", b)
  // aBCDE

}