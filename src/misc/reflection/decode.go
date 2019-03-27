package reflection

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"text/scanner"
)

type lexer struct {
	scan  scanner.Scanner
	token rune
}

func main() {
	var m = make(map[string]string)
	var data = make([]byte, 100)
	Unmarshal(data, m)
}

// Unmarshal ...
func Unmarshal(data []byte, out interface{}) (err error) {
	lex := &lexer{scan: scanner.Scanner{}}
	lex.scan.Init(bytes.NewReader(data))
	lex.next()
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("Error at %s: %v", lex.scan.Position, x)
		}
	}()
	read(lex, reflect.ValueOf(out).Elem())
	return nil
}

func read(lex *lexer, v reflect.Value) {
	switch lex.token {
	case scanner.Ident:
		if lex.text() == "nil" {
			v.Set(reflect.Zero(v.Type()))
			lex.next()
			return
		}
	case scanner.Int:
		i, _ := strconv.Atoi(lex.text())
		v.SetInt(int64(i))
		return
	case scanner.String:
		s, _ := strconv.Unquote(lex.text())
		v.SetString(s)
		return
	case '(':
		lex.next()
		readList(lex, v)
		lex.consume(')')
		return
	}
	panic(fmt.Sprintf("unexpected token %q\n", lex.text()))
}

func readList(lex *lexer, v reflect.Value) {
	switch v.Kind() {
	case reflect.Array:
		for i := 0; i < v.Len(); i++ {
			read(lex, v.Index(i))
		}
	case reflect.Slice:
		for !endList(lex) {
			item := reflect.New(v.Type().Elem()).Elem()
			read(lex, item)
			v.Set(reflect.Append(v, item))
		}
	case reflect.Map:
		v.Set(reflect.MakeMap(v.Type()))
		for !endList(lex) {
			lex.consume('(')
			key := reflect.New(v.Type().Key()).Elem()
			read(lex, key)
			value := reflect.New(v.Type().Elem()).Elem()
			read(lex, value)
			v.SetMapIndex(key, value)
			lex.consume(')')
		}
	case reflect.Struct:

		for !endList(lex) {
			lex.consume('(')
			if lex.token != scanner.Ident {
				panic(fmt.Sprintf("want field name, got token %q", lex.text()))
			}
			name := lex.text()
			lex.next()
			read(lex, v.FieldByName(name))
			lex.consume(')')
		}
	default:
		panic(fmt.Sprintf("Cannot decode list into %v", v.Type()))
	}
}

func (lex *lexer) next() {
	lex.token = lex.scan.Scan()
}

func (lex *lexer) text() string {
	return lex.scan.TokenText()
}

func (lex *lexer) consume(want rune) {
	if lex.token != want {
		panic(fmt.Sprintf("want %q, got %q", want, lex.token))
	}
	lex.next()
}

func endList(lex *lexer) bool {
	switch lex.token {
	case scanner.EOF:
		panic("end of file")
	case ')':
		return true
	}
	return false
}
