# APIP - Free public IP API

*Get your public IP address in one **HTTP** request.*


## What is APIP

**APIP** is an **open source** public IP API that focuses on simplicity. It can be really useful if you're looking for a way to get your public IP from **command line**, **script** or **programs**. There is a public server for test purpose : [apip.negrel.dev](https://apip.negrel.dev/) hosted on a **free** dynos on [heroku.com](https://heroku.com). If your looking for a more reliable public server for public IP API I recommend you [**ipify.org**](https://ipify.org).

## Setup a local APIP server

**Requirements** :
You must have [Golang](https://golang.org/dl/) installed on your server.

Clone the repository :

```bash
$ git clone https://github.com/negrel/apip.git $GOPATH/your/path
```

Run the code for test :

```bash
$ cd apip
$ go run .
```

Build the project :

```bash
$ go build
```

This will build the go code to a binary file named "apip". To launch the server :

```bash
$ ./apip
```



## :stars: Show your support

Please give a :star: if this project helped you!

## :scroll: License
MIT © Alexandre Negrel

