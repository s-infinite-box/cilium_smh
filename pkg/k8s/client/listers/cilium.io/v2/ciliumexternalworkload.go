// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by lister-gen. DO NOT EDIT.

package v2

import (
	ciliumiov2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// CiliumExternalWorkloadLister helps list CiliumExternalWorkloads.
// All objects returned here must be treated as read-only.
type CiliumExternalWorkloadLister interface {
	// List lists all CiliumExternalWorkloads in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*ciliumiov2.CiliumExternalWorkload, err error)
	// Get retrieves the CiliumExternalWorkload from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*ciliumiov2.CiliumExternalWorkload, error)
	CiliumExternalWorkloadListerExpansion
}

// ciliumExternalWorkloadLister implements the CiliumExternalWorkloadLister interface.
type ciliumExternalWorkloadLister struct {
	listers.ResourceIndexer[*ciliumiov2.CiliumExternalWorkload]
}

// NewCiliumExternalWorkloadLister returns a new CiliumExternalWorkloadLister.
func NewCiliumExternalWorkloadLister(indexer cache.Indexer) CiliumExternalWorkloadLister {
	return &ciliumExternalWorkloadLister{listers.New[*ciliumiov2.CiliumExternalWorkload](indexer, ciliumiov2.Resource("ciliumexternalworkload"))}
}
