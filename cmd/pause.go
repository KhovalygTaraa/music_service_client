package cmd

import (
	"context"
	"fmt"
	"github.com/KhovalygTaraa/music_service/api"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var pauseCmd = &cobra.Command{
	Use:   "pause",
	Short: "A brief description of your command",
	Long: ``,
	Run: pause,
}

func pause(cmd *cobra.Command, args []string) {
	fmt.Println("pause called")
	conn, err := grpc.Dial("0.0.0.0:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := api.NewMusicServiceClient(conn)
	_, err = client.Pause(context.Background(), &api.Empty{})
	if err != nil {
		panic(err)
	}
}

func init() {
	rootCmd.AddCommand(pauseCmd)
}
