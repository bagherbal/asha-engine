package generation2movingu1phaselineandhyperchargenormalizationfirewallaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2MovingU1PhaseLineAndHyperchargeNormalizationFirewallAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: AuditID, Name: "Gate 717 — Moving U(1) Phase Line and Hypercharge Normalization Firewall Audit", Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: AuditID, Name: "Gate 717 — Moving U(1) Phase Line and Hypercharge Normalization Firewall Audit", Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate716 SU2-side airlock", Passed: a.Inherited.SU2AirlockInherited && a.Inherited.SU2SideStructurallyReady && a.Inherited.InternalCSelectorIndependent && a.Inherited.ComplexCarrierSelectorDependent && a.Inherited.U1PhaseSelectorDependent && !a.Inherited.HyperchargeDerived && !a.Inherited.HyperchargeNormalized && !a.Inherited.FullHiggsDoubletMap && !a.Inherited.HiggsMassCertified && !a.Inherited.YukawaCertified && a.Inherited.Verdict == StatusGate716SU2IntertwinerAirlockInherited, Detail: FormatInherited(a.Inherited)},
			{Name: "define moving phase line", Passed: strings.Contains(a.PhaseLine.Definition, "L_n") && a.PhaseLine.Dimension == movingPhaseLineDimension && a.PhaseLine.DependsOnSelectorN && a.PhaseLine.FixedJHRequired && !a.PhaseLine.SelectorIndependent && !a.PhaseLine.PhysicalHypercharge && strings.Contains(a.PhaseLine.Verdict, StatusMovingPhaseLineDefined), Detail: FormatPhaseLine(a.PhaseLine)},
			{Name: "audit centrality in fixed U2 socket", Passed: strings.Contains(a.Central.Commutator, "[J_H") && a.Central.CommutesWithC && a.Central.LiesInCenterOfU2 && a.Central.FixedJHOnly && !a.Central.PhysicalU1Y && strings.Contains(a.Central.Verdict, StatusLNIsCentralInU2SocketForFixedJH) && strings.Contains(a.Central.Verdict, StatusLNInternalU1PhaseSocketAfterJHChoice), Detail: FormatCentral(a.Central)},
			{Name: "audit uniform phase action", Passed: a.Uniform.Generator == "J_H(n)" && a.Uniform.ActsAsMultiplicationByI && strings.Contains(a.Uniform.ExponentialAction, "exp") && a.Uniform.UniformOnFullC2 && a.Uniform.ComplexDimension == k7PlusComplexDimension && !a.Uniform.PhysicalChargeFixed && strings.Contains(a.Uniform.Verdict, StatusJHExponentiatesToUniformPhaseOnC2), Detail: FormatUniform(a.Uniform)},
			{Name: "audit charge normalization ambiguity", Passed: a.Charge.PhaseLineFixed && a.Charge.NaturalDirection == "J_H(n)" && len(a.Charge.CandidateNormalizations) == 3 && a.Charge.SameLineDifferentCharges && !a.Charge.PhysicalHyperchargeNormalization && a.Charge.ThetaYRequired && strings.Contains(a.Charge.Verdict, StatusChargeNormalizationAudited) && strings.Contains(a.Charge.Verdict, StatusNoHyperchargeAssignmentOrNormalization), Detail: FormatCharge(a.Charge)},
			{Name: "audit selector dependence", Passed: a.Selector.PhaseLineDependsOnN && !a.Selector.NativeTwistorPointSelector && !a.Selector.SelectorIndependentU1Line && !a.Selector.ComplexStructureSelected && strings.Contains(a.Selector.Verdict, StatusSelectorDependenceAudited) && strings.Contains(a.Selector.Verdict, StatusNoNativeTwistorPointSelector), Detail: FormatSelector(a.Selector)},
			{Name: "record SU2/U1 asymmetry", Passed: strings.Contains(a.Asymmetry.SU2Side, "twistor-invariant") && strings.Contains(a.Asymmetry.U1Side, "moves") && a.Asymmetry.SU2SelectorIndependent && a.Asymmetry.U1SelectorDependent && a.Asymmetry.U1NormalizationOpen && strings.Contains(a.Asymmetry.Verdict, StatusSU2U1AsymmetryRecorded), Detail: FormatAsymmetry(a.Asymmetry)},
			{Name: "enforce physical hypercharge firewall", Passed: !a.Physical.LnPhysicalU1Y && !a.Physical.JHHyperchargeGenerator && !a.Physical.InternalPhaseChargePhysicalHiggsHypercharge && !a.Physical.FullPhysicalHiggsDoublet && !a.Physical.HyperchargeAssignment && !a.Physical.HyperchargeNormalization && !a.Physical.HiggsMass && !a.Physical.ScalarRuntime && !a.Physical.YukawaOperator && !a.Physical.YukawaEigenvalues && len(a.Physical.MissingMaps) == 3 && strings.Contains(a.Physical.Verdict, StatusPhysicalHyperchargeFirewallEnforced) && strings.Contains(a.Physical.Verdict, StatusGate717MovingU1PhaseHyperchargeBoundary), Detail: FormatPhysical(a.Physical)},
		}
		notes := append(Statuses(), a.Truth)
		return theorem.Result{ID: AuditID, Name: "Gate 717 — Moving U(1) Phase Line and Hypercharge Normalization Firewall Audit", Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
