package tourmanager

import (
	"encoding/json"
	"testing"
)

func TestDecimal_UnmarshalJSON_string(t *testing.T) {
	var d Decimal
	if err := json.Unmarshal([]byte(`"1240.00"`), &d); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if d != "1240.00" {
		t.Fatalf("Decimal = %q, want %q", d, "1240.00")
	}
}

func TestDecimal_UnmarshalJSON_bareNumber(t *testing.T) {
	var d Decimal
	if err := json.Unmarshal([]byte(`1240.00`), &d); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if d != "1240.00" {
		t.Fatalf("Decimal = %q, want %q", d, "1240.00")
	}
}

func TestDecimal_UnmarshalJSON_preservesTrailingZeros(t *testing.T) {
	// The motivating case: a float64 round-trip would normalize "10.50" to
	// "10.5", silently losing the currency's minor-unit precision.
	var d Decimal
	if err := json.Unmarshal([]byte(`"10.50"`), &d); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if d != "10.50" {
		t.Fatalf("Decimal = %q, want %q (trailing zero lost)", d, "10.50")
	}
}

func TestDecimal_MarshalJSON(t *testing.T) {
	out, err := json.Marshal(Decimal("1240.00"))
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	if string(out) != `"1240.00"` {
		t.Fatalf("Marshal = %s, want %q", out, `"1240.00"`)
	}
}
