package protobuf_test

import (
	"testing"

	pb "github.com/golang/protobuf/ptypes/struct"
	"github.com/x-punch/go-protobuf"
)

func TestNumberValue(t *testing.T) {
	num := float64(123)
	p := &pb.Value{Kind: &pb.Value_NumberValue{NumberValue: num}}
	v := protobuf.GetValue(p)
	if v != num {
		t.Fail()
	}
}

func TestStringValue(t *testing.T) {
	str := "string"
	p := &pb.Value{Kind: &pb.Value_StringValue{StringValue: str}}
	v := protobuf.GetValue(p)
	if v != str {
		t.Fail()
	}
}

func TestNullValue(t *testing.T) {
	p := &pb.Value{Kind: &pb.Value_NullValue{}}
	v := protobuf.GetValue(p)
	if v != nil {
		t.Fail()
	}
}

func TestBoolValue(t *testing.T) {
	p := &pb.Value{Kind: &pb.Value_BoolValue{BoolValue: true}}
	v := protobuf.GetValue(p)
	if v != true {
		t.Fail()
	}
}

func TestStructValue(t *testing.T) {
	v1 := float64(123)
	v2 := "strings"
	p := &pb.Value{Kind: &pb.Value_StructValue{StructValue: &pb.Struct{
		Fields: map[string]*pb.Value{
			"v1": &pb.Value{Kind: &pb.Value_NumberValue{NumberValue: v1}},
			"v2": &pb.Value{Kind: &pb.Value_StringValue{StringValue: v2}},
		},
	}}}
	v := protobuf.GetValue(p)
	if maps, ok := v.(map[string]interface{}); ok {
		if len(maps) == 2 && maps["v1"] == v1 && maps["v2"] == v2 {
			return
		}
	}
	t.Fail()
}

func TestListValue(t *testing.T) {
	v1 := float64(123)
	v2 := "strings"
	p := &pb.Value{Kind: &pb.Value_ListValue{ListValue: &pb.ListValue{
		Values: []*pb.Value{
			&pb.Value{Kind: &pb.Value_NumberValue{NumberValue: v1}},
			&pb.Value{Kind: &pb.Value_StringValue{StringValue: v2}},
		},
	}}}
	v := protobuf.GetValue(p)
	if array, ok := v.([]interface{}); ok {
		if len(array) == 2 && array[0] == v1 && array[1] == v2 {
			return
		}
	}
	t.Fail()
}
