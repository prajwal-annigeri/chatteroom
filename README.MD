# Chatteroom

A chat room application where users can create and join groups

# Getting Started
## Requirements
* `Go` >= 1.21.4

## Fetch the application
`go get github.com/prajwal-annigeri/chatteroom`

# Usage
## Starting the backend
* Set environment variable 'secretkey' to be used for hashing 
`$ go run server/cmd/main.go --conn-string="{postgres_connection_string}"`

## Starting the frontend
* Navigate to the client directory
* `$ npm run dev`

You can access the application at `http://localhost:8080`