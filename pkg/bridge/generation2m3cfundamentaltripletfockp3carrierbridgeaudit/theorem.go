package generation2m3cfundamentaltripletfockp3carrierbridgeaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "GATE-833-M3C-FUNDAMENTAL-TRIPLET-FOCK-P3-CARRIER-BRIDGE"
	theoremName = "Gate 833 — M_3(C) Fundamental Triplet / Fock P_3 Carrier-Bridge Audit"
)

func Generation2M3CFundamentalTripletFockP3CarrierBridgeAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate 833 M3/P3 carrier-bridge audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 832 R2++ source/carrier obstruction", Passed: a.Ledger.R2PlusPlusConsolidated && a.Ledger.AlphaSealed && !a.Ledger.R3SectorLedgerCertified && a.Ledger.AggregateAtomCount == AggregateAtomCount && math.Abs(a.Ledger.OperatorNEff-OperatorNEff) < 5e-16 && containsAll(a.Impact.Verdicts, []string{StatusGate832Inherited, StatusR2PlusPlusRetained}), Detail: FormatLedger(a.Ledger)},
			{Name: "audit M3(C) fundamental triplet carrier", Passed: a.M3Carrier.CanonicalCarrier && a.M3Carrier.MatrixUnitsAudited && a.M3Carrier.Dim == M3FundamentalDim && a.M3Carrier.MatrixUnitCount == M3MatrixUnitCount && a.M3Carrier.IdentityTrace == M3FundamentalDim && a.M3Carrier.SuppliesCarrierProjectors && !a.M3Carrier.SuppliesYukawaMagnitudes && !a.M3Carrier.UsesObservedYukawaData && containsAll(a.M3Carrier.Supports, []string{SupportM3SuppliesCanonicalFundamentalTriplet, SupportM3MatrixUnitsSourceCarrierProjectors}), Detail: FormatM3Carrier(a.M3Carrier)},
			{Name: "audit Fock P3 rank-three carrier without M3 action", Passed: a.FockP3.RankThreeCarrier && a.FockP3.BMinusLSelectorTyped && a.FockP3.P3Rank == FockP3Rank && !a.FockP3.M3ActionCertified && !a.FockP3.RepresentsM3 && containsAll(a.FockP3.Failures, []string{FailureNoM3ActionOnFockP3, FailureNoFockP3RepresentationOfM3}), Detail: FormatFockP3(a.FockP3)},
			{Name: "compare M3 fundamental and Fock P3 carrier shape without typed identification", Passed: a.ShapeBridge.DimensionMatches && a.ShapeBridge.FormalIsomorphismsExist && a.ShapeBridge.NonCircular && !a.ShapeBridge.CanonicalIntertwinerCertified && !a.ShapeBridge.TypedBridgeCertified && !a.ShapeBridge.TripletsIdentified && containsAll(a.ShapeBridge.Supports, []string{SupportCarrierShapesMatchDimensionThree, SupportFormalIsomorphismsExist}) && containsAll(a.ShapeBridge.Failures, []string{FailureNoCanonicalM3ToFockP3Intertwiner, FailureM3ColorTripletNotIdentifiedWithFockP3, FailureShapeMatchOnlyNotTypedBridge}), Detail: FormatShapeBridge(a.ShapeBridge)},
			{Name: "audit finite algebra, Morita, and trace-representation bridge routes", Passed: a.Routes.FiniteAlgebraRouteAudited && a.Routes.MoritaRouteAudited && a.Routes.TraceRepresentationRouteAudited && !a.Routes.ActionLawCertified && !a.Routes.CommutationLawCertified && !a.Routes.RepresentationLawCertified && !a.Routes.CanonicalIntertwinerCertified && !a.Routes.CarrierBridgeCertified && containsAll(a.Routes.Failures, []string{FailureNoM3ActionOnFockP3, FailureNoIntertwiningLaw, FailureNoMoritaBridge, FailureNoTraceRepresentationBridge}), Detail: FormatRoutes(a.Routes)},
			{Name: "audit top I3 carrier compatibility without identifying carriers", Passed: a.TopI3.IdentityShapeMatches && a.TopI3.TopDim == TopBlockDim && a.TopI3.M3Dim == M3FundamentalDim && a.TopI3.FockP3Dim == FockP3Rank && !a.TopI3.SameAsM3FundamentalCertified && !a.TopI3.SameAsFockP3Certified && !a.TopI3.TopToM3ToP3ChainCertified && containsAll(a.TopI3.Failures, []string{FailureTopI3NotCarrierCompatibleWithM3Fundamental, FailureTopI3NotIdentifiedWithFockP3}), Detail: FormatTopI3(a.TopI3)},
			{Name: "block sector-ledger and trace-magnitude promotion", Passed: a.Impact.CarrierShapeSupport && !a.Impact.CarrierBridgeCertified && !a.Impact.SectorProjectorMapCertified && !a.Impact.SectorTraceLedgerCertified && !a.Impact.TraceMagnitudeReadoutCertified && !a.Impact.CanPromoteToR3 && !a.Impact.CanPromoteToR4 && !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && strings.Contains(a.Impact.NextMissingObject, "canonical M_3(C) fundamental -> Fock P_3") && containsAll(a.Impact.Failures, []string{FailureCarrierBridgeNotSectorLedger, FailureNoSectorProjectorMap, FailureNoSectorMagnitudeReadout, FailureNoNEffUpdate, FailureNoCYukawaUpdate}), Detail: FormatImpact(a.Impact)},
			{Name: "preserve Gate 833 physical firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.ShapeNotBridge && a.Firewalls.NoM3P3Intertwiner && a.Firewalls.NoM3ActionOnP3 && a.Firewalls.TopI3Separated && a.Firewalls.CarrierBridgeNotSectorLedger && a.Firewalls.NoSectorProjectorMap && a.Firewalls.NoSectorLedgerMap && a.Firewalls.NoMagnitudeReadout && a.Firewalls.AlphaSealed && a.Firewalls.NoBoundaryAlphaMap && a.Firewalls.NotR3 && a.Firewalls.NotR4 && a.Firewalls.NoNEffUpdate && a.Firewalls.NoCYukawaUpdate && a.Firewalls.NoObservedYukawaFit && a.Firewalls.NoPMNSCKM && a.Firewalls.NoParticleAssignment && a.Firewalls.Verdict == StatusFirewallGate833, Detail: a.Firewalls.Verdict},
		}
		ok := true
		for _, c := range checks {
			if !c.Passed {
				ok = false
				break
			}
		}
		status := theorem.BridgeRequired
		if !ok {
			status = theorem.FailedRoute
		}
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatM3Carrier(a.M3Carrier), FormatFockP3(a.FockP3), FormatShapeBridge(a.ShapeBridge), FormatRoutes(a.Routes), FormatTopI3(a.TopI3), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
