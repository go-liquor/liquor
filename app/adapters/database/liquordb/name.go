package liquordb

import (
	"fmt"
	"github.com/go-liquor/liquor/v2/pkg/lqstring"
	"reflect"
)

func GetCollectionName(collection any) (string, error) {
	t := reflect.TypeOf(collection)
	if t.Kind() == reflect.Ptr {
		if t.Elem().Name() != "" {
			return lqstring.ToSnakeCase(lqstring.ToPlural(t.Elem().Name())), nil
		}
		// is a slice
		if t.Elem().Elem().Name() != "" {
			return lqstring.ToSnakeCase(lqstring.ToPlural(t.Elem().Elem().Name())), nil
		}

	}
	if t.Kind() == reflect.Slice {
		return lqstring.ToSnakeCase(lqstring.ToPlural(t.Elem().Name())), nil
	}
	return "", fmt.Errorf("collection must be a pointer or a slice! Is: %v", t.Kind())
}
