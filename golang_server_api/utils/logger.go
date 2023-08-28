package utils

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

const (
	InvalidToken = "Invalid Token !"
)

type Zlog interface {
	LogInfo(ctx *fiber.Ctx, strFormat string, v ...interface{})
	LogDebug(ctx *fiber.Ctx, strFormat string, v ...interface{})
	LogWarn(ctx *fiber.Ctx, strFormat string, v ...interface{})
	LogError(err error, ctx *fiber.Ctx, strFormat string, v ...interface{})
	LogFatal(err error, ctx *fiber.Ctx, strFormat string, v ...interface{})
	LogPanic(err error, ctx *fiber.Ctx, strFormat string, v ...interface{})
}

type Logger interface {
	Zlog
	myBoost() Zlog
}

type logger struct {
	logger *zerolog.Logger
	*sync.Mutex
}

func NewLogger(cfg Config) Zlog {
	path := cfg.GetString(LoggingPath)
	name := cfg.GetString(AppName)
	folderName := fmt.Sprintf("%s/%s", path, name)

	if err := os.MkdirAll(folderName, 0664); err != nil {
		fmt.Println("Failed to create folder :", err)
	}

	today := time.Now().Format("20060102")
	fileName := fmt.Sprintf("%s/%s.log", folderName, today)
	runLogFile, _ := os.OpenFile(
		fileName,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)

	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
	multi := zerolog.MultiLevelWriter(consoleWriter, runLogFile)
	zlog := zerolog.New(multi).With().Timestamp().Logger()

	return &logger{
		&zlog,
		&sync.Mutex{},
	}
}

func (l *logger) LogInfo(ctx *fiber.Ctx, strFormat string, v ...interface{}) {
	l.generateLog(zerolog.InfoLevel, ctx, nil, strFormat, v...)
}

func (l *logger) LogDebug(ctx *fiber.Ctx, strFormat string, v ...interface{}) {
	l.generateLog(zerolog.DebugLevel, ctx, nil, strFormat, v...)
}

func (l *logger) LogWarn(ctx *fiber.Ctx, strFormat string, v ...interface{}) {
	l.generateLog(zerolog.WarnLevel, ctx, nil, strFormat, v...)
}

func (l *logger) LogError(err error, ctx *fiber.Ctx, strFormat string, v ...interface{}) {
	l.generateLog(zerolog.ErrorLevel, ctx, err, strFormat, v...)
}

func (l *logger) LogFatal(err error, ctx *fiber.Ctx, strFormat string, v ...interface{}) {
	l.generateLog(zerolog.FatalLevel, ctx, err, strFormat, v...)
}

func (l *logger) LogPanic(err error, ctx *fiber.Ctx, strFormat string, v ...interface{}) {
	l.generateLog(zerolog.PanicLevel, ctx, err, strFormat, v...)
}

func (l *logger) generateLog(
	level zerolog.Level,
	ctx *fiber.Ctx,
	err error,
	strFormat string,
	v ...interface{},
) {
	var event *zerolog.Event
	switch level {
	case zerolog.InfoLevel:
		event = l.logger.Info()
	case zerolog.DebugLevel:
		event = l.logger.Debug()
	case zerolog.WarnLevel:
		event = l.logger.Warn()
	case zerolog.ErrorLevel:
		event = l.logger.Error()
	case zerolog.FatalLevel:
		event = l.logger.Fatal()
	case zerolog.PanicLevel:
		event = l.logger.Panic()
	}

	if err != nil {
		event.Err(err)
	}

	if ctx != nil {
		r := ctx.Context()

		event.Str("Connection_ID", strconv.FormatUint(r.ConnID(), 10)).
			Str("IP", ctx.IP()).
			Str("Path", string(r.Path())).
			Str("Method", ctx.Method()).
			Str("Url", string(r.URI().RequestURI())).
			Str("UserAgent", string(r.UserAgent())).
			Msgf(strFormat, v...)
	} else {
		event.Msgf(strFormat, v...)
	}
}

// func gcpLogging(c *fiber.Ctx, strFormat string, v interface{}) {
// 	ctx := context.Background()

// 	env, _ := NewViper()
// 	var gcpLog = env.GetBool(GCPLogging)
// 	var projectID = env.GetString(GCPProjectId)

// 	if !gcpLog || (projectID == "") {
// 		return
// 	}

// 	credential, err := GetGCPSettings()

// 	if err != nil {
// 		return
// 	}

// 	client, err := logging.NewClient(ctx, projectID, option.WithCredentialsJSON([]byte(credential)))

// 	if err != nil {
// 		log.Fatalf("Failed to create client : %v", err)
// 	}
// 	defer func(client *logging.Client) {
// 		_ = client.Close()
// 	}(client)

// 	appName := env.GetString(AppName)
// 	logName := fmt.Sprintf("projects/%s/logs/%s", projectID, appName)
// 	logger := client.Logger(logName)

// 	entry := logging.Entry{
// 		Payload:  fmt.Sprintf(strFormat, v),
// 		Severity: logging.Info,
// 	}

// 	if err := logger.LogSync(ctx, entry); err != nil {
// 		log.Fatalf("Failed to log entry: %v", err)
// 	}
// }
