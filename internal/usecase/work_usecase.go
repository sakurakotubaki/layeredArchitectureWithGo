package usecase

import (
	"context"

	"github.com/example/layeredArchitectureWithGo/internal/domain"
)

// WorkUseCase は作業に関するユースケースを実装します
type WorkUseCase struct {
	repo domain.WorkRepository
}

// NewWorkUseCase は新しいWorkUseCaseを作成します
func NewWorkUseCase(repo domain.WorkRepository) *WorkUseCase {
	return &WorkUseCase{
		repo: repo,
	}
}

// GetAll は全ての作業を取得します
func (u *WorkUseCase) GetAll(ctx context.Context) ([]*domain.Work, error) {
	return u.repo.FindAll(ctx)
}

// GetByID は指定されたIDの作業を取得します
func (u *WorkUseCase) GetByID(ctx context.Context, id int) (*domain.Work, error) {
	return u.repo.FindByID(ctx, id)
}

// Create は新しい作業を作成します
func (u *WorkUseCase) Create(ctx context.Context, work *domain.Work) error {
	if err := work.Validate(); err != nil {
		return err
	}
	return u.repo.Create(ctx, work)
}

// Update は既存の作業を更新します
func (u *WorkUseCase) Update(ctx context.Context, work *domain.Work) error {
	if err := work.Validate(); err != nil {
		return err
	}
	return u.repo.Update(ctx, work)
}

// Delete は指定されたIDの作業を削除します
func (u *WorkUseCase) Delete(ctx context.Context, id int) error {
	return u.repo.Delete(ctx, id)
}
