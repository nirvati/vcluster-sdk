package syncers

import (
	"context"
	"fmt"

	"github.com/nirvati/vcluster-sdk/plugin"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func NewPodHook() plugin.ClientHook {
	return &podHook{}
}

type podHook struct{}

func (p *podHook) Name() string {
	return "pod-hook"
}

func (p *podHook) Resource() client.Object {
	return &corev1.Pod{}
}

var _ plugin.MutateCreatePhysical = &podHook{}

func (p *podHook) MutateCreatePhysical(ctx context.Context, obj client.Object) (client.Object, error) {
	pod, ok := obj.(*corev1.Pod)
	if !ok {
		return nil, fmt.Errorf("object %v is not a pod", obj)
	}

	if pod.Labels == nil {
		pod.Labels = map[string]string{}
	}
	pod.Labels["created-by-plugin"] = "pod-hook"
	return pod, nil
}

var _ plugin.MutateUpdatePhysical = &podHook{}

func (p *podHook) MutateUpdatePhysical(ctx context.Context, obj client.Object) (client.Object, error) {
	pod, ok := obj.(*corev1.Pod)
	if !ok {
		return nil, fmt.Errorf("object %v is not a pod", obj)
	}

	if pod.Labels == nil {
		pod.Labels = map[string]string{}
	}
	pod.Labels["created-by-plugin"] = "pod-hook"
	return pod, nil
}
