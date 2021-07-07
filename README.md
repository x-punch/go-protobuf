# Go Protobuf
A go protobuf library, used to help parse data from protobuf or parse to protobuf.

# Usage
```
package main

import (
	"fmt"

	"github.com/x-punch/go-protobuf"
	pb "google.golang.org/protobuf/types/known/structpb"
)

func main() {
	req, err := protobuf.MarshalValue(666)
	if err != nil {
		panic(err)
	}
	fmt.Println(req)

	resp := &pb.Value{Kind: &pb.Value_NumberValue{NumberValue: 666}}
	v := protobuf.UnmarshalValue(resp)
	fmt.Println(v)
}
```