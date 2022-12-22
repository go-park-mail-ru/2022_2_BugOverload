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

func easyjson2f096870DecodeGoParkMailRu20222BugOverloadInternalModels(in *jlexer.Lexer, out *Review) {
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
		case "id":
			out.ID = int(in.Int())
		case "name":
			out.Name = string(in.String())
		case "type":
			out.Type = string(in.String())
		case "body":
			out.Body = string(in.String())
		case "count_likes":
			out.CountLikes = int(in.Int())
		case "create_time":
			out.CreateTime = string(in.String())
		case "author":
			(out.Author).UnmarshalEasyJSON(in)
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
func easyjson2f096870EncodeGoParkMailRu20222BugOverloadInternalModels(out *jwriter.Writer, in Review) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != 0 {
		const prefix string = ",\"id\":"
		first = false
		out.RawString(prefix[1:])
		out.Int(int(in.ID))
	}
	if in.Name != "" {
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	if in.Type != "" {
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Type))
	}
	if in.Body != "" {
		const prefix string = ",\"body\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Body))
	}
	if in.CountLikes != 0 {
		const prefix string = ",\"count_likes\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.CountLikes))
	}
	if in.CreateTime != "" {
		const prefix string = ",\"create_time\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.CreateTime))
	}
	if true {
		const prefix string = ",\"author\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Author).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Review) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson2f096870EncodeGoParkMailRu20222BugOverloadInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Review) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson2f096870EncodeGoParkMailRu20222BugOverloadInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Review) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson2f096870DecodeGoParkMailRu20222BugOverloadInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Review) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson2f096870DecodeGoParkMailRu20222BugOverloadInternalModels(l, v)
}
