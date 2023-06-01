package logger

import (
	"context"
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"graded/config"

	"github.com/gofrs/uuid"
)

const (
	DEBUG = iota
	INFO
	WARNING
	ERROR
	FATAL
)

type logContext string

const (
	logIDContext logContext = "logIDContext"
)

type Logger struct {
	output   io.Writer
	outputMu sync.Mutex

	level int

	track   map[string]string
	trackMu sync.Mutex
}

func Init(config *config.Config) (*Logger, error) {
	file, err := os.OpenFile("./logs/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return &Logger{
		output:   file,
		outputMu: sync.Mutex{},

		level: config.LOGGER.LEVEL,

		track:   make(map[string]string),
		trackMu: sync.Mutex{},
	}, nil
}

func (l *Logger) StartTrack(ctx context.Context) (context.Context, error) {
	logID, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	l.trackMu.Lock()
	defer l.trackMu.Unlock()
	l.track[logID.String()] = ""
	return context.WithValue(ctx, logIDContext, logID), nil
}

func (l *Logger) StopTrack(ctx context.Context) {
	l.trackMu.Lock()
	defer l.trackMu.Unlock()

	logID := ctx.Value(logIDContext).(string)
	logs := l.track[logID]

	delete(l.track, logID)

	l.outputMu.Lock()
	defer l.outputMu.Unlock()

	l.output.Write([]byte(logs))
}

func (l *Logger) Push(ctx context.Context, lvl int, msg string) {
	if lvl < l.level {
		return
	}

	header := ""
	switch lvl {
	case DEBUG:
		header = "DEBUG"
	case INFO:
		header = "INFO"
	case WARNING:
		header = "WARNING"
	case ERROR:
		header = "ERROR"
	case FATAL:
		header = "FATAL"
	}

	logID := ctx.Value(logIDContext).(string)
	time := time.Now().Format(time.TimeOnly)

	message := fmt.Sprintf("%s %s |\t%s:\t%s\n", time, logID, header, msg)

	l.trackMu.Lock()
	defer l.trackMu.Unlock()
	l.track[logID] += message
}

func (l *Logger) Debug(ctx context.Context, msg string) {
	l.Push(ctx, DEBUG, msg)
}

func (l *Logger) Info(ctx context.Context, msg string) {
	l.Push(ctx, INFO, msg)
}

func (l *Logger) Warning(ctx context.Context, msg string) {
	l.Push(ctx, WARNING, msg)
}

func (l *Logger) Error(ctx context.Context, msg string) {
	l.Push(ctx, ERROR, msg)
}

func (l *Logger) Fatal(ctx context.Context, msg string) {
	l.Push(ctx, FATAL, msg)
}
