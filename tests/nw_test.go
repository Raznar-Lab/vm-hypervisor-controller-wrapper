package tests

import (
	"github.com/Raznar-Lab/vm-hypervisor-controller-wrapper/services/nw"
	"testing"
)

type NWTest struct {
	BaseTest
}

func (b NWTest) Start() (err error) {
	nwService := b.wrapper.NW()

	// Reset the network
	if err = b.resetNetwork(nwService); err != nil {
		return err
	}

	// Create a network rule
	if err = b.createNetworkRule(nwService); err != nil {
		return err
	}

	// Delete a network rule
	if err = b.deleteNetworkRule(nwService); err != nil {
		return err
	}

	// Create multiple network rules
	if err = b.createMultipleNetworkRules(nwService); err != nil {
		return err
	}

	// Reset the network again
	if err = b.resetNetwork(nwService); err != nil {
		return err
	}

	return nil
}

func (b NWTest) resetNetwork(nwService *nw.NetworkService) (err error) {
	b.t.Log("Resetting the network")
	success, err := nwService.Reset()
	if err != nil || !success {
		b.t.Logf("Network reset failed: %v", err)
		return err
	}
	b.t.Log("Network reset successfully")
	return nil
}

func (b NWTest) createNetworkRule(nwService *nw.NetworkService) (err error) {
	b.t.Log("Creating a network rule")
	success, err := nwService.Add("139.99.4.25/32", "02:00:00:be:c4:6c")
	if err != nil || !success {
		b.t.Logf("Failed to create network rule: %v", err)
		return err
	}
	b.t.Log("Network rule created successfully")
	return nil
}

func (b NWTest) deleteNetworkRule(nwService *nw.NetworkService) (err error) {
	b.t.Log("Deleting network rule")
	success, err := nwService.Delete("139.99.4.25", "02:00:00:be:c4:6c")
	if err != nil || !success {
		b.t.Logf("Failed to delete network rule: %v", err)
		return err
	}
	b.t.Log("Network rule deleted successfully")
	return nil
}

func (b NWTest) createMultipleNetworkRules(nwService *nw.NetworkService) (err error) {
	b.t.Log("Creating multiple network rules")
	success, err := nwService.CreateMultiple([]string{"139.99.4.25/32", "139.99.4.21/32", "139.99.4.23/32"}, "02:00:00:be:c4:6c")
	if err != nil || !success {
		b.t.Logf("Failed to create multiple network rules: %v", err)
		return err
	}
	b.t.Log("Multiple network rules created successfully")
	return nil
}

func TestNW(t *testing.T) {
	wrapper, err := loadEnv()
	if err != nil {
		t.Fatalf("cannot load wrapper!: %s", err.Error())
		return
	}

	nwTest := NWTest{}
	nwTest.wrapper = wrapper
	nwTest.t = t

	// Start the test
	err = nwTest.Start()
	if err != nil {
		panic(err)
	}
}
