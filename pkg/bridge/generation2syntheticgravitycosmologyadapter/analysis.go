// Package generation2syntheticgravitycosmologyadapter implements Gate 515:
// Bridge-Only Gravity/Cosmology Adapter Dry-Run.
//
// Gate 514 defined the fail-closed airlock for spectral cutoff, moment,
// Planck/Newton, cosmological, vacuum-subtraction, and renormalization rows.
// Gate 515 deliberately uses fake bridge numbers to test formula plumbing and
// residual reporting. The numerical values in this package are synthetic test
// fixtures, not observed gravity/cosmology data and not native ASHA results.
package generation2syntheticgravitycosmologyadapter

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2spectralcutoffrenormalizationairlock"
)

const (
	AuditID = "GATE515-BRIDGE-ONLY-GRAVITY-COSMOLOGY-ADAPTER-DRY-RUN"

	StatusGate514AirlockInherited                 = "CONDITIONAL_SUPPORT_GATE514_CUTOFF_RENORMALIZATION_AIRLOCK_INHERITED"
	StatusSyntheticAdapterExecuted                = "CONDITIONAL_SUPPORT_SYNTHETIC_GRAVITY_COSMOLOGY_ADAPTER_EXECUTED"
	StatusSyntheticInputsFake                     = "CONDITIONAL_SUPPORT_SYNTHETIC_CUTOFF_MOMENT_INPUTS_EXPLICITLY_FAKE"
	StatusA2EHComputedBridgeOnly                  = "CONDITIONAL_SUPPORT_A2_EINSTEIN_HILBERT_COEFFICIENT_COMPUTED_BRIDGE_ONLY"
	StatusA0CosmologicalComputedBridgeOnly        = "CONDITIONAL_SUPPORT_A0_COSMOLOGICAL_VOLUME_COEFFICIENT_COMPUTED_BRIDGE_ONLY"
	StatusA4CurvatureComputedBridgeOnly           = "CONDITIONAL_SUPPORT_A4_CURVATURE_SQUARED_COEFFICIENT_COMPUTED_BRIDGE_ONLY"
	StatusVacuumSubtractionResidualPlumbingTested = "CONDITIONAL_SUPPORT_VACUUM_SUBTRACTION_RESIDUAL_PLUMBING_TESTED"
	StatusSyntheticResidualsComputed              = "CONDITIONAL_SUPPORT_COMPARATOR_RESIDUALS_COMPUTED_SYNTHETICALLY"
	StatusNoObservedGravityCosmologyDataImported  = "CONDITIONAL_SUPPORT_NO_OBSERVED_GRAVITY_COSMOLOGY_DATA_IMPORTED"
	StatusSyntheticAdapterBridgeOnly              = "CONDITIONAL_SUPPORT_SYNTHETIC_ADAPTER_BRIDGE_ONLY"

	StatusFailedSyntheticOutputsNotNative               = "FAILED_ROUTE_SYNTHETIC_GRAVITY_OUTPUTS_ARE_NOT_NATIVE_PREDICTIONS"
	StatusFailedSyntheticOutputsNotNewtonOrCosmological = "FAILED_ROUTE_SYNTHETIC_OUTPUTS_ARE_NOT_NEWTON_OR_COSMOLOGICAL_CONSTANT"
	StatusFailedLambdaF2F4StillNotDerived               = "FAILED_ROUTE_LAMBDA_F2_F4_STILL_NOT_DERIVED"
	StatusFailedVacuumSubtractionStillNotNative         = "FAILED_ROUTE_VACUUM_SUBTRACTION_STILL_NOT_NATIVE"
	StatusFailedNativeNormalizationStillBlocked         = "FAILED_ROUTE_NATIVE_GRAVITY_COSMOLOGY_NORMALIZATION_STILL_BLOCKED"
	StatusFailedObservedComparatorNotUsed               = "FAILED_ROUTE_OBSERVED_GRAVITY_COSMOLOGY_COMPARATOR_NOT_USED"
	StatusFirewallNoObservedImport                      = "FIREWALL_PRESERVED_NO_NEWTON_PLANCK_COSMOLOGY_OR_DARK_ENERGY_DATA_IMPORTED"
	StatusFirewallSyntheticNativeWriteBlocked           = "FIREWALL_BLOCKED_SYNTHETIC_GRAVITY_COSMOLOGY_OUTPUT_NATIVE_WRITE"
)

