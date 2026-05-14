// Package cosmologicalobservables implements Gate 375:
// Cosmological Observables & Dark Sector Prediction Sieve.
//
// Gate 375 deliberately breaks the Gate-374 scoped closure in the only honest
// direction left: observable cosmology.  It asks whether the already-derived
// ASHA boundary/threshold ledger is sufficient to compute hard numbers for
// dark-matter abundance, vacuum lifetime, and the cosmological constant.
//
// The answer is an audit, not a fit.  A cosmological observable is promoted only
// if the ledger fixes every physical ingredient of the relevant continuum
// model: particle identity, stability/decay channel, coupling, initial/reheating
// state, Boltzmann kernel, RG trajectory, bounce action, counterterm, and
// normalization.  Where those ingredients are missing, the gate records a
// failed route rather than importing observed cosmology.
package cosmologicalobservables

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate374 "github.com/bagherbal/asha-engine/pkg/bridge/asha_final_closing_theorem"
)

const (
	AuditID = "GATE375-COSMOLOGICAL-OBSERVABLES-DARK-SECTOR-PREDICTION-SIEVE"

	StatusGate374Inherited                    = "CONDITIONAL_SUPPORT_GATE374_SCOPED_CLOSURE_INHERITED"
	StatusCosmologicalObservableSieveOpened   = "CONDITIONAL_SUPPORT_COSMOLOGICAL_OBSERVABLE_SIEVE_OPENED"
	StatusASHAHeavyScaleLedgerFormalized      = "CONDITIONAL_SUPPORT_ASHA_HEAVY_SCALE_LEDGER_FORMALIZED"
	StatusDarkMatterBoltzmannSystemFormalized = "CONDITIONAL_SUPPORT_DARK_MATTER_BOLTZMANN_SYSTEM_FORMALIZED"
	StatusRelicDensityInputAuditExecuted      = "CONDITIONAL_SUPPORT_RELIC_DENSITY_INPUT_AUDIT_EXECUTED"
	StatusVacuumLifetimeFormalized            = "CONDITIONAL_SUPPORT_VACUUM_LIFETIME_BOUNCE_FORMALIZED"
	StatusMetastabilityInputAuditExecuted     = "CONDITIONAL_SUPPORT_METASTABILITY_INPUT_AUDIT_EXECUTED"
	StatusCosmologicalConstantAuditExecuted   = "CONDITIONAL_SUPPORT_COSMOLOGICAL_CONSTANT_CAPACITY_AUDITED"
	StatusObservableFirewallPreserved         = "CONDITIONAL_SUPPORT_COSMOLOGICAL_OBSERVABLE_FIREWALL_PRESERVED"
	StatusNoCosmologicalPredictionDerived     = "CONDITIONAL_SUPPORT_NO_HARD_COSMOLOGICAL_OBSERVABLE_DERIVED_IN_CURRENT_LEDGER"

	StatusTensionHeavyScaleNotRelicModel              = "CONDITIONAL_TENSION_HEAVY_SCALE_IS_NOT_A_RELIC_MODEL"
	StatusTensionBGapThresholdConditional             = "CONDITIONAL_TENSION_BGAP_THRESHOLD_IS_CONDITIONAL_TRANSPORT_LEDGER"
	StatusTensionBoltzmannNeedsRatesAndInitialState   = "CONDITIONAL_TENSION_BOLTZMANN_RELIC_NEEDS_RATES_AND_INITIAL_STATE"
	StatusTensionMajoranaStabilityNotDerived          = "CONDITIONAL_TENSION_MAJORANA_DARK_STABILITY_NOT_DERIVED"
	StatusTensionVacuumLifetimeNeedsFullRGTrajectory  = "CONDITIONAL_TENSION_VACUUM_LIFETIME_NEEDS_FULL_RG_TRAJECTORY"
	StatusTensionBounceNeedsLambdaMinimumAndPrefactor = "CONDITIONAL_TENSION_BOUNCE_ACTION_NEEDS_LAMBDA_MINIMUM_AND_PREFACTOR"
	StatusTensionDarkEnergyNeedsCounterterm           = "CONDITIONAL_TENSION_DARK_ENERGY_NEEDS_RENORMALIZED_COUNTERTERM"
	StatusTensionHierarchyNotCosmologicalConstant     = "CONDITIONAL_TENSION_PFAFFIAN_HIERARCHY_IS_NOT_A_NATIVE_LAMBDA_COSMO_PREDICTION"

	StatusFailedCosmologicalObservablesNotDerived = "FAILED_ROUTE_COSMOLOGICAL_OBSERVABLES_NOT_DERIVED"
	StatusFailedRelicDensityNotDerived            = "FAILED_ROUTE_DARK_MATTER_RELIC_DENSITY_NOT_DERIVED"
	StatusFailedOmegaDMNotComputed                = "FAILED_ROUTE_OMEGA_DM_H2_NOT_COMPUTED_FROM_ASHA_LEDGER"
	StatusFailedStableDarkCandidateNotDerived     = "FAILED_ROUTE_STABLE_DARK_SECTOR_CANDIDATE_NOT_DERIVED"
	StatusFailedBoltzmannKernelNotDerived         = "FAILED_ROUTE_BOLTZMANN_KERNEL_NOT_DERIVED"
	StatusFailedVacuumLifetimeNotDerived          = "FAILED_ROUTE_UNIVERSE_LIFETIME_NOT_DERIVED"
	StatusFailedBounceActionNotComputed           = "FAILED_ROUTE_EUCLIDEAN_BOUNCE_ACTION_NOT_COMPUTED"
	StatusFailedCosmologicalConstantNotDerived    = "FAILED_ROUTE_COSMOLOGICAL_CONSTANT_NOT_DERIVED"
	StatusFailedDarkEnergyDensityNotComputed      = "FAILED_ROUTE_DARK_ENERGY_DENSITY_NOT_COMPUTED"
)

