package ingress

import (
	"os"

	"sigs.k8s.io/yaml"

	networkingv1 "k8s.io/api/networking/v1"
)

func NewFromFile(filename string) (*networkingv1.Ingress, error) {
	var ingress networkingv1.Ingress

	yfile, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yfile, &ingress)
	if err != nil {
		return nil, err
	}

	return &ingress, nil
}
