package generation2octonionicfanocalibrationnormalformaudit

import (
	"strings"
	"testing"
)

func TestGate652Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.CalibrationInherited || !a.Inherited.AAAUnit || !a.Inherited.AABEqualNegativeUnits || !a.Inherited.ReconstructsPPlusMinus3 || a.Inherited.FullSymbolicCalibration || a.Inherited.SplitG2Certified || a.Inherited.BoundaryStressAssignment || a.Inherited.SevenOver72Theorem || a.Inherited.ScalarFlavorTransport || a.Inherited.PhysicalMetric || !a.Inherited.Gate651FirewallPreserved {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if !a.BVolume.BIsVolumeForm || a.BVolume.Beta != 1 || a.BVolume.ResidualAgainstVolume > tol || len(a.BVolume.Basis) != 3 {
		t.Fatalf("bad B volume audit: %+v", a.BVolume)
	}
	if !a.AExtract.AllExtracted || !a.AExtract.OrthogonalTriple || !a.AExtract.EqualNorms || !a.AExtract.WedgeOrthonormal || a.AExtract.Residual > tol || len(a.AExtract.Rows) != 3 {
		t.Fatalf("bad A extraction: %+v", a.AExtract)
	}
	for _, r := range a.AExtract.Rows {
		if !r.Extracted || r.NormSquared != 1 || r.InnerWithOthers != 0 || r.WedgeSelfCoefficient != 1 || r.WedgeCrossResidual > tol || r.SelfDualSign != 1 {
			t.Fatalf("bad two-form row: %+v", r)
		}
	}
	if !a.Quaternionic.FormsDefineEndomorphisms || !a.Quaternionic.WedgeIdentityPasses || !a.Quaternionic.QuaternionicIdentities || a.Quaternionic.IdentityResidual > tol {
		t.Fatalf("bad quaternionic audit: %+v", a.Quaternionic)
	}
	if a.AAA.CPositive != 1 || !a.AAA.ScalarMultipleOfP || a.AAA.AnisotropyResidual > tol {
		t.Fatalf("bad AAA derivation: %+v", a.AAA)
	}
	if !a.AAB.EqualToMinusPositive || a.AAB.CombinedCoefficient != -3 || a.AAB.CombinedResidual > tol || len(a.AAB.Rows) != 3 {
		t.Fatalf("bad AAB derivation: %+v", a.AAB)
	}
	for _, r := range a.AAB.Rows {
		if r.Coefficient != -1 || r.Target != -1 || !r.ScalarMultipleOfP || r.AnisotropyResidual > tol {
			t.Fatalf("bad negative channel row: %+v", r)
		}
	}
	if !a.EqualUnit.SameAlphaBetaNormalization || !a.EqualUnit.FanoIncidenceSymmetry || !a.EqualUnit.QuaternionicNormalization || a.EqualUnit.RouteSpecificOnly || a.EqualUnit.BasisFreeProofCertified {
		t.Fatalf("bad equal unit audit: %+v", a.EqualUnit)
	}
	if !a.Routes.AllRoutesReduce || !a.Routes.SameNormalFormAfterNorm || a.Routes.RouteDependentScale || len(a.Routes.Rows) != 3 {
		t.Fatalf("bad routes: %+v", a.Routes)
	}
	if !a.Theorem.FiniteNormalFormIdentitiesPass || a.Theorem.FullSymbolicOctonionicTheorem {
		t.Fatalf("bad theorem: %+v", a.Theorem)
	}
	if a.Firewalls.ClaimsFullSymbolicOctonionicTheorem || a.Firewalls.ClaimsSplitG2 || a.Firewalls.ClaimsBoundaryStress || a.Firewalls.ClaimsSevenOver72 || a.Firewalls.ClaimsScalarFlavor || a.Firewalls.ClaimsPhysicalMetric || a.Firewalls.ClaimsHiggsMass || a.Firewalls.ClaimsCKMPMNS || a.Firewalls.ClaimsGaugeUnification || a.Firewalls.Verdict != StatusGate652Boundary {
		t.Fatalf("firewall breach: %+v", a.Firewalls)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2OctonionicFanoCalibrationNormalFormIdentityAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusGate651CalibrationInherited, StatusBNegativeVolumeAudited, StatusATwoFormTripleExtracted, StatusOmegaWedgeOrthonormalityAudited, StatusQuaternionicTripleAudited, StatusAAADerivedFromTriple, StatusAABDerivedFromVolumeTriple, StatusEqualUnitFromCalibration, StatusNegativeSignSourceTraced, StatusCalibrationTheoremSharpened, StatusNoFullSymbolicOctonionicTheorem, StatusNoSplitG2, StatusNoBoundaryStress, StatusNoSevenOver72, StatusNoScalarFlavor, StatusNoPhysicalMetric, StatusNoHiggsFlavorGauge, StatusGate652Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