const (
	finiteTraceDimension = 96.0
	a0Prefactor          = 6.0 / (math.Pi * math.Pi)       // TrF/(16π²)
	a2Prefactor          = 1.0 / (2.0 * math.Pi * math.Pi) // TrF/(192π²)
	a4Prefactor          = 1.0 / (60.0 * math.Pi * math.Pi)
)

type Inheritance struct {
	Executed                   bool
	Gate514Inherited           bool
	RedactedSchemaAccepted     bool
	RequiredRows               int
	AcceptedCases              int
	RejectedCases              int
	Gate514NoAdapterExecuted   bool
	Gate514NativeWriteBlocked  bool
	Gate514LambdaSelected      bool
	Gate514F2Selected          bool
	Gate514F4Selected          bool
	Gate514NewtonDerived       bool
	Gate514CosmologicalDerived bool
	Verdict                    string
	Reason                     string
}

type SyntheticInputs struct {
	Executed                  bool
	LambdaCutoff              float64
	F2Moment                  float64
	F4Moment                  float64
	F0Moment                  float64
	VacuumSubtraction         float64
	EHComparator              float64
	CosmologicalComparator    float64
	AllInputsSynthetic        bool
	AllRowsBridgeOnly         bool
	AllNativePromotionBlocked bool
	ObservedDataImported      bool
	Verdict                   string
	Reason                    string
}

type AdapterOutput struct {
	Executed                      bool
	FiniteTraceDimension          float64
	A2PrefactorPerF2LambdaSquared float64
	A0PrefactorPerF4LambdaFourth  float64
	A4PrefactorPerF0              float64
	F2LambdaSquared               float64
	F4LambdaFourth                float64
	EinsteinHilbertCoefficient    float64
	CosmologicalVolumeRaw         float64
	CosmologicalAfterSubtraction  float64
	CurvatureSquaredCoefficient   float64
	NativeGravityPrediction       bool
	NativeCosmologyPrediction     bool
	Verdict                       string
	Reason                        string
}

type ResidualLedger struct {
	Executed                       bool
	EHComparatorResidual           float64
	CosmologicalComparatorResidual float64
	ResidualsAreSynthetic          bool
	ResidualsBridgeOnly            bool
	ResidualsZeroByConstruction    bool
	ObservedComparatorUsed         bool
	Verdict                        string
	Reason                         string
}

type Airlock struct {
	Executed                    bool
	NumericalAdapterExecuted    bool
	SyntheticOnly               bool
	ObservedComparatorImported  bool
	LambdaNativeSelected        bool
	F2NativeSelected            bool
	F4NativeSelected            bool
	F2LambdaNative              bool
	F4LambdaNative              bool
	PlanckNewtonNative          bool
	CosmologicalConstantNative  bool
	VacuumSubtractionNative     bool
	RenormalizationSchemeNative bool
	NewtonConstantDerived       bool
	CosmologicalConstantDerived bool
	NativeNormalizationWrite    bool
	Verdict                     string
	Reason                      string
}

type Firewall struct {
	Executed                     bool
	NewtonConstantImported       bool
	PlanckScaleImported          bool
	CutoffLambdaImported         bool
	F2MomentImported             bool
	F4MomentImported             bool
	F2LambdaProductImported      bool
	F4LambdaProductImported      bool
	CosmologicalConstantImported bool
	DarkEnergyImported           bool
	VacuumSubtractionImported    bool
	ObservedComparatorImported   bool
	SyntheticOutputNativeWrite   bool
	Verdict                      string
	Reason                       string
}

type RegistryUpdate struct {
	NativeEntries        []string
	BridgeEntries        []string
	EnvironmentalEntries []string
	FailedRoutes         []string
	OpenTheorems         []string
}

type NextStep struct {
	Gate                       int
	Title, Reason, PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Inputs      SyntheticInputs
	Output      AdapterOutput
	Residuals   ResidualLedger
	Airlock     Airlock
	Firewall    Firewall
	Registry    RegistryUpdate
	Next        NextStep
	Truth       string
}

var cache struct {
	sync.Once
	a   Analysis
	err error
}

func BuildDefault() (Analysis, error) {
	cache.Once.Do(func() { cache.a, cache.err = Build() })
	return cache.a, cache.err
}

