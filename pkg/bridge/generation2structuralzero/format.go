package generation2structuralzero

import (
	"fmt"
	"strings"
)

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("Gate420Atlas=%t nativeFlavorDim=%d conditionalKXYDim=%d flavorFirewall=%t Gate412K=%s traceless=%t threeLevel=%t quarantined=%t empiricalImported=%t verdict=%s", x.Gate420PublicationAtlasRead, x.Gate420NativeFlavorDim, x.Gate420ConditionalKXYDim, x.Gate420FlavorFirewall, x.Gate412KGen.String(), x.Gate412Traceless, x.Gate412ThreeLevel, x.Gate412Quarantined, !x.NoEmpiricalInputsImported, x.Verdict)
}

func FormatBoundary(x Boundary) string {
	return fmt.Sprintf("%s formula=%q applied=%t passed=%t verdict=%s reason=%s", x.Name, x.Formula, x.Applied, x.Passed, x.Verdict, x.Reason)
}

func FormatCandidate(x Candidate) string {
	return fmt.Sprintf("spectrum=%s trace=%d gaps=(%d,%d) gcd=%d distinct=%t traceless=%t even=%t primitive=%t passes=%t canonical=%t", x.Spectrum.String(), x.Trace, x.Gap01, x.Gap12, x.GCD, x.Distinct, x.Traceless, x.EvenlySpaced, x.Primitive, x.PassesAll, x.CanonicalMinimal)
}

func FormatCandidates(xs []Candidate) string {
	parts := make([]string, len(xs))
	for i, x := range xs {
		parts[i] = FormatCandidate(x)
	}
	return strings.Join(parts, " | ")
}

func FormatEnumeration(x Enumeration) string {
	return fmt.Sprintf("radius=%d raw=%d uniqueSorted=%d passing=%d primitive=%d rejectedZero=%t onlyScaleVariants=%t verdict=%s", x.SearchRadius, x.RawTriplesVisited, x.UniqueSortedCandidates, len(x.PassingFamilies), len(x.PrimitivePassing), x.RejectedDegenerateZero, x.OnlyScaleVariants, x.Verdict)
}

func FormatCollapse(x AnalyticCollapse) string {
	return fmt.Sprintf("ansatz=%q traceEq=%q family=%q primitive=%q primitiveSolution=%s arbitraryScale=%t uniquePermutation=%t uniqueSign=%t uniqueMinimal=%t middleZero=%t verdict=%s", x.SortedAnsatz, x.TracelessEquation, x.SolutionFamily, x.PrimitiveRule, x.PrimitiveSolution.String(), x.ArbitraryScale, x.UniqueUpToPermutation, x.UniqueUpToSign, x.UniqueMinimal, x.ForcesMiddleZero, x.Verdict)
}

func FormatAxiom(x Axiom) string {
	return fmt.Sprintf("%s spectrum=%s trace=%d traceSquare=%d rank=%d gen2=%d bareZero=%t forced=%t scaleEmpirical=%t colliderData=%t yukawaPredicted=%t mixingPredicted=%t verdict=%s", x.Name, x.Spectrum.String(), x.Trace, x.TraceSquare, x.Rank, x.MiddleEigenvalue, x.Generation2BareZero, x.GeometricallyForced, x.ScaleEmpirical, x.ColliderDataUsed, x.YukawaValuesPredicted, x.MixingAnglesPredicted, x.Verdict)
}

