package db

import (
	"fmt"
	"reflect"
	"regexp"
)

/*InitCfg update config to a string */
func InitCfg(cfg *string, cfgStruct interface{}) {
	t, v := reflect.TypeOf(cfgStruct), reflect.ValueOf(cfgStruct)
	for k := 0; k < t.NumField(); k++ {
		r, _ := regexp.Compile(t.Field(k).Name)
		*cfg = r.ReplaceAllString(*cfg, fmt.Sprintf(
			"%v", v.Field(k).Interface()))
	}
}
