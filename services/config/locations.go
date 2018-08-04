package config

import (
	"errors"
	"fmt"
	"reflect"
)

const locationsKey = "locations"

type Location struct {
	Name      string  `conf:"name"`
	InName    string  `conf:"in_name"`
	Anchor    string  `conf:"anchor"`
	Url       string  `conf:"url"`
	Latitude  float64 `conf:"latitude"`
	Longitude float64 `conf:"longitude"`
	SortOrder int     `conf:"sort_order"`
}

type locArr []Location

func (ls locArr) Len() int           { return len(ls) }
func (ls locArr) Swap(i, j int)      { ls[i], ls[j] = ls[j], ls[i] }
func (ls locArr) Less(i, j int) bool { return ls[i].SortOrder < ls[j].SortOrder }

func setField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()

	for i := 0; i < structValue.NumField(); i++ {
		fieldValue := structValue.Field(i)
		fieldType := structValue.Type().Field(i)
		tag := fieldType.Tag
		if tag.Get("conf") != name {
			continue
		}

		if !fieldValue.IsValid() {
			return fmt.Errorf("No such field: %s in obj", name)
		}

		if !fieldValue.CanSet() {
			return fmt.Errorf("Cannot set %s field value", name)
		}
		structFieldType := fieldValue.Type()
		val := reflect.ValueOf(value)
		if structFieldType != val.Type() {
			return errors.New("Provided value type didn't match obj field type")
		}

		fieldValue.Set(val)
	}
	return nil
}

func (l *Location) Fill(m map[string]interface{}) error {
	for k, v := range m {
		err := setField(l, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
