package main

import "github.com/Barrioslopezfd/gator/internal/config"

func main(){
    c:=config.Read()
    c.SetUser("Franco")
    config.Read()
}
