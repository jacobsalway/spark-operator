package yunikorn

import (
	"github.com/kubeflow/spark-operator/api/v1beta2"
	"testing"
)

func pointer[T any](v T) *T {
	return &v
}

func TestGetInitialExecutors(t *testing.T) {
	testCases := []struct {
		app      *v1beta2.SparkApplication
		expected int32
	}{
		{
			app: &v1beta2.SparkApplication{
				Spec: v1beta2.SparkApplicationSpec{
					Executor: v1beta2.ExecutorSpec{
						Instances: pointer(int32(5)),
					},
				},
			},
			expected: 5,
		},
	}
}
