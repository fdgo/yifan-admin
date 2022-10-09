package logger

// Debug ..
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

// Debugf ..
func Debugf(msg string, args ...interface{}) {
	logger.Debugf(msg, args...)
}

// Debugw ..
func Debugw(msg string, keysAndValues ...interface{}) {
	logger.Debugw(msg, keysAndValues...)
}

// Info ..
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Infof ..
func Infof(msg string, args ...interface{}) {
	logger.Infof(msg, args...)
}

// Infow ..
func Infow(msg string, keysAndValues ...interface{}) {
	logger.Infow(msg, keysAndValues...)
}

// Warn ..
func Warn(args ...interface{}) {
	logger.Warn(args...)
}

// Warnf ..
func Warnf(msg string, args ...interface{}) {
	logger.Warnf(msg, args...)
}

// Warnw ..
func Warnw(msg string, keysAndValues ...interface{}) {
	logger.Warnw(msg, keysAndValues...)
}

// Error ..
func Error(args ...interface{}) {
	logger.Error(args...)
}

// Errorf ..
func Errorf(msg string, args ...interface{}) {
	logger.Errorf(msg, args...)
}

// Errorw ..
func Errorw(msg string, keysAndValues ...interface{}) {
	logger.Errorw(msg, keysAndValues...)
}

// DPanic log a message, In development, the logger then panics
// DPanic ..
func DPanic(args ...interface{}) {
	logger.DPanic(args...)
}

// DPanicf ..
func DPanicf(msg string, args ...interface{}) {
	logger.DPanicf(msg, args...)
}

// DPanicw ..
func DPanicw(msg string, keysAndValues ...interface{}) {
	logger.DPanicw(msg, keysAndValues...)
}

// Panic log a message, then panics.
// Panic ..
func Panic(args ...interface{}) {
	logger.Panic(args)
}

// Panicf ..
func Panicf(msg string, args ...interface{}) {
	logger.Panicf(msg, args...)
}

// Panicw ..
func Panicw(msg string, keysAndValues ...interface{}) {
	logger.Panicw(msg, keysAndValues...)
}

// Fatal log a message, then calls os.Exit.
// Fatal ..
func Fatal(args ...interface{}) {
	logger.Fatal(args)
}

// Fatalf ..
func Fatalf(msg string, args ...interface{}) {
	logger.Fatalf(msg, args...)
}

// Fatalw ..
func Fatalw(msg string, keysAndValues ...interface{}) {
	logger.Fatalw(msg, keysAndValues...)
}
