package asha

import "fmt"

type Engine struct {
	Metadata    Metadata      `json:"metadata"`
	Empirical   EmpiricalData `json:"empirical_data"`
	Geometry    Geometry      `json:"geometry"`
	Electroweak Electroweak   `json:"electroweak"`
	Bridge      Bridge        `json:"bridge"`
	Family      Family        `json:"family"`
	Environment Environment   `json:"environment"`
}

func New(opts ...Option) Engine {
	o := defaultOptions()
	for _, opt := range opts {
		opt(&o)
	}
	if o.PlanckMassGeV <= 0 {
		o.PlanckMassGeV = DefaultPlanckMassGeV
	}
	if o.Beta == 0 {
		o.Beta = DefaultBeta
	}
	geo := NewGeometry()
	env := NewEnvironment()
	return Engine{
		Metadata: Metadata{
			RuntimeVersion: RuntimeVersion,
			LatestGate:     LatestGate,
			Title:          "ASHA final runtime board",
			Keywords:       env.Keywords,
			PACS:           env.PACS,
			Source:         "github.com/bagherbal/asha-engine",
		},
		Empirical:   EmpiricalData{PlanckMassGeV: o.PlanckMassGeV, Beta: o.Beta, ObservedHiggsMassGeV: o.ObservedHiggsMassGeV, ObservedOmegaDMH2: o.ObservedOmegaDMH2, UseObservedComparators: o.UseObservedComparators},
		Geometry:    geo,
		Electroweak: NewElectroweak(),
		Bridge:      NewBridge(o.PlanckMassGeV),
		Family:      NewFamily(o.Beta),
		Environment: env,
	}
}

type Report struct {
	Metadata   Metadata   `json:"metadata"`
	Scenario   Scenario   `json:"scenario"`
	Sections   []Section  `json:"sections"`
	Checks     []Check    `json:"checks"`
	Boundaries []Boundary `json:"boundaries"`
	Verdict    string     `json:"verdict"`
}

func (e Engine) Report(s Scenario) (Report, error) {
	if s == "" {
		s = ScenarioAll
	}
	sections := make([]Section, 0, 8)
	add := func(sec Section) { sections = append(sections, sec) }
	if s == ScenarioAll || s == ScenarioNative || s == ScenarioCI {
		add(e.nativeSection())
	}
	if s == ScenarioAll || s == ScenarioHiggs || s == ScenarioCI {
		add(e.higgsSection())
	}
	if s == ScenarioAll || s == ScenarioFamily || s == ScenarioCI {
		add(e.familySection())
	}
	if s == ScenarioAll || s == ScenarioDarkStableThermal || s == ScenarioCI {
		add(e.darkSection())
	}
	if s == ScenarioAll || s == ScenarioCosmology || s == ScenarioEnvironment || s == ScenarioCI {
		add(e.cosmologySection())
	}
	if s == ScenarioAll || s == ScenarioCosmology || s == ScenarioEnvironment || s == ScenarioCI {
		add(e.vacuumFateSection())
	}
	if len(sections) == 0 {
		return Report{}, fmt.Errorf("scenario %q produced no sections", s)
	}
	checks := e.Checks()
	for _, sec := range sections {
		checks = append(checks, sec.Checks...)
	}
	boundaries := e.Boundaries()
	verdict := "PASS: ASHA runtime board is internally consistent; native law-space separated from bridge, quarantined, and environmental data."
	for _, c := range checks {
		if !c.Passed {
			verdict = "FAIL: runtime consistency check failed"
			break
		}
	}
	return Report{Metadata: e.Metadata, Scenario: s, Sections: sections, Checks: checks, Boundaries: boundaries, Verdict: verdict}, nil
}

func (e Engine) Checks() []Check {
	out := []Check{}
	out = append(out, e.Geometry.Checks()...)
	out = append(out, e.Electroweak.Checks()...)
	out = append(out, e.Bridge.Checks()...)
	out = append(out, e.Family.Checks()...)
	out = append(out, Check{Name: "Latest gate marker", Passed: e.Metadata.LatestGate == LatestGate, Detail: fmt.Sprintf("gate=%d", e.Metadata.LatestGate)})
	return out
}

func (e Engine) Boundaries() []Boundary {
	out := []Boundary{}
	out = append(out, e.Environment.EmpiricalCoordinates...)
	out = append(out, e.Environment.CosmologyScenarios...)
	out = append(out, e.Environment.DarkSectorScenarios...)
	out = append(out, e.Environment.EpistemologicalSeals...)
	return out
}

