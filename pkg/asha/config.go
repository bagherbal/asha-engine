package asha

const (
	RuntimeVersion = "gate425-runtime-env-scenarios-20260514"
	LatestGate     = 425
	DefaultBeta    = 1.0
	// Project convention: non-reduced Planck mass in GeV used by the Pfaffian lane.
	DefaultPlanckMassGeV = 1.22089e19
)

type Metadata struct {
	RuntimeVersion string   `json:"runtime_version"`
	LatestGate     int      `json:"latest_gate"`
	Title          string   `json:"title"`
	Keywords       []string `json:"keywords"`
	PACS           []string `json:"pacs"`
	Source         string   `json:"source"`
}

type EmpiricalData struct {
	PlanckMassGeV          float64 `json:"planck_mass_gev"`
	Beta                   float64 `json:"beta"`
	ObservedHiggsMassGeV   float64 `json:"observed_higgs_mass_gev,omitempty"`
	ObservedOmegaDMH2      float64 `json:"observed_omega_dm_h2,omitempty"`
	UseObservedComparators bool    `json:"use_observed_comparators"`
}

type Options struct {
	PlanckMassGeV          float64
	Beta                   float64
	ObservedHiggsMassGeV   float64
	ObservedOmegaDMH2      float64
	UseObservedComparators bool
}

type Option func(*Options)

func WithPlanckMassGeV(x float64) Option { return func(o *Options) { o.PlanckMassGeV = x } }
func WithBeta(x float64) Option          { return func(o *Options) { o.Beta = x } }
func WithObservedHiggsMassGeV(x float64) Option {
	return func(o *Options) { o.ObservedHiggsMassGeV = x; o.UseObservedComparators = true }
}
func WithObservedOmegaDMH2(x float64) Option {
	return func(o *Options) { o.ObservedOmegaDMH2 = x; o.UseObservedComparators = true }
}

func defaultOptions() Options { return Options{PlanckMassGeV: DefaultPlanckMassGeV, Beta: DefaultBeta} }
