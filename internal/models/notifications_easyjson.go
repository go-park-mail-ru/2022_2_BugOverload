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

func easyjsonCcad4d1aDecodeGoParkMailRu20222BugOverloadInternalModels(in *jlexer.Lexer, out *Notification) {
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
		case "action":
			out.Action = string(in.String())
		case "payload":
			if m, ok := out.Payload.(easyjson.Unmarshaler); ok {
				m.UnmarshalEasyJSON(in)
			} else if m, ok := out.Payload.(json.Unmarshaler); ok {
				_ = m.UnmarshalJSON(in.Raw())
			} else {
				out.Payload = in.Interface()
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
func easyjsonCcad4d1aEncodeGoParkMailRu20222BugOverloadInternalModels(out *jwriter.Writer, in Notification) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Action != "" {
		const prefix string = ",\"action\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Action))
	}
	if in.Payload != nil {
		const prefix string = ",\"payload\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if m, ok := in.Payload.(easyjson.Marshaler); ok {
			m.MarshalEasyJSON(out)
		} else if m, ok := in.Payload.(json.Marshaler); ok {
			out.Raw(m.MarshalJSON())
		} else {
			out.Raw(json.Marshal(in.Payload))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Notification) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCcad4d1aEncodeGoParkMailRu20222BugOverloadInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Notification) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCcad4d1aEncodeGoParkMailRu20222BugOverloadInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Notification) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCcad4d1aDecodeGoParkMailRu20222BugOverloadInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Notification) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCcad4d1aDecodeGoParkMailRu20222BugOverloadInternalModels(l, v)
}
