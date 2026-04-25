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

func TestGlobalUpgradeSuccessUsesSuccessUnion(t *testing.T) {
	var resp opencode.GlobalUpgradeResponse
	err := json.Unmarshal([]byte(`{"success":true,"version":"1.0.0"}`), &resp)
	if err != nil {
		t.Fatal(err)
	}

	success, ok := resp.AsUnion().(opencode.GlobalUpgradeResponseSuccess)
	if !ok {
		t.Fatalf("unexpected union type: %#v", resp.AsUnion())
	}
	if success.Version != "1.0.0" {
		t.Fatalf("unexpected version: %q", success.Version)
	}
}

func TestGlobalUpgradeRejectsMalformedFalseBranch(t *testing.T) {
	var resp opencode.GlobalUpgradeResponse
	err := json.Unmarshal([]byte(`{"success":false,"version":"1.0.0"}`), &resp)
	if err == nil {
		t.Fatal("expected malformed failure branch to be rejected")
	}
}

func TestGlobalUpgradeRejectsMissingSuccess(t *testing.T) {
	var resp opencode.GlobalUpgradeResponse
	err := json.Unmarshal([]byte(`{"error":"boom"}`), &resp)
	if err == nil {
		t.Fatal("expected missing success to be rejected")
	}
}

func TestGlobalUpgradeRejectsMalformedTrueBranch(t *testing.T) {
	var resp opencode.GlobalUpgradeResponse
	err := json.Unmarshal([]byte(`{"success":true,"error":"boom"}`), &resp)
	if err == nil {
		t.Fatal("expected malformed success branch to be rejected")
	}
}

func TestGlobalUpgradeRejectsNonBooleanSuccess(t *testing.T) {
	var resp opencode.GlobalUpgradeResponse
	err := json.Unmarshal([]byte(`{"success":"true","version":"1.0.0"}`), &resp)
	if err == nil {
		t.Fatal("expected non-boolean success to be rejected")
	}
}
