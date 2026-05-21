package generation2sectortraceledgermapcandidatesourceandcarriercompatibilityaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "GATE-832-SECTOR-TRACE-LEDGER-MAP-CANDIDATE-SOURCE-CARRIER-COMPATIBILITY"
	theoremName = "Gate 832 — SectorTraceLedgerMap Candidate Source and Carrier-Compatibility Audit"
)

func Generation2SectorTraceLedgerMapCandidateSourceAndCarrierCompatibilityAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate 832 sector-ledger compatibility audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit R2++ aggregate carrier without R3 promotion", Passed: a.Ledger.R2PlusPlusConsolidated && a.Ledger.AlphaSealed && !a.Ledger.R3SectorLedgerCertified && a.Ledger.AggregateAtomCount == AggregateAtomCount && math.Abs(a.Ledger.OperatorNEff-OperatorNEff) < 5e-16 && containsAll(a.Impact.Verdicts, []string{StatusGate831Inherited, StatusR2PlusPlusRetained}), Detail: FormatLedger(a.Ledger)},
			{Name: "audit A_F as strongest finite-sector projector source", Passed: a.FiniteSector.StrongestLawfulSource && a.FiniteSector.SectorProjectorCandidatesAudited && a.FiniteSector.SuppliesGaugeInternalCarriers && !a.FiniteSector.SuppliesYukawaMagnitudes && !a.FiniteSector.SuppliesTraceMagnitudeReadout && !a.FiniteSector.UsesObservedYukawaData && containsAll(a.FiniteSector.Supports, []string{SupportAFStrongestSectorProjectorSource, SupportFiniteSectorProjectorCandidates}) && containsAll(a.FiniteSector.Failures, []string{FailureSectorProjectorsDoNotSupplyMagnitudes, FailureSectorProjectorNotYukawaValue}), Detail: FormatFiniteSector(a.FiniteSector)},
			{Name: "test aggregate carrier compatibility with finite sector projectors", Passed: !a.Compatibility.CarrierMapCertified && !a.Compatibility.CarrierCompatibilityCertified && !a.Compatibility.FiniteAlgebraCommutationProven && a.Compatibility.NonCircular && !a.Compatibility.Compatible && containsAll(a.Compatibility.Failures, []string{FailureAggregateNotCompatibleWithSectorProjectors, FailureNoTypedAggregateToSectorMap, FailureNoCarrierCompatibilityTheorem, FailureNoSectorTraceLedgerMap}), Detail: FormatCompatibility(a.Compatibility)},
			{Name: "audit seven-count route without K7 promotion", Passed: a.K7Route.CountMatchesK7 && !a.K7Route.TypedMapCertified && !a.K7Route.ProjectorIdentityCertified && !a.K7Route.PromotedToK7 && containsAll(a.K7Route.Failures, []string{FailureSevenAtomAggregateNotK7, FailureNoAggregateToK7Map}), Detail: FormatK7Route(a.K7Route)},
			{Name: "audit dual-triplet bridge route through finite algebra, Morita, and trace representation", Passed: a.DualTriplet.FiniteAlgebraRouteAudited && a.DualTriplet.MoritaRouteAudited && a.DualTriplet.TraceRepresentationRouteAudited && !a.DualTriplet.TypedBridgeCertified && !a.DualTriplet.TripletsIdentified && containsAll(a.DualTriplet.Failures, []string{FailureColorTripletToFockTripletNotCertified, FailureFiniteAlgebraDoesNotIdentifyTriplets}), Detail: FormatDualTriplet(a.DualTriplet)},
			{Name: "separate sector projectors from trace magnitudes", Passed: a.Magnitude.SectorProjectorsFoundAsCandidates && a.Magnitude.ProjectorsAreCarriers && !a.Magnitude.ProjectorsAreMagnitudes && !a.Magnitude.PositiveTraceAtomsDerived && !a.Magnitude.ReadoutMapCertified && !a.Magnitude.YukawaValuesDerived && containsAll(a.Magnitude.Failures, []string{FailureSectorProjectorsDoNotSupplyMagnitudes, FailureNoSectorMagnitudeReadout, FailureSectorProjectorNotYukawaValue, FailureNoNativeYukawaOperator}), Detail: FormatMagnitude(a.Magnitude)},
			{Name: "block R3/R4 promotion and ledger updates", Passed: !a.Impact.CanPromoteToR3 && !a.Impact.CanPromoteToR4 && !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && strings.Contains(a.Impact.NextMissingObject, "SectorProjectorMap") && strings.Contains(a.Impact.NextMissingObject, "SectorTraceMagnitudeReadoutMap") && containsAll(a.Impact.Failures, []string{FailureAlphaStillSealed, FailureNoBoundaryAlphaMap, FailureNoSectorTraceLedgerMap, FailureNoSectorMagnitudeReadout, FailureNoNEffUpdate, FailureNoCYukawaUpdate}), Detail: FormatImpact(a.Impact)},
			{Name: "preserve sector-ledger and physical firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.AFProjectorsNotMagnitudes && a.Firewalls.AggregateNotSectorLedger && a.Firewalls.NoCarrierMap && a.Firewalls.NoSectorLedgerMap && a.Firewalls.NoMagnitudeReadout && a.Firewalls.DualTripletSeparated && a.Firewalls.SevenNotK7 && a.Firewalls.AlphaSealed && a.Firewalls.NoBoundaryAlphaMap && a.Firewalls.NotR3 && a.Firewalls.NotR4 && a.Firewalls.NoNEffUpdate && a.Firewalls.NoCYukawaUpdate && a.Firewalls.NoObservedYukawaFit && a.Firewalls.NoPMNSCKM && a.Firewalls.NoParticleAssignment && a.Firewalls.Verdict == StatusFirewallGate832, Detail: a.Firewalls.Verdict},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatFiniteSector(a.FiniteSector), FormatCompatibility(a.Compatibility), FormatK7Route(a.K7Route), FormatDualTriplet(a.DualTriplet), FormatMagnitude(a.Magnitude), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