const (
	lambdaHOverGStarSquared  = 1197.0 / 4624.0
	alphaGUTInverseBranch    = 8.0 * math.Pi
	sin2ThetaWBoundary       = 3.0 / 8.0
	vevOverPlanckHierarchy   = 2.024161131083123e-17
	peVThresholdGeV          = 1.46774973718e6
	sealedIntermediateGeV    = 6.650726476871e11
	requiredQuarticJump      = -0.0978
	chargedFiniteDiracModuli = 13
	externalMinimalLedger    = 15
)

type Inheritance struct {
	Executed               bool
	HighestInheritedGate   int
	ScopedClosureInherited bool
	FiniteKinematicsClosed bool
	FlavorVacuumUnselected bool
	ChargedModuli          int
	ExternalLedger         int
	PreviousTruth          string
	Question               string
	Verdict                string
}

type ScaleLedger struct {
	Executed                      bool
	Sin2ThetaWBoundary            float64
	AlphaGUTInverseBranch         float64
	LambdaHOverGStarSquared       float64
	VEVOverPlanckHierarchy        float64
	PeVThresholdGeV               float64
	SealedIntermediateGeV         float64
	RequiredQuarticJump           float64
	PeVThresholdDerivedAsMass     bool
	IntermediateScaleRelicTheorem bool
	HeavyDarkStabilityDerived     bool
	DirectAnswer                  string
	Verdict                       string
}

type RelicInput struct {
	Name        string
	Required    bool
	Native      bool
	Value       string
	MissingWhy  string
	BlocksOmega bool
}

type DarkMatterAudit struct {
	Executed                    bool
	Candidate                   string
	BoltzmannSystem             string
	MassScaleGeV                float64
	AlternativeScaleGeV         float64
	Inputs                      []RelicInput
	NativeMassScaleAvailable    bool
	StabilityDerived            bool
	AnnihilationCrossSection    bool
	DecayWidthDerived           bool
	ReheatingTemperatureDerived bool
	InitialAbundanceDerived     bool
	EntropyDilutionDerived      bool
	BoltzmannKernelClosed       bool
	OmegaH2Derived              bool
	OmegaH2Prediction           float64
	IndependentPredictions      int
	DirectAnswer                string
	Verdict                     string
}

type MetastabilityInput struct {
	Name           string
	Required       bool
	Native         bool
	Value          string
	MissingWhy     string
	BlocksLifetime bool
}

