package cmd

import (
	"context"
	"fmt"
	"github.com/KhovalygTaraa/music_service/api"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"bufio"
	"os"
	"strings"
	"strconv"
)

var interactiveCmd = &cobra.Command{
	Use:   "interactive",
	Short: "A brief description of your command",
	Long: ``,
	Run: interactive,
}

func interactive(cmd *cobra.Command, args []string) {
	fmt.Println("interactive called")
	var act string

	conn, err := grpc.Dial("0.0.0.0:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := api.NewMusicServiceClient(conn)

	for act != "stop" {
		scanner  := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanLines)
		scanner.Scan()
		act = scanner.Text()
		if act == "play" {
			fmt.Println("Playing")
			_, err = client.Play(context.Background(), &api.Empty{})
			if err != nil {
				panic(err)
			}
		} else if act == "pause" {
			fmt.Println("Paused")
			_, err = client.Pause(context.Background(), &api.Empty{})
			if err != nil {
				panic(err)
			}
		} else if act == "prev" {
			fmt.Println("Prev song")
			_, err = client.Prev(context.Background(), &api.Empty{})
			if err != nil {
				panic(err)
			}
		} else if act == "next" {
			fmt.Println("Next song")
			_, err = client.Next(context.Background(), &api.Empty{})
			if err != nil {
				panic(err)
			}
		} else if strings.Contains(act, "add") {
			params := strings.Split(act, " ")
			if (len(params) != 4 || params[0] != "add") {
				fmt.Println("can't add new song, wrong format. Use: add <songName> <duration>")
				continue
			}
			dur, err := strconv.Atoi(params[3])
			if err != nil {
				panic("Atoi error")
			}
			song := &api.Song{Name: params[1], Author: params[2], Duration: int64(dur)}
			_, err = client.AddSong(context.Background(), song)
			if err != nil {
				panic(err)
			}
		}
	}
}

func init() {
	rootCmd.AddCommand(interactiveCmd)
}
