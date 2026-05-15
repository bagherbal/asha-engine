package generation2specialbranchselector

import (
	"fmt"
	"strings"
)

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("G444K=%t G444Zero=%t G445Triangle=%t G446PhaseSealed=%t G447CoeffSealed=%t G450SumRule=%t G450RatioSealed=%t noEmpirical=%t verdict=%s", x.Gate444KGenForced, x.Gate444Generation2Zero, x.Gate445TriangleForced, x.Gate446PhaseQuarantined, x.Gate447CoefficientsSealed, x.Gate450TextureZeroSumRule, x.Gate450RatioSealed, x.NoEmpiricalInputsImported, x.Verdict)
}

func FormatEdge(x Edge) string {
	return fmt.Sprintf("edge=%s %d-%d DeltaK=%d is13=%t KMSInteger=%t allowed=%t reason=%s", x.Name, x.From, x.To, x.DeltaK, x.Is13, x.KMSInteger, x.Allowed, x.Reason)
}

func FormatNativeLaw(x NativeLawAudit) string {
	return fmt.Sprintf("law=%s layer=%s suppress13=%t fixesPhase=%t allows=[12:%t,23:%t,13:%t] edgeBlind=%t phaseBlind=%t triangle=%t reason=%s", x.Law, x.NativeLayer, x.Suppresses13Edge, x.FixesPhaseRay, x.Allows12, x.Allows23, x.Allows13, x.EdgeBlind, x.PhaseBlind, x.CompatibleWithTriangle, x.Reason)
}

func FormatEdgeAudit(x EdgeSuppressionAudit) string {
	return fmt.Sprintf("edges=%d laws=%d X=%s NN=%s detFull=%s coeffFull=%d detNN=%s coeffNN=%d allAllow13=%t anySuppress13=%t trianglePreserved=%t NNForced=%t NNFailsLift=%t verdict=%s reason=%s", len(x.Edges), len(x.Laws), x.XTriangleFormula, x.NearestNeighborFormula, x.FullTriangleDeterminant, x.FullTriangleDeterminantCoeff, x.NearestNeighborDeterminant, x.NearestNeighborDeterminantCoeff, x.AllNativeLawsAllow13, x.AnyNativeLawSuppresses13, x.FullTrianglePreserved, x.NearestNeighborNativelyForced, x.NearestNeighborFailsMassLift, x.Verdict, x.Reason)
}

func FormatPhaseCandidate(x PhaseCandidate) string {
	return fmt.Sprintf("%s phi=%.6g b=%.6g c=%.6g cZero=%t detShape=%.6g nonzeroLift=%t hermitian=%t trace0=%t M22Zero=%t KMS=%t anomaly=%t firstOrder=%t empirical=%t survives=%t", x.Label, x.Phi, x.B, x.C, x.CZero, x.DeterminantShape, x.NonzeroMassLift, x.Hermitian, x.TraceZero, x.StructuralZero22, x.KMSCompatible, x.AnomalyCompatible, x.FirstOrderCompatible, x.ImportsEmpiricalData, x.SurvivesNativeConstraints)
}

func FormatPhaseAudit(x PhaseRayAudit) string {
	return fmt.Sprintf("candidates=%d phaseBlind=%t survivors=%d cZeroSurvivor=%t nonzeroCSurvivor=%t pureYDegenerate=%t uniqueRay=%t fixesCZero=%t fixesPiOverTwo=%t verdict=%s reason=%s", len(x.Candidates), x.NativeConstraintsPhaseBlind, x.SurvivingNonzeroLiftRays, x.ContainsCZeroSurvivor, x.ContainsNonzeroCSurvivor, x.PureYRayLiftDegenerate, x.UniqueRayForced, x.FixesCZero, x.FixesPiOverTwo, x.Verdict, x.Reason)
}

