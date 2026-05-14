// Package cosmologicalobservablesdarksector implements Gate 386:
// Cosmological Observables & Dark Sector Prediction Sieve.
//
// Gate 385 geometrically selected the finite one-form edge measure and sealed
// the CCM+Pfaffian tree-level Higgs proxy near 125 GeV.  Gate 386 asks whether
// that sealed Higgs lane plus the B-gap / heavy-sector threshold ledger is now
// sufficient to compute hard macroscopic observables: a dark-matter relic
// density and the cosmological fate of the electroweak vacuum.
//
// The answer is deliberately channel-sensitive.  A sealed Higgs boundary gives
// a valid initial quartic for continuum calculations, but a relic abundance
// still requires a dark-sector Lagrangian, stability/decay theorem, production
// mechanism, reheating state, and Boltzmann kernel.  Vacuum fate still requires
// the physical top/Yukawa and gauge running, threshold matching, lambda minimum,
// and bounce prefactor.  The package therefore produces a computability ledger
// rather than fitting observed cosmology.
package cosmologicalobservablesdarksector

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate385 "github.com/bagherbal/asha-engine/pkg/bridge/innerfluctuationedgemeasure"
)

const (
	AuditID = "GATE386-COSMOLOGICAL-OBSERVABLES-DARK-SECTOR-PREDICTION-SIEVE"

	StatusGate385SealedHiggsInherited   = "CONDITIONAL_SUPPORT_GATE385_SEALED_HIGGS_PROXY_INHERITED"
	StatusBGapThresholdLedgerInherited  = "CONDITIONAL_SUPPORT_BGAP_THRESHOLD_LEDGER_INHERITED"
	StatusDarkSectorCandidateFormalized = "CONDITIONAL_SUPPORT_DARK_SECTOR_CANDIDATE_FORMALIZED"
	StatusBoltzmannEquationFormalized   = "CONDITIONAL_SUPPORT_BOLTZMANN_EQUATION_FORMALIZED"
	StatusVacuumStabilityRGFormalized   = "CONDITIONAL_SUPPORT_VACUUM_STABILITY_RG_FORMALIZED"
	StatusThresholdJumpFormalized       = "CONDITIONAL_SUPPORT_THRESHOLD_JUMP_FORMALIZED"
	StatusBounceActionFormalized        = "CONDITIONAL_SUPPORT_BOUNCE_ACTION_FORMALIZED"
	StatusObservableFirewallPreserved   = "CONDITIONAL_SUPPORT_COSMOLOGICAL_OBSERVABLE_FIREWALL_PRESERVED"
	StatusComputableTargetsOpened       = "CONDITIONAL_SUPPORT_COMPUTABLE_COSMOLOGICAL_TARGETS_OPENED"

	StatusTensionMajoranaScaleNotStability      = "CONDITIONAL_TENSION_MAJORANA_SCALE_IS_NOT_DARK_MATTER_STABILITY"
	StatusTensionRelicNeedsProductionHistory    = "CONDITIONAL_TENSION_RELIC_DENSITY_NEEDS_PRODUCTION_HISTORY"
	StatusTensionSealedHiggsNotFullRGTrajectory = "CONDITIONAL_TENSION_SEALED_HIGGS_PROXY_NOT_FULL_RG_TRAJECTORY"
	StatusTensionTopYukawaStillFlavorSeal       = "CONDITIONAL_TENSION_TOP_YUKAWA_REMAINS_EMPIRICAL_FLAVOR_SEAL"
	StatusTensionThresholdSignNeedsMatching     = "CONDITIONAL_TENSION_THRESHOLD_JUMP_SIGN_NEEDS_MATCHING_CONVENTION"
	StatusTensionBounceNeedsLambdaMinimum       = "CONDITIONAL_TENSION_BOUNCE_ACTION_NEEDS_LAMBDA_MINIMUM"

	StatusFailedCosmologicalObservablesNotDerived = "FAILED_ROUTE_COSMOLOGICAL_OBSERVABLES_NOT_DERIVED"
	StatusFailedDarkMatterRelicNotDerived         = "FAILED_ROUTE_DARK_MATTER_RELIC_DENSITY_NOT_DERIVED"
	StatusFailedOmegaDMNotComputed                = "FAILED_ROUTE_OMEGA_DM_H2_NOT_COMPUTED"
	StatusFailedStableDarkCandidateNotDerived     = "FAILED_ROUTE_STABLE_DARK_CANDIDATE_NOT_DERIVED"
	StatusFailedBoltzmannKernelNotClosed          = "FAILED_ROUTE_BOLTZMANN_KERNEL_NOT_CLOSED"
	StatusFailedVacuumStabilityNotDerived         = "FAILED_ROUTE_VACUUM_STABILITY_NOT_DERIVED"
	StatusFailedUniverseLifetimeNotDerived        = "FAILED_ROUTE_UNIVERSE_LIFETIME_NOT_DERIVED"
	StatusFailedEuclideanBounceNotComputed        = "FAILED_ROUTE_EUCLIDEAN_BOUNCE_ACTION_NOT_COMPUTED"
	StatusFailedFullNumericalTOENotClosed         = "FAILED_ROUTE_FULL_NUMERICAL_TOE_CLOSURE_STILL_NOT_REACHED"
)

