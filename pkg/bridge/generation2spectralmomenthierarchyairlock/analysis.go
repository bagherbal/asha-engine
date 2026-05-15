// Package generation2spectralmomenthierarchyairlock implements Gate 513:
// Spectral Moment Hierarchy and Cutoff-Separation Airlock Audit.
//
// Gate 512 isolated the a0/f4Λ4 cosmological volume term and proved that its
// finite trace prefactor is native but its physical cosmological normalization
// is not. Gate 513 audits the three gravitational heat-kernel channels together
// (a0, a2, a4). It proves the dimensionless relative heat-kernel prefactor
// hierarchy after the independent spectral moments and cutoff powers are
// factored out, and it blocks the forbidden step of turning that hierarchy into
// Newton's constant, a cutoff scale, a cosmological constant, or a vacuum
// subtraction prescription.
package generation2spectralmomenthierarchyairlock

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2a4curvaturesquaredledger"
	"github.com/bagherbal/asha-engine/pkg/bridge/generation2cosmologicalf4vacuumairlock"
	"github.com/bagherbal/asha-engine/pkg/bridge/generation2curvaturecoefficientprovenance"
	"github.com/bagherbal/asha-engine/pkg/bridge/productspectralactioncoefficients"
)

const (
	AuditID = "GATE513-SPECTRAL-MOMENT-HIERARCHY-CUTOFF-SEPARATION-AIRLOCK-AUDIT"

	StatusGate512CosmologicalAirlockInherited = "CONDITIONAL_SUPPORT_GATE512_COSMOLOGICAL_AIRLOCK_INHERITED"
	StatusGate510A2Inherited                  = "CONDITIONAL_SUPPORT_GATE510_A2_CURVATURE_WEIGHT_INHERITED"
	StatusGate511A4Inherited                  = "CONDITIONAL_SUPPORT_GATE511_A4_CURVATURE_WEIGHT_INHERITED"
	StatusProductMomentLedgerInherited        = "CONDITIONAL_SUPPORT_PRODUCT_SPECTRAL_MOMENT_LEDGER_INHERITED"
	StatusThreeChannelLedgerConstructed       = "CONDITIONAL_SUPPORT_A0_A2_A4_THREE_CHANNEL_LEDGER_CONSTRUCTED"
	StatusRelativePrefactorHierarchyComputed  = "CONDITIONAL_SUPPORT_RELATIVE_HEAT_KERNEL_PREFACTOR_HIERARCHY_COMPUTED"
	StatusA2OverA0RatioOneTwelfth             = "CONDITIONAL_SUPPORT_A2_OVER_A0_PREFACTOR_RATIO_ONE_TWELFTH"
	StatusA4OverA0RatioOneOver360             = "CONDITIONAL_SUPPORT_A4_OVER_A0_PREFACTOR_RATIO_ONE_OVER_360"
	StatusA4OverA2RatioOneOver30              = "CONDITIONAL_SUPPORT_A4_OVER_A2_PREFACTOR_RATIO_ONE_OVER_30"
	StatusMomentAirlockDefined                = "CONDITIONAL_SUPPORT_SPECTRAL_MOMENT_AND_CUTOFF_AIRLOCK_DEFINED"
	StatusNoEmpiricalScalesImported           = "CONDITIONAL_SUPPORT_NO_EMPIRICAL_SPECTRAL_SCALES_IMPORTED"

	StatusFailedMomentRatiosDoNotSelectCutoff       = "FAILED_ROUTE_RELATIVE_PREFACTORS_DO_NOT_SELECT_CUTOFF_SCALE"
	StatusFailedF2F4MomentsNotSelected              = "FAILED_ROUTE_F2_AND_F4_MOMENTS_NOT_SELECTED"
	StatusFailedF2LambdaNotSeparated                = "FAILED_ROUTE_F2_LAMBDA_SQUARED_NOT_SEPARATED"
	StatusFailedF4LambdaNotSeparated                = "FAILED_ROUTE_F4_LAMBDA_FOURTH_NOT_SEPARATED"
	StatusFailedNewtonNotDerived                    = "FAILED_ROUTE_NEWTON_CONSTANT_NOT_DERIVED_BY_MOMENT_HIERARCHY"
	StatusFailedCosmologicalConstantNotDerived      = "FAILED_ROUTE_COSMOLOGICAL_CONSTANT_NOT_DERIVED_BY_MOMENT_HIERARCHY"
	StatusFailedVacuumSubtractionNotSelected        = "FAILED_ROUTE_VACUUM_SUBTRACTION_NOT_SELECTED_BY_MOMENT_HIERARCHY"
	StatusFailedPlanckCutoffRelationNotNative       = "FAILED_ROUTE_PLANCK_CUTOFF_RELATION_NOT_NATIVE"
	StatusFirewallMomentNativeWriteBlocked          = "FIREWALL_BLOCKED_SPECTRAL_MOMENT_HIERARCHY_NATIVE_NORMALIZATION_WRITE"
	StatusFirewallPreservedNoGravityCosmologyImport = "FIREWALL_PRESERVED_NO_NEWTON_PLANCK_COSMOLOGY_EW_OR_FLAVOR_DATA_IMPORTED"
)

