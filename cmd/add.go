package cmd

import (
	"context"
	"fmt"
	"github.com/KhovalygTaraa/music_service/api"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add new song to playlist",
	Long: ``,
	Run: add,
}

var songName string
var songAuthor string
var duration int

func add(cmd *cobra.Command, args []string) {
	host, port := getHostPort()
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer conn.Close()
	client := api.NewMusicServiceClient(conn)
	song := new(api.Song)
	song.Author = songAuthor
	song.Name = songName
	song.Duration = int64(duration)
	response, err := client.AddSong(context.Background(), song)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(response.Response)
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVar(&songName, "songName", "unknown", "")
	addCmd.Flags().StringVar(&songAuthor, "songAuthor", "unknown", "")
	addCmd.Flags().IntVar(&duration, "songDuration", 30, "")
}
