package api

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/buonotti/apisense/filesystem/locations/directories"
	"github.com/go-resty/resty/v2"
)

type GHListing struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Url         string `json:"url"`
	DownloadUrl string `json:"download_url"`
}

type GHListingResponse struct {
	Entries []GHListing `json:"entries"`
}

func formatResponse(resp string) string {
	return fmt.Sprintf("{\"entries\": %s}", resp)
}

var client = resty.New()

func cd(base string, dir string) string {
	return base + "/" + dir
}

func InstallUI() error {
	baseDir := directories.UiDirectory()
	baseUrl := "https://api.github.com/repos/buonotti/apisense/contents/ui"

	listing, err := downloadListings(baseUrl)
	if err != nil {
		return err
	}

	for _, entry := range listing.Entries {
		err := writeToDisk(entry, baseDir, baseUrl)
		if err != nil {
			return err
		}
	}

	return nil
}

func downloadRaw(url string) (string, error) {
	resp, err := client.R().Get(url)
	if err != nil {
		return "", err
	}
	return string(resp.Body()), nil
}

func downloadListings(url string) (GHListingResponse, error) {
	resp, err := client.R().Get(url + "?ref=dev")
	if err != nil {
		return GHListingResponse{}, err
	}

	jsonStr := formatResponse(string(resp.Body()))

	var listing GHListingResponse
	err = json.Unmarshal([]byte(jsonStr), &listing)
	if err != nil {
		return GHListingResponse{}, err
	}
	return listing, nil
}

func writeToDisk(listing GHListing, workdir string, fetchdir string) error {
	if listing.Type == "file" {
		content, err := downloadRaw(listing.DownloadUrl)
		if err != nil {
			return err
		}
		err = os.WriteFile(workdir+"/"+listing.Name, []byte(content), os.ModePerm)
		if err != nil {
			return err
		}
	} else {
		err := os.Mkdir(listing.Name, os.ModePerm)
		if err != nil {
			return err
		}
		children, err := downloadListings(fetchdir + "/" + listing.Name)
		if err != nil {
			return err
		}

		for _, entry := range children.Entries {
			err := writeToDisk(entry, cd(workdir, listing.Name), cd(fetchdir, listing.Name))
			if err != nil {
				return err
			}
		}

	}
	return nil
}
