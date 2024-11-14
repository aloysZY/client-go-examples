package deployment

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
)

func DeleteDeployment(dpClient v1.DeploymentInterface) error {
	// Kubernetes 提供了三种删除传播策略：
	// metav1.DeletePropagationOrphan：仅删除父资源，子资源将变成孤儿（orphan），不再受父资源管理。
	// metav1.DeletePropagationBackground：立即删除父资源，子资源将在后台异步删除。
	// metav1.DeletePropagationForeground：先删除子资源，然后再删除父资源。这是默认的行为。
	deletepolicy := metav1.DeletePropagationForeground
	return dpClient.Delete(context.TODO(), "my-deployment", metav1.DeleteOptions{PropagationPolicy: &deletepolicy})
}
