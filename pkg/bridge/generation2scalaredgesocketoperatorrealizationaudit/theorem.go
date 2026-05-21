package generation2scalaredgesocketoperatorrealizationaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_SCALAR_EDGE_SOCKET_OPERATOR_REALIZATION_AUDIT"
	theoremName = "Gate 860 — Scalar Edge-Socket Operator Realization Audit"
)

func Generation2ScalarEdgeSocketOperatorRealizationAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "define operator-valued scalar edge-socket Y", Passed: a.Y.OperatorValued && a.Y.SymbolicSocketMatrix && a.Y.ColorCentrality && a.Y.LeptonTriviality && a.Y.PunctureZero && a.Y.RankIfActiveSocketsNonzero == YRankFull && containsAll(a.Y.Supports, []string{StatusYOperatorDefined, SupportOperatorYFirstOrderIfScalar}), Detail: FormatY(a.Y)},
			{Name: "realize color edges as scalar identity on P3", Passed: colorEdgesScalar(a.Edges), Detail: FormatEdges(a.Edges)},
			{Name: "realize lepton edge as scalar on P1 and keep puncture edge zero", Passed: leptonAndPunctureOK(a.Edges), Detail: FormatEdges(a.Edges)},
			{Name: "rebuild D_F^sym from Y and recompute rank/kernel ledger", Passed: a.D.BuiltFromY && a.D.SelfAdjointByBlockForm && a.D.ChiralBlockForm && a.D.PostOrientationOnly && a.D.RankIfActiveSocketsNonzero == DSymRankFull && a.D.KernelRankIfNonzero == KernelRank && a.D.KernelSingleton == "h_+ tensor P_1" && containsAll(a.D.Supports, []string{StatusDRebuilt, StatusRankKernelRecomputed}), Detail: FormatD(a.D)},
			{Name: "inherit Gate 859 edge-centrality first-order support position without operator theorem", Passed: a.FirstOrder.Gate859Inherited && a.FirstOrder.EdgeCentralityInstalled && !a.FirstOrder.OperatorLevelFirstOrderCertified && !a.FirstOrder.CompleteJOppositeCertified && containsAll(a.FirstOrder.Failures, []string{FailureNoFullOperatorFirstOrder, FailureNoCompleteJOppositeProof}), Detail: FormatFirstOrder(a.FirstOrder)},
			{Name: "freeze ledgers and prevent magnitude/R3/R4 promotion", Passed: a.Ledger.OfficialFrozen && !a.Ledger.AlphaNative && !a.Ledger.R3 && !a.Ledger.R4 && a.Impact.AlphaStillSealed && a.Impact.MagnitudesStillMissing && !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && !a.Impact.CanPromoteToR3 && !a.Impact.CanPromoteToR4, Detail: FormatLedger(a.Ledger) + " | " + FormatImpact(a.Impact)},
			{Name: "preserve Gate 860 scalar-edge socket firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.OperatorYSymbolicNotYukawa && a.Firewalls.NoNumericalYukawa && a.Firewalls.NoAlphaSource && a.Firewalls.NoTraceReadout && a.Firewalls.NoFullUnbrokenAFTheorem && a.Firewalls.AForientNotFullAF && a.Firewalls.NoFullFirstOrder && a.Firewalls.NoCompleteJOpposite && a.Firewalls.NoBimoduleProof && a.Firewalls.ScalarSocketSupportOnly && a.Firewalls.NoOfficialNEffUpdate && a.Firewalls.NoCYukawaCHiggsUpdate && a.Firewalls.NotR3 && a.Firewalls.NotR4 && a.Firewalls.NoParticleAssignment && a.Firewalls.NoNeutrinoTheorem && a.Firewalls.NoThreeGenerationTheorem && a.Firewalls.Verdict == StatusFirewallVerdict, Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatY(a.Y), FormatEdges(a.Edges), FormatD(a.D), FormatFirstOrder(a.FirstOrder), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}

func colorEdgesScalar(edges []EdgeOperator) bool {
	seen := map[string]bool{}
	for _, e := range edges {
		if e.ColorEdge {
			seen[e.Name] = true
			if !e.Present || !e.CoefficientSymbolic || !e.ScalarOnVisibleFactor || !e.IdentityOnVisibleFactor || e.NumericalValue || e.YukawaMagnitude || !containsAll(e.Supports, []string{StatusColorEdgesScalar, SupportColorCentralityRepairsM3}) {
				return false
			}
			if e.OperatorForm != "y_+3 |h_+><e_+| tensor I_{P_3}" && e.OperatorForm != "y_-3 |h_-><e_-| tensor I_{P_3}" {
				return false
			}
		}
	}
	return seen["Y_+3"] && seen["Y_-3"]
}

func leptonAndPunctureOK(edges []EdgeOperator) bool {
	lepton, puncture := false, false
	for _, e := range edges {
		if e.LeptonEdge {
			lepton = e.Name == "Y_-1" && e.Present && e.CoefficientSymbolic && e.ScalarOnVisibleFactor && e.IdentityOnVisibleFactor && !e.NumericalValue && !e.YukawaMagnitude && e.OperatorForm == "y_-1 |h_-><e_-| tensor I_{P_1}"
		}
		if e.PunctureEdge {
			puncture = e.Name == "Y_+1" && !e.Present && !e.CoefficientSymbolic && e.OperatorForm == "0" && e.ScalarOnVisibleFactor && !e.IdentityOnVisibleFactor
		}
	}
	return lepton && puncture
}
