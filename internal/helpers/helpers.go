package helpers

import (
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertStringToTimestamp(timeStr string) (pgtype.Timestamp, error) {
	var ts pgtype.Timestamp
	parsedTime, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return ts, err
	}

	ts.Time = parsedTime
	ts.Valid = true

	return ts, nil
}

func ConvertToPgNumeric(floatValue float64) (pgtype.Numeric, error) {

	stringVal := strconv.FormatFloat(floatValue, 'G', 10, 64)

	numeric := pgtype.Numeric{}

	if err := numeric.Scan(stringVal); err != nil {
		return numeric, err
	}

	return numeric, nil
}

func ConvertToTime(time pgtype.Timestamp) *timestamppb.Timestamp {
	// Assume v.JummahTime is of type pgtype.Timestamp
	var timestamp *timestamppb.Timestamp

	timeValue := time.Time

	timestamp = timestamppb.New(timeValue)

	return timestamp
}
