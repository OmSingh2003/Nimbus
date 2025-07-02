package util

import (
	"time"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// TimeToTimestamppb converts time.Time to *timestamppb.Timestamp
func TimeToTimestamppb(t time.Time) *timestamppb.Timestamp {
	return timestamppb.New(t)
}
