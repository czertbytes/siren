package siren

import (
	"fmt"
	"reflect"
	"strings"
)

type SirenMarshaller interface {
	MarshalSirenJSON() ([]byte, error)
}

func Marshal(v SirenMarshaller) ([]byte, error) {
	return v.MarshalSirenJSON()
}

func ParseProperties(v interface{}) (Properties, error) {
	properties := make(Properties)

	sv := reflect.ValueOf(v)
	st := reflect.TypeOf(v)

	for i := 0; i < sv.NumField(); i++ {
		f := sv.Field(i)
		sirenTagValue := st.Field(i).Tag.Get("siren")

		if len(sirenTagValue) > 0 {
			fieldJSONKey := st.Field(i).Name

			jsonTagValue := st.Field(i).Tag.Get("json")
			if len(jsonTagValue) > 0 {
				if jsonTagValue == "-" {
					return properties, fmt.Errorf("'siren': Field '%s' is ignored by JSON Tag!", fieldJSONKey)
				}

				if strings.Contains(jsonTagValue, ",") {
					jsonTagValues := strings.Split(jsonTagValue, ",")
					if len(jsonTagValues) > 0 && len(jsonTagValues[0]) > 0 {
						fieldJSONKey = jsonTagValue
					}
				} else {
					fieldJSONKey = jsonTagValue
				}
			}

			if sirenTagValue == "property" {
				properties[fieldJSONKey] = f.Interface()
			}
		}
	}

	return properties, nil
}