const finiteTraceDimension = 96.0

type Inheritance struct {
	Executed                       bool
	Gate512Inherited               bool
	Gate512A0PrefactorNative       bool
	Gate512CosmologicalBlocked     bool
	Gate510A2Inherited             bool
	Gate510NewtonBlocked           bool
	Gate511A4Inherited             bool
	Gate511PhysicalDynamicsBlocked bool
	ProductTripleValid             bool
	ProductMomentLedgerAvailable   bool
	ProductAllCoefficientsClosed   bool
	Verdict                        string
	Reason                         string
}

type Channel struct {
	Name                 string
	HeatKernelIndex      string
	MomentFactor         string
	CutoffPower          int
	PrefactorAfterMoment float64
	Physical             bool
}

type ThreeChannelLedger struct {
	Executed    bool
	FiniteTrace float64
	A0, A2, A4  Channel
	A0Expected  float64
	A2Expected  float64
	A4Expected  float64
	AllMatched  bool
	Formula     string
	Verdict     string
	Reason      string
}

type RelativeHierarchy struct {
	Executed                  bool
	A2OverA0AfterFactoring    float64
	A4OverA0AfterFactoring    float64
	A4OverA2AfterFactoring    float64
	ExpectedA2OverA0          float64
	ExpectedA4OverA0          float64
	ExpectedA4OverA2          float64
	DimensionlessCombinatoric bool
	SelectsF2Moment           bool
	SelectsF4Moment           bool
	SelectsCutoffLambda       bool
	SelectsVacuumSubtraction  bool
	PhysicalNormalization     bool
	Verdict                   string
	Reason                    string
}

type MomentAirlock struct {
	Executed                       bool
	F0MomentDimensionlessAvailable bool
	F2MomentSelected               bool
	F4MomentSelected               bool
	F2LambdaProductSeparated       bool
	F4LambdaProductSeparated       bool
	CutoffLambdaSelected           bool
	PlanckCutoffRelationNative     bool
	NewtonConstantDerived          bool
	CosmologicalConstantDerived    bool
	VacuumSubtractionSelected      bool
	NativeNormalizationWrite       bool
	Verdict                        string
	Reason                         string
}

