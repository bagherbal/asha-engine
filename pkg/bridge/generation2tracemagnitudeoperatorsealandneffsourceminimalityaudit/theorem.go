package generation2tracemagnitudeoperatorsealandneffsourceminimalityaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "GATE-807-TRACE-MAGNITUDE-OPERATOR-SEAL-AND-N-EFF-SOURCE"
	theoremName = "Gate 807 — TraceMagnitudeOperatorSeal and N_eff Source Audit"
)

func Generation2TraceMagnitudeOperatorSealAndNEffSourceMinimalityAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Notes: []string{err.Error()}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 806 and N_eff pressure", Passed: a.Inheritance.Gate806Inherited && strings.Contains(a.Inheritance.Formula, "C_Higgs") && math.Abs(a.Inheritance.NEff-NEff) < 1e-15 && containsAll(a.Inheritance.Verdicts, []string{StatusGate806Inherited, StatusNEffSubproblem, StatusUnitLeverageInherited}), Detail: a.Inheritance.UnitLeverage},
			{Name: "define TraceMagnitudeOperatorSeal", Passed: a.Seal.Defined && a.Seal.Name == "TraceMagnitudeOperatorSeal" && containsAll(a.Seal.Components, []string{"sector Hermitian operators H_u,H_d,H_e,H_nu", "positive spectra Spec(H_f)", "top-dominant block selector", "rest-pressure spectral measure", "scale and scheme convention"}) && containsAll(a.Seal.Supports, []string{StatusNeedsSpectraNotMixing}) && containsAll(a.Seal.Failures, []string{StatusSealNotNative}), Detail: FormatSeal(a.Seal)},
			{Name: "record trace formulas and inverse participation identity", Passed: a.Formulas.Recorded && strings.Contains(a.Formulas.A, "3H_u") && strings.Contains(a.Formulas.B, "3H_d²") && strings.Contains(a.Formulas.Identity, "N_eff=1/sum_i w_i²") && containsAll(a.Formulas.Verdicts, []string{StatusTraceFormulas, StatusParticipationIdentity}) && containsAll(a.Formulas.Supports, []string{StatusEffectiveTraceAtomCount}), Detail: FormatFormulas(a.Formulas)},
			{Name: "audit orientation invisibility", Passed: a.Orientation.Audited && strings.Contains(a.Orientation.Transformation, "U_f") && containsAll(a.Orientation.InvariantTraces, []string{"Tr(H_f)", "Tr(H_f²)", "N_eff"}) && containsAll(a.Orientation.Supports, []string{StatusSpectralMagnitudeOnly}) && containsAll(a.Orientation.Failures, []string{StatusNoPMNSCKMFromNEff, StatusNoKappaOrientFromTrace}), Detail: FormatOrientation(a.Orientation)},
			{Name: "define rank-three top-color block and exact limit", Passed: a.TopColor.Defined && math.Abs(a.TopColor.NEffTop-3) < 1e-12 && containsAll(a.TopColor.Verdicts, []string{StatusRankThreeBlock, StatusTopLimit}) && containsAll(a.TopColor.Supports, []string{StatusNearThreeTopColorBlock}) && containsAll(a.TopColor.Failures, []string{StatusTopColorNotGenerationTriality, StatusTopBlockNoT}), Detail: FormatTop(a.TopColor)},
			{Name: "derive rest-pressure formulas", Passed: a.Rest.Derived && strings.Contains(a.Rest.NEffFormula, "3(1+alpha)²/(1+beta)") && strings.Contains(a.Rest.DeltaFormula, "2alpha") && math.Abs(a.Rest.CurrentDelta-NEffDelta) < 1e-15 && containsAll(a.Rest.Supports, []string{StatusRestPressureAboveTop}) && containsAll(a.Rest.Failures, []string{StatusNoAlphaBetaWithoutT, StatusNoSectorRestAssignment}), Detail: FormatRest(a.Rest)},
			{Name: "reaffirm aggregate non-identifiability", Passed: a.NonIdentifiability.Reaffirmed && containsAll(a.NonIdentifiability.CannotFind, []string{"TraceMagnitudeOperatorSeal", "top channel T", "alpha,beta", "sector fractions"}) && containsAll(a.NonIdentifiability.Failures, []string{StatusABNoOperators, StatusABNoTopChannel, StatusABNoRestSectors}), Detail: FormatNonID(a.NonIdentifiability)},
			{Name: "audit finite triple trace template", Passed: a.FiniteTriple.Audited && containsAll(a.FiniteTriple.Supports, []string{StatusFSTSuppliesTraceShape}) && containsAll(a.FiniteTriple.Failures, []string{StatusFSTNoOperators}), Detail: FormatSource(a.FiniteTriple)},
			{Name: "audit external ledger magnitude source", Passed: a.External.Audited && containsAll(a.External.Supports, []string{StatusExternalCanPopulate}) && containsAll(a.External.Failures, []string{StatusExternalNotNative}), Detail: FormatSource(a.External)},
			{Name: "audit D4 trilinear magnitude source", Passed: a.TD4.Audited && containsAll(a.TD4.Failures, []string{StatusTD4NoMagnitudes, StatusTD4NoNEff}), Detail: FormatSource(a.TD4)},
			{Name: "audit K7/projective magnitude source", Passed: a.K7Projective.Audited && containsAll(a.K7Projective.Failures, []string{StatusK7NotMagnitudeOperator, StatusProjectiveNotNEff}), Detail: FormatSource(a.K7Projective)},
			{Name: "record scale locality", Passed: a.Scale.Audited && strings.Contains(a.Scale.Differential, "2 d ln a - d ln b") && containsAll(a.Scale.Verdicts, []string{StatusScaleAudited, StatusScaleDifferential}) && containsAll(a.Scale.Failures, []string{StatusNoScaleStability, StatusMZScaleSealed}), Detail: FormatScale(a.Scale)},
			{Name: "audit C_Higgs impact", Passed: a.CHiggs.Audited && strings.Contains(a.CHiggs.Formula, "3/N_eff") && containsAll(a.CHiggs.Supports, []string{StatusImprovesCYukawaTestability}) && containsAll(a.CHiggs.Failures, []string{StatusSealAloneNoNativeCHiggs, StatusCHiggsLevelB}), Detail: FormatCHiggs(a.CHiggs)},
			{Name: "record outcome classification", Passed: a.Outcome.Recorded && len(a.Outcome.Items) == 6 && containsAll(a.Outcome.Items, []string{"Hermitian trace spectra", "rank-three top-color", "rest spectral pressure", "aggregate a,b do not identify", "C_Higgs remains Level B"}) && containsAll(a.Outcome.Supports, []string{StatusMagnitudeSharper}), Detail: FormatOutcome(a.Outcome)},
			{Name: "record branch decision", Passed: a.Branch.Recorded && strings.Contains(a.Branch.Next, "Gate 808") && strings.Contains(a.Branch.Next, "RankThreeTopColorBlock") && containsAll(a.Branch.Supports, []string{StatusNextTopRestPressure}), Detail: a.Branch.Next},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.NoPMNSCKM && a.Firewalls.NoYukawa && a.Firewalls.NoEigenvalues && a.Firewalls.NoFlavor && a.Firewalls.NoScalar && a.Firewalls.NoPoleMass && a.Firewalls.NoVEVGF && a.Firewalls.NoGJ && a.Firewalls.NoTriality && a.Firewalls.NoHistoryLoop && a.Firewalls.Verdict == StatusFirewallGate807, Detail: a.Firewalls.Verdict},
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
		notes := []string{a.Truth, FormatSeal(a.Seal), FormatFormulas(a.Formulas), FormatOrientation(a.Orientation), FormatTop(a.TopColor), FormatRest(a.Rest), FormatNonID(a.NonIdentifiability), FormatSource(a.FiniteTriple), FormatSource(a.External), FormatSource(a.TD4), FormatSource(a.K7Projective), FormatScale(a.Scale), FormatCHiggs(a.CHiggs), FormatOutcome(a.Outcome), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
