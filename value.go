package protobuf

import (
	"reflect"

	pb "github.com/golang/protobuf/ptypes/struct"
)

// GetValue represents get value from protobuf value type
func GetValue(v *pb.Value) interface{} {
	if v == nil {
		return nil
	}
	k := reflect.ValueOf(v).Elem().Field(0).Elem().Elem().Field(0)
	if wkt, ok := k.Interface().(wkt); ok {
		switch wkt.XXX_WellKnownType() {
		case "NullValue":
			return nil
		case "Struct":
			s := v.GetStructValue()
			if s == nil {
				return nil
			}
			vals := make(map[string]interface{})
			for k, v := range s.Fields {
				vals[k] = GetValue(v)
			}
			return vals
		case "ListValue":
			s := v.GetListValue()
			if s == nil {
				return nil
			}
			vals := make([]interface{}, 0, len(s.Values))
			for _, v := range s.Values {
				vals = append(vals, GetValue(v))
			}
			return vals
		}
	}
	return k.Interface()
}
