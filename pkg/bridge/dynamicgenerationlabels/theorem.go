package dynamicgenerationlabels

import "github.com/bagherbal/asha-engine/pkg/theorem"

func RepresentationOriginDynamicGenerationLabelsTheorem() theorem.Theorem {
	const id = "BRIDGE-REPRESENTATION-ORIGIN-DYNAMIC-GENERATION-LABELS"
	const name = "Representation-Origin Search for Dynamic Generation Labels"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 395 audit", Passed: false, Detail: err.Error()}}}
		}
		chirality := findLabel(a.Labels.Candidates, "spinor chirality split")
		branch := findLabel(a.Labels.Candidates, "triality representation-type triple")
		n := findLabel(a.Labels.Candidates, "sealed branch number operator")
		broadcast := findLabel(a.Labels.Candidates, "finite-Dirac generation broadcast")
		checks := []theorem.Check{
			{Name: "inherits Gate 394 centrality firewall", Passed: a.Inheritance.Gate394CentralityFirewall && a.Inheritance.Gate394NativeNoncentralOperators == 0 && a.Inheritance.Gate394NativeNoncommutingPairs == 0, Detail: FormatInheritance(a.Inheritance)},
			{Name: "Cl(1,7) spinor decomposition is 16 = 8 + 8", Passed: a.Spinor.FullSpinorRealDimension == 16 && len(a.Spinor.ChiralSplit) == 2 && a.Spinor.ChiralSplit[0] == 8 && a.Spinor.ChiralSplit[1] == 8, Detail: FormatSpinor(a.Spinor)},
			{Name: "spinor split does not derive three generation labels", Passed: !a.Spinor.HasThreeNativeSectors && !a.Spinor.GenerationLabelsDerived && chirality.Native && chirality.SectorCount == 2 && !chirality.GenerationLabelsDerived, Detail: FormatLabel(chirality)},
			{Name: "triality is category-level arena, not native finite-Dirac flavor carrier", Passed: a.Triality.CategoryLevelTriple && !a.Triality.NativeFunctorToC3Gen && !a.Triality.ExplicitThetaOnFiniteDiracFlavor && branch.Sealed && branch.Circular && !branch.Native, Detail: FormatTriality(a.Triality) + "\n" + FormatLabel(branch)},
			{Name: "sealed branch operators show capacity but remain quarantined", Passed: branch.NonCentral && branch.Mixing && n.NonCentral && n.DiagonalOnly && a.Labels.SealedNoncentralCount >= 2, Detail: FormatLabel(n)},
			{Name: "native finite-Dirac broadcast remains central", Passed: broadcast.Native && broadcast.Central && !broadcast.GenerationLabelsDerived, Detail: FormatLabel(broadcast)},
			{Name: "no native dynamic generation labels were derived", Passed: a.Labels.NativeGenerationLabelCount == 0 && a.Labels.NativeNoncentralCount == 1 && !a.Operators.CKMCapacityNative, Detail: FormatLabels(a.Labels)},
			{Name: "no native noncommuting texture pair exists", Passed: a.Operators.NativeNoncommutingPairs == 0 && a.Operators.MaxNativeCommutatorNorm < eps && !a.Operators.CKMCapacityNative, Detail: FormatOperators(a.Operators)},
			{Name: "sealed noncommuting branch capacity is not promoted", Passed: a.Operators.SealedNoncommutingPairs > 0 && a.Operators.MaxSealedCommutatorNorm > eps && a.Firewall.NoTrialityLabelsPromoted && a.Firewall.NoNPromoted, Detail: FormatOperators(a.Operators)},
			{Name: "13 charged moduli firewall remains preserved", Passed: a.Moduli.StartingChargedDim == 13 && !a.Moduli.NativeReductionBelow13 && a.Moduli.BestNativeDim == 13, Detail: FormatModuli(a.Moduli)},
			{Name: "firewalls remain clean", Passed: a.Firewall.NoMassesImported && a.Firewall.NoCKMImported && a.Firewall.NoPMNSImported && a.Firewall.NoManualGenerationAssignment && a.Firewall.NoNativeFlavorClaimed && a.Firewall.NoModuliReductionClaimed, Detail: FormatFirewall(a.Firewall)},
			{Name: "next gate moves beyond spinor chirality to endogenous three-object source", Passed: a.Next.Gate == 396 && a.Next.Title != "", Detail: FormatNext(a.Next)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: checks, Notes: []string{a.Truth}}
	}}
}
