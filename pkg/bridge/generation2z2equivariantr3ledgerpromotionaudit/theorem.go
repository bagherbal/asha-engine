package generation2z2equivariantr3ledgerpromotionaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_Z2_EQUIVARIANT_R3_LEDGER_PROMOTION_AUDIT"
	theoremName = "Gate 908 — Z2-Equivariant R3 Ledger Promotion Audit"
)

func Generation2Z2EquivariantR3LedgerPromotionAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 907 phase-orientation gauge candidate inherited", Passed: a.Inherited.QPhiSignGaugeCandidate && !a.Inherited.NativeZ2Certified && containsAll(a.Inherited.Supports, []string{SupportPhaseSignGaugeAggregate}) && containsAll(a.Inherited.Failures, []string{FailureNoNativeGlobalZ2}), Detail: FormatInherited(a.Inherited)},
			{Name: "orientation-class airlock preserves BoundaryAlpha flag ranks", Passed: a.Airlock.RanksInvariant && a.Airlock.AlphaRankSourceClassLevel && !a.Airlock.RequiresAbsolutePhaseSign && !a.Airlock.NativeZ2AirlockFunctor && containsAll(a.Airlock.Supports, []string{SupportAirlockDescendsZ2, SupportBoundaryAlphaFlagZ2Invariant, SupportAlphaRankNoAbsolutePhaseSign}) && containsAll(a.Airlock.Failures, []string{FailureNoNativeZ2AirlockFunctor, FailureAlphaStillSealed}), Detail: FormatAirlock(a.Airlock)},
			{Name: "edge ledger descends to Z2 orientation class with invariant rank and kernel", Passed: a.Edge.TauExchangesRepresentatives && a.Edge.RankPatternsInvariant && a.Edge.ImageRankInvariant && a.Edge.LeftKernelInvariant && a.Edge.OrientationClassLedgerExists && !a.Edge.NativeOperatorTheorem && containsAll(a.Edge.Supports, []string{SupportEdgeLedgerDescendsZ2, SupportEdgeRankKernelInvariant}), Detail: FormatEdge(a.Edge)},
			{Name: "trace-magnitude row multiset and operator diagnostics are Z2 invariant", Passed: a.Trace.TraceInvariant && a.Trace.SquareInvariant && a.Trace.NEffInvariant && a.Trace.CYukawaInvariant && a.Trace.CHiggsInvariant && a.Trace.DescendsToZ2Class && containsAll(a.Trace.Supports, []string{SupportTraceMultisetDescendsZ2, SupportNEffZ2Invariant, SupportCYukawaZ2Invariant}), Detail: FormatTrace(a.Trace)},
			{Name: "R3 trace-ledger requirements can be restated without absolute phase sign", Passed: a.R3.ProjectorLedgerZ2Class && a.R3.PositiveReadoutOnZ2Class && a.R3.TraceReconstructionOnZ2Class && a.R3.RequirementsRestatedWithoutSign && !a.R3.NativeSourceCertified && !a.R3.FullNativeR3 && containsAll(a.R3.Supports, []string{SupportProjectorLedgerZ2Class, SupportPositiveReadoutZ2Class, SupportZ2ClassReconstructsNEff, SupportR3TraceNoAbsolutePhaseSign}) && containsAll(a.R3.Failures, []string{FailureNoNativeZ2AirlockFunctor, FailureNotNativeR3}), Detail: FormatR3(a.R3)},
			{Name: "socket, physical-sector, generation, flavor, and individual-Yukawa labels do not descend", Passed: !a.NonDescending.SocketNamesDescend && !a.NonDescending.PhysicalSectorLabelsDescend && !a.NonDescending.GenerationLabelsDescend && !a.NonDescending.FlavorLabelsDescend && !a.NonDescending.IndividualYukawaValuesDescend && a.NonDescending.AggregateTraceLedgerDescends && containsAll(a.NonDescending.Failures, []string{FailureSocketLabelsNotPhysical, FailureNoPhysicalParticleAssign, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues}), Detail: FormatNonDescending(a.NonDescending)},
			{Name: "operator diagnostics remain separated from official frozen ledgers", Passed: a.Freeze.Frozen && a.Freeze.DiagnosticOnly && !a.Freeze.CanUpdate && near(a.Freeze.OperatorNEff, OperatorNEffDiagnostic) && near(a.Freeze.OperatorCYukawa, OperatorCYukawaDiagnostic) && !near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff) && containsAll(a.Freeze.Failures, []string{FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate}), Detail: FormatFreeze(a.Freeze)},
			{Name: "native Z2, native airlock, alpha, Higgs, full descent, R3/R4, physical assignment, generation/flavor, Yukawa, and official-update firewalls preserved", Passed: firewallsOK(a.Firewalls) && containsAll(a.FirewallsList(), []string{FailureNoNativeGlobalZ2, FailureNoNativeZ2AirlockFunctor, FailureAlphaStillSealed, FailureNoNativeBoundaryIncidence, FailureHiggsOrientationClassSealed, FailureFullAFDescentStillBlocked, FailureNotNativeR3, FailureNoNativeYukawaOperator}), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, a.Classification, a.ShortStatus, FormatInherited(a.Inherited), FormatAirlock(a.Airlock), FormatEdge(a.Edge), FormatTrace(a.Trace), FormatR3(a.R3), FormatNonDescending(a.NonDescending), FormatFreeze(a.Freeze), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
