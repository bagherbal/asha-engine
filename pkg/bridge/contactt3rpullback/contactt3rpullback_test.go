package contactt3rpullback

import "testing"

func TestGate137FockToContactIntertwinerObstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.MatterT3ROperatorFound || !a.MatterChiralRestricted || !a.MatterMirrorAmbiguity {
		t.Fatalf("matter T3R diagnostics should be inherited: %+v", a.Summary)
	}
	if a.MatterDimension != 16 || a.ScalarDimension != 4 || a.TensorDimension != 64 || a.ContactRows != 7 {
		t.Fatalf("unexpected dimensions: %+v", a.Summary)
	}
	if !a.GenericFockToContactMapsExist || !a.GenericTensorToContactMapsExist || a.Summary.FockToContactGenericKernelDim != 9 || a.Summary.TensorToContactGenericKernelDim != 57 {
		t.Fatalf("generic map/kernel audit failed: %+v", a.Summary)
	}
	if a.CanonicalFockToContactMaps != 0 || a.FockToContactIntertwinersDerived != 0 {
		t.Fatalf("no canonical intertwiner should be derived: %+v", a.Summary)
	}
	if a.T3RPullbackRowsDerived != 0 || a.ChiralityPullbackRowsDerived != 0 || a.BMinusLPullbackRowsDerived != 0 || a.SU2LPullbackRowsDerived != 0 || a.HyperchargeRowsDerived != 0 {
		t.Fatalf("unexpected contact pullback rows: %+v", a.Summary)
	}
	if !a.BetaPermissionFirewallClosed || a.RepresentationCompleteRows != 0 || a.ContactBetaRowsAllowed != 0 || a.ContactZeroRowsProved != 0 {
		t.Fatalf("beta firewall should remain closed: %+v", a.Summary)
	}
}

func TestGate137RowsRemainUnpulled(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if len(a.Rows) != 6 {
		t.Fatalf("expected six leptoquark pullback rows, got %d", len(a.Rows))
	}
	for _, r := range a.Rows {
		if len(r.MatterT3RCandidateValues) != 2 || !r.RequiresFockContactMap || !r.RequiresS6Choice {
			t.Fatalf("row should retain Fock/contact and S6 obstruction: %+v", r)
		}
		if r.T3RPullbackDerived || r.ChiralityPullbackDerived || r.BMinusLPullbackDerived || r.SU2LPullbackDerived || r.HyperchargeRowDerived || r.LocalFieldDerived || r.RepresentationComplete || r.BetaPermitted {
			t.Fatalf("row was over-promoted: %+v", r)
		}
	}
}
