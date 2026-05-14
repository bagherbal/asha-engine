package ko6twistedrealstructure

import "github.com/bagherbal/asha-engine/pkg/theorem"

func KO6TwistedRealStructurePhysicalJDerivationAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-KO6-TWISTED-REAL-STRUCTURE-PHYSICAL-J-DERIVATION-AUDIT"
	const name = "KO-6 Twisted Real Structure / Physical J_F Derivation Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 293 KO6 twisted J audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 292 KO0-like fiber J is inherited", Passed: a.Input.J2Sign == 1 && a.Input.JGammaSign == 1 && !a.Input.KOSixLike, Detail: FormatInput(a.Input)},
			{Name: "twist candidates are constructed and audited", Passed: len(a.Twists.Candidates) >= 4, Detail: FormatTwistAudit(a.Twists)},
			{Name: "even grading/volume twist does not produce KO6", Passed: !a.Twists.Candidates[1].KOSixLike && a.Twists.Candidates[1].JGammaSign == +1, Detail: FormatCandidate(a.Twists.Candidates[1])},
			{Name: "odd one-mode twists satisfy KO6 signs", Passed: a.Twists.KO6Candidates == 2, Detail: FormatTwistAudit(a.Twists)},
			{Name: "no canonical odd twist selector is derived", Passed: !a.Twists.CanonicalKO6Found, Detail: a.Twists.Degeneracy},
			{Name: "JD=DJ is a sieve, not a canonical D_F selector", Passed: len(a.DiracSieve) == 2 && a.DiracSieve[0].JDRealityFreeParams == 3 && !a.DiracSieve[0].CanonicalDFSelected, Detail: FormatDiracSieve(a.DiracSieve[0]) + " || " + FormatDiracSieve(a.DiracSieve[1])},
			{Name: "formal doubled swap has KO6 signs but lacks physical representation", Passed: a.DoubledSwap.KOSixLike && !a.DoubledSwap.PhysicalRepAvailable && !a.DoubledSwap.OppositeActionPossible, Detail: FormatDoubled(a.DoubledSwap)},
			{Name: "opposite algebra action remains missing", Passed: !a.Opposite.OppositeActionConstructed && !a.Opposite.PhysicalJAvailable, Detail: FormatOpposite(a.Opposite)},
			{Name: "firewalls keep Path B and Path C blocked", Passed: !a.Firewalls.FiniteCorePolluted && a.Firewalls.DoesNotUnlockHiggs && a.Firewalls.DoesNotUnlockBGap, Detail: FormatFirewalls(a.Firewalls)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth, FormatSummary(a.Summary)}}
	}}
}
