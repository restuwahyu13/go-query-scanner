package queryscan

import (
	helper "github.com/restuwahyu13/go-query-scanner/helpers"
)

func Scan(query string, dest interface{}) error {
	valueof, store, err := helper.Parser(query, dest)
	if err != nil {
		return err
	}

	structType := valueof.Elem().Type()
	structValue := valueof.Elem()

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)

		tag := field.Tag.Get("query")
		if tag == "" {
			continue
		}

		if val, ok := store[tag]; ok {
			if structValue.Field(i).CanSet() {
				fieldValue := structValue.Field(i)

				if err := helper.Condition(field, fieldValue, val, tag); err != nil {
					return err
				}
			}
		}
	}

	store = nil
	return nil
}
