//     Copyright 2020 Aicura Nexus Client and/or its authors
//
//     This file is part of Aicura Nexus Client.
//
//     Aicura Nexus Client is free software: you can redistribute it and/or modify
//     it under the terms of the GNU Lesser General Public License as published by
//     the Free Software Foundation, either version 3 of the License, or
//     (at your option) any later version.
//
//     Aicura Nexus Client is distributed in the hope that it will be useful,
//     but WITHOUT ANY WARRANTY; without even the implied warranty of
//     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//     GNU Lesser General Public License for more details.
//
//     You should have received a copy of the GNU Lesser General Public License
//     along with Aicura Nexus Client.  If not, see <https://www.gnu.org/licenses/>.

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
