package workflows

import (
	"github.com/bilaltahirx/cadence-workflows/activites"
	"go.uber.org/cadence/workflow"
	"time"
)

func init()  {
	workflow.Register(CreateUserWorkFlow)

}

func CreateUserWorkFlow(ctx workflow.Context)  error{
	logger := workflow.GetLogger(ctx)


	logger.Info("IOWSKNCkndjvds dvbsdkjvbsdjb")

	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var resp string
	if err := workflow.ExecuteActivity(ctx, activites.InsertUserToPg, "podsRequest").Get(ctx, &resp); err != nil {
		//Message not published on nats if error occurs
		logger.Error(err.Error())
		return err
	}
	return nil
}