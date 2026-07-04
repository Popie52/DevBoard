package service

import (
	"errors"
	"sync"
	"time"

	"github.com/Popie52/devboard/backend/internal/model"
)

type NoteService struct {
	mu     sync.Mutex
	notes  []model.Note
	nextID int
}

func NewNoteService() *NoteService {
	return &NoteService{
		notes:  make([]model.Note, 0),
		nextID: 1,
	}
}

func (s *NoteService) GetAll() []model.Note {
	s.mu.Lock()
	defer s.mu.Unlock()

	return append([]model.Note{}, s.notes...)
}

func (s *NoteService) Create(title, content string) model.Note {
	s.mu.Lock()
	defer s.mu.Unlock()

	note := model.Note{
		ID:        s.nextID,
		Title:     title,
		Content:   content,
		CreatedAt: time.Now().UTC(),
	}

	s.nextID++
	s.notes = append(s.notes, note)

	return note
}

func (s *NoteService) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, n := range s.notes {
		if n.ID == id {
			s.notes = append(s.notes[:i], s.notes[i+1:]...)
			return nil
		}
	}

	return errors.New("note not found")
}