package helper

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func ProtoToTime(field string, val *timestamp.Timestamp) (time.Time, error) {
	if val == nil {
		return time.Time{}, status.Errorf(codes.InvalidArgument, "timestamp %s must not be nil", field)
	}

	if !val.IsValid() {
		return time.Time{}, status.Errorf(codes.InvalidArgument, "timestamp %s must be valid", field)
	}

	return val.AsTime(), nil
}

func TimeToProto(val time.Time) *timestamp.Timestamp {
	if val.IsZero() {
		return nil
	}
	return timestamppb.New(val)
}
