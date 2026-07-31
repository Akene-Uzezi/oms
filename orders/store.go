package main

import "context"

type store struct {
	// add mongo instace
}

func NewStore() *store {
	return &store{}
}

func (s *store) Create(context.Context) error {
	return nil
}
