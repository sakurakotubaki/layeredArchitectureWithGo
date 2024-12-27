package domain

import "errors"

var (
	// ErrWorkNotFound は作業が見つからない場合のエラーです
	ErrWorkNotFound = errors.New("work not found")
)
