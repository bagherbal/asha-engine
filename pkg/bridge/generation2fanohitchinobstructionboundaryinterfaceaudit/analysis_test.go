package generation2fanohitchinobstructionboundaryinterfaceaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate655Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.InternalMechanismSourced || !a.Inherited.PGForcesFanoNormalForm || !a.Inherited.GaugeControlledSource || a.Inherited.BasisFreeSourceTheorem || !near(a.Inherited.CosTheta, 13/math.Sqrt(217)) || !near(a.Inherited.RhoSquared, 48.0/217.0) || a.Inherited.ClaimsSplitG2 || a.Inherited.ClaimsBoundaryStress || a.Inherited.ClaimsSevenOver72 || a.Inherited.ClaimsScalarFlavor || a.Inherited.ClaimsPhysicalMetric || !a.Inherited.Gate654FirewallPreserved {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if len(a.Invariants.Rows) < 10 || a.Invariants.TraceSK != 1 || a.Invariants.TraceGUn != -5 || a.Invariants.Norm2SK != 7 || a.Invariants.Norm2GUn != 31 || a.Invariants.DetGUn != -27 || !near(a.Invariants.ProjectiveInner, 13/math.Sqrt(217)) || !near(a.Invariants.ObstructionSquare, 48.0/217.0) || a.Invariants.RankK7 != 7 || a.Invariants.RankPlus != 4 || a.Invariants.RankMinus != 3 || !a.Invariants.AllNativeFinite || !a.Invariants.AllGaugeClassified || a.Invariants.BoundaryDataPresent {
		t.Fatalf("bad invariant ledger: %+v", a.Invariants)
	}
	if !near(a.SevenOver72.CandidateWeight, 7.0/72.0) || !a.SevenOver72.FanoAddsBeyondNumerator || a.SevenOver72.BoundaryPairSupplied || a.SevenOver72.TraceMapSupplied || !a.SevenOver72.StructuresNumerator7 || a.SevenOver72.CertifiedSevenOver72Theorem {
		t.Fatalf("bad 7/72 audit: %+v", a.SevenOver72)
	}
	if len(a.BoundaryStress.Rows) != 5 || a.BoundaryStress.CertifiedBoundaryStressSource || !a.BoundaryStress.NearBridgeClueOnly || !a.BoundaryStress.NoArbitrarySearch {
		t.Fatalf("bad boundary stress audit: %+v", a.BoundaryStress)
	}
	if a.HistoryLoop.SuppliesPiOrS1 || a.HistoryLoop.SuppliesHeatKernel || a.HistoryLoop.SuppliesAngularMeasure || !a.HistoryLoop.FiniteAlgebraicOnly || a.HistoryLoop.CertifiedSource {
		t.Fatalf("bad history interface audit: %+v", a.HistoryLoop)
	}
	if a.Flavor.UsesFlavorData || a.Flavor.TypedIntertwinerSupplied || a.Flavor.ObstructionAngleMappedToFlavor || !a.Flavor.RejectsNumericalProximityWithoutMap || a.Flavor.CertifiedFlavorMap {
		t.Fatalf("bad flavor audit: %+v", a.Flavor)
	}
	if a.BoundaryMap.HasPsi || a.BoundaryMap.HasTau || a.BoundaryMap.CanAssignBoundaryPair || a.BoundaryMap.CanAssignSevenOver72 || a.BoundaryMap.CanAssignScalarFlavor {
		t.Fatalf("bad boundary map audit: %+v", a.BoundaryMap)
	}
	if a.Seal.Name != "FanoHitchinObstructionSeal" || !a.Seal.InternalOnly || !strings.Contains(a.Seal.Verdict, StatusFanoHitchinSealDefined) || !strings.Contains(a.Seal.Verdict, StatusNoBoundaryInterface) {
		t.Fatalf("bad seal: %+v", a.Seal)
	}
	if a.Firewalls.ClaimsSplitG2 || a.Firewalls.ClaimsBoundaryStress || a.Firewalls.ClaimsSevenOver72 || a.Firewalls.ClaimsScalarFlavor || a.Firewalls.ClaimsHistoryLoopUnit || a.Firewalls.ClaimsPhysicalMetric || a.Firewalls.ClaimsHiggsMass || a.Firewalls.ClaimsCKMPMNS || a.Firewalls.ClaimsGaugeUnification || a.Firewalls.Verdict != StatusGate655Boundary {
		t.Fatalf("firewall breach: %+v", a.Firewalls)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2FanoHitchinObstructionBoundaryInterfaceAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusGate654InternalHitchinInherited, StatusInternalInvariantLedgerBuilt, StatusSevenOver72InterfaceAudited, StatusBoundaryStressInterfaceAudited, StatusHistoryLoopInterfaceAudited, StatusFlavorOrientationAudited, StatusBoundaryMapObstructionAudited, StatusFanoStructuresNumerator7, StatusFanoHitchinSealDefined, StatusNoBoundaryInterface, StatusNoSevenOver72Theorem, StatusNoBoundaryStress, StatusNoScalarFlavorMap, StatusNoHistoryLoopSource, StatusNoSplitG2, StatusNoPhysicalMetric, StatusNoHiggsCKMGauge, StatusGate655Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
