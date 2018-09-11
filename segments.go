package router

import (
	"strconv"
	"strings"
)

func Segments(path string) []string {
	return strings.Split(path, "/")[1:]
}

func Segment(path string, index int) string {
	return Segments(path)[index]
}

func IntSegment(path string, index int) (int, error) {
	return strconv.Atoi(Segment(path, index))
}

func Int64Segment(path string, index int) (int64, error) {
	return strconv.ParseInt(Segment(path, index), 10, 64)
}

func Float64Segment(path string, index int) (float64, error) {
	return strconv.ParseFloat(Segment(path, index), 64)
}
