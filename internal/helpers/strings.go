package helpers

import (
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func ApplyVar(s string, vars map[string]interface{}) string {
	r := regexp.MustCompile(`\%[A-Z]+\%`)
	res := r.FindAllString(s, -1)
	for _, v := range res {
		existVar := v
		if val, ok := vars[existVar]; ok {
			val = interfaceToStringInterface(val)
			s = strings.Replace(s, existVar, val.(string), -1)
		}
	}

	return s
}

func interfaceToStringInterface(val interface{}) interface{} {
	if reflect.TypeOf(val).String() == "string" {
		return val
	}

	if reflect.TypeOf(val).String() == "float64" {
		f := val.(float64)
		val = strconv.FormatFloat(f, 'f', -1, 64)
		return val
	}

	if reflect.TypeOf(val).String() == "int" {
		f := val.(int)
		val = strconv.Itoa(f)
		return val
	}

	return ""
}

func ApplyMap(m map[string]string, vars map[string]interface{}) map[string]string {
	nm := make(map[string]string)

	for k, v := range m {
		nm[k] = ApplyVar(v, vars)
	}

	return nm
}

func ApplySlice(m []string, vars map[string]interface{}) []string {
	for k, v := range m {
		m[k] = ApplyVar(v, vars)
	}

	return m
}
