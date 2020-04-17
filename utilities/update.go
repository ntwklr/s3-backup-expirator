package utilities

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/go-github/github"
	"github.com/ntwklr/s3-backup-expirator/error"
)

func Update(projectOwner, projectRepo, binaryPlatform, version string) {
	ctx := context.Background()
	client := github.NewClient(nil)

	latestRelease, _, err := client.Repositories.GetLatestRelease(ctx, projectOwner, projectRepo)

	if err != nil {
		error.Exitf("Error while fetching releases: %s/%s, %v", projectOwner, projectRepo, err)
	}

	if *latestRelease.TagName <= "v"+version {
		fmt.Printf("You are already using %s version %s\n", projectRepo, version)
		os.Exit(0)
	}

	download := ""

	for _, asset := range latestRelease.Assets {
		if strings.Contains(*asset.Name, binaryPlatform) {
			download = *asset.BrowserDownloadURL
		}
	}

	if download == "" {
		error.Exitf("Download failed. Binary not found.")
	}

	fileOriginal := GetExec()
	fileOriginalObj, _ := os.Stat(fileOriginal)
	fileOriginalMode := fileOriginalObj.Mode()
	fileBak := fileOriginal + ".bak"
	fileNew := fileOriginal + ".new"

	fmt.Printf("Download %v\n", *latestRelease.Name)

	DownloadFile(download, fileNew)

	os.Rename(fileOriginal, fileBak)
	os.Rename(fileNew, fileOriginal)
	os.Chmod(fileOriginal, fileOriginalMode)

	os.Exit(0)
}

// DownloadFile will download a url and store it in local filepath.
// It writes to the destination file as it downloads it, without
// loading the entire file into memory.
func DownloadFile(url string, filepath string) {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		error.Exitf("Error opening file: %s, %v", filepath, err)
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		error.Exitf("Error while downloading file: %s, %v", url, err)
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		error.Exitf("Error while writing file: %s, %v", filepath, err)
	}
}

func GetExec() string {
	ex, _ := os.Executable()
	fi, _ := os.Lstat(ex)

	if fi.Mode()&os.ModeSymlink == os.ModeSymlink {
		link, _ := filepath.EvalSymlinks(ex)

		return link
	} else {
		return ex
	}
}