type VacuumLifetimeAudit struct {
	Executed                     bool
	BoundaryLambdaRatioAvailable bool
	ThresholdJumpAvailable       bool
	BounceFormula                string
	Inputs                       []MetastabilityInput
	FullRGTrajectoryDerived      bool
	LambdaMinimumDerived         bool
	NegativeLambdaRegionDerived  bool
	EuclideanBounceActionDerived bool
	PrefactorDerived             bool
	LifetimeYearsDerived         bool
	EuclideanBounceAction        float64
	LifetimeYears                float64
	IndependentPredictions       int
	DirectAnswer                 string
	Verdict                      string
}

type DarkEnergyAudit struct {
	Executed                     bool
	HierarchyAvailable           bool
	NativeSuppressionPowers      []SuppressionPower
	VacuumCountertermDerived     bool
	RenormalizedVacuumFunctional bool
	HolographicSaturationDerived bool
	DarkEnergyDensityDerived     bool
	CosmologicalConstantDerived  bool
	IndependentPredictions       int
	DirectAnswer                 string
	Verdict                      string
}

type SuppressionPower struct {
	Expression string
	Value      float64
	Meaning    string
	Derived    bool
	PredictsCC bool
}

type ObservableCensus struct {
	Executed                        bool
	RequestedObservables            int
	HardPredictionsDerived          int
	DarkMatterPredictions           int
	VacuumLifetimePredictions       int
	CosmologicalConstantPredictions int
	RemainingChargedModuli          int
	FinalStatement                  string
	Verdict                         string
}

type Firewall struct {
	Executed                         bool
	NoObservedOmegaDMFitted          bool
	NoObservedDarkEnergyFitted       bool
	NoObservedLifetimeTargetFitted   bool
	NoReheatingTemperatureInserted   bool
	NoAnnihilationCrossSectionFitted bool
	NoDecayWidthFitted               bool
	NoRGTrajectoryFitted             bool
	NoLambdaMinimumInserted          bool
	NoVacuumCountertermInserted      bool
	NoHolographicSaturationAssumed   bool
	NoClaimBeyondInputs              bool
	Verdict                          string
}

type Analysis struct {
	Inheritance Inheritance
	Scales      ScaleLedger
	DarkMatter  DarkMatterAudit
	Lifetime    VacuumLifetimeAudit
	DarkEnergy  DarkEnergyAudit
	Census      ObservableCensus
	Firewall    Firewall
	Truth       string
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
	prev, err := gate374.BuildDefault()
	if err != nil {
		return Analysis{}, err
	}
	inheritance := inherit(prev)
	scales := formalizeScaleLedger()
	dm := auditDarkMatter(scales)
	life := auditVacuumLifetime(scales)
	de := auditDarkEnergy(scales)
	census := updateObservableCensus(dm, life, de)
	fw := auditFirewall()
	truth := "Gate 375 breaks the Gate-374 scoped closure only as an observable-prediction audit.  The ASHA ledger contains powerful boundary and threshold data, but those data are not yet a closed cosmological model.  A dark-matter relic abundance requires a stable dark candidate, interaction rates, production history, reheating temperature, entropy dilution, and a closed Boltzmann kernel; the current B-gap/heavy-sector ledger supplies candidate scales and overlap/threshold structure but not those inputs.  Vacuum lifetime requires the full continuum RG trajectory of λ(μ), the negative λ minimum, a bounce prefactor, and matching/gravity corrections; the boundary ratio and threshold jump do not determine a lifetime by themselves.  The Pfaffian hierarchy is a scale relation, not a renormalized cosmological-constant counterterm or holographic saturation theorem.  Therefore no hard predictions for Ω_DM h², universe lifetime, or dark-energy density are derived in the current ledger.  The correct next program is a continuum cosmology extension deriving the dark-sector Lagrangian, Boltzmann kernel, and RG/bounce functional, not fitting observed cosmology into the finite ledger."
	return Analysis{inheritance, scales, dm, life, de, census, fw, truth}, nil
}

