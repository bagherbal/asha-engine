package generation2gaugescalarboundarystresssourcetypeaudit

import (
	"math"
	"testing"
)

func TestGate614BoundaryStressSourceTyping(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.Inherited.XiBoundary-0.0503471644870914) > 1e-12 {
		t.Fatalf("unexpected xi_boundary %.15g", a.Inherited.XiBoundary)
	}
	if len(a.SourceTypes) != 4 {
		t.Fatalf("expected four source types, got %d", len(a.SourceTypes))
	}
	if !hasSourceVerdict(a.SourceTypes, StatusSpectralActionSlotRelevant) || !hasSourceVerdict(a.SourceTypes, StatusNoNativeXi) {
		t.Fatalf("missing source-type verdicts: %+v", a.SourceTypes)
	}
	if a.BoundaryEquation.AbsResidualOverXi > 0.03 || a.BoundaryEquation.HalfResidualOverXi > 0.02 {
		t.Fatalf("stress residual too large: %+v", a.BoundaryEquation)
	}
}

func TestGate614SpectralActionLaneAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !hasLane("gauge_kinetic_C_i_Tr_F_i2", a.SpectralActionLanes) || !hasLane("scalar_quartic_lambda", a.SpectralActionLanes) || !hasLane("f0_cutoff_moment", a.SpectralActionLanes) {
		t.Fatalf("missing expected lanes: %+v", a.SpectralActionLanes)
	}
	if !a.KineticQuarticPairing.SymbolicPairingSlot || a.KineticQuarticPairing.NativeCoefficientLaw {
		t.Fatalf("bad kinetic/quartic pairing: %+v", a.KineticQuarticPairing)
	}
	if a.NativeStatus.NativeXiBoundary || a.NativeStatus.NativeF0SectorSplit || a.NativeStatus.NativeThresholdSpectrum || a.NativeStatus.NativeGaugeScalarCoefficientLaw {
		t.Fatalf("native theorem violation: %+v", a.NativeStatus)
	}
	if a.Firewalls.ClaimsGaugeUnification || a.Firewalls.ClaimsThresholdExists || a.Firewalls.ClaimsHiggsMassPrediction || a.Firewalls.ClaimsNativeCorrection {
		t.Fatalf("firewall violation: %+v", a.Firewalls)
	}
}