func FormatGST(x GSTBranchVerdict) string {
	return fmt.Sprintf("edgeSelector=%t phaseSelector=%t nativeGST=%t reevaluated=%t empiricalAssumption=%t assumptions=%s verdict=%s reason=%s", x.EdgeSelectorFound, x.PhaseSelectorFound, x.GSTLikeBranchNativelyForced, x.MassAngleRatiosReevaluated, x.GSTFritzschEmpiricalAssumption, strings.Join(x.NecessaryNonNativeAssumptions, "; "), x.Verdict, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("noMuon=%t noCharm=%t noYukawa=%t noCKM=%t noPMNS=%t noFit=%t K=%t Gen2Zero=%t XTriangle=%t YPhaseSealed=%t coeffSealed=%t GSTSealed=%t nativeDim=%d KXY=%d verdict=%s reason=%s", x.NoObservedMuonMassImported, x.NoObservedCharmMassImported, x.NoObservedYukawaImported, x.NoCKMImported, x.NoPMNSImported, x.NoCurveFit, x.KGenStillForced, x.Generation2ZeroStillForced, x.XTriangleStillForced, x.YPhaseStillQuarantined, x.SectorCoefficientsStillSealed, x.GSTFritzschRelationsQuarantined, x.NativeFlavorDimAfter, x.KXYCoeffDimAfter, x.Verdict, x.Reason)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s Task=%s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 451 Registry Audit — Texture-Zero Special-Branch Selector / Necessary Boundary Audit\n\n")
	b.WriteString("## Scope\n\n")
	b.WriteString("Gate 451 audits whether the native ASHA law-space secretly contains the extra selector needed to turn the Gate-450 texture-zero identity into a GST/Fritzsch branch. Two selectors are tested: native suppression of the 1-3 edge and native fixation of the complex phase ray. No observed lepton, quark, CKM, PMNS, or Yukawa data is imported.\n\n")

	b.WriteString("## Inheritance\n\n")
	b.WriteString(FormatInheritance(a.Inheritance) + "\n\n")

	b.WriteString("## Edge suppression audit\n\n")
	b.WriteString(FormatEdgeAudit(a.EdgeAudit) + "\n\n")
	b.WriteString("| Edge | Delta K | KMS integer? | Allowed? | Reason |\n")
	b.WriteString("|---|---:|---|---|---|\n")
	for _, e := range a.EdgeAudit.Edges {
		b.WriteString(fmt.Sprintf("| %s | %d | %t | %t | %s |\n", e.Name, e.DeltaK, e.KMSInteger, e.Allowed, esc(e.Reason)))
	}
	b.WriteString("\n")
	b.WriteString("| Native boundary | Layer | Allows 1-3? | Suppresses 1-3? | Fixes phase? | Triangle compatible? | Reason |\n")
	b.WriteString("|---|---|---|---|---|---|---|\n")
	for _, law := range a.EdgeAudit.Laws {
		b.WriteString(fmt.Sprintf("| %s | %s | %t | %t | %t | %t | %s |\n", esc(law.Law), esc(law.NativeLayer), law.Allows13, law.Suppresses13Edge, law.FixesPhaseRay, law.CompatibleWithTriangle, esc(law.Reason)))
	}
	b.WriteString("\n")
	b.WriteString("```text\n")
	b.WriteString(a.EdgeAudit.FullTriangleDeterminant + "\n")
	b.WriteString(a.EdgeAudit.NearestNeighborDeterminant + "\n")
	b.WriteString("```\n\n")
	b.WriteString("The 1-3 edge is not a forbidden edge under KMS quantization; it is the integer second harmonic between the `-1` and `+1` levels. Removing it creates the nearest-neighbor chain, but that chain is not natively selected and fails the primitive mass-lift determinant test.\n\n")

	b.WriteString("## Phase ray audit\n\n")
	b.WriteString(FormatPhaseAudit(a.PhaseAudit) + "\n\n")
	b.WriteString("| Candidate ray | phi | b | c | det shape cos(3phi) | Nonzero lift? | Survives native constraints? |\n")
	b.WriteString("|---|---:|---:|---:|---:|---|---|\n")
	for _, c := range a.PhaseAudit.Candidates {
		b.WriteString(fmt.Sprintf("| %s | %.6g | %.6g | %.6g | %.6g | %t | %t |\n", esc(c.Label), c.Phi, c.B, c.C, c.DeterminantShape, c.NonzeroMassLift, c.SurvivesNativeConstraints))
	}
	b.WriteString("\n")
	b.WriteString("At least one `c=0` ray and multiple `c!=0` rays survive the native constraints with nonzero determinant. The pure `Y` ray is included as a diagnostic: it is native-constraint compatible but lift-degenerate, so it also cannot be the forced GST selector.\n\n")

	b.WriteString("## GST/Fritzsch branch verdict\n\n")
	b.WriteString(FormatGST(a.GST) + "\n\n")
	b.WriteString("Non-native assumptions required before a GST/Fritzsch branch can be studied:\n\n")
	for _, x := range a.GST.NecessaryNonNativeAssumptions {
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
