package generation2boundaryalphadomaintransportmapconstructionandobstructionaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "GATE-828-BOUNDARY-ALPHA-DOMAIN-TRANSPORT-MAP-CONSTRUCTION-OBSTRUCTION"
	theoremName = "Gate 828 — BoundaryAlphaDomainTransportMap Construction/Obstruction Audit"
)

func Generation2BoundaryAlphaDomainTransportMapConstructionAndObstructionAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate 828 boundary-alpha transport audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 827 candidate alpha law and rebuild alpha_B", Passed: math.Abs(a.Ledger.S-SBoundary) < 1e-18 && math.Abs(a.Ledger.AlphaB-a.Ledger.ExpectedAlphaB) < 1e-18 && containsAll(a.Weights.Verdicts, []string{StatusGate827Inherited, StatusCandidateLawRebuilt}), Detail: FormatLedger(a.Ledger)},
			{Name: "verify normalized support-trace weights in both domains", Passed: a.Weights.WeightsVerified && a.Weights.LinearNumerator == 3 && a.Weights.LinearDenominator == 10 && a.Weights.QuadraticNumerator == 7 && a.Weights.QuadraticDenominator == 72 && math.Abs(a.Weights.LinearWeight-0.3) < 1e-15 && math.Abs(a.Weights.QuadraticWeight-7.0/72.0) < 1e-15 && containsAll(a.Weights.Supports, []string{SupportLinearTraceWeight, SupportQuadraticTraceWeight}), Detail: FormatWeights(a.Weights)},
			{Name: "specify linear vector-boundary triplet lane without over-promoting transport", Passed: a.LinearLane.SourcePower == 1 && a.LinearLane.CarrierTyped && a.LinearLane.SupportProjectorTyped && !a.LinearLane.ConcreteTransportMap && !a.LinearLane.ResponseOrderDerived && containsAll(a.LinearLane.Failures, []string{FailureNoLinearTransport, FailurePowerLawNotDerived}), Detail: FormatLane(a.LinearLane)},
			{Name: "specify quadratic H72 K7 lane without over-promoting transport", Passed: a.QuadraticLane.SourcePower == 2 && a.QuadraticLane.CarrierTyped && a.QuadraticLane.SupportProjectorTyped && !a.QuadraticLane.ConcreteTransportMap && !a.QuadraticLane.ResponseOrderDerived && containsAll(a.QuadraticLane.Failures, []string{FailureNoQuadraticTransport, FailurePowerLawNotDerived}), Detail: FormatLane(a.QuadraticLane)},
			{Name: "evaluate BoundaryAlphaDomainTransportMap certification criteria", Passed: a.Criteria.HasSourceScalar && a.Criteria.HasTypedTargetCarriers && a.Criteria.HasSupportTraceWeights && a.Criteria.HasNonCircularDirection && !a.Criteria.HasConcreteLinearMap && !a.Criteria.HasConcreteQuadraticMap && !a.Criteria.HasSharedFunctor && !a.Criteria.HasPowerLawDerivation && !a.Criteria.HasVariationalPrinciple && !a.Criteria.CertifiesNativeAlphaTheorem && strings.Contains(a.Criteria.Classification, "NOT_CERTIFIED_TRANSPORT_MAP") && containsAll(a.Criteria.Failures, []string{FailureNoBoundaryAlphaMap, FailureNoSharedFunctor, FailureNoNativeAlphaTheorem}), Detail: FormatCriteria(a.Criteria)},
			{Name: "enforce noncircular direction and no N_eff backfitting", Passed: a.NonCircular.ComputesAlphaBeforeReadout && !a.NonCircular.UsesNEffToDefineAlpha && !a.NonCircular.UsesObservedYukawas && !a.NonCircular.UsesHiggsMass && strings.Contains(a.NonCircular.Direction, "s -> candidate alpha_B"), Detail: FormatNonCircularity(a.NonCircular)},
			{Name: "defer total operator promotion and preserve coefficient freeze", Passed: !a.Impact.CanPromoteTotalOperator && !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && math.Abs(a.Impact.OfficialCYukawa-CYukawa) < 1e-15 && math.Abs(a.Impact.OfficialCHiggs-CHiggs) < 1e-15 && containsAll(a.Impact.Failures, []string{FailureNoTraceMagnitudeReadout, FailureNoSectorLedger, FailureNoCYukawaUpdate}), Detail: FormatImpact(a.Impact)},
			{Name: "enforce physical and theorem-layer firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.NoBoundaryAlphaMap && a.Firewalls.DimensionRatioOnly && a.Firewalls.NoLinearTransport && a.Firewalls.NoQuadraticTransport && a.Firewalls.NoSharedFunctor && a.Firewalls.PowerLawNotDerived && a.Firewalls.NoTraceMagnitudeReadout && a.Firewalls.NoSectorLedger && a.Firewalls.NotR4 && a.Firewalls.NoCYukawaUpdate && a.Firewalls.NoPMNSCKM && a.Firewalls.NoHiggs && a.Firewalls.Verdict == StatusFirewallGate828, Detail: a.Firewalls.Verdict},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatWeights(a.Weights), FormatLane(a.LinearLane), FormatLane(a.QuadraticLane), FormatCriteria(a.Criteria), FormatNonCircularity(a.NonCircular), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
