package contactlqcharge

import (
	"math"
	"testing"
)

func TestGate135BLChargeLatticeObstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.BMinusLChargeBridgeValid || !a.BMinusLPolarizesOnePlusThree {
		t.Fatalf("B-L charge bridge should remain valid: %+v", a.Summary)
	}
	if math.Abs(a.LeptonColorBLDifference-4.0/3.0) > 1e-10 {
		t.Fatalf("unexpected lepton-color B-L difference: %.12f", a.LeptonColorBLDifference)
	}
	if !a.BLDifferenceDiagnostic || a.Summary.BLDiagnosticRows != 6 || a.SignedBLRowsDerived != 0 {
		t.Fatalf("B-L should be diagnostic but not signed contact hypercharge: %+v", a.Summary)
	}
	if a.T3RRowsDerived != 0 || a.WeakChiralityRowsDerived != 0 || a.WeakSU2RowsDerived != 0 || a.HyperchargeRowsDerived != 0 || a.ElectricChargeRowsDerived != 0 {
		t.Fatalf("unexpected charge/hypercharge rows derived: %+v", a.Summary)
	}
	if !a.BetaPermissionFirewallClosed || a.RepresentationCompleteRows != 0 || a.ContactBetaRowsAllowed != 0 || a.ContactZeroRowsProved != 0 {
		t.Fatalf("beta firewall should remain closed: %+v", a.Summary)
	}
	if a.ResidualS6Choices != 720 || a.ResidualNullityBefore != 3 || a.ResidualNullityAfter != 3 {
		t.Fatalf("residual obstruction changed: %+v", a.Summary)
	}
}

func TestGate135RowsRemainDiagnosticsOnly(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if len(a.Rows) != 6 {
		t.Fatalf("expected six leptoquark rows, got %d", len(a.Rows))
	}
	for _, r := range a.Rows {
		if !r.BLMagnitudeDiagnostic || math.Abs(r.BMinusLDifference-4.0/3.0) > 1e-10 || !r.RequiresS6Choice {
			t.Fatalf("row should retain only B-L diagnostic and S6 obstruction: %+v", r)
		}
		if r.SignedBLDerived || r.T3RDerived || r.WeakChiralityDerived || r.WeakSU2Derived || r.HyperchargeDerived || r.ElectricChargeDerived || r.LocalFieldDerived || r.RepresentationComplete || r.BetaPermitted {
			t.Fatalf("row was over-promoted: %+v", r)
		}
	}
}
