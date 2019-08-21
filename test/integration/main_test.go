package integration

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/yahoo/k8s-athenz-istio-auth/test/integration/fixtures"
	"github.com/yahoo/k8s-athenz-istio-auth/test/integration/framework"
)

func TestMain(m *testing.M) {
	err := framework.Setup()
	if err != nil {
		log.Println("Error setting up test framework:", err)
		os.Exit(1)
	}

	time.Sleep(time.Second * 15)
	exitCode := m.Run()
	fixtures.CreateAthenzDomain(framework.Global.AthenzDomainClientset)
	framework.Teardown()
	os.Exit(exitCode)
}
