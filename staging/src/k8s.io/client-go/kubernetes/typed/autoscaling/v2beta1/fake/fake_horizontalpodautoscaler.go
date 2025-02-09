/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v2beta1 "k8s.io/api/autoscaling/v2beta1"
	autoscalingv2beta1 "k8s.io/client-go/applyconfigurations/autoscaling/v2beta1"
	gentype "k8s.io/client-go/gentype"
	typedautoscalingv2beta1 "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta1"
)

// fakeHorizontalPodAutoscalers implements HorizontalPodAutoscalerInterface
type fakeHorizontalPodAutoscalers struct {
	*gentype.FakeClientWithListAndApply[*v2beta1.HorizontalPodAutoscaler, *v2beta1.HorizontalPodAutoscalerList, *autoscalingv2beta1.HorizontalPodAutoscalerApplyConfiguration]
	Fake *FakeAutoscalingV2beta1
}

func newFakeHorizontalPodAutoscalers(fake *FakeAutoscalingV2beta1, namespace string) typedautoscalingv2beta1.HorizontalPodAutoscalerInterface {
	return &fakeHorizontalPodAutoscalers{
		gentype.NewFakeClientWithListAndApply[*v2beta1.HorizontalPodAutoscaler, *v2beta1.HorizontalPodAutoscalerList, *autoscalingv2beta1.HorizontalPodAutoscalerApplyConfiguration](
			fake.Fake,
			namespace,
			v2beta1.SchemeGroupVersion.WithResource("horizontalpodautoscalers"),
			v2beta1.SchemeGroupVersion.WithKind("HorizontalPodAutoscaler"),
			func() *v2beta1.HorizontalPodAutoscaler { return &v2beta1.HorizontalPodAutoscaler{} },
			func() *v2beta1.HorizontalPodAutoscalerList { return &v2beta1.HorizontalPodAutoscalerList{} },
			func(dst, src *v2beta1.HorizontalPodAutoscalerList) { dst.ListMeta = src.ListMeta },
			func(list *v2beta1.HorizontalPodAutoscalerList) []*v2beta1.HorizontalPodAutoscaler {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v2beta1.HorizontalPodAutoscalerList, items []*v2beta1.HorizontalPodAutoscaler) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