const (
	bGapMajoranaScaleGeV      = 1.46774973718e6
	heavyIntermediateScaleGeV = 6.650726476871e11
	thresholdJumpDeltaLambda  = -0.097846792207
	omegaDMObservedReference  = 0.120
	targetHiggsGeV            = 125.10
)

type InheritedHiggsBoundary struct {
	Executed          bool
	SourceGate        int
	LambdaEW          float64
	HiggsMassProxyGeV float64
	PfaffianVEVGeV    float64
	TreeProxySealed   bool
	PoleMassDerived   bool
	UsesEdgeMeasure   bool
	DirectAnswer      string
	Verdict           string
}

type HeavySectorLedger struct {
	Executed                  bool
	Candidate                 string
	BGapScaleGeV              float64
	HeavyIntermediateScaleGeV float64
	ThresholdJumpDeltaLambda  float64
	GeometricallyMandated     bool
	StableRelicTheorem        bool
	DecayWidthDerived         bool
	AnnihilationRateDerived   bool
	DirectAnswer              string
	Verdict                   string
}

type RequiredInput struct {
	Name        string
	Required    bool
	Native      bool
	Value       string
	MissingWhy  string
	BlocksClaim bool
}

type DarkMatterPrediction struct {
	Executed                 bool
	BoltzmannEquation        string
	Candidate                string
	MassScaleGeV             float64
	Inputs                   []RequiredInput
	NativeMassScaleAvailable bool
	StableCandidateDerived   bool
	BoltzmannKernelClosed    bool
	OmegaH2Derived           bool
	OmegaH2                  float64
	ReferenceOmegaH2         float64
	IndependentPrediction    bool
	DirectAnswer             string
	Verdict                  string
}

type RGInput struct {
	Name       string
	Required   bool
	Native     bool
	Value      string
	MissingWhy string
	BlocksFate bool
}

type VacuumStabilityPrediction struct {
	Executed                 bool
	LambdaEW                 float64
	HiggsMassProxyGeV        float64
	ThresholdScaleGeV        float64
	ThresholdDeltaLambda     float64
	OneLoopBetaLambda        string
	Inputs                   []RGInput
	FullRGTrajectoryDerived  bool
	ThresholdMatchingDerived bool
	LambdaMinimumDerived     bool
	AbsoluteStabilityDerived bool
	MetastabilityDerived     bool
	EuclideanBounceFormula   string
	EuclideanBounceAction    float64
	LifetimeYears            float64
	LifetimeDerived          bool
	IndependentPrediction    bool
	DirectAnswer             string
	Verdict                  string
}

type ObservableCensus struct {
	Executed                 bool
	RequestedObservables     int
	HardPredictionsDerived   int
	ConditionalTargetsOpened int
	DarkMatterDerived        bool
	VacuumFateDerived        bool
	FullNumericalTOEClosed   bool
	FinalStatement           string
	Verdict                  string
}

type Firewall struct {
	Executed                       bool
	NoObservedOmegaFitted          bool
	NoReheatingTemperatureInserted bool
	NoCrossSectionInserted         bool
	NoDecayWidthInserted           bool
	NoTopYukawaInserted            bool
	NoGaugeTrajectoryInserted      bool
	NoThresholdSignAssumed         bool
	NoBounceMinimumInserted        bool
	NoLifetimeTargetFitted         bool
	NoClaimBeyondNativeInputs      bool
	Verdict                        string
}

