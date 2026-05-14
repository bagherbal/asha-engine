package contactlqt3r

import (
	"math"
	"testing"
)

func TestGate136ContactT3RChiralityPullbackObstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.MatterT3ROperatorFound || !a.MatterChiralRestricted || !a.MatterHyperauditSelectsBranch {
		t.Fatalf("matter-side T3R diagnostics should be available: %+v", a.Summary)
	}
	if a.MatterFullSMTableDerived {
		t.Fatalf("full SM table should not be derived by the matter audit")
	}
	if math.Abs(a.Summary.HalfBMinusLDifference-2.0/3.0) > 1e-10 || a.Summary.HypotheticalYValueCount != 4 {
		t.Fatalf("unexpected hypothetical Y diagnostic: %+v", a.Summary)
	}
	if a.ContactPullbackRowsDerived != 0 || a.ContactT3RRowsDerived != 0 || a.ContactChiralityRowsDerived != 0 || a.SignedBLRowsDerived != 0 || a.WeakSU2RowsDerived != 0 || a.HyperchargeRowsDerived != 0 || a.ElectricChargeRowsDerived != 0 {
		t.Fatalf("unexpected contact charge rows derived: %+v", a.Summary)
	}
	if !a.BetaPermissionFirewallClosed || a.RepresentationCompleteRows != 0 || a.ContactBetaRowsAllowed != 0 || a.ContactZeroRowsProved != 0 {
		t.Fatalf("beta firewall should remain closed: %+v", a.Summary)
	}
	if a.ResidualS6Choices != 720 || a.ResidualNullityBefore != 3 || a.ResidualNullityAfter != 3 {
		t.Fatalf("residual obstruction changed: %+v", a.Summary)
	}
}

func TestGate136RowsRemainHypotheticalOnly(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if len(a.Rows) != 6 {
		t.Fatalf("expected six leptoquark rows, got %d", len(a.Rows))
	}
	for _, r := range a.Rows {
		if !r.MatterT3RDiagnostic || len(r.MatterT3RCandidateValues) != 2 || len(r.HypotheticalYValues) != 4 || !r.RequiresS6Choice {
			t.Fatalf("row should retain matter diagnostic and S6 obstruction: %+v", r)
		}
		if r.ContactT3RDerived || r.ContactChiralityDerived || r.SignedBLDerived || r.WeakSU2Derived || r.HyperchargeDerived || r.ElectricChargeDerived || r.LocalFieldDerived || r.RepresentationComplete || r.BetaPermitted {
			t.Fatalf("row was over-promoted: %+v", r)
		}
	}
}
