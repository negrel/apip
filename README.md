# APIP - Free public IP API

*Get your public IP address in one **HTTP** request.*


## What is APIP

**APIP** is an **open source** public IP API that focuses on simplicity. It can be really useful if you're looking for a way to get your public IP from **command line**, **script** or **programs**. There is a public server for test purpose : [demo-apip.herokuapp.com](https://demo-apip.herokuapp.com/) hosted on a **free** dynos at [heroku.com](https://heroku.com). If your looking for a more reliable public server for public IP API I recommend you [**ipify.org**](https://ipify.org).


[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

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

### Using the Makefile

_Take a look at the exemple.env file and write your own .env file._

Build the image:

```
$ make build
# Or without caching
$ make build-nc
```

Run it:

```
$ make run
# Or build it then run it
$ make up
```

To stop & remove the container :
```
$ make stop
```

Compile go to binary :

```
# Build the binary to $BIN_OUTPUT
$ make bin
```


Clean the binary :

```
$ make clean
```

Deploy to heroku :

```
$ make heroku
```

## :stars: Show your support

Please give a :star: if this project helped you!

## :scroll: License
MIT © Alexandre Negrel

