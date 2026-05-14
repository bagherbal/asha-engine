package realstructurekofactorization

import "github.com/bagherbal/asha-engine/pkg/theorem"

func RealStructureJFactorizationKODimensionAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-REAL-STRUCTURE-J-FACTORIZATION-KO-DIMENSION-AUDIT"
	const name = "Paths B & C Convergence / Real Structure J KO-Dimension Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 292 J factorization audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 234 occupation-complement J candidate is inherited as preflight only", Passed: a.Inherited.J2Sign == 1 && a.Inherited.JGammaSign == 1 && !a.Inherited.PhysicalChargeConjugation, Detail: FormatInherited(a.Inherited)},
			{Name: "Gate 3 spacetime/fiber Witt split is represented as 2+2 modes", Passed: len(a.Split.SpacetimeWittModes) == 2 && len(a.Split.FiberWittModes) == 2, Detail: FormatSplit(a.Split)},
			{Name: "full occupation complement factorizes across spacetime/fiber split", Passed: a.Factor.FullComplementMatchesTensor && a.Factor.Residual < 1e-12, Detail: FormatFactor(a.Factor)},
			{Name: "fiber J signs are computed explicitly", Passed: a.KO.J2Sign == 1 && a.KO.JGammaSign == 1 && a.KO.KOZeroLike, Detail: FormatKO(a.KO)},
			{Name: "fiber J is not KO6 / SM physical real structure", Passed: !a.KO.KOSixLike && a.Firewalls.DoesNotClaimKO6, Detail: FormatKO(a.KO)},
			{Name: "J-reality remains a sieve, not a canonical finite Dirac selector", Passed: a.DReality.JRealityFreeParams < a.DReality.GenericOddBlockParams && !a.DReality.CanonicalDFSelected, Detail: FormatDReality(a.DReality)},
			{Name: "opposite algebra action and dynamics remain blocked", Passed: !a.Opposite.OppositeActionConstructed && !a.Opposite.HeatKernelUnblocked && !a.Opposite.BGapInstantonUnblocked, Detail: FormatOpposite(a.Opposite)},
			{Name: "firewalls prevent KO/preflight data from unlocking Path B or Path C", Passed: !a.Firewalls.FiniteCorePolluted && a.Firewalls.DoesNotUnlockHiggs && a.Firewalls.DoesNotUnlockBGap, Detail: FormatFirewalls(a.Firewalls)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth, FormatSummary(a.Summary)}}
	}}
}
