# music_service_client

## Клиент реализован на фреймворке cobra-cli.


### Команда сборки:
make

### Список доступных комманд:
./music_service -help

### Просмотр необходимых опций комманд:
./music_servise add -help

### Интерактивный режим:
./music_servise interactive

### Примеры:
./music_servise play \
./music_servise add --songName "Hello" --authorName "Adele" --duration 31 

### Интерактивный режим пример:
./music_servise interactive \
play \
add wrong \
add "Hello" "Adele" 90 \
