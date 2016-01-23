package main

import (
	"os"
	"reflect"
	"testing"

	mp "github.com/gophergala2016/magopie"
)

func TestKATParse(t *testing.T) {
	data, err := os.Open("testdata/kat_ubuntu.xml")
	if err != nil {
		t.Fatal("Error opening fixture", err)
	}
	defer data.Close()

	actual, err := katParse(data)

	expected := []mp.Torrent{
		{
			ID:       "0C5B427C2F833B09EA0E3DC7C624F6C187125267",
			Title:    "Ubuntu Linux Go from Beginner to Power User!",
			FileURL:  "https://torcache.net/torrent/0C5B427C2F833B09EA0E3DC7C624F6C187125267.torrent?title=[kat.cr]ubuntu.linux.go.from.beginner.to.power.user",
			SiteID:   "kat",
			Size:     1137244499,
			Seeders:  85,
			Leechers: 105,
		},
		{
			ID:       "21236117B7A773639BD5C7C771E66A045BD51A8A",
			Title:    "Learning Ubuntu Linux Server",
			FileURL:  "https://torcache.net/torrent/21236117B7A773639BD5C7C771E66A045BD51A8A.torrent?title=[kat.cr]learning.ubuntu.linux.server",
			SiteID:   "kat",
			Size:     581653341,
			Seeders:  48,
			Leechers: 12,
		},
		{
			ID:       "13DBA979D53F20E6A73D4EE939952D1C367B64C7",
			Title:    "CodeWeavers Crossover 15.0.1 with crack for ubuntu fedora linux",
			FileURL:  "https://torcache.net/torrent/13DBA979D53F20E6A73D4EE939952D1C367B64C7.torrent?title=[kat.cr]codeweavers.crossover.15.0.1.with.crack.for.ubuntu.fedora.linux",
			SiteID:   "kat",
			Size:     240330669,
			Seeders:  1,
			Leechers: 1,
		},
		{
			ID:       "EB7B040141407F150E32FF366CD624403387B5C1",
			Title:    "Ubuntu 16.04 (Xenial Xerus) Alpha Desktop AMD64 (64-bit PC)",
			FileURL:  "https://torcache.net/torrent/EB7B040141407F150E32FF366CD624403387B5C1.torrent?title=[kat.cr]ubuntu.16.04.xenial.xerus.alpha.desktop.amd64.64.bit.pc",
			SiteID:   "kat",
			Size:     1480048640,
			Seeders:  1,
			Leechers: 3,
		},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("tpbParse actual = %v\nexpected %v", actual, expected)
	}
}
