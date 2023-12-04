package queries

import "fmt"

type query string

const (
	match  query = "match"
	create query = "create"
)

type (
	Match struct {
		label string
		where map[string]any
	}

	Query interface {
		Serialize() string
	}
)

func NewMatch(label string, where map[string]any) *Match {
	return &Match{
		label: label,
		where: where,
	}
}

func (m *Match) Serialize() string {
	where := ""
	if len(m.where) > 0 {
		where += " WHERE"
		for key, val := range m.where {
			where += fmt.Sprintf(" x.%s = '%s'", key, val)
		}
	}
	return fmt.Sprintf("MATCH (x:%s%s) RETURN x", m.label, where)
}
