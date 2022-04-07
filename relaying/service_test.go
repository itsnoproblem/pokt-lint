package relaying_test

import (
	"context"
	"github.com/itsnoproblem/pokt-lint/mock"
	"github.com/itsnoproblem/pokt-lint/pocket"
	"github.com/itsnoproblem/pokt-lint/relaying"
	"testing"
)

func TestNodeChecker_RunRelayTests(t *testing.T) {
	nodeId := "123abc"
	nodeAddress := "https://node-000.mynode.com"
	chains := []string{"0001"}

	client := mock.NewFakeHTTPClient(true)
	prv := pocket.NewProvider("https://1.2.3.4:443", client)
	svc, err := relaying.NewNodeChecker(nodeId, nodeAddress, chains, prv)
	if err != nil {
		t.Fatalf("got error instantiating node checker: %s", err)
	}

	_, err = svc.RunRelayTests(context.Background(), 5)
	if err != nil {
		t.Fatalf("got error running relay tests: %s", err)
	}
}