func Build() (Analysis, error) {
	g514, err := generation2spectralcutoffrenormalizationairlock.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate514 cutoff-renormalization airlock: %w", err)
	}
	a := Analysis{}
	a.Inheritance = buildInheritance(g514)
	a.Inputs = buildInputs()
	a.Output = buildOutput(a.Inputs)
	a.Residuals = buildResiduals(a.Inputs, a.Output)
	a.Airlock = buildAirlock()
	a.Firewall = buildFirewall()
	a.Registry = buildRegistry(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g514 generation2spectralcutoffrenormalizationairlock.Analysis) Inheritance {
	return Inheritance{
		Executed:                   true,
		Gate514Inherited:           g514.Inheritance.Gate513Inherited && g514.Schema.RowsBridgeOnly && g514.Schema.RowsRejectNativePromotion,
		RedactedSchemaAccepted:     g514.Schema.AcceptedRedactedRows == 10 && g514.Schema.NumericalRows == 0,
		RequiredRows:               g514.Schema.RequiredRowCount,
		AcceptedCases:              g514.Preflight.AcceptedCases,
		RejectedCases:              g514.Preflight.RejectedCases,
		Gate514NoAdapterExecuted:   !g514.Airlock.NumericalAdapterExecuted,
		Gate514NativeWriteBlocked:  !g514.Airlock.NativeNormalizationWrite,
		Gate514LambdaSelected:      g514.Airlock.LambdaCutoffSelected,
		Gate514F2Selected:          g514.Airlock.F2MomentSelected,
		Gate514F4Selected:          g514.Airlock.F4MomentSelected,
		Gate514NewtonDerived:       g514.Airlock.NewtonConstantDerived,
		Gate514CosmologicalDerived: g514.Airlock.CosmologicalConstantDerived,
		Verdict:                    StatusGate514AirlockInherited,
		Reason:                     "Gate515 inherits Gate514's fail-closed cutoff/moment/renormalization airlock, then deliberately leaves the redacted preflight layer and runs only a synthetic, fake-number bridge adapter.",
	}
}

func buildInputs() SyntheticInputs {
	lambda := 2.0
	f2 := 3.0
	f4 := 5.0
	f0 := 7.0
	subtraction := 11.0
	eh := f2 * lambda * lambda * a2Prefactor
	cosmo := f4*math.Pow(lambda, 4)*a0Prefactor - subtraction
	return SyntheticInputs{
		Executed:                  true,
		LambdaCutoff:              lambda,
		F2Moment:                  f2,
		F4Moment:                  f4,
		F0Moment:                  f0,
		VacuumSubtraction:         subtraction,
		EHComparator:              eh,
		CosmologicalComparator:    cosmo,
		AllInputsSynthetic:        true,
		AllRowsBridgeOnly:         true,
		AllNativePromotionBlocked: true,
		ObservedDataImported:      false,
		Verdict:                   strings.Join([]string{StatusSyntheticInputsFake, StatusSyntheticAdapterBridgeOnly}, ";"),
		Reason:                    "the adapter uses fake positive numbers Λ=2, f₂=3, f₄=5, f₀=7, and δρ=11 only to test bridge formula plumbing; the comparator rows are synthetic and reject native promotion.",
	}
}

func buildOutput(in SyntheticInputs) AdapterOutput {
	f2lambda2 := in.F2Moment * in.LambdaCutoff * in.LambdaCutoff
	f4lambda4 := in.F4Moment * math.Pow(in.LambdaCutoff, 4)
	eh := f2lambda2 * a2Prefactor
	volume := f4lambda4 * a0Prefactor
	return AdapterOutput{
		Executed:                      true,
		FiniteTraceDimension:          finiteTraceDimension,
		A2PrefactorPerF2LambdaSquared: a2Prefactor,
		A0PrefactorPerF4LambdaFourth:  a0Prefactor,
		A4PrefactorPerF0:              a4Prefactor,
		F2LambdaSquared:               f2lambda2,
		F4LambdaFourth:                f4lambda4,
		EinsteinHilbertCoefficient:    eh,
		CosmologicalVolumeRaw:         volume,
		CosmologicalAfterSubtraction:  volume - in.VacuumSubtraction,
		CurvatureSquaredCoefficient:   in.F0Moment * a4Prefactor,
		NativeGravityPrediction:       false,
		NativeCosmologyPrediction:     false,
		Verdict:                       strings.Join([]string{StatusSyntheticAdapterExecuted, StatusA2EHComputedBridgeOnly, StatusA0CosmologicalComputedBridgeOnly, StatusA4CurvatureComputedBridgeOnly}, ";"),
		Reason:                        "the adapter computes the a2, a0, and a4 bridge coefficients from fake Λ/f₂/f₄/f₀ inputs using the native prefactor hierarchy, but every output remains synthetic bridge arithmetic.",
	}
}

