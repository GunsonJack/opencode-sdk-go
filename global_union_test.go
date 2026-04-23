package opencode_test

import (
	"encoding/json"
	"testing"

	"github.com/GunsonJack/opencode-sdk-go"
)

func TestGlobalUpgradeFailureUsesFailureUnion(t *testing.T) {
	var resp opencode.GlobalUpgradeResponse
	err := json.Unmarshal([]byte(`{"success":false,"error":"boom"}`), &resp)
	if err != nil {
		t.Fatal(err)
	}

	failure, ok := resp.AsUnion().(opencode.GlobalUpgradeResponseFailure)
	if !ok {
		t.Fatalf("unexpected union type: %#v", resp.AsUnion())
	}
	if failure.Error != "boom" {
		t.Fatalf("unexpected error: %q", failure.Error)
	}
}

func TestGlobalUpgradeRejectsMalformedFalseBranch(t *testing.T) {
	var resp opencode.GlobalUpgradeResponse
	err := json.Unmarshal([]byte(`{"success":false,"version":"1.0.0"}`), &resp)
	if err == nil {
		t.Fatal("expected malformed failure branch to be rejected")
	}
}
