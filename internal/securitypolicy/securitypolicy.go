package securitypolicy

import (
	eg "github.com/envoyproxy/gateway/api/v1alpha1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	gatewayv1 "sigs.k8s.io/gateway-api/apis/v1"
	gatewayv1a2 "sigs.k8s.io/gateway-api/apis/v1alpha2"
)

const (
	IngressNginxAuthType      = "nginx.ingress.kubernetes.io/auth-type"
	IngressNginxAuthSecret    = "nginx.ingress.kubernetes.io/auth-secret"
	IngressNginxAuthTypeBasic = "basic"
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

func (r *SecurityPolicy) WithIngress(ingress *networkingv1.Ingress) *SecurityPolicy {
	annotations := ingress.GetAnnotations()
	if authType, ok := annotations[IngressNginxAuthType]; ok {
		return r.WithAuth(authType, annotations[IngressNginxAuthSecret])
	}
	return nil
}

func (r *SecurityPolicy) WithAuth(authType, authSecret string) *SecurityPolicy {
	if authType == IngressNginxAuthTypeBasic {
		kind := gatewayv1.Kind("Secret")
		group := gatewayv1.Group("")
		r.Resource.Spec.BasicAuth = &eg.BasicAuth{
			Users: gatewayv1.SecretObjectReference{
				Group: &group,
				Kind:  &kind,
				Name:  gatewayv1.ObjectName(authSecret),
			},
		}
	}

	return r
}
