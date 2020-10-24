package rules

import (
	"clinne/internal/pkg/printer"
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {

	return &cobra.Command{
		Use:     "rules",
		Short:   "NNE Rules",
		Long:    "This command is used to show all the rules of NNE that you need to take care while coding.",
		Example: "clinne rules",
		Args:    cobra.MaximumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			printer.Println("The NNE you need to take care are:-", color.Underline, color.FgHiGreen)
			fmt.Println()
			printer.Println("1. Indentation and spacing between code constructs (classes/methods/specs) must be consistent")
			fmt.Println()
			printer.Println("2. Use only spaces (no tabs) for indentation")
			fmt.Println()
			printer.Println("3. Newlines at end of file")
			fmt.Println()
			printer.Println("4. Follow accepted naming conventions for your language/framework")
			fmt.Println()
			printer.Println("5. Follow accepted naming file and Directory structure for your language/framework")
			fmt.Println()
			printer.Println("6. Use namespaces")
			fmt.Println()
			printer.Println("7. No comments/Unused Code must ever be checked in")
			fmt.Println()
			printer.Println("8. Runtime environment should be consistent with IDE environment - i.e there should be no difference in running a build or a spec from your IDE and from the command line")
			fmt.Println()
			printer.Println("9. Use .gitignore")
			fmt.Println()
			printer.Println("10. Ensure there is a README.md that includes: Problem Description, Dev Environment Setup, How to run test, Build Instructions, Run instructions.")
			fmt.Println()
			printer.Println("11. Test Driven Development (this should show clear pattern in the commit log: red, green, commit; refactor commit;)")
			fmt.Println()
		},
	}
}