type Analysis struct {
	Higgs      InheritedHiggsBoundary
	Heavy      HeavySectorLedger
	DarkMatter DarkMatterPrediction
	Vacuum     VacuumStabilityPrediction
	Census     ObservableCensus
	Firewall   Firewall
	Statuses   []string
	Truth      string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() { defaultA, defaultErr = Build() })
	return defaultA, defaultErr
}

func Build() (Analysis, error) {
	g385, err := gate385.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate 385 Higgs edge-measure theorem: %w", err)
	}
	h := inheritHiggs(g385)
	if !h.TreeProxySealed || h.LambdaEW <= 0 || h.HiggsMassProxyGeV <= 0 {
		return Analysis{}, fmt.Errorf("Gate 385 did not expose a sealed tree-level Higgs proxy")
	}
	heavy := formalizeHeavySector()
	dm := auditDarkMatter(heavy)
	vacuum := auditVacuumFate(h, heavy)
	census := buildCensus(dm, vacuum)
	fw := auditFirewall()
	statuses := []string{
		StatusGate385SealedHiggsInherited,
		StatusBGapThresholdLedgerInherited,
		StatusDarkSectorCandidateFormalized,
		StatusBoltzmannEquationFormalized,
		StatusVacuumStabilityRGFormalized,
		StatusThresholdJumpFormalized,
		StatusBounceActionFormalized,
		StatusObservableFirewallPreserved,
		StatusComputableTargetsOpened,
		StatusTensionMajoranaScaleNotStability,
		StatusTensionRelicNeedsProductionHistory,
		StatusTensionSealedHiggsNotFullRGTrajectory,
		StatusTensionTopYukawaStillFlavorSeal,
		StatusTensionThresholdSignNeedsMatching,
		StatusTensionBounceNeedsLambdaMinimum,
		StatusFailedCosmologicalObservablesNotDerived,
		StatusFailedDarkMatterRelicNotDerived,
		StatusFailedOmegaDMNotComputed,
		StatusFailedStableDarkCandidateNotDerived,
		StatusFailedBoltzmannKernelNotClosed,
		StatusFailedVacuumStabilityNotDerived,
		StatusFailedUniverseLifetimeNotDerived,
		StatusFailedEuclideanBounceNotComputed,
		StatusFailedFullNumericalTOENotClosed,
	}
	truth := "Gate 386 inherits the Gate-385 CCM+Pfaffian Higgs proxy and the B-gap threshold ledger, so the continuum cosmology questions are now mathematically well posed.  However, they are not closed native predictions.  The B-gap/Majorana scale is a candidate heavy sector, not a stable dark relic theorem; Omega_DM h^2 requires stability, rates, production history, reheating temperature, entropy evolution, and a closed Boltzmann kernel.  The sealed Higgs quartic supplies an initial boundary for RG, but vacuum fate still depends on the top Yukawa/flavor seal, gauge running, threshold matching convention, lambda minimum, and bounce prefactor.  Therefore Gate 386 opens the correct observable-calculation program but does not derive dark matter abundance, absolute stability/metastability, or universe lifetime from the finite ledger alone."
	return Analysis{Higgs: h, Heavy: heavy, DarkMatter: dm, Vacuum: vacuum, Census: census, Firewall: fw, Statuses: statuses, Truth: truth}, nil
}

func inheritHiggs(g gate385.Analysis) InheritedHiggsBoundary {
	c := g.Calculation
	return InheritedHiggsBoundary{
		Executed:          true,
		SourceGate:        385,
		LambdaEW:          c.Higgs.LambdaEdge,
		HiggsMassProxyGeV: c.Higgs.MassPfaffianGeV,
		PfaffianVEVGeV:    c.Higgs.MassPfaffianGeV / math.Sqrt(2*c.Higgs.LambdaEdge),
		TreeProxySealed:   c.HiggsTreeProxySealed,
		PoleMassDerived:   c.PhysicalPoleMassDerived,
		UsesEdgeMeasure:   c.EdgeMeasureSelected,
		DirectAnswer:      "Sealed CCM+Pfaffian tree-level proxy inherited; physical pole mass/RG matching still open.",
		Verdict:           join(StatusGate385SealedHiggsInherited, StatusTensionSealedHiggsNotFullRGTrajectory),
	}
}

