package go_benchmark

import (
	"github.com/buger/jsonparser"
	"testing"
	"github.com/json-iterator/go"
	"encoding/json"
	"github.com/mailru/easyjson/jlexer"
)

func BenchmarkJsonParserEachKeyStructMedium(b *testing.B) {
	paths := [][]string{
		[]string{"person", "name", "fullName"},
		[]string{"person", "github", "followers"},
		[]string{"company"},
		[]string{"person", "gravatar", "avatars"},
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data := MediumPayload{
			Person: &CBPerson{
				Name:     &CBName{},
				Github:   &CBGithub{},
				Gravatar: &CBGravatar{},
			},
		}

		jsonparser.EachKey(mediumFixture, func(idx int, value []byte, vt jsonparser.ValueType, err error) {
			switch idx {
			case 0:
				data.Person.Name.FullName, _ = jsonparser.ParseString(value)
			case 1:
				v, _ := jsonparser.ParseInt(value)
				data.Person.Github.Followers = int(v)
			case 2:
				v, _ := jsonparser.ParseString(value)
				data.Company = v
			case 3:
				var avatars []*CBAvatar
				jsonparser.ArrayEach(value, func(avalue []byte, dataType jsonparser.ValueType, offset int, err error) {
					url, _ := jsonparser.GetString(avalue, "url")
					avatars = append(avatars, &CBAvatar{Url: url})
				})
				data.Person.Gravatar.Avatars = avatars
			}
		}, paths...)
	}
}

func BenchmarkJsoniterStructMedium(b *testing.B) {
	iter := jsoniter.ParseBytes(mediumFixture)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data := MediumPayload{
			Person: &CBPerson{
				Name:     &CBName{},
				Github:   &CBGithub{},
				Gravatar: &CBGravatar{},
			},
		}
		iter.ResetBytes(mediumFixture)
		for field := iter.ReadObject(); field != ""; field = iter.ReadObject() {
			switch field {
			case "company":
				data.Company = iter.ReadString()
			case "person":
				for field := iter.ReadObject(); field != ""; field = iter.ReadObject() {
					switch field {
					case "name":
						for field := iter.ReadObject(); field != ""; field = iter.ReadObject() {
							if "fullName" != field {
								iter.Skip()
								continue
							}
							data.Person.Name.FullName = iter.ReadString()
						}
					case "github":
						for field := iter.ReadObject(); field != ""; field = iter.ReadObject() {
							if "followers" != field {
								iter.Skip()
								continue
							}
							data.Person.Github.Followers = iter.ReadInt()
						}
					case "gravatar":
						for field := iter.ReadObject(); field != ""; field = iter.ReadObject() {
							if "avatars" != field {
								iter.Skip()
								continue
							}
							var avatars []*CBAvatar
							for iter.ReadArray() {
								for field := iter.ReadObject(); field != ""; field = iter.ReadObject() {
									if "url" != field {
										iter.Skip()
										continue
									}
									avatars = append(avatars, &CBAvatar{Url: iter.ReadString()})
								}
							}
							data.Person.Gravatar.Avatars = avatars
						}
					default:
						iter.Skip()
					}
				}
			default:
				iter.Skip()
			}
		}
	}
}


func BenchmarkJsoniterReflectStructMedium(b *testing.B) {
	iter := jsoniter.ParseBytes(mediumFixture)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data := MediumPayload{
			Person: &CBPerson{
				Name:     &CBName{},
				Github:   &CBGithub{},
				Gravatar: &CBGravatar{},
			},
		}
		iter.ResetBytes(mediumFixture)
		iter.ReadVal(&data)
	}
}

/*
   encoding/json
*/
func BenchmarkEncodingJsonStructMedium(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var data MediumPayload
		json.Unmarshal(mediumFixture, &data)
	}
}

func BenchmarkEasyJsonMedium(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		lexer := &jlexer.Lexer{Data: mediumFixture}
		data := &MediumPayload{
			Person: &CBPerson{
				Name:     &CBName{},
				Github:   &CBGithub{},
				Gravatar: &CBGravatar{},
			},
		}
		data.UnmarshalEasyJSON(lexer)
	}
}