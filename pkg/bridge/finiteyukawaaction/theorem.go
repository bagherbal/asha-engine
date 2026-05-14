package finiteyukawaaction

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func FiniteYukawaActionFunctionalTrialityHopfAmplitudeQualificationAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-FINITE-YUKAWA-ACTION-FUNCTIONAL-TRIALITY-HOPF-AMPLITUDE-QUALIFICATION-AUDIT"
	const name = "Finite Yukawa Action Functional / Triality-Hopf Amplitude Qualification Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{{Name: "build Gate 263 finite Yukawa action audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 262 Hermitian triality mixing basis is inherited", Passed: a.Inheritance.BilinearCarrierDefined && a.Inheritance.HermitianTrialityBasisExposed && a.Inheritance.RawNonCommutingPartnerExists && !a.Inheritance.PreviousQualifiedPartnerFound, Detail: FormatInheritance(a.Inheritance)},
			{Name: "exact M3(C) trace diagnostics evaluate the real and phase bases", Passed: a.Summary.TraceFunctionalsEvaluated && len(a.TraceAudits) == 3 && a.TraceAudits[1].RealBasisValue == 6 && a.TraceAudits[1].PhaseBasisValue == 6 && a.TraceAudits[1].CrossValue == 0, Detail: FormatTraceAudit(a.TraceAudits[1])},
			{Name: "trace and commutator diagnostics do not select Yukawa amplitudes", Passed: a.Summary.TraceMetricDegenerate && !anyTraceSelectsAmplitude(a.TraceAudits) && a.TraceAudits[2].RealBasisValue == a.TraceAudits[2].PhaseBasisValue, Detail: FormatTraceAudit(a.TraceAudits[2])},
			{Name: "native action candidates are audited without promotion", Passed: a.Summary.NativeActionCandidateCount == 5 && !a.Summary.ActionCandidateQualified && !anyQualifiedAction(a.ActionCandidates), Detail: "candidates=" + fmtActionCandidates(a.ActionCandidates)},
			{Name: "B_gap and Hopf phase do not yet integrate into the M3 off-diagonal basis", Passed: a.ScalarPhase.BGapAvailableAsScale && !a.ScalarPhase.BGapCanWeightTrialityBasis && a.ScalarPhase.HopfPhaseLedgerAvailable && !a.ScalarPhase.HopfCanFixCPPhase && !a.ScalarPhase.ScalarPhaseIntegrationDerived, Detail: FormatScalarPhase(a.ScalarPhase)},
			{Name: "physical Yukawa texture remains an ansatz with unselected coefficients", Passed: a.Texture.DiagonalTauSourceAvailable && a.Texture.HermitianOffDiagonalBasisExists && a.Texture.TraceMetricAvailable && !a.Texture.FiniteActionCoefficientRule && !a.Texture.PhysicalTextureConstructed && a.Texture.EmpiricalYukawaSealRequired, Detail: FormatTexture(a.Texture)},
			{Name: "CKM/PMNS, fermion masses, and empirical Yukawa seal remain blocked", Passed: !a.Summary.CKMPMNSDerived && !a.Summary.FermionMassesDerived && a.Firewall.EmpiricalYukawaSealPreserved && !a.Firewall.FiniteCorePolluted, Detail: FormatFirewall(a.Firewall)},
			{Name: "firewall prevents trace metrics, B_gap, Hopf phases, or symmetry bases from becoming amplitudes", Passed: a.Firewall.DoesNotPromoteTraceMetricToDynamics && a.Firewall.DoesNotPromoteSymmetryToAmplitude && a.Firewall.DoesNotUseBGapWithoutMap && a.Firewall.DoesNotUseHopfWithoutProjection && a.Firewall.DoesNotClaimSpectralTripleComplete, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{
			a.TruthStatement,
			"Gate 263 evaluates the finite trace/action landscape around the Gate 262 Hermitian triality basis, but no native action functional assigns physical Yukawa amplitudes.",
			"The lawful texture ansatz is exposed; activating it requires either a future finite D_F/order-one action selector or an explicitly quarantined EmpiricalYukawaSeal.",
		}}
	}}
}

func fmtActionCandidates(c []NativeActionCandidate) string {
	return "[" + strings.Join(candidateNames(c), ", ") + "]"
}
