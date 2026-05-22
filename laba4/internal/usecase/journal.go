package usecase

import (
	"errors"
	"laba4/internal/entity"
)

type StudentRepository interface {
	Add(name string)
	Delete(name string)
	UpdateMark(name string, mark int)
	GetMark(name string) (int, error)
	GetAll() []entity.Stud
	GetAvg() float64
	GetMax() *entity.Stud
	GetMin() *entity.Stud
	GetAllMoreThen(needMark int) []entity.Stud
	CountBelow(threshold int) int
	Found(name string) bool
	GetAllSortedByName() []entity.Stud
}

type FileStorage interface {
	Save(students []entity.Stud, filename string) error
	Load(filename string) ([]entity.Stud, error)
}

type JournalUsecase struct {
	repo StudentRepository
	file FileStorage
}

func NewJournalUsecase(repo StudentRepository, file FileStorage) *JournalUsecase {
	return &JournalUsecase{
		repo: repo,
		file: file,
	}
}

func (uc *JournalUsecase) AddStudent(name string) error {
	if name == "" {
		return errors.New("student name cannot be empty")
	}
	uc.repo.Add(name)
	return nil
}

func (uc *JournalUsecase) DeleteStudent(name string) error {
	if !uc.repo.Found(name) {
		return errors.New("student not found")
	}
	uc.repo.Delete(name)
	return nil
}

func (uc *JournalUsecase) UpdateMark(name string, mark int) error {
	if !uc.repo.Found(name) {
		return errors.New("student not found")
	}
	if mark < 0 || mark > 100 {
		return errors.New("mark must be between 0 and 100")
	}
	uc.repo.UpdateMark(name, mark)
	return nil
}

func (uc *JournalUsecase) GetMark(name string) (int, error) {
	return uc.repo.GetMark(name)
}

func (uc *JournalUsecase) GetAll() []entity.Stud {
	return uc.repo.GetAll()
}

func (uc *JournalUsecase) GetAvg() float64 {
	return uc.repo.GetAvg()
}

func (uc *JournalUsecase) GetMax() *entity.Stud {
	return uc.repo.GetMax()
}

func (uc *JournalUsecase) GetMin() *entity.Stud {
	return uc.repo.GetMin()
}

func (uc *JournalUsecase) GetAllMoreThen(threshold int) []entity.Stud {
	return uc.repo.GetAllMoreThen(threshold)
}

func (uc *JournalUsecase) CountBelow(threshold int) int {
	return uc.repo.CountBelow(threshold)
}

func (uc *JournalUsecase) FoundStudent(name string) bool {
	return uc.repo.Found(name)
}

func (uc *JournalUsecase) GetAllSortedByName() []entity.Stud {
	return uc.repo.GetAllSortedByName()
}

func (uc *JournalUsecase) SaveToFile(filename string) error {
	students := uc.repo.GetAll()
	return uc.file.Save(students, filename)
}

func (uc *JournalUsecase) LoadFromFile(filename string) error {
	students, err := uc.file.Load(filename)
	if err != nil {
		return err
	}
	all := uc.repo.GetAll()
	for _, s := range all {
		uc.repo.Delete(s.Name)
	}
	for _, s := range students {
		uc.repo.Add(s.Name)
		uc.repo.UpdateMark(s.Name, s.Mark)
	}
	return nil
}
