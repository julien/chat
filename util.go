package main

import (
	"strconv"
	"strings"
	"time"
)

func formatDigit(value int) string {
	res := strconv.FormatInt(int64(value), 10)
	if value < 10 {
		return "0" + res
	}
	return res
}

func addTimeStamp() string {
	now := time.Now()
	times := []string{
		formatDigit(now.Hour()),
		formatDigit(now.Minute()),
		formatDigit(now.Second()),
	}
	res := strings.Join(times, ":")
	return res
}
