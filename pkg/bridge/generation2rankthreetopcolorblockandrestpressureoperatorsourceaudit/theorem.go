package generation2rankthreetopcolorblockandrestpressureoperatorsourceaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "GATE-808-RANK-THREE-TOP-COLOR-BLOCK-AND-REST-PRESSURE-OPERATOR-SOURCE"
	theoremName = "Gate 808 — RankThreeTopColorBlock and RestPressureOperator Source Audit"
)

func Generation2RankThreeTopColorBlockAndRestPressureOperatorSourceAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Notes: []string{err.Error()}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 807 trace-magnitude audit", Passed: a.Inheritance.Gate807Inherited && math.Abs(a.Inheritance.NEff-NEff) < 1e-15 && math.Abs(a.Inheritance.CYukawa-CYukawa) < 1e-15 && containsAll(a.Inheritance.Verdicts, []string{StatusGate807Inherited, StatusTopColorSelected, StatusRestPressureSelected}), Detail: a.Inheritance.CHiggsFormula},
			{Name: "define and rederive rank-three top-color block", Passed: a.TopColor.Defined && a.TopColor.Name == "RankThreeTopColorBlockSeal" && math.Abs(a.TopColor.NEffTop-3) < 1e-12 && containsAll(a.TopColor.Verdicts, []string{StatusTopColorSealDefined, StatusTopColorLimit}) && containsAll(a.TopColor.Supports, []string{StatusExactTopNEffThree, StatusColorThreeStrongest}) && containsAll(a.TopColor.Failures, []string{StatusTopBlockNoT, StatusTopThreeNotGeneration, StatusTopBlockNotHierarchyTheorem}), Detail: FormatTop(a.TopColor)},
			{Name: "define rest pressure operator and decomposition", Passed: a.RestPressure.Defined && strings.Contains(a.RestPressure.NEffFormula, "3(1+alpha)²/(1+beta)") && strings.Contains(a.RestPressure.DeltaFormula, "2alpha") && containsAll(a.RestPressure.Verdicts, []string{StatusRestSealDefined, StatusRestDecomposition}) && containsAll(a.RestPressure.Supports, []string{StatusRestPressureAboveTop, StatusRestPressureDilutesCYukawa}) && containsAll(a.RestPressure.Failures, []string{StatusRestNoSector, StatusRestNotNative}), Detail: FormatRest(a.RestPressure)},
			{Name: "compute aggregate positivity corridor", Passed: a.Corridor.Computed && math.Abs(a.Corridor.AOver3-0.9474698380779695) < 1e-15 && math.Abs(a.Corridor.SqrtBOver3-0.9471025365183062) < 1e-15 && math.Abs(a.Corridor.TLower-0.9471023226011707) < 1e-15 && math.Abs(a.Corridor.AlphaAtUpper-0.00038781604472679744) < 1e-12 && math.Abs(a.Corridor.BetaAtLower-4.5172977535955994e-7) < 1e-16 && containsAll(a.Corridor.Supports, []string{StatusNarrowCorridor, StatusAlphaScale}) && containsAll(a.Corridor.Failures, []string{StatusCorridorNotTopDerivation, StatusNoBackwardT, StatusCorridorNoSector}), Detail: FormatCorridor(a.Corridor)},
			{Name: "record rest concentration ratio and bounds", Passed: a.Concentration.Recorded && strings.Contains(a.Concentration.Formula, "q_rest") && strings.Contains(a.Concentration.BetaFormula, "3 alpha² q_rest") && containsAll(a.Concentration.Verdicts, []string{StatusRestConcentration, StatusRestConcentrationBound}) && containsAll(a.Concentration.Supports, []string{StatusAlphaQRestSplit}) && containsAll(a.Concentration.Failures, []string{StatusNoRestAtomCount, StatusNoQRestFromAggregate}), Detail: FormatConcentration(a.Concentration)},
			{Name: "record rest-pressure sector candidates", Passed: a.SectorCandidates.Audited && containsAll(a.SectorCandidates.Supports, []string{StatusPlausibleRestSources}) && containsAll(a.SectorCandidates.Failures, []string{StatusNoSectorAssignment, StatusNeutrinoImplicit, StatusScaleSchemeUntyped}), Detail: FormatAudit(a.SectorCandidates)},
			{Name: "preserve pattern diagnostic firewall", Passed: a.PatternDiagnostics.Audited && containsAll(a.PatternDiagnostics.Supports, []string{StatusFNRelevant, StatusGJRelevant}) && containsAll(a.PatternDiagnostics.Failures, []string{StatusKoideNotNEff, StatusFNNotNativeRest, StatusGJNotLowScaleTop}), Detail: FormatAudit(a.PatternDiagnostics)},
			{Name: "re-audit D4 triality firewall", Passed: a.D4Firewall.Audited && containsAll(a.D4Firewall.Supports, []string{StatusTrialityMotivation}) && containsAll(a.D4Firewall.Failures, []string{StatusNEffNotD4, StatusNoTrialityTraceReadout, StatusNoTrialityRestPressure, StatusNoTrialityRealDescent}), Detail: FormatAudit(a.D4Firewall)},
			{Name: "audit finite triple top-color source", Passed: a.FiniteTriple.Audited && containsAll(a.FiniteTriple.Supports, []string{StatusFiniteTripleColorShape}) && containsAll(a.FiniteTriple.Failures, []string{StatusFSTNoTopEigenvalue, StatusFSTNoRestOperator}), Detail: FormatAudit(a.FiniteTriple)},
			{Name: "audit external ledger rest-pressure source", Passed: a.ExternalLedger.Audited && containsAll(a.ExternalLedger.Supports, []string{StatusExternalCanTest}) && containsAll(a.ExternalLedger.Failures, []string{StatusExternalNotNative}), Detail: FormatAudit(a.ExternalLedger)},
			{Name: "audit K7/projective resonance", Passed: a.K7Projective.Audited && containsAll(a.K7Projective.Failures, []string{StatusK7NotTopBlock, StatusProjectiveNotRest}), Detail: FormatAudit(a.K7Projective)},
			{Name: "audit complex D4 trilinear source", Passed: a.ComplexD4.Audited && containsAll(a.ComplexD4.Failures, []string{StatusTD4NotTraceMagnitude, StatusTD4NotTopDominance, StatusTD4NotRestPressure}), Detail: FormatAudit(a.ComplexD4)},
			{Name: "record C_Higgs impact", Passed: a.CHiggs.Recorded && math.Abs(a.CHiggs.TopLimit-CHistory) < 1e-15 && math.Abs(a.CHiggs.Current-CHiggs) < 1e-15 && math.Abs(a.CHiggs.DeltaCHiggs-0.0008046575187645733) < 1e-15 && math.Abs(a.CHiggs.TreeShift-TreeProxyShift) < 1e-10 && containsAll(a.CHiggs.Supports, []string{StatusRestPressureNumericallySmall}) && containsAll(a.CHiggs.Failures, []string{StatusTreeProxyNotPole, StatusNoNewSpectralData}), Detail: FormatCHiggs(a.CHiggs)},
			{Name: "select hierarchy-breaking obstruction", Passed: a.Hierarchy.Selected && strings.Contains(a.Hierarchy.Name, "HierarchyBreakingOperatorSeal") && containsAll(a.Hierarchy.Needs, []string{"one large colored eigenvalue", "suppressed rest spectrum", "small N_eff-3", "T", "q_rest"}) && containsAll(a.Hierarchy.Supports, []string{StatusNeedsHierarchyMechanism}) && containsAll(a.Hierarchy.Failures, []string{StatusNoHierarchyBreakingOperator, StatusNoNativeTopDominance, StatusNoLightSuppression, StatusNoNativeRestSource}), Detail: FormatHierarchy(a.Hierarchy)},
			{Name: "record outcome classification", Passed: a.Outcome.Recorded && len(a.Outcome.Items) == 6 && containsAll(a.Outcome.Items, []string{"exact N_eff=3", "positive rest spectral pressure", "narrow top-dominant positivity corridor", "C_Higgs remains Level B"}) && containsAll(a.Outcome.Supports, []string{StatusBestTypedModel}), Detail: FormatOutcome(a.Outcome)},
			{Name: "record branch decision", Passed: a.Branch.Recorded && strings.Contains(a.Branch.Next, "Gate 809") && strings.Contains(a.Branch.Next, "HierarchyBreakingOperatorSeal") && containsAll(a.Branch.Supports, []string{StatusNextHierarchyGate}), Detail: a.Branch.Next},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.NoNativeYukawa && a.Firewalls.NoPMNSCKM && a.Firewalls.NoFlavor && a.Firewalls.NoScalar && a.Firewalls.NoPoleMass && a.Firewalls.NoVEVGF && a.Firewalls.NoGJ && a.Firewalls.NoD4Triality && a.Firewalls.NoHistoryLoop && a.Firewalls.Verdict == StatusFirewallGate808, Detail: a.Firewalls.Verdict},
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
		notes := []string{a.Truth, FormatTop(a.TopColor), FormatRest(a.RestPressure), FormatCorridor(a.Corridor), FormatConcentration(a.Concentration), FormatAudit(a.SectorCandidates), FormatAudit(a.PatternDiagnostics), FormatAudit(a.D4Firewall), FormatAudit(a.FiniteTriple), FormatAudit(a.ExternalLedger), FormatAudit(a.K7Projective), FormatAudit(a.ComplexD4), FormatCHiggs(a.CHiggs), FormatHierarchy(a.Hierarchy), FormatOutcome(a.Outcome), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
