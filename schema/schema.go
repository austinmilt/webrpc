package schema

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
)

const (
	SCHEMA_VERSION = "v2"
)

// schema of webrpc json file, and validations
type WebRPCSchema struct {
	WebRPCVersion string    `json:"webrpc"`
	Name          string    `json:"name"`
	SchemaVersion string    `json:"version"`
	Imports       []*Import `json:"imports"`

	Types  []*Type  `json:"types"`
	Errors []*Error `json:"errors"`
	// Values   []*Value   `json:"values"` // TODO: future
	Services []*Service `json:"services"`
}

type Import struct {
	Path    string   `json:"path"`
	Members []string `json:"members"`
}

// Validate validates the schema through the AST, intended to be called after
// the json has been unmarshalled
func (s *WebRPCSchema) Validate() error {
	if s.WebRPCVersion != SCHEMA_VERSION {
		return fmt.Errorf("invalid webrpc schema version '%s', expecting '%s'", s.WebRPCVersion, SCHEMA_VERSION)
	}

	for _, typ := range s.Types {
		err := typ.Parse(s)
		if err != nil {
			return err
		}
	}
	for _, serr := range s.Errors {
		err := serr.Parse(s)
		if err != nil {
			return err
		}
	}
	// TODO: future feature
	// for _, sval := range s.Values {
	// 	err := sval.Parse(s)
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	for _, svc := range s.Services {
		err := svc.Parse(s)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *WebRPCSchema) SchemaHash() (string, error) {
	// TODO: lets later make this even more deterministic in face of re-ordering
	// definitions within the ridl file
	jsonString, err := s.ToJSON(false)
	if err != nil {
		return "", err
	}

	h := sha1.New()
	h.Write([]byte(jsonString))
	return hex.EncodeToString(h.Sum(nil)), nil
}

func (s *WebRPCSchema) ToJSON(optIndent ...bool) (string, error) {
	indent := false
	if len(optIndent) > 0 {
		indent = optIndent[0]
	}

	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	if indent {
		enc.SetIndent("", " ")
	}

	err := enc.Encode(s)
	if err != nil {
		return "", err
	}

	return string(buf.Bytes()), nil
}

func (s *WebRPCSchema) GetTypeByName(name string) *Type {
	name = strings.ToLower(name)
	for _, typ := range s.Types {
		if strings.ToLower(string(typ.Name)) == name {
			return typ
		}
	}
	return nil
}

func (s *WebRPCSchema) GetServiceByName(name string) *Service {
	name = strings.ToLower(name)
	for _, service := range s.Services {
		if strings.ToLower(string(service.Name)) == name {
			return service
		}
	}
	return nil
}

func (s *WebRPCSchema) HasFieldType(fieldType string) (bool, error) {
	fieldType = strings.ToLower(fieldType)
	_, ok := CoreTypeFromString[fieldType]
	if !ok {
		return false, fmt.Errorf("webrpc: invalid data type '%s'", fieldType)
	}

	for _, m := range s.Types {
		for _, f := range m.Fields {
			if DataTypeToString[f.Type.Type] == fieldType {
				return true, nil
			}
		}
	}

	for _, s := range s.Services {
		for _, m := range s.Methods {
			for _, i := range m.Inputs {
				if DataTypeToString[i.Type.Type] == fieldType {
					return true, nil
				}
			}
			for _, o := range m.Outputs {
				if DataTypeToString[o.Type.Type] == fieldType {
					return true, nil
				}
			}
		}
	}

	return false, nil
}
