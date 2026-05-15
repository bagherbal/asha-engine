package generation2massliftbridge

import (
	"fmt"
	"strings"
)

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("Gate444KForced=%t gen2BareZero=%t noColliderData=%t KXYStillFree=%t nativeFlavorDim=%d conditionalKXYDim=%d verdict=%s", x.Gate444KGenForced, x.Gate444Generation2BareZero, x.Gate444NoColliderData, x.Gate444KXYStillFree, x.NativeFlavorDim, x.ConditionalKXYDim, x.Verdict)
}

func FormatArena(x BridgeArena) string {
	return fmt.Sprintf("ansatz=%q epsilonSymbolic=%t hermitian=%t zeroDiagonal=%t traceless=%t familyFiber=%t primitiveScan=%t noYukawa=%t verdict=%s", x.BridgeAnsatz, x.EpsilonSymbolic, x.Hermitian, x.ZeroDiagonal, x.Traceless, x.ActsOnlyOnFamilyFiber, x.IntegerPrimitiveScan, x.NoYukawaImported, x.Verdict)
}

func FormatBoundary(x Boundary) string {
	return fmt.Sprintf("%s formula=%q applied=%t passed=%t verdict=%s reason=%s", x.Name, x.Formula, x.Applied, x.Passed, x.Verdict, x.Reason)
}

func FormatCandidate(x Candidate) string {
	return fmt.Sprintf("weights=%s support=%d gcd=%d balanced=%t connected=%t closed=%t open=%t det=%s nonzero=%t balancedLift=%t canonical=%t", x.Weights.String(), x.SupportEdges, x.GCD, x.EndpointBalanced, x.Connected, x.ClosedTriangle, x.OpenChain, x.DeterminantFormula, x.DeterminantNonZero, x.BalancedMassLift, x.CanonicalTopology)
}

func FormatCandidates(xs []Candidate) string {
	parts := make([]string, len(xs))
	for i, x := range xs {
		parts[i] = FormatCandidate(x)
	}
	return strings.Join(parts, " | ")
}

func FormatSieve(x Sieve) string {
	return fmt.Sprintf("radius=%d raw=%d primitive=%d balanced=%d balancedLift=%d openFailures=%d unbalancedLift=%d uniqueUnsigned=%t signedVariants=%d verdict=%s", x.WeightRadius, x.RawCandidates, len(x.PrimitiveOffDiagonal), len(x.EndpointBalancedCandidates), len(x.BalancedLiftCandidates), len(x.OpenChainFailures), len(x.UnbalancedLiftCandidates), x.UniqueUnsignedTopology, x.SignedVariants, x.Verdict)
}

func FormatCollapse(x DeterminantCollapse) string {
	return fmt.Sprintf("K=%q bridge=%q det=%q balance=%q reduction=%q open=%q triangle=%q middleOrder=%q triangleForced=%t xSupport=%t fixesAmplitude=%t fixesSign=%t verdict=%s", x.KGen, x.Bridge, x.DeterminantIdentity, x.EndpointBalance, x.BalancedReduction, x.OpenChainResult, x.TriangleResult, x.MiddleEigenvalueOrder, x.ForcesClosedTriangle, x.ForcesXGenSupport, x.FixesAmplitude, x.FixesSignedOrientation, x.Verdict)
}

