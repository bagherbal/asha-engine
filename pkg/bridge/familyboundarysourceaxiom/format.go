package familyboundarysourceaxiom

import (
	"fmt"
	"strings"
)

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%v gate413_capacity=%v pair_not_native=%v gate414_no_selector=%v trace_curvature_exhausted=%v coefficients_free=%v charged_moduli=%d verdict=%s", x.Executed, x.Gate413NoncommutingCapacity, x.Gate413PairNotNative, x.Gate414NoCoefficientSelector, x.Gate414TraceCurvatureExhausted, x.Gate414CoefficientsFree, x.ChargedModuliDim, x.Verdict)
}
func FormatArena(x SourceArena) string {
	return fmt.Sprintf("executed=%v family_pair=%q sectors=%s coeff_per_sector=%d baseline_rays=%d gauge_compatible=%v native_selector=%v empirical_yukawa=%v verdict=%s reason=%s", x.Executed, x.FamilyPair, strings.Join(x.Sectors, ","), x.CoefficientsPerSector, x.BaselineCoefficientRays, x.GaugeCompatibleIfFamily, x.NativeSelectorPresent, x.EmpiricalYukawaImported, x.Verdict, x.Reason)
}
func FormatCandidate(x AxiomCandidate) string {
	return fmt.Sprintf("name=%q executed=%v kind=%q cost=%d gauge=%v JGamma=%v empirical_independent=%v imports_yukawa=%v selects_ray=%v fixes_values=%v ckm=%v pmns=%v diagonal_only=%v free_params=%d native=%v promoted=%v verdict=%s reason=%s", x.Name, x.Executed, x.AxiomKind, x.MathematicalCost, x.GaugeCompatible, x.CompatibleWithJGamma, x.EmpiricalIndependent, x.ImportsObservedYukawa, x.SelectsCoefficientRay, x.FixesCoefficientValues, x.CKMCapacity, x.PMNSCapacity, x.DiagonalOnly, x.FreeRealParameters, x.NativeToCurrentAsha, x.PromotedToTheorem, x.Verdict, x.Reason)
}
func FormatRanking(x MinimalityRanking) string {
	return fmt.Sprintf("executed=%v ranked=%s least_cost_name=%q least_cost=%d least_still_axiom=%v least_ckm=%v least_fixes_angles=%v no_candidate_native=%v verdict=%s", x.Executed, strings.Join(x.RankedNames, "; "), x.LeastCostName, x.LeastCost, x.LeastCostStillAxiom, x.LeastCostCKMCapacity, x.LeastCostFixesAngles, x.NoCandidateNative, x.Verdict)
}
func FormatCapacity(x CapacityAudit) string {
	return fmt.Sprintf("executed=%v conditional_ckm=%v conditional_pmns=%v fixes_angles=%v native=%v curve_fitting_candidate=%v best_empirical_independent=%q required_extra_data=%q verdict=%s reason=%s", x.Executed, x.ConditionalCKMAvailable, x.ConditionalPMNSAvailable, x.AnyCandidateFixesAngles, x.AnyCandidateNative, x.AnyCandidateCurveFitting, x.BestEmpiricalIndependent, x.RequiredExtraData, x.Verdict, x.Reason)
}
func FormatScenario(x ModuliScenario) string {
	return fmt.Sprintf("scenario=%q status=%s dim=%d masses3=%v ckm=%v pmns=%v coefficients_fixed=%v native_reduction=%v conditional=%v empirical_fitting=%v reason=%s", x.Name, x.Status, x.ModuliDim, x.ThreeMassesPossible, x.CKMPossible, x.PMNSPossible, x.CoefficientsFixed, x.NativeReduction, x.ConditionalOnly, x.EmpiricalFitting, x.Reason)
}
func FormatModuli(x ModuliImpact) string {
	return fmt.Sprintf("start_dim=%d best_native_dim=%d native_reduction=%v conditional_mixing=%v coefficients_free=%v firewall=%v verdict=%s", x.StartDim, x.BestNativeDim, x.NativeReductionBelow13, x.ConditionalMixingCapacity, x.CoefficientsRemainFree, x.FirewallPreserved, x.Verdict)
}
func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("executed=%v no_masses=%v no_ckm=%v no_pmns=%v no_yukawa_matrices=%v axioms_quarantined=%v no_native_derivation=%v verdict=%s", x.Executed, x.NoObservedMassesImported, x.NoCKMImported, x.NoPMNSImported, x.NoYukawaMatricesInserted, x.AxiomsQuarantined, x.NoNativeDerivationClaimed, x.Verdict)
}
func FormatNext(x NextStep) string {
	return fmt.Sprintf("gate=%d title=%q reason=%s primary_task=%s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}

func RenderCandidateSummary(cs []AxiomCandidate) string {
	var b strings.Builder
	for _, c := range cs {
		b.WriteString(FormatCandidate(c) + "\n")
	}
	return strings.TrimRight(b.String(), "\n")
}

func RenderMarkdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 415 Registry Audit — Family Boundary Condition / Sector Source Axiom Minimality Sieve\n\n")
	b.WriteString("## Claim tested\n\nGate 415 asks what minimal boundary/source axiom would be required after Gate 414 failed to derive family coefficients from trace or curvature functionals. It ranks candidate axioms without promoting any of them to native ASHA theorems and without importing observed Yukawa matrices.\n\n")
	sections := []struct{ title, body string }{
		{"Prior boundary inherited", FormatInheritance(a.Inheritance)},
		{"Boundary/source arena", FormatArena(a.Arena)},
		{"Minimality ranking", FormatRanking(a.Ranking)},
		{"CKM/PMNS capacity audit", FormatCapacity(a.Capacity)},
		{"Empirical firewall", FormatFirewall(a.Firewall)},
	}
	for _, s := range sections {
		b.WriteString("## " + s.title + "\n\n```text\n" + s.body + "\n```\n\n")
	}
	b.WriteString("## Candidate axiom ledger\n\n```text\n")
	b.WriteString(RenderCandidateSummary(a.Candidates))
	b.WriteString("\n```\n\n")
	b.WriteString("## Moduli impact\n\n```text\n")
	b.WriteString(FormatModuli(a.Moduli))
	for _, s := range a.Moduli.Scenarios {
		b.WriteString("\n" + FormatScenario(s))
	}
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
