package generation2structuralzero

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2StructuralZeroIntersectionSieveTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 structural zero intersection sieve"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate444 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate420/Gate412 flavor boundary", Passed: a.Inheritance.Executed && a.Inheritance.Gate420PublicationAtlasRead && a.Inheritance.Gate420FlavorFirewall && a.Inheritance.Gate420NativeFlavorDim == NativeFlavorDim && a.Inheritance.Gate412Traceless && a.Inheritance.Gate412ThreeLevel && a.Inheritance.NoEmpiricalInputsImported, Detail: FormatInheritance(a.Inheritance)},
			{Name: "all three boundary walls are applied", Passed: len(a.Boundaries) == 3 && a.Boundaries[0].Passed && a.Boundaries[1].Passed && a.Boundaries[2].Passed, Detail: FormatBoundary(a.Boundaries[0]) + " | " + FormatBoundary(a.Boundaries[1]) + " | " + FormatBoundary(a.Boundaries[2])},
			{Name: "sieve enumeration isolates scale family", Passed: a.Enumeration.Executed && a.Enumeration.RejectedDegenerateZero && a.Enumeration.OnlyScaleVariants && len(a.Enumeration.PrimitivePassing) == 1 && a.Enumeration.PrimitivePassing[0].CanonicalMinimal, Detail: FormatEnumeration(a.Enumeration)},
			{Name: "analytic collapse proves unique primitive triplet", Passed: a.Collapse.Executed && a.Collapse.UniqueMinimal && a.Collapse.UniqueUpToPermutation && a.Collapse.UniqueUpToSign && a.Collapse.ForcesMiddleZero && a.Collapse.PrimitiveSolution.String() == (Spectrum{-1, 0, 1}).String(), Detail: FormatCollapse(a.Collapse)},
			{Name: "K_gen geometrically forced axiom installed", Passed: a.Axiom.Executed && a.Axiom.GeometricallyForced && a.Axiom.Generation2BareZero && a.Axiom.MiddleEigenvalue == 0 && a.Axiom.Trace == 0 && a.Axiom.TraceSquare == 2 && !a.Axiom.ColliderDataUsed && !a.Axiom.ScaleEmpirical, Detail: FormatAxiom(a.Axiom)},
			{Name: "empirical flavor firewall preserved", Passed: a.Firewall.Executed && a.Firewall.NoObservedMuonMassImported && a.Firewall.NoObservedCharmMassImported && a.Firewall.NoObservedYukawaImported && a.Firewall.NoCKMImported && a.Firewall.NoPMNSImported && a.Firewall.BareStructuralStatementOnly && a.Firewall.PhysicalMassRequiresBridgeData && a.Firewall.NativeFlavorDimAfter == NativeFlavorDim && a.Firewall.KXYCoeffDimStillFree == KXYCoeffDim, Detail: FormatFirewall(a.Firewall)},
			{Name: "next gate audits mass-lift bridge", Passed: a.Next.Gate == 445 && a.Next.Title == "Seesaw Bridge Mass-Lift / Structural-Zero Compatibility Audit", Detail: FormatNext(a.Next)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth}}
	}}
}
