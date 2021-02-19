package query

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"

	"github.com/Pylons-tech/pylons/x/pylons/types"
	"github.com/spf13/cobra"
)

// GetExecution get an execution by GUID
func GetExecution(queryRoute string, ) *cobra.Command {
	ccb := &cobra.Command{
		Use:   "get_execution <id>",
		Short: "get an execution by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := client.GetClientContextFromCmd(cmd)

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/get_execution/%s", queryRoute, args[0]), nil)
			if err != nil {
				return fmt.Errorf(err.Error())
			}

			var out types.Execution
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintProto()
		},
	}
	return ccb
}
