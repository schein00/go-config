package main



import (
    "go-conrfig/src/app/configuration"
)

func main() {
    configuration := &configuration.Configuration{}
    configuration.Start()
}

