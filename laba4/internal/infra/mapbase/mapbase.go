package mapbase

import (
	"errors"
	"fmt"
	"laba4/internal/entity"
	"slices"
	"strings"
	"sync"
)

const dcp = 10

type DB struct {
	mp sync.Map
}

func NewDB() *DB {
	return &DB{}
}

func (d *DB) Add(stud string) {
	d.mp.Store(stud, 0)
}

func (d *DB) Delete(stud string) {
	d.mp.Delete(stud)
}

func (d *DB) UpdateMark(stud string, mark int) {
	d.mp.Store(stud, mark)
}

var ErrNotFound = errors.New("mark not found")

func (d *DB) GetMark(stud string) (int, error) {
	v, ok := d.mp.Load(stud)
	if !ok {
		return 0, fmt.Errorf("%w: %s", ErrNotFound, stud)
	}

	mark, ok := v.(int)
	if !ok {
		return 0, fmt.Errorf("invalid mark type for %s: %T", stud, v)
	}
	return mark, nil
}

func (d *DB) GetAll() []entity.Stud {
	studs := make([]entity.Stud, 0, dcp)
	d.mp.Range(func(key, value any) bool {
		mark, ok := value.(int)
		if !ok {
			return true
		}

		stud := entity.Stud{
			Name: key.(string),
			Mark: mark,
		}
		studs = append(studs, stud)
		return true
	})

	return studs
}

func (d *DB) GetAvg() float64 {
	sum := 0.
	len := 0.
	d.mp.Range(func(key, value any) bool {
		mark, ok := value.(int)
		if !ok {
			return true
		}
		sum += float64(mark)
		len++
		return true
	})

	if len == 0 {
		return 0
	}

	avg := sum / len
	return avg
}

func (d *DB) GetMax() *entity.Stud {
	var stud *entity.Stud
	maxMark := 0
	d.mp.Range(func(key, value any) bool {
		mark, ok := value.(int)
		if !ok {
			return true
		}

		if mark > maxMark {
			maxMark = mark
			stud = &entity.Stud{
				Name: key.(string),
				Mark: mark,
			}
		}

		return true
	})

	return stud
}

func (d *DB) GetMin() *entity.Stud {
	var stud *entity.Stud
	minMark := 10000
	d.mp.Range(func(key, value any) bool {
		mark, ok := value.(int)
		if !ok {
			return true
		}

		if mark < minMark {
			minMark = mark
			stud = &entity.Stud{
				Name: key.(string),
				Mark: mark,
			}
		}

		return true
	})

	return stud
}

func (d *DB) GetAllMoreThen(needMark int) []entity.Stud {
	studs := make([]entity.Stud, 0, dcp)
	d.mp.Range(func(key, value any) bool {
		mark, ok := value.(int)
		if !ok {
			return true
		}

		if mark < needMark {
			return true
		}

		stud := entity.Stud{
			Name: key.(string),
			Mark: mark,
		}
		studs = append(studs, stud)
		return true
	})

	return studs
}

func (d *DB) CountBelow(threshold int) int {
	count := 0
	d.mp.Range(func(key, value any) bool {
		mark, ok := value.(int)
		if !ok {
			return true
		}
		if mark < threshold {
			count++
		}
		return true
	})
	return count
}

func (d *DB) Found(stud string) bool {
	_, ok := d.mp.Load(stud)
	return ok
}

func (d *DB) GetAllSortedByName() []entity.Stud {
	studs := make([]entity.Stud, 0, dcp)
	d.mp.Range(func(key, value any) bool {
		mark, ok := value.(int)
		if !ok {
			return true
		}

		stud := entity.Stud{
			Name: key.(string),
			Mark: mark,
		}
		studs = append(studs, stud)
		return true
	})

	slices.SortFunc(studs, func(f entity.Stud, s entity.Stud) int {
		return strings.Compare(f.Name, s.Name)
	})

	return studs
}
