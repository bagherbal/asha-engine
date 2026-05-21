package generation2leptocolorfockcarrierrepresentationsealaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "GATE-837-LEPTOCOLOR-FOCK-CARRIER-REPRESENTATION-SEAL"
	theoremName = "Gate 837 — LeptoColor Fock Carrier Representation Seal Audit"
)

func Generation2LeptoColorFockCarrierRepresentationSealAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate 837 lepto-color carrier seal audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 836 finite triple data-completion obstruction", Passed: a.Ledger.R2PlusPlusConsolidated && a.Ledger.AlphaSealed && !a.Ledger.R3SectorLedgerCertified && a.Ledger.AggregateAtomCount == AggregateAtomCount && math.Abs(a.Ledger.OperatorNEff-OperatorNEff) < 5e-16 && containsAll(a.Impact.Verdicts, []string{StatusGate836Inherited, StatusR2PlusPlusRetained}), Detail: FormatLedger(a.Ledger)},
			{Name: "construct shared lepto-color Fock carrier W=C plus C^3", Passed: a.Carrier.Dim == LeptoColorDim && a.Carrier.LeptonBlockDim == LeptonBlockDim && a.Carrier.ColorBlockDim == ColorBlockDim && a.Carrier.P1Rank == LeptonBlockDim && a.Carrier.P3Rank == ColorBlockDim && a.Carrier.P1P3Orthogonal && a.Carrier.P1PlusP3CompletesW && a.Carrier.FockOnePlusThreeCarrier && a.Carrier.LeptoColorCarrier && a.Carrier.BMinusLTraceZero && containsAll(a.Carrier.Supports, []string{SupportSharedWUnifiesFockOnePlusThreeAndM3ColorModule, SupportP1P3SourceBMinusLInternally}), Detail: FormatCarrier(a.Carrier)},
			{Name: "certify block-level M3(C) action on P3W by representation seal", Passed: a.M3.ActsOnP3WBySealDefinition && a.M3.P3WIsM3Fundamental && a.M3.M3IdentityTrace == ColorBlockDim && a.M3.MatrixUnitsExist && a.M3.MatrixUnitsActWithinP3W && a.M3.P1InvariantUnderM3 && a.M3.BlockLevelCanonical && !a.M3.IndividualColorAtomsCanonical && !a.M3.CanonicalColorFrameCertified && a.M3.NoSeparateTripletBridgeNeeded && !a.M3.ContradictsGate833 && containsAll(a.M3.Supports, []string{SupportP3WIsM3FundamentalByRepresentationSeal, SupportCarrierProblemSolvedAtBlockLevel}) && containsAll(a.M3.Failures, []string{FailureNoCanonicalColorAtomFrame, FailureM3MatrixUnitsBasisDependent}), Detail: FormatM3(a.M3)},
			{Name: "build one-generation-like finite carrier skeleton without complete representation matrices", Passed: a.Body.RightSlotDim == RightWeakSlotDim && a.Body.LeftSlotDim == LeftWeakSlotDim && a.Body.ChiralitySlotDim == ChiralitySlotDim && a.Body.WDim == LeptoColorDim && a.Body.HPartDim == HPartDim && a.Body.RealOppositeCopies == RealOppositeCopies && a.Body.HFDim == HFSealDim && a.Body.RhoFRoleDeclared && a.Body.GammaFRoleDeclared && a.Body.JFRoleDeclared && a.Body.DFRoleDeclared && !a.Body.CompleteRhoFActionLedger && !a.Body.ExplicitGammaFOperator && !a.Body.ExplicitJFOperator && !a.Body.SymbolicDFEdgeMatrix && !a.Body.ObservedDataUsed && containsAll(a.Body.Failures, []string{FailureNoCompleteRhoFActionLedger, FailureNoCompleteFiniteTripleData, FailureNoGammaFOperatorMatrices, FailureNoJFOperatorMatrices, FailureNoDFSymbolicEdgeMatrix, FailureDFSocketsNotYukawaMagnitudes}), Detail: FormatBody(a.Body)},
			{Name: "audit projector prospects without certifying Pi_sector^F or magnitudes", Passed: a.Projectors.CanConstructBlockProjectors && a.Projectors.CanConstructChiralityTimesLeptoColorSupports && a.Projectors.CanConstructOppositeCopySupports && !a.Projectors.PiSectorFCertified && !a.Projectors.SupportRankLedgerCertified && !a.Projectors.BimoduleFirstOrderCertified && !a.Projectors.DFEdgeSupportLedgerCertified && !a.Projectors.TraceMagnitudeReadoutCertified && !a.Projectors.SectorProjectorsAreMagnitudes && strings.Contains(a.Projectors.NextRequiredObject, "Pi_sector^F") && containsAll(a.Projectors.Failures, []string{FailureNoPiSectorFLedgerYet, FailureNoTraceMagnitudeReadout, FailureCarrierSealNotMagnitudeReadout}), Detail: FormatProjectors(a.Projectors)},
			{Name: "correct direction to finite body -> aggregate trace-compression shadow", Passed: a.Compression.DirectionCorrected && a.Compression.FiniteBodyToAggregateShadow && !a.Compression.AggregateToFiniteBody && !a.Compression.R2OperatorSectorLedger && !a.Compression.CompressionMapCertified && !a.Compression.SigmaPullbackCertified && a.Compression.FiniteSectorBodyRequiredFirst && len(a.Compression.AggregateClasses) == 3 && containsAll(a.Compression.Supports, []string{SupportTraceCompressionShadowDirection}) && containsAll(a.Compression.Failures, []string{FailureNoAggregateCompressionMapYet, FailureAggregateOperatorNotSectorLedger, FailureNoSigmaPullback}), Detail: FormatCompression(a.Compression)},
			{Name: "classify impact as carrier seal only, not R3/R4 or official ledger update", Passed: a.Impact.CarrierSealConstructed && a.Impact.CarrierProblemSolvedAtBlockLevel && !a.Impact.CompleteFiniteTripleData && !a.Impact.PiSectorFCertified && !a.Impact.CompressionMapCertified && !a.Impact.TraceMagnitudeReadoutCertified && !a.Impact.CanPromoteToR3 && !a.Impact.CanPromoteToR4 && !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && strings.Contains(a.Impact.NextGate, "Pi_sector^F") && containsAll(a.Impact.Failures, []string{FailureNoCompleteFiniteTripleData, FailureNoPiSectorFLedgerYet, FailureNoAggregateCompressionMapYet, FailureNoTraceMagnitudeReadout, FailureNotR3, FailureNotR4, FailureNoNEffUpdate, FailureNoCYukawaUpdate}), Detail: FormatImpact(a.Impact)},
			{Name: "preserve Gate 837 physical and ledger firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.CarrierSealNotNativeDerivation && a.Firewalls.NoCompleteRhoF && a.Firewalls.NoCompleteFiniteTriple && a.Firewalls.NoCanonicalColorAtoms && a.Firewalls.MatrixUnitsBasisDependent && a.Firewalls.NoGammaF && a.Firewalls.NoJF && a.Firewalls.NoDFEdgeMatrix && a.Firewalls.DFSocketsNotMagnitudes && a.Firewalls.NoPiSectorF && a.Firewalls.NoCompressionMap && a.Firewalls.AggregateNotSectorLedger && a.Firewalls.NoSigma && a.Firewalls.NoMagnitudeReadout && a.Firewalls.AlphaSealed && a.Firewalls.NoBoundaryAlphaMap && a.Firewalls.NotR3 && a.Firewalls.NotR4 && a.Firewalls.NoNEffUpdate && a.Firewalls.NoCYukawaUpdate && a.Firewalls.NoObservedYukawaFit && a.Firewalls.NoPMNSCKM && a.Firewalls.NoParticleAssignment && a.Firewalls.NoThreeGeneration && a.Firewalls.Verdict == StatusFirewallGate837, Detail: a.Firewalls.Verdict},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatCarrier(a.Carrier), FormatM3(a.M3), FormatBody(a.Body), FormatProjectors(a.Projectors), FormatCompression(a.Compression), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
