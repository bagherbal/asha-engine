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
	if s == ScenarioAll || s == ScenarioCosmology || s == ScenarioCI {
		add(e.cosmologySection())
	}
	if len(sections) == 0 {
		return Report{}, fmt.Errorf("scenario %q produced no sections", s)
	}
	checks := e.Checks()
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
	return Section{Name: "Dark-sector scenarios", Summary: "Heavy finite sectors are scenario-classified. Stable thermal B-gap Majorana relics are rejected by overclosure; decay/nonthermal routes require extra history.", Quantities: []Quantity{{Symbol: "M_B", Value: e.Bridge.HeavyBGapMajoranaGeV, Unit: "GeV", Status: StatusBridge}, {Symbol: "Ω_candidate/Ω_DM", Value: e.Bridge.MajoranaOverclosure, Status: StatusFailedRoute}}, Boundaries: e.Environment.DarkSectorScenarios}
}

func (e Engine) cosmologySection() Section {
	return Section{Name: "Cosmology boundary", Summary: "The spectral-action cosmological term is present, but observed cosmology requires continuum subtraction, history, or holographic/dilaton bridge data.", Boundaries: e.Environment.CosmologyScenarios}
}
