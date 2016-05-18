package main

import (
	"log"

	"github.com/y-okubo/gogfapi/gfapi"
)

func main() {
	vol := new(gfapi.Volume)
	if vol == nil {
		log.Fatal("Can't create volume")
	}

	ret := vol.Init("server", "gfs_volume")
	if ret != 0 {
		log.Fatal("Failed to initialize volume")
	}

	vol.Mount()

	_, err := vol.Create("create_by_go.txt")
	if err != nil {
		log.Fatal("Failed to create file")
	}

	err = vol.Mkdir("/TestRmdir", 0755)
	if err != nil {
		log.Printf("Failed to create file. Error = %v", err)
	}

	ffd, err := vol.Open("/TestRmdir")
	if err != nil {
		log.Fatal("Failed to open file")
	}

	entries, err := ffd.Readdir(100)
	if err != nil {
		log.Fatal("Failed to read directory")
	}

	for _, entry := range entries {
		log.Printf("%s\t%d\t%s\t%s", entry.Name(), entry.Size(), entry.IsDir(), entry.ModTime())
		log.Print(entry.Sys())
	}

	// log.Println(f)
}
