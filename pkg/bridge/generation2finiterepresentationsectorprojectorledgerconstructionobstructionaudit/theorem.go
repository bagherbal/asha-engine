package generation2finiterepresentationsectorprojectorledgerconstructionobstructionaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "GATE-835-FINITE-REPRESENTATION-SECTOR-PROJECTOR-LEDGER-CONSTRUCTION-OBSTRUCTION"
	theoremName = "Gate 835 — Finite Representation SectorProjectorLedger Construction/Obstruction Audit"
)

func Generation2FiniteRepresentationSectorProjectorLedgerConstructionObstructionAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate 835 finite representation sector-projector ledger audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 834 representation-layer requirement and R2++ status", Passed: a.Ledger.R2PlusPlusConsolidated && a.Ledger.AlphaSealed && !a.Ledger.R3SectorLedgerCertified && !a.Ledger.PiSectorFCertified && !a.Ledger.SigmaCertified && math.Abs(a.Ledger.OperatorNEff-OperatorNEff) < 5e-16 && strings.Contains(a.Truth, "Pi_sector^F"), Detail: FormatLedger(a.Ledger)},
			{Name: "audit represented central support projector recipe without instantiating full ledger", Passed: a.Central.SummandCount == AFSummandCount && a.Central.CentralIdempotentsOrthogonal && a.Central.CentralIdempotentsSumI && a.Central.SupportProjectorRecipeDefined && !a.Central.SupportProjectorsInstantiated && !a.Central.SupportRanksCertified && a.Central.BasisIndependentAtCoarseLevel && !a.Central.CompleteFiniteSectorLedger && !a.Central.TraceMagnitudeCertified && containsAll(a.Central.Supports, []string{SupportRepresentedFiniteTripleIsCorrectLayer, SupportCentralSupportsWouldSourceCoarseBlocks, SupportSectorLedgerRequiresHFPackage}) && containsAll(a.Central.Failures, []string{FailureNoCompleteFiniteRepresentationLedger, FailureNoCompleteRhoFSupportRankLedger, FailureCentralSupportsOnlyCoarseNotFullLedger}), Detail: FormatCentral(a.Central)},
			{Name: "audit chirality and real-structure refinements as required but uncertified", Passed: a.Chirality.GammaFAvailable && a.Chirality.JFAvailable && !a.Chirality.GammaRefinementCertified && !a.Chirality.JRefinementCertified && !a.Chirality.LeftRightSplitCertified && !a.Chirality.ParticleOppositeSplitCertified && !a.Chirality.CompatibleWithCentralSupports && !a.Chirality.CompleteRefinementLedger && !a.Chirality.TraceMagnitudeCertified && containsAll(a.Chirality.Failures, []string{FailureNoGammaFProjectorRefinementCertified, FailureNoJFRealStructureRefinementCertified, FailureNoParticleAntiparticleSplitCertified}), Detail: FormatChirality(a.Chirality)},
			{Name: "audit bimodule, commutant, and first-order typing requirements", Passed: a.Bimodule.LeftActionRequired && a.Bimodule.RightActionRequired && !a.Bimodule.CommutantStableProjectorsCertified && !a.Bimodule.FirstOrderStableCertified && !a.Bimodule.BimoduleDecompositionCertified && !a.Bimodule.CompleteTypedLedger && containsAll(a.Bimodule.Supports, []string{SupportBimoduleCommutantIsNecessaryTypingCondition}) && containsAll(a.Bimodule.Failures, []string{FailureNoBimoduleCommutantLedgerCertified, FailureNoFirstOrderStableSectorProjectorsCertified}), Detail: FormatBimodule(a.Bimodule)},
			{Name: "audit finite Dirac edge support without converting edges into magnitudes", Passed: a.DiracEdges.RequiresProjectorLedger && a.DiracEdges.UsesDF && a.DiracEdges.EdgeSupportAudited && !a.DiracEdges.EdgeSupportLedgerCertified && a.DiracEdges.CouplingGraphOnly && !a.DiracEdges.TraceMagnitudeReadoutCertified && !a.DiracEdges.YukawaValuesCertified && !a.DiracEdges.ObservedMassDataUsed && containsAll(a.DiracEdges.Failures, []string{FailureNoDFEdgeSupportLedgerCertified, FailureDFEdgeSupportNotMagnitudeReadout, FailureNoSectorTraceMagnitudeReadout}), Detail: FormatDiracEdges(a.DiracEdges)},
			{Name: "reinforce M3 matrix-unit color-frame firewall", Passed: a.Matrix.MatrixUnitsExist && a.Matrix.DiagonalProjectorsExist && a.Matrix.MatrixUnitCount == M3MatrixUnitCount && a.Matrix.DiagonalProjectorCount == M3ColorAtomCount && !a.Matrix.CanonicalColorFrameCertified && !a.Matrix.BasisIndependentAtoms && !a.Matrix.CanonicalColorAtomsCertified && !a.Matrix.CompleteSectorLedger && containsAll(a.Matrix.Failures, []string{FailureM3MatrixUnitsNotCanonicalColorAtomsWithoutFrame, FailureNoCanonicalColorFrame}), Detail: FormatMatrix(a.Matrix)},
			{Name: "defer aggregate-carrier pullback until Pi_sector^F codomain exists", Passed: !a.Pullback.PiSectorFCodomainCertified && !a.Pullback.PullbackAllowedToRun && !a.Pullback.SigmaCertified && !a.Pullback.TopI3PulledBack && !a.Pullback.FockP1P3PulledBack && a.Pullback.NonCircular && containsAll(a.Pullback.Failures, []string{FailureNoPiSectorFCodomainCertified, FailureNoAggregateCarrierPullbackYet, FailureNoSigmaMap}), Detail: FormatPullback(a.Pullback)},
			{Name: "block trace magnitudes, R3/R4 promotion, and official ledger updates", Passed: a.Impact.CentralSupportRecipe && !a.Impact.ChiralityRealRefinement && !a.Impact.BimoduleTyping && a.Impact.DFEdgeAudit && !a.Impact.PiSectorFCertified && !a.Impact.SigmaCertified && !a.Impact.TraceMagnitudeReadoutCertified && !a.Impact.CanPromoteToR3 && !a.Impact.CanPromoteToR4 && !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && strings.Contains(a.Impact.NextMissingObject, "Pi_sector^F") && containsAll(a.Impact.Failures, []string{FailureNoCompleteFiniteRepresentationLedger, FailureNoPiSectorFCodomainCertified, FailureNoSectorTraceMagnitudeReadout, FailureAggregateOperatorNotR3, FailureNoCYukawaUpdate}), Detail: FormatImpact(a.Impact)},
			{Name: "preserve Gate 835 physical firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.AFAloneNotLedger && a.Firewalls.NoCompletePiSectorF && a.Firewalls.CentralOnlyCoarse && a.Firewalls.NoGammaRefinement && a.Firewalls.NoJRefinement && a.Firewalls.NoBimoduleLedger && a.Firewalls.NoFirstOrder && a.Firewalls.NoDFEdgeLedger && a.Firewalls.DFEdgesNotMagnitudes && a.Firewalls.MatrixUnitsBasisDependent && a.Firewalls.NoColorFrame && a.Firewalls.NoSigmaMap && a.Firewalls.PullbackPremature && a.Firewalls.ProjectorsNotMagnitudes && a.Firewalls.NoMagnitudeReadout && a.Firewalls.AlphaSealed && a.Firewalls.NoBoundaryAlphaMap && a.Firewalls.NotR3 && a.Firewalls.NotR4 && a.Firewalls.NoNEffUpdate && a.Firewalls.NoCYukawaUpdate && a.Firewalls.NoObservedYukawaFit && a.Firewalls.NoPMNSCKM && a.Firewalls.NoParticleAssignment && a.Firewalls.Verdict == StatusFirewallGate835, Detail: a.Firewalls.Verdict},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatCentral(a.Central), FormatChirality(a.Chirality), FormatBimodule(a.Bimodule), FormatDiracEdges(a.DiracEdges), FormatMatrix(a.Matrix), FormatPullback(a.Pullback), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
