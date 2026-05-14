package quaternionicscalarbundleidentity

import (
	"fmt"
	"math"
	"strings"
)

func FormatInheritance(a Inheritance) string {
	return fmt.Sprintf("executed=%t gate398NoHphiID=%t qdim=%d hphiDim=%d q4=%q companionSealed=%t localH=%t Hclosure=%t globalH=%t physicalJ=%t weakHLeft=%t exactAF=%t firstOrder=%t pairComplex=%t canonicalComplex=%t qTriple=%t qSelected=%t oneForm=%t edges=%d chargedModuli=%d noEmpirical=%t (%s)", a.Executed, a.Gate398NoCanonicalHphiID, a.Gate398QuarticDim, a.Gate398HphiDim, a.Gate398Q4Polynomial, a.Gate398CompanionStressSealed, a.Gate274LocalHExtracted, a.Gate274QuaternionClosureExact, a.Gate274GlobalHDerived, a.Gate274PhysicalJDerived, a.Gate295WeakHLeftActionIsolated, a.Gate295ExactAFDerived, a.Gate295FirstOrderComplete, a.Gate50PairComplexAvailable, a.Gate50CanonicalComplexDerived, a.Gate50QuaternionicTripleAvailable, a.Gate50QuaternionicTripleSelected, a.Gate385OneFormEdgeSupportDerived, a.Gate385JDoubledEdgeCount, a.Gate372ChargedModuliDim, a.NoEmpiricalValuesImported, a.Verdict)
}

func FormatQ4(a Q4Fingerprint) string {
	return fmt.Sprintf("q4=%q coeffs=%s degree=%d irreducibleQ=%t branchFree=%t contactDatum=%t (%s)", a.Polynomial, formatFloatSlice(a.MonicCoefficients), a.Degree, a.IrreducibleOverQ, a.BranchFreePrimary, a.ContactSpectralDatum, a.Verdict)
}

func FormatModule(a QuaternionicModuleAudit) string {
	return fmt.Sprintf("carrier=%q realDim=%d complexDoublet=%d algebra=%q localH=%t moritaWeakH=%t globalH=%t pairComplex=%t canonicalComplex=%t qTriple=%t qSelectedByScalar=%t fullSU2=%t AF=%t J=%t firstOrder=%t oneForm=%t (%s)", a.Carrier, a.RealDimension, a.ComplexDoubletDimension, a.Algebra, a.LocalHExtracted, a.MoritaWeakHAction, a.GlobalHUnsealed, a.PairComplexAvailable, a.CanonicalComplexDerived, a.AbstractQuaternionicTripleAvailable, a.QuaternionicTripleSelectedByScalar, a.FullScalarSU2Recovered, a.CompatibleWithAF, a.CompatibleWithJ, a.CompatibleWithFirstOrder, a.CompatibleWithOneFormEdges, a.Verdict)
}

func FormatEndomorphism(a EndomorphismFingerprint) string {
	sq := fmt.Sprintf("%.6e", a.SquareResidual)
	cl := fmt.Sprintf("%.6e", a.ClosureResidual)
	comm := fmt.Sprintf("%.6e", a.ScalarCommutatorNorm)
	q4r := fmt.Sprintf("%.6e", a.Q4CoefficientResidual)
	if math.IsInf(a.ScalarCommutatorNorm, 0) {
		comm = "structural/not-computed"
	}
	if math.IsInf(a.Q4CoefficientResidual, 0) {
		q4r = "structural/not-equal"
	}
	return fmt.Sprintf("%s source=%q dim=%d native=%t sealed=%t circular=%t Haction=%t square=-I:%t sqRes=%s closure=%s minPoly=%q deg=%d char=%q coeffs=%s squareQuad=%t q4Residual=%s q4Exact=%t q4Factor=%t scalarCommutes=%t scalarComm=%s AF=%t J=%t firstOrder=%t oneForm=%t promotable=%t reason=%q (%s)", a.Name, a.Source, a.MatrixDim, a.Native, a.Sealed, a.Circular, a.QuaternionicAction, a.SquaresToMinusIdentity, sq, cl, a.MinimalPolynomial, a.MinimalDegree, a.CharacteristicPolynomial, formatFloatSlice(a.CharacteristicCoefficients), a.CharPolyIsSquareOfQuadratic, q4r, a.Q4ExactMatch, a.Q4FactorMatch, a.CommutesWithScalarResponse, comm, a.CompatibleWithAF, a.CompatibleWithJ, a.CompatibleWithFirstOrder, a.CompatibleWithOneFormEdges, a.PromotableAsQ4Selector, a.Reason, a.Verdict)
}

