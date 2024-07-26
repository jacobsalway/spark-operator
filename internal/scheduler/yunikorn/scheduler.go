package yunikorn

import (
	"github.com/kubeflow/spark-operator/api/v1beta2"
	"github.com/kubeflow/spark-operator/internal/scheduler"
	"github.com/kubeflow/spark-operator/pkg/common"
)

type Scheduler struct{}

// Ensure the Yunikorn scheduler implements the required interface
var _ scheduler.Interface = &Scheduler{}

func Factory(config scheduler.Config) (scheduler.Interface, error) {
	return &Scheduler{}, nil
}

func (s *Scheduler) Name() string {
	return common.YunikornSchedulerName
}

func (s *Scheduler) ShouldSchedule(app *v1beta2.SparkApplication) bool {
	// Yunikorn gets all the information it needs from pod annotations, so
	// there are no additional resources to be created
	return true
}

func (s *Scheduler) Schedule(app *v1beta2.SparkApplication) error {
	return nil
}

func (s *Scheduler) Cleanup(app *v1beta2.SparkApplication) error {
	// No additional resources are created so there's nothing to be cleaned up
	return nil
}
