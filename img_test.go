package main

import (
	"testing"
)

type IMGTest struct {
	BaseTest
}

func (b IMGTest) Start() (err error) {
	b.t.Log("downloading image..")
	// Download Image
	success, err := b.wrapper.IMG().Download("https://gist.githubusercontent.com/NotYusta/2254074e61513cc36fd27b9a3624dd99/raw/53dbc5d471bf3558508affec75e47b0062d8b153/proxmox_automatic_boot_order.sh", "test.img")
	if !success || err != nil {
		b.t.Log("failed to download image test.img")
		return
	}

	b.t.Log("downloaded image test.img")
	success, err = b.wrapper.IMG().Delete("test.img")
	if !success || err != nil {
		b.t.Log("failed to delete image test.img")
		return
	}

	b.t.Log("successfully deleted image test.img")
	return
	// Delete Image
}

func TestIMG(t *testing.T) {
	wrapper := New("http://15.235.160.20:8080", "cikoHNJOA2", "Caiktjhoi@tha@ithai@tha@izz")
	imgTest := IMGTest{}
	imgTest.wrapper = wrapper
	imgTest.t = t

	err := imgTest.Start()
	if err != nil {
		panic(err)
	}
}
