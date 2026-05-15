package generation2massmixinginvariants

import (
	"fmt"
	"strings"
)

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("G444K=%t G444Zero=%t G445Triangle=%t G446PhaseSealed=%t G447CoeffSealed=%t G449Board=%t nativeDim=%d KXY=%d noEmpirical=%t verdict=%s", x.Gate444KGenForced, x.Gate444Generation2Zero, x.Gate445TriangleForced, x.Gate446PhaseQuarantined, x.Gate447CoefficientsSealed, x.Gate449BoardExported, x.NativeFlavorDim, x.KXYCoeffDim, x.NoEmpiricalInputsImported, x.Verdict)
}

func FormatArena(x MatrixArena) string {
	return fmt.Sprintf("%s; %s; %s; %s hermitian=%t trace0=%t M22Zero=%t triangle=%t endpointBalanced=%t symbolicOnly=%t noCollider=%t verdict=%s reason=%s", x.KGenFormula, x.XTriangleFormula, x.YCycleFormula, x.MassMatrixFormula, x.Hermitian, x.TraceZero, x.StructuralZero22, x.ClosedTriangleSupport, x.EndpointBalanced, x.UsesSymbolicABCOnly, x.NoColliderDataUsed, x.Verdict, x.Reason)
}

func FormatEigen(x SymbolicEigenAnalysis) string {
	return fmt.Sprintf("chi=%s P=%s D=%s cardano=%s eigenvector=%s trace=%s det=%s convention=%s freeScaleRatio=%t freePhase=%t coeffRequired=%t verdicts=[%s,%s] reason=%s", x.CharacteristicPolynomial, x.PInvariant, x.DInvariant, x.CardanoEigenvalueFormula, x.EigenvectorFormula, x.TraceIdentity, x.DeterminantIdentity, x.PhysicalMassConvention, x.ContainsFreeScaleRatio, x.ContainsFreeCyclePhase, x.CoefficientsStillRequired, x.VerdictPolynomial, x.VerdictEigenvectors, x.Reason)
}

func FormatIdentity(x TextureZeroIdentity) string {
	return fmt.Sprintf("zero=%s sumRule=%s localAngles=%s GST=%s exact=%t specificRatio=%t needsU=%t needsCoeff=%t verdict=%s reason=%s", x.StructuralZeroFormula, x.SpectralSumRule, x.LocalJacobiAngleFormula, x.GSTCandidate, x.SumRuleExact, x.SpecificMassAngleRatio, x.RequiresEigenvectorData, x.RequiresCoefficientData, x.Verdict, x.Reason)
}

func FormatCounterexample(x Counterexample) string {
	return fmt.Sprintf("%s: a=%.6g b=%.6g c=%.6g r=%.6g phi=%.6g P=%.6g D=%.6g q=%.6g theta=%.6g eig/sqrtP=%s absRatios=%s boundary=%t empirical=%t demonstrates=%s", x.Label, x.A, x.B, x.C, x.R, x.Phase, x.P, x.D, x.ShapeQ, x.LocalTheta, floatList(x.NormalizedEigenvalues), floatList(x.AbsMassRatios), x.BoundaryCompatible, x.ImportsEmpiricalData, x.Demonstrates)
}

func FormatSieve(x RatioSieve) string {
	return fmt.Sprintf("examples=%d sameAngleDifferentMass=%t sameMassDifferentAngle=%t zeroSumRule=%t uniqueInvariant=%t coeffRequired=%t phaseRequired=%t scaleIrrelevant=%t verdict=%s reason=%s", len(x.Counterexamples), x.SameAngleDifferentMassShape, x.SameMassShapeDifferentAngle, x.StructuralZeroSumRuleSurvives, x.UniqueMassAngleInvariant, x.CoefficientsRequiredForRatios, x.PhaseRequiredForRatios, x.AbsoluteScaleIrrelevantButRatios, x.Verdict, x.Reason)
}