func formalizeHeavySector() HeavySectorLedger {
	return HeavySectorLedger{
		Executed:                  true,
		Candidate:                 "B-gap / heavy Majorana right-handed sector ν_R ↔ ν_R^c",
		BGapScaleGeV:              bGapMajoranaScaleGeV,
		HeavyIntermediateScaleGeV: heavyIntermediateScaleGeV,
		ThresholdJumpDeltaLambda:  thresholdJumpDeltaLambda,
		GeometricallyMandated:     true,
		StableRelicTheorem:        false,
		DecayWidthDerived:         false,
		AnnihilationRateDerived:   false,
		DirectAnswer:              "The heavy sector is geometrically present, but its relic stability, decay width, and annihilation/production rates are not derived.",
		Verdict:                   join(StatusBGapThresholdLedgerInherited, StatusTensionMajoranaScaleNotStability),
	}
}

func auditDarkMatter(h HeavySectorLedger) DarkMatterPrediction {
	inputs := []RequiredInput{
		{"dark-sector stability / protecting symmetry", true, h.StableRelicTheorem, "not available", "a heavy Majorana state can decay unless a native stability theorem forbids it", true},
		{"annihilation or scattering cross section <σv>", true, h.AnnihilationRateDerived, "not available", "freeze-out abundance depends on interaction strength, not mass alone", true},
		{"decay width Γ", true, h.DecayWidthDerived, "not available", "unstable heavy states cannot be treated as present-day dark matter", true},
		{"production channel", true, false, "not available", "thermal freeze-out, freeze-in, misalignment, or nonthermal production give different abundance laws", true},
		{"reheating temperature T_R", true, false, "not available", "a 10^6 GeV state is populated only if the early thermal history reaches the relevant scale", true},
		{"effective relativistic degrees of freedom g*(T)", true, false, "not native", "needed for H(T), entropy density, and freeze-out/decoupling", true},
		{"initial abundance / entropy dilution", true, false, "not available", "Boltzmann integration requires initial conditions", true},
	}
	return DarkMatterPrediction{
		Executed:                 true,
		BoltzmannEquation:        "dn/dt + 3Hn = -<σv>(n²-n_eq²) - Γn + source(T)",
		Candidate:                h.Candidate,
		MassScaleGeV:             h.BGapScaleGeV,
		Inputs:                   inputs,
		NativeMassScaleAvailable: true,
		StableCandidateDerived:   false,
		BoltzmannKernelClosed:    false,
		OmegaH2Derived:           false,
		OmegaH2:                  math.NaN(),
		ReferenceOmegaH2:         omegaDMObservedReference,
		IndependentPrediction:    false,
		DirectAnswer:             "Dark matter abundance is not predicted.  ASHA supplies a candidate heavy sector scale, but not the stability/rate/history data required for Ω_DM h².",
		Verdict:                  join(StatusDarkSectorCandidateFormalized, StatusBoltzmannEquationFormalized, StatusFailedDarkMatterRelicNotDerived, StatusFailedOmegaDMNotComputed),
	}
}

func auditVacuumFate(h InheritedHiggsBoundary, heavy HeavySectorLedger) VacuumStabilityPrediction {
	inputs := []RGInput{
		{"top Yukawa y_t or top mass scheme", true, false, "not native after 13-moduli census", "the sign of β_λ is dominated by the -6y_t^4 term", true},
		{"gauge couplings g1,g2,g3 and normalization scheme", true, false, "partially structural ratios only", "full λ(μ) needs absolute running couplings", true},
		{"two-loop or selected one-loop beta system", true, false, "not selected here", "vacuum fate is precision-sensitive", true},
		{"threshold matching convention at B-gap", true, false, fmt.Sprintf("Δλ=%.12f available", heavy.ThresholdJumpDeltaLambda), "the sign and side of the jump must be fixed for upward vs downward transport", true},
		{"lambda minimum λ_min and scale μ_min", true, false, "not derived", "bounce action requires the deepest negative value if λ turns negative", true},
		{"bounce prefactor and gravitational correction", true, false, "not derived", "lifetime is not only S_E; prefactor and spacetime volume enter", true},
	}
	return VacuumStabilityPrediction{
		Executed:                 true,
		LambdaEW:                 h.LambdaEW,
		HiggsMassProxyGeV:        h.HiggsMassProxyGeV,
		ThresholdScaleGeV:        heavy.BGapScaleGeV,
		ThresholdDeltaLambda:     heavy.ThresholdJumpDeltaLambda,
		OneLoopBetaLambda:        "β_λ=(16π²)^-1[24λ² -6y_t^4 + (9/8)g2^4 + (3/4)g2²g1² + (3/8)g1^4 + (-9g2²-3g1²+12y_t²)λ] + threshold terms",
		Inputs:                   inputs,
		FullRGTrajectoryDerived:  false,
		ThresholdMatchingDerived: false,
		LambdaMinimumDerived:     false,
		AbsoluteStabilityDerived: false,
		MetastabilityDerived:     false,
		EuclideanBounceFormula:   "S_E ≈ 8π²/(3|λ_min|) when λ_min<0, plus prefactor and gravitational corrections",
		EuclideanBounceAction:    math.NaN(),
		LifetimeYears:            math.NaN(),
		LifetimeDerived:          false,
		IndependentPrediction:    false,
		DirectAnswer:             "Vacuum fate is not predicted.  The sealed Higgs proxy gives λ(v), but y_t, gauge running, matching, λ_min, and bounce data are still open.",
		Verdict:                  join(StatusVacuumStabilityRGFormalized, StatusThresholdJumpFormalized, StatusBounceActionFormalized, StatusFailedVacuumStabilityNotDerived, StatusFailedUniverseLifetimeNotDerived),
	}
}

