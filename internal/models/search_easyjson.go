// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

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

func easyjsonD4176298DecodeGoParkMailRu20222BugOverloadInternalModels(in *jlexer.Lexer, out *Search) {
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
		case "films":
			if in.IsNull() {
				in.Skip()
				out.Films = nil
			} else {
				in.Delim('[')
				if out.Films == nil {
					if !in.IsDelim(']') {
						out.Films = make([]Film, 0, 0)
					} else {
						out.Films = []Film{}
					}
				} else {
					out.Films = (out.Films)[:0]
				}
				for !in.IsDelim(']') {
					var v1 Film
					(v1).UnmarshalEasyJSON(in)
					out.Films = append(out.Films, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "serials":
			if in.IsNull() {
				in.Skip()
				out.Serials = nil
			} else {
				in.Delim('[')
				if out.Serials == nil {
					if !in.IsDelim(']') {
						out.Serials = make([]Film, 0, 0)
					} else {
						out.Serials = []Film{}
					}
				} else {
					out.Serials = (out.Serials)[:0]
				}
				for !in.IsDelim(']') {
					var v2 Film
					(v2).UnmarshalEasyJSON(in)
					out.Serials = append(out.Serials, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "persons":
			if in.IsNull() {
				in.Skip()
				out.Persons = nil
			} else {
				in.Delim('[')
				if out.Persons == nil {
					if !in.IsDelim(']') {
						out.Persons = make([]Person, 0, 0)
					} else {
						out.Persons = []Person{}
					}
				} else {
					out.Persons = (out.Persons)[:0]
				}
				for !in.IsDelim(']') {
					var v3 Person
					(v3).UnmarshalEasyJSON(in)
					out.Persons = append(out.Persons, v3)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.AddError(&jlexer.LexerError{
				Offset: in.GetPos(),
				Reason: "unknown field",
				Data:   key,
			})
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD4176298EncodeGoParkMailRu20222BugOverloadInternalModels(out *jwriter.Writer, in Search) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Films) != 0 {
		const prefix string = ",\"films\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('[')
			for v4, v5 := range in.Films {
				if v4 > 0 {
					out.RawByte(',')
				}
				(v5).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	if len(in.Serials) != 0 {
		const prefix string = ",\"serials\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v6, v7 := range in.Serials {
				if v6 > 0 {
					out.RawByte(',')
				}
				(v7).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	if len(in.Persons) != 0 {
		const prefix string = ",\"persons\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v8, v9 := range in.Persons {
				if v8 > 0 {
					out.RawByte(',')
				}
				(v9).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Search) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD4176298EncodeGoParkMailRu20222BugOverloadInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Search) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD4176298EncodeGoParkMailRu20222BugOverloadInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Search) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD4176298DecodeGoParkMailRu20222BugOverloadInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Search) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD4176298DecodeGoParkMailRu20222BugOverloadInternalModels(l, v)
}