func (e Engine) nativeSection() Section {
	q := []Quantity{
		{Symbol: "dim Cℓ(1,7)", Value: float64(e.Geometry.CliffordDimension), Formula: "2^8", Status: StatusNative},
		{Symbol: "grades", Text: fmt.Sprint(e.Geometry.GradeDimensions), Formula: "dim Λ^k R^8 = C(8,k)", Status: StatusNative},
		{Symbol: "rank(P_B)", Value: float64(e.Geometry.RankPB), Status: StatusNative},
		{Symbol: "rank(P_G)", Value: float64(e.Geometry.RankPG), Status: StatusNative},
		{Symbol: "dim K", Value: float64(e.Geometry.DimK), Formula: "Im(P_B)∩Im(P_G)", Status: StatusNative},
		{Symbol: "I_BG", Text: e.Geometry.ContactIndex.String(), Formula: "dim K / 7", Status: StatusNative},
		{Symbol: "k_Y", Text: e.Electroweak.KY.String(), Formula: "Tr(Y²)/Tr(T₃²)", Status: StatusNative},
		{Symbol: "sin²θ*", Text: e.Electroweak.Sin2ThetaStar.String(), Formula: "1/(1+k_Y)", Status: StatusNative},
	}
	return Section{Name: "Native finite law-space", Summary: "Finite measurement ladder, Boolean/G₂ contact vacuum, charge skeleton, and inner-fluctuation field inventory.", Quantities: q, Checks: append(e.Geometry.Checks(), e.Electroweak.Checks()...)}
}

func (e Engine) higgsSection() Section {
	return Section{Name: "Higgs and coefficient bridge", Summary: "One-form edge measure and Pfaffian scale lane yield the tree-level Higgs proxy; pole-mass and RG thresholds remain bridge work.", Quantities: e.Bridge.Quantities(), Checks: e.Bridge.Checks(), Boundaries: []Boundary{{Name: "Higgs pole mass", Formula: "m_H^tree + RG + thresholds + self-energy", Status: StatusBridge, Meaning: "runtime reports tree proxy only"}}}
}

func (e Engine) familySection() Section {
	return Section{Name: "Family/flavor frontier", Summary: "Native flavor remains 13-dimensional; quarantined K/X/Y axioms activate hierarchy, mixing, and CP capacity with 9 symbolic charged coefficients.", Quantities: e.Family.Quantities(), Checks: e.Family.Checks(), Boundaries: e.Environment.EmpiricalCoordinates}
}

func (e Engine) darkSection() Section {
	dm := ComputeDarkMatterConditional(e.Bridge.HeavyBGapMajoranaGeV)
	q := []Quantity{
		{Symbol: "M_B", Value: dm.MassGeV, Unit: "GeV", Formula: "B-gap heavy Majorana scale", Status: StatusBridge},
		{Symbol: "Ω_DM h² target", Value: dm.TargetOmegaH2, Formula: "observational comparator", Status: StatusEnvironmental},
		{Symbol: "Y_required", Value: dm.RequiredYield, Formula: "Ω h² ρ_c/(m s_0)", Status: StatusBridge, Note: "yield needed if B-gap particle were all dark matter"},
		{Symbol: "Y_thermal", Value: dm.ThermalRelativisticYield, Formula: "135 ζ(3)/(8π⁴) · g/g_*S", Status: StatusBridge, Note: "stable relativistic thermal abundance stress test"},
		{Symbol: "Ω_thermal h²", Value: dm.ThermalStableOmegaH2, Formula: "m Y_thermal s_0/ρ_c", Status: StatusFailedRoute},
		{Symbol: "Ω_thermal/Ω_DM", Value: dm.OverclosureFactor, Formula: "stable thermal B-gap relic overclosure", Status: StatusFailedRoute},
		{Symbol: "Y_required/Y_thermal", Value: dm.RequiredFractionOfThermal, Formula: "required suppression / dilution fraction", Status: StatusBridge, Note: "viable only with nonthermal production, dilution, or decay history"},
	}
	checks := []Check{
		{Name: "Stable thermal B-gap Majorana rejected", Passed: dm.OverclosureFactor > 1e12, Detail: "overclosure ratio computed, not guessed"},
		{Name: "Suppressed/nonthermal target computed", Passed: dm.RequiredFractionOfThermal > 0 && dm.RequiredFractionOfThermal < 1e-10, Detail: "conditional yield fraction exists, production history remains sealed"},
	}
	return Section{Name: "Dark-sector conditional scenarios", Summary: "The runtime now computes the viable and rejected dark-sector paths. A stable thermal B-gap Majorana relic is ruled out by overclosure; a suppressed/nonthermal or decaying route remains a conditional cosmological-history bridge.", Quantities: q, Checks: checks, Boundaries: e.Environment.DarkSectorScenarios}
}