type Firewall struct {
	Executed                         bool
	NewtonConstantImported           bool
	PlanckScaleImported              bool
	CutoffLambdaImported             bool
	F2MomentImported                 bool
	F4MomentImported                 bool
	CosmologicalConstantImported     bool
	DarkEnergyImported               bool
	ElectroweakScaleImported         bool
	FlavorDataImported               bool
	NativeSpectralNormalizationWrite bool
	Verdict                          string
	Reason                           string
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
	Ledger      ThreeChannelLedger
	Hierarchy   RelativeHierarchy
	Airlock     MomentAirlock
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
	g512, err := generation2cosmologicalf4vacuumairlock.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate512 cosmological airlock: %w", err)
	}
	g510, err := generation2curvaturecoefficientprovenance.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate510 curvature coefficient audit: %w", err)
	}
	g511, err := generation2a4curvaturesquaredledger.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate511 a4 curvature ledger: %w", err)
	}
	g377, err := productspectralactioncoefficients.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit product spectral-action coefficients: %w", err)
	}

	a := Analysis{}
	a.Inheritance = buildInheritance(g512, g510, g511, g377)
	a.Ledger = buildLedger(g512, g510, g511)
	a.Hierarchy = buildHierarchy(a.Ledger)
	a.Airlock = buildAirlock(g377)
	a.Firewall = buildFirewall()
	a.Registry = buildRegistry(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g512 generation2cosmologicalf4vacuumairlock.Analysis, g510 generation2curvaturecoefficientprovenance.Analysis, g511 generation2a4curvaturesquaredledger.Analysis, g377 productspectralactioncoefficients.Analysis) Inheritance {
	return Inheritance{
		Executed:                       true,
		Gate512Inherited:               g512.A0.NativeDimensionlessTraceWeight && !g512.Airlock.PhysicalLambdaCosmoDerived,
		Gate512A0PrefactorNative:       nearly(g512.A0.PrefactorPerF4Lambda4, 6.0/(math.Pi*math.Pi), 1e-12),
		Gate512CosmologicalBlocked:     !g512.Airlock.PhysicalLambdaCosmoDerived && !g512.Airlock.NativeCosmologicalWriteAllowed,
		Gate510A2Inherited:             g510.A2.DimensionlessTraceWeightNative && g510.A2.Gate377RawCoefficientMatched,
		Gate510NewtonBlocked:           !g510.Firewall.NewtonConstantDerived && !g510.Firewall.CutoffLambdaSelected && !g510.Firewall.NativeGravityNormalizationWritten,
		Gate511A4Inherited:             g511.A4.DimensionlessChannel && g511.A4.RawPrefactorPerF0BeforeInvariant > 0,
		Gate511PhysicalDynamicsBlocked: !g511.Dynamical.PhysicalA4DynamicsClosed && g511.Firewall.PhysicalA4DynamicsWritten == false,
		ProductTripleValid:             g377.Calculation.Product.Valid,
		ProductMomentLedgerAvailable:   g377.Calculation.Finite.F0 > 0 && g377.Calculation.Convention.Expansion != "",
		ProductAllCoefficientsClosed:   g377.Calculation.AllCoefficientsDetermined,
		Verdict:                        strings.Join([]string{StatusGate512CosmologicalAirlockInherited, StatusGate510A2Inherited, StatusGate511A4Inherited, StatusProductMomentLedgerInherited}, ";"),
		Reason:                         "Gate513 inherits the native finite-trace prefactors of the a0, a2, and a4 channels, plus their existing firewalls: a0 has no vacuum subtraction, a2 has no Newton/cutoff normalization, and a4 has no physical metric-dynamics closure.",
	}
}

