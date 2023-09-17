# Go Htmx Frontend

The goal of this repo is to play around with golang in combination with htmx.

## Required tools

- [Golang](https://go.dev/doc/install)
- [air](https://github.com/cosmtrek/air) Hot reloading for golang
- [nodejs](https://github.com/nvm-sh/nvm) For tailwind styles
- [CMake](https://cmake.org/install/) Run build scripts with make

## Staring the server

Start the database using

`docker-compose up`

Start tailwind in watch mode using

`make tailwind`

Start the server in watch mode with

`air`