func inherit(prev gate374.Analysis) Inheritance {
	return Inheritance{
		Executed:               true,
		HighestInheritedGate:   374,
		ScopedClosureInherited: true,
		FiniteKinematicsClosed: prev.Closing.KinematicsComplete,
		FlavorVacuumUnselected: prev.Closing.DynamicsOfFlavorUnselected,
		ChargedModuli:          prev.Moduli.ChargedFlavorModuli,
		ExternalLedger:         prev.Moduli.ExternalLedger,
		PreviousTruth:          prev.Truth,
		Question:               "do the sealed ASHA boundary/threshold data produce hard cosmological observables without new fitted continuum inputs?",
		Verdict:                join(StatusGate374Inherited, StatusCosmologicalObservableSieveOpened),
	}
}

func formalizeScaleLedger() ScaleLedger {
	return ScaleLedger{
		Executed:                      true,
		Sin2ThetaWBoundary:            sin2ThetaWBoundary,
		AlphaGUTInverseBranch:         alphaGUTInverseBranch,
		LambdaHOverGStarSquared:       lambdaHOverGStarSquared,
		VEVOverPlanckHierarchy:        vevOverPlanckHierarchy,
		PeVThresholdGeV:               peVThresholdGeV,
		SealedIntermediateGeV:         sealedIntermediateGeV,
		RequiredQuarticJump:           requiredQuarticJump,
		PeVThresholdDerivedAsMass:     false,
		IntermediateScaleRelicTheorem: false,
		HeavyDarkStabilityDerived:     false,
		DirectAnswer:                  "ASHA supplies boundary ratios, hierarchy, a conditional PeV threshold lane, and a sealed intermediate-scale ledger; these are not yet a dark-sector mass/coupling/stability theorem.",
		Verdict:                       join(StatusASHAHeavyScaleLedgerFormalized, StatusTensionBGapThresholdConditional, StatusTensionHeavyScaleNotRelicModel),
	}
}

func auditDarkMatter(scales ScaleLedger) DarkMatterAudit {
	inputs := []RelicInput{
		{"dark particle identity", true, false, "candidate: heavy Majorana/B-gap sector", "candidate semantics exist, but a stable cosmological relic species is not derived", true},
		{"physical mass", true, false, fmt.Sprintf("PeV lane %.12e GeV; sealed M_int %.12e GeV", scales.PeVThresholdGeV, scales.SealedIntermediateGeV), "scale ledgers are conditional/transport/resonance data, not a relic-particle mass theorem", true},
		{"stability or lifetime", true, false, "not fixed", "no symmetry or decay-width theorem proves survival to cosmological times", true},
		{"annihilation/scattering cross section <σv>", true, false, "not fixed", "requires dark-sector interactions and continuum matrix elements", true},
		{"decay width Γ", true, false, "not fixed", "requires couplings, available final states, and mixing angles", true},
		{"reheating temperature / production history", true, false, "not fixed", "thermal vs freeze-in vs nonthermal production is not selected", true},
		{"entropy dilution and effective g_*", true, false, "not fixed", "requires cosmological thermal history", true},
	}
	return DarkMatterAudit{
		Executed:                    true,
		Candidate:                   "Heavy Majorana / B-gap sector as a possible dark relic channel",
		BoltzmannSystem:             "dY/dx = -s <σv>/(Hx) (Y²-Y_eq²) + source/decay terms; Ωh² ∝ m Y∞",
		MassScaleGeV:                scales.PeVThresholdGeV,
		AlternativeScaleGeV:         scales.SealedIntermediateGeV,
		Inputs:                      inputs,
		NativeMassScaleAvailable:    false,
		StabilityDerived:            false,
		AnnihilationCrossSection:    false,
		DecayWidthDerived:           false,
		ReheatingTemperatureDerived: false,
		InitialAbundanceDerived:     false,
		EntropyDilutionDerived:      false,
		BoltzmannKernelClosed:       false,
		OmegaH2Derived:              false,
		OmegaH2Prediction:           math.NaN(),
		IndependentPredictions:      0,
		DirectAnswer:                "No Ω_DM h² prediction follows from the current ASHA ledger: a mass scale alone is not a relic-density model.",
		Verdict:                     join(StatusDarkMatterBoltzmannSystemFormalized, StatusRelicDensityInputAuditExecuted, StatusTensionBoltzmannNeedsRatesAndInitialState, StatusTensionMajoranaStabilityNotDerived, StatusFailedRelicDensityNotDerived, StatusFailedOmegaDMNotComputed, StatusFailedStableDarkCandidateNotDerived, StatusFailedBoltzmannKernelNotDerived),
	}
}

