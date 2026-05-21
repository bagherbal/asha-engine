package generation2leptocolorfiniteactionprojectorledgeraudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "GATE-838-LEPTOCOLOR-FINITE-ACTION-PROJECTOR-LEDGER"
	theoremName = "Gate 838 — LeptoColor Finite Representation Action and ProjectorLedger Audit"
)

func Generation2LeptoColorFiniteActionProjectorLedgerAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate 838 lepto-color finite action/projector ledger audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit R2++ ledger and Gate 837 lepto-color carrier seal", Passed: a.Ledger.R2PlusPlus && !a.Ledger.R3 && !a.Ledger.R4 && !a.Ledger.AlphaIsNative && a.Ledger.OfficialFrozen && math.Abs(a.Ledger.OperatorNEff-OperatorNEff) < 5e-16 && containsAll(a.Impact.Verdicts, []string{StatusGate837Inherited, StatusR2PlusPlusRetained}), Detail: FormatLedger(a.Ledger)},
			{Name: "reverify W=C_lepton plus C^3_color with P1/P3 and internal B-L", Passed: a.W.Dim == WDim && a.W.P1Rank == LeptonBlockDim && a.W.P3Rank == ColorBlockDim && a.W.P1P3Orthogonal && a.W.P1PlusP3CompletesW && a.W.BMinusLTraceZero && a.W.P3WIsM3Fundamental && a.W.ColorBlockCanonical && !a.W.ColorAtomsCanonical && !a.W.CanonicalColorFrame && containsAll(a.W.Supports, []string{SupportSharedLeptoColorCarrierInherited, SupportP3WAsM3Fundamental, SupportBMinusLInternalToW}) && containsAll(a.W.Failures, []string{FailureNoCanonicalColorAtomFrame, FailureM3MatrixUnitsBasisDependent}), Detail: FormatW(a.W)},
			{Name: "audit E=C_R^2 plus C_L^2 right/left socket roles without particle assignment", Passed: a.E.Dim == ElectroweakSlotDim && a.E.RightSlotDim == RightSlotDim && a.E.LeftSlotDim == LeftSlotDim && a.E.RightSocketPairDeclared && a.E.LeftDoubleSocketDeclared && a.E.QuaternionicWeakRoleDeclared && a.E.ComplexRightRoleDeclared && a.E.SourceTypingCertifiedAsSeal && !a.E.ObservedParticleNamesUsed && containsAll(a.E.Supports, []string{SupportESlotRightLeftSocketBody}), Detail: FormatE(a.E)},
			{Name: "certify schematic rho_F action consistency on sealed carrier", Passed: a.Rho.M3ActsOnP3W && a.Rho.M3TrivialOnP1W && a.Rho.HActsOnLeftDoubleSocket && a.Rho.CActsOnRightSocketPair && a.Rho.ActionPreservesP1P3 && a.Rho.ActionPreservesBMinusL && a.Rho.RepresentationLawConsistentAtBlockLevel && !a.Rho.NativeDerivationCertified && !a.Rho.ExplicitMatricesCertified && !a.Rho.FirstOrderConditionCertified && !a.Rho.BimoduleCommutantProof && !a.Rho.CompleteRhoFActionLedger && containsAll(a.Rho.Supports, []string{SupportRhoFActionSeal}) && containsAll(a.Rho.Failures, []string{FailureRepresentationSealNotNative, FailureNoFullFiniteTripleProof, FailureNoExplicitMatrices, FailureNoFirstOrderProof, FailureNoBimoduleCommutantProof}), Detail: FormatRho(a.Rho)},
			{Name: "construct coarse particle-side projector ledger and J-copy rank doubling", Passed: len(a.Projectors.ParticleProjectors) == 4 && a.Projectors.ParticleRankSum == HPartDim && a.Projectors.ExpectedHPartDim == HPartDim && a.Projectors.HFProjectorRankSum == HFDim && a.Projectors.ExpectedHFDim == HFDim && a.Projectors.Orthogonal && a.Projectors.CompleteOnHPart && a.Projectors.JCopyIncluded && a.Projectors.CoarsePiSectorFSealCertified && !a.Projectors.FullNativePiSectorFCertified && !a.Projectors.FineColorAtomLedgerCertified && !a.Projectors.CanonicalColorFrameCertified && !a.Projectors.TraceMagnitudeReadoutCertified && containsAll(a.Projectors.Supports, []string{SupportCoarsePiSectorFSeal, SupportJCopyFiniteBody}) && containsAll(a.Projectors.Failures, []string{FailureNoTraceMagnitudeReadout, FailureNoFineColorAtomLedger}), Detail: FormatProjectors(a.Projectors)},
			{Name: "audit D_F symbolic edge skeleton as coupling graph only", Passed: a.DF.SymbolicEdgeSupportAudited && a.DF.CouplingGraphOnly && !a.DF.NumericalMagnitudes && !a.DF.UsesObservedMasses && !a.DF.UsesCKM && !a.DF.UsesPMNS && containsAll(a.DF.Supports, []string{SupportDFEdgesAsSocketGraph}) && containsAll(a.DF.Failures, []string{FailureDFEdgesNotMagnitudes, FailureNoNumericalYukawaSockets, FailureNoObservedYukawaFit}), Detail: FormatDF(a.DF)},
			{Name: "preserve direction: sector body before aggregate trace compression", Passed: a.Compression.SectorBodyBeforeCompression && !a.Compression.AggregateCompressionMapCertified && !a.Compression.AggregateToSectorPullbackCertified && !a.Compression.AggregateOperatorIsSectorLedger && !a.Compression.TraceMagnitudeReadoutCertified && containsAll(a.Compression.Supports, []string{SupportSectorBodyBeforeCompression}) && containsAll(a.Compression.Failures, []string{FailureNoAggregateCompressionMap, FailureAggregateOperatorNotSectorLedger, FailureNoTraceMagnitudeReadout}), Detail: FormatCompression(a.Compression)},
			{Name: "classify result as sealed coarse finite-sector body, not R3/R4 or official update", Passed: a.Impact.RhoActionSealConstructed && a.Impact.CoarseLedgerConstructed && a.Impact.CarrierProblemSolvedAtBlockLevel && !a.Impact.FullNativeFiniteTriple && !a.Impact.FullPiSectorF && !a.Impact.CanPromoteToR3 && !a.Impact.CanPromoteToR4 && !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && strings.Contains(a.Impact.Classification, "sealed coarse") && containsAll(a.Impact.Failures, []string{FailureRepresentationSealNotNative, FailureNoTraceMagnitudeReadout, FailureNoAggregateCompressionMap, FailureNotR3, FailureNotR4, FailureNoNEffUpdate, FailureNoCYukawaUpdate}), Detail: FormatImpact(a.Impact)},
			{Name: "preserve Gate 838 firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.RepresentationSealNotNative && a.Firewalls.NoFullFiniteTripleProof && a.Firewalls.NoExplicitMatrices && a.Firewalls.NoFirstOrderProof && a.Firewalls.NoBimoduleProof && a.Firewalls.NoCanonicalColorAtoms && a.Firewalls.MatrixUnitsBasisDependent && a.Firewalls.NoFineColorLedger && a.Firewalls.DFEdgesNotMagnitudes && a.Firewalls.NoNumericalYukawaSockets && a.Firewalls.NoTraceMagnitudeReadout && a.Firewalls.NoCompressionMap && a.Firewalls.AggregateNotSectorLedger && a.Firewalls.AlphaSealed && a.Firewalls.NoNEffUpdate && a.Firewalls.NoCYukawaUpdate && a.Firewalls.NoObservedYukawaFit && a.Firewalls.NoParticleAssignment && a.Firewalls.NoThreeGeneration && a.Firewalls.NotR3 && a.Firewalls.NotR4 && a.Firewalls.Verdict == StatusFirewallGate838, Detail: a.Firewalls.Verdict},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatW(a.W), FormatE(a.E), FormatRho(a.Rho), FormatProjectors(a.Projectors), FormatDF(a.DF), FormatCompression(a.Compression), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
