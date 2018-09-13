package router

import (
	"strconv"
	"strings"
)

// Segments returns all segments in path.
func Segments(path string) []string {
	return strings.Split(path, "/")[1:]
}

// Segment returns the path's segment at the specified index.
func Segment(path string, index int) string {
	return Segments(path)[index]
}

// IntSegment returns the path's segment at the specified index as an int.
func IntSegment(path string, index int) (int, error) {
	return strconv.Atoi(Segment(path, index))
}

// Int64Segment returns the path's segment at the specified index as an int64.
func Int64Segment(path string, index int) (int64, error) {
	return strconv.ParseInt(Segment(path, index), 10, 64)
}
