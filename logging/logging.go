package logging

import "time"

type Entry interface {
	String(key, value string) Entry
	Int(key string, value int) Entry
	Err(err error) Entry
	Bool(key string, value bool) Entry
	Time(key string, value time.Time) Entry
	Duration(key string, duration time.Duration) Entry
	Float64(key string, value float64) Entry
	Uint64(key string, value uint64) Entry

	Logf(message string, args ...interface{})
	Log(message string)
}

type Logger interface {
	Info() Entry
	Error() Entry
	Debug() Entry
	Warn() Entry
	Critical() Entry
}
