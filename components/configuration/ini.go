package configuration

import (
	"errors"
	"strconv"
	"strings"
)

type Entry = map[string]string

type Section = map[string]Entry

type Config struct {
	AllowMultiLineValues bool
}

type Ini struct {
	sections Section
	config   Config
	column   uint
	row      uint
}

func New() *Ini {
	return &Ini{
		sections: make(map[string]map[string]string),
		config: Config{
			AllowMultiLineValues: true,
		},
		column: 0,
		row:    0,
	}
}

func NewWithConfig(config Config) *Ini {
	return &Ini{
		sections: make(map[string]map[string]string),
		config:   config,
	}
}

func (ini *Ini) From(in *string) error {
	var isInMultiLine = false
	lines := strings.Split(*in, "\n")
	for _, line := range lines {

		ini.incrementRow()

		for _, char := range line {
			if (char == '#' || char == '\'') && !isInMultiLine {
				ini.incrementColumn()
				continue
			}

			if char == '\n' {
				ini.incrementRow()
			}

		}
	}

	return nil
}

func (ini *Ini) GetString(section string, key string) (*string, error) {
	s := ini.sections[section]
	if s == nil {
		return nil, errors.New("no section found")
	}

	k := s[key]
	return &k, nil
}

func (ini *Ini) GetInt(section string, key string) (*int, error) {
	v, err := ini.GetString(section, key)
	if err != nil {
		return nil, err
	}
	i, err := strconv.Atoi(*v)
	return &i, err
}

func (ini *Ini) incrementRow() {
	ini.row++
	ini.column = 0
}

func (ini *Ini) incrementColumn() {
	ini.column++
}
