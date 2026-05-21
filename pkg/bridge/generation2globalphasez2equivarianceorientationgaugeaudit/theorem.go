package generation2globalphasez2equivarianceorientationgaugeaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_GLOBAL_PHASE_Z2_EQUIVARIANCE_ORIENTATION_GAUGE_AUDIT"
	theoremName = "Gate 907 — GlobalPhaseZ2 Equivariance and OrientationGauge Audit"
)

func Generation2GlobalPhaseZ2EquivarianceOrientationGaugeAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 906 Q_phi sign wound inherited", Passed: a.Inherited.QPhiExists && !a.Inherited.PositiveSignSelected && containsAll(a.Inherited.Supports, []string{SupportGate906Inherited}), Detail: FormatInherited(a.Inherited)},
			{Name: "global phase Z2 flips Q_phi and exchanges lambda/barlambda projectors", Passed: a.Z2.SwapsLambdaBarLambda && a.Z2.SwapsProjectors && a.Z2.FlipsQPhi && !a.Z2.NativeTheorem && containsAll(a.Z2.Supports, []string{SupportGlobalZ2Exchanges}) && containsAll(a.Z2.Failures, []string{FailureNoNativeGlobalZ2}), Detail: FormatZ2(a.Z2)},
			{Name: "airlock representatives and alpha ranks are Z2 invariant as an orientation class", Passed: a.Airlock.AlphaRanksInvariant && a.Airlock.OrientationClassCandidate && !a.Airlock.NativeEquivariance && sameInts(a.Airlock.PlusFlagRanks, a.Airlock.MinusFlagRanks) && containsAll(a.Airlock.Supports, []string{SupportGlobalZ2Exchanges, SupportAlphaInvariant, SupportAirlockClass}), Detail: FormatAirlock(a.Airlock)},
			{Name: "edge table has Z2 mirror with invariant rank and kernel count", Passed: a.Edge.MirrorRepresentativeExists && a.Edge.RankKernelInvariant && !a.Edge.NativeOperatorTheorem && sameInts(a.Edge.CurrentRankPattern, a.Edge.MirrorRankPattern) && a.Edge.CurrentKernelCount == a.Edge.MirrorKernelCount && containsAll(a.Edge.Supports, []string{SupportEdgeMirror, SupportEdgeRankKernel}), Detail: FormatEdge(a.Edge)},
			{Name: "trace row multiset and operator diagnostics are phase-orientation invariant", Passed: a.Trace.TraceInvariant && a.Trace.SquareTraceInvariant && a.Trace.NEffInvariant && a.Trace.CYukawaInvariant && !a.Trace.NativeTraceTheorem && containsAll(a.Trace.Supports, []string{SupportRowMultiset, SupportNEffInvariant, SupportCYukawaInvariant}), Detail: FormatTrace(a.Trace)},
			{Name: "socket labels are gauge dependent and trace ledger needs no physical labels", Passed: a.Labels.SocketLabelsChange && !a.Labels.PhysicalLabelsCertified && !a.Labels.TraceLedgerNeedsPhysical && containsAll(a.Labels.Supports, []string{SupportLabelsGauge, SupportNoPhysicalLabelsNeed}), Detail: FormatLabels(a.Labels)},
			{Name: "absolute Q_phi sign pressure weakens to Z2-equivariant airlock theorem", Passed: !a.Wound.AbsoluteSignNeeded && !a.Wound.NativeSolved && containsAll(a.Wound.Supports, []string{SupportQPhiGauge, SupportNativePressureZ2, SupportGeneratorWoundWeakens}) && containsAll(a.Wound.Failures, []string{FailureNoNativeGlobalZ2}), Detail: FormatWound(a.Wound)},
			{Name: "operator diagnostics remain coherent and official ledgers frozen", Passed: a.Freeze.Frozen && a.Freeze.DiagnosticOnly && !a.Freeze.CanUpdate && near(a.Freeze.OperatorNEff, OperatorNEffDiagnostic) && !near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff), Detail: FormatFreeze(a.Freeze)},
			{Name: "native R3/R4, native Z2, phase module, airlock, alpha, Higgs, full descent, physical-sector, generation/flavor, and official-ledger firewalls preserved", Passed: firewallsOK(a.Firewalls) && containsAll(a.FirewallsList(), []string{FailureNoNativeGlobalZ2, FailureNoNativePhaseModuleCR2, FailureNoNativeAirlock, FailureAlphaStillSealed, FailureHiggsStillSealed, FailureNotNativeR3}), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatInherited(a.Inherited), FormatZ2(a.Z2), FormatAirlock(a.Airlock), FormatEdge(a.Edge), FormatTrace(a.Trace), FormatLabels(a.Labels), FormatWound(a.Wound), FormatFreeze(a.Freeze), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