func buildCensus(dm DarkMatterPrediction, v VacuumStabilityPrediction) ObservableCensus {
	hard := 0
	if dm.IndependentPrediction {
		hard++
	}
	if v.IndependentPrediction {
		hard++
	}
	return ObservableCensus{
		Executed:                 true,
		RequestedObservables:     2,
		HardPredictionsDerived:   hard,
		ConditionalTargetsOpened: 2,
		DarkMatterDerived:        dm.OmegaH2Derived,
		VacuumFateDerived:        v.LifetimeDerived || v.AbsoluteStabilityDerived || v.MetastabilityDerived,
		FullNumericalTOEClosed:   false,
		FinalStatement:           "Gate 386 turns dark matter and vacuum fate into well-defined conditional calculations, but derives zero hard macroscopic observables from native data alone.",
		Verdict:                  join(StatusComputableTargetsOpened, StatusFailedCosmologicalObservablesNotDerived, StatusFailedFullNumericalTOENotClosed),
	}
}

func auditFirewall() Firewall {
	return Firewall{
		Executed:                       true,
		NoObservedOmegaFitted:          true,
		NoReheatingTemperatureInserted: true,
		NoCrossSectionInserted:         true,
		NoDecayWidthInserted:           true,
		NoTopYukawaInserted:            true,
		NoGaugeTrajectoryInserted:      true,
		NoThresholdSignAssumed:         true,
		NoBounceMinimumInserted:        true,
		NoLifetimeTargetFitted:         true,
		NoClaimBeyondNativeInputs:      true,
		Verdict:                        join(StatusObservableFirewallPreserved, StatusFailedCosmologicalObservablesNotDerived),
	}
}

func StatusLine(a Analysis) string { return strings.Join(a.Statuses, "\n") }

func NativeCosmologyConstants() map[string]float64 {
	a, err := BuildDefault()
	if err != nil {
		return map[string]float64{}
	}
	return map[string]float64{
		"lambda_EW_edge":             a.Higgs.LambdaEW,
		"higgs_mass_proxy_GeV":       a.Higgs.HiggsMassProxyGeV,
		"pfaffian_vev_GeV":           a.Higgs.PfaffianVEVGeV,
		"b_gap_scale_GeV":            a.Heavy.BGapScaleGeV,
		"heavy_intermediate_GeV":     a.Heavy.HeavyIntermediateScaleGeV,
		"delta_lambda_threshold":     a.Heavy.ThresholdJumpDeltaLambda,
		"omega_dm_h2_derived":        boolFloat(a.DarkMatter.OmegaH2Derived),
		"vacuum_lifetime_derived":    boolFloat(a.Vacuum.LifetimeDerived),
		"hard_predictions_derived":   float64(a.Census.HardPredictionsDerived),
		"conditional_targets_opened": float64(a.Census.ConditionalTargetsOpened),
	}
}

func join(parts ...string) string { return strings.Join(parts, ";") }
func boolFloat(v bool) float64 {
	if v {
		return 1
	}
	return 0
}
