package log

import (
	"fmt"
	"net"
	"os"
	"runtime"

	"github.com/kiem-toan/pkg/config"

	logrustash "github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/openzipkin/zipkin-go"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

const (
	appNameFieldKey = "applicationName"
)

func SetLoglevel(level string) {
	switch level {
	case "trace":
		log.SetLevel(log.TraceLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	default:
		log.SetLevel(log.ErrorLevel)
	}
}

func RegisterLogStash() {
	SetLoglevel("info")
	logst := logrus.New()
	conn, err := net.Dial("tcp", config.GetAppConfig().LogStash.IP+`:`+config.GetAppConfig().LogStash.Port)
	if err != nil {
		Fatal(err, nil, nil)
	}
	hook := logrustash.New(conn, &log.TextFormatter{
		TimestampFormat: "02-01-2006 15:04:05",
		FullTimestamp:   true,
	})
	logst.Hooks.Add(hook)
	log.AddHook(hook)
}

func Trace(err error, span zipkin.Span, fields map[string]interface{}) {
	if err != nil {
		logFields := getLogFields(span, fields)
		log.WithFields(logFields).Trace(err)
	}
}

func Debug(err error, span zipkin.Span, fields map[string]interface{}) {
	if err != nil {
		logFields := getLogFields(span, fields)
		log.WithFields(logFields).Debug(err)
	}
}

func Info(message string, span zipkin.Span, fields map[string]interface{}) {
	logFields := getLogFields(span, fields)
	log.WithFields(logFields).Info(message)
}

func Warn(message string, span zipkin.Span, fields map[string]interface{}) {
	logFields := getLogFields(span, fields)
	log.WithFields(logFields).Warn(message)
}

func Error(err error, span zipkin.Span, fields map[string]interface{}) {
	if err != nil {
		logFields := getLogFields(span, fields)
		log.WithFields(logFields).Error(err)
	}
}

func Fatal(err error, span zipkin.Span, fields map[string]interface{}) {
	if err != nil {
		logFields := getLogFields(span, fields)
		log.WithFields(logFields).Fatal(err)
	}
}

func Panic(err error, span zipkin.Span, fields map[string]interface{}) {
	if err != nil {
		logFields := getLogFields(span, fields)
		log.WithFields(logFields).Panic(err)
	}
}

func getLogFields(span zipkin.Span, fields map[string]interface{}) log.Fields {
	// Init log fields
	f := log.Fields{
		appNameFieldKey: os.Getenv("APPLICATION_NAME"),
	}

	// Always include the original location
	_, file, line, _ := runtime.Caller(2)
	f["LocationFile"] = fmt.Sprintf("%v:%v", file, line)

	// Assign log fields
	for k, v := range fields {
		f[k] = v
	}

	// Assign trace_id field
	if span != nil {
		if span.Context().ParentID == nil {
			f["TraceID"] = span.Context().TraceID.String()
		} else {
			f["TraceID"] = span.Context().ParentID.String()
		}
	}

	return f
}
