package sbuff

import "errors"

// error
var (
	ErrCorrupt = errors.New("[SBuff]: Capacity is full")
)

// Smart buffer
type SBuff struct {
	buff     []byte //buffer
	pointer  int    //last free byte in buffer
	capacity int    //capacity of buffer
}

// create constructor for smart buffer
func NewSBuff(capacity int) *SBuff {
	return &SBuff{
		buff:     make([]byte, capacity),
		capacity: capacity,
	}
}

// append new pice of slice in buffer
func (s *SBuff) Append(pice []byte) error {
	if s.capacity < len(pice)+s.pointer {
		return ErrCorrupt
	}

	//copy data to our smart buffer
	s.pointer += copy(s.buff[s.pointer:], pice)
	return nil
}

// reset pointer
func (s *SBuff) Reset() {
	s.pointer = 0
}

// read from our buffer
func (s SBuff) Read(start, end int) []byte {
	return s.buff[start:end]
}

// get value of buffer
func (s SBuff) GetBuff() []byte {
	return s.buff
}
