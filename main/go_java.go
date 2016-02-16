package  main

import (
"net"
"fmt"
"os"
proto "github.com/golang/protobuf/proto"
"github.com/logoocc/protobuf_example/example"
)


func main() {
regMessage := &example.Person{
Id:       proto.Int32(10001),
Name: proto.String("samchen"),
Email: proto.String("samchen@google.com"),
Friends:[]string{"person1","person2"},
}

fmt.Println(regMessage)

buffer, err := proto.Marshal(regMessage)
if err != nil {
fmt.Printf("failed: %s\n", err)
return
}


pTCPAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:12000")
if err != nil {
fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
return
}
pTCPConn, err := net.DialTCP("tcp", nil, pTCPAddr)
if err != nil {
fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
return
}
pTCPConn.Write(buffer)

}
