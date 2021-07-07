package protobuf_test

import (
	"testing"

	"github.com/x-punch/go-protobuf"
	pb "google.golang.org/protobuf/types/known/structpb"
)

func TestUnmarshalNumberValue(t *testing.T) {
	num := float64(123)
	p := &pb.Value{Kind: &pb.Value_NumberValue{NumberValue: num}}
	v := protobuf.UnmarshalValue(p)
	if v != num {
		t.Fail()
	}
}

func TestUnmarshalStringValue(t *testing.T) {
	str := "string"
	p := &pb.Value{Kind: &pb.Value_StringValue{StringValue: str}}
	v := protobuf.UnmarshalValue(p)
	if v != str {
		t.Fail()
	}
}

func TestUnmarshalNullValue(t *testing.T) {
	p := &pb.Value{Kind: &pb.Value_NullValue{}}
	v := protobuf.UnmarshalValue(p)
	if v != nil {
		t.Fail()
	}
}

func TestUnmarshalBoolValue(t *testing.T) {
	p := &pb.Value{Kind: &pb.Value_BoolValue{BoolValue: true}}
	v := protobuf.UnmarshalValue(p)
	if v != true {
		t.Fail()
	}
}

func TestUnmarshalStructValue(t *testing.T) {
	v1 := float64(123)
	v2 := "strings"
	p := &pb.Value{Kind: &pb.Value_StructValue{StructValue: &pb.Struct{
		Fields: map[string]*pb.Value{
			"v1": {Kind: &pb.Value_NumberValue{NumberValue: v1}},
			"v2": {Kind: &pb.Value_StringValue{StringValue: v2}},
		},
	}}}
	v := protobuf.UnmarshalValue(p)
	if maps, ok := v.(map[string]interface{}); ok {
		if len(maps) == 2 && maps["v1"] == v1 && maps["v2"] == v2 {
			return
		}
	}
	t.Fail()
}

func TestUnmarshalListValue(t *testing.T) {
	v1 := float64(123)
	v2 := "strings"
	p := &pb.Value{Kind: &pb.Value_ListValue{ListValue: &pb.ListValue{
		Values: []*pb.Value{
			{Kind: &pb.Value_NumberValue{NumberValue: v1}},
			{Kind: &pb.Value_StringValue{StringValue: v2}},
		},
	}}}
	v := protobuf.UnmarshalValue(p)
	if array, ok := v.([]interface{}); ok {
		if len(array) == 2 && array[0] == v1 && array[1] == v2 {
			return
		}
	}
	t.Fail()
}

func TestMarshalNull(t *testing.T) {
	v, err := protobuf.MarshalValue(nil)
	if err != nil {
		t.Fatal(err.Error())
	}
	r := &pb.Value{Kind: &pb.Value_NullValue{}}
	if v.String() != r.String() {
		t.Fatal(v.String())
	}
}

func TestMarshalNumber(t *testing.T) {
	v, err := protobuf.MarshalValue(123)
	if err != nil {
		t.Fatal(err.Error())
	}
	r := &pb.Value{Kind: &pb.Value_NumberValue{NumberValue: 123}}
	if v.String() != r.String() {
		t.Fatal(v.String())
	}
}

func TestMarshalBool(t *testing.T) {
	b := true
	v, err := protobuf.MarshalValue(b)
	if err != nil {
		t.Fatal(err.Error())
	}
	r := &pb.Value{Kind: &pb.Value_BoolValue{BoolValue: b}}
	if v.String() != r.String() {
		t.Fatal(v.String())
	}
}

func TestMarshalString(t *testing.T) {
	s := ""
	v, err := protobuf.MarshalValue(s)
	if err != nil {
		t.Fatal(err.Error())
	}
	r := &pb.Value{Kind: &pb.Value_StringValue{StringValue: s}}
	if v.String() != r.String() {
		t.Fatal(v.String())
	}
}

func TestMarshalUnsupportType(t *testing.T) {
	v, err := protobuf.MarshalValue(make([]string, 0))
	if v != nil || err == nil || err.Error() != "Unsupported marshal type: slice" {
		t.Fatal(err)
	}
	v, err = protobuf.MarshalValue(make(map[string]interface{}))
	if !(v == nil && err != nil && err.Error() == "Unsupported marshal type: map") {
		t.Fatal(err)
	}
	v, err = protobuf.MarshalValue(struct{}{})
	if !(v == nil && err != nil && err.Error() == "Unsupported marshal type: struct") {
		t.Fatal(err)
	}
}
