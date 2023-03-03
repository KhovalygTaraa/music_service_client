package cmd

import (
	"context"
	"fmt"
	"github.com/KhovalygTaraa/music_service/api"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var playCmd = &cobra.Command{
	Use:   "play",
	Short: "A brief description of your command",
	Long: ``,
	Run: play,
}

var _config string

func play(cmd *cobra.Command, args []string) {
		fmt.Println("play called")
		conn, err := grpc.Dial("0.0.0.0:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		client := api.NewMusicServiceClient(conn)
		_, err = client.Play(context.Background(), &api.Empty{})
		if err != nil {
			panic(err)
		}
}

func init() {
	rootCmd.AddCommand(playCmd)
	playCmd.Flags().StringVar(&_config, "config", "./config.yaml", "")
}
