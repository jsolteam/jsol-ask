# JSOL Ask


<img align="right" src="https://i.ibb.co/fHHNWjL/jsol-team-white.png" height="30px" alt="JSOL Team Logo"/>

[//]: # (![Golang Badge]&#40;https://img.shields.io/badge/Go-1.23-blue&#41;)

[//]: # (![GitHub License]&#40;https://img.shields.io/github/license/jsolteam/slm-bot-publisher&#41;)

[//]: # (![GitHub Release]&#40;https://img.shields.io/github/v/release/jsolteam/slm-bot-publisher&#41;)


<div align="center">
  <img src="https://i.ibb.co/vhkDKDg/jsol-ask.png" alt="Project Logo" height="140px"/>
   <p align="center">Бот для принятия обратной связи от пользователей</p>
</div>

<p align="center">
  <a href="https://jsol-team.ru">JSOL Team</a> •
  <a href="https://t.me/jsol_team">Telegram</a> •
  <a href="https://github.com/jsolteam">Github</a> •
</p>

<br/>

### Запуск

1. Установите Go, следуя официальной инструкции: [https://go.dev/doc/install](https://go.dev/doc/install)
2. Убедитесь, что Go установлен корректно:
   ```sh
   go version
   ```
3. Клонируйте репозиторий
   ```sh
   git clone https://github.com/jsolteam/jsol-ask.git
   ```
4. Перейдите в папку проекта
   ```sh
   cd jsol-ask
   ```
5. Установите зависимости из файла go.mod
   ```sh
   go mod tidy
   ```
6. Создайте .env файл, файл config, согласно информации ниже
7. Запустите проект
    ```sh
   go run cmd/main.go
   ```

### ENV файл


```shell
DATABASE_PATH=*путь к файлу sqlite*

TELEGRAM_BOT_TOKEN=*токен бота*
GROUP_ID=*ID супер-группы*
GENERAL_THREAD=*ID главной темы*
SUPER_ADMIN_ID=*ID главного администратора*
```

### Релизы

Все доступные релизы можно найти в разделе [Releases](https://github.com/jsolteam/jsol-ask/releases).

### Контрибуторы

<a href="https://github.com/jsolteam/jsol-ask/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=jsolteam/jsol-ask" alt="contrib.rocks image" />
</a>

### Лицензия

Проект представлен MIT лицензией. Смотрите `LICENSE` для подробной информации.
