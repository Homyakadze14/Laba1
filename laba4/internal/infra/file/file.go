package file

import (
	"encoding/json"
	"laba4/internal/entity"
	"os"
)

type FileStore struct{}

func NewFileStore() *FileStore {
	return &FileStore{}
}

func (fs *FileStore) Save(students []entity.Stud, filename string) error {
	if len(students) == 0 {
		return writeJSON(filename, []entity.Stud{})
	}
	return writeJSON(filename, students)
}

func (fs *FileStore) Load(filename string) ([]entity.Stud, error) {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []entity.Stud{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var students []entity.Stud
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&students); err != nil {
		return nil, err
	}
	return students, nil
}

func writeJSON(filename string, data interface{}) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}
