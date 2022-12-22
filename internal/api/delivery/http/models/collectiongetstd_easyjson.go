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

func easyjsonE94b7a4fDecodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels(in *jlexer.Lexer, out *GetStdCollectionResponse) {
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
						out.Films = make([]FilmTagCollectionResponse, 0, 0)
					} else {
						out.Films = []FilmTagCollectionResponse{}
					}
				} else {
					out.Films = (out.Films)[:0]
				}
				for !in.IsDelim(']') {
					var v1 FilmTagCollectionResponse
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
func easyjsonE94b7a4fEncodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels(out *jwriter.Writer, in GetStdCollectionResponse) {
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
func (v GetStdCollectionResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE94b7a4fEncodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetStdCollectionResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE94b7a4fEncodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetStdCollectionResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE94b7a4fDecodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetStdCollectionResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE94b7a4fDecodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels(l, v)
}
func easyjsonE94b7a4fDecodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels1(in *jlexer.Lexer, out *FilmTagCollectionResponse) {
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
		case "prod_year":
			out.ProdYear = string(in.String())
		case "end_year":
			out.EndYear = string(in.String())
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
					var v4 string
					v4 = string(in.String())
					out.Genres = append(out.Genres, v4)
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
func easyjsonE94b7a4fEncodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels1(out *jwriter.Writer, in FilmTagCollectionResponse) {
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
	if in.ProdYear != "" {
		const prefix string = ",\"prod_year\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ProdYear))
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
			for v5, v6 := range in.Genres {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FilmTagCollectionResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE94b7a4fEncodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FilmTagCollectionResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE94b7a4fEncodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FilmTagCollectionResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE94b7a4fDecodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FilmTagCollectionResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE94b7a4fDecodeGoParkMailRu20222BugOverloadInternalApiDeliveryHttpModels1(l, v)
}