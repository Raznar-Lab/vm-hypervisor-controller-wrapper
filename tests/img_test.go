package tests

import (
	"testing"
)

type IMGTest struct {
	BaseTest
}

func (b IMGTest) Start() error {
	if err := b.DownloadImage("https://gist.githubusercontent.com/NotYusta/2254074e61513cc36fd27b9a3624dd99/raw/53dbc5d471bf3558508affec75e47b0062d8b153/proxmox_automatic_boot_order.sh", "test.qcow2"); err != nil {
		return err
	}

	if err := b.DownloadImage("https://gist.githubusercontent.com/NotYusta/2254074e61513cc36fd27b9a3624dd99/raw/53dbc5d471bf3558508affec75e47b0062d8b153/proxmox_automatic_boot_order.sh", "test-2.qcow2"); err != nil {
		return err
	}

	if err := b.DeleteImage("test.qcow2"); err != nil {
		return err
	}

	if err := b.FetchAllImages(); err != nil {
		return err
	}

	// Uncomment with caution if clearing images is required
	// if err := b.ClearImages(); err != nil {
	//     return err
	// }
	return nil
}

func (b IMGTest) DownloadImage(url, filename string) error {
	success, err := b.wrapper.IMG().Download(url, filename)
	if err != nil || !success {
		b.t.Errorf("Failed to download image '%s': %v", filename, err)
		return err
	}
	return nil
}

func (b IMGTest) DeleteImage(filename string) error {
	success, err := b.wrapper.IMG().Delete(filename)
	if err != nil || !success {
		b.t.Errorf("Failed to delete image '%s': %v", filename, err)
		return err
	}
	return nil
}

func (b IMGTest) FetchAllImages() error {
	data, err := b.wrapper.IMG().GetAll()
	if err != nil {
		b.t.Errorf("Failed to fetch images: %v", err)
		return err
	}
	b.t.Logf("Fetched images: %s", data.Data)
	return nil
}

// Uncomment with caution if clearing all images is required
// func (b IMGTest) ClearImages() error {
// 	success, err := b.wrapper.IMG().Clear()
// 	if err != nil || !success {
// 		b.t.Errorf("Failed to clear images: %v", err)
// 		return err
// 	}
// 	return nil
// }

func TestIMG(t *testing.T) {
	// Load environment variables
	wrapper, err := loadEnv()
	if err != nil {
		t.Fatalf("cannot load wrapper!: %s", err.Error())
		return
	}

	imgTest := IMGTest{
		BaseTest: BaseTest{
			wrapper: wrapper,
			t:       t,
		},
	}

	// Start the test
	if err := imgTest.Start(); err != nil {
		t.Fatalf("TestIMG failed: %v", err)
	}
}
