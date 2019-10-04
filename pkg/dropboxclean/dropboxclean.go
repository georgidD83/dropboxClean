package dropboxclean

import (
	"encoding/json"
	"log"

	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
)

// Metadata contains all metadata for the Dropbox account file system.
type Metadata struct {
	Entries []struct {
		Tag         string `json:".tag"`
		Name        string `json:"name"`
		PathLower   string `json:"path_lower"`
		PathDisplay string `json:"path_display"`
		ID          string `json:"id"`
	} `json:"entries"`
	Cursor  string `json:"cursor"`
	HasMore bool   `json:"has_more"`
}

// GetMetadata returns all metadata for the Dropbox account file system.
func GetMetadata(client files.Client) *Metadata {
	root := &files.ListFolderArg{
		Path:      "/Apps/Netatmo",
		Recursive: true,
	}
	resp, err := dbx.ListFolder(root)
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
}
