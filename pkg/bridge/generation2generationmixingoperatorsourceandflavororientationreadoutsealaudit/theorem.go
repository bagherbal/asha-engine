package generation2generationmixingoperatorsourceandflavororientationreadoutsealaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2GenerationMixingOperatorSourceAndFlavorOrientationReadoutSealAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 789 — Generation Mixing Operator Source and FlavorOrientationReadoutSeal Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}, Notes: []string{StatusFirewallPreservedGate789}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate788 flavor-orientation bottleneck", Passed: a.Gate788.Inherited && a.Gate788.KappaOrientFocus && a.Gate788.Formula == "sin^2(theta13)/4 - J_CKM" && strings.Contains(a.Gate788.PMNSTyping, "PMNS") && strings.Contains(a.Gate788.CKMTyping, "J_CKM"), Detail: a.Gate788.Formula},
			{Name: "define required generation-mixing objects", Passed: a.Required.Defined && containsAll(a.Required.RequiredObjects, []string{"generation carrier G_gen", "sector Yukawa or mass operators on G_gen", "typed diagonalization maps", "misalignment unitaries between sectors", "readout maps theta13 and J_CKM", "orientation/sign convention explaining sin^2(theta13)/4 - J_CKM"}) && strings.Contains(a.Required.LeptonMisalignment, "U_PMNS") && strings.Contains(a.Required.QuarkMisalignment, "V_CKM") && a.Required.SectorMisalignmentNeed && !a.Required.NativeUPMNSOrVCKMExists, Detail: a.Required.Verdict},
			{Name: "audit Yukawa trace pair", Passed: a.YukawaTrace.Audited && closeRel(a.YukawaTrace.NEff, nEffSnapshot, 1e-15) && a.YukawaTrace.AggregateParticipation && !a.YukawaTrace.DeterminesPMNSOrCKM && !a.YukawaTrace.SuppliesEigenvectorMisalignment, Detail: a.YukawaTrace.Verdict},
			{Name: "audit Yukawa singular-value ledger", Passed: a.SingularLedger.Audited && a.SingularLedger.SingularValuesCanSourceTraces && !a.SingularLedger.SingularValuesDetermineMixing && !a.SingularLedger.NativeEigenvectorOrientation, Detail: a.SingularLedger.Verdict},
			{Name: "audit finite spectral triple flavor source", Passed: a.FiniteTriple.Audited && a.FiniteTriple.AllowedYukawaEdgeShapes && !a.FiniteTriple.GenerationMixingOperatorSourced, Detail: a.FiniteTriple.Verdict},
			{Name: "audit K7 Hodge polarity source", Passed: a.K7Polarity.Audited && strings.Contains(a.K7Polarity.Split, "4") && strings.Contains(a.K7Polarity.Split, "3") && a.K7Polarity.SelectorResonance && !a.K7Polarity.DefinesGenerationMixingOperator && !a.K7Polarity.QuarterWeightDerivesTheta13, Detail: a.K7Polarity.Verdict},
			{Name: "audit Fock/projective selector source", Passed: a.FockSelector.Audited && containsAll(a.FockSelector.Patterns, []string{"4 = 1 + 3", "CP3", "projective selector patterns"}) && a.FockSelector.FutureGenerationCandidate && !a.FockSelector.TypedSelectorToPMNSCKMMap, Detail: a.FockSelector.Verdict},
			{Name: "audit triality generation carrier candidate", Passed: a.Triality.Audited && a.Triality.ThreefoldRelevantCandidate && !a.Triality.SuppliesSectorOperators && !a.Triality.SuppliesRelativeOrientations && !a.Triality.SuppliesPhaseData && !a.Triality.SuppliesMixingReadoutMaps && !a.Triality.SectorMisalignmentOperatorFound, Detail: a.Triality.Verdict},
			{Name: "audit boundary data source", Passed: a.BoundaryData.Audited && containsAll(a.BoundaryData.Coordinates, []string{"s", "xi_boundary", "lambda(Lambda12)", "R3-1"}) && a.BoundaryData.SmallCorrectionToReadout && !a.BoundaryData.DerivesFlavorMixing, Detail: a.BoundaryData.Verdict},
			{Name: "define GenerationMixingOperatorSeal", Passed: a.Seal.Defined && a.Seal.Name == "GenerationMixingOperatorSeal" && containsAll(a.Seal.Components, []string{"G_gen", "U_PMNS", "V_CKM", "readout maps theta13 and J_CKM", "orientation/sign convention"}) && strings.Contains(a.Seal.Readout, "sin^2(theta13)/4 - J_CKM") && !a.Seal.Native, Detail: FormatSeal(a.Seal)},
			{Name: "audit seal minimality", Passed: a.Minimality.Audited && a.Minimality.Minimal && strings.Contains(a.Minimality.RemoveEffects["G_gen"], "no generation carrier") && strings.Contains(a.Minimality.RemoveEffects["U_PMNS"], "theta13") && strings.Contains(a.Minimality.RemoveEffects["V_CKM"], "J_CKM"), Detail: a.Minimality.Verdict},
			{Name: "audit runtime target absence", Passed: a.Runtime.Audited && containsAll(a.Runtime.ForbiddenDirectVariables, []string{"lambda_runtime", "m_H_tree", "C_Higgs", "G_F", "v"}) && !a.Runtime.ContainsForbidden && a.Runtime.RuntimeTargetIndependent && !a.Runtime.TheoremLevelIndependent, Detail: a.Runtime.Verdict},
			{Name: "record status propagation", Passed: a.Propagation.Recorded && strings.Contains(a.Propagation.KappaOrient, "GenerationMixingOperatorSeal") && strings.Contains(a.Propagation.KappaERed, "mixed flavor-boundary") && strings.Contains(a.Propagation.FWall3, "Level B+") && strings.Contains(a.Propagation.CHistory, "Level B") && strings.Contains(a.Propagation.CHiggs, "not Level C"), Detail: a.Propagation.Verdict},
			{Name: "record branch decision", Passed: a.Branch.Recorded && strings.Contains(a.Branch.SuccessBranch, "Native Generation Mixing") && strings.Contains(a.Branch.FailureBranch, "C_Higgs Dependency Freeze") && a.Branch.Selected == "failure branch", Detail: FormatBranch(a.Branch)},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Enforced && !a.Firewalls.NEffPMNSCKMTheorem && !a.Firewalls.YukawaSingularValuesMixingTheorem && !a.Firewalls.K7PolarityMixingTheorem && !a.Firewalls.RadialQuarterTheta13Theorem && !a.Firewalls.ProjectiveSelectorPMNSCKMTheorem && !a.Firewalls.TrialityPMNSCKMTheorem && !a.Firewalls.BoundaryPairFlavorMixingTheorem && !a.Firewalls.KappaOrientNativeFlavorTheorem && !a.Firewalls.FlavorOrientationSealNative && !a.Firewalls.TreeProxyPoleMass && a.Firewalls.Verdict == StatusFirewallPreservedGate789, Detail: a.Firewalls.Verdict},
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
		notes := append([]string{a.Truth, FormatSeal(a.Seal), FormatBranch(a.Branch), a.FinalStatement}, Statuses()...)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
