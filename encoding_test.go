package encoding

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name    string
		traceID string
		low     uint64
		high    uint64
		so      func(actual interface{}, expected ...interface{}) string
		se      func(actual interface{}, expected ...interface{}) string
	}{
		{name: "large format", traceID: "000000000000000f000000000000000a", low: 10, high: 15, so: ShouldBeNil, se: ShouldEqual},
		{name: "large format zero lower", traceID: "000000000000000a0000000000000000", low: 0, high: 10, so: ShouldBeNil, se: ShouldEqual},
		{name: "small format", traceID: "000000000000000f", low: 15, high: 0, so: ShouldBeNil, se: ShouldEqual},
		{name: "bad format", traceID: "blarg", low: 0, high: 0, so: ShouldNotBeNil, se: ShouldNotEqual},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			Convey("test stuff", t, func() {
				id, err := GetID(test.traceID)
				So(err, test.so)
				So(id[0], ShouldEqual, test.low)
				So(id[1], ShouldEqual, test.high)
				newid := id.String()
				So(newid, test.se, test.traceID)
			})
		})
	}
}

func TestAdditionalDims(t *testing.T) {
	Convey("test additional dims", t, func() {
		si := &SpanIdentity{Service: "service", Operation: "operation"}
		nsi := NewExtendedSpanIdentity(si, map[string]string{"foo": "bar", "bar": "baz"})
		So(nsi.Dims(), ShouldResemble, map[string]string{"foo": "bar", "bar": "baz", "service": "service", "operation": "operation", "sf_dimensionalized": "true", "error": "false"})
	})
	Convey("test forbidden dims", t, func() {
		si := &SpanIdentity{Service: "service", Operation: "operation"}
		nsi := NewExtendedSpanIdentity(si, map[string]string{"service": "foo", "operation": "baz"})
		So(nsi.Dims(), ShouldResemble, map[string]string{"service": "service", "operation": "operation", "sf_dimensionalized": "true", "error": "false"})
	})
}