func FormatFirewall(x PhenomenologyFirewall) string {
	return fmt.Sprintf("muonImported=%t charmImported=%t yukawaImported=%t CKM=%t PMNS=%t bareOnly=%t seesawBridge=%t physicalMassNeedsBridge=%t nativeDim=%d→%d KXYFree=%d verdict=%s", !x.NoObservedMuonMassImported, !x.NoObservedCharmMassImported, !x.NoObservedYukawaImported, !x.NoCKMImported, !x.NoPMNSImported, x.BareStructuralStatementOnly, x.SeesawBridgeInterpretation, x.PhysicalMassRequiresBridgeData, x.NativeFlavorDimBefore, x.NativeFlavorDimAfter, x.KXYCoeffDimStillFree, x.Verdict)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s", x.Gate, x.Title, x.Reason)
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 444 Registry Audit — Generation 2 Structural Zero / Intersection Sieve\n\n")
	b.WriteString("## Claim tested\n\n")
	b.WriteString("Gate 444 tests whether the diagonal family generator `K_gen` is merely a clean quarantined choice or is forced by the intersection of three existing boundaries: traceless family-source balance, modular/KMS integer spacing, and exactly three distinct generation eigenlevels. No collider masses, Yukawa matrices, CKM data, or PMNS data are imported.\n\n")

	b.WriteString("## Prior boundary inherited\n\n")
	b.WriteString(FormatInheritance(a.Inheritance) + "\n\n")

	b.WriteString("## Boundary stack\n\n")
	b.WriteString("| Boundary | Formula | Applied | Passed | Verdict | Reason |\n")
	b.WriteString("|---|---|---:|---:|---|---|\n")
	for _, x := range a.Boundaries {
		b.WriteString(fmt.Sprintf("| %s | `%s` | %t | %t | `%s` | %s |\n", x.Name, x.Formula, x.Applied, x.Passed, x.Verdict, x.Reason))
	}
	b.WriteString("\n")

	b.WriteString("## Sieve enumeration\n\n")
	b.WriteString(FormatEnumeration(a.Enumeration) + "\n\n")
	b.WriteString("The bounded search is only an implementation witness. The proof of uniqueness is the analytic collapse below; therefore the conclusion is not a finite-range artifact.\n\n")
	b.WriteString("### Passing spectra inside the audit radius\n\n")
	b.WriteString("| Spectrum | Trace | Gaps | GCD | Primitive | Canonical minimal |\n")
	b.WriteString("|---|---:|---:|---:|---:|---:|\n")
	for _, x := range a.Enumeration.PassingFamilies {
		b.WriteString(fmt.Sprintf("| `%s` | %d | `%d,%d` | %d | %t | %t |\n", x.Spectrum.String(), x.Trace, x.Gap01, x.Gap12, x.GCD, x.Primitive, x.CanonicalMinimal))
	}
	b.WriteString("\n")

	b.WriteString("## Analytic boundary collapse\n\n")
	b.WriteString(FormatCollapse(a.Collapse) + "\n\n")
	b.WriteString("Let the sorted integer-spaced three-level spectrum be `(a, a+q, a+2q)` with `q>0`. Tracelessness gives `3a+3q=0`, hence `a=-q`. Therefore every survivor is exactly `(-q,0,q)`. Without primitive quantization this leaves an arbitrary integer scale; with `gcd=1`, the unique minimal representative is `(-1,0,1)`, up to sign and permutation.\n\n")

	b.WriteString("## Geometrically forced axiom\n\n")
	b.WriteString(FormatAxiom(a.Axiom) + "\n\n")
	b.WriteString("Canonical matrix:\n\n")
	b.WriteString("```text\nK_gen = diag(-1, 0, 1)\n```\n\n")

	b.WriteString("## Phenomenology/firewall audit\n\n")
	b.WriteString(FormatFirewall(a.Firewall) + "\n\n")
	b.WriteString("Gate 444 refines the flavor frontier but does not erase it. The middle diagonal bare level is structurally zero, so the Generation-2 charged family is classified as a bare resonance/seesaw bridge location. Physical muon/charm mass values still require a mass-lift bridge and are not predicted by this gate.\n\n")

	b.WriteString("## Result statuses\n\n")
	for _, s := range []string{StatusGate420FirewallInherited, StatusGate412KGenBoundaryInherited, StatusTracelessBoundaryApplied, StatusKMSQuantizationBoundaryApplied, StatusThreeGenerationBoundaryApplied, StatusIntersectionCollapsed, StatusGen2StructuralZeroProved, StatusKGenGeometricallyForcedAxiom, StatusMuonSeesawResonanceDerived, StatusEmpiricalFirewallPreserved, StatusFailedScaleArbitraryWithoutPrimitive, StatusNoYukawaPrediction, StatusNoMixingPrediction, StatusFlavorFirewallPartiallyRefined} {
		b.WriteString("- `" + s + "`\n")
	}

	b.WriteString("\n## Next gate\n\n")
	b.WriteString(FormatNext(a.Next) + "\n\n")
	b.WriteString("## Truth statement\n\n")
	b.WriteString(a.Truth + "\n")
	return b.String()
}
