package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
)

var (
	token = "-e2kZ4ViybAAAAAAAAAXjLDOIlquPemHFeH2pUMnPfxoKfrJT9ttZEvE53uEz1rx"
)

func main() {
	config := dropbox.Config{
		Token: token,
	}
	dbx := files.New(config)

	dropboxclean.GetMetadata(dbx)

	arg := &files.ListFolderArg{
		Path:      "/Apps/Netatmo",
		Recursive: true,
	}
	resp, err := dbx.ListFolder(arg)
	if err != nil {
		log.Printf("%v\n", err)
	}

	b, err := json.Marshal(resp)
	if err != nil {
		log.Printf("%v\n", err)
	}
	// os.Stdout.Write(b)

	data := new(metadata)
	err = json.Unmarshal(b, &data)
	if err != nil {
		log.Printf("%v\n", err)
	}
	// fmt.Printf("%+v\n", data)

	if data.HasMore == false {
		return
	}
	arg := &files.ListFolderContinueArg{
		Cursor: data.Cursor,
	}
	resp, err := dbx.ListFolderContinue(arg)
	if err != nil {
		log.Printf("%v\n", err)
	}
	b, err := json.Marshal(resp)
	if err != nil {
		log.Printf("%v\n", err)
	}
	// os.Stdout.Write(b)
	err = json.Unmarshal(b, &data)
	if err != nil {
		log.Printf("%v\n", err)
	}
	fmt.Printf("%+v\n", data)
}