func FormatAxiom(x BridgeAxiom) string {
	return fmt.Sprintf("%s topology=%q support=%d trace=%d traceSquare=%d detLeading=%q gen2DiagZero=%t liftsZero=%t forcedTopology=%t amplitudeSealed=%t signSealed=%t yukawaPredicted=%t muonCharmPredicted=%t verdict=%s", x.Name, x.TopologyName, x.SupportEdges, x.Trace, x.TraceSquare, x.DeterminantLeadingOrder, x.Generation2DiagonalStillZero, x.LiftsGeneration2Zero, x.GeometricallyForcedTopology, x.AmplitudeEmpirical, x.SignedOrientationEmpirical, x.YukawaValuesPredicted, x.MuonCharmMassPredicted, x.Verdict)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("muonImported=%t charmImported=%t yukawaImported=%t CKM=%t PMNS=%t topologyConditional=%t amplitudeSealed=%t signedPhaseSealed=%t physicalMassNeedsData=%t nativeDim=%d→%d KXYFree=%d verdict=%s", !x.NoObservedMuonMassImported, !x.NoObservedCharmMassImported, !x.NoObservedYukawaImported, !x.NoCKMImported, !x.NoPMNSImported, x.BridgeTopologyNativeConditional, x.BridgeAmplitudeSealed, x.SignedPhaseSealed, x.PhysicalMassRequiresBridgeData, x.NativeFlavorDimBefore, x.NativeFlavorDimAfter, x.KXYCoeffDimStillFree, x.Verdict)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s", x.Gate, x.Title, x.Reason)
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 445 Registry Audit — Seesaw Bridge Mass-Lift / Structural-Zero Compatibility Audit\n\n")
	b.WriteString("## Claim tested\n\n")
	b.WriteString("Gate 445 tests whether the Gate-444 Generation-2 structural zero can be lifted by a purely off-diagonal family bridge without inserting empirical Yukawa values, muon/charm masses, CKM data, or PMNS data. The gate selects topology only; it does not select a bridge amplitude.\n\n")

	b.WriteString("## Prior boundary inherited\n\n")
	b.WriteString(FormatInheritance(a.Inheritance) + "\n\n")

	b.WriteString("## Bridge arena\n\n")
	b.WriteString(FormatArena(a.Arena) + "\n\n")
	b.WriteString("The tested bridge is:\n\n")
	b.WriteString("```text\n")
	b.WriteString("K_gen = diag(-1, 0, 1)\n")
	b.WriteString("B(a,b,c) = [[0,a,c],[a,0,b],[c,b,0]]\n")
	b.WriteString("M(ε) = K_gen + ε B(a,b,c)\n")
	b.WriteString("```\n\n")

	b.WriteString("## Boundary stack\n\n")
	b.WriteString("| Boundary | Formula | Applied | Passed | Verdict | Reason |\n")
	b.WriteString("|---|---|---:|---:|---|---|\n")
	for _, x := range a.Boundaries {
		b.WriteString(fmt.Sprintf("| %s | `%s` | %t | %t | `%s` | %s |\n", x.Name, x.Formula, x.Applied, x.Passed, x.Verdict, x.Reason))
	}
	b.WriteString("\n")

	b.WriteString("## Sieve enumeration\n\n")
	b.WriteString(FormatSieve(a.Sieve) + "\n\n")
	b.WriteString("The bounded scan uses primitive integer edge weights in `{−1,0,1}`. This is not an amplitude scan; it is a support-topology sieve.\n\n")

	b.WriteString("### Endpoint-balanced mass-lift survivors\n\n")
	b.WriteString("| Weights | Support | Det polynomial | Closed triangle | Canonical topology |\n")
	b.WriteString("|---|---:|---|---:|---:|\n")
	for _, x := range a.Sieve.BalancedLiftCandidates {
		b.WriteString(fmt.Sprintf("| `%s` | %d | `%s` | %t | %t |\n", x.Weights.String(), x.SupportEdges, x.DeterminantFormula, x.ClosedTriangle, x.CanonicalTopology))
	}
	b.WriteString("\n")

	b.WriteString("### Open-chain failures\n\n")
	b.WriteString("| Weights | Support | Det polynomial | Failure status |\n")
	b.WriteString("|---|---:|---|---|\n")
	for _, x := range a.Sieve.OpenChainFailures {
		b.WriteString(fmt.Sprintf("| `%s` | %d | `%s` | `%s` |\n", x.Weights.String(), x.SupportEdges, x.DeterminantFormula, StatusFailedOpenChainNoLift))
	}
	b.WriteString("\n")

	b.WriteString("## Analytic determinant collapse\n\n")
	b.WriteString(FormatCollapse(a.Collapse) + "\n\n")
	b.WriteString("For `B(a,b,c)` the exact symbolic identity is:\n\n")
	b.WriteString("```text\n")
	b.WriteString("det(K_gen + ε B) = (b^2 - a^2) ε^2 + 2abc ε^3\n")
	b.WriteString("```\n\n")
	b.WriteString("Endpoint balance gives `|a|=|b|`, so the `ε²` term cancels. An open chain has `c=0`, hence determinant zero and the middle structural zero survives. A balanced nonzero lift therefore requires `a,b,c` all nonzero. Primitive normalization reduces the unsigned support to the three-edge triangle.\n\n")

	b.WriteString("## Geometrically forced bridge topology\n\n")
	b.WriteString(FormatAxiom(a.Axiom) + "\n\n")
	b.WriteString("Canonical positive representative:\n\n")
	b.WriteString("```text\n")
	b.WriteString("B_lift = [[0,1,1],[1,0,1],[1,1,0]]\n")
	b.WriteString("det(K_gen + ε B_lift) = 2 ε^3\n")
	b.WriteString("```\n\n")
	b.WriteString("This is the real `X_gen` support topology. The gate does not fix the sign orientation, complex phase, or sector coefficient multiplying this topology.\n\n")

	b.WriteString("## Phenomenology/firewall audit\n\n")
	b.WriteString(FormatFirewall(a.Firewall) + "\n\n")
	b.WriteString("The result is a topology theorem, not a mass theorem. It says the second-family bare zero is lift-compatible through a closed triangular mixing bridge. The observed muon/charm mass values still require bridge amplitude and sector-source data outside Gate 445.\n\n")

	b.WriteString("## Result statuses\n\n")
	for _, s := range []string{StatusGate444StructuralZeroInherited, StatusBridgeArenaFormalized, StatusEndpointBalanceBoundaryApplied, StatusDeterminantLiftBoundaryApplied, StatusTriangleBridgeTopologyForced, StatusSeesawMassLiftCompatible, StatusXGenSupportSelectedAsTopology, StatusEmpiricalFirewallPreserved, StatusFailedOpenChainNoLift, StatusFailedUnbalancedBridgeLopsided, StatusFailedAmplitudeNotPredicted, StatusFailedSignedCycleOrientationUnfixed, StatusFailedNoMuonCharmMassPrediction, StatusFlavorFirewallPartiallyRefined} {
		b.WriteString("- `" + s + "`\n")
	}

	b.WriteString("\n## Next gate\n\n")
	b.WriteString(FormatNext(a.Next) + "\n\n")
	b.WriteString("## Truth statement\n\n")
	b.WriteString(a.Truth + "\n")
	return b.String()
}
