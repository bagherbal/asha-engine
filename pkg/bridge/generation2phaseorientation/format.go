package generation2phaseorientation

import (
	"fmt"
	"strings"
)

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("Gate445K=%t XSupport=%t amplitudeSealed=%t signSealed=%t noEmpiricalMasses=%t nativeDim=%d KXYFree=%d verdict=%s", x.Gate445KGenForced, x.Gate445XSupportForced, x.Gate445AmplitudeSealed, x.Gate445SignedOrientationSealed, x.Gate445NoEmpiricalMasses, x.NativeFlavorDim, x.KXYCoeffDimStillFree, x.Verdict)
}

func FormatArena(x OrientationArena) string {
	return fmt.Sprintf("K=%q ansatz=%q hermitian=%t zeroDiagonal=%t triangle=%t endpointBalanced=%t vertexGauge=%t invariant=%q empiricalImported=%t verdict=%s", x.KGen, x.BridgeAnsatz, x.Hermitian, x.ZeroDiagonal, x.TriangleSupportInherited, x.EndpointBalanced, x.VertexRephasingAllowed, x.GaugeInvariantCyclePhase, x.EmpiricalDataImported, x.Verdict)
}

func FormatBoundary(x Boundary) string {
	return fmt.Sprintf("%s formula=%q applied=%t passed=%t verdict=%s reason=%s", x.Name, x.Formula, x.Applied, x.Passed, x.Verdict, x.Reason)
}

func FormatRealSignCandidate(x RealSignCandidate) string {
	return fmt.Sprintf("weights=%s product=%d class=%q Phi=%s det=%q CPpreserving=%t etaNeutral=%t JGamma=%t massLift=%t representative=%t", FormatSign(x.A, x.B, x.C), x.Product, x.GaugeClass, x.CyclePhase, x.DeterminantLeading, x.CPPreserving, x.EtaTraceNeutral, x.JGammaCompatible, x.MassLiftCompatible, x.Representative)
}

func FormatRealSignSieve(x RealSignSieve) string {
	return fmt.Sprintf("candidates=%d positive=%d negative=%d Z2classes=%d uniqueSigned=%t verdict=%s reason=%s", len(x.Candidates), x.PositiveCycleCount, x.NegativeCycleCount, x.Z2GaugeClasses, x.UniqueSignedCycle, x.Verdict, x.Reason)
}

func FormatPhaseSample(x PhaseSample) string {
	return fmt.Sprintf("Phi=%s radians=%.6g det=%q CPwitness=%s massLift=%t CPCapable=%t CPconjugate=%s", x.Label, x.PhiRadians, x.DeterminantLeading, x.CPWitness, x.MassLiftCompatible, x.CPCapable, x.CPConjugateLabel)
}

func FormatComplexPhaseSieve(x ComplexPhaseSieve) string {
	return fmt.Sprintf("invariant=%q determinant=%q CPmap=%q massCondition=%q witness=%q continuum=%t CPpairs=%t uniquePhase=%t CPValuePredicted=%t verdict=%s samples=[%s]", x.CyclePhaseInvariant, x.EndpointBalancedDeterminant, x.CPMap, x.MassLiftCondition, x.CPWitnessFormula, x.ContinuumSurvives, x.CPConjugatePairsSurvive, x.UniqueComplexPhase, x.CPPhaseValuePredicted, x.Verdict, phaseSampleSummary(x.Samples))
}

