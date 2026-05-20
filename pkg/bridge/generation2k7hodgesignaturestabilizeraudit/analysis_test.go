package generation2k7hodgesignaturestabilizeraudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate634Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.StarPreservesK7 || !a.Inherited.K7HodgeStable || !a.Inherited.NoBoundaryAssignment || !a.Inherited.Gate633FirewallPreserved {
		t.Fatalf("bad Gate633 inheritance: %+v", a.Inherited)
	}
	if !a.RestrictedOperator.Symmetric || !a.RestrictedOperator.Orthogonal || !a.RestrictedOperator.Involutive || a.RestrictedOperator.Rows != 7 || a.RestrictedOperator.Cols != 7 {
		t.Fatalf("bad restricted operator certificate: %+v", a.RestrictedOperator)
	}
	if a.RestrictedOperator.SymmetryResidual > 1e-12 || a.RestrictedOperator.OrthogonalityResidual > 1e-12 || a.RestrictedOperator.InvolutionResidual > 1e-12 || math.Abs(a.RestrictedOperator.Trace-1) > 1e-10 || math.Abs(a.RestrictedOperator.Determinant+1) > 1e-8 {
		t.Fatalf("bad restricted operator diagnostics: %+v", a.RestrictedOperator)
	}
	if a.Spectrum.PlusRank != 4 || a.Spectrum.MinusRank != 3 || !a.Spectrum.Mixed || a.Spectrum.FullySelfDual || a.Spectrum.FullyAntiSelfDual || len(a.Spectrum.Eigenvalues) != 7 {
		t.Fatalf("unexpected K7 Hodge spectrum: %+v", a.Spectrum)
	}
	for i, v := range a.Spectrum.Eigenvalues {
		want := 1.0
		if i >= 4 {
			want = -1.0
		}
		if math.Abs(v-want) > 1e-8 {
			t.Fatalf("eigenvalue %d got %.16g want %.16g; spectrum=%+v", i, v, want, a.Spectrum)
		}
	}
	if !a.InternalProjectors.ProjectorsCertified || a.InternalProjectors.PlusProjectorRank != 4 || a.InternalProjectors.MinusProjectorRank != 3 || a.InternalProjectors.PlusProjectorIdempotence > 1e-12 || a.InternalProjectors.MinusProjectorIdempotence > 1e-12 {
		t.Fatalf("bad internal projectors: %+v", a.InternalProjectors)
	}
	if a.AmbientProjection.AmbientSelfDualRank != 35 || a.AmbientProjection.AmbientAntiSelfDualRank != 35 || a.AmbientProjection.AmbientHodgeStarSquaredResidual > 1e-12 || math.Abs(a.AmbientProjection.K7SelfDualFrobeniusSquared-4) > 1e-8 || math.Abs(a.AmbientProjection.K7AntiSelfDualFrobeniusSquared-3) > 1e-8 {
		t.Fatalf("bad ambient projection audit: %+v", a.AmbientProjection)
	}
	if !a.Classification.K7MixedHodgePolarity || a.Classification.K7FullySelfDual || a.Classification.K7FullyAntiSelfDual || a.Classification.PlusDimension != 4 || a.Classification.MinusDimension != 3 {
		t.Fatalf("bad classification: %+v", a.Classification)
	}
	if a.Consequences.K7ToW7PairingReopened || a.Consequences.OctonionicResidualReopened || a.Consequences.BoundaryAssignmentPromoted || a.Consequences.SevenOver72Promoted {
		t.Fatalf("prior-route firewall reopened: %+v", a.Consequences)
	}
	if a.Firewalls.ClaimsBoundaryStressAssignment || a.Firewalls.ClaimsSevenOver72Theorem || a.Firewalls.ClaimsScalarRGMatching || a.Firewalls.ClaimsHiggsMassDerivation || a.Firewalls.ClaimsFlavorDerivation || a.Firewalls.ClaimsCKMPMNSDerivation || a.Firewalls.ClaimsGaugeUnification || a.Firewalls.ClaimsPhysicalOrientation {
		t.Fatalf("firewall breach: %+v", a.Firewalls)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2K7HodgeSignatureStabilizerAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusGate633Inherited, StatusRestrictedHodgeOperatorDefined, StatusSKOrthogonalSymmetricInvolutive, StatusSpectrumComputed, StatusMixedHodgeSignature, StatusNotFullySelfDual, StatusNotFullyAntiSelfDual, StatusNoBoundaryStressAssignment, StatusGate634Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
