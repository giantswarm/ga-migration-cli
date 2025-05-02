package httproute

import (
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	gatewayv1 "sigs.k8s.io/gateway-api/apis/v1"
)

func NewFromIngress(ingress *networkingv1.Ingress) *gatewayv1.HTTPRoute {
	httproute := &gatewayv1.HTTPRoute{
		TypeMeta: metav1.TypeMeta{
			Kind:       "HTTPRoute",
			APIVersion: "gateway.networking.k8s.io/v1alpha1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      ingress.Name,
			Namespace: ingress.Namespace,
		},
		Spec: gatewayv1.HTTPRouteSpec{
			Rules: []gatewayv1.HTTPRouteRule{},
		},
	}

	httproute.Spec.ParentRefs = []gatewayv1.ParentReference{
		{
			Name: "TODO: pass by argument",
		},
	}

	return httproute
}
