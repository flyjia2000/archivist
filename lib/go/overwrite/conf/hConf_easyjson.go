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

func easyjsonA1edf190DecodeGithubComKingsgrouposArchivistLibGoOverwriteConf(in *jlexer.Lexer, out *HConf) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
	} else {
		in.Delim('{')
		*out = make(HConf)
		for !in.IsDelim('}') {
			key := string(in.String())
			in.WantColon()
			var v1 []int64
			if in.IsNull() {
				in.Skip()
				v1 = nil
			} else {
				in.Delim('[')
				if v1 == nil {
					if !in.IsDelim(']') {
						v1 = make([]int64, 0, 8)
					} else {
						v1 = []int64{}
					}
				} else {
					v1 = (v1)[:0]
				}
				for !in.IsDelim(']') {
					var v2 int64
					v2 = int64(in.Int64())
					v1 = append(v1, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
			(*out)[key] = v1
			in.WantComma()
		}
		in.Delim('}')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonA1edf190EncodeGithubComKingsgrouposArchivistLibGoOverwriteConf(out *jwriter.Writer, in HConf) {
	if in == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v3First := true
		for v3Name, v3Value := range in {
			if v3First {
				v3First = false
			} else {
				out.RawByte(',')
			}
			out.String(string(v3Name))
			out.RawByte(':')
			if v3Value == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
				out.RawString("null")
			} else {
				out.RawByte('[')
				for v4, v5 := range v3Value {
					if v4 > 0 {
						out.RawByte(',')
					}
					out.Int64(int64(v5))
				}
				out.RawByte(']')
			}
		}
		out.RawByte('}')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v HConf) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA1edf190EncodeGithubComKingsgrouposArchivistLibGoOverwriteConf(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v HConf) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA1edf190EncodeGithubComKingsgrouposArchivistLibGoOverwriteConf(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *HConf) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA1edf190DecodeGithubComKingsgrouposArchivistLibGoOverwriteConf(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *HConf) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA1edf190DecodeGithubComKingsgrouposArchivistLibGoOverwriteConf(l, v)
}
