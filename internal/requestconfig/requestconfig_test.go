package requestconfig

import (
	"testing"
)

func TestEncodePathSegmentEscapesReservedCharacters(t *testing.T) {
	segment, err := EncodePathSegment("provider/with space", "providerID")
	if err != nil {
		t.Fatal(err)
	}
	if segment != "provider%2Fwith%20space" {
		t.Fatalf("unexpected segment: %s", segment)
	}
}

func TestEncodePathSegmentRejectsEmptyValue(t *testing.T) {
	_, err := EncodePathSegment("", "providerID")
	if err == nil {
		t.Fatal("expected missing providerID error")
	}
}

func TestEncodePathSegmentEscapesDotDotValue(t *testing.T) {
	segment, err := EncodePathSegment("..", "providerID")
	if err != nil {
		t.Fatal(err)
	}
	if segment != "%2E%2E" {
		t.Fatalf("unexpected segment: %s", segment)
	}
}
