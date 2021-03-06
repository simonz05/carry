// Copyright 2014 Simon Zimmermann. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package types

import (
	"fmt"
	"time"
)

type StatKind int

const (
	_                    = iota
	CounterKind StatKind = iota
	ValueKind
)

func (s *StatKind) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"value"`:
		*s = ValueKind
	case `"count"`:
		*s = CounterKind
	default:
		return fmt.Errorf("invalid StatKind %s", string(data))
	}
	return nil
}

func (s *StatKind) MarshalJSON() ([]byte, error) {
	switch *s {
	case ValueKind:
		return []byte(`"value"`), nil
	case CounterKind:
		return []byte(`"count"`), nil
	}
	panic("not reached")
}

func (s *StatKind) UnmarshalText(data []byte) error {
	switch string(data) {
	case "value":
		*s = ValueKind
	case "count":
		*s = CounterKind
	default:
		return fmt.Errorf("invalid StatKind %s, len %d", string(data), len(data))
	}
	return nil
}

func (s *StatKind) MarshalText() ([]byte, error) {
	switch *s {
	case ValueKind:
		return []byte("value"), nil
	case CounterKind:
		return []byte("count"), nil
	}
	panic("not reached")
}

func (s StatKind) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

type Stat struct {
	Key       string   `json:"k"`
	Value     float64  `json:"v"`
	Timestamp int64    `json:"t"`
	Type      StatKind `json:"c"`
}

func (s *Stat) String() string {
	return fmt.Sprintf("stat: %s, status: %f, timestamp: %v, type: %d",
		s.Key, s.Value, time.Unix(s.Timestamp, 0).UTC(), s.Type)
}