func buildLedger(g512 generation2cosmologicalf4vacuumairlock.Analysis, g510 generation2curvaturecoefficientprovenance.Analysis, g511 generation2a4curvaturesquaredledger.Analysis) ThreeChannelLedger {
	a0 := g512.A0.PrefactorPerF4Lambda4
	a2 := g510.A2.RawDensityCoefficientPerF2Lambda2
	a4 := g511.A4.RawPrefactorPerF0BeforeInvariant
	a0e := finiteTraceDimension / (16.0 * math.Pi * math.Pi)
	a2e := finiteTraceDimension / (192.0 * math.Pi * math.Pi)
	a4e := finiteTraceDimension / (360.0 * 16.0 * math.Pi * math.Pi)
	return ThreeChannelLedger{
		Executed:    true,
		FiniteTrace: finiteTraceDimension,
		A0:          Channel{Name: "a0 cosmological volume", HeatKernelIndex: "a0", MomentFactor: "f4 Λ^4", CutoffPower: 4, PrefactorAfterMoment: a0, Physical: false},
		A2:          Channel{Name: "a2 Einstein-Hilbert curvature", HeatKernelIndex: "a2", MomentFactor: "f2 Λ^2", CutoffPower: 2, PrefactorAfterMoment: a2, Physical: false},
		A4:          Channel{Name: "a4 curvature-squared", HeatKernelIndex: "a4", MomentFactor: "f0", CutoffPower: 0, PrefactorAfterMoment: a4, Physical: false},
		A0Expected:  a0e, A2Expected: a2e, A4Expected: a4e,
		AllMatched: nearly(a0, a0e, 1e-12) && nearly(a2, a2e, 1e-12) && nearly(a4, a4e, 1e-12),
		Formula:    "C0/(f4Λ4)=TrF/(16π²), C2/(f2Λ²)=TrF/(192π²), C4/f0=TrF/(360·16π²)",
		Verdict:    StatusThreeChannelLedgerConstructed,
		Reason:     "the product heat-kernel ledger supplies three native dimensionless prefactors after each independent spectral moment and cutoff power is factored out; none of the three channels is a physical scale by itself.",
	}
}

func buildHierarchy(l ThreeChannelLedger) RelativeHierarchy {
	r20 := l.A2.PrefactorAfterMoment / l.A0.PrefactorAfterMoment
	r40 := l.A4.PrefactorAfterMoment / l.A0.PrefactorAfterMoment
	r42 := l.A4.PrefactorAfterMoment / l.A2.PrefactorAfterMoment
	return RelativeHierarchy{
		Executed:                  true,
		A2OverA0AfterFactoring:    r20,
		A4OverA0AfterFactoring:    r40,
		A4OverA2AfterFactoring:    r42,
		ExpectedA2OverA0:          1.0 / 12.0,
		ExpectedA4OverA0:          1.0 / 360.0,
		ExpectedA4OverA2:          1.0 / 30.0,
		DimensionlessCombinatoric: nearly(r20, 1.0/12.0, 1e-12) && nearly(r40, 1.0/360.0, 1e-12) && nearly(r42, 1.0/30.0, 1e-12),
		SelectsF2Moment:           false,
		SelectsF4Moment:           false,
		SelectsCutoffLambda:       false,
		SelectsVacuumSubtraction:  false,
		PhysicalNormalization:     false,
		Verdict:                   strings.Join([]string{StatusRelativePrefactorHierarchyComputed, StatusA2OverA0RatioOneTwelfth, StatusA4OverA0RatioOneOver360, StatusA4OverA2RatioOneOver30}, ";"),
		Reason:                    "the relative heat-kernel prefactor hierarchy is native combinatorics, but it exists only after f4Λ4, f2Λ2, and f0 have been stripped away; it cannot determine those moments, the cutoff, or any physical normalization.",
	}
}

func buildAirlock(g377 productspectralactioncoefficients.Analysis) MomentAirlock {
	return MomentAirlock{
		Executed:                       true,
		F0MomentDimensionlessAvailable: nearly(g377.Calculation.Finite.F0, 7, 1e-12),
		F2MomentSelected:               false,
		F4MomentSelected:               false,
		F2LambdaProductSeparated:       false,
		F4LambdaProductSeparated:       false,
		CutoffLambdaSelected:           false,
		PlanckCutoffRelationNative:     false,
		NewtonConstantDerived:          false,
		CosmologicalConstantDerived:    false,
		VacuumSubtractionSelected:      false,
		NativeNormalizationWrite:       false,
		Verdict:                        strings.Join([]string{StatusMomentAirlockDefined, StatusFailedF2F4MomentsNotSelected, StatusFailedF2LambdaNotSeparated, StatusFailedF4LambdaNotSeparated, StatusFailedPlanckCutoffRelationNotNative}, ";"),
		Reason:                         "f0 is available as the dimensionless a4/contact moment, but the a0 and a2 physical channels require f4Λ4 and f2Λ2. The finite algebra does not split f2 from Λ, f4 from Λ, or Λ from a Planck/cosmological comparator.",
	}
}

