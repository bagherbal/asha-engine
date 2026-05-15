package generation2massliftbridge

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2MassLiftBridgeStructuralZeroCompatibilitySieveTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 mass-lift bridge structural-zero compatibility sieve"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate445 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate444 structural-zero family axis", Passed: a.Inheritance.Executed && a.Inheritance.Gate444KGenForced && a.Inheritance.Gate444Generation2BareZero && a.Inheritance.Gate444NoColliderData && a.Inheritance.Gate444KXYStillFree, Detail: FormatInheritance(a.Inheritance)},
			{Name: "off-diagonal bridge arena preserves bare zero", Passed: a.Arena.Executed && a.Arena.Hermitian && a.Arena.ZeroDiagonal && a.Arena.Traceless && a.Arena.ActsOnlyOnFamilyFiber && a.Arena.NoYukawaImported, Detail: FormatArena(a.Arena)},
			{Name: "all mass-lift boundaries are applied", Passed: len(a.Boundaries) == 4 && a.Boundaries[0].Passed && a.Boundaries[1].Passed && a.Boundaries[2].Passed && a.Boundaries[3].Passed, Detail: FormatBoundary(a.Boundaries[0]) + " | " + FormatBoundary(a.Boundaries[1]) + " | " + FormatBoundary(a.Boundaries[2]) + " | " + FormatBoundary(a.Boundaries[3])},
			{Name: "sieve rejects balanced open chains and isolates triangle support", Passed: a.Sieve.Executed && a.Sieve.UniqueUnsignedTopology && len(a.Sieve.BalancedLiftCandidates) == 8 && len(a.Sieve.OpenChainFailures) == 4 && a.Sieve.SignedVariants == 8, Detail: FormatSieve(a.Sieve)},
			{Name: "determinant identity proves cubic seesaw lift", Passed: a.Collapse.Executed && a.Collapse.ForcesClosedTriangle && a.Collapse.ForcesXGenSupport && !a.Collapse.FixesAmplitude && !a.Collapse.FixesSignedOrientation, Detail: FormatCollapse(a.Collapse)},
			{Name: "bridge topology is forced but amplitude remains sealed", Passed: a.Axiom.Executed && a.Axiom.GeometricallyForcedTopology && a.Axiom.Generation2DiagonalStillZero && a.Axiom.LiftsGeneration2Zero && a.Axiom.AmplitudeEmpirical && a.Axiom.SignedOrientationEmpirical && !a.Axiom.MuonCharmMassPredicted, Detail: FormatAxiom(a.Axiom)},
			{Name: "empirical flavor firewall remains intact", Passed: a.Firewall.Executed && a.Firewall.NoObservedMuonMassImported && a.Firewall.NoObservedCharmMassImported && a.Firewall.NoObservedYukawaImported && a.Firewall.NoCKMImported && a.Firewall.NoPMNSImported && a.Firewall.BridgeAmplitudeSealed && a.Firewall.NativeFlavorDimAfter == NativeFlavorDim && a.Firewall.KXYCoeffDimStillFree == KXYCoeffDim, Detail: FormatFirewall(a.Firewall)},
			{Name: "next gate audits signed/complex orientation", Passed: a.Next.Gate == 446 && a.Next.Title == "Signed-Cycle / Complex Phase Orientation Sieve", Detail: FormatNext(a.Next)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth}}
	}}
}
