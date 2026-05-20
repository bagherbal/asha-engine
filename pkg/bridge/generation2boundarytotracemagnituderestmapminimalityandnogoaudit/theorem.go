package generation2boundarytotracemagnituderestmapminimalityandnogoaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "GATE-814-BOUNDARY-TO-TRACE-MAGNITUDE-REST-MAP-MINIMALITY-NO-GO"
	theoremName = "Gate 814 — BoundaryToTraceMagnitudeRestMap Minimality and No-Go Audit"
)

func Generation2BoundaryToTraceMagnitudeRestMapMinimalityAndNoGoAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Notes: []string{err.Error()}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 813 second-moment rest-pressure result", Passed: a.Inheritance.Gate813Inherited && a.Inheritance.SecondMomentSelected && a.Inheritance.PositiveCompatibilityInherited && math.Abs(a.Inheritance.M2-1.624013231638281e-7) < 1e-20 && math.Abs(a.Inheritance.ResidualImprovement-34.2924) < 1e-3 && containsAll(a.Inheritance.Verdicts, []string{StatusGate813Inherited}) && containsAll(a.Inheritance.Failures, []string{"FAILED_ROUTE_GATE813_DID_NOT_CONSTRUCT_NATIVE_TRACE_MAGNITUDE_REST_MAP"}), Detail: FormatInheritance(a.Inheritance)},
			{Name: "define BoundaryToTraceMagnitudeRestMap and target chain", Passed: a.Map.Defined && len(a.Map.Objects) >= 12 && strings.Contains(a.Map.TargetChain, "C_Yukawa") && containsAll(a.Map.Verdicts, []string{StatusMapDefined, StatusTargetChain, StatusNonCircularityReq}) && containsAll(a.Map.Supports, []string{StatusExactMissingObject}), Detail: FormatMap(a.Map)},
			{Name: "audit minimality of all required subobjects", Passed: a.Minimality.Audited && len(a.Minimality.RemovalFailures) >= 10 && containsAll(a.Minimality.Supports, []string{StatusAllSubobjectsNonCosmetic}) && containsAll(a.Minimality.Failures, []string{StatusMapNotCoefficientFit, StatusNoScaleWithoutS, StatusNoM2WithoutP, StatusNoNineFiveWithoutHyper, StatusNoTopBaselineWithoutColor, StatusNoSixWithoutBoundaryTwo, StatusNoAlphaBetaWithoutTop, StatusDirectClosureNoAtoms, StatusNoYukawaWithoutPositive, StatusNoScaleLocalWithoutScheme, StatusNoPredictionNoNonCirc}), Detail: FormatMinimality(a.Minimality)},
			{Name: "audit existing ASHA objects for map source", Passed: a.Sources.Audited && containsAll(a.Sources.Verdicts, []string{StatusExistingObjects, StatusBoundaryPairSource, StatusK7Source, StatusFiniteTripleSource, StatusHyperchargeSource, StatusExternalLedgerSource, StatusD4Source, StatusChiralityFirewall}) && containsAll(a.Sources.Supports, []string{StatusBoundaryPairSupplies, StatusK7SuppliesM2, StatusFiniteTripleSupplies, StatusHyperchargeSupplies}) && containsAll(a.Sources.Failures, []string{StatusBoundaryPairNoMap, StatusK7NoYukawaAtoms, StatusFiniteTripleNoMap, StatusHyperchargeNotTheorem, StatusExternalNotNative, StatusD4NotMap, StatusChiralityNotSource}), Detail: FormatSources(a.Sources)},
			{Name: "separate scalar Delta_N closure from trace construction", Passed: a.Closure.Defined && strings.Contains(a.Closure.WeakAchievement, "Delta_N") && strings.Contains(a.Closure.StrongAchievement, "positive rest atoms") && containsAll(a.Closure.Supports, []string{StatusSecondMomentStrong}) && containsAll(a.Closure.Failures, []string{StatusClosureNotTraceTheorem, StatusClosureNotSpectrum}), Detail: a.Closure.WeakAchievement + " || " + a.Closure.StrongAchievement},
			{Name: "audit positive spectrum non-uniqueness", Passed: a.Spectrum.Audited && len(a.Spectrum.Examples) >= 3 && containsAll(a.Spectrum.Supports, []string{StatusPositiveWeaker}) && containsAll(a.Spectrum.Failures, []string{StatusPositiveNoSectors, StatusQNoLedger}), Detail: strings.Join(a.Spectrum.Examples, " | ")},
			{Name: "classify rest-map status levels", Passed: a.Levels.Defined && len(a.Levels.Levels) == 5 && strings.Contains(a.Levels.CurrentStatus, "R1") && strings.Contains(a.Levels.CurrentStatus, "partial R2") && containsAll(a.Levels.Supports, []string{StatusR1PartialR2}) && containsAll(a.Levels.Failures, []string{StatusNotR3, StatusNotR4}), Detail: FormatLevels(a.Levels)},
			{Name: "record noncircularity requirements and post-hoc risk", Passed: a.NonCircularity.Defined && len(a.NonCircularity.Rules) >= 4 && containsAll(a.NonCircularity.Supports, []string{StatusTypedSources}) && containsAll(a.NonCircularity.Failures, []string{StatusPostHocRisk, StatusNoPriorCoeffTheorem}), Detail: strings.Join(a.NonCircularity.Rules, " | ")},
			{Name: "record C_Yukawa and C_Higgs candidate impact without updating ledger", Passed: a.Impact.Recorded && math.Abs(a.Impact.CYukawaBoundaryB2-0.9992248096849) < 1e-10 && math.Abs(a.Impact.CYukawaBoundaryB2-a.Impact.OfficialCYukawa) < 2e-8 && math.Abs(a.Impact.OfficialCYukawa-CYukawa) < 1e-15 && containsAll(a.Impact.Supports, []string{StatusCertifiedMapWouldReduce}) && containsAll(a.Impact.Failures, []string{StatusNoCYukawaUpdate, StatusCHiggsLevelB}), Detail: FormatImpact(a.Impact)},
			{Name: "record outcome and branch decision", Passed: a.Outcome.Recorded && len(a.Outcome.Items) == 6 && a.Branch.Recorded && strings.Contains(a.Branch.Recommended, "Gate 815") && containsAll([]string{a.Branch.Support}, []string{StatusNextFreezeProtocol}), Detail: strings.Join(a.Outcome.Items, " | ") + " -> " + a.Branch.Recommended},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.NoCoefficientPromotion && a.Firewalls.NoDeltaAsTraceTheorem && a.Firewalls.NoPositiveAsSectorLedger && a.Firewalls.NoPostHoc && a.Firewalls.NoLedgerUpdate && a.Firewalls.NoPoleMass && a.Firewalls.Verdict == StatusFirewallGate814, Detail: a.Firewalls.Verdict},
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
		notes := []string{a.Truth, FormatInheritance(a.Inheritance), FormatMap(a.Map), FormatMinimality(a.Minimality), FormatSources(a.Sources), a.Closure.WeakAchievement + " != " + a.Closure.StrongAchievement, FormatLevels(a.Levels), strings.Join(a.NonCircularity.Rules, " | "), FormatImpact(a.Impact), strings.Join(a.Outcome.Items, " | "), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