func FormatGST(x GSTFritzschAudit) string {
	return fmt.Sprintf("historical=%s extras=%s fullTriangle=%t phaseFree=%t coeffRayFree=%t forcedGST=%t universalApprox=%t specialLater=%t verdict=%s reason=%s", x.HistoricalRelation, strings.Join(x.NecessaryExtraAssumptions, "; "), x.ASHATopologyHasFullTriangle, x.ASHAPhaseContinuumUnfixed, x.ASHACoefficientRayUnfixed, x.RecognizableGSTRelationForced, x.ApproximateGSTRelationUniversal, x.SpecialBranchesMayBeTestedLater, x.Verdict, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("noMuon=%t noCharm=%t noYukawa=%t noCKM=%t noPMNS=%t noFit=%t K=%t X=%t YSealed=%t coeffSealed=%t ratioSealed=%t nativeDim=%d KXY=%d verdict=%s reason=%s", x.NoObservedMuonMassImported, x.NoObservedCharmMassImported, x.NoObservedYukawaImported, x.NoCKMImported, x.NoPMNSImported, x.NoCurveFit, x.KGenStillForced, x.XTriangleStillForced, x.YPhaseStillQuarantined, x.CoefficientsStillSealed, x.RatioPredictionSealed, x.NativeFlavorDimAfter, x.KXYCoeffDimAfter, x.Verdict, x.Reason)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s Task=%s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 450 Registry Audit — Structural Zero Mass-Mixing Invariants / Ratio Sieve\n\n")
	b.WriteString("## Scope\n\n")
	b.WriteString("Gate 450 cancels the publication-only lane and tests the strongest post-444 flavor temptation: whether the forced Generation-2 structural zero plus the closed-triangle bridge topology already implies a GST/Fritzsch-style mass-angle ratio. It uses symbolic matrix algebra and counterexample sieves only; no observed lepton, quark, CKM, PMNS, or Yukawa data is imported.\n\n")

	b.WriteString("## Inheritance\n\n")
	b.WriteString(FormatInheritance(a.Inheritance) + "\n\n")

	b.WriteString("## Symbolic matrix arena\n\n")
	b.WriteString(FormatArena(a.Arena) + "\n\n")
	b.WriteString("```text\n")
	b.WriteString(a.Arena.MassMatrixFormula + "\n")
	b.WriteString("```\n\n")

	b.WriteString("## Symbolic eigen-analysis\n\n")
	b.WriteString(FormatEigen(a.Eigen) + "\n\n")
	b.WriteString("| Object | Formula | Status |\n")
	b.WriteString("|---|---|---|\n")
	b.WriteString(fmt.Sprintf("| Characteristic polynomial | `%s` | `%s` |\n", esc(a.Eigen.CharacteristicPolynomial), a.Eigen.VerdictPolynomial))
	b.WriteString(fmt.Sprintf("| Quadratic invariant | `%s` | symbolic |\n", esc(a.Eigen.PInvariant)))
	b.WriteString(fmt.Sprintf("| Cubic/determinant invariant | `%s` | symbolic |\n", esc(a.Eigen.DInvariant)))
	b.WriteString(fmt.Sprintf("| Eigenvalues | `%s` | Cardano form |\n", esc(a.Eigen.CardanoEigenvalueFormula)))
	b.WriteString(fmt.Sprintf("| Eigenvectors | `%s` | `%s` |\n", esc(a.Eigen.EigenvectorFormula), a.Eigen.VerdictEigenvectors))
	b.WriteString("\n")

	b.WriteString("## Texture-zero identity\n\n")
	b.WriteString(FormatIdentity(a.Identity) + "\n\n")
	b.WriteString("The structural zero does prove the exact spectral sum rule\n\n")
	b.WriteString("```text\n")
	b.WriteString(a.Identity.SpectralSumRule + "\n")
	b.WriteString("```\n\n")
	b.WriteString("but this is not the same thing as a pairwise relation such as `sin(theta_ij)=sqrt(m_i/m_j)`. The eigenvector weights remain part of the equation.\n\n")

	b.WriteString("## Ratio sieve counterexamples\n\n")
	b.WriteString(FormatSieve(a.Sieve) + "\n\n")
	b.WriteString("| Witness | a | b | c | phi | q=D/P^(3/2) | theta_local | normalized eigenvalues | abs mass ratios | Demonstrates |\n")
	b.WriteString("|---|---:|---:|---:|---:|---:|---:|---|---|---|\n")
	for _, x := range a.Sieve.Counterexamples {
		b.WriteString(fmt.Sprintf("| %s | %.6g | %.6g | %.6g | %.6g | %.6g | %.6g | `%s` | `%s` | %s |\n", esc(x.Label), x.A, x.B, x.C, x.Phase, x.ShapeQ, x.LocalTheta, floatList(x.NormalizedEigenvalues), floatList(x.AbsMassRatios), esc(x.Demonstrates)))
	}
	b.WriteString("\n")
	b.WriteString("Two independent obstructions are visible: one pair keeps the same local mixing angle while changing the mass-shape invariant, and another pair keeps the same normalized mass spectrum while changing the local mixing angle. This kills a universal topology-only mass-angle ratio.\n\n")

	b.WriteString("## GST / Fritzsch test\n\n")
	b.WriteString(FormatGST(a.GST) + "\n\n")
	b.WriteString("Extra assumptions that would be needed before a GST-like branch could be audited:\n\n")
	for _, x := range a.GST.NecessaryExtraAssumptions {
		b.WriteString("- " + x + "\n")
	}
	b.WriteString("\n")

	b.WriteString("## Result statuses\n\n")
	for _, s := range statuses() {
		b.WriteString("- `" + s + "`\n")
	}
	b.WriteString("\n")

	b.WriteString("## Firewall\n\n")
	b.WriteString(FormatFirewall(a.Firewall) + "\n\n")

	b.WriteString("## Next gate\n\n")
	b.WriteString(FormatNext(a.Next) + "\n\n")

	b.WriteString("## Truth statement\n\n")
	b.WriteString(a.Truth + "\n")
	return b.String()
}

func esc(s string) string {
	s = strings.ReplaceAll(s, "|", "\\|")
	s = strings.ReplaceAll(s, "\n", "<br>")
	return s
}