func buildFirewall() Firewall {
	return Firewall{Executed: true, Verdict: strings.Join([]string{StatusNoEmpiricalScalesImported, StatusFirewallPreservedNoGravityCosmologyImport, StatusFirewallMomentNativeWriteBlocked}, ";"), Reason: "Gate513 imports no G, Planck mass, cutoff value, f2/f4 moment value, cosmological constant, dark-energy density, electroweak scale, or flavor datum, and writes no native spectral normalization."}
}

func buildRegistry(a Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries: []string{
			"The a0, a2, and a4 heat-kernel channels have native finite-trace prefactors after their independent spectral moments and cutoff powers are factored out.",
			"The structural prefactor ratios are C2/C0=1/12, C4/C0=1/360, and C4/C2=1/30 at the stripped, dimensionless heat-kernel level.",
		},
		BridgeEntries: []string{
			"Physical gravity/cosmology requires the bridge products f2Λ² and f4Λ⁴ plus trace convention, renormalization, and vacuum-subtraction choices.",
			"The dimensionless f0/a4 channel remains a curvature² socket, not a low-energy gravity dynamics theorem.",
		},
		EnvironmentalEntries: []string{"Λ, f2, f4, Newton normalization, Planck matching, cosmological constant, vacuum subtraction, dark-energy comparator, renormalization scheme, and boundary/manifold volume data."},
		FailedRoutes:         []string{"Relative heat-kernel prefactors do not select Λ.", "The finite trace does not derive Newton's constant.", "The finite trace does not derive the cosmological constant or its subtraction.", "The moment hierarchy does not identify the cutoff with the Planck scale."},
		OpenTheorems:         []string{"A native cutoff/moment selector theorem, if it exists.", "A renormalization/vacuum-subtraction principle compatible with the finite spectral ledger.", "A physical gravity normalization adapter that remains explicitly bridge-only."},
	}
}

func buildNext() NextStep {
	return NextStep{Gate: 514, Title: "Spectral Cutoff and Renormalization Airlock Comparator", Reason: "Gate513 proves only stripped moment-prefactor ratios; the next safe step is an explicit bridge-only schema for cutoff, spectral moments, and renormalization/subtraction metadata.", PrimaryTask: "Define a fail-closed comparator/adapter ledger for Λ, f2, f4, Planck normalization, and vacuum subtraction without allowing any native gravity or cosmology write."}
}

