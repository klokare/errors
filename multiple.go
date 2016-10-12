package errors

import (
	"bytes"
	"fmt"
)

type Multiple struct {
	items []error
}

func (m *Multiple) Add(err error) {
	if err != nil {
		m.items = append(m.items, err)
	}
}

func (m Multiple) Error() string {
	switch len(m.items) {
	case 0:
		return ""
	case 1:
		return m.items[0].Error()
	default:
		b := bytes.NewBufferString("Multiple errors:\n")
		for i, err := range m.items {
			b.WriteString(fmt.Sprintf("%d - %v\n", i, err))
		}
		return b.String()
	}
}

func (m Multiple) Err() error {
	msg := m.Error()
	if msg == "" {
		return nil
	}
	return fatal{message: msg, isfatal: m.IsFatal()}
}

func (m Multiple) IsFatal() bool {
	for _, err := range m.items {
		if IsFatal(err) {
			return true
		}
	}
	return false
}
