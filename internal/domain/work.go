package domain

import (
	"context"
	"errors"
)

// Work は作業の進捗を表すドメインモデルです
type Work struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Progress int    `json:"progress"`
}

// WorkRepository はWorkの永続化を担当するリポジトリのインターフェースです
type WorkRepository interface {
	FindAll(ctx context.Context) ([]*Work, error)
	FindByID(ctx context.Context, id int) (*Work, error)
	Create(ctx context.Context, work *Work) error
	Update(ctx context.Context, work *Work) error
	Delete(ctx context.Context, id int) error
}

// Validate は作業の妥当性を検証します
func (w *Work) Validate() error {
	if w.Title == "" {
		return errors.New("title is required")
	}
	if w.Progress < 0 || w.Progress > 100 {
		return errors.New("progress must be between 0 and 100")
	}
	return nil
}
