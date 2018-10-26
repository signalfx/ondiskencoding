// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package encoding

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	trace "github.com/signalfx/golib/trace"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonE453ad8fDecodeGithubComSignalfxOndiskencoding(in *jlexer.Lexer, out *SpanIdentity) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Service":
			out.Service = string(in.String())
		case "Operation":
			out.Operation = string(in.String())
		case "Error":
			out.Error = bool(in.Bool())
		case "HttpMethod":
			out.HttpMethod = string(in.String())
		case "Kind":
			out.Kind = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonE453ad8fEncodeGithubComSignalfxOndiskencoding(out *jwriter.Writer, in SpanIdentity) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Service != "" {
		const prefix string = ",\"Service\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Service))
	}
	if in.Operation != "" {
		const prefix string = ",\"Operation\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Operation))
	}
	if in.Error {
		const prefix string = ",\"Error\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Error))
	}
	if in.HttpMethod != "" {
		const prefix string = ",\"HttpMethod\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.HttpMethod))
	}
	if in.Kind != "" {
		const prefix string = ",\"Kind\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Kind))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SpanIdentity) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE453ad8fEncodeGithubComSignalfxOndiskencoding(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SpanIdentity) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE453ad8fEncodeGithubComSignalfxOndiskencoding(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SpanIdentity) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE453ad8fDecodeGithubComSignalfxOndiskencoding(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SpanIdentity) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE453ad8fDecodeGithubComSignalfxOndiskencoding(l, v)
}
func easyjsonE453ad8fDecodeGithubComSignalfxOndiskencoding1(in *jlexer.Lexer, out *HistoOnDiskEntry) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Digest":
			if in.IsNull() {
				in.Skip()
				out.Digest = nil
			} else {
				out.Digest = in.Bytes()
			}
		case "Last":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Last).UnmarshalJSON(data))
			}
		case "Count":
			out.Count = int64(in.Int64())
		case "DecayedCount":
			out.DecayedCount = float64(in.Float64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonE453ad8fEncodeGithubComSignalfxOndiskencoding1(out *jwriter.Writer, in HistoOnDiskEntry) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Digest) != 0 {
		const prefix string = ",\"Digest\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Base64Bytes(in.Digest)
	}
	if true {
		const prefix string = ",\"Last\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Last).MarshalJSON())
	}
	if in.Count != 0 {
		const prefix string = ",\"Count\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Count))
	}
	if in.DecayedCount != 0 {
		const prefix string = ",\"DecayedCount\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.DecayedCount))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v HistoOnDiskEntry) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE453ad8fEncodeGithubComSignalfxOndiskencoding1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v HistoOnDiskEntry) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE453ad8fEncodeGithubComSignalfxOndiskencoding1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *HistoOnDiskEntry) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE453ad8fDecodeGithubComSignalfxOndiskencoding1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *HistoOnDiskEntry) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE453ad8fDecodeGithubComSignalfxOndiskencoding1(l, v)
}
func easyjsonE453ad8fDecodeGithubComSignalfxOndiskencoding2(in *jlexer.Lexer, out *HistoOnDisk) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Entries":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Entries = make(map[SpanIdentity]HistoOnDiskEntry)
				} else {
					out.Entries = nil
				}
				for !in.IsDelim('}') {
					var key SpanIdentity
					(key).UnmarshalEasyJSON(in)
					in.WantColon()
					var v4 HistoOnDiskEntry
					(v4).UnmarshalEasyJSON(in)
					(out.Entries)[key] = v4
					in.WantComma()
				}
				in.Delim('}')
			}
		case "MetricsReservoirSize":
			out.MetricsReservoirSize = int(in.Int())
		case "MetricsAlphaFactor":
			out.MetricsAlphaFactor = float64(in.Float64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonE453ad8fEncodeGithubComSignalfxOndiskencoding2(out *jwriter.Writer, in HistoOnDisk) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Entries) != 0 {
		const prefix string = ",\"Entries\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v5First := true
			for v5Name, v5Value := range in.Entries {
				if v5First {
					v5First = false
				} else {
					out.RawByte(',')
				}
				(v5Name).MarshalEasyJSON(out)
				out.RawByte(':')
				(v5Value).MarshalEasyJSON(out)
			}
			out.RawByte('}')
		}
	}
	if in.MetricsReservoirSize != 0 {
		const prefix string = ",\"MetricsReservoirSize\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.MetricsReservoirSize))
	}
	if in.MetricsAlphaFactor != 0 {
		const prefix string = ",\"MetricsAlphaFactor\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.MetricsAlphaFactor))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v HistoOnDisk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE453ad8fEncodeGithubComSignalfxOndiskencoding2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v HistoOnDisk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE453ad8fEncodeGithubComSignalfxOndiskencoding2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *HistoOnDisk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE453ad8fDecodeGithubComSignalfxOndiskencoding2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *HistoOnDisk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE453ad8fDecodeGithubComSignalfxOndiskencoding2(l, v)
}
func easyjsonE453ad8fDecodeGithubComSignalfxOndiskencoding3(in *jlexer.Lexer, out *BufferOnDisk) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Traces":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Traces = make(map[string]*BufferEntry)
				} else {
					out.Traces = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v6 *BufferEntry
					if in.IsNull() {
						in.Skip()
						v6 = nil
					} else {
						if v6 == nil {
							v6 = new(BufferEntry)
						}
						(*v6).UnmarshalEasyJSON(in)
					}
					(out.Traces)[key] = v6
					in.WantComma()
				}
				in.Delim('}')
			}
		case "NumSpans":
			out.NumSpans = int64(in.Int64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonE453ad8fEncodeGithubComSignalfxOndiskencoding3(out *jwriter.Writer, in BufferOnDisk) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Traces) != 0 {
		const prefix string = ",\"Traces\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v7First := true
			for v7Name, v7Value := range in.Traces {
				if v7First {
					v7First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v7Name))
				out.RawByte(':')
				if v7Value == nil {
					out.RawString("null")
				} else {
					(*v7Value).MarshalEasyJSON(out)
				}
			}
			out.RawByte('}')
		}
	}
	if in.NumSpans != 0 {
		const prefix string = ",\"NumSpans\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.NumSpans))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BufferOnDisk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE453ad8fEncodeGithubComSignalfxOndiskencoding3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BufferOnDisk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE453ad8fEncodeGithubComSignalfxOndiskencoding3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BufferOnDisk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE453ad8fDecodeGithubComSignalfxOndiskencoding3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BufferOnDisk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE453ad8fDecodeGithubComSignalfxOndiskencoding3(l, v)
}
func easyjsonE453ad8fDecodeGithubComSignalfxOndiskencoding4(in *jlexer.Lexer, out *BufferEntry) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Spans":
			if in.IsNull() {
				in.Skip()
				out.Spans = nil
			} else {
				in.Delim('[')
				if out.Spans == nil {
					if !in.IsDelim(']') {
						out.Spans = make([]*trace.Span, 0, 8)
					} else {
						out.Spans = []*trace.Span{}
					}
				} else {
					out.Spans = (out.Spans)[:0]
				}
				for !in.IsDelim(']') {
					var v8 *trace.Span
					if in.IsNull() {
						in.Skip()
						v8 = nil
					} else {
						if v8 == nil {
							v8 = new(trace.Span)
						}
						easyjsonE453ad8fDecodeGithubComSignalfxGolibTrace(in, &*v8)
					}
					out.Spans = append(out.Spans, v8)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "Last":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Last).UnmarshalJSON(data))
			}
		case "ToBeReleased":
			out.ToBeReleased = bool(in.Bool())
		case "ReleaseImmediately":
			out.ReleaseImmediately = bool(in.Bool())
		case "LatestEndTime":
			out.LatestEndTime = float64(in.Float64())
		case "StartTime":
			out.StartTime = float64(in.Float64())
		case "Initiating":
			if in.IsNull() {
				in.Skip()
				out.Initiating = nil
			} else {
				if out.Initiating == nil {
					out.Initiating = new(trace.Span)
				}
				easyjsonE453ad8fDecodeGithubComSignalfxGolibTrace(in, &*out.Initiating)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonE453ad8fEncodeGithubComSignalfxOndiskencoding4(out *jwriter.Writer, in BufferEntry) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Spans) != 0 {
		const prefix string = ",\"Spans\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v9, v10 := range in.Spans {
				if v9 > 0 {
					out.RawByte(',')
				}
				if v10 == nil {
					out.RawString("null")
				} else {
					easyjsonE453ad8fEncodeGithubComSignalfxGolibTrace(out, *v10)
				}
			}
			out.RawByte(']')
		}
	}
	if true {
		const prefix string = ",\"Last\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Last).MarshalJSON())
	}
	if in.ToBeReleased {
		const prefix string = ",\"ToBeReleased\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.ToBeReleased))
	}
	if in.ReleaseImmediately {
		const prefix string = ",\"ReleaseImmediately\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.ReleaseImmediately))
	}
	if in.LatestEndTime != 0 {
		const prefix string = ",\"LatestEndTime\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.LatestEndTime))
	}
	if in.StartTime != 0 {
		const prefix string = ",\"StartTime\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.StartTime))
	}
	if in.Initiating != nil {
		const prefix string = ",\"Initiating\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjsonE453ad8fEncodeGithubComSignalfxGolibTrace(out, *in.Initiating)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BufferEntry) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE453ad8fEncodeGithubComSignalfxOndiskencoding4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BufferEntry) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE453ad8fEncodeGithubComSignalfxOndiskencoding4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BufferEntry) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE453ad8fDecodeGithubComSignalfxOndiskencoding4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BufferEntry) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE453ad8fDecodeGithubComSignalfxOndiskencoding4(l, v)
}
func easyjsonE453ad8fDecodeGithubComSignalfxGolibTrace(in *jlexer.Lexer, out *trace.Span) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "traceId":
			out.TraceID = string(in.String())
		case "name":
			if in.IsNull() {
				in.Skip()
				out.Name = nil
			} else {
				if out.Name == nil {
					out.Name = new(string)
				}
				*out.Name = string(in.String())
			}
		case "parentId":
			if in.IsNull() {
				in.Skip()
				out.ParentID = nil
			} else {
				if out.ParentID == nil {
					out.ParentID = new(string)
				}
				*out.ParentID = string(in.String())
			}
		case "id":
			out.ID = string(in.String())
		case "kind":
			if in.IsNull() {
				in.Skip()
				out.Kind = nil
			} else {
				if out.Kind == nil {
					out.Kind = new(string)
				}
				*out.Kind = string(in.String())
			}
		case "timestamp":
			if in.IsNull() {
				in.Skip()
				out.Timestamp = nil
			} else {
				if out.Timestamp == nil {
					out.Timestamp = new(float64)
				}
				*out.Timestamp = float64(in.Float64())
			}
		case "duration":
			if in.IsNull() {
				in.Skip()
				out.Duration = nil
			} else {
				if out.Duration == nil {
					out.Duration = new(float64)
				}
				*out.Duration = float64(in.Float64())
			}
		case "debug":
			if in.IsNull() {
				in.Skip()
				out.Debug = nil
			} else {
				if out.Debug == nil {
					out.Debug = new(bool)
				}
				*out.Debug = bool(in.Bool())
			}
		case "shared":
			if in.IsNull() {
				in.Skip()
				out.Shared = nil
			} else {
				if out.Shared == nil {
					out.Shared = new(bool)
				}
				*out.Shared = bool(in.Bool())
			}
		case "localEndpoint":
			if in.IsNull() {
				in.Skip()
				out.LocalEndpoint = nil
			} else {
				if out.LocalEndpoint == nil {
					out.LocalEndpoint = new(trace.Endpoint)
				}
				easyjsonE453ad8fDecodeGithubComSignalfxGolibTrace1(in, &*out.LocalEndpoint)
			}
		case "remoteEndpoint":
			if in.IsNull() {
				in.Skip()
				out.RemoteEndpoint = nil
			} else {
				if out.RemoteEndpoint == nil {
					out.RemoteEndpoint = new(trace.Endpoint)
				}
				easyjsonE453ad8fDecodeGithubComSignalfxGolibTrace1(in, &*out.RemoteEndpoint)
			}
		case "annotations":
			if in.IsNull() {
				in.Skip()
				out.Annotations = nil
			} else {
				in.Delim('[')
				if out.Annotations == nil {
					if !in.IsDelim(']') {
						out.Annotations = make([]*trace.Annotation, 0, 8)
					} else {
						out.Annotations = []*trace.Annotation{}
					}
				} else {
					out.Annotations = (out.Annotations)[:0]
				}
				for !in.IsDelim(']') {
					var v11 *trace.Annotation
					if in.IsNull() {
						in.Skip()
						v11 = nil
					} else {
						if v11 == nil {
							v11 = new(trace.Annotation)
						}
						easyjsonE453ad8fDecodeGithubComSignalfxGolibTrace2(in, &*v11)
					}
					out.Annotations = append(out.Annotations, v11)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "tags":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Tags = make(map[string]string)
				} else {
					out.Tags = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v12 string
					v12 = string(in.String())
					(out.Tags)[key] = v12
					in.WantComma()
				}
				in.Delim('}')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonE453ad8fEncodeGithubComSignalfxGolibTrace(out *jwriter.Writer, in trace.Span) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"traceId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.TraceID))
	}
	{
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Name == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Name))
		}
	}
	{
		const prefix string = ",\"parentId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.ParentID == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.ParentID))
		}
	}
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"kind\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Kind == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Kind))
		}
	}
	{
		const prefix string = ",\"timestamp\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Timestamp == nil {
			out.RawString("null")
		} else {
			out.Float64(float64(*in.Timestamp))
		}
	}
	{
		const prefix string = ",\"duration\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Duration == nil {
			out.RawString("null")
		} else {
			out.Float64(float64(*in.Duration))
		}
	}
	{
		const prefix string = ",\"debug\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Debug == nil {
			out.RawString("null")
		} else {
			out.Bool(bool(*in.Debug))
		}
	}
	{
		const prefix string = ",\"shared\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Shared == nil {
			out.RawString("null")
		} else {
			out.Bool(bool(*in.Shared))
		}
	}
	{
		const prefix string = ",\"localEndpoint\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.LocalEndpoint == nil {
			out.RawString("null")
		} else {
			easyjsonE453ad8fEncodeGithubComSignalfxGolibTrace1(out, *in.LocalEndpoint)
		}
	}
	{
		const prefix string = ",\"remoteEndpoint\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.RemoteEndpoint == nil {
			out.RawString("null")
		} else {
			easyjsonE453ad8fEncodeGithubComSignalfxGolibTrace1(out, *in.RemoteEndpoint)
		}
	}
	{
		const prefix string = ",\"annotations\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Annotations == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v13, v14 := range in.Annotations {
				if v13 > 0 {
					out.RawByte(',')
				}
				if v14 == nil {
					out.RawString("null")
				} else {
					easyjsonE453ad8fEncodeGithubComSignalfxGolibTrace2(out, *v14)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"tags\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Tags == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v15First := true
			for v15Name, v15Value := range in.Tags {
				if v15First {
					v15First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v15Name))
				out.RawByte(':')
				out.String(string(v15Value))
			}
			out.RawByte('}')
		}
	}
	out.RawByte('}')
}
func easyjsonE453ad8fDecodeGithubComSignalfxGolibTrace2(in *jlexer.Lexer, out *trace.Annotation) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "timestamp":
			if in.IsNull() {
				in.Skip()
				out.Timestamp = nil
			} else {
				if out.Timestamp == nil {
					out.Timestamp = new(float64)
				}
				*out.Timestamp = float64(in.Float64())
			}
		case "value":
			if in.IsNull() {
				in.Skip()
				out.Value = nil
			} else {
				if out.Value == nil {
					out.Value = new(string)
				}
				*out.Value = string(in.String())
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonE453ad8fEncodeGithubComSignalfxGolibTrace2(out *jwriter.Writer, in trace.Annotation) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"timestamp\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Timestamp == nil {
			out.RawString("null")
		} else {
			out.Float64(float64(*in.Timestamp))
		}
	}
	{
		const prefix string = ",\"value\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Value == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Value))
		}
	}
	out.RawByte('}')
}
func easyjsonE453ad8fDecodeGithubComSignalfxGolibTrace1(in *jlexer.Lexer, out *trace.Endpoint) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "serviceName":
			if in.IsNull() {
				in.Skip()
				out.ServiceName = nil
			} else {
				if out.ServiceName == nil {
					out.ServiceName = new(string)
				}
				*out.ServiceName = string(in.String())
			}
		case "ipv4":
			if in.IsNull() {
				in.Skip()
				out.Ipv4 = nil
			} else {
				if out.Ipv4 == nil {
					out.Ipv4 = new(string)
				}
				*out.Ipv4 = string(in.String())
			}
		case "ipv6":
			if in.IsNull() {
				in.Skip()
				out.Ipv6 = nil
			} else {
				if out.Ipv6 == nil {
					out.Ipv6 = new(string)
				}
				*out.Ipv6 = string(in.String())
			}
		case "port":
			if in.IsNull() {
				in.Skip()
				out.Port = nil
			} else {
				if out.Port == nil {
					out.Port = new(int32)
				}
				*out.Port = int32(in.Int32())
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonE453ad8fEncodeGithubComSignalfxGolibTrace1(out *jwriter.Writer, in trace.Endpoint) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"serviceName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.ServiceName == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.ServiceName))
		}
	}
	{
		const prefix string = ",\"ipv4\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Ipv4 == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Ipv4))
		}
	}
	{
		const prefix string = ",\"ipv6\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Ipv6 == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Ipv6))
		}
	}
	{
		const prefix string = ",\"port\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Port == nil {
			out.RawString("null")
		} else {
			out.Int32(int32(*in.Port))
		}
	}
	out.RawByte('}')
}
func easyjsonE453ad8fDecodeGithubComSignalfxOndiskencoding5(in *jlexer.Lexer, out *BufferEntries) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(BufferEntries, 0, 8)
			} else {
				*out = BufferEntries{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v16 *BufferEntry
			if in.IsNull() {
				in.Skip()
				v16 = nil
			} else {
				if v16 == nil {
					v16 = new(BufferEntry)
				}
				(*v16).UnmarshalEasyJSON(in)
			}
			*out = append(*out, v16)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonE453ad8fEncodeGithubComSignalfxOndiskencoding5(out *jwriter.Writer, in BufferEntries) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v17, v18 := range in {
			if v17 > 0 {
				out.RawByte(',')
			}
			if v18 == nil {
				out.RawString("null")
			} else {
				(*v18).MarshalEasyJSON(out)
			}
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v BufferEntries) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE453ad8fEncodeGithubComSignalfxOndiskencoding5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BufferEntries) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE453ad8fEncodeGithubComSignalfxOndiskencoding5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BufferEntries) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE453ad8fDecodeGithubComSignalfxOndiskencoding5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BufferEntries) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE453ad8fDecodeGithubComSignalfxOndiskencoding5(l, v)
}
