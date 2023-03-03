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
	Short: "A brief description of your command",
	Long: ``,
	Run: add,
}

var songName string
var songAuthor string
var duration int

func add(cmd *cobra.Command, args []string) {
	fmt.Println("add called")
	conn, err := grpc.Dial("0.0.0.0:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := api.NewMusicServiceClient(conn)
	song := new(api.Song)
	song.Author = songAuthor
	song.Name = songName
	song.Duration = int64(duration)
	_, err = client.AddSong(context.Background(), song)
	if err != nil {
		panic(err)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVar(&songName, "song name", "unknown", "")
	addCmd.Flags().StringVar(&songAuthor, "song author", "unknown", "")
	addCmd.Flags().IntVar(&duration, "song duration", 30, "")
}
