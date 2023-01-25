package kace

import (
	"regexp"
	"strings"
	"text/template"
	"unicode"
)

type (
	Case uint
)

const (
	Camel Case = iota
	Pascal
	Snake
	Kedab
	Unknown
)

func DetectCatse(s string) Case {
	if IsPascal(s) {
		return Pascal
	}
	if IsCamel(s) {
		return Camel
	}
	if IsSnake(s) {
		return Snake
	}
	if Iskedab(s) {
		return Kedab
	}
	return Unknown
}

func IsPascal(s string) bool {
	return regexp.MustCompile(`^[A-Z]+[A-Za-z0-9]*$`).MatchString(s)
}

func IsCamel(s string) bool {
	return regexp.MustCompile(`^[a-z][A-Za-z0-9]+$`).MatchString(s)
}

func IsSnake(s string) bool {
	return regexp.MustCompile(`^([A-Za-z]+[_][A-Za-z]+)+$`).MatchString(s)
}

func Iskedab(s string) bool {
	return regexp.MustCompile(`^([A-Za-z]+[-][A-Za-z]+)+$`).MatchString(s)
}

func ToCamel(s string) string {
	return toPascalOrCamel(s, Camel)
}

func ToSnake(s string) string {
	return snake(s)
}

func ToPascal(s string) string {
	return toPascalOrCamel(s, Pascal)
}

func ToKedab(s string) string {
	return strings.Join(strings.Split(s, "_"), "-")
}

func toPascalOrCamel(s string, c Case) string {
	s = snake(s)
	rs := []rune(s)
	res := []rune{}

	for i, r := range rs {
		if s[i] == '_' {
			continue
		}

		if c == Pascal && i == 0 {
			res = append(res, unicode.ToUpper(r))
			continue
		}

		if i > 0 && s[i-1] == '_' {
			res = append(res, unicode.ToUpper(r))
			continue
		}

		res = append(res, r)
	}

	return string(res)
}

func TemplateFunc() template.FuncMap {
	m := template.FuncMap{
		"camel":  ToCamel,
		"pascal": ToPascal,
		"snake":  ToSnake,
		"kedab":  ToKedab,
	}
	return m
}

// Copyright entgo.io/ent
func snake(s string) string {
	var (
		j int
		b strings.Builder
	)
	for i := 0; i < len(s); i++ {
		r := rune(s[i])
		if i > 0 && i < len(s)-1 && unicode.IsUpper(r) {
			if unicode.IsLower(rune(s[i-1])) ||
				j != i-1 && unicode.IsLower(rune(s[i+1])) && unicode.IsLetter(rune(s[i-1])) {
				j = i
				b.WriteString("_")
			}
		}
		b.WriteRune(unicode.ToLower(r))
	}
	return b.String()
}
