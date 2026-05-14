package asha

type Environment struct {
	PACS                 []string   `json:"pacs"`
	Keywords             []string   `json:"keywords"`
	EmpiricalCoordinates []Boundary `json:"empirical_coordinates"`
	CosmologyScenarios   []Boundary `json:"cosmology_scenarios"`
	DarkSectorScenarios  []Boundary `json:"dark_sector_scenarios"`
	EpistemologicalSeals []Boundary `json:"epistemological_seals"`
}

func NewEnvironment() Environment {
	return Environment{
		PACS:     []string{"02.10.Ud", "04.50.Kd", "11.10.Nx", "11.15.-q", "12.10.-g", "12.15.Ff"},
		Keywords: []string{"Clifford algebra", "noncommutative geometry", "spectral action", "Standard Model", "Higgs", "flavor firewall", "finite geometry"},
		EmpiricalCoordinates: []Boundary{
			{Name: "charged flavor", Formula: "dim M_charged^native = 13", Status: StatusEnvironmental, Meaning: "Yukawa values and CKM coordinates are not native ASHA outputs."},
			{Name: "K/X/Y coefficients", Formula: "{a_s,b_s,c_s}_{s=u,d,e}", Status: StatusEnvironmental, Meaning: "conditional family source coefficients remain boundary data."},
		},
		CosmologyScenarios: []Boundary{
			{Name: "spectral-action cosmological term", Formula: "48 f₄ Λ⁴ + subtraction/renormalization", Status: StatusBridge, Meaning: "bare term exists; observed ρΛ needs continuum/history rule."},
			{Name: "holographic/dilaton bridge", Formula: "ρΛ ~ M_P²/L², Λ → Λ(x)", Status: StatusBridge, Meaning: "possible pathway, not native prediction."},
			{Name: "cosmological coordinates", Formula: "(Ω_DM h², ρΛ, t_universe, η_B)", Status: StatusEnvironmental, Meaning: "history/state dependent and not predicted by current law-space."},
		},
		DarkSectorScenarios: []Boundary{
			{Name: "B-gap Majorana stable thermal relic", Formula: "Ω_candidate/Ω_DM ~ 1.3×10¹³", Status: StatusFailedRoute, Meaning: "simple stable thermal interpretation is rejected by overclosure."},
			{Name: "decaying/portal heavy sector", Formula: "Ω_heavy h² = 0 after decay assumptions", Status: StatusBridge, Meaning: "allowed only with sealed decay/portal dynamics."},
			{Name: "nonthermal/axion-like routes", Formula: "requires production-history axiom", Status: StatusEnvironmental, Meaning: "not native ASHA output."},
		},
		EpistemologicalSeals: []Boundary{
			{Name: "q4 sector", Formula: "q4 ∈ contact spectral sector, q4 ∉ HΦ selector", Status: StatusFailedRoute, Meaning: "contact invariant not promoted to scalar/flavor theorem."},
			{Name: "HΦ flavor blindness", Formula: "HΦ ⇒ weak doublet + scalar potential, not flavor selector", Status: StatusFailedRoute, Meaning: "native scalar selectors are central or pair-degenerate."},
			{Name: "family axioms", Formula: "K/X/Y ⇒ capacity, not coefficient prediction", Status: StatusQuarantined, Meaning: "family extension is explicit and quarantined."},
		},
	}
}
