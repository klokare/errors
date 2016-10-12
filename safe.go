package errors

import "sync"

type Safe struct {
	Multiple
	sync.Mutex
}

func (s *Safe) Add(err error) {
	s.Lock()
	defer s.Unlock()
	s.Multiple.Add(err)
}

func (s *Safe) Error() string {
	s.Lock()
	defer s.Unlock()
	return s.Multiple.Error()
}

func (s *Safe) Err() error {
	s.Lock()
	defer s.Unlock()
	return s.Multiple.Err()
}

func (s *Safe) IsFatal() bool {
	s.Lock()
	defer s.Unlock()
	return s.Multiple.IsFatal()
}