func auditVacuumLifetime(scales ScaleLedger) VacuumLifetimeAudit {
	inputs := []MetastabilityInput{
		{"UV quartic boundary", true, true, "λ_H/g_*² = 1197/4624", "available as a boundary ratio, but not a full low-energy λ(μ) trajectory", false},
		{"threshold jump", true, true, fmt.Sprintf("Δλ ≈ %.4f", scales.RequiredQuarticJump), "available as a transport target/witness, not a complete heavy-sector matching theorem", false},
		{"full βλ, βyt, βg trajectory", true, false, "not closed", "requires continuum RG system and thresholds with calibrated boundary conditions", true},
		{"top Yukawa and gauge coupling transport", true, false, "not fixed", "the 13 moduli include the top/flavor sector; Gate 372 says it is not selected", true},
		{"λ_min and instability scale", true, false, "not fixed", "bounce action depends on the most negative running quartic", true},
		{"bounce prefactor and gravitational corrections", true, false, "not fixed", "lifetime requires determinant prefactor and spacetime-volume convention", true},
	}
	return VacuumLifetimeAudit{
		Executed:                     true,
		BoundaryLambdaRatioAvailable: true,
		ThresholdJumpAvailable:       true,
		BounceFormula:                "S_E ≈ 8π²/(3 |λ_min|) for the thin/scale-invariant quartic approximation; Γ/V ≈ μ_B⁴ exp(-S_E)",
		Inputs:                       inputs,
		FullRGTrajectoryDerived:      false,
		LambdaMinimumDerived:         false,
		NegativeLambdaRegionDerived:  false,
		EuclideanBounceActionDerived: false,
		PrefactorDerived:             false,
		LifetimeYearsDerived:         false,
		EuclideanBounceAction:        math.NaN(),
		LifetimeYears:                math.NaN(),
		IndependentPredictions:       0,
		DirectAnswer:                 "No universe lifetime is derived: the ASHA boundary/jump data do not determine λ_min, bounce scale, or prefactor.",
		Verdict:                      join(StatusVacuumLifetimeFormalized, StatusMetastabilityInputAuditExecuted, StatusTensionVacuumLifetimeNeedsFullRGTrajectory, StatusTensionBounceNeedsLambdaMinimumAndPrefactor, StatusFailedVacuumLifetimeNotDerived, StatusFailedBounceActionNotComputed),
	}
}

func auditDarkEnergy(scales ScaleLedger) DarkEnergyAudit {
	powers := []SuppressionPower{
		{"(v/M_P)^2", scales.VEVOverPlanckHierarchy * scales.VEVOverPlanckHierarchy, "native hierarchy squared; not a vacuum-energy theorem", true, false},
		{"(v/M_P)^4", math.Pow(scales.VEVOverPlanckHierarchy, 4), "electroweak vacuum-energy dimensional scaling; still requires counterterm/sign/normalization", true, false},
		{"exp(-8π²)", math.Exp(-8.0 * math.Pi * math.Pi), "instanton-like suppression related to hierarchy square; not uniquely Λ_cosmo", true, false},
	}
	return DarkEnergyAudit{
		Executed:                     true,
		HierarchyAvailable:           true,
		NativeSuppressionPowers:      powers,
		VacuumCountertermDerived:     false,
		RenormalizedVacuumFunctional: false,
		HolographicSaturationDerived: false,
		DarkEnergyDensityDerived:     false,
		CosmologicalConstantDerived:  false,
		IndependentPredictions:       0,
		DirectAnswer:                 "No cosmological constant prediction is derived: hierarchy suppressions exist, but no renormalized vacuum counterterm or saturation equation selects dark energy.",
		Verdict:                      join(StatusCosmologicalConstantAuditExecuted, StatusTensionDarkEnergyNeedsCounterterm, StatusTensionHierarchyNotCosmologicalConstant, StatusFailedCosmologicalConstantNotDerived, StatusFailedDarkEnergyDensityNotComputed),
	}
}

