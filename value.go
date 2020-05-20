package protobuf

import (
	"errors"
	"reflect"

	pb "github.com/golang/protobuf/ptypes/struct"
)

// UnmarshalValue convert protobuf value type into generic value type
func UnmarshalValue(v *pb.Value) interface{} {
	if v == nil {
		return nil
	}
	value := reflect.ValueOf(v.Kind).Elem().Field(0).Interface()
	switch value.(type) {
	case pb.NullValue: //compatiable with golang/protobuf@1.3.x
		return nil
	case *pb.NullValue:
		return nil
	case *pb.Struct:
		s := v.GetStructValue()
		if s == nil {
			return nil
		}
		vals := make(map[string]interface{})
		for k, v := range s.Fields {
			vals[k] = UnmarshalValue(v)
		}
		return vals
	case *pb.ListValue:
		s := v.GetListValue()
		if s == nil {
			return nil
		}
		vals := make([]interface{}, 0, len(s.Values))
		for _, v := range s.Values {
			vals = append(vals, UnmarshalValue(v))
		}
		return vals
	}
	return value
}

// MarshalValue convert generic type into protobuf value type
// Only support NullValue, NumberValue, BoolValue, StringValue
func MarshalValue(input interface{}) (*pb.Value, error) {
	if input == nil {
		return &pb.Value{Kind: &pb.Value_NullValue{}}, nil
	}
	value := reflect.ValueOf(input)
	switch value.Kind() {
	case reflect.Bool:
		return &pb.Value{Kind: &pb.Value_BoolValue{BoolValue: value.Bool()}}, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return &pb.Value{Kind: &pb.Value_NumberValue{NumberValue: float64(value.Int())}}, nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return &pb.Value{Kind: &pb.Value_NumberValue{NumberValue: float64(value.Uint())}}, nil
	case reflect.Float32, reflect.Float64:
		return &pb.Value{Kind: &pb.Value_NumberValue{NumberValue: value.Float()}}, nil
	case reflect.String:
		return &pb.Value{Kind: &pb.Value_StringValue{StringValue: value.String()}}, nil
	default:
		return nil, errors.New("Unsupported marshal type: " + value.Kind().String())
	}
}
