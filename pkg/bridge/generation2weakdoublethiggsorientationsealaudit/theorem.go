package generation2weakdoublethiggsorientationsealaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_WEAKDOUBLET_HIGGS_ORIENTATION_SEAL_AUDIT"
	theoremName = "Gate 853 — WeakDoublet / HiggsOrientationSeal Audit"
)

func Generation2WeakDoubletHiggsOrientationSealAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 852 firewall and audit quaternionic weak module", Passed: a.Impact.Gate852Inherited && a.QuaternionicFirewall.FullModuleHStable && !a.QuaternionicFirewall.NativeRankOneEigensplit && !a.QuaternionicFirewall.GenericHActionPreservesLines && containsAll(a.QuaternionicFirewall.Failures, []string{FailureWeakSplitNotNativeHEigensplit, FailureRankOneLinesNotGloballyHStable}), Detail: FormatQuaternionicFirewall(a.QuaternionicFirewall)},
			{Name: "define Higgs/weak orientation seal without native derivation", Passed: a.OrientationSeal.DefinedAtSealLevel && a.OrientationSeal.ProjectorsComplete && a.OrientationSeal.Orthogonal && a.OrientationSeal.HPlusRank == 1 && a.OrientationSeal.HMinusRank == 1 && !a.OrientationSeal.NativeDerivation && a.OrientationSeal.RequiresGaugeOrientation && containsAll(a.OrientationSeal.Supports, []string{SupportWeakSocketSplitAfterHiggsSeal}) && containsAll(a.OrientationSeal.Failures, []string{FailureHiggsOrientationSealNotNative, FailureNoNativeHiggsVacuumTheorem}), Detail: FormatOrientationSeal(a.OrientationSeal)},
			{Name: "separate full weak-module stability from oriented socket stability", Passed: a.Stability.FullModuleStableUnderH && !a.Stability.SocketsStableUnderFullH && a.Stability.SocketsStableAfterOrientation && !a.Stability.GlobalQuaternionicEigenspaces && containsAll(a.Stability.Failures, []string{FailureWeakSplitNotNativeHEigensplit, FailureRankOneLinesNotGloballyHStable}), Detail: FormatStability(a.Stability)},
			{Name: "rewrite Gate 847 symbolic edges in the oriented frame", Passed: a.EdgeRewrite.ActiveEdgeCount == ActiveEdgeCount && a.EdgeRewrite.PunctureEdgeAbsent && a.EdgeRewrite.LeptoColorPreserved && a.EdgeRewrite.UsesOrientationFrame && a.EdgeRewrite.CompatibleWithGate847Skeleton && containsAll(a.EdgeRewrite.Supports, []string{SupportEdgeSkeletonCompatibleWithSeal}), Detail: FormatEdgeRewrite(a.EdgeRewrite)},
			{Name: "classify left kernel as orientation-relative only", Passed: a.Kernel.Kernel == "h_+ tensor P_1" && a.Kernel.Rank == 1 && a.Kernel.OrientationRelative && a.Kernel.StableUnderOrientationBlocks && !a.Kernel.StableUnderFullRhoJ && !a.Kernel.PhysicalNeutrino && !a.Kernel.Masslessness && containsAll(a.Kernel.Failures, []string{FailureLeftKernelNotRepresentationStable, FailureNoPhysicalNeutrinoTheorem, FailureNoMasslessnessTheorem}), Detail: FormatKernel(a.Kernel)},
			{Name: "prepare but do not execute first-order calculation", Passed: a.FirstOrderPreparation.OrientationSealAvailable && !a.FirstOrderPreparation.OperatorRealizationReady && !a.FirstOrderPreparation.FirstOrderExecutable && !a.FirstOrderPreparation.FirstOrderCertified && containsAll(a.FirstOrderPreparation.Failures, []string{FailureNoOperatorLevelRhoFJFDFF, FailureNoFirstOrderProofYet, FailureNoJOppositeProofYet}), Detail: FormatFirstOrderPreparation(a.FirstOrderPreparation)},
			{Name: "preserve ledgers and R3/R4 firewalls", Passed: a.Ledger.OfficialFrozen && !a.Ledger.AlphaNative && !a.Ledger.R3 && !a.Ledger.R4 && a.Impact.AlphaStillSealed && a.Impact.MagnitudesStillMissing && !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && !a.Impact.CanPromoteToR3 && !a.Impact.CanPromoteToR4 && !a.Impact.FirstOrderProved && !a.Impact.OperatorRealizationReady, Detail: FormatLedger(a.Ledger) + " | " + FormatImpact(a.Impact)},
			{Name: "preserve Gate 853 orientation firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.WeakSplitNotNative && a.Firewalls.HiggsOrientationNotNative && a.Firewalls.RankOneLinesNotHStable && a.Firewalls.NoNativeHiggsVacuum && a.Firewalls.NoOperatorPackage && a.Firewalls.NoFirstOrderProof && a.Firewalls.NoJOppositeProof && a.Firewalls.NoBimoduleProof && a.Firewalls.KernelNotRepresentationStable && a.Firewalls.OrientationDoesNotDeriveAlpha && a.Firewalls.NoYukawaMagnitudes && a.Firewalls.NoTraceMagnitudeReadout && a.Firewalls.NoOfficialNEffUpdate && a.Firewalls.NoCYukawaCHiggsUpdate && a.Firewalls.NotR3 && a.Firewalls.NotR4 && a.Firewalls.NoPhysicalNeutrino && a.Firewalls.NoRightNeutrino && a.Firewalls.NoMasslessness && a.Firewalls.NoWeakMixingOrHiggsMass && a.Firewalls.NoParticleAssign && a.Firewalls.Verdict == StatusFirewallVerdict, Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatQuaternionicFirewall(a.QuaternionicFirewall), FormatOrientationSeal(a.OrientationSeal), FormatStability(a.Stability), FormatEdgeRewrite(a.EdgeRewrite), FormatKernel(a.Kernel), FormatFirstOrderPreparation(a.FirstOrderPreparation), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
