package jsonschema

import (
	"encoding/json"
	"fmt"

	"cuelang.org/go/cue"
	"github.com/google/jsonschema-go/jsonschema"
)

func Validator(schemav cue.Value) (cue.Value, error) {
	var s jsonschema.Schema
	if err := schemav.Decode(&s); err != nil {
		return cue.Value{}, fmt.Errorf("cannot decode JSON Schema: %v", err)
	}
	resolved, err := s.Resolve(nil)
	if err != nil {
		return cue.Value{}, fmt.Errorf("cannot resolve JSON Schema: %v", err)
	}
	return cue.NewPureValidatorFunc(func(inst cue.Value) error {
		data, err := inst.MarshalJSON()
		if err != nil {
			return err
		}
		var x any
		if err := json.Unmarshal(data, &x); err != nil {
			return err
		}
		return resolved.Validate(x)
	}), nil
}
