package main

import (
	"errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	for i:= 0; i < 1000000 ; i++  {
		sampleLogTest()
	}
}
var (
	sampleLogger *zap.Logger
	notSampleLogger *zap.Logger
)


func init() {
	sampleConf := zap.NewProductionConfig()
	var err error
	sampleConf.Sampling = &zap.SamplingConfig{
		Initial:    1,
		Thereafter: 100,
		Hook:       nil,
	}
	sampleConf.OutputPaths= []string{"zap.log"}
	sampleLogger, err = sampleConf.Build()
	if err != nil {
		panic(err)
	}

	notSampleConf := zap.NewProductionConfig()
	notSampleConf.Sampling = nil
	notSampleLogger, err = notSampleConf.Build()
	if err != nil {
		panic(err)
	}
}

func sampleLogTest() {
	sampleLogger.Info("zap info test")

}

func notSampleLogTest() {
	notSampleLogger.Info("zap info test")
	err := errors.New("internal server error")
	notSampleLogger.Error("zap error test", zap.Error(err))
}

func SampleHook(entry zapcore.Entry,dec zapcore.SamplingDecision) {
	if dec&zapcore.LogDropped > 0 && entry.Level == zapcore.ErrorLevel {
		zap.L().Error(entry.Message,)
	}
}
