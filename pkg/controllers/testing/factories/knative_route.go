/*
Copyright 2019 the original author or authors.

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

package factories

import (
	"fmt"

	"github.com/vmware-labs/reconciler-runtime/apis"
	rtesting "github.com/vmware-labs/reconciler-runtime/testing"

	duckv1 "github.com/projectriff/system/pkg/apis/duck/v1"
	knativeservingv1 "github.com/projectriff/system/pkg/apis/thirdparty/knative/serving/v1"
)

type knativeRoute struct {
	target *knativeservingv1.Route
}

var (
	_ rtesting.Factory = (*knativeRoute)(nil)
)

func KnativeRoute(seed ...*knativeservingv1.Route) *knativeRoute {
	var target *knativeservingv1.Route
	switch len(seed) {
	case 0:
		target = &knativeservingv1.Route{}
	case 1:
		target = seed[0]
	default:
		panic(fmt.Errorf("expected exactly zero or one seed, got %v", seed))
	}
	return &knativeRoute{
		target: target,
	}
}

func (f *knativeRoute) deepCopy() *knativeRoute {
	return KnativeRoute(f.target.DeepCopy())
}

func (f *knativeRoute) Create() *knativeservingv1.Route {
	return f.deepCopy().target
}

func (f *knativeRoute) CreateObject() apis.Object {
	return f.Create()
}

func (f *knativeRoute) mutation(m func(*knativeservingv1.Route)) *knativeRoute {
	f = f.deepCopy()
	m(f.target)
	return f
}

func (f *knativeRoute) NamespaceName(namespace, name string) *knativeRoute {
	return f.mutation(func(route *knativeservingv1.Route) {
		route.ObjectMeta.Namespace = namespace
		route.ObjectMeta.Name = name
	})
}

func (f *knativeRoute) ObjectMeta(nf func(ObjectMeta)) *knativeRoute {
	return f.mutation(func(route *knativeservingv1.Route) {
		omf := ObjectMetaFactory(route.ObjectMeta)
		nf(omf)
		route.ObjectMeta = omf.Create()
	})
}

func (f *knativeRoute) Traffic(traffic ...knativeservingv1.TrafficTarget) *knativeRoute {
	return f.mutation(func(route *knativeservingv1.Route) {
		route.Spec.Traffic = traffic
	})
}

func (f *knativeRoute) StatusConditions(conditions ...ConditionFactory) *knativeRoute {
	return f.mutation(func(route *knativeservingv1.Route) {
		c := make([]apis.Condition, len(conditions))
		for i, cg := range conditions {
			c[i] = cg.Create()
		}
		route.Status.Conditions = c
	})
}

func (f *knativeRoute) StatusReady() *knativeRoute {
	return f.StatusConditions(
		Condition().Type(knativeservingv1.RouteConditionReady).True(),
	)
}

func (f *knativeRoute) StatusObservedGeneration(generation int64) *knativeRoute {
	return f.mutation(func(route *knativeservingv1.Route) {
		route.Status.ObservedGeneration = generation
	})
}

func (f *knativeRoute) StatusAddressURL(url string) *knativeRoute {
	return f.mutation(func(route *knativeservingv1.Route) {
		route.Status.Address = &duckv1.Addressable{
			URL: url,
		}
	})
}

func (f *knativeRoute) StatusURL(url string) *knativeRoute {
	return f.mutation(func(route *knativeservingv1.Route) {
		route.Status.URL = url
	})
}
