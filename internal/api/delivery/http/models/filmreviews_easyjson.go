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

func easyjson7b77f903DecodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels(in *jlexer.Lexer, out *ReviewResponse) {
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
			out.Name = string(in.String())
		case "type":
			out.Type = string(in.String())
		case "create_time":
			out.CreateTime = string(in.String())
		case "body":
			out.Body = string(in.String())
		case "count_likes":
			out.CountLikes = int(in.Int())
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
func easyjson7b77f903EncodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels(out *jwriter.Writer, in ReviewResponse) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Name != "" {
		const prefix string = ",\"name\":"
		first = false
		out.RawString(prefix[1:])
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
func (v ReviewResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7b77f903EncodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ReviewResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7b77f903EncodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ReviewResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7b77f903DecodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ReviewResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7b77f903DecodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels(l, v)
}
func easyjson7b77f903DecodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels1(in *jlexer.Lexer, out *ReviewList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(ReviewList, 0, 0)
			} else {
				*out = ReviewList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 ReviewResponse
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson7b77f903EncodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels1(out *jwriter.Writer, in ReviewList) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v ReviewList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7b77f903EncodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ReviewList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7b77f903EncodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ReviewList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7b77f903DecodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ReviewList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7b77f903DecodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels1(l, v)
}
func easyjson7b77f903DecodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels2(in *jlexer.Lexer, out *ReviewAuthorResponse) {
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
		case "nickname":
			out.Nickname = string(in.String())
		case "count_reviews":
			out.CountReviews = int(in.Int())
		case "avatar":
			out.Avatar = string(in.String())
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
func easyjson7b77f903EncodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels2(out *jwriter.Writer, in ReviewAuthorResponse) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != 0 {
		const prefix string = ",\"id\":"
		first = false
		out.RawString(prefix[1:])
		out.Int(int(in.ID))
	}
	if in.Nickname != "" {
		const prefix string = ",\"nickname\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Nickname))
	}
	if in.CountReviews != 0 {
		const prefix string = ",\"count_reviews\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.CountReviews))
	}
	if in.Avatar != "" {
		const prefix string = ",\"avatar\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Avatar))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ReviewAuthorResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7b77f903EncodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ReviewAuthorResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7b77f903EncodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ReviewAuthorResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7b77f903DecodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ReviewAuthorResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7b77f903DecodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels2(l, v)
}