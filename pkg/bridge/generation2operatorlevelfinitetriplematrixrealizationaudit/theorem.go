package generation2operatorlevelfinitetriplematrixrealizationaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_OPERATOR_LEVEL_FINITE_TRIPLE_MATRIX_REALIZATION_AUDIT"
	theoremName = "Gate 854 — Operator-Level FiniteTriple Matrix Realization Audit"
)

func Generation2OperatorLevelFiniteTripleMatrixRealizationAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "define ordered minimal H_F basis and separate ambient/active carriers", Passed: a.Basis.Complete && a.Basis.JCopyIncluded && a.Basis.HLRank == HLRank && a.Basis.HRMinRank == HRMinRank && a.Basis.HPartMinRank == HPartMinRank && a.Basis.HFMinRank == HFMinRank && a.Basis.AmbientPartRank == AmbientPartRank && a.Basis.AmbientF == AmbientFRank && a.Basis.RightPunctureOutside && a.Basis.LeftKernelInHL && containsAll(a.Basis.Supports, []string{StatusOrderedBasisDefined, StatusAmbientActiveSeparated}), Detail: FormatBasis(a.Basis)},
			{Name: "instantiate rho_F block-action matrix seal on minimal carrier", Passed: a.Rho.DefinedAtSealLevel && a.Rho.PreservesMinimalCarrier && !a.Rho.PunctureForcedBackIntoCarrier && a.Rho.M3ActsOnP3 && a.Rho.M3TrivialOnP1 && a.Rho.CActsOnRightCharacters && a.Rho.HActsOnFullWeakDoublet && a.Rho.HMayMixHPlusHMinus && !a.Rho.HPlusHMinusNativeHEigensplit && !a.Rho.NativeRepresentationProof && containsAll(a.Rho.Supports, []string{SupportRhoPreservesMinimalCarrier}) && containsAll(a.Rho.Failures, []string{FailureWeakFrameNotNativeHInvariant, FailureDFAfterOrientationNotUnbrokenH}), Detail: FormatRho(a.Rho)},
			{Name: "instantiate gamma_F matrix seal and preserve KO-sign firewall", Passed: a.Gamma.ParticleSideDefined && a.Gamma.SquareIdentity && a.Gamma.ChiralityOddWithDFByBlock && !a.Gamma.KOExtensionCertified && a.Gamma.SupportLevelOnly && containsAll(a.Gamma.Failures, []string{FailureKOSignExtensionNotCertified}), Detail: FormatGamma(a.Gamma)},
			{Name: "instantiate formal J_F particle/opposite exchange seal", Passed: a.J.ParticleOppositeExchange && a.J.AntiunitaryFormal && a.J.OppositeCopyDimension == HPartMinRank && !a.J.KOSignsCertified && !a.J.OppositeActionCompatibilityProved && !a.J.FullRealStructureProof && containsAll(a.J.Failures, []string{FailureJStructureFormalOnly, FailureNoJOppositeProofYet}), Detail: FormatJ(a.J)},
			{Name: "instantiate symbolic D_F matrix seal", Passed: a.D.YPlus1Zero && len(a.D.YTerms) == YActiveEdges && a.D.HLRank == HLRank && a.D.HRMinRank == HRMinRank && a.D.ParticleDim == HPartMinRank && a.D.Rank == DSymRank && a.D.KernelRank == DSymKernelRank && a.D.SelfAdjointByBlock && a.D.ChiralityOddByBlock && a.D.LeptoColorPreserving && a.D.ExtendedToJCopy && !a.D.OperatorValuedMatrixCertified && !a.D.NumericalYukawaMagnitudesCertified && a.D.OrientationFrameObject && !a.D.UnbrokenHEquivariantTheorem && containsAll(a.D.Failures, []string{FailureDOperatorValuedNotCertified, FailureYCoefficientsSymbolicOnly}), Detail: FormatD(a.D)},
			{Name: "verify matrix-level consistency checks but defer first-order proof", Passed: a.Checks.BasisComplete && a.Checks.DimensionConsistent && a.Checks.RhoPreservesCarrier && a.Checks.GammaSquaredIdentity && a.Checks.DSelfAdjoint && a.Checks.DChiralityOdd && a.Checks.JMapsParticleToOpposite && a.Checks.PunctureOutside && a.Checks.LeftKernelPresent && a.Checks.FirstOrderExecutableNextGate && !a.Checks.FirstOrderProvedThisGate && containsAll(a.Checks.Failures, []string{FailureNoFirstOrderProofYet, FailureNoFirstOrderCalculationYet}), Detail: FormatChecks(a.Checks)},
			{Name: "preserve official ledger and R3/R4 firewalls", Passed: a.Ledger.OfficialFrozen && !a.Ledger.AlphaNative && !a.Ledger.R3 && !a.Ledger.R4 && a.Impact.OperatorMatrixSeal && a.Impact.FirstOrderPrepared && !a.Impact.FirstOrderProved && !a.Impact.NativeFiniteTripleProof && a.Impact.AlphaStillSealed && a.Impact.MagnitudesStillMissing && !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && !a.Impact.CanPromoteToR3 && !a.Impact.CanPromoteToR4, Detail: FormatLedger(a.Ledger) + " | " + FormatImpact(a.Impact)},
			{Name: "preserve Gate 854 matrix-seal firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.MatrixSealNotNative && a.Firewalls.NoFirstOrderProof && a.Firewalls.NoFirstOrderCalculation && a.Firewalls.NoJOppositeProof && a.Firewalls.NoBimoduleProof && a.Firewalls.KOSignNotCertified && a.Firewalls.JFormalOnly && a.Firewalls.DSymbolicOnly && a.Firewalls.YSymbolicOnly && a.Firewalls.WeakFrameNotNativeHInvariant && a.Firewalls.DFOrientationNotUnbrokenH && a.Firewalls.KernelNotStable && a.Firewalls.PunctureNotNative && a.Firewalls.NoAlphaSource && a.Firewalls.NoYukawaMagnitudes && a.Firewalls.NoTraceReadout && a.Firewalls.NoOfficialNEffUpdate && a.Firewalls.NoCYukawaCHiggsUpdate && a.Firewalls.NotR3 && a.Firewalls.NotR4 && a.Firewalls.NoParticleAssign && a.Firewalls.NoNeutrinoTheorem && a.Firewalls.NoThreeGenerationTheorem && a.Firewalls.Verdict == StatusFirewallVerdict, Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatBasis(a.Basis), FormatRho(a.Rho), FormatGamma(a.Gamma), FormatJ(a.J), FormatD(a.D), FormatChecks(a.Checks), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
