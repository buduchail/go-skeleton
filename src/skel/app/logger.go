package app

type (
	// The stdlib logger does not provide an interface, so it's not possible to use
	// interchangeable implementations by default. The logrus library provides this
	// kind of interface. StdLogger is an interface compatible with the stdlib logger,
	// while FieldLogger is an enhanced logger interface that adds structured logging
	// capabilities.
	// To be able to easily use a context when logging, we'll define a new interface
	// inspired in FieldLogger and add a LoggerContext type on top of it.
	// This interface defines a structured logger that accepts a context in each
	// invocation. It is not compatible with the stdlib logger.
	// Implementations can also accept a default context at creation time

	Logger interface {
		Debug(message string, context *LoggerContext)
		Info(message string, context *LoggerContext)
		Print(message string, context *LoggerContext)
		Warn(message string, context *LoggerContext)
		Warning(message string, context *LoggerContext)
		Error(message string, context *LoggerContext)
		Fatal(message string, context *LoggerContext)
		Panic(message string, context *LoggerContext)
	}

	LoggerContext map[string]interface{}
)
