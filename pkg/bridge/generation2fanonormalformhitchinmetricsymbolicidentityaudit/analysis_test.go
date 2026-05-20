package generation2fanonormalformhitchinmetricsymbolicidentityaudit

import (
	"strings"
	"testing"
)

func TestGate653Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.FanoNormalFormInherited || !a.Inherited.BVolumeForm || !a.Inherited.ATwoFormTriple || !a.Inherited.WedgeOrthonormality || !a.Inherited.QuaternionicTriple || !a.Inherited.AAAChannelFinite || !a.Inherited.AABChannelsFinite || !a.Inherited.FiniteNormalFormIdentities || a.Inherited.FullBasisFreeFanoTheorem || a.Inherited.ClaimsSplitG2 || a.Inherited.ClaimsBoundaryStress || a.Inherited.ClaimsSevenOver72 || a.Inherited.ClaimsScalarFlavor || a.Inherited.ClaimsPhysicalMetric || !a.Inherited.Gate652FirewallPreserved {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if !a.Positive.UsesNormalForm || !a.Positive.UsesWedgeIdentity || a.Positive.CPositive != unitC || !a.Positive.ScalarMultipleOfP || a.Positive.AnisotropyResidual > tol || !a.Positive.SymbolicDerivation {
		t.Fatalf("bad positive block: %+v", a.Positive)
	}
	if !a.Negative.EachEqualsMinusC || a.Negative.CombinedCoefficient != -3 || a.Negative.CombinedTarget != -3 || a.Negative.CombinedResidual > tol || !a.Negative.NegativeSignLocated || !a.Negative.SymbolicDerivation || len(a.Negative.Rows) != 3 {
		t.Fatalf("bad negative block: %+v", a.Negative)
	}
	for _, r := range a.Negative.Rows {
		if r.Coefficient != -1 || r.Target != -1 || !r.UsesVolumeForm || !r.UsesWedgeIdentity || !r.ScalarMultipleOfP || r.AnisotropyResidual > tol {
			t.Fatalf("bad negative channel: %+v", r)
		}
	}
	if !a.Mixed.SymbolicallyZero || a.Mixed.MixedBlockNorm > tol || len(a.Mixed.Cases) != 2 {
		t.Fatalf("bad mixed block: %+v", a.Mixed)
	}
	if a.Normalization.CPositive != 1 || a.Normalization.CAAB != -1 || a.Normalization.CABA != -1 || a.Normalization.CBAA != -1 || !a.Normalization.AllEqualAbs || a.Normalization.RequiresRescale || a.Normalization.Residual > tol {
		t.Fatalf("bad normalization: %+v", a.Normalization)
	}
	if !a.Routes.AllRoutesReduce || !a.Routes.SameSymbolicIdentity || a.Routes.RouteDependentOnly || len(a.Routes.Rows) != 3 {
		t.Fatalf("bad routes: %+v", a.Routes)
	}
	if !a.FinalIdentity.PositiveBlockPasses || !a.FinalIdentity.NegativeBlockPasses || !a.FinalIdentity.MixedBlockPasses || !a.FinalIdentity.EqualNormalizationPasses || !a.FinalIdentity.InternalMechanismClosed || a.FinalIdentity.FullPGToFanoSourceTheorem {
		t.Fatalf("bad final identity: %+v", a.FinalIdentity)
	}
	if a.Firewalls.ClaimsSplitG2 || a.Firewalls.ClaimsBoundaryStress || a.Firewalls.ClaimsSevenOver72 || a.Firewalls.ClaimsScalarFlavor || a.Firewalls.ClaimsPhysicalMetric || a.Firewalls.ClaimsHiggsMass || a.Firewalls.ClaimsCKMPMNS || a.Firewalls.ClaimsGaugeUnification || a.Firewalls.Verdict != StatusGate653Boundary {
		t.Fatalf("firewall breach: %+v", a.Firewalls)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2FanoNormalFormHitchinMetricSymbolicIdentityAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusGate652FanoNormalFormInherited, StatusSymbolicPositiveBlockDerived, StatusSymbolicNegativeBlockDerived, StatusSymbolicMixedBlockVanishingDerived, StatusEqualCNormalizationAudited, StatusRouteNormalizationSingleFano, StatusFanoForcesPPlusMinus3, StatusInternalHitchinMechanismClosed, StatusFanoHitchinSymbolicIdentitySharpened, StatusNoBasisFreePGToFanoTheorem, StatusNoSplitG2, StatusNoBoundaryStress, StatusNoSevenOver72, StatusNoScalarFlavor, StatusNoPhysicalMetric, StatusNoHiggsFlavorGauge, StatusGate653Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
