package runner

import (
	"fmt"
	"log"
	"os"
	"service/internal/app/grpc"
	"service/internal/pkg/config"

	xtremegrpc "github.com/globalxtreme/go-core/v2/grpc"
	xtremepkg "github.com/globalxtreme/go-core/v2/pkg"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/reflection"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:  "xtreme:grpc",
		Long: "Running gRPC",
		Run: func(cmd *cobra.Command, args []string) {
			xtremepkg.InitDevMode()

			config.InitTZ()

			DBClose := config.InitDB()
			defer DBClose()

			logCleanup := xtremepkg.InitLogRPC()
			defer logCleanup()

			addr := fmt.Sprintf("%s", os.Getenv("GRPC_DEV_TEST_HOST"))

			server := xtremegrpc.GRPCServer{}
			server.NewServer(addr)

			grpc.Register(&server)
			reflection.Register(server.Server)

			fmt.Println(fmt.Sprintf("gRPC server is running: %s", addr))
			if err := server.Serve(); err != nil {
				log.Fatalf("Failed to server: %v", err)
			}
		},
	})
}
