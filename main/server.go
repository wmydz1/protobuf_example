package main



import (

    "fmt"

    "net"

    "os"
    "github.com/golang/protobuf/proto"
    "io"
    "bytes"
    "github.com/logoocc/protobuf_example/example"

)

func main() {

    service := ":10000"
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    checkError(err)
    listener, err := net.ListenTCP("tcp", tcpAddr)
    checkError(err)
    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        go handleClient(conn)
    }
}

func handleClient(conn net.Conn) {
    defer conn.Close()
    var buf bytes.Buffer
    _, err := io.Copy(&buf, conn)
    if err != nil {
        // Error handler

    }
    newTest := &example.RegMessage{}
    err = proto.Unmarshal(buf.Bytes(), newTest)
    fmt.Println(*newTest.Username)

}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}