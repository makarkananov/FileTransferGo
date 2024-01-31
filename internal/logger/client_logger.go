package logger

// ClientLogger is an interface for logging client-related messages.
type ClientLogger interface {
	// Printf logs a formatted message with optional parameters.
	Printf(format string, v ...interface{})
}
