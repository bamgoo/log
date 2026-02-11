package log

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

type Instance struct {
	connect Connect

	Name    string
	Config  Config
	Setting map[string]any
}

func (inst *Instance) Format(entry Log) string {
	if inst.Config.Json {
		body, _ := json.Marshal(map[string]any{
			"time":  entry.Time.Format("2006-01-02 15:04:05.000"),
			"unix":  entry.Time.Unix(),
			"nano":  entry.Time.UnixNano(),
			"level": levelStrings[entry.Level],
			"name":  inst.Name,
			"flag":  inst.Config.Flag,
			"body":  entry.Body,
		})
		return string(body)
	}

	message := inst.Config.Format
	if message == "" {
		message = "%time% [%level%] %body%"
	}

	message = strings.ReplaceAll(message, "%nano%", strconv.FormatInt(entry.Time.UnixNano(), 10))
	message = strings.ReplaceAll(message, "%unix%", strconv.FormatInt(entry.Time.Unix(), 10))
	message = strings.ReplaceAll(message, "%time%", entry.Time.Format("2006-01-02 15:04:05.000"))
	message = strings.ReplaceAll(message, "%name%", inst.Name)
	message = strings.ReplaceAll(message, "%flag%", inst.Config.Flag)
	message = strings.ReplaceAll(message, "%level%", levelStrings[entry.Level])
	message = strings.ReplaceAll(message, "%body%", entry.Body)

	return message
}

func (inst *Instance) Allow(level Level) bool {
	return inst.Config.Levels[level]
}

func normalizeLevels(cfg Config) Config {
	if len(cfg.Levels) > 0 {
		return cfg
	}

	cfg.Levels = map[Level]bool{}
	for level := range levelStrings {
		if level <= cfg.Level {
			cfg.Levels[level] = true
		}
	}
	return cfg
}

func normalizeConfig(cfg Config) Config {
	if cfg.Driver == "" {
		cfg.Driver = "default"
	}
	if cfg.Levels == nil {
		cfg.Levels = map[Level]bool{}
	}
	if cfg.Level < LevelFatal || cfg.Level > LevelDebug {
		cfg.Level = LevelDebug
	}
	if cfg.Format == "" {
		cfg.Format = "%time% [%level%] %body%"
	}
	if cfg.Buffer <= 0 {
		cfg.Buffer = 1024
	}
	if cfg.Timeout <= 0 {
		cfg.Timeout = time.Millisecond * 200
	}
	cfg = normalizeLevels(cfg)
	return cfg
}
