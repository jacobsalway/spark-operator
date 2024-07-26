package yunikorn

import (
	v1 "k8s.io/api/core/v1"

	"github.com/kubeflow/spark-operator/api/v1beta2"
	"github.com/kubeflow/spark-operator/internal/scheduler"
	"github.com/kubeflow/spark-operator/pkg/common"
)

// Defined separately rather than imported to include tags for JSON marshalling
// https://github.com/apache/yunikorn-k8shim/blob/207e4031c6484c965fca4018b6b8176afc5956b4/pkg/cache/amprotocol.go#L47-L56
type taskGroup struct {
	Name         string            `json:"name"`
	MinMember    int32             `json:"minMember"`
	MinResource  map[string]string `json:"minResource,omitempty"`
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`
	Tolerations  []v1.Toleration   `json:"tolerations,omitempty"`
	Affinity     *v1.Affinity      `json:"affinity,omitempty"`
	Labels       map[string]string `json:"labels,omitempty"`
}

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
