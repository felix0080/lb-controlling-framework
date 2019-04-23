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

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "git.tencent.com/tke/lb-controlling-framework/pkg/apis/lbcf.tke.cloud.tencent.com/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// LoadBalancerLister helps list LoadBalancers.
type LoadBalancerLister interface {
	// List lists all LoadBalancers in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.LoadBalancer, err error)
	// LoadBalancers returns an object that can list and get LoadBalancers.
	LoadBalancers(namespace string) LoadBalancerNamespaceLister
	LoadBalancerListerExpansion
}

// loadBalancerLister implements the LoadBalancerLister interface.
type loadBalancerLister struct {
	indexer cache.Indexer
}

// NewLoadBalancerLister returns a new LoadBalancerLister.
func NewLoadBalancerLister(indexer cache.Indexer) LoadBalancerLister {
	return &loadBalancerLister{indexer: indexer}
}

// List lists all LoadBalancers in the indexer.
func (s *loadBalancerLister) List(selector labels.Selector) (ret []*v1beta1.LoadBalancer, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.LoadBalancer))
	})
	return ret, err
}

// LoadBalancers returns an object that can list and get LoadBalancers.
func (s *loadBalancerLister) LoadBalancers(namespace string) LoadBalancerNamespaceLister {
	return loadBalancerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LoadBalancerNamespaceLister helps list and get LoadBalancers.
type LoadBalancerNamespaceLister interface {
	// List lists all LoadBalancers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.LoadBalancer, err error)
	// Get retrieves the LoadBalancer from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.LoadBalancer, error)
	LoadBalancerNamespaceListerExpansion
}

// loadBalancerNamespaceLister implements the LoadBalancerNamespaceLister
// interface.
type loadBalancerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LoadBalancers in the indexer for a given namespace.
func (s loadBalancerNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.LoadBalancer, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.LoadBalancer))
	})
	return ret, err
}

// Get retrieves the LoadBalancer from the indexer for a given namespace and name.
func (s loadBalancerNamespaceLister) Get(name string) (*v1beta1.LoadBalancer, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("loadbalancer"), name)
	}
	return obj.(*v1beta1.LoadBalancer), nil
}
