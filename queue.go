package queue // import "kkn.fi/queue"

import "errors"

//go:generate gends templates.json queue.tmpl

// ErrEmptyQueue error signals that the queue is empty.
var ErrEmptyQueue = errors.New("queue is empty")
