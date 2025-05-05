package httproute

import (
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	gatewayv1 "sigs.k8s.io/gateway-api/apis/v1"
)

type HTTPRoute struct {
	Resource *gatewayv1.HTTPRoute
	Ingress  *networkingv1.Ingress
}

var (
	HTTPRouteGVK = schema.GroupVersionKind{
		Group:   "gateway.networking.k8s.io",
		Version: "v1",
		Kind:    "HTTPRoute",
	}
)

func New() *HTTPRoute {
	hr := &gatewayv1.HTTPRoute{
		ObjectMeta: metav1.ObjectMeta{},
		Spec: gatewayv1.HTTPRouteSpec{
			Rules: []gatewayv1.HTTPRouteRule{},
		},
	}

	hr.SetGroupVersionKind(HTTPRouteGVK)
	hr.Status = gatewayv1.HTTPRouteStatus{}

	return &HTTPRoute{
		Resource: hr,
	}
}

// Setup HTTPRoute values from an Ingress resource
func (r *HTTPRoute) WithIngress(ingress *networkingv1.Ingress) *HTTPRoute {
	r.Ingress = ingress
	r.Resource.ObjectMeta.Name = ingress.Name
	r.Resource.ObjectMeta.Namespace = ingress.Namespace
	return r
}
