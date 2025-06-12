package cmd

import (
	"fmt"
	"log"
	"os"

	"k8s.io/cli-runtime/pkg/printers"

	"github.com/giantswarm/ga-migration-cli/internal/httproute"
	"github.com/giantswarm/ga-migration-cli/internal/ingress"
	"github.com/giantswarm/ga-migration-cli/internal/securitypolicy"
	"github.com/spf13/cobra"
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert Ingress resource into HTTPRoute",
	Run:   runConvert,
}

func init() {
	rootCmd.AddCommand(convertCmd)

	convertCmd.Flags().StringP("filename", "f", "", "Filename containing the ingress resource")
	convertCmd.Flags().StringP("gateway", "", "giantswarm-default", "Name of the Gateway")
	convertCmd.Flags().StringP("gateway-namespace", "", "envoy-gateway-system", "Namespace of the Gateway")
}

func runConvert(cmd *cobra.Command, args []string) {
	resourcePrinter := &printers.YAMLPrinter{}

	filename, _ := cmd.Flags().GetString("filename")
	source, err := ingress.NewFromFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	gatewayName, _ := cmd.Flags().GetString("gateway")
	gatewayNamespace, _ := cmd.Flags().GetString("gateway-namespace")

	httproute := httproute.New().WithIngress(source).WithGateway(gatewayName, gatewayNamespace, "https")

	// TODO: Clean resource of unwanted fields like: Status, ResourceVersion, CreationTimestamp, etc.
	err = resourcePrinter.PrintObj(httproute.Resource, os.Stdout)
	if err != nil {
		fmt.Printf("# Error printing %s HTTPRoute: %v\n", httproute.Resource.Name, err)
	}

	securityPolicy := securitypolicy.New().WithHTTPRoute(httproute.Resource).WithIngress(source)

	if securityPolicy != nil {
		err = resourcePrinter.PrintObj(securityPolicy.Resource, os.Stdout)
		if err != nil {
			fmt.Printf("# Error printing %s HTTPRoute: %v\n", httproute.Resource.Name, err)
		}
	}
}
