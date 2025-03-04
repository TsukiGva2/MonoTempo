package dbman

import (
	"fmt"
	"os"

	"database/sql"
	"envio/athlete"
	_ "modernc.org/sqlite"
)

type Baselet struct {
	db     *sql.DB
	Path   string
	opened bool
}

type MADB struct {
	databases []Baselet
	groupSize int // group length in which athletes are divided (0-10, 0-100, ...)
	maxValue  int // max number that fits
}

func NewBaselet(path string) (b Baselet, err error) {

	_, err = os.Create(path)

	if err != nil {

		return
	}

	db, err := sql.Open("sqlite", path)

	if err != nil {

		return
	}

	defer db.Close() // flush changes

	_, err = db.Exec(CREATE_TIME_DATABASE)

	if err != nil {

		return
	}

	b.db = db
	b.Path = path

	return
}

func (b *Baselet) Open() (err error) {

	if b.opened {

		return
	}

	db, err := sql.Open("sqlite", b.Path)

	if err != nil {

		return
	}

	b.db = db

	b.opened = true

	return
}

func (b *Baselet) Insert(a *athlete.Atleta) (err error) {

	err = b.Open()

	if err != nil {

		return
	}

	_, err = b.db.Exec(
		INSERT_TIME,

		a.Antena,
		a.Numero,
		a.Staff,
		a.Tempo,
	)

	return
}

func (b *Baselet) Close() {

	if !b.opened {

		return
	}

	b.db.Close()

	b.opened = false
}

func (m *MADB) GroupSize(s int) {
	m.groupSize = s
}

func (m *MADB) Add() (err error) {

	var (
		b Baselet
	)

	b, err = NewBaselet(
		fmt.Sprintf("/var/monotempo-data/%d.db", len(m.databases)))

	if err != nil {

		return
	}

	m.databases = append(m.databases, b)

	return
}

func (m *MADB) Grow(amount int) (err error) {

	for range amount {

		err = m.Add()

		if err != nil {

			return
		}
	}

	m.maxValue += m.groupSize * amount

	return
}

func (m *MADB) Insert(a *athlete.Atleta) (err error) {

	if a.Numero > m.maxValue {

		err = m.Grow((a.Numero - m.maxValue) / m.groupSize)

		if err != nil {

			return
		}
	}

	err = m.databases[a.Numero].Insert(a)

	return
}

func (m *MADB) Close() {

	for _, v := range m.databases {

		v.Close()
	}
}
