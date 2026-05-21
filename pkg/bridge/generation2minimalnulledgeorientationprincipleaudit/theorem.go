package generation2minimalnulledgeorientationprincipleaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_MINIMAL_NULL_EDGE_ORIENTATION_PRINCIPLE_AUDIT"
	theoremName = "Gate 894 — MinimalNullEdge Orientation Principle Audit"
)

func Generation2MinimalNullEdgeOrientationPrincipleAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 893 weak-socket selector obstruction inherited", Passed: containsAll(Statuses(), []string{StatusGate893Inherited}) && a.NullEdge.YPlus1Zero, Detail: FormatNullEdgeMinimization(a.NullEdge)},
			{Name: "minimal null-edge candidate reduces right rectangle 8 to 7", Passed: a.NullEdge.AmbientRightRank == 8 && a.NullEdge.ActiveRightRank == 7 && a.NullEdge.RankReduction == 1 && a.NullEdge.MinimalRankSevenCandidate && !a.NullEdge.NativePrinciple, Detail: FormatNullEdgeMinimization(a.NullEdge)},
			{Name: "image/kernel reconstruction yields h_+ tensor P_1 candidate", Passed: a.ImageKernel.HLeftRank == 8 && a.ImageKernel.ImageRank == 7 && a.ImageKernel.KernelRank == 1 && a.ImageKernel.QuotientIsHPlusP1 && a.ImageKernel.ReconstructsKernel && !a.ImageKernel.SelectsFrameNonCircularly, Detail: FormatImageKernel(a.ImageKernel)},
			{Name: "noncircular selector functional remains missing", Passed: !a.NonCircularity.CanDefineHPlusFromKernelWithoutPriorFrame && a.NonCircularity.EdgeSupportAssumesFrame && a.NonCircularity.RequiresVariationalMinimality && !a.NonCircularity.NativeSelectorFunctional && containsAll(a.NonCircularity.Failures, []string{FailureNoNonCircularWeakSocketSelector, FailureNoNativeVariationalMinimalityTheorem}), Detail: FormatNonCircularity(a.NonCircularity)},
			{Name: "official diagnostics remain frozen", Passed: a.Freeze.Frozen && a.Freeze.DiagnosticOnly && !a.Freeze.CanUpdate && near(a.Freeze.OperatorNEff, OperatorNEffDiagnostic) && !near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff), Detail: FormatFreeze(a.Freeze)},
			{Name: "native R3/R4 and physical-sector firewalls preserved", Passed: firewallsOK(a.Firewalls) && containsAll(a.FirewallsList(), []string{FailureNotNativeR3, FailureNoNativeMinimalNullEdgeOrientation, FailureNoNonCircularWeakSocketSelector, FailureNoNativeDescentFullToOrient}), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatNullEdgeMinimization(a.NullEdge), FormatImageKernel(a.ImageKernel), FormatNonCircularity(a.NonCircularity), FormatFreeze(a.Freeze), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
