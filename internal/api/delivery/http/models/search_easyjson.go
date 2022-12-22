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

func easyjsonD4176298DecodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels(in *jlexer.Lexer, out *SearchResponse) {
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
						out.Films = make([]SearchFilmResponse, 0, 0)
					} else {
						out.Films = []SearchFilmResponse{}
					}
				} else {
					out.Films = (out.Films)[:0]
				}
				for !in.IsDelim(']') {
					var v1 SearchFilmResponse
					(v1).UnmarshalEasyJSON(in)
					out.Films = append(out.Films, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "serials":
			if in.IsNull() {
				in.Skip()
				out.Series = nil
			} else {
				in.Delim('[')
				if out.Series == nil {
					if !in.IsDelim(']') {
						out.Series = make([]SearchFilmResponse, 0, 0)
					} else {
						out.Series = []SearchFilmResponse{}
					}
				} else {
					out.Series = (out.Series)[:0]
				}
				for !in.IsDelim(']') {
					var v2 SearchFilmResponse
					(v2).UnmarshalEasyJSON(in)
					out.Series = append(out.Series, v2)
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
						out.Persons = make([]SearchPersonResponse, 0, 0)
					} else {
						out.Persons = []SearchPersonResponse{}
					}
				} else {
					out.Persons = (out.Persons)[:0]
				}
				for !in.IsDelim(']') {
					var v3 SearchPersonResponse
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
func easyjsonD4176298EncodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels(out *jwriter.Writer, in SearchResponse) {
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
	if len(in.Series) != 0 {
		const prefix string = ",\"serials\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v6, v7 := range in.Series {
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
func (v SearchResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD4176298EncodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SearchResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD4176298EncodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SearchResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD4176298DecodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SearchResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD4176298DecodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels(l, v)
}
func easyjsonD4176298DecodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels1(in *jlexer.Lexer, out *SearchPersonResponse) {
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
		case "original_name":
			out.OriginalName = string(in.String())
		case "birthday":
			out.Birthday = string(in.String())
		case "avatar":
			out.Avatar = string(in.String())
		case "count_films":
			out.CountFilms = int(in.Int())
		case "professions":
			if in.IsNull() {
				in.Skip()
				out.Professions = nil
			} else {
				in.Delim('[')
				if out.Professions == nil {
					if !in.IsDelim(']') {
						out.Professions = make([]string, 0, 4)
					} else {
						out.Professions = []string{}
					}
				} else {
					out.Professions = (out.Professions)[:0]
				}
				for !in.IsDelim(']') {
					var v10 string
					v10 = string(in.String())
					out.Professions = append(out.Professions, v10)
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
func easyjsonD4176298EncodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels1(out *jwriter.Writer, in SearchPersonResponse) {
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
	if in.OriginalName != "" {
		const prefix string = ",\"original_name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.OriginalName))
	}
	if in.Birthday != "" {
		const prefix string = ",\"birthday\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Birthday))
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
	if in.CountFilms != 0 {
		const prefix string = ",\"count_films\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.CountFilms))
	}
	if len(in.Professions) != 0 {
		const prefix string = ",\"professions\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v11, v12 := range in.Professions {
				if v11 > 0 {
					out.RawByte(',')
				}
				out.String(string(v12))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SearchPersonResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD4176298EncodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SearchPersonResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD4176298EncodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SearchPersonResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD4176298DecodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SearchPersonResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD4176298DecodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels1(l, v)
}
func easyjsonD4176298DecodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels2(in *jlexer.Lexer, out *SearchFilmResponse) {
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
		case "original_name":
			out.OriginalName = string(in.String())
		case "prod_date":
			out.ProdYear = string(in.String())
		case "poster_ver":
			out.PosterVer = string(in.String())
		case "rating":
			out.Rating = float32(in.Float32())
		case "genres":
			if in.IsNull() {
				in.Skip()
				out.Genres = nil
			} else {
				in.Delim('[')
				if out.Genres == nil {
					if !in.IsDelim(']') {
						out.Genres = make([]string, 0, 4)
					} else {
						out.Genres = []string{}
					}
				} else {
					out.Genres = (out.Genres)[:0]
				}
				for !in.IsDelim(']') {
					var v13 string
					v13 = string(in.String())
					out.Genres = append(out.Genres, v13)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "prod_countries":
			if in.IsNull() {
				in.Skip()
				out.ProdCountries = nil
			} else {
				in.Delim('[')
				if out.ProdCountries == nil {
					if !in.IsDelim(']') {
						out.ProdCountries = make([]string, 0, 4)
					} else {
						out.ProdCountries = []string{}
					}
				} else {
					out.ProdCountries = (out.ProdCountries)[:0]
				}
				for !in.IsDelim(']') {
					var v14 string
					v14 = string(in.String())
					out.ProdCountries = append(out.ProdCountries, v14)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "directors":
			if in.IsNull() {
				in.Skip()
				out.Directors = nil
			} else {
				in.Delim('[')
				if out.Directors == nil {
					if !in.IsDelim(']') {
						out.Directors = make([]SearchFilmPersonResponse, 0, 2)
					} else {
						out.Directors = []SearchFilmPersonResponse{}
					}
				} else {
					out.Directors = (out.Directors)[:0]
				}
				for !in.IsDelim(']') {
					var v15 SearchFilmPersonResponse
					(v15).UnmarshalEasyJSON(in)
					out.Directors = append(out.Directors, v15)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "end_year":
			out.EndYear = string(in.String())
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
func easyjsonD4176298EncodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels2(out *jwriter.Writer, in SearchFilmResponse) {
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
	if in.OriginalName != "" {
		const prefix string = ",\"original_name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.OriginalName))
	}
	if in.ProdYear != "" {
		const prefix string = ",\"prod_date\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ProdYear))
	}
	if in.PosterVer != "" {
		const prefix string = ",\"poster_ver\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.PosterVer))
	}
	if in.Rating != 0 {
		const prefix string = ",\"rating\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float32(float32(in.Rating))
	}
	if len(in.Genres) != 0 {
		const prefix string = ",\"genres\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v16, v17 := range in.Genres {
				if v16 > 0 {
					out.RawByte(',')
				}
				out.String(string(v17))
			}
			out.RawByte(']')
		}
	}
	if len(in.ProdCountries) != 0 {
		const prefix string = ",\"prod_countries\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v18, v19 := range in.ProdCountries {
				if v18 > 0 {
					out.RawByte(',')
				}
				out.String(string(v19))
			}
			out.RawByte(']')
		}
	}
	if len(in.Directors) != 0 {
		const prefix string = ",\"directors\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v20, v21 := range in.Directors {
				if v20 > 0 {
					out.RawByte(',')
				}
				(v21).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	if in.EndYear != "" {
		const prefix string = ",\"end_year\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.EndYear))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SearchFilmResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD4176298EncodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SearchFilmResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD4176298EncodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SearchFilmResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD4176298DecodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SearchFilmResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD4176298DecodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels2(l, v)
}
func easyjsonD4176298DecodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels3(in *jlexer.Lexer, out *SearchFilmPersonResponse) {
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
func easyjsonD4176298EncodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels3(out *jwriter.Writer, in SearchFilmPersonResponse) {
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
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SearchFilmPersonResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD4176298EncodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SearchFilmPersonResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD4176298EncodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SearchFilmPersonResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD4176298DecodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SearchFilmPersonResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD4176298DecodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels3(l, v)
}