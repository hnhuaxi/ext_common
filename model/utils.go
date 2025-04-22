package model

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func Ptimestamp(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}
	return timestamppb.New(*t)
}

func Timestamp(t time.Time) *timestamppb.Timestamp {
	if t.IsZero() {
		return nil
	}

	return timestamppb.New(t)
}

// Ptime converts a protobuf timestamp to a time.Time.
func Ptime(t *timestamppb.Timestamp) *time.Time {
	if t == nil {
		return nil
	}
	tt := t.AsTime()
	return &tt
}

// time converts a protobuf timestamp to a time.Time.
func Ti(t *timestamppb.Timestamp) time.Time {
	if t == nil {
		return time.Time{}
	}
	return t.AsTime()
}