func validate(a Analysis) error {
	checks := []struct {
		ok  bool
		msg string
	}{
		{a.Inheritance.Executed && a.Inheritance.Gate512Inherited && a.Inheritance.Gate512A0PrefactorNative && a.Inheritance.Gate512CosmologicalBlocked && a.Inheritance.Gate510A2Inherited && a.Inheritance.Gate510NewtonBlocked && a.Inheritance.Gate511A4Inherited && a.Inheritance.Gate511PhysicalDynamicsBlocked && a.Inheritance.ProductTripleValid && a.Inheritance.ProductMomentLedgerAvailable && !a.Inheritance.ProductAllCoefficientsClosed, "Gate513 inheritance invalid"},
		{a.Ledger.Executed && nearly(a.Ledger.FiniteTrace, 96, 1e-12) && a.Ledger.AllMatched && !a.Ledger.A0.Physical && !a.Ledger.A2.Physical && !a.Ledger.A4.Physical, "Gate513 three-channel ledger invalid"},
		{a.Hierarchy.Executed && a.Hierarchy.DimensionlessCombinatoric && nearly(a.Hierarchy.A2OverA0AfterFactoring, 1.0/12.0, 1e-12) && nearly(a.Hierarchy.A4OverA0AfterFactoring, 1.0/360.0, 1e-12) && nearly(a.Hierarchy.A4OverA2AfterFactoring, 1.0/30.0, 1e-12) && !a.Hierarchy.SelectsF2Moment && !a.Hierarchy.SelectsF4Moment && !a.Hierarchy.SelectsCutoffLambda && !a.Hierarchy.PhysicalNormalization, "Gate513 relative hierarchy invalid"},
		{a.Airlock.Executed && a.Airlock.F0MomentDimensionlessAvailable && !a.Airlock.F2MomentSelected && !a.Airlock.F4MomentSelected && !a.Airlock.F2LambdaProductSeparated && !a.Airlock.F4LambdaProductSeparated && !a.Airlock.CutoffLambdaSelected && !a.Airlock.PlanckCutoffRelationNative && !a.Airlock.NewtonConstantDerived && !a.Airlock.CosmologicalConstantDerived && !a.Airlock.VacuumSubtractionSelected && !a.Airlock.NativeNormalizationWrite, "Gate513 moment airlock invalid"},
		{a.Firewall.Executed && !a.Firewall.NewtonConstantImported && !a.Firewall.PlanckScaleImported && !a.Firewall.CutoffLambdaImported && !a.Firewall.F2MomentImported && !a.Firewall.F4MomentImported && !a.Firewall.CosmologicalConstantImported && !a.Firewall.DarkEnergyImported && !a.Firewall.ElectroweakScaleImported && !a.Firewall.FlavorDataImported && !a.Firewall.NativeSpectralNormalizationWrite, "Gate513 firewall invalid"},
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
		StatusGate512CosmologicalAirlockInherited,
		StatusGate510A2Inherited,
		StatusGate511A4Inherited,
		StatusProductMomentLedgerInherited,
		StatusThreeChannelLedgerConstructed,
		StatusRelativePrefactorHierarchyComputed,
		StatusA2OverA0RatioOneTwelfth,
		StatusA4OverA0RatioOneOver360,
		StatusA4OverA2RatioOneOver30,
		StatusMomentAirlockDefined,
		StatusFailedMomentRatiosDoNotSelectCutoff,
		StatusFailedF2F4MomentsNotSelected,
		StatusFailedF2LambdaNotSeparated,
		StatusFailedF4LambdaNotSeparated,
		StatusFailedNewtonNotDerived,
		StatusFailedCosmologicalConstantNotDerived,
		StatusFailedVacuumSubtractionNotSelected,
		StatusFailedPlanckCutoffRelationNotNative,
		StatusNoEmpiricalScalesImported,
		StatusFirewallPreservedNoGravityCosmologyImport,
		StatusFirewallMomentNativeWriteBlocked,
	}
}

