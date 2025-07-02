package task

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Manager struct {
	tasks sync.Map
}

type taskEntry struct {
	task   *Task
	cancel context.CancelFunc
}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) CreateTask(_ context.Context) *Task {
	id := uuid.NewString()

	t := &Task{
		ID:        id,
		Status:    StatusPending,
		CreatedAt: time.Now(),
	}

	ctx, cancel := context.WithCancel(context.Background())

	m.tasks.Store(id, &taskEntry{
		task:   t,
		cancel: cancel,
	})

	go m.runTask(ctx, id)

	return t
}

func (m *Manager) runTask(ctx context.Context, id string) {
	entry, ok := m.loadTask(id)
	if !ok {
		return
	}

	task := entry.task
	task.Status = StatusRunning
	start := time.Now()

	select {
	case <-time.After(3 * time.Minute):
		task.Status = StatusCompleted
		task.Result = fmt.Sprintf("Task %s completed successfully", id)

	case <-ctx.Done():
		task.Status = StatusCanceled
		task.ErrorMessage = "Task was canceled"
	}

	task.Duration = time.Since(start)
}

func (m *Manager) GetTask(id string) (*Task, error) {
	entry, ok := m.loadTask(id)
	if !ok {
		return nil, fmt.Errorf("task with id %s not found", id)
	}

	return entry.task, nil
}

func (m *Manager) DeleteTask(id string) error {
	entry, ok := m.loadTask(id)
	if !ok {
		return fmt.Errorf("task with id %s not found", id)
	}

	entry.cancel()
	m.tasks.Delete(id)
	return nil
}

func (m *Manager) loadTask(id string) (*taskEntry, bool) {
	raw, ok := m.tasks.Load(id)
	if !ok {
		return nil, false
	}

	return raw.(*taskEntry), true
}
