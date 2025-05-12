package httproute

import (
	gatewayv1 "sigs.k8s.io/gateway-api/apis/v1"
)

// toRules

func toRules() []gatewayv1.HTTPRouteRule {
	rules := make([]gatewayv1.HTTPRouteRule, 0)

	rule := &gatewayv1.HTTPRouteRule{
		BackendRefs: toBackendRefs(),
	}

	rules = append(rules, *rule)
	return rules
}

// toBackendRef
func toBackendRefs() []gatewayv1.HTTPBackendRef {
	kind := gatewayv1.Kind("Service")
	port := gatewayv1.PortNumber(80)
	weith := int32(1)

	backendRef := &gatewayv1.HTTPBackendRef{
		BackendRef: gatewayv1.BackendRef{
			BackendObjectReference: gatewayv1.BackendObjectReference{
				Kind: &kind,
				Name: gatewayv1.ObjectName("service1"),
				Port: &port,
			},
			Weight: &weith,
		},
	}

	backendRefs := make([]gatewayv1.HTTPBackendRef, 0)
	backendRefs = append(backendRefs, *backendRef)

	return backendRefs
}

// toMatches
