package generation2nativez2equivariantairlockfunctorandboundaryalphasourceaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_NATIVE_Z2_EQUIVARIANT_AIRLOCK_FUNCTOR_AND_BOUNDARY_ALPHA_SOURCE_AUDIT"
	theoremName = "Gate 909 — Native Z2-Equivariant Airlock Functor and BoundaryAlpha Source Audit"
)

func Generation2NativeZ2EquivariantAirlockFunctorAndBoundaryAlphaSourceAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 908 sealed Z2 trace-ledger plateau inherited", Passed: a.Inherited.TraceLedgerZ2Class && a.Inherited.AlphaStillSealed && !a.Inherited.NativeR3 && containsAll(a.Inherited.Supports, []string{StatusGate908Inherited, SupportR3OnZ2AirlockClass}) && containsAll(a.Inherited.Failures, []string{FailureAlphaStillSealed, FailureNotNativeR3}), Detail: FormatInherited(a.Inherited)},
			{Name: "BoundaryAlpha candidate is Z2 well-defined at rank/class level", Passed: a.WellDefined.TauMapsPunctures && a.WellDefined.TauMapsDegreeOneTarget && a.WellDefined.TauMapsDegreeTwoTarget && a.WellDefined.IBZ2CommutesWithTau && a.WellDefined.RankPairInvariant && a.WellDefined.ClassLevelWellDefined && !a.WellDefined.NativeFunctorTheorem && containsAll(a.WellDefined.Supports, []string{SupportBoundaryAlphaFunctorZ2Class, SupportIBZ2CommutesTauRank, SupportAlphaRankPairInvariant}) && containsAll(a.WellDefined.Failures, []string{FailureZ2RankNotNativeFunctor, FailureNoNativeZ2BoundaryAlphaFunctor}), Detail: FormatWellDefined(a.WellDefined)},
			{Name: "reduced B2 response feeds Z2 quotient targets at shape level only", Passed: a.Response.ZeroOrderSuppressed && a.Response.CubicTermStopped && a.Response.Lambda3B2Zero && a.Response.CompatibleWithZ2Class && !a.Response.NativeBoundaryFunctional && !a.Response.NativeDegreeToFlagFunctor && a.Response.DegreeOneRank == RankF1OverF0 && a.Response.DegreeTwoRank == RankF2OverF0 && containsAll(a.Response.Supports, []string{SupportReducedB2Compatible, SupportZeroCubicFirewalls}) && containsAll(a.Response.Failures, []string{FailureReducedB2NotNativeFunctional, FailureNoNativeDegreeToZ2FlagFunctor}), Detail: FormatResponse(a.Response)},
			{Name: "cross-lanes remain excluded only conditionally on a certified I_B_Z2 functor", Passed: !a.CrossLane.DegreeOneToDegreeTwoAllowed && !a.CrossLane.DegreeTwoToDegreeOneAllowed && a.CrossLane.ExcludedIfFunctorCertified && !a.CrossLane.NativeCrossLaneTheorem && !a.CrossLane.NativeLinearDomainClassExclusion && !a.CrossLane.NativeQuadraticFaceClassExclusion && containsAll(a.CrossLane.Supports, []string{SupportCrossLanesIfFunctor}) && containsAll(a.CrossLane.Failures, []string{FailureNoNativeZ2CrossLane, FailureNoLinearActiveDomainExclusion, FailureNoQuadraticFaceExclusion}), Detail: FormatCrossLane(a.CrossLane)},
			{Name: "BoundaryAlpha reconstructs on the Z2 class and no longer depends on phase representative", Passed: !a.Alpha.RepresentativeAlphaRequired && a.Alpha.Z2ClassAlphaSupported && !a.Alpha.NativeAlphaCertified && a.Alpha.RepresentativeIndependent && a.Alpha.SealWeakenedToClassSeal && near(a.Alpha.ReconstructedAlpha, AlphaB) && a.Alpha.RankPair[0] == RankF1OverF0 && a.Alpha.RankPair[1] == RankF2OverF0 && containsAll(a.Alpha.Supports, []string{SupportAlphaNoPhaseRepresentative, SupportAlphaSealWeakensZ2}) && containsAll(a.Alpha.Failures, []string{FailureAlphaStillSealed, FailureNoNativeBoundaryAlphaSource, FailureNoNativeTransportS}), Detail: FormatAlpha(a.Alpha)},
			{Name: "R3 consequence reduces pressure to native Z2 BoundaryAlpha functor but does not certify native R3", Passed: a.R3.R3LedgerOnZ2AirlockClass && !a.R3.PhaseSignBlocksTraceLedger && !a.R3.NativeZ2AlphaFunctor && !a.R3.NativeR3 && !a.R3.FullAFDescent && !a.R3.OfficialLedgerUpdate && containsAll(a.R3.Supports, []string{SupportR3OnZ2AirlockClass, SupportPhaseSignNoLongerBlocksR3, SupportR3PressureToZ2AlphaFunctor}) && containsAll(a.R3.Failures, []string{FailureNotNativeR3, FailureFullAFDescentStillBlocked, FailureNoPhysicalParticleAssign, FailureNoIndividualYukawaValues}), Detail: FormatR3(a.R3)},
			{Name: "operator diagnostics remain separated from official frozen ledgers", Passed: a.Freeze.Frozen && a.Freeze.DiagnosticOnly && !a.Freeze.CanUpdate && near(a.Freeze.OperatorNEff, OperatorNEffDiagnostic) && near(a.Freeze.OperatorCYukawa, OperatorCYukawaDiagnostic) && near(a.Freeze.OperatorCHiggs, OperatorCHiggsDiagnostic) && !near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff) && containsAll(a.Freeze.Failures, []string{FailureNoOfficialNEffUpdate}), Detail: FormatFreeze(a.Freeze)},
			{Name: "native functor/source, response, cross-lane, Higgs, full descent, physical, generation/flavor, Yukawa, and official-update firewalls preserved", Passed: firewallsOK(a.Firewalls) && containsAll(a.FirewallsList(), []string{FailureNoNativeZ2AirlockFunctor, FailureNoNativeZ2BoundaryAlphaFunctor, FailureReducedB2NotNativeFunctional, FailureNoNativeBoundaryAlphaSource, FailureNoNativeTransportS, FailureNotNativeR3, FailureNoNativeYukawaOperator}), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, a.Classification, a.ShortStatus, FormatInherited(a.Inherited), FormatWellDefined(a.WellDefined), FormatResponse(a.Response), FormatCrossLane(a.CrossLane), FormatAlpha(a.Alpha), FormatR3(a.R3), FormatFreeze(a.Freeze), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
