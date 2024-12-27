package memory

import (
	"context"
	"sync"

	"github.com/example/layeredArchitectureWithGo/internal/domain"
)

// WorkRepository はインメモリでWorkを管理するリポジトリの実装です
type WorkRepository struct {
	sync.RWMutex
	works map[int]*domain.Work
	maxID int
}

// NewWorkRepository は新しいWorkRepositoryを作成します
func NewWorkRepository() *WorkRepository {
	return &WorkRepository{
		works: make(map[int]*domain.Work),
		maxID: 0,
	}
}

// FindAll は全ての作業を取得します
func (r *WorkRepository) FindAll(ctx context.Context) ([]*domain.Work, error) {
	r.RLock()
	defer r.RUnlock()

	works := make([]*domain.Work, 0, len(r.works))
	for _, work := range r.works {
		works = append(works, work)
	}
	return works, nil
}

// FindByID は指定されたIDの作業を取得します
func (r *WorkRepository) FindByID(ctx context.Context, id int) (*domain.Work, error) {
	r.RLock()
	defer r.RUnlock()

	work, ok := r.works[id]
	if !ok {
		return nil, domain.ErrWorkNotFound
	}
	return work, nil
}

// Create は新しい作業を作成します
func (r *WorkRepository) Create(ctx context.Context, work *domain.Work) error {
	r.Lock()
	defer r.Unlock()

	r.maxID++
	work.ID = r.maxID
	r.works[work.ID] = work
	return nil
}

// Update は既存の作業を更新します
func (r *WorkRepository) Update(ctx context.Context, work *domain.Work) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.works[work.ID]; !ok {
		return domain.ErrWorkNotFound
	}
	r.works[work.ID] = work
	return nil
}

// Delete は指定されたIDの作業を削除します
func (r *WorkRepository) Delete(ctx context.Context, id int) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.works[id]; !ok {
		return domain.ErrWorkNotFound
	}
	delete(r.works, id)
	return nil
}
