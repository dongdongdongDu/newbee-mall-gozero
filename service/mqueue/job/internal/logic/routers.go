package logic

import (
	"context"
	"github.com/hibiken/asynq"
	"newbee-mall-gozero/service/mqueue/job/internal/svc"
	"newbee-mall-gozero/service/mqueue/job/jobtype"
)

type CronJob struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCronJob(ctx context.Context, svcCtx *svc.ServiceContext) *CronJob {
	return &CronJob{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// register job
func (l *CronJob) Register() *asynq.ServeMux {

	mux := asynq.NewServeMux()

	//defer job
	mux.Handle(jobtype.DeferCloseOrder, NewCloseOrderHandler(l.svcCtx))

	//queue job , asynq support queue job
	// wait you fill..

	return mux
}