func buildResiduals(in SyntheticInputs, out AdapterOutput) ResidualLedger {
	ehResidual := math.Abs(out.EinsteinHilbertCoefficient - in.EHComparator)
	cosmoResidual := math.Abs(out.CosmologicalAfterSubtraction - in.CosmologicalComparator)
	return ResidualLedger{
		Executed:                       true,
		EHComparatorResidual:           ehResidual,
		CosmologicalComparatorResidual: cosmoResidual,
		ResidualsAreSynthetic:          true,
		ResidualsBridgeOnly:            true,
		ResidualsZeroByConstruction:    nearly(ehResidual, 0, 1e-12) && nearly(cosmoResidual, 0, 1e-12),
		ObservedComparatorUsed:         false,
		Verdict:                        strings.Join([]string{StatusVacuumSubtractionResidualPlumbingTested, StatusSyntheticResidualsComputed}, ";"),
		Reason:                         "residuals are computed only against synthetic comparator rows generated from the same fake bridge ledger; zero residuals test plumbing, not physics.",
	}
}

func buildAirlock() Airlock {
	return Airlock{
		Executed: true, NumericalAdapterExecuted: true, SyntheticOnly: true, ObservedComparatorImported: false,
		LambdaNativeSelected: false, F2NativeSelected: false, F4NativeSelected: false, F2LambdaNative: false, F4LambdaNative: false,
		PlanckNewtonNative: false, CosmologicalConstantNative: false, VacuumSubtractionNative: false, RenormalizationSchemeNative: false,
		NewtonConstantDerived: false, CosmologicalConstantDerived: false, NativeNormalizationWrite: false,
		Verdict: strings.Join([]string{StatusFailedSyntheticOutputsNotNative, StatusFailedSyntheticOutputsNotNewtonOrCosmological, StatusFailedLambdaF2F4StillNotDerived, StatusFailedVacuumSubtractionStillNotNative, StatusFailedNativeNormalizationStillBlocked}, ";"),
		Reason:  "the numerical adapter is allowed only because every input is fake and bridge-only; the run derives no Λ, f₂, f₄, f₂Λ², f₄Λ⁴, Newton constant, cosmological constant, or subtraction rule.",
	}
}

func buildFirewall() Firewall {
	return Firewall{Executed: true, Verdict: strings.Join([]string{StatusNoObservedGravityCosmologyDataImported, StatusFailedObservedComparatorNotUsed, StatusFirewallNoObservedImport, StatusFirewallSyntheticNativeWriteBlocked}, ";"), Reason: "Gate515 imports no Newton constant, Planck scale, cutoff, spectral moments, moment products, cosmological constant, dark energy density, observed comparator, or vacuum-subtraction prescription; synthetic outputs are blocked from the native registry."}
}

func buildRegistry(a Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries:        []string{"No new native gravity/cosmology normalization is added. Gate513's stripped prefactor hierarchy and Gate510/Gate512 dimensionless coefficients remain the native results."},
		BridgeEntries:        []string{"A synthetic bridge adapter computes a2 Einstein-Hilbert, a0 cosmological-volume, and a4 curvature-squared coefficients from fake Λ/f₂/f₄/f₀ rows.", "Synthetic residual plumbing for Planck/Newton-normalization and cosmological-subtraction comparator rows is tested."},
		EnvironmentalEntries: []string{"Any real numerical Λ, f₂, f₄, f₀, Newton/Planck normalization, cosmological comparator, dark-energy density, vacuum subtraction, or renormalization scheme."},
		FailedRoutes:         []string{"Treating synthetic adapter outputs as native predictions.", "Treating zero residuals against fake comparator rows as physical success.", "Promoting Λ, f₂, f₄, f₂Λ², f₄Λ⁴, or vacuum subtraction to native data."},
		OpenTheorems:         []string{"A native regulator/profile selector, if one exists.", "A native renormalization or vacuum-subtraction principle, if one exists.", "A return to scale-free topological gravity invariants such as Euler/Gauss-Bonnet and Pontryagin/signature ledgers."},
	}
}

