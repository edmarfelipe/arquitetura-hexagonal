package cmd

import (
	"fmt"
	"os"

	"github.com/edmarfelipe/go-hexagonal/adapters/cli"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-hexagonal",
	Short: "A brief description of your application",
}

var service = cli.CreateProductService()

func Create(cmd *cobra.Command, args []string) {
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		fmt.Println(err)
	}

	price, err := cmd.Flags().GetFloat64("price")
	if err != nil {
		fmt.Println(err)
	}

	result, err := cli.Create(service, name, price)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}

func Enable(cmd *cobra.Command, args []string) {
	result, err := cli.Enable(service, args[0])
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}

func Disable(cmd *cobra.Command, args []string) {
	result, err := cli.Disable(service, args[0])
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}

func Get(cmd *cobra.Command, args []string) {
	result, err := cli.Get(service, args[0])
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}

func init() {
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new product",
		Run:   Create,
	}

	createCmd.Flags().StringP("name", "n", "", "Product Name")
	createCmd.Flags().Float64P("price", "p", 0, "Product Price")

	rootCmd.AddCommand(createCmd)

	rootCmd.AddCommand(&cobra.Command{
		Use:   "get",
		Short: "Get a product by ID",
		Args:  cobra.MinimumNArgs(1),
		Run:   Get,
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "disable",
		Short: "Disable a product",
		Args:  cobra.MinimumNArgs(1),
		Run:   Disable,
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "enable",
		Short: "Enable a product",
		Args:  cobra.MinimumNArgs(1),
		Run:   Enable,
	})
}

func ExecuteCli() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
