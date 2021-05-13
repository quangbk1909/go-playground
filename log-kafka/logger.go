package loggers

import (
	"fmt"
	"go.uber.org/zap"
	"strings"
)

const (
	KafkaLogPrefix     = "sarama: "
	NonFormatDelimiter = "; "
)

var (
	defaultZapFields = []zap.Field{zap.String("log_type", "sarama")}
)

type KafkaLogger struct {
	zapLogger *zap.Logger
}

func NewKafkaLogger() (*KafkaLogger, error) {
	zl, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	return &KafkaLogger{
		zapLogger: zl,
	}, nil
}

//func initSaramaLog() {
//	l, err := loggers.NewKafkaLogger()
//	if err != nil {
//		log.Error(nil, "cannot init kafka logger, using noop log", err)
//		return
//	}
//	sarama.Logger = l
//
//}

func (l *KafkaLogger) Print(v ...interface{}) {
	l.zapLogger.Info(l.buildLogString(v...), defaultZapFields...)
}
func (l *KafkaLogger) Printf(format string, v ...interface{}) {
	f := fmt.Sprintf("%s%s", KafkaLogPrefix, format)
	l.zapLogger.Info(fmt.Sprintf(f, v...), defaultZapFields...)
}
func (l *KafkaLogger) Println(v ...interface{}) {
	l.zapLogger.Info(l.buildLogString(v...), defaultZapFields...)
}
func (l *KafkaLogger) buildLogString(v ...interface{}) string {
	var ret strings.Builder
	ret.WriteString(KafkaLogPrefix)
	for i := 0; i < len(v); i++ {
		if i > 0 {
			ret.WriteString(NonFormatDelimiter)
		}
		ret.WriteString(fmt.Sprint(v[i]))
	}
	return ret.String()
}
