package log

import . "github.com/bamgoo/base"

func Levels() map[Level]string {
	out := make(map[Level]string, len(levelStrings))
	for k, v := range levelStrings {
		out[k] = v
	}
	return out
}

func Write(level Level, args ...Any) {
	module.Logging(level, args...)
}

func Debug(args ...Any)   { module.Logging(LevelDebug, args...) }
func Trace(args ...Any)   { module.Logging(LevelTrace, args...) }
func Info(args ...Any)    { module.Logging(LevelInfo, args...) }
func Notice(args ...Any)  { module.Logging(LevelNotice, args...) }
func Warning(args ...Any) { module.Logging(LevelWarning, args...) }
func Error(args ...Any)   { module.Logging(LevelError, args...) }
func Panic(args ...Any) {
	module.Logging(LevelPanic, args...)
	panic(module.parseBody(args...))
}
func Fatal(args ...Any) { module.Logging(LevelFatal, args...) }
