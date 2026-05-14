package contactquarticscalaryukawabundle

import (
	"fmt"
	"math"
	"strings"
)

func FormatInheritance(a Inheritance) string {
	return fmt.Sprintf("executed=%t singletonFlavorBlocked=%t qdim=%d hphiDim=%d galois=%t abstractModule=%t companion=%t hphiID=%t physicalScalarBundle=%t scalar=%d/%d protected=%d normalForm=%t oneForm=%t edges=%d yukawaChannels=%d fibers=%d massMatrix=%t chargedModuli=%d noEmpirical=%t (%s)", a.Executed, a.Gate397SingletonFlavorBlocked, a.Gate183QuarticPrimaryDim, a.Gate183ScalarCarrierDim, a.Gate183GaloisSafePrimaryIdeal, a.Gate183AbstractRankOneModule, a.Gate183CompanionRepresentation, a.Gate183CanonicalHphiID, a.Gate183PhysicalScalarBundle, a.Gate37ActiveScalarDim, a.Gate37ComplexDoubletDim, a.Gate37ProtectedDirections, a.Gate37ScalarNormalFormAvailable, a.Gate385OneFormEdgeSupportDerived, a.Gate385JDoubledEdgeCount, a.Gate26MinimalYukawaChannels, a.Gate26ScalarFiberEntries, a.Gate26MassMatrixDerived, a.Gate372ChargedModuliDim, a.NoEmpiricalFlavorValuesImported, a.Verdict)
}

func FormatQuartic(a QuarticPrimaryAudit) string {
	return fmt.Sprintf("algebra=%q q=%q dim=%d field=%q galois=%t branchFree=%t selectedBranches=%d companion=%t abstractModule=%t contactExact=%t scalarExact=%t hphiID=%t scalarMinPoly=%t semantics=%q (%s)", a.Algebra, a.Polynomial, a.Dimension, a.BaseField, a.GaloisSafePrimary, a.BranchFreeBlock, a.IndividualBranchesSelected, a.CompanionRepresentation, a.AbstractRankOneModule, a.ExactAsContactSpectralDatum, a.ExactAsScalarHiggsDatum, a.CanonicalHphiIdentification, a.ScalarMinimalPolynomialDerived, a.SpectrumSemantics, a.Verdict)
}

func FormatScalar(a ScalarCarrierAudit) string {
	res := fmt.Sprintf("%.12g", a.QuarticMinimalResidual)
	if math.IsInf(a.QuarticMinimalResidual, 1) {
		res = "not-defined"
	}
	return fmt.Sprintf("carrier=%q realDim=%d complexDim=%d protected=%d pairDeg=%t r0^2=%.12g lambda=%.12g normalForm=%t eating=%t ewScale=%t higgsMass=%t quarticAction=%t minPoly=%q residual=%s (%s)", a.Carrier, a.ActiveRealDim, a.ComplexDoubletDim, a.ProtectedDirections, a.PairDegenerate, a.VacuumRadiusSquared, a.LambdaShape, a.NormalFormAvailable, a.GaugeEatingTheoremDerived, a.ElectroweakScaleDerived, a.HiggsMassDerived, a.CanonicalQuarticAction, a.QuarticMinimalPolynomial, res, a.Verdict)
}

func FormatTarget(a BundleTargetAudit) string {
	return fmt.Sprintf("oneForm=%t edges=%d edgeMeasure=%t yukawaChannels=%d fibers=%d branches=%d massMatrix=%t couplings=%t quarticEdges=%t quarticFibers=%t yukawaReduced=%t (%s)", a.OneFormEdgeSupportDerived, a.JDoubledEdgeCount, a.OneFormMeasureSelected, a.YukawaChannels, a.ScalarFiberEntries, a.ScalarBranches, a.MassMatrixDerived, a.CouplingConstantsDerived, a.QuarticActsOnEdges, a.QuarticWeightsYukawaFibers, a.YukawaBundleReduced, a.Verdict)
}

func FormatFunctor(c FunctorCandidate) string {
	res := fmt.Sprintf("%.12g", c.Residual)
	if math.IsInf(c.Residual, 1) {
		res = "not-defined"
	}
	return fmt.Sprintf("%s domain=%q target=%q native=%t sealed=%t circular=%t dim=%t branchFree=%t hom=%t module=%t physical=%t AF=%t J=%t firstOrder=%t EW=%t oneForm=%t minPoly=%t yukawa=%t moduli=%t arbitraryID=%t rank=%d spectrum=%v residual=%s promotable=%t reason=%q (%s)", c.Name, c.Domain, c.Target, c.Native, c.Sealed, c.Circular, c.DimensionCompatible, c.BranchFree, c.AlgebraHomomorphism, c.ProjectiveModule, c.PhysicalCarrierAction, c.CompatibleWithAF, c.CompatibleWithJ, c.CompatibleWithFirstOrder, c.CompatibleWithElectroweak, c.CompatibleWithOneFormEdges, c.ScalarMinimalPolynomial, c.ReducesYukawaCouplings, c.ReducesFlavorModuli, c.ArbitraryBasisIdentification, c.Rank, c.Spectrum, res, c.PromotableAsNativeFunctor, c.Reason, c.Verdict)
}

