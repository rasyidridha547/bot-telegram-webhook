package helper

import (
	"log"
	"time"
)

func ConvertUnixToDate(raw float64) string {
	// convert to int64 cause Unix only allow int64
	intDateTime := int64(raw)

	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Println(err)
	}

	// convert raw to Unix
	dateTime := time.Unix(intDateTime, 0).In(loc)
	// Monday, 02-Jan-06 15:04:05 MST
	formatDate := dateTime.Format(time.RFC850)

	return formatDate
}
