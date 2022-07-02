package util

/**
 * @Author: kylo_cheok
 * @Email:  maggic0816@gmail.com
 * @Date:   2022/5/22 23:41
 * @Desc:   Grace under pressure
 */

import (
	"go.uber.org/zap"
)

func InitLogger() (l *zap.Logger, e error) {
	logger, err := zap.NewProduction()
	if err != nil {
		logger.Info("db connect error: ", zap.Error(err))
		return
	}
	return logger, nil
}
