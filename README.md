# music_service_client

Клиент реализован на фреймворке cobra-cli.

Использование:
1. go run main.go play
2. go run main.go pause
3. go run main.go next
4. go run main.go prev
5. go run main.go add --songName "Hello" --authorName "Adele" --duration 31
6. go run main.go interactive(включает интерактивный режим, при котором можно вводить все комманды)
7. play
8. pause
9. play
10. add "Hello" "Adele" "31"

    rpc DeleteSong(Song) returns(Response);
    rpc GetPlaylist(Empty) returns(Playlist);
    rpc GetSong(Song) returns(Song);
    rpc UpdateSong(Song) returns(Response);
