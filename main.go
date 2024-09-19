package main

import (
	"fmt"
	"net/http"

	"github.com/sonlis/technapi/internal/api"
)


type techniClient struct {
    Url string
    sessionToken string
    c http.Client
}

func main() {
    client := api.TechniClient{
        Url: "http://192.168.0.13:5380",
    } 
    client.GetSessionToken("", "")
    zones, err := client.ListZones()
    if err != nil {
        fmt.Println("Error: ", err)
    }
    fmt.Println(zones)
}
