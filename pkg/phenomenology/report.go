package phenomenology

import (
	"fmt"
	"math"
	"strings"
)

type PhenomenologyReport struct {
	Executed       bool
	VacuumFate     VacuumFateResult
	VacuumFates    []VacuumFateResult
	DarkMatter     DarkMatterResult
	CosmologicalCC CosmologicalConstantResult
	Statuses       []string
	Accuracy       map[string]string
	BottomLine     string
}

func ComputeReport() PhenomenologyReport {
	vfs := ComputeVacuumFateEnsemble()
	vf := vfs[len(vfs)-1]
	dm := ComputeDarkMatterConstraint()
	cc := ComputeCosmologicalConstantSubtraction()
	statuses := []string{StatusPhenomenologyLayerComplete}
	statuses = append(statuses, vf.Statuses...)
	statuses = append(statuses, dm.Statuses...)
	statuses = append(statuses, cc.Statuses...)
	return PhenomenologyReport{
		Executed:       true,
		VacuumFate:     vf,
		VacuumFates:    vfs,
		DarkMatter:     dm,
		CosmologicalCC: cc,
		Statuses:       uniqueStrings(statuses),
		Accuracy: map[string]string{
			"vacuum_fate":        "qualitative/conditional: one-loop RG with pole-mass seeds; not a precision SM vacuum-lifetime computation",
			"dark_matter":        "constraint-level: exact yield arithmetic for the assumed B-gap mass; no native production kernel",
			"cosmological_const": "order-of-magnitude fine-tuning diagnostic; no native subtraction theorem",
		},
		BottomLine: "The empirical-seal layer computes conditional consequences. It predicts ASHA+empirical vacuum metastability in the simple one-loop audit, rules out an unsuppressed stable thermal B-gap relic by overclosure, and confirms that the cosmological constant still needs an environmental subtraction theorem.",
	}
}

func uniqueStrings(in []string) []string {
	seen := make(map[string]bool, len(in))
	out := make([]string, 0, len(in))
	for _, s := range in {
		if !seen[s] {
			seen[s] = true
			out = append(out, s)
		}
	}
	return out
}

func (r PhenomenologyReport) Markdown() string {
	var b strings.Builder
	fmt.Fprintf(&b, "# ASHA Phenomenology Package — Empirical-Seal Prediction Audit\n\n")
	fmt.Fprintf(&b, "## Bottom line\n\n%s\n\n", r.BottomLine)
	fmt.Fprintf(&b, "## Environmental quarantine inputs\n\n")
	fmt.Fprintf(&b, "| Input | Value |\n|---|---:|\n")
	fmt.Fprintf(&b, "| Top mass | %.2f GeV |\n", TopMassGeV)
	fmt.Fprintf(&b, "| Higgs mass | %.2f GeV |\n", HiggsMassGeV)
	fmt.Fprintf(&b, "| alpha_s(mZ) | %.4f |\n", StrongCoupling)
	fmt.Fprintf(&b, "| Z mass | %.4f GeV |\n", ZBosonMassGeV)
	fmt.Fprintf(&b, "| Target Omega_c h^2 | %.3f |\n", TargetRelicDensity)
	fmt.Fprintf(&b, "| Target rho_Lambda/M_Pl^4 | %.1e |\n\n", TargetDarkEnergy)

	fmt.Fprintf(&b, "## Vacuum fate: one-loop conditional audit\n\n")
	fmt.Fprintf(&b, "| Seed mode | y_t start | λ before threshold | λ after threshold | instability scale | λ_min | log10 lifetime/yr |\n")
	fmt.Fprintf(&b, "|---|---:|---:|---:|---:|---:|---:|\n")
	for _, vf := range r.VacuumFates {
		lt := "+Inf"
		if !math.IsInf(vf.Log10LifetimeYears, 1) {
			lt = fmt.Sprintf("%.6f", vf.Log10LifetimeYears)
		}
		fmt.Fprintf(&b, "| %s | %.9f | %.9f | %.9f | %.6e GeV | %.9f | %s |\n", vf.SeedMode, vf.InitialYTop, vf.LambdaBeforeThreshold, vf.LambdaAfterThreshold, vf.InstabilityScaleGeV, vf.LambdaMin, lt)
	}
	vf := r.VacuumFate
	fmt.Fprintf(&b, "\nDefault precision lane: `%s`. Initial λ(mZ)=%.9f, g3(mZ)=%.9f, bounce action proxy=%.6f, log10(age/yr)=%.6f.\n\n", vf.SeedMode, vf.InitialLambda, vf.InitialG3, vf.BounceAction, vf.AgeUniverseLog10Years)
	fmt.Fprintf(&b, "**Verdict:** conditional metastability = %t. %s\n\n", vf.Metastable, vf.PrecisionWarning)

	dm := r.DarkMatter
	fmt.Fprintf(&b, "## Dark matter: B-gap Majorana constraint\n\n")
	fmt.Fprintf(&b, "| Quantity | Value |\n|---|---:|\n")
	fmt.Fprintf(&b, "| candidate mass | %.6e GeV |\n", dm.MassGeV)
	fmt.Fprintf(&b, "| required yield Y=n/s for Omega=0.120 | %.6e |\n", dm.RequiredYield)
	fmt.Fprintf(&b, "| relativistic thermal yield, g=2, g*=106.75 | %.6e |\n", dm.ThermalRelativisticYield)
	fmt.Fprintf(&b, "| stable thermal Omega h^2 | %.6e |\n", dm.ThermalStableOmegaH2)
	fmt.Fprintf(&b, "| overclosure factor | %.6e |\n", dm.OverclosureFactor)
	fmt.Fprintf(&b, "| required fraction of thermal yield | %.6e |\n\n", dm.RequiredFractionOfThermal)
	fmt.Fprintf(&b, "**Verdict:** %s\n\n", dm.Interpretation)

	cc := r.CosmologicalCC
	fmt.Fprintf(&b, "## Cosmological constant: subtraction severity\n\n")
	fmt.Fprintf(&b, "| Quantity | Value |\n|---|---:|\n")
	fmt.Fprintf(&b, "| convention | %s |\n", cc.Convention)
	fmt.Fprintf(&b, "| bare rho / M_Pl^4 | %.6e |\n", cc.BareVacuumPlanckUnits)
	fmt.Fprintf(&b, "| target rho_Lambda / M_Pl^4 | %.6e |\n", cc.TargetDarkEnergy)
	fmt.Fprintf(&b, "| required counterterm | %.6e |\n", cc.RequiredCounterterm)
	fmt.Fprintf(&b, "| cancellation ratio | %.6e |\n", cc.CancellationRatio)
	fmt.Fprintf(&b, "| decimal digits of cancellation | %.3f |\n\n", cc.DigitsOfCancellation)
	fmt.Fprintf(&b, "**Verdict:** %s\n\n", cc.Interpretation)

	fmt.Fprintf(&b, "## Status ledger\n\n")
	for _, s := range r.Statuses {
		fmt.Fprintf(&b, "- `%s`\n", s)
	}
	return b.String()
}