func updateObservableCensus(dm DarkMatterAudit, life VacuumLifetimeAudit, de DarkEnergyAudit) ObservableCensus {
	dmN, lifeN, deN := dm.IndependentPredictions, life.IndependentPredictions, de.IndependentPredictions
	total := dmN + lifeN + deN
	return ObservableCensus{
		Executed:                        true,
		RequestedObservables:            3,
		HardPredictionsDerived:          total,
		DarkMatterPredictions:           dmN,
		VacuumLifetimePredictions:       lifeN,
		CosmologicalConstantPredictions: deN,
		RemainingChargedModuli:          chargedFiniteDiracModuli,
		FinalStatement:                  "0 hard cosmological observables are derived; the 13 charged finite-Dirac flavor moduli remain unselected.",
		Verdict:                         join(StatusNoCosmologicalPredictionDerived, StatusFailedCosmologicalObservablesNotDerived),
	}
}

func auditFirewall() Firewall {
	return Firewall{
		Executed:                         true,
		NoObservedOmegaDMFitted:          true,
		NoObservedDarkEnergyFitted:       true,
		NoObservedLifetimeTargetFitted:   true,
		NoReheatingTemperatureInserted:   true,
		NoAnnihilationCrossSectionFitted: true,
		NoDecayWidthFitted:               true,
		NoRGTrajectoryFitted:             true,
		NoLambdaMinimumInserted:          true,
		NoVacuumCountertermInserted:      true,
		NoHolographicSaturationAssumed:   true,
		NoClaimBeyondInputs:              true,
		Verdict:                          StatusObservableFirewallPreserved,
	}
}

func NativeCosmologicalScales() map[string]float64 {
	return map[string]float64{
		"sin2_thetaW_boundary":     sin2ThetaWBoundary,
		"alpha_GUT_inverse_branch": alphaGUTInverseBranch,
		"lambdaH_over_gstar2":      lambdaHOverGStarSquared,
		"v_over_MP_hierarchy":      vevOverPlanckHierarchy,
		"PeV_threshold_GeV":        peVThresholdGeV,
		"sealed_intermediate_GeV":  sealedIntermediateGeV,
		"required_quartic_jump":    requiredQuarticJump,
		"hierarchy_squared":        vevOverPlanckHierarchy * vevOverPlanckHierarchy,
		"hierarchy_fourth_power":   math.Pow(vevOverPlanckHierarchy, 4),
	}
}

func FormatInheritance(i Inheritance) string {
	return fmt.Sprintf("executed=%t inheritedGate=%d scoped=%t kinematicsClosed=%t flavorUnselected=%t moduli=%d external=%d verdict=%s", i.Executed, i.HighestInheritedGate, i.ScopedClosureInherited, i.FiniteKinematicsClosed, i.FlavorVacuumUnselected, i.ChargedModuli, i.ExternalLedger, i.Verdict)
}

func FormatScales(s ScaleLedger) string {
	return fmt.Sprintf("sin2=%.12g alphaInv=%.12g lambdaRatio=%.12g hierarchy=%.12e PeV=%.12e M_int=%.12e deltaLambda=%.6g PeVMass=%t relicTheorem=%t stability=%t verdict=%s", s.Sin2ThetaWBoundary, s.AlphaGUTInverseBranch, s.LambdaHOverGStarSquared, s.VEVOverPlanckHierarchy, s.PeVThresholdGeV, s.SealedIntermediateGeV, s.RequiredQuarticJump, s.PeVThresholdDerivedAsMass, s.IntermediateScaleRelicTheorem, s.HeavyDarkStabilityDerived, s.Verdict)
}

func FormatRelicInputs(inputs []RelicInput) string {
	parts := make([]string, 0, len(inputs))
	for _, in := range inputs {
		parts = append(parts, fmt.Sprintf("%s[native=%t value=%s blocks=%t]", in.Name, in.Native, in.Value, in.BlocksOmega))
	}
	return strings.Join(parts, " | ")
}

