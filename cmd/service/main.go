package main

import (
    "encoding/json"
    "io/ioutil"

    "github.com/Gorynychdo/bstreee/internal/bstree"
    "github.com/Gorynychdo/bstreee/internal/service"
    log "github.com/sirupsen/logrus"
)

func init() {
    log.SetFormatter(&log.JSONFormatter{})
    log.SetLevel(log.DebugLevel)
}

func main() {
    data, err := ioutil.ReadFile("data/data.json")
    if err != nil {
        log.Fatal(err)
    }

    var slice []int
    if err = json.Unmarshal(data, &slice); err != nil{
        log.Fatal(err)
    }

    tree := bstree.NewTree(slice)
    srv := service.NewServer(tree)
    if err = srv.Serve(":8080"); err != nil {
        log.Fatal(err)
    }
}
