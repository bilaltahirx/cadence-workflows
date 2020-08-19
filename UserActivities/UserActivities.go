package UserActivities

import (
	"context"
	"go.uber.org/cadence/activity"
	"go.uber.org/zap"
)
func init()  {

	activity.Register(insertUserToPg)

}

func insertUserToPg(_ context.Context , abc string) (* string,error) {

	logger, _ := zap.NewProduction()

	logger.Info("I ma here")

	return nil,nil

}