func FormatDarkMatter(d DarkMatterAudit) string {
	return fmt.Sprintf("candidate=%q mass=%.12e alt=%.12e stable=%t sigma=%t decay=%t reheating=%t kernel=%t omegaDerived=%t omega=%v predictions=%d inputs={%s} verdict=%s", d.Candidate, d.MassScaleGeV, d.AlternativeScaleGeV, d.StabilityDerived, d.AnnihilationCrossSection, d.DecayWidthDerived, d.ReheatingTemperatureDerived, d.BoltzmannKernelClosed, d.OmegaH2Derived, d.OmegaH2Prediction, d.IndependentPredictions, FormatRelicInputs(d.Inputs), d.Verdict)
}

func FormatLifetimeInputs(inputs []MetastabilityInput) string {
	parts := make([]string, 0, len(inputs))
	for _, in := range inputs {
		parts = append(parts, fmt.Sprintf("%s[native=%t value=%s blocks=%t]", in.Name, in.Native, in.Value, in.BlocksLifetime))
	}
	return strings.Join(parts, " | ")
}

func FormatLifetime(l VacuumLifetimeAudit) string {
	return fmt.Sprintf("boundary=%t jump=%t fullRG=%t lambdaMin=%t negativeRegion=%t bounce=%t prefactor=%t lifetime=%t SE=%v years=%v predictions=%d formula=%q inputs={%s} verdict=%s", l.BoundaryLambdaRatioAvailable, l.ThresholdJumpAvailable, l.FullRGTrajectoryDerived, l.LambdaMinimumDerived, l.NegativeLambdaRegionDerived, l.EuclideanBounceActionDerived, l.PrefactorDerived, l.LifetimeYearsDerived, l.EuclideanBounceAction, l.LifetimeYears, l.IndependentPredictions, l.BounceFormula, FormatLifetimeInputs(l.Inputs), l.Verdict)
}

func FormatDarkEnergy(d DarkEnergyAudit) string {
	parts := make([]string, 0, len(d.NativeSuppressionPowers))
	for _, p := range d.NativeSuppressionPowers {
		parts = append(parts, fmt.Sprintf("%s=%.12e predictsCC=%t", p.Expression, p.Value, p.PredictsCC))
	}
	return fmt.Sprintf("hierarchy=%t counterterm=%t functional=%t saturation=%t rhoDE=%t Lambda=%t predictions=%d powers={%s} verdict=%s", d.HierarchyAvailable, d.VacuumCountertermDerived, d.RenormalizedVacuumFunctional, d.HolographicSaturationDerived, d.DarkEnergyDensityDerived, d.CosmologicalConstantDerived, d.IndependentPredictions, strings.Join(parts, " | "), d.Verdict)
}

func FormatCensus(c ObservableCensus) string {
	return fmt.Sprintf("requested=%d hardPredictions=%d dm=%d lifetime=%d cc=%d remainingModuli=%d final=%q verdict=%s", c.RequestedObservables, c.HardPredictionsDerived, c.DarkMatterPredictions, c.VacuumLifetimePredictions, c.CosmologicalConstantPredictions, c.RemainingChargedModuli, c.FinalStatement, c.Verdict)
}

func FormatFirewall(f Firewall) string {
	return fmt.Sprintf("omegaFit=%t deFit=%t lifetimeFit=%t reheatingInserted=%t sigmaFit=%t decayFit=%t rgFit=%t lambdaMin=%t counterterm=%t saturation=%t noOverclaim=%t verdict=%s", !f.NoObservedOmegaDMFitted, !f.NoObservedDarkEnergyFitted, !f.NoObservedLifetimeTargetFitted, !f.NoReheatingTemperatureInserted, !f.NoAnnihilationCrossSectionFitted, !f.NoDecayWidthFitted, !f.NoRGTrajectoryFitted, !f.NoLambdaMinimumInserted, !f.NoVacuumCountertermInserted, !f.NoHolographicSaturationAssumed, f.NoClaimBeyondInputs, f.Verdict)
}

func join(statuses ...string) string { return strings.Join(statuses, "; ") }