func (e Engine) cosmologySection() Section {
	cc := ComputeCosmologyConditional(e.Bridge.PlanckMassGeV, e.Bridge.VPfGeV)
	q := []Quantity{
		{Symbol: "ρ_bare/M_P⁴", Value: cc.BareVacuumPlanckUnits, Formula: "48/π²", Status: StatusBridge, Note: "diagnostic CCM bare vacuum convention f₄=1, Λ=M_P"},
		{Symbol: "ρ_Λ/M_P⁴ target", Value: cc.TargetDarkEnergyPlanckUnits, Formula: "diagnostic observed-scale comparator", Status: StatusEnvironmental},
		{Symbol: "counterterm severity", Value: cc.CancellationRatio, Formula: "ρ_bare/ρ_target", Status: StatusEnvironmental},
		{Symbol: "digits cancellation", Value: cc.DigitsOfCancellation, Formula: "log₁₀(ρ_bare/ρ_target)", Status: StatusEnvironmental},
		{Symbol: "L·M_P target", Value: cc.HolographicLMpForTarget, Formula: "1/sqrt(ρ_Λ/M_P⁴)", Status: StatusBridge, Note: "holographic/dilaton bridge scale for 10^-120 target"},
		{Symbol: "L·M_P Gate344 target", Value: cc.Gate344HolographicLMpForTarget, Formula: "1/sqrt(10^-122)", Status: StatusBridge, Note: "alternate Gate-344 target convention"},
		{Symbol: "(v_Pf/M_P)^4", Value: cc.ElectroweakVacuumPlanckFourth, Formula: "ρ^4", Status: StatusBridge},
		{Symbol: "EW vacuum / target", Value: cc.EWVacuumOverTarget, Formula: "(v_Pf/M_P)^4 / 10^-120", Status: StatusFailedRoute},
		{Symbol: "EW vacuum / Gate344 target", Value: cc.EWVacuumOverGate344Target, Formula: "(v_Pf/M_P)^4 / 10^-122", Status: StatusFailedRoute},
	}
	checks := []Check{
		{Name: "Cosmological constant not solved natively", Passed: cc.DigitsOfCancellation > 120, Detail: "bare spectral term needs subtraction/history rule"},
		{Name: "Holographic/dilaton bridge computable", Passed: cc.HolographicLMpForTarget > 1e59, Detail: "conditional IR-UV scale is numerical but not native saturation theorem"},
	}
	return Section{Name: "Cosmology conditional scenarios", Summary: "The runtime reports conditional cosmology numbers with warnings: bare spectral-action vacuum severity, holographic/dilaton target scales, and electroweak-vacuum tension. These are bridge diagnostics, not native predictions of dark energy.", Quantities: q, Checks: checks, Boundaries: e.Environment.CosmologyScenarios}
}

func (e Engine) vacuumFateSection() Section {
	vfs := ComputeVacuumFateConditionals(e.Bridge.PlanckMassGeV, e.Bridge.HeavyBGapMajoranaGeV)
	q := make([]Quantity, 0, 18)
	for _, vf := range vfs {
		prefix := vf.SeedMode
		q = append(q,
			Quantity{Symbol: prefix + " λ_before", Value: vf.LambdaBeforeThreshold, Formula: "one-loop RG to M_B", Status: StatusBridge},
			Quantity{Symbol: prefix + " λ_after", Value: vf.LambdaAfterThreshold, Formula: "λ_before + Δλ_ASHA", Status: StatusBridge},
			Quantity{Symbol: prefix + " μ_inst", Value: vf.InstabilityScaleGeV, Unit: "GeV", Formula: "λ crossing scale", Status: StatusBridge},
			Quantity{Symbol: prefix + " λ_min", Value: vf.LambdaMin, Formula: "conditional one-loop minimum", Status: StatusBridge},
			Quantity{Symbol: prefix + " S_E", Value: vf.BounceAction, Formula: "8π²/(3|λ_min|)", Status: StatusBridge},
			Quantity{Symbol: prefix + " log10 τ/yr", Value: vf.Log10LifetimeYears, Formula: "conditional bounce proxy", Status: StatusBridge},
		)
	}
	checks := []Check{
		{Name: "Vacuum-fate ensemble computed", Passed: len(vfs) == 2, Detail: "pole and one-loop-QCD top seeds audited"},
		{Name: "Vacuum fate remains conditional", Passed: true, Detail: "requires empirical top/Higgs inputs and continuum RG scheme"},
	}
	return Section{Name: "Vacuum-fate conditional scenario", Summary: "A conditional one-loop RG/bounce stress test can be computed once empirical top/Higgs inputs and the ASHA B-gap threshold jump are supplied. It is useful phenomenology, but it is not a native ASHA universe-lifetime theorem.", Quantities: q, Checks: checks, Boundaries: []Boundary{{Name: "vacuum lifetime", Formula: "top/Higgs/RG scheme + threshold convention + bounce prefactor", Status: StatusEnvironmental, Meaning: "conditional scenario only; no native lifetime prediction"}}}
}
