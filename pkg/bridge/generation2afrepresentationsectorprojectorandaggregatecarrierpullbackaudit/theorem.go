package generation2afrepresentationsectorprojectorandaggregatecarrierpullbackaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "GATE-834-AF-REPRESENTATION-SECTOR-PROJECTOR-AGGREGATE-CARRIER-PULLBACK"
	theoremName = "Gate 834 — A_F-Representation Sector Projector and Aggregate-Carrier Pullback Audit"
)

func Generation2AFRepresentationSectorProjectorAndAggregateCarrierPullbackAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate 834 A_F representation projector audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 833 direct triplet bridge obstruction", Passed: a.Ledger.R2PlusPlusConsolidated && a.Ledger.AlphaSealed && !a.Ledger.R3SectorLedgerCertified && !a.Ledger.SectorProjectorMapCertified && math.Abs(a.Ledger.OperatorNEff-OperatorNEff) < 5e-16 && strings.Contains(a.Truth, "represented finite algebra"), Detail: FormatLedger(a.Ledger)},
			{Name: "audit A_F central idempotents as coarse sector block candidates", Passed: a.Central.SummandCount == AFSummandCount && a.Central.Orthogonal && a.Central.SumToIdentity && a.Central.CoarseSectorBlocks && a.Central.RepresentationIndependent && !a.Central.SectorLedgerCertified && !a.Central.TraceMagnitudeCertified && containsAll(a.Central.Supports, []string{SupportAFStrongestSectorSource, SupportCentralIdempotentsSourceCoarseBlocks}) && containsAll(a.Central.Failures, []string{FailureAFAloneNotSectorLedgerWithoutRepresentation, FailureSectorProjectorsNotTraceMagnitudes}), Detail: FormatCentral(a.Central)},
			{Name: "audit representation-induced projector requirement without completing rho_F package", Passed: a.Representation.UsesHF && a.Representation.UsesRhoF && a.Representation.UsesJF && a.Representation.UsesGammaF && a.Representation.UsesDF && a.Representation.PartialPredataAvailable && !a.Representation.CompletePackageCertified && !a.Representation.RepresentationInducedProjectorsCertified && a.Representation.CanSourceCoarseProjectorCandidates && !a.Representation.CanSourceSectorLedger && !a.Representation.CanSourceTraceMagnitudes && containsAll(a.Representation.Supports, []string{SupportRepresentationCanInduceProjectors, SupportSectorProjectorsRequireRhoF}) && containsAll(a.Representation.Failures, []string{FailureNoCompleteRhoFRepresentationCertified, FailureNoCompleteFiniteHilbertPackage}), Detail: FormatRepresentation(a.Representation)},
			{Name: "enforce M3 matrix-unit basis-dependence firewall", Passed: a.MatrixUnits.MatrixUnitsExist && a.MatrixUnits.DiagonalProjectorsExist && a.MatrixUnits.MatrixUnitCount == M3MatrixUnitCount && !a.MatrixUnits.CanonicalColorFrameCertified && !a.MatrixUnits.BasisIndependent && !a.MatrixUnits.CanonicalColorAtomsCertified && a.MatrixUnits.SuppliesCarrierProjectors && !a.MatrixUnits.SuppliesSectorLedger && containsAll(a.MatrixUnits.Failures, []string{FailureM3MatrixUnitsNotCanonicalColorAtomsWithoutFrame, FailureNoCanonicalColorFrame}), Detail: FormatMatrixUnits(a.MatrixUnits)},
			{Name: "test aggregate-carrier pullback into represented sector projectors", Passed: a.Pullback.CentralBlocksAvailable && a.Pullback.RepresentationProjectorRecipe && a.Pullback.NonCircular && !a.Pullback.PullbackCertified && !a.Pullback.TopI3PulledBack && !a.Pullback.FockP1P3PulledBack && !a.Pullback.M3P3IntertwinerCertified && containsAll(a.Pullback.Failures, []string{FailureNoAggregateCarrierToRepresentationProjectorPullback, FailureNoSigmaMap, FailureTopI3NotPulledBackToRepresentationSector, FailureFockP1P3NotPulledBackToRepresentationSector, FailureNoM3P3Intertwiner}), Detail: FormatPullback(a.Pullback)},
			{Name: "block sector ledger, trace-magnitude readout, and ledger updates", Passed: a.Impact.CentralProjectorSource && a.Impact.RepresentationProjectorRecipe && !a.Impact.AggregatePullbackCertified && !a.Impact.SectorProjectorMapCertified && !a.Impact.SectorTraceLedgerCertified && !a.Impact.TraceMagnitudeReadoutCertified && !a.Impact.CanPromoteToR3 && !a.Impact.CanPromoteToR4 && !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && strings.Contains(a.Impact.NextMissingObject, "SectorProjectorMap") && containsAll(a.Impact.Failures, []string{FailureNoSigmaMap, FailureSectorProjectorsNotTraceMagnitudes, FailureNoSectorTraceMagnitudeReadout, FailureAggregateOperatorNotR3, FailureNoCYukawaUpdate}), Detail: FormatImpact(a.Impact)},
			{Name: "preserve Gate 834 physical firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.AFAloneNotLedger && a.Firewalls.RequiresRepresentation && a.Firewalls.NoCompleteRhoF && a.Firewalls.MatrixUnitsBasisDependent && a.Firewalls.NoColorFrame && a.Firewalls.NoAggregatePullback && a.Firewalls.NoSigmaMap && a.Firewalls.NoM3P3Intertwiner && a.Firewalls.ProjectorsNotMagnitudes && a.Firewalls.NoMagnitudeReadout && a.Firewalls.AlphaSealed && a.Firewalls.NoBoundaryAlphaMap && a.Firewalls.NotR3 && a.Firewalls.NotR4 && a.Firewalls.NoNEffUpdate && a.Firewalls.NoCYukawaUpdate && a.Firewalls.NoObservedYukawaFit && a.Firewalls.NoPMNSCKM && a.Firewalls.NoParticleAssignment && a.Firewalls.Verdict == StatusFirewallGate834, Detail: a.Firewalls.Verdict},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatCentral(a.Central), FormatRepresentation(a.Representation), FormatMatrixUnits(a.MatrixUnits), FormatPullback(a.Pullback), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
