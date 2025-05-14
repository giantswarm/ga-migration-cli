package httproute

import (
	networkingv1 "k8s.io/api/networking/v1"
	gatewayv1 "sigs.k8s.io/gateway-api/apis/v1"
)

func toRules(ingressRule []networkingv1.IngressRule) []gatewayv1.HTTPRouteRule {
	rules := make([]gatewayv1.HTTPRouteRule, 0)

	// Only rule 0 is supported for now
	rule := &gatewayv1.HTTPRouteRule{
		BackendRefs: toBackendRefs(ingressRule[0].HTTP.Paths[0].Backend),
		Matches:     toMatches(ingressRule[0].HTTP.Paths[0]),
	}

	rules = append(rules, *rule)
	return rules
}

func toBackendRefs(backend networkingv1.IngressBackend) []gatewayv1.HTTPBackendRef {
	name := backend.Service.Name
	// namespace := backend.Service.Namespace
	port := backend.Service.Port.Number
	weight := int32(100)
	kind := gatewayv1.Kind("Service")

	backendRef := &gatewayv1.HTTPBackendRef{
		BackendRef: gatewayv1.BackendRef{
			BackendObjectReference: gatewayv1.BackendObjectReference{
				Kind: &kind,
				Name: gatewayv1.ObjectName(name),
				Port: (*gatewayv1.PortNumber)(&port),
			},
			Weight: &weight,
		},
	}

	backendRefs := make([]gatewayv1.HTTPBackendRef, 0)
	backendRefs = append(backendRefs, *backendRef)

	return backendRefs
}

// toMatches
func toMatches(ingressRule networkingv1.HTTPIngressPath) []gatewayv1.HTTPRouteMatch {
	matches := make([]gatewayv1.HTTPRouteMatch, 0)

	path := ingressRule.Path
	pathMatchType := ingressRule.PathType

	matchRule := gatewayv1.HTTPRouteMatch{
		Path: matchPath(path, pathMatchType),
	}
	matches = append(matches, matchRule)

	return matches
}

func matchPath(path string, pathMatchType *networkingv1.PathType) *gatewayv1.HTTPPathMatch {
	// Ingress pathType must be translated to https://pkg.go.dev/sigs.k8s.io/gateway-api@v1.3.0/apis/v1#PathMatchType
	// The value ImplementationSpecific is not valid in Gateway API.
	matchRule := &gatewayv1.HTTPPathMatch{
		Type:  (*gatewayv1.PathMatchType)(pathMatchType),
		Value: &path,
	}
	return matchRule
}
