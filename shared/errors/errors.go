// カスタムエラーを提供します
package errors

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
)

var (
	NotFoundErr = fmt.Errorf("NorFoundErr")
)

// エラーをログ出力します
func LogErr(msg string, err error) {
	log.Println(fmt.Sprintf("%s: %v", msg, err))
}

// カスタムエラーです
type Error struct {
	Err      error
	Previous error
}

// エラーを新規作成します
//
// 一つ前のエラーを保持しているので、このエラーを再帰的に仕様することで簡易的なスタックトレースを出力できます。
//
// 引数のパターンは以下のとおりです。
//
// 1. エラーメッセージを指定する
//
// NewError("message")
//
// 2. エラーメッセージと一つ前のエラーを指定する
//
// NewError("message", err)
func NewError(msg string, arg ...interface{}) *Error {
	var (
		prev error
	)

	switch len(arg) {
	case 0:
		prev = nil
	case 1:
		switch a := arg[0].(type) {
		case *Error:
			prev = a
		case error:
			prev = a
		}
	}

	_, file, line, _ := runtime.Caller(1)

	if prev != nil {
		return &Error{
			Err:      fmt.Errorf(msg+" file: %s line: %d prev: [%w]", filepath.ToSlash(file), line, prev),
			Previous: prev,
		}
	}

	return &Error{
		Err:      fmt.Errorf(msg+" file: %s line: %d", filepath.ToSlash(file), line),
		Previous: prev,
	}
}

// エラーメッセージを返します
func (e *Error) Error() string {
	return e.Err.Error()
}

// １つ前のエラーを返します
func (e *Error) Unwrap() error {
	return e.Previous
}
