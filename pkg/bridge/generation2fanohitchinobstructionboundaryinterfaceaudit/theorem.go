package generation2fanohitchinobstructionboundaryinterfaceaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2FanoHitchinObstructionBoundaryInterfaceAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 655 — Fano-Hitchin Obstruction Boundary-Interface Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate655 Fano-Hitchin boundary-interface audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate654 internal Fano-Hitchin mechanism and firewall", Passed: a.Inherited.InternalMechanismSourced && a.Inherited.PGForcesFanoNormalForm && a.Inherited.GaugeControlledSource && !a.Inherited.BasisFreeSourceTheorem && near(a.Inherited.CosTheta, 13.0/math.Sqrt(217)) && near(a.Inherited.RhoSquared, 48.0/217.0) && !a.Inherited.ClaimsSplitG2 && !a.Inherited.ClaimsBoundaryStress && !a.Inherited.ClaimsSevenOver72 && !a.Inherited.ClaimsScalarFlavor && !a.Inherited.ClaimsPhysicalMetric && a.Inherited.Gate654FirewallPreserved, Detail: FormatInherited(a.Inherited)},
			{Name: "construct internal invariant ledger", Passed: len(a.Invariants.Rows) >= 10 && near(a.Invariants.TraceSK, 1) && near(a.Invariants.TraceGUn, -5) && near(a.Invariants.Norm2SK, 7) && near(a.Invariants.Norm2GUn, 31) && near(a.Invariants.DetGUn, -27) && near(a.Invariants.ProjectiveInner, 13.0/math.Sqrt(217)) && near(a.Invariants.ObstructionSquare, 48.0/217.0) && a.Invariants.RankK7 == 7 && a.Invariants.RankPlus == 4 && a.Invariants.RankMinus == 3 && a.Invariants.AllNativeFinite && a.Invariants.AllGaugeClassified && !a.Invariants.BoundaryDataPresent, Detail: FormatInvariants(a.Invariants)},
			{Name: "audit 7/72 interface", Passed: near(a.SevenOver72.CandidateWeight, 7.0/72.0) && a.SevenOver72.FanoAddsBeyondNumerator && !a.SevenOver72.BoundaryPairSupplied && !a.SevenOver72.TraceMapSupplied && a.SevenOver72.StructuresNumerator7 && !a.SevenOver72.CertifiedSevenOver72Theorem, Detail: FormatSevenOver72(a.SevenOver72)},
			{Name: "audit boundary-stress interface", Passed: len(a.BoundaryStress.Rows) == 5 && !a.BoundaryStress.CertifiedBoundaryStressSource && a.BoundaryStress.NearBridgeClueOnly && a.BoundaryStress.NoArbitrarySearch, Detail: FormatBoundaryStress(a.BoundaryStress)},
			{Name: "audit HistoryLoopUnit interface", Passed: near(a.HistoryLoop.TargetL, LHistory) && !a.HistoryLoop.SuppliesPiOrS1 && !a.HistoryLoop.SuppliesHeatKernel && !a.HistoryLoop.SuppliesAngularMeasure && a.HistoryLoop.FiniteAlgebraicOnly && !a.HistoryLoop.CertifiedSource, Detail: FormatHistoryLoop(a.HistoryLoop)},
			{Name: "audit flavor orientation interface", Passed: len(a.Flavor.Targets) == 5 && !a.Flavor.UsesFlavorData && !a.Flavor.TypedIntertwinerSupplied && !a.Flavor.ObstructionAngleMappedToFlavor && a.Flavor.RejectsNumericalProximityWithoutMap && !a.Flavor.CertifiedFlavorMap, Detail: FormatFlavor(a.Flavor)},
			{Name: "audit missing boundary map", Passed: !a.BoundaryMap.HasPsi && !a.BoundaryMap.HasTau && !a.BoundaryMap.CanAssignBoundaryPair && !a.BoundaryMap.CanAssignSevenOver72 && !a.BoundaryMap.CanAssignScalarFlavor, Detail: FormatBoundaryMap(a.BoundaryMap)},
			{Name: "define internal Fano-Hitchin obstruction seal", Passed: a.Seal.Name == "FanoHitchinObstructionSeal" && a.Seal.InternalOnly && strings.Contains(a.Seal.Verdict, StatusFanoHitchinSealDefined) && strings.Contains(a.Seal.Verdict, StatusNoBoundaryInterface), Detail: FormatSeal(a.Seal)},
			{Name: "preserve split-G2, boundary, scalar/flavor, history, physical, and 7/72 firewalls", Passed: !a.Firewalls.ClaimsSplitG2 && !a.Firewalls.ClaimsBoundaryStress && !a.Firewalls.ClaimsSevenOver72 && !a.Firewalls.ClaimsScalarFlavor && !a.Firewalls.ClaimsHistoryLoopUnit && !a.Firewalls.ClaimsPhysicalMetric && !a.Firewalls.ClaimsHiggsMass && !a.Firewalls.ClaimsCKMPMNS && !a.Firewalls.ClaimsGaugeUnification && a.Firewalls.Verdict == StatusGate655Boundary, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		notes = append(notes, "Gate655 classifies the Fano-Hitchin package as an internal obstruction seal unless a future theorem constructs Psi:K_7/FanoHitchinPackage->R^2_boundary or tau_defect with normalized trace 7/72.")
		if !strings.Contains(a.Seal.Verdict, StatusNoBoundaryInterface) {
			notes = append(notes, "WARNING_MISSING_BOUNDARY_INTERFACE_FIREWALL")
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
