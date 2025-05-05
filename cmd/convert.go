package cmd

import (
	"fmt"
	"log"
	"os"

	"k8s.io/cli-runtime/pkg/printers"

	"github.com/giantswarm/ga-migration-cli/internal/httproute"
	"github.com/giantswarm/ga-migration-cli/internal/ingress"
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
}

func runConvert(cmd *cobra.Command, args []string) {
	filename, _ := cmd.Flags().GetString("filename")
	source, err := ingress.NewFromFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	httproute := httproute.New().WithIngress(source)

	resourcePrinter := &printers.YAMLPrinter{}

	// TODO: Clean resource of unwanted fields like: Status, ResourceVersion, CreationTimestamp, etc.
	err = resourcePrinter.PrintObj(httproute.Resource, os.Stdout)
	if err != nil {
		fmt.Printf("# Error printing %s HTTPRoute: %v\n", httproute.Resource.Name, err)
	}
}
