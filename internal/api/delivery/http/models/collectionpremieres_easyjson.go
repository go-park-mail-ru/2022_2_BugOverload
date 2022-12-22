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

func easyjson4f9fdeaDecodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels(in *jlexer.Lexer, out *PremieresCollectionResponse) {
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
		case "description":
			out.Description = string(in.String())
		case "films":
			if in.IsNull() {
				in.Skip()
				out.Films = nil
			} else {
				in.Delim('[')
				if out.Films == nil {
					if !in.IsDelim(']') {
						out.Films = make([]PremieresCollectionFilm, 0, 0)
					} else {
						out.Films = []PremieresCollectionFilm{}
					}
				} else {
					out.Films = (out.Films)[:0]
				}
				for !in.IsDelim(']') {
					var v1 PremieresCollectionFilm
					(v1).UnmarshalEasyJSON(in)
					out.Films = append(out.Films, v1)
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
func easyjson4f9fdeaEncodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels(out *jwriter.Writer, in PremieresCollectionResponse) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Name != "" {
		const prefix string = ",\"name\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	if in.Description != "" {
		const prefix string = ",\"description\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Description))
	}
	if len(in.Films) != 0 {
		const prefix string = ",\"films\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v2, v3 := range in.Films {
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
func (v PremieresCollectionResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4f9fdeaEncodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PremieresCollectionResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4f9fdeaEncodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PremieresCollectionResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4f9fdeaDecodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PremieresCollectionResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4f9fdeaDecodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels(l, v)
}
func easyjson4f9fdeaDecodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels1(in *jlexer.Lexer, out *PremieresCollectionFilm) {
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
		case "prod_date":
			out.ProdDate = string(in.String())
		case "poster_ver":
			out.PosterVer = string(in.String())
		case "rating":
			out.Rating = float32(in.Float32())
		case "duration_minutes":
			out.DurationMinutes = int(in.Int())
		case "description":
			out.Description = string(in.String())
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
					var v4 string
					v4 = string(in.String())
					out.Genres = append(out.Genres, v4)
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
					var v5 string
					v5 = string(in.String())
					out.ProdCountries = append(out.ProdCountries, v5)
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
						out.Directors = make([]FilmPersonPremiersResponse, 0, 2)
					} else {
						out.Directors = []FilmPersonPremiersResponse{}
					}
				} else {
					out.Directors = (out.Directors)[:0]
				}
				for !in.IsDelim(']') {
					var v6 FilmPersonPremiersResponse
					(v6).UnmarshalEasyJSON(in)
					out.Directors = append(out.Directors, v6)
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
func easyjson4f9fdeaEncodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels1(out *jwriter.Writer, in PremieresCollectionFilm) {
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
	if in.ProdDate != "" {
		const prefix string = ",\"prod_date\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ProdDate))
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
	if in.DurationMinutes != 0 {
		const prefix string = ",\"duration_minutes\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.DurationMinutes))
	}
	if in.Description != "" {
		const prefix string = ",\"description\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Description))
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
			for v7, v8 := range in.Genres {
				if v7 > 0 {
					out.RawByte(',')
				}
				out.String(string(v8))
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
			for v9, v10 := range in.ProdCountries {
				if v9 > 0 {
					out.RawByte(',')
				}
				out.String(string(v10))
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
			for v11, v12 := range in.Directors {
				if v11 > 0 {
					out.RawByte(',')
				}
				(v12).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PremieresCollectionFilm) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4f9fdeaEncodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PremieresCollectionFilm) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4f9fdeaEncodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PremieresCollectionFilm) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4f9fdeaDecodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PremieresCollectionFilm) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4f9fdeaDecodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels1(l, v)
}
func easyjson4f9fdeaDecodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels2(in *jlexer.Lexer, out *FilmPersonPremiersResponse) {
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
func easyjson4f9fdeaEncodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels2(out *jwriter.Writer, in FilmPersonPremiersResponse) {
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
func (v FilmPersonPremiersResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4f9fdeaEncodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FilmPersonPremiersResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4f9fdeaEncodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FilmPersonPremiersResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4f9fdeaDecodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FilmPersonPremiersResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4f9fdeaDecodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels2(l, v)
}
