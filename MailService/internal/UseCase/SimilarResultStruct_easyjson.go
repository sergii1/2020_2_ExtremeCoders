// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package UseCase

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

func easyjsonEf683459DecodeMailServiceInternalUseCase(in *jlexer.Lexer, out *SearchResult) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Senders":
			if in.IsNull() {
				in.Skip()
				out.Senders = nil
			} else {
				in.Delim('[')
				if out.Senders == nil {
					if !in.IsDelim(']') {
						out.Senders = make([]string, 0, 4)
					} else {
						out.Senders = []string{}
					}
				} else {
					out.Senders = (out.Senders)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Senders = append(out.Senders, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "Receivers":
			if in.IsNull() {
				in.Skip()
				out.Receivers = nil
			} else {
				in.Delim('[')
				if out.Receivers == nil {
					if !in.IsDelim(']') {
						out.Receivers = make([]string, 0, 4)
					} else {
						out.Receivers = []string{}
					}
				} else {
					out.Receivers = (out.Receivers)[:0]
				}
				for !in.IsDelim(']') {
					var v2 string
					v2 = string(in.String())
					out.Receivers = append(out.Receivers, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "Themes":
			if in.IsNull() {
				in.Skip()
				out.Themes = nil
			} else {
				in.Delim('[')
				if out.Themes == nil {
					if !in.IsDelim(']') {
						out.Themes = make([]string, 0, 4)
					} else {
						out.Themes = []string{}
					}
				} else {
					out.Themes = (out.Themes)[:0]
				}
				for !in.IsDelim(']') {
					var v3 string
					v3 = string(in.String())
					out.Themes = append(out.Themes, v3)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "Texts":
			if in.IsNull() {
				in.Skip()
				out.Texts = nil
			} else {
				in.Delim('[')
				if out.Texts == nil {
					if !in.IsDelim(']') {
						out.Texts = make([]string, 0, 4)
					} else {
						out.Texts = []string{}
					}
				} else {
					out.Texts = (out.Texts)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.Texts = append(out.Texts, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "SimilarTo":
			out.SimilarTo = string(in.String())
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
func easyjsonEf683459EncodeMailServiceInternalUseCase(out *jwriter.Writer, in SearchResult) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Senders\":"
		out.RawString(prefix[1:])
		if in.Senders == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Senders {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"Receivers\":"
		out.RawString(prefix)
		if in.Receivers == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v7, v8 := range in.Receivers {
				if v7 > 0 {
					out.RawByte(',')
				}
				out.String(string(v8))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"Themes\":"
		out.RawString(prefix)
		if in.Themes == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v9, v10 := range in.Themes {
				if v9 > 0 {
					out.RawByte(',')
				}
				out.String(string(v10))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"Texts\":"
		out.RawString(prefix)
		if in.Texts == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v11, v12 := range in.Texts {
				if v11 > 0 {
					out.RawByte(',')
				}
				out.String(string(v12))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"SimilarTo\":"
		out.RawString(prefix)
		out.String(string(in.SimilarTo))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SearchResult) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEf683459EncodeMailServiceInternalUseCase(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SearchResult) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEf683459EncodeMailServiceInternalUseCase(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SearchResult) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEf683459DecodeMailServiceInternalUseCase(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SearchResult) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEf683459DecodeMailServiceInternalUseCase(l, v)
}