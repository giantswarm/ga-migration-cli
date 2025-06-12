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

// // Setup HTTPRoute values from an Ingress resource
// func (r *HTTPRoute) WithIngress(ingress *networkingv1.Ingress) *HTTPRoute {
// 	r.Ingress = ingress
// 	r.Resource.ObjectMeta.Name = ingress.Name
// 	r.Resource.ObjectMeta.Namespace = ingress.Namespace
//
// 	r.Resource.Spec.Hostnames = []gatewayv1.Hostname{gatewayv1.Hostname(ingress.Spec.Rules[0].Host)}
//
// 	rules := toRules(ingress.Spec.Rules)
// 	r.Resource.Spec.Rules = rules
//
// 	return r
// }
//
// func (r *HTTPRoute) WithGateway(name, namespace, sectionName string) *HTTPRoute {
// 	r.Resource.Spec.ParentRefs = []gatewayv1.ParentReference{
// 		{
// 			Name:      gatewayv1.ObjectName(name),
// 			Namespace: (*gatewayv1.Namespace)(&namespace),
// 		},
// 	}
// 	if sectionName != "" {
// 		r.Resource.Spec.ParentRefs[0].SectionName = (*gatewayv1.SectionName)(&sectionName)
// 	}
//
// 	return r
// }
