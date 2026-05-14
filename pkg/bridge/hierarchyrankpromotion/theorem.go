package hierarchyrankpromotion

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Rank56HalfInstantonHierarchyPromotionSieveTheorem() theorem.Theorem {
	const id = "BRIDGE-RANK-56-HALF-INSTANTON-HIERARCHY-PROMOTION-SIEVE"
	const name = "Rank-56 / Half-Instanton Hierarchy Promotion Sieve"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 340 hierarchy promotion sieve", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 339 hierarchy audit inherited", Passed: a.Inputs.HighestInheritedGate == inheritedHighestGate && a.Inputs.BooleanProjectorRank == booleanProjectorRank && a.Inputs.STop > 78, Detail: FormatInputs(a.Inputs)},
			{Name: "effective hierarchy exponents computed", Passed: a.Targets.Log2Unreduced > 55 && a.Targets.Log2Unreduced < 56 && a.Targets.Log2Reduced > 53 && a.Targets.Log2Reduced < 54, Detail: FormatTargets(a.Targets)},
			{Name: "rank-56 near miss audited but not promoted", Passed: a.Ledger.Rank56.Name != "" && a.Ledger.Rank56.RatioToUnreduced > 0.68 && a.Ledger.Rank56.RatioToUnreduced < 0.70 && !a.Ledger.Rank56.Promotable, Detail: FormatCandidate(a.Ledger.Rank56)},
			{Name: "half-topological action audited but not promoted", Passed: a.Ledger.HalfTopological.Name != "" && a.Ledger.HalfTopological.RatioToUnreduced > 0.35 && a.Ledger.HalfTopological.RatioToUnreduced < 0.36 && !a.Ledger.HalfTopological.Promotable, Detail: FormatCandidate(a.Ledger.HalfTopological)},
			{Name: "prefactor sieve rejects accidental repair", Passed: !a.Prefactors.NativePrefactorDerived && a.Prefactors.SqrtTwoRank56Ratio > 0.97 && a.Prefactors.SqrtTwoRank56Ratio < 0.98, Detail: FormatPrefactors(a.Prefactors)},
			{Name: "category firewalls preserved", Passed: !a.Firewalls.RankExponentControlsMassScale && !a.Firewalls.HalfTopologicalActionControlsVEV && !a.Firewalls.PrefactorSelectedByFiniteGeometry && a.Firewalls.ArbitraryExponentFittingRejected, Detail: FormatFirewalls(a.Firewalls)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth, FormatSummary(a.Summary), FormatStatuses(Statuses(a))}}
	}}
}
