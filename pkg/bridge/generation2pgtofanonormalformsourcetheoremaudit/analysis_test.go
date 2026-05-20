package generation2pgtofanonormalformsourcetheoremaudit

import (
	"strings"
	"testing"
)

func TestGate654Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.FanoHitchinIdentityInherited || !a.Inherited.NormalFormInherited || !a.Inherited.SymbolicPositive || !a.Inherited.SymbolicNegative || !a.Inherited.SymbolicMixedZero || !a.Inherited.InternalMechanismClosed || a.Inherited.PGToFanoAlreadyBasisFree || a.Inherited.SplitG2Certified || a.Inherited.BoundaryStressAssignment || a.Inherited.SevenOver72Theorem || a.Inherited.ScalarFlavorTransport || a.Inherited.PhysicalMetric || !a.Inherited.Gate653FirewallPreserved || !a.Inherited.Gate652FiniteSourceVisible || a.Inherited.Gate652FullSourceTheorem {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if !a.Support.OmegaPPPZero || !a.Support.OmegaPPMNonzero || !a.Support.OmegaPMMZero || !a.Support.OmegaMMMNonzero || !a.Support.ReducesToLambda21Plus03 || a.Support.Residual > tol || len(a.Support.Rows) != 4 {
		t.Fatalf("bad support decomposition: %+v", a.Support)
	}
	if a.BVolume.Beta != unit || a.BVolume.OrientationSign != 1 || !a.BVolume.SO3VolumeCovariant || a.BVolume.ResidualAgainstVolMinus > tol || !a.BVolume.BasisIndependentVolume {
		t.Fatalf("bad B volume: %+v", a.BVolume)
	}
	if a.AMap.Rank != minusDim || a.AMap.ScaleAlpha != unit || !a.AMap.IsometryUpToScale || !a.AMap.ImageInSelfDualForms || a.AMap.ImageDimension != minusDim || !a.AMap.WedgeOrthonormal || a.AMap.Residual > tol {
		t.Fatalf("bad A map: %+v", a.AMap)
	}
	if !a.Quaternionic.FormsDefineEndomorphisms || !a.Quaternionic.QuaternionicTriple || a.Quaternionic.JIdentityResidual > tol || a.Quaternionic.WedgeIdentityResidual > tol {
		t.Fatalf("bad quaternionic source: %+v", a.Quaternionic)
	}
	if !a.Gauge.AInvariant || !a.Gauge.BVolumeInvariant || !a.Gauge.FMapEquivariant || !a.Gauge.NormalFormGaugeCovariant || a.Gauge.BasisArbitrary {
		t.Fatalf("bad gauge covariance: %+v", a.Gauge)
	}
	if !a.Routes.AllRoutesReduce || !a.Routes.SamePGSourcePackage || a.Routes.RouteDependentOnly || len(a.Routes.Rows) != 3 {
		t.Fatalf("bad route source audit: %+v", a.Routes)
	}
	if !a.SourceTheorem.PGForcesFanoNormalForm || !a.SourceTheorem.GaugeControlledSource || a.SourceTheorem.BasisFreeSourceTheorem || !a.SourceTheorem.Gate653ImplicationAvailable || !a.SourceTheorem.InternalMechanismSourced {
		t.Fatalf("bad source theorem readiness: %+v", a.SourceTheorem)
	}
	if a.Firewalls.ClaimsSplitG2 || a.Firewalls.ClaimsBoundaryStress || a.Firewalls.ClaimsSevenOver72 || a.Firewalls.ClaimsScalarFlavor || a.Firewalls.ClaimsPhysicalMetric || a.Firewalls.ClaimsHiggsMass || a.Firewalls.ClaimsCKMPMNS || a.Firewalls.ClaimsGaugeUnification || a.Firewalls.Verdict != StatusGate654Boundary {
		t.Fatalf("firewall breach: %+v", a.Firewalls)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2PGToFanoNormalFormSourceTheoremAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusGate653FanoHitchinInherited, StatusPGSupportDecompositionAudited, StatusPGSupportLambda21Plus03, StatusPGForcesNegativeVolume, StatusAAsK7MinusToTwoFormsMapAudited, StatusOmegaWedgeOrthonormalitySource, StatusQuaternionicTripleSourceAudited, StatusSO3GaugeCovarianceAudited, StatusRouteSourceIndependenceAudited, StatusPGForcesFanoNormalForm, StatusInternalHitchinFullySourced, StatusPGToFanoSourceTheoremSharpened, StatusNoFullBasisFreePGToFanoTheorem, StatusNoSplitG2, StatusNoBoundaryStress, StatusNoSevenOver72, StatusNoScalarFlavor, StatusNoPhysicalMetric, StatusNoHiggsFlavorGauge, StatusGate654Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
