package api

import (
	"fmt"
	"net/http"
	"reflect"
)

type RequestBody struct {
	Path string
	Name string
	Type []TypeReflector
}

type TypeReflector struct {
	field string
	name  string
}

type ResponseBody struct {
	Path string
	Name string
	Type []TypeReflector
}

func Get(router *http.ServeMux, path string, handler func(w http.ResponseWriter, r *http.Request), type_ interface{}) RequestBody {
	router.HandleFunc("GET "+path, handler)
	if type_ == nil {
		return RequestBody{Path: path, Type: nil}
	}
	var typeReflector []TypeReflector
	userType := reflect.TypeOf(type_)
	if userType.Kind() == reflect.Slice {
		userType = userType.Elem()
	}
	for i := 0; i < userType.NumField(); i++ {
		field := userType.Field(i)
		fmt.Println(field)
		fmt.Println(field.Type)
		fmt.Println(field.Name)
		fmt.Println(field.Type.Kind())
		fmt.Println(field.Type.String())
		fmt.Println(field.Tag.Get("json"))
		typeReflector = append(typeReflector, TypeReflector{field: field.Type.String(), name: field.Tag.Get("json")})
	}
	return RequestBody{Path: path, Type: typeReflector, Name: userType.String()}
}

func Post(router *http.ServeMux, path string, handler func(w http.ResponseWriter, r *http.Request)) {
	router.HandleFunc("POST "+path, handler)
}

func Put(router *http.ServeMux, path string, handler func(w http.ResponseWriter, r *http.Request)) {
	router.HandleFunc("PUT "+path, handler)
}

func Delete(router *http.ServeMux, path string, handler func(w http.ResponseWriter, r *http.Request)) {
	router.HandleFunc("DELETE "+path, handler)
}

func Patch(router *http.ServeMux, path string, handler func(w http.ResponseWriter, r *http.Request)) {
	router.HandleFunc("PATCH "+path, handler)
}

func Options(router *http.ServeMux, path string, handler func(w http.ResponseWriter, r *http.Request)) {
	router.HandleFunc("OPTIONS "+path, handler)
}
