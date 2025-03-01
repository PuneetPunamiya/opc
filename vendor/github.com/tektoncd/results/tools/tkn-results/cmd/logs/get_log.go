package logs

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	pb "github.com/tektoncd/results/proto/v1alpha2/results_go_proto"
	"github.com/tektoncd/results/tools/tkn-results/internal/flags"
	"github.com/tektoncd/results/tools/tkn-results/internal/format"
)

func GetLogCommand(params *flags.Params) *cobra.Command {
	opts := &flags.GetOptions{}

	cmd := &cobra.Command{
		Use: `get [flags] <log>

  <log path>: Log full name to query. This is typically "<namespace>/results/<result name>/logs/<log name>".`,
		Short: "Get Log",
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := params.LogsClient.GetLog(cmd.Context(), &pb.GetLogRequest{
				Name: args[0],
			})
			if err != nil {
				fmt.Printf("GetLog: %v\n", err)
				return err
			}
			data, err := resp.Recv()
			if err != nil {
				fmt.Printf("Get Log Client Resp: %v\n", err)
				return err
			}
			return format.PrintProto(os.Stdout, data, opts.Format)
		},
		Args: cobra.ExactArgs(1),
		Annotations: map[string]string{
			"commandType": "main",
		},
	}

	flags.AddGetFlags(opts, cmd)

	return cmd
}
