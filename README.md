# Go Protobuf
A go protobuf library, used to help parse data from protobuf or parse to protobuf.

# Usage
```
package main

import (
    "fmt"

	pb "github.com/golang/protobuf/ptypes/struct"
    "github.com/x-punch/go-protobuf"
)

func main() {
    resp := &pb.Value{Kind: &pb.Value_NumberValue{NumberValue: 666}}

    v := protobuf.GetValue(resp)
    fmt.Println(v)
}
```