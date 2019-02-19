package fixtures

import (
	"gopkg.in/src-d/hercules.v8/internal/plumbing"
	"gopkg.in/src-d/hercules.v8/internal/test"
)

// FileDiff initializes a new plumbing.FileDiff item for testing.
func FileDiff() *plumbing.FileDiff {
	fd := &plumbing.FileDiff{}
	fd.Initialize(test.Repository)
	return fd
}