func FormatConclusion(x OrientationConclusion) string {
	return fmt.Sprintf("XSupport=%t signedForced=%t complexForced=%t YNative=%t phaseCoeffFixed=%t CPViolationPredicted=%t massLiftCompatible=%t verdict=%s reason=%s", x.XSupportTopologyPreserved, x.SignedCycleForced, x.ComplexPhaseForced, x.YGenPromotedToNative, x.PhaseCoefficientFixed, x.CPViolationPredicted, x.MassLiftStillCompatible, x.Verdict, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("muonImported=%t charmImported=%t yukawaImported=%t CKM=%t PMNS=%t XForced=%t amplitudeSealed=%t signSealed=%t phaseSealed=%t YQuarantined=%t nativeDim=%d→%d KXYFree=%d verdict=%s", !x.NoObservedMuonMassImported, !x.NoObservedCharmMassImported, !x.NoObservedYukawaImported, !x.NoCKMImported, !x.NoPMNSImported, x.XSupportTopologyForced, x.BridgeAmplitudeSealed, x.SignedCycleOrientationSealed, x.ComplexPhaseSealed, x.YGenRemainsQuarantined, x.NativeFlavorDimBefore, x.NativeFlavorDimAfter, x.KXYCoeffDimStillFree, x.Verdict)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s", x.Gate, x.Title, x.Reason)
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 446 Registry Audit — Signed-Cycle / Complex Phase Orientation Sieve\n\n")
	b.WriteString("## Claim tested\n\n")
	b.WriteString("Gate 446 tests whether the Gate-445 triangular bridge support collapses further to one signed real cycle or one complex CP phase. It deliberately does not import muon/charm masses, Yukawa matrices, CKM data, or PMNS data.\n\n")

	b.WriteString("## Prior boundary inherited\n\n")
	b.WriteString(FormatInheritance(a.Inheritance) + "\n\n")

	b.WriteString("## Orientation arena\n\n")
	b.WriteString(FormatArena(a.Arena) + "\n\n")
	b.WriteString("```text\n")
	b.WriteString("K_gen = diag(-1,0,1)\n")
	b.WriteString("B = [[0,z12,z13],[conj(z12),0,z23],[conj(z13),conj(z23),0]]\n")
	b.WriteString("Phi = arg(z12 z23 conjugate(z13))\n")
	b.WriteString("det(K_gen + eps B) = (|z23|^2-|z12|^2) eps^2 + 2 Re(z12 z23 conjugate(z13)) eps^3\n")
	b.WriteString("```\n\n")

	b.WriteString("## Boundary stack\n\n")
	b.WriteString("| Boundary | Formula | Applied | Passed | Verdict | Reason |\n")
	b.WriteString("|---|---|---:|---:|---|---|\n")
	for _, x := range a.Boundaries {
		b.WriteString(fmt.Sprintf("| %s | `%s` | %t | %t | `%s` | %s |\n", x.Name, x.Formula, x.Applied, x.Passed, x.Verdict, x.Reason))
	}
	b.WriteString("\n")

	b.WriteString("## Real signed-cycle sieve\n\n")
	b.WriteString(FormatRealSignSieve(a.RealSieve) + "\n\n")
	b.WriteString("| Weights | Product | Gauge class | Cycle phase | Determinant leading term | Representative |\n")
	b.WriteString("|---|---:|---|---|---|---:|\n")
	for _, x := range a.RealSieve.Candidates {
		b.WriteString(fmt.Sprintf("| `%s` | %d | %s | `%s` | `%s` | %t |\n", FormatSign(x.A, x.B, x.C), x.Product, x.GaugeClass, x.CyclePhase, x.DeterminantLeading, x.Representative))
	}
	b.WriteString("\n")
	b.WriteString("After quotienting by vertex sign flips, the eight sign assignments do not collapse to one class. They collapse only to two invariant cycle-product classes: `abc=+1` and `abc=-1`. Both are trace neutral, J/Gamma compatible, and mass-lift compatible.\n\n")

	b.WriteString("## Complex phase sieve\n\n")
	b.WriteString(FormatComplexPhaseSieve(a.PhaseSieve) + "\n\n")
	b.WriteString("Endpoint balance gives the exact reduction:\n\n")
	b.WriteString("```text\n")
	b.WriteString("det(K_gen + eps B) = 2 r^3 cos(Phi) eps^3\n")
	b.WriteString("CP-odd cycle witness ∝ sin(Phi)\n")
	b.WriteString("CP maps Phi -> -Phi\n")
	b.WriteString("```\n\n")
	b.WriteString("| Phase sample | det leading | CP witness | Mass lift | CP capable | CP conjugate |\n")
	b.WriteString("|---|---|---|---:|---:|---|\n")
	for _, x := range a.PhaseSieve.Samples {
		b.WriteString(fmt.Sprintf("| `%s` | `%s` | `%s` | %t | %t | `%s` |\n", x.Label, x.DeterminantLeading, x.CPWitness, x.MassLiftCompatible, x.CPCapable, x.CPConjugateLabel))
	}
	b.WriteString("\n")
	b.WriteString("The mass-lift boundary removes the purely imaginary cycle products with `cos(Phi)=0`, but it does not select a unique value of `Phi`. CP-capable pairs such as `Phi=±pi/4` survive as conjugate orientations.\n\n")

	b.WriteString("## Orientation conclusion\n\n")
	b.WriteString(FormatConclusion(a.Conclusion) + "\n\n")
	b.WriteString("Gate 446 therefore does not promote `Y_gen` or any CP phase value to native law-space. It only confirms that Gate-445 `X_gen` support is compatible with both CP-even and CP-odd phase orientations.\n\n")

	b.WriteString("## Phenomenology/firewall audit\n\n")
	b.WriteString(FormatFirewall(a.Firewall) + "\n\n")
	b.WriteString("The firewall remains intact: no observed masses, Yukawa coefficients, CKM angles, CKM CP phase, or PMNS data were used. The native charged flavor dimension remains 13, and the conditional K/X/Y coefficient ledger remains 9.\n\n")

	b.WriteString("## Result statuses\n\n")
	for _, s := range []string{StatusGate445TopologyInherited, StatusHermitianCycleArenaFormalized, StatusJGammaTraceBoundariesApplied, StatusRealSignedSieveCompleted, StatusCyclePhaseInvariantIdentified, StatusCPPhaseCapacityAudited, StatusEmpiricalFirewallPreserved, StatusFailedSignedOrientationNotUnique, StatusFailedComplexPhaseContinuum, StatusFailedCPPhaseValueNotPredicted, StatusFailedYGenNotNative, StatusFailedNoMuonCharmMassPrediction, StatusFirewallPhaseOrientationQuarantine} {
		b.WriteString("- `" + s + "`\n")
	}

	b.WriteString("\n## Next gate\n\n")
	b.WriteString(FormatNext(a.Next) + "\n\n")
	b.WriteString("## Truth statement\n\n")
	b.WriteString(a.Truth + "\n")
	return b.String()
}
