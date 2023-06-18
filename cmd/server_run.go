package cmd

// import (
// 	"github.com/spf13/cobra"
// 	"github.com/u-next/papy/app"
// 	"github.com/u-next/papy/app/api"
// 	"github.com/u-next/papy/app/infra/repo"
// 	"golang.org/x/xerrors"
// )

// newServerRunCmd represents the serverPapy command
// func newServerRunCmd() *cobra.Command {
// 	serverPapyCmd := &cobra.Command{
// 		Use:   "run",
// 		Short: "run server commmand",
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			db := infrastructure.NewDB()
// 			app.Init(db)
// 			return nil
// 		},
// 	}

// 	return serverPapyCmd
// }