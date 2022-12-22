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

func easyjson9e1087fdDecodeGoParkMailRu20222BugOverloadInternalModels(in *jlexer.Lexer, out *UserActivity) {
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
		case "count_reviews":
			out.CountReviews = int(in.Int())
		case "rating":
			out.Rating = int(in.Int())
		case "date_rating":
			out.DateRating = string(in.String())
		case "collections":
			if in.IsNull() {
				in.Skip()
				out.Collections = nil
			} else {
				in.Delim('[')
				if out.Collections == nil {
					if !in.IsDelim(']') {
						out.Collections = make([]NodeInUserCollection, 0, 2)
					} else {
						out.Collections = []NodeInUserCollection{}
					}
				} else {
					out.Collections = (out.Collections)[:0]
				}
				for !in.IsDelim(']') {
					var v1 NodeInUserCollection
					(v1).UnmarshalEasyJSON(in)
					out.Collections = append(out.Collections, v1)
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
func easyjson9e1087fdEncodeGoParkMailRu20222BugOverloadInternalModels(out *jwriter.Writer, in UserActivity) {
	out.RawByte('{')
	first := true
	_ = first
	if in.CountReviews != 0 {
		const prefix string = ",\"count_reviews\":"
		first = false
		out.RawString(prefix[1:])
		out.Int(int(in.CountReviews))
	}
	if in.Rating != 0 {
		const prefix string = ",\"rating\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Rating))
	}
	if in.DateRating != "" {
		const prefix string = ",\"date_rating\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.DateRating))
	}
	if len(in.Collections) != 0 {
		const prefix string = ",\"collections\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v2, v3 := range in.Collections {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserActivity) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeGoParkMailRu20222BugOverloadInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserActivity) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeGoParkMailRu20222BugOverloadInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserActivity) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeGoParkMailRu20222BugOverloadInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserActivity) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeGoParkMailRu20222BugOverloadInternalModels(l, v)
}
func easyjson9e1087fdDecodeGoParkMailRu20222BugOverloadInternalModels1(in *jlexer.Lexer, out *User) {
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
			out.ID = int(in.Int())
		case "nickname":
			out.Nickname = string(in.String())
		case "email":
			out.Email = string(in.String())
		case "password":
			out.Password = string(in.String())
		case "avatar":
			out.Avatar = string(in.String())
		case "joined_date":
			out.JoinedDate = string(in.String())
		case "count_views_films":
			out.CountViewsFilms = int(in.Int())
		case "count_collections":
			out.CountCollections = int(in.Int())
		case "count_reviews":
			out.CountReviews = int(in.Int())
		case "count_ratings":
			out.CountRatings = int(in.Int())
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
func easyjson9e1087fdEncodeGoParkMailRu20222BugOverloadInternalModels1(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != 0 {
		const prefix string = ",\"user_id\":"
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
	if in.Email != "" {
		const prefix string = ",\"email\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Email))
	}
	if in.Password != "" {
		const prefix string = ",\"password\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Password))
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
	if in.JoinedDate != "" {
		const prefix string = ",\"joined_date\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.JoinedDate))
	}
	if in.CountViewsFilms != 0 {
		const prefix string = ",\"count_views_films\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.CountViewsFilms))
	}
	if in.CountCollections != 0 {
		const prefix string = ",\"count_collections\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.CountCollections))
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
	if in.CountRatings != 0 {
		const prefix string = ",\"count_ratings\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.CountRatings))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeGoParkMailRu20222BugOverloadInternalModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeGoParkMailRu20222BugOverloadInternalModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeGoParkMailRu20222BugOverloadInternalModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeGoParkMailRu20222BugOverloadInternalModels1(l, v)
}
func easyjson9e1087fdDecodeGoParkMailRu20222BugOverloadInternalModels2(in *jlexer.Lexer, out *NodeInUserCollection) {
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
		case "is_used":
			out.IsUsed = bool(in.Bool())
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
func easyjson9e1087fdEncodeGoParkMailRu20222BugOverloadInternalModels2(out *jwriter.Writer, in NodeInUserCollection) {
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
	if in.IsUsed {
		const prefix string = ",\"is_used\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.IsUsed))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v NodeInUserCollection) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeGoParkMailRu20222BugOverloadInternalModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v NodeInUserCollection) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeGoParkMailRu20222BugOverloadInternalModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *NodeInUserCollection) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeGoParkMailRu20222BugOverloadInternalModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *NodeInUserCollection) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeGoParkMailRu20222BugOverloadInternalModels2(l, v)
}