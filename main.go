package main

import (
	"net/http"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugarLogger *zap.SugaredLogger

func main() {
	InitLogger()
	defer sugarLogger.Sync()
	simpleHttpGet("www.phub.com")
	simpleHttpGet("http://www.p*rnhub.com")
}

func InitLogger() {
	writeSyncer := getLogwriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder        //pang to nga oras pls
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder //dakua ang ERRORS
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogwriter() zapcore.WriteSyncer {
	file, _ := os.Create("./test.log")
	return zapcore.AddSync(file)
}

func simpleHttpGet(url string) {
	sugarLogger.Debugf("Trying to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Errorf("Error fetching URL %s : Error = %s\n\n", url, err)
	} else {
		sugarLogger.Infof("Success! statusCode = %s for URL %s\n\n", resp.Status, url)
		resp.Body.Close()
	}
}
