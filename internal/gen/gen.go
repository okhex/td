package gen

import (
	"fmt"
	"io"
	"strings"
	"text/template"

	"github.com/ernado/tl"
)

type Config struct {
	Package string
	Structs []Struct
}

type Struct struct {
	Name     string
	Comment  string
	Receiver string
	HexID    string
	BufArg   string

	Fields []Field
}

type Field struct {
	Name       string
	Comment    string
	Type       string
	PutFunc    string
	PutEncoder bool
}

func Generate(w io.Writer, t *template.Template, s *tl.Schema) error {
	cfg := Config{
		Package: "td",
	}

	// Searching for all types with single constructor.
	// This can be used to reduce interfaces.
	constructors := map[string]int{}
	for _, d := range s.Definitions {
		if d.Category != tl.CategoryType {
			continue
		}
		// TODO: Namespaces?
		constructors[d.Definition.Type.Name] += 1
	}
	singular := map[string]struct{}{}
	for k, v := range constructors {
		if v == 1 {
			singular[k] = struct{}{}
		}
	}

	for _, d := range s.Definitions {
		switch d.Category {
		case tl.CategoryType:
			s := Struct{
				Name:     pascal(d.Definition.Name),
				Receiver: strings.ToLower(d.Definition.Name[0:1]),
				HexID:    fmt.Sprintf("%x", d.Definition.ID),
				BufArg:   "b",
			}
			if s.Receiver == "b" {
				// bin.Buffer argument collides with reciever.
				s.BufArg = "buf"
			}
			for _, a := range d.Annotations {
				if a.Name == tl.AnnotationDescription {
					s.Comment = a.Value
				}
			}
			if s.Comment == "" {
				s.Comment = fmt.Sprintf("%s represents TL type %s#%x.",
					s.Name,
					d.Definition.Name,
					d.Definition.ID,
				)
			}
			for _, param := range d.Definition.Params {
				f := Field{
					Name: pascal(param.Name),
					Type: param.Type.Name,
				}
				for _, a := range d.Annotations {
					if a.Name == param.Name {
						f.Comment = a.Value
					}
				}
				switch param.Type.Name {
				case "int":
					f.PutFunc = "PutInt"
				case "int32":
					f.PutFunc = "PutInt32"
				case "string":
					f.PutFunc = "PutString"
				case "Bool":
					f.PutFunc = "PutBool"
					f.Type = "bool"
				default:
					f.PutEncoder = true
				}
				if f.Comment == "" {
					f.Comment = fmt.Sprintf("%s field of %s.", f.Name, s.Name)
				}
				s.Fields = append(s.Fields, f)
			}
			cfg.Structs = append(cfg.Structs, s)
		}
	}
	return t.ExecuteTemplate(w, "simple", cfg)
}