func FormatFunctors(a FunctorAudit) string {
	rows := []string{fmt.Sprintf("executed=%t dimMatches=%d native=%d sealed=%d abstractModules=%d physicalScalar=%d oneFormActions=%d yukawaReducers=%d promotable=%d best=%q (%s)", a.Executed, a.DimensionCompatibleCount, a.NativeCandidateCount, a.SealedCandidateCount, a.AbstractModuleCount, a.PhysicalScalarActions, a.OneFormEdgeActions, a.YukawaReducingActions, a.PromotableNativeCount, a.BestNativeCandidate, a.Verdict)}
	for _, c := range a.Candidates {
		rows = append(rows, "  - "+FormatFunctor(c))
	}
	return strings.Join(rows, "\n")
}

func FormatScenario(s ImpactScenario) string {
	return fmt.Sprintf("%s assumption=%q native=%t conditional=%t failed=%t scalarBundle=%t yukawaReduced=%t moduli=%d->%d higgsChanged=%t firewall=%t reason=%q (%s)", s.Name, s.AssumptionClass, s.Native, s.Conditional, s.Failed, s.ScalarBundleDerived, s.YukawaCouplingsReduced, s.ChargedModuliStart, s.ChargedModuliResult, s.HiggsLaneChanged, s.FlavorFirewallPreserved, s.Reason, s.Verdict)
}

func FormatImpact(a ImpactAudit) string {
	rows := []string{fmt.Sprintf("executed=%t start=%d nativeFlavorReduction=%t bestNative=%d bestConditional=%d scalarLane=%t physicalHiggs=%t (%s)", a.Executed, a.ChargedModuliStart, a.NativeFlavorReduction, a.BestNativeModuliDim, a.BestConditionalModuliDim, a.ScalarHiggsLanePreserved, a.PhysicalHiggsMassDerived, a.Verdict)}
	for _, s := range a.Scenarios {
		rows = append(rows, "  - "+FormatScenario(s))
	}
	return strings.Join(rows, "\n")
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("executed=%t masses=%t ckm=%t pmns=%t ordering=%t observedHiggs=%t manualHphiID=%t companion=%t arbitraryBasis=%t yukawaClaim=%t moduliClaim=%t (%s)", a.Executed, a.NoMassesImported, a.NoCKMImported, a.NoPMNSImported, a.NoEmpiricalOrderingImported, a.NoObservedHiggsUsedForFunctor, a.NoManualQuarticHphiID, a.NoCompanionOperatorPromoted, a.NoArbitraryBasisMapPromoted, a.NoYukawaCouplingClaimed, a.NoFlavorModuliReductionClaimed, a.Verdict)
}

func FormatNext(a NextStep) string {
	return fmt.Sprintf("Gate %d — %s\nReason: %s\nPrimary task: %s", a.Gate, a.Title, a.Reason, a.PrimaryTask)
}

func RenderMarkdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 398 Registry Audit — Contact Quartic Primary to Scalar/Yukawa Bundle Functor Audit\n\n")
	b.WriteString("## Claim tested\n\n")
	b.WriteString("Can the exact four-dimensional quartic contact primary block be canonically promoted to the scalar/Higgs carrier and then to the finite one-form/Yukawa bundle? Equivalently, does the project derive `rho_4: Q[x]/(q4) -> End(H_phi)` compatible with `A_F`, `J`, first-order, electroweak charges, and the one-form edge module?\n\n")
	b.WriteString("## Previous gates used\n\n```text\n")
	b.WriteString(FormatInheritance(a.Inheritance))
	b.WriteString("\n```\n\n")
	b.WriteString("## Quartic primary block\n\n```text\n")
	b.WriteString(FormatQuartic(a.Quartic))
	b.WriteString("\n```\n\n")
	b.WriteString("## Scalar/Higgs carrier\n\n```text\n")
	b.WriteString(FormatScalar(a.Scalar))
	b.WriteString("\n```\n\n")
	b.WriteString("## One-form/Yukawa target\n\n```text\n")
	b.WriteString(FormatTarget(a.Target))
	b.WriteString("\n```\n\n")
	b.WriteString("## Candidate functor table\n\n")
	b.WriteString("| Candidate | Native | Sealed | Dimension | Hom | Physical action | One-form compatible | Yukawa reduction | Promotable | Verdict |\n")
	b.WriteString("|---|---:|---:|---:|---:|---:|---:|---:|---:|---|\n")
	for _, c := range a.Functors.Candidates {
		b.WriteString(fmt.Sprintf("| %s | %t | %t | %t | %t | %t | %t | %t | %t | `%s` |\n", c.Name, c.Native, c.Sealed, c.DimensionCompatible, c.AlgebraHomomorphism, c.PhysicalCarrierAction, c.CompatibleWithOneFormEdges, c.ReducesYukawaCouplings, c.PromotableAsNativeFunctor, c.Verdict))
	}
	b.WriteString("\n```text\n")
	b.WriteString(FormatFunctors(a.Functors))
	b.WriteString("\n```\n\n")
	b.WriteString("## Impact audit\n\n")
	b.WriteString("| Scenario | Class | Scalar bundle | Yukawa reduced | Moduli result | Native | Conditional | Verdict |\n")
	b.WriteString("|---|---|---:|---:|---:|---:|---:|---|\n")
	for _, s := range a.Impact.Scenarios {
		b.WriteString(fmt.Sprintf("| %s | %s | %t | %t | %d | %t | %t | `%s` |\n", s.Name, s.AssumptionClass, s.ScalarBundleDerived, s.YukawaCouplingsReduced, s.ChargedModuliResult, s.Native, s.Conditional, s.Verdict))
	}
	b.WriteString("\n```text\n")
	b.WriteString(FormatImpact(a.Impact))
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
