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
		Status: gatewayv1.HTTPRouteStatus{},
	}

	hr.SetGroupVersionKind(HTTPRouteGVK)

	return &HTTPRoute{
		Resource: hr,
	}
}

// Setup HTTPRoute values from an Ingress resource
func (r *HTTPRoute) WithIngress(ingress *networkingv1.Ingress) *HTTPRoute {
	r.Ingress = ingress
	r.Resource.ObjectMeta.Name = ingress.Name
	r.Resource.ObjectMeta.Namespace = ingress.Namespace

	r.Resource.Spec.Hostnames = []gatewayv1.Hostname{gatewayv1.Hostname(ingress.Spec.Rules[0].Host)}
	return r
}

func (r *HTTPRoute) WithGateway(name, namespace, sectionName string) *HTTPRoute {
	r.Resource.Spec.ParentRefs = []gatewayv1.ParentReference{
		{
			Name:      gatewayv1.ObjectName(name),
			Namespace: (*gatewayv1.Namespace)(&namespace),
		},
	}
	if sectionName != "" {
		r.Resource.Spec.ParentRefs[0].SectionName = (*gatewayv1.SectionName)(&sectionName)
	}

	return r
}
