package mock

import (
	"context"

	"github.com/lyft/flinkk8soperator/pkg/apis/app/v1alpha1"
)

type JobManagerController struct {
	CreateIfNotExistFunc CreateIfNotExistFunc
}

func (m *JobManagerController) CreateIfNotExist(
	ctx context.Context,
	application *v1alpha1.FlinkApplication) (bool, error) {
	if m.CreateIfNotExistFunc != nil {
		return m.CreateIfNotExistFunc(ctx, application)
	}
	return false, nil
}
