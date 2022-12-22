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

func easyjsonEf032899DecodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels(in *jlexer.Lexer, out *NewFilmReviewRequest) {
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
		case "name":
			out.ReviewName = string(in.String())
		case "type":
			out.ReviewType = string(in.String())
		case "body":
			out.ReviewBody = string(in.String())
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
func easyjsonEf032899EncodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels(out *jwriter.Writer, in NewFilmReviewRequest) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ReviewName != "" {
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ReviewName))
	}
	if in.ReviewType != "" {
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ReviewType))
	}
	if in.ReviewBody != "" {
		const prefix string = ",\"body\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ReviewBody))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v NewFilmReviewRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEf032899EncodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v NewFilmReviewRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEf032899EncodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *NewFilmReviewRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEf032899DecodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *NewFilmReviewRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEf032899DecodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels(l, v)
}
func easyjsonEf032899DecodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels1(in *jlexer.Lexer, out *FilmReviewResponse) {
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
		case "user_id":
			out.UserID = int(in.Int())
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
func easyjsonEf032899EncodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels1(out *jwriter.Writer, in FilmReviewResponse) {
	out.RawByte('{')
	first := true
	_ = first
	if in.UserID != 0 {
		const prefix string = ",\"user_id\":"
		first = false
		out.RawString(prefix[1:])
		out.Int(int(in.UserID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FilmReviewResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEf032899EncodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FilmReviewResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEf032899EncodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FilmReviewResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEf032899DecodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FilmReviewResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEf032899DecodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels1(l, v)
}