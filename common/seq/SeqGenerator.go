package seq

import "github.com/google/uuid"

/**
  ηζεΊε
*/
func Generator() string {
	seq := uuid.New()
	return seq.String()
}
