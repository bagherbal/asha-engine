package generation2finitetriplerepresentationcompletionandprojectorledgerdataaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "generation2-gate836-finite-triple-representation-completion-and-projector-ledger-data-audit"
	theoremName = "Gate 836 — Finite Triple Representation Completion and Projector-Ledger Data Audit"
)

func Generation2FiniteTripleRepresentationCompletionAndProjectorLedgerDataAuditTheorem() theorem.Theorem {
	return theorem.Theorem{
		ID:     theoremID,
		Name:   theoremName,
		Layer:  theorem.LayerBridge,
		Status: theorem.BridgeRequired,
		Verify: func() theorem.Result {
			a, err := BuildDefault()
			if err != nil {
				return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
			}
			checks := []theorem.Check{
				{Name: "inherit Gate 835 codomain obstruction and R2++ status", Passed: a.Ledger.R2PlusPlusConsolidated && !a.Ledger.R3SectorLedgerCertified && a.Ledger.AlphaSealed && !a.Ledger.PiSectorFCertified && !a.Ledger.SigmaCertified && !a.Ledger.TraceMagnitudeCertified && a.Ledger.OfficialLedgerFrozen, Detail: FormatLedger(a.Ledger)},
				{Name: "audit minimal finite triple representation data without inventing it", Passed: a.Data.AlgebraKnown && !a.Data.ExplicitHF && !a.Data.ExplicitRhoF && !a.Data.ExplicitJF && !a.Data.ExplicitGammaF && !a.Data.ExplicitDF && !a.Data.CompletePackageCertified && !a.Data.CanConstructPiSectorF && !a.Data.ObservedDataUsed && containsAll(a.Data.Supports, []string{SupportAFKnownButRepresentationDataIncomplete, SupportCompletePackageRequiredForPiSectorF, SupportFiniteRepresentationDataSealRequired}) && containsAll(a.Data.Failures, []string{FailureNoCompleteFiniteTripleRepresentationData, FailureNoExplicitHFCarrier, FailureNoExplicitRhoFRepresentation, FailureNoExplicitJFRealStructure, FailureNoExplicitGammaFChirality, FailureNoExplicitDFOperator}), Detail: FormatData(a.Data)},
				{Name: "audit represented central support ranks as unavailable", Passed: a.Central.CentralIdempotentsOrthogonal && a.Central.CentralIdempotentsSumI && !a.Central.RhoFExplicit && !a.Central.SupportProjectorsInstantiated && !a.Central.SupportRanksCertified && !a.Central.OrthogonalityCertified && !a.Central.CompletenessCertified && !a.Central.RankLedgerCertified && a.Central.CoarseRecipeOnly && !a.Central.CompleteLedger && containsAll(a.Central.Failures, []string{FailureNoCentralSupportRankLedger, FailureNoRepresentedSupportProjectors, FailureNoSupportOrthogonalityCompleteness}), Detail: FormatCentral(a.Central)},
				{Name: "audit chirality and real-structure refinement data as missing", Passed: !a.Chirality.GammaFExplicit && !a.Chirality.JFExplicit && !a.Chirality.ChiralityProjectorsInstantiated && !a.Chirality.ChiralityRanksCertified && !a.Chirality.RealStructureImagesInstantiated && !a.Chirality.ParticleOppositeSplitCertified && !a.Chirality.LeftRightSplitCertified && !a.Chirality.CompatibleWithCentralSupports && !a.Chirality.YukawaMagnitudeCertified && !a.Chirality.ObservedParticleAssignment && containsAll(a.Chirality.Failures, []string{FailureNoChiralityProjectorLedger, FailureChiralitySplitNotYukawaMagnitude, FailureNoRealStructureImageLedger, FailureJRefinementNotParticleAssignment}), Detail: FormatChirality(a.Chirality)},
				{Name: "audit bimodule and first-order stability data as missing", Passed: a.Bimodule.RequiresLeftAction && a.Bimodule.RequiresRightAction && !a.Bimodule.RhoFExplicit && !a.Bimodule.JFExplicit && !a.Bimodule.DFExplicit && !a.Bimodule.LeftActionMatricesCertified && !a.Bimodule.RightActionMatricesCertified && !a.Bimodule.BimoduleStabilityCertified && !a.Bimodule.CommutantDecompositionCertified && !a.Bimodule.FirstOrderCompatibilityCertified && !a.Bimodule.TypedProjectorLedgerCertified && containsAll(a.Bimodule.Failures, []string{FailureNoBimoduleStabilityData, FailureNoFirstOrderCompatibilityCertificate}), Detail: FormatBimodule(a.Bimodule)},
				{Name: "audit finite Dirac edge graph as unavailable and non-magnitude", Passed: a.Edges.RequiresPiSectorF && a.Edges.RequiresDF && !a.Edges.PiSectorFExists && !a.Edges.DFExplicit && !a.Edges.EdgeBlocksComputed && !a.Edges.EdgeSupportGraphCertified && a.Edges.CouplingGraphOnly && !a.Edges.TraceMagnitudeReadoutCertified && !a.Edges.YukawaValuesCertified && !a.Edges.ObservedDataUsed && containsAll(a.Edges.Failures, []string{FailureNoDFEdgeGraph, FailureDFEdgesNotTraceMagnitudeReadout, FailureNoSectorTraceMagnitudeReadout}), Detail: FormatEdges(a.Edges)},
				{Name: "preserve M3 matrix-unit color-frame data firewall", Passed: a.ColorFrame.M3MatrixUnitsExist && a.ColorFrame.DiagonalProjectorsExist && a.ColorFrame.MatrixUnitCount == M3MatrixUnitCount && a.ColorFrame.DiagonalProjectorCount == M3ColorAtomCount && !a.ColorFrame.CanonicalFrameCertified && !a.ColorFrame.BasisIndependentAtoms && !a.ColorFrame.ColorAtomLedgerCertified && !a.ColorFrame.GaugeFrameChoiceUsed && containsAll(a.ColorFrame.Failures, []string{FailureNoCanonicalM3ColorFrame, FailureM3MatrixUnitsBasisDependent}), Detail: FormatColorFrame(a.ColorFrame)},
				{Name: "block Pi_sector^F construction, Sigma, trace magnitudes, and ledger updates", Passed: !a.Impact.FiniteTripleDataComplete && !a.Impact.CentralRanksComplete && !a.Impact.ChiralityRealComplete && !a.Impact.BimoduleFirstOrderComplete && !a.Impact.DiracEdgeGraphComplete && !a.Impact.ColorFrameComplete && !a.Impact.PiSectorFConstructible && !a.Impact.PiSectorFCertified && !a.Impact.SigmaAllowed && !a.Impact.TraceMagnitudeAllowed && !a.Impact.CanPromoteToR3 && !a.Impact.CanPromoteToR4 && !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && strings.Contains(a.Impact.NextRequiredObject, "FiniteRepresentationDataSeal") && containsAll(a.Impact.Failures, []string{FailureNoCompleteFiniteTripleRepresentationData, FailureNoPiSectorFConstruction, FailureNoPiSectorFCodomain, FailureNoAggregateCarrierPullbackYet, FailureNoSigmaMap, FailureSectorProjectorsNotTraceMagnitudes, FailureAggregateOperatorNotR3, FailureNoCYukawaUpdate}), Detail: FormatImpact(a.Impact)},
				{Name: "preserve Gate 836 physical and ledger firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.NoCompleteFiniteTripleData && a.Firewalls.NoExplicitHF && a.Firewalls.NoExplicitRhoF && a.Firewalls.NoExplicitJF && a.Firewalls.NoExplicitGammaF && a.Firewalls.NoExplicitDF && a.Firewalls.NoCentralRanks && a.Firewalls.NoSupportProjectors && a.Firewalls.NoChiralityLedger && a.Firewalls.NoRealStructureLedger && a.Firewalls.NoBimoduleStability && a.Firewalls.NoFirstOrder && a.Firewalls.NoDFEdgeGraph && a.Firewalls.DFEdgesNotMagnitudes && a.Firewalls.NoColorFrame && a.Firewalls.MatrixUnitsBasisDependent && a.Firewalls.NoPiSectorF && a.Firewalls.PullbackPremature && a.Firewalls.NoSigmaMap && a.Firewalls.ProjectorsNotMagnitudes && a.Firewalls.NoMagnitudeReadout && a.Firewalls.AlphaSealed && a.Firewalls.NoBoundaryAlphaMap && a.Firewalls.NotR3 && a.Firewalls.NotR4 && a.Firewalls.NoNEffUpdate && a.Firewalls.NoCYukawaUpdate && a.Firewalls.NoObservedYukawaFit && a.Firewalls.NoPMNSCKM && a.Firewalls.NoParticleAssignment && a.Firewalls.Verdict == StatusFirewallGate836, Detail: a.Firewalls.Verdict},
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
			notes := []string{a.Truth, FormatLedger(a.Ledger), FormatData(a.Data), FormatCentral(a.Central), FormatChirality(a.Chirality), FormatBimodule(a.Bimodule), FormatEdges(a.Edges), FormatColorFrame(a.ColorFrame), FormatImpact(a.Impact), a.Final}
			notes = append(notes, Statuses()...)
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
		},
	}
}
