package generation2decomposedyukawatraceledgerandneffscalestabilityaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "GATE-793-DECOMPOSED-YUKAWA-TRACE-LEDGER-N-EFF-SCALE-STABILITY"
	theoremName = "Gate 793 — Decomposed Yukawa Trace Ledger and N_eff Source-Stability Audit"
)

func Generation2DecomposedYukawaTraceLedgerAndNEffScaleStabilityAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 793 analysis", Passed: false, Detail: err.Error()}}, Notes: []string{StatusFirewallPreservedGate793}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 792 error-budget result", Passed: a.Gate792.Inherited && a.Gate792.NEffTopNumericalLeverage && a.Gate792.Verdict == StatusGate792Inherited, Detail: a.Gate792.Verdict},
			{Name: "record trace-atom participation identity", Passed: a.TraceAtom.Recorded && closeAbs(a.TraceAtom.Ratio, ratioSnapshot, 1e-16) && closeAbs(a.TraceAtom.NEff, nEffSnapshot, 5e-16) && strings.Contains(a.TraceAtom.Formula, "N_eff"), Detail: FormatTraceAtom(a.TraceAtom)},
			{Name: "define sector decomposition requirement", Passed: a.Sector.Defined && !a.Sector.SectorTracesAvailable && a.Sector.MissingSeal == "DecomposedYukawaTraceLedgerSeal" && containsAll(a.Sector.RequiredQuadratic, []string{"a_u", "a_d", "a_e", "a_nu"}) && containsAll(a.Sector.RequiredQuartic, []string{"b_u", "b_d", "b_e", "b_nu"}), Detail: a.Sector.MissingSeal},
			{Name: "inherit top-color dominance limit", Passed: a.TopColor.Inherited && closeAbs(a.TopColor.RatioTop, 1.0/3.0, 1e-16) && closeAbs(a.TopColor.NEffTop, 3, 1e-15) && closeAbs(a.TopColor.DeltaRatio, -0.0002583937062663466, 1e-15) && closeAbs(a.TopColor.NEffMinusThree, 0.0023273474722147, 1e-15), Detail: FormatTopColor(a.TopColor)},
			{Name: "inherit top-rest decomposition formula", Passed: a.TopRest.FormulaInherited && !a.TopRest.TypedTopChannelAvailable && !a.TopRest.DecomposedLedgerAvailable && strings.Contains(a.TopRest.FormulaRatio, "beta") && a.TopRest.Verdict == StatusNoAlphaBetaWithoutTopLedger, Detail: a.TopRest.FormulaDelta},
			{Name: "audit generation participation", Passed: a.Generation.Audited && !a.Generation.GenerationCarrierCertified && !a.Generation.GenerationResolvedTraceLedger && containsAll(a.Generation.RequiredObjects, []string{"G_gen", "generation-resolved", "a,b"}), Detail: a.Generation.Verdict},
			{Name: "define D4/triality candidate requirements", Passed: a.D4.RequirementsDefined && a.D4.RealFormFirewallAudited && a.D4.StrongFutureCandidate && !a.D4.CurrentCertified && containsAll(a.D4.RequiredPackage, []string{"D4", "triality", "trace-readout", "real-form"}), Detail: strings.Join(a.D4.RequiredPackage, "; ")},
			{Name: "record N_eff scale-stability requirements", Passed: a.Scale.RequirementsDefined && a.Scale.DifferentialRecorded && a.Scale.Scale == "M_Z" && strings.Contains(a.Scale.Differential, "d ln N_eff") && !a.Scale.MultiScaleLedgerAvailable, Detail: FormatScale(a.Scale)},
			{Name: "record N_eff baseline impact on C_Higgs", Passed: a.Impact.Recorded && closeAbs(a.Impact.CHiggsTopColor, cHistorySnapshot, 1e-15) && closeAbs(a.Impact.DeltaCHiggs, 0.0008046575187645733, 1e-16) && closeAbs(a.Impact.DeltaTreeProxy, 0.04862437568908, 5e-14), Detail: FormatImpact(a.Impact)},
			{Name: "classify possible sources of three", Passed: a.SourceClassification.Completed && a.SourceClassification.TopColorCurrent && !a.SourceClassification.GenerationCertified && !a.SourceClassification.D4Current && a.SourceClassification.AggregateSealedCurrent && containsAll(a.SourceClassification.Sources, []string{"top-color", "generation", "D4", "aggregate"}), Detail: strings.Join(a.SourceClassification.Sources, "; ")},
			{Name: "audit symbolic-pattern firewall", Passed: a.Symbolic.Audited && a.Symbolic.D4MotivationOnly && !a.Symbolic.SymbolicPatternEvidence && a.Symbolic.Verdict == StatusSymbolicD4MotivationOnly, Detail: a.Symbolic.Verdict},
			{Name: "record branch decision", Passed: a.Branch.Recorded && !a.Branch.DecomposedLedgerAvailable && !a.Branch.D4PackageIntroduced && strings.Contains(a.Branch.Recommended, "DecomposedYukawaTraceLedgerSeal") && containsAll(a.Branch.Alternatives, []string{"Sector Contribution", "D4 Triality"}), Detail: FormatBranch(a.Branch)},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Enforced && !a.Firewalls.NEffGenerationTheorem && !a.Firewalls.NEffD4TrialityTheorem && !a.Firewalls.TopColorGeneration && !a.Firewalls.D4ResonanceReadoutTheorem && !a.Firewalls.Spin8AutomaticNative && !a.Firewalls.SymbolicProof && !a.Firewalls.ScaleStabilityAssumed && !a.Firewalls.CHiggsPoleMass && !a.Firewalls.TreeProxyShiftPoleCorrection && a.Firewalls.Verdict == StatusFirewallPreservedGate793, Detail: a.Firewalls.Verdict},
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
		notes := append([]string{a.Truth, FormatTraceAtom(a.TraceAtom), FormatTopColor(a.TopColor), FormatImpact(a.Impact), FormatBranch(a.Branch), a.FinalStatement}, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
