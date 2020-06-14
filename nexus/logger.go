package nexus

import "go.uber.org/zap"

func getLogger(verbose bool) *zap.SugaredLogger {
	var logger *zap.Logger
	var err error
	if verbose {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		panic(err)
	}

	defer syncLogger(logger)

	return logger.Sugar()
}

func syncLogger(logger *zap.Logger) {
	err := logger.Sync()
	if err != nil {
		// Let the messages in DEBUG mode only
		// see: https://github.com/uber-go/zap/issues/772
		// see: https://github.com/uber-go/zap/issues/370
		logger.Sugar().Debugf("Failed to sync Sugered log: ", err)
	}
}
