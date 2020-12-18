// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package conf

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonEae1800fDecodeGitlabEeFunplusIoWatcherWatcherArchivistRuntimeGoOverwriteConf(in *jlexer.Lexer, out *GConf) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
	} else {
		in.Delim('{')
		*out = make(GConf)
		for !in.IsDelim('}') {
			key := string(in.String())
			in.WantColon()
			var v1 int64
			v1 = int64(in.Int64())
			(*out)[key] = v1
			in.WantComma()
		}
		in.Delim('}')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonEae1800fEncodeGitlabEeFunplusIoWatcherWatcherArchivistRuntimeGoOverwriteConf(out *jwriter.Writer, in GConf) {
	if in == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v2First := true
		for v2Name, v2Value := range in {
			if v2First {
				v2First = false
			} else {
				out.RawByte(',')
			}
			out.String(string(v2Name))
			out.RawByte(':')
			out.Int64(int64(v2Value))
		}
		out.RawByte('}')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v GConf) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEae1800fEncodeGitlabEeFunplusIoWatcherWatcherArchivistRuntimeGoOverwriteConf(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GConf) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEae1800fEncodeGitlabEeFunplusIoWatcherWatcherArchivistRuntimeGoOverwriteConf(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GConf) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEae1800fDecodeGitlabEeFunplusIoWatcherWatcherArchivistRuntimeGoOverwriteConf(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GConf) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEae1800fDecodeGitlabEeFunplusIoWatcherWatcherArchivistRuntimeGoOverwriteConf(l, v)
}