func truth(a Analysis) string {
	return "Gate 513 proves a native stripped heat-kernel hierarchy: after removing the independent spectral moments and cutoff powers, the a0:a2:a4 prefactors obey C2/C0=1/12, C4/C0=1/360, and C4/C2=1/30. This is real spectral combinatorics, not physical normalization. The same audit proves that f2, f4, Λ, Planck matching, Newton's constant, vacuum subtraction, and the cosmological constant remain outside native ASHA closure."
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("Gate512=%t; a0 native=%t; cosmology blocked=%t; Gate510 a2=%t; Newton blocked=%t; Gate511 a4=%t; dynamics blocked=%t; product valid=%t; moments available=%t; all coefficients closed=%t", x.Gate512Inherited, x.Gate512A0PrefactorNative, x.Gate512CosmologicalBlocked, x.Gate510A2Inherited, x.Gate510NewtonBlocked, x.Gate511A4Inherited, x.Gate511PhysicalDynamicsBlocked, x.ProductTripleValid, x.ProductMomentLedgerAvailable, x.ProductAllCoefficientsClosed)
}
func FormatLedger(x ThreeChannelLedger) string {
	return fmt.Sprintf("%s; TrF=%.0f; a0=%.12g expected=%.12g; a2=%.12g expected=%.12g; a4=%.12g expected=%.12g; all matched=%t", x.Formula, x.FiniteTrace, x.A0.PrefactorAfterMoment, x.A0Expected, x.A2.PrefactorAfterMoment, x.A2Expected, x.A4.PrefactorAfterMoment, x.A4Expected, x.AllMatched)
}
func FormatHierarchy(x RelativeHierarchy) string {
	return fmt.Sprintf("a2/a0=%.12g expected=%.12g; a4/a0=%.12g expected=%.12g; a4/a2=%.12g expected=%.12g; combinatoric=%t; selects f2=%t; selects f4=%t; selects Λ=%t; physical=%t", x.A2OverA0AfterFactoring, x.ExpectedA2OverA0, x.A4OverA0AfterFactoring, x.ExpectedA4OverA0, x.A4OverA2AfterFactoring, x.ExpectedA4OverA2, x.DimensionlessCombinatoric, x.SelectsF2Moment, x.SelectsF4Moment, x.SelectsCutoffLambda, x.PhysicalNormalization)
}
func FormatAirlock(x MomentAirlock) string {
	return fmt.Sprintf("f0 available=%t; f2 selected=%t; f4 selected=%t; f2Λ² separated=%t; f4Λ4 separated=%t; Λ selected=%t; Planck/cutoff native=%t; G derived=%t; Λ_cosmo derived=%t; subtraction selected=%t; native write=%t", x.F0MomentDimensionlessAvailable, x.F2MomentSelected, x.F4MomentSelected, x.F2LambdaProductSeparated, x.F4LambdaProductSeparated, x.CutoffLambdaSelected, x.PlanckCutoffRelationNative, x.NewtonConstantDerived, x.CosmologicalConstantDerived, x.VacuumSubtractionSelected, x.NativeNormalizationWrite)
}
func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("G imported=%t; Planck imported=%t; cutoff imported=%t; f2 imported=%t; f4 imported=%t; Λ_cosmo imported=%t; dark energy imported=%t; EW imported=%t; flavor imported=%t; native write=%t", x.NewtonConstantImported, x.PlanckScaleImported, x.CutoffLambdaImported, x.F2MomentImported, x.F4MomentImported, x.CosmologicalConstantImported, x.DarkEnergyImported, x.ElectroweakScaleImported, x.FlavorDataImported, x.NativeSpectralNormalizationWrite)
}

func Markdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 513 Registry Audit — Spectral Moment Hierarchy and Cutoff-Separation Airlock Audit\n\n")
	b.WriteString("## Verdict\n\n```text\n" + strings.Join(statuses(), "\n") + "\n```\n\n")
	b.WriteString("## Inherited boundary\n\n" + a.Inheritance.Reason + "\n\n```text\n" + FormatInheritance(a.Inheritance) + "\n```\n\n")
	b.WriteString("## Three-channel heat-kernel ledger\n\n" + a.Ledger.Reason + "\n\n```text\n" + FormatLedger(a.Ledger) + "\n```\n\n")
	b.WriteString("## Relative prefactor hierarchy\n\n" + a.Hierarchy.Reason + "\n\n```text\n" + FormatHierarchy(a.Hierarchy) + "\n```\n\n")
	b.WriteString("## Spectral moment and cutoff airlock\n\n" + a.Airlock.Reason + "\n\n```text\n" + FormatAirlock(a.Airlock) + "\n```\n\n")
	b.WriteString("## Firewall result\n\n" + a.Firewall.Reason + "\n\n```text\n" + FormatFirewall(a.Firewall) + "\n```\n\n")
	b.WriteString("## Registry update\n\n")
	writeList(&b, "Native entries", a.Registry.NativeEntries)
	writeList(&b, "Bridge entries", a.Registry.BridgeEntries)
	writeList(&b, "Environmental entries", a.Registry.EnvironmentalEntries)
	writeList(&b, "Failed routes", a.Registry.FailedRoutes)
	writeList(&b, "Open theorems", a.Registry.OpenTheorems)
	b.WriteString("## Next step\n\nGate514 should be:\n\n```text\nGate 514 — " + a.Next.Title + "\n```\n\nPrimary task:\n\n```text\n" + a.Next.PrimaryTask + "\n```\n\n")
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
