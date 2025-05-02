package cmd

import (
	"fmt"
	"log"

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
	i, err := ingress.NewFromFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(i)
	// create a new converter
	// and convert resource

	// output the new resource
}
