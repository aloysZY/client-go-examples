package deployment

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	"k8s.io/client-go/util/retry"
	"k8s.io/utils/pointer"
)

func UpdateDeployment(dpClient v1.DeploymentInterface) error {
	dp, err := dpClient.Get(context.TODO(), "my-deployment", metav1.GetOptions{})
	if err != nil {
		return err
	}
	dp.Spec.Replicas = pointer.Int32Ptr(5)
	// retry.RetryOnConflict 是一个重试机制，当更新操作遇到冲突时会自动重试。
	// retry.DefaultRetry 是一个默认的重试策略，通常包括多次重试和指数退避。
	// func() error 是一个匿名函数，用于执行更新操作并返回错误
	return retry.RetryOnConflict(retry.DefaultRetry, func() error {
		_, err := dpClient.Update(context.TODO(), dp, metav1.UpdateOptions{})
		return err
	})
}
