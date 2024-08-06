package pocketlog

type Level byte

const (
	// represents the lowest level of log, used for debuging
	// iota lets the compiler know we are starting an enumeration
	LevelDebug Level = iota
	
	// represents a level that contains info 
	LevelInfo

	// useed to trace errors, represents highest level of logging
	LevelError
)