func buildNext() NextStep {
	return NextStep{Gate: 516, Title: "Topological Gravity Characteristic-Class Ledger", Reason: "Gate515 proves that numerical gravity/cosmology matching is only bridge plumbing; the next native lane should avoid scales and audit curvature topology.", PrimaryTask: "Audit Euler/Gauss-Bonnet, Pontryagin/signature, and boundary characteristic-class sockets as scale-free gravitational topological invariants, without importing Newton, Λ, or cosmological data."}
}

func validate(a Analysis) error {
	checks := []struct {
		ok  bool
		msg string
	}{
		{a.Inheritance.Executed && a.Inheritance.Gate514Inherited && a.Inheritance.RedactedSchemaAccepted && a.Inheritance.RequiredRows == 10 && a.Inheritance.AcceptedCases == 1 && a.Inheritance.RejectedCases == 8 && a.Inheritance.Gate514NoAdapterExecuted && a.Inheritance.Gate514NativeWriteBlocked && !a.Inheritance.Gate514LambdaSelected && !a.Inheritance.Gate514F2Selected && !a.Inheritance.Gate514F4Selected && !a.Inheritance.Gate514NewtonDerived && !a.Inheritance.Gate514CosmologicalDerived, "Gate515 inheritance invalid"},
		{a.Inputs.Executed && a.Inputs.AllInputsSynthetic && a.Inputs.AllRowsBridgeOnly && a.Inputs.AllNativePromotionBlocked && !a.Inputs.ObservedDataImported && a.Inputs.LambdaCutoff == 2 && a.Inputs.F2Moment == 3 && a.Inputs.F4Moment == 5 && a.Inputs.F0Moment == 7 && a.Inputs.VacuumSubtraction == 11, "Gate515 synthetic inputs invalid"},
		{a.Output.Executed && nearly(a.Output.F2LambdaSquared, 12, 1e-12) && nearly(a.Output.F4LambdaFourth, 80, 1e-12) && nearly(a.Output.EinsteinHilbertCoefficient, 6.0/(math.Pi*math.Pi), 1e-12) && nearly(a.Output.CosmologicalVolumeRaw, 480.0/(math.Pi*math.Pi), 1e-12) && nearly(a.Output.CosmologicalAfterSubtraction, 480.0/(math.Pi*math.Pi)-11, 1e-12) && nearly(a.Output.CurvatureSquaredCoefficient, 7.0/(60.0*math.Pi*math.Pi), 1e-12) && !a.Output.NativeGravityPrediction && !a.Output.NativeCosmologyPrediction, "Gate515 adapter outputs invalid"},
		{a.Residuals.Executed && a.Residuals.ResidualsAreSynthetic && a.Residuals.ResidualsBridgeOnly && a.Residuals.ResidualsZeroByConstruction && !a.Residuals.ObservedComparatorUsed && nearly(a.Residuals.EHComparatorResidual, 0, 1e-12) && nearly(a.Residuals.CosmologicalComparatorResidual, 0, 1e-12), "Gate515 residuals invalid"},
		{a.Airlock.Executed && a.Airlock.NumericalAdapterExecuted && a.Airlock.SyntheticOnly && !a.Airlock.ObservedComparatorImported && !a.Airlock.LambdaNativeSelected && !a.Airlock.F2NativeSelected && !a.Airlock.F4NativeSelected && !a.Airlock.F2LambdaNative && !a.Airlock.F4LambdaNative && !a.Airlock.PlanckNewtonNative && !a.Airlock.CosmologicalConstantNative && !a.Airlock.VacuumSubtractionNative && !a.Airlock.RenormalizationSchemeNative && !a.Airlock.NewtonConstantDerived && !a.Airlock.CosmologicalConstantDerived && !a.Airlock.NativeNormalizationWrite, "Gate515 airlock invalid"},
		{a.Firewall.Executed && !a.Firewall.NewtonConstantImported && !a.Firewall.PlanckScaleImported && !a.Firewall.CutoffLambdaImported && !a.Firewall.F2MomentImported && !a.Firewall.F4MomentImported && !a.Firewall.F2LambdaProductImported && !a.Firewall.F4LambdaProductImported && !a.Firewall.CosmologicalConstantImported && !a.Firewall.DarkEnergyImported && !a.Firewall.VacuumSubtractionImported && !a.Firewall.ObservedComparatorImported && !a.Firewall.SyntheticOutputNativeWrite, "Gate515 firewall invalid"},
	}
	for _, c := range checks {
		if !c.ok {
			return fmt.Errorf(c.msg)
		}
	}
	return nil
}

func statuses() []string {
	return []string{
		StatusGate514AirlockInherited,
		StatusSyntheticAdapterExecuted,
		StatusSyntheticInputsFake,
		StatusA2EHComputedBridgeOnly,
		StatusA0CosmologicalComputedBridgeOnly,
		StatusA4CurvatureComputedBridgeOnly,
		StatusVacuumSubtractionResidualPlumbingTested,
		StatusSyntheticResidualsComputed,
		StatusSyntheticAdapterBridgeOnly,
		StatusNoObservedGravityCosmologyDataImported,
		StatusFailedSyntheticOutputsNotNative,
		StatusFailedSyntheticOutputsNotNewtonOrCosmological,
		StatusFailedLambdaF2F4StillNotDerived,
		StatusFailedVacuumSubtractionStillNotNative,
		StatusFailedNativeNormalizationStillBlocked,
		StatusFailedObservedComparatorNotUsed,
		StatusFirewallNoObservedImport,
		StatusFirewallSyntheticNativeWriteBlocked,
	}
}

func truth(a Analysis) string {
	return "Gate 515 proves only that the gravity/cosmology adapter plumbing works when fed explicitly fake bridge numbers. It computes synthetic a2, a0, and a4 coefficients and synthetic residuals, but derives no cutoff, moments, Newton constant, cosmological constant, dark-energy value, vacuum subtraction, or physical gravity prediction. The native result remains the scale-free spectral prefactor hierarchy; all numerical normalization stays behind the bridge airlock."
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("Gate514 inherited=%t; redacted schema=%t; required rows=%d; accepted cases=%d; rejected cases=%d; prior adapter executed=%t; native write blocked=%t; Λ selected=%t; f2 selected=%t; f4 selected=%t; G derived=%t; Λ_cosmo derived=%t", x.Gate514Inherited, x.RedactedSchemaAccepted, x.RequiredRows, x.AcceptedCases, x.RejectedCases, !x.Gate514NoAdapterExecuted, x.Gate514NativeWriteBlocked, x.Gate514LambdaSelected, x.Gate514F2Selected, x.Gate514F4Selected, x.Gate514NewtonDerived, x.Gate514CosmologicalDerived)
}
func FormatInputs(x SyntheticInputs) string {
	return fmt.Sprintf("Λ=%.12g; f2=%.12g; f4=%.12g; f0=%.12g; δρ=%.12g; synthetic=%t; bridge-only=%t; native-promotion blocked=%t; observed imported=%t", x.LambdaCutoff, x.F2Moment, x.F4Moment, x.F0Moment, x.VacuumSubtraction, x.AllInputsSynthetic, x.AllRowsBridgeOnly, x.AllNativePromotionBlocked, x.ObservedDataImported)
}
func FormatOutput(x AdapterOutput) string {
	return fmt.Sprintf("TrF=%.12g; f2Λ²=%.12g; f4Λ4=%.12g; C_EH=%.12g; C_Λ_raw=%.12g; C_Λ_after_subtraction=%.12g; C_a4=%.12g; native gravity prediction=%t; native cosmology prediction=%t", x.FiniteTraceDimension, x.F2LambdaSquared, x.F4LambdaFourth, x.EinsteinHilbertCoefficient, x.CosmologicalVolumeRaw, x.CosmologicalAfterSubtraction, x.CurvatureSquaredCoefficient, x.NativeGravityPrediction, x.NativeCosmologyPrediction)
}
func FormatResiduals(x ResidualLedger) string {
	return fmt.Sprintf("EH residual=%.12g; cosmological residual=%.12g; synthetic=%t; bridge-only=%t; zero by construction=%t; observed comparator used=%t", x.EHComparatorResidual, x.CosmologicalComparatorResidual, x.ResidualsAreSynthetic, x.ResidualsBridgeOnly, x.ResidualsZeroByConstruction, x.ObservedComparatorUsed)
}
func FormatAirlock(x Airlock) string {
	return fmt.Sprintf("adapter executed=%t; synthetic only=%t; observed imported=%t; Λ native=%t; f2 native=%t; f4 native=%t; f2Λ² native=%t; f4Λ4 native=%t; Planck/Newton native=%t; Λ_cosmo native=%t; subtraction native=%t; renormalization native=%t; G derived=%t; Λ_cosmo derived=%t; native write=%t", x.NumericalAdapterExecuted, x.SyntheticOnly, x.ObservedComparatorImported, x.LambdaNativeSelected, x.F2NativeSelected, x.F4NativeSelected, x.F2LambdaNative, x.F4LambdaNative, x.PlanckNewtonNative, x.CosmologicalConstantNative, x.VacuumSubtractionNative, x.RenormalizationSchemeNative, x.NewtonConstantDerived, x.CosmologicalConstantDerived, x.NativeNormalizationWrite)
}
func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("G imported=%t; Planck imported=%t; Λ imported=%t; f2 imported=%t; f4 imported=%t; f2Λ² imported=%t; f4Λ4 imported=%t; Λ_cosmo imported=%t; dark energy imported=%t; subtraction imported=%t; observed comparator imported=%t; synthetic native write=%t", x.NewtonConstantImported, x.PlanckScaleImported, x.CutoffLambdaImported, x.F2MomentImported, x.F4MomentImported, x.F2LambdaProductImported, x.F4LambdaProductImported, x.CosmologicalConstantImported, x.DarkEnergyImported, x.VacuumSubtractionImported, x.ObservedComparatorImported, x.SyntheticOutputNativeWrite)
}

func Markdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 515 Registry Audit — Bridge-Only Gravity/Cosmology Adapter Dry-Run\n\n")
	b.WriteString("## Verdict\n\n```text\n" + strings.Join(statuses(), "\n") + "\n```\n\n")
	b.WriteString("## Inherited boundary\n\n" + a.Inheritance.Reason + "\n\n```text\n" + FormatInheritance(a.Inheritance) + "\n```\n\n")
	b.WriteString("## Synthetic bridge inputs\n\n" + a.Inputs.Reason + "\n\n```text\n" + FormatInputs(a.Inputs) + "\n```\n\n")
	b.WriteString("## Adapter calculation\n\n" + a.Output.Reason + "\n\n```text\n" + FormatOutput(a.Output) + "\n```\n\n")
	b.WriteString("Formula ledger:\n\n```text\nC_EH = f2 Λ² · 1/(2π²)\nC_Λ  = f4 Λ⁴ · 6/π²\nC_a4 = f0 · 1/(60π²)\nC_Λ,sub = C_Λ - δρ\n```\n\n")
	b.WriteString("## Residual ledger\n\n" + a.Residuals.Reason + "\n\n```text\n" + FormatResiduals(a.Residuals) + "\n```\n\n")
	b.WriteString("## Airlock result\n\n" + a.Airlock.Reason + "\n\n```text\n" + FormatAirlock(a.Airlock) + "\n```\n\n")
	b.WriteString("## Firewall result\n\n" + a.Firewall.Reason + "\n\n```text\n" + FormatFirewall(a.Firewall) + "\n```\n\n")
	b.WriteString("## Registry update\n\n")
	writeList(&b, "Native entries", a.Registry.NativeEntries)
	writeList(&b, "Bridge entries", a.Registry.BridgeEntries)
	writeList(&b, "Environmental entries", a.Registry.EnvironmentalEntries)
	writeList(&b, "Failed routes", a.Registry.FailedRoutes)
	writeList(&b, "Open theorems", a.Registry.OpenTheorems)
	b.WriteString("## Next step\n\nGate516 should be:\n\n```text\nGate 516 — " + a.Next.Title + "\n```\n\nPrimary task:\n\n```text\n" + a.Next.PrimaryTask + "\n```\n\n")
	b.WriteString("## Truth statement\n\n" + a.Truth + "\n")
	return b.String()
}

func writeList(b *strings.Builder, title string, xs []string) {
	b.WriteString("### " + title + "\n\n")
	if len(xs) == 0 {
		b.WriteString("- None.\n\n")
		return
	}
	for _, x := range xs {
		b.WriteString("- " + x + "\n")
	}
	b.WriteString("\n")
}

func nearly(a, b, eps float64) bool { return math.Abs(a-b) <= eps }
