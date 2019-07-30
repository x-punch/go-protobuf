package protobuf

// wkt represents well known type interface
type wkt interface {
	XXX_WellKnownType() string
}
