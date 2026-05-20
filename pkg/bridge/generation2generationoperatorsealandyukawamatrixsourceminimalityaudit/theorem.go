package generation2generationoperatorsealandyukawamatrixsourceminimalityaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "GATE-806-GENERATION-OPERATOR-SEAL-YUKAWA-MATRIX-SOURCE-MINIMALITY"
	theoremName = "Gate 806 — GenerationOperatorSeal and Yukawa Matrix Source Minimality Audit"
)

func Generation2GenerationOperatorSealAndYukawaMatrixSourceMinimalityAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Notes: []string{err.Error()}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 805 edge-triality no-go", Passed: a.Inheritance.Gate805NoGo && strings.Contains(a.Inheritance.FactorizedNormal, "Edge_f ⊗ Y_f") && containsAll(a.Inheritance.Verdicts, []string{StatusGate805Inherited, StatusFactorizationRecorded, StatusEdgeNoYF}), Detail: a.Inheritance.FactorizedNormal},
			{Name: "define GenerationOperatorSeal", Passed: a.Seal.Defined && a.Seal.Name == "GenerationOperatorSeal" && containsAll(a.Seal.Components, []string{"G_gen", "sector Yukawa operators Y_u,Y_d,Y_e,Y_nu", "Hermitian trace operators H_f=Y_f†Y_f", "singular-value spectrum", "diagonalization frames", "PMNS/CKM misalignment readouts", "hierarchy/breaking operator", "color multiplicity rule", "neutrino convention", "noncircularity proof"}) && containsAll(a.Seal.Supports, []string{StatusYukawaRequiresOperators}) && containsAll(a.Seal.Failures, []string{StatusNoNativeGenerationOperator}), Detail: FormatSeal(a.Seal)},
			{Name: "audit minimality of generation operator subobjects", Passed: a.Minimality.Audited && len(a.Minimality.Items) >= 10 && containsAll([]string{FormatMinimality(a.Minimality)}, []string{"remove G_gen", "remove Y_f", "remove H_f", "remove diagonalization frames", "remove hierarchy/breaking operator", "remove noncircularity"}) && containsAll(a.Minimality.Supports, []string{StatusSubobjectsNoncosmetic}) && containsAll(a.Minimality.Failures, []string{StatusCannotCompress}), Detail: FormatMinimality(a.Minimality)},
			{Name: "separate magnitude and orientation layers", Passed: a.Layers.Separated && len(a.Layers.Layers) == 2 && containsAll(a.Layers.Supports, []string{StatusNEffNeedsSpectra, StatusKappaOrientNeedsFrames}) && containsAll(a.Layers.Failures, []string{StatusSingularNoMixing}), Detail: FormatLayers(a.Layers)},
			{Name: "audit finite spectral triple source", Passed: a.FST.Audited && containsAll(a.FST.Supplies, []string{"sector edge skeleton", "trace-form templates"}) && containsAll(a.FST.Missing, []string{"Y_f entries", "generation carrier", "eigenvalues", "mixing frames"}) && containsAll(a.FST.Supports, []string{StatusFSTSuppliesEdgeDomain}) && containsAll(a.FST.Failures, []string{StatusFSTNoYF}), Detail: FormatSource(a.FST)},
			{Name: "audit T_D4 source", Passed: a.TD4.Audited && containsAll(a.TD4.Supplies, []string{"airlocked trilinear kernel shape"}) && containsAll(a.TD4.Failures, []string{StatusTD4NoGenOperator, StatusTD4NoTraceAtoms}), Detail: FormatSource(a.TD4)},
			{Name: "audit aggregate trace ledger source", Passed: a.Aggregate.Audited && containsAll(a.Aggregate.Supplies, []string{"sealed aggregate trace values"}) && containsAll(a.Aggregate.Failures, []string{StatusAggregateNoOperator}), Detail: FormatSource(a.Aggregate)},
			{Name: "audit external ledger source", Passed: a.External.Audited && containsAll(a.External.Supports, []string{StatusExternalCanPopulateSeal}) && containsAll(a.External.Failures, []string{StatusExternalNotNative}), Detail: FormatSource(a.External)},
			{Name: "audit K7/projective candidates", Passed: a.K7Projective.Audited && containsAll(a.K7Projective.Supports, []string{StatusK7ProjectiveCandidates}) && containsAll(a.K7Projective.Failures, []string{StatusK7NotOperator, StatusProjectiveNotSource}), Detail: FormatSource(a.K7Projective)},
			{Name: "record Yukawa edge times generation operator normal form", Passed: a.NormalForm.Recorded && containsAll(a.NormalForm.Forms, []string{"D_u = Edge_u ⊗ Y_u", "D_d = Edge_d ⊗ Y_d", "D_e = Edge_e ⊗ Y_e", "D_nu = Edge_nu ⊗ Y_nu"}) && containsAll(a.NormalForm.Supports, []string{StatusFSTWhereNotWhat}) && containsAll(a.NormalForm.Failures, []string{StatusNoNativeWhatOperator}), Detail: FormatNormal(a.NormalForm)},
			{Name: "audit trace readout obstruction", Passed: a.Trace.Audited && containsAll(a.Trace.Failures, []string{StatusCarrierAloneNoSpectra, StatusGenerationsAloneNoNEff, StatusNoTraceAtomsWithoutH}), Detail: FormatObstruction(a.Trace)},
			{Name: "audit mixing readout obstruction", Passed: a.Mixing.Audited && containsAll(a.Mixing.Failures, []string{StatusTraceLedgerNoKappaOrient, StatusNoPMNSCKMFrames}), Detail: FormatObstruction(a.Mixing)},
			{Name: "audit hierarchy breaking obstruction", Passed: a.Hierarchy.Audited && containsAll(a.Hierarchy.Failures, []string{StatusNoTopDominanceOperator, StatusNoLightSuppressionOperator, StatusNoNEffMinusThreeSource}), Detail: FormatObstruction(a.Hierarchy)},
			{Name: "preserve C_Higgs firewall", Passed: a.CHiggs.Preserved && strings.Contains(a.CHiggs.Formula, "C_Higgs") && containsAll(a.CHiggs.Unchanged, []string{"N_eff", "C_Yukawa", "C_History", "C_Higgs"}) && containsAll(a.CHiggs.Failures, []string{StatusNoCYukawaUpdate, StatusCHiggsLevelB}), Detail: FormatCHiggs(a.CHiggs)},
			{Name: "record outcome classification", Passed: a.Outcome.Recorded && len(a.Outcome.Items) == 6 && containsAll(a.Outcome.Items, []string{"lawful Yukawa edge locations", "actual Yukawa matrices require independent generation-sector operators", "trace magnitude and mixing orientation", "current ASHA does not supply the native GenerationOperatorSeal", "C_Higgs remains Level B"}) && containsAll(a.Outcome.Supports, []string{StatusGenerationOperatorBottleneck}), Detail: FormatOutcome(a.Outcome)},
			{Name: "record branch decision", Passed: a.Branch.Recorded && strings.Contains(a.Branch.Next, "Gate 807") && strings.Contains(a.Branch.Next, "TraceMagnitudeOperatorSeal") && containsAll(a.Branch.Supports, []string{StatusNextTraceMagnitude}), Detail: a.Branch.Next},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.NoYukawa && a.Firewalls.NoEigenvalues && a.Firewalls.NoPMNSCKM && a.Firewalls.NoFlavor && a.Firewalls.NoNEff && a.Firewalls.NoGJ && a.Firewalls.NoScalar && a.Firewalls.NoPoleMass && a.Firewalls.NoVEVGF && a.Firewalls.NoTriality && a.Firewalls.NoHistoryLoop && a.Firewalls.Verdict == StatusFirewallGate806, Detail: a.Firewalls.Verdict},
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
		notes := []string{a.Truth, FormatSeal(a.Seal), FormatMinimality(a.Minimality), FormatLayers(a.Layers), FormatSource(a.FST), FormatSource(a.TD4), FormatSource(a.Aggregate), FormatSource(a.External), FormatSource(a.K7Projective), FormatNormal(a.NormalForm), FormatObstruction(a.Trace), FormatObstruction(a.Mixing), FormatObstruction(a.Hierarchy), FormatCHiggs(a.CHiggs), FormatOutcome(a.Outcome), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
