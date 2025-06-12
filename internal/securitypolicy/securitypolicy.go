package securitypolicy

import (
	// networkingv1 "k8s.io/api/networking/v1"
	eg "github.com/envoyproxy/gateway/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	gatewayv1 "sigs.k8s.io/gateway-api/apis/v1"
	gatewayv1a2 "sigs.k8s.io/gateway-api/apis/v1alpha2"
)

type SecurityPolicy struct {
	Resource *eg.SecurityPolicy
}

var (
	SecurityPolicyGVK = schema.GroupVersionKind{
		Group:   "gateway.envoyproxy.io",
		Version: "v1alpha1",
		Kind:    "SecurityPolicy",
	}
)

func New() *SecurityPolicy {
	r := &eg.SecurityPolicy{
		ObjectMeta: metav1.ObjectMeta{},
		Spec:       eg.SecurityPolicySpec{},
		Status:     gatewayv1a2.PolicyStatus{},
	}

	r.SetGroupVersionKind(SecurityPolicyGVK)

	return &SecurityPolicy{
		Resource: r,
	}
}

func (r *SecurityPolicy) WithHTTPRoute(httpRoute *gatewayv1.HTTPRoute) *SecurityPolicy {
	return r
}

func (r *SecurityPolicy) WithAnnotations() *SecurityPolicy {
	return r
}
