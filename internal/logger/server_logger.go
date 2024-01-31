package logger

// ServerLogger is an interface for logging server-related messages.
type ServerLogger interface {
	// Printf logs a formatted message with optional parameters.
	Printf(format string, v ...interface{})
}
