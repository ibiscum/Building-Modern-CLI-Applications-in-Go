//go:build !int

package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/ibiscum/Building-Modern-CLI-Applications-in-Go/Chapter13/audiofile/models"
)

func TestGet(t *testing.T) {
	ConfigureTest()
	b := bytes.NewBufferString("")
	rootCmd.SetOut(b)
	rootCmd.SetArgs([]string{"get", "--id", "123", "--json"})
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println("err: ", err)
	}
	actualBytes, err := io.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	expectedBytes, err := os.ReadFile("./testfiles/get.json")
	if err != nil {
		t.Fatal(err)
	}
	var audio1, audio2 models.Audio
	err = json.Unmarshal(actualBytes, &audio1)
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(expectedBytes, &audio2)
	if err != nil {
		t.Fatal(err)
	}
	if !(audio1.Id == audio2.Id &&
		audio1.Metadata.Tags.Album == audio2.Metadata.Tags.Album &&
		audio1.Metadata.Tags.AlbumArtist == audio2.Metadata.Tags.AlbumArtist &&
		audio1.Metadata.Tags.Artist == audio2.Metadata.Tags.Artist &&
		audio1.Metadata.Tags.Comment == audio2.Metadata.Tags.Comment &&
		audio1.Metadata.Tags.Composer == audio2.Metadata.Tags.Composer &&
		audio1.Metadata.Tags.Genre == audio2.Metadata.Tags.Genre &&
		audio1.Metadata.Tags.Lyrics == audio2.Metadata.Tags.Lyrics &&
		audio1.Metadata.Tags.Year == audio2.Metadata.Tags.Year) {
		t.Fatalf("expected %q got %q", string(expectedBytes), string(actualBytes))
	}
}