func FormatEndomorphisms(a EndomorphismAudit) string {
	rows := []string{fmt.Sprintf("executed=%t quaternionic=%d q4Exact=%d q4Factor=%d promotable=%d maxScalarComm=%.6e best=%q (%s)", a.Executed, a.QuaternionicCandidateCount, a.Q4ExactMatchCount, a.Q4FactorMatchCount, a.PromotableNativeCount, a.MaxScalarCommutator, a.BestNativeCandidate, a.Verdict)}
	for _, c := range a.Candidates {
		rows = append(rows, "  - "+FormatEndomorphism(c))
	}
	return strings.Join(rows, "\n")
}

func FormatIdentity(a BundleIdentityAudit) string {
	return fmt.Sprintf("executed=%t hphiQuarticID=%t scalarSealed=%t oneFormFunctor=%t yukawaReduced=%t moduli=%d->%d flavorFirewall=%t higgsLanePreserved=%t (%s)", a.Executed, a.HphiQuarticIdentified, a.ScalarBundleGeometricallySealed, a.OneFormEdgeFunctorDerived, a.YukawaCouplingsReduced, a.ChargedModuliStart, a.ChargedModuliResult, a.FlavorFirewallPreserved, a.HiggsLanePreserved, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("executed=%t masses=%t CKM=%t PMNS=%t observedHiggs=%t manualQ4ID=%t companion=%t arbitraryBasis=%t yukawaClaim=%t moduliClaim=%t (%s)", a.Executed, a.NoObservedMassesImported, a.NoCKMImported, a.NoPMNSImported, a.NoObservedHiggsInserted, a.NoManualQ4HphiID, a.NoCompanionOperatorPromoted, a.NoArbitraryBasisMapPromoted, a.NoYukawaCouplingClaimed, a.NoFlavorModuliReductionClaimed, a.Verdict)
}

func FormatNext(a NextStep) string {
	return fmt.Sprintf("Gate %d — %s\nReason: %s\nPrimary task: %s", a.Gate, a.Title, a.Reason, a.PrimaryTask)
}

func RenderMarkdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 399 Registry Audit — Quaternionic (H) Endomorphism / Scalar Bundle Identity Sieve\n\n")
	b.WriteString("## Claim tested\n\n")
	b.WriteString("Can the weak quaternionic `H` action on the four-real-dimensional scalar carrier `H_phi` provide the missing basis-free identity selector from Gate 398 by producing the contact quartic primary polynomial `q4` as a native minimal or characteristic polynomial?\n\n")
	b.WriteString("## Previous gates used\n\n```text\n")
	b.WriteString(FormatInheritance(a.Inheritance))
	b.WriteString("\n```\n\n")
	b.WriteString("## Contact quartic fingerprint\n\n```text\n")
	b.WriteString(FormatQ4(a.Q4))
	b.WriteString("\n```\n\n")
	b.WriteString("## Quaternionic scalar module audit\n\n```text\n")
	b.WriteString(FormatModule(a.Module))
	b.WriteString("\n```\n\n")
	b.WriteString("## Endomorphism fingerprint table\n\n")
	b.WriteString("| Candidate | Native | Sealed | H-action | Minimal degree | Characteristic | q4 exact | Promotable | Verdict |\n")
	b.WriteString("|---|---:|---:|---:|---:|---|---:|---:|---|\n")
	for _, c := range a.Endomorphisms.Candidates {
		b.WriteString(fmt.Sprintf("| %s | %t | %t | %t | %d | %s | %t | %t | `%s` |\n", c.Name, c.Native, c.Sealed, c.QuaternionicAction, c.MinimalDegree, c.CharacteristicPolynomial, c.Q4ExactMatch, c.PromotableAsQ4Selector, c.Verdict))
	}
	b.WriteString("\n```text\n")
	b.WriteString(FormatEndomorphisms(a.Endomorphisms))
	b.WriteString("\n```\n\n")
	b.WriteString("## Identity / impact audit\n\n```text\n")
	b.WriteString(FormatIdentity(a.Identity))
	b.WriteString("\n```\n\n")
	b.WriteString("## Firewall status\n\n```text\n")
	b.WriteString(FormatFirewall(a.Firewall))
	b.WriteString("\n```\n\n")
	b.WriteString("## Statuses\n\n```text\n")
	b.WriteString(strings.Join(Statuses(a), "\n"))
	b.WriteString("\n```\n\n")
	b.WriteString("## Conclusion\n\n")
	b.WriteString(a.Truth)
	b.WriteString("\n\n## Next gate\n\n```text\n")
	b.WriteString(FormatNext(a.Next))
	b.WriteString("\n```\n")
	return b.String()
}

func formatFloatSlice(xs []float64) string {
	if len(xs) == 0 {
		return "[]"
	}
	parts := make([]string, len(xs))
	for i, x := range xs {
		parts[i] = fmt.Sprintf("%.12g", x)
	}
	return "[" + strings.Join(parts, ", ") + "]"
}
