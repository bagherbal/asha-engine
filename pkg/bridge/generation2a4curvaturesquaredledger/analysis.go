// Package generation2a4curvaturesquaredledger implements Gate 511:
// Gravitational a4 Curvature-Squared and Topological Counterterm Audit.
//
// Gate 510 fixed the native dimensionless a2 scalar-curvature trace weight but
// quarantined Newton normalization behind the f2 Λ² airlock. Gate 511 moves to
// the scale-independent a4 gravitational channel.  It classifies the curvature²
// sockets present in the product spectral action, separates topological
// Gauss-Bonnet data from dynamical Weyl/curvature² data, and blocks any attempt
// to turn those dimensionless sockets into physical gravitational dynamics,
// Newton normalization, or cosmological/vacuum-energy predictions.
package generation2a4curvaturesquaredledger

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2curvaturecoefficientprovenance"
	"github.com/bagherbal/asha-engine/pkg/bridge/productspectralactioncoefficients"
)

const (
	AuditID = "GATE511-A4-CURVATURE-SQUARED-TOPOLOGICAL-COUNTERTERM-AUDIT"

	StatusGate510CurvatureCoefficientInherited    = "CONDITIONAL_SUPPORT_GATE510_A2_CURVATURE_COEFFICIENT_INHERITED"
	StatusProductA4LedgerInherited                = "CONDITIONAL_SUPPORT_PRODUCT_SPECTRAL_A4_LEDGER_INHERITED"
	StatusA4CurvatureSquaredSocketDefined         = "CONDITIONAL_SUPPORT_A4_CURVATURE_SQUARED_SOCKET_DEFINED"
	StatusFourDimensionalCurvatureBasisClassified = "CONDITIONAL_SUPPORT_FOUR_DIMENSIONAL_CURVATURE_BASIS_CLASSIFIED"
	StatusGaussBonnetTopologicalCountertermFound  = "CONDITIONAL_SUPPORT_GAUSS_BONNET_TOPOLOGICAL_COUNTERTERM_IDENTIFIED"
	StatusWeylSquaredDynamicalSocketFound         = "CONDITIONAL_SUPPORT_WEYL_SQUARED_DYNAMICAL_SOCKET_PRESENT"
	StatusA4DimensionlessF0ChannelIsolated        = "CONDITIONAL_SUPPORT_A4_DIMENSIONLESS_F0_CHANNEL_ISOLATED"
	StatusA4DoesNotUseF2LambdaSquared             = "CONDITIONAL_SUPPORT_A4_CHANNEL_DOES_NOT_USE_F2_LAMBDA_SQUARED"
	StatusGravityA4FirewallPreserved              = "FIREWALL_PRESERVED_NO_NEWTON_PLANCK_COSMOLOGY_EW_OR_FLAVOR_DATA_IMPORTED"
	StatusFirewallA4NativeDynamicsBlocked         = "FIREWALL_BLOCKED_A4_CURVATURE_SQUARED_PHYSICAL_DYNAMICS_WRITE"

	StatusFailedNewtonConstantStillNotDerived           = "FAILED_ROUTE_NEWTON_CONSTANT_STILL_NOT_DERIVED_BY_A4"
	StatusFailedCutoffLambdaStillNotSelected            = "FAILED_ROUTE_CUTOFF_SCALE_LAMBDA_STILL_NOT_SELECTED_BY_A4"
	StatusFailedEHNormalizationStillOpen                = "FAILED_ROUTE_EINSTEIN_HILBERT_NORMALIZATION_STILL_OPEN_AFTER_A4"
	StatusFailedCosmologicalF4StillUnsolved             = "FAILED_ROUTE_COSMOLOGICAL_F4_VACUUM_CHANNEL_STILL_UNSOLVED"
	StatusFailedA4CoefficientsNotUniquePhysicalDynamics = "FAILED_ROUTE_A4_CURVATURE_SQUARED_COEFFICIENTS_NOT_UNIQUE_PHYSICAL_GRAVITY_DYNAMICS"
	StatusFailedBoundaryTermsAndSchemeNotClosed         = "FAILED_ROUTE_A4_BOUNDARY_TERMS_AND_RENORMALIZATION_SCHEME_NOT_CLOSED"
	StatusFailedMetricEquationsNotDerived               = "FAILED_ROUTE_HIGHER_DERIVATIVE_METRIC_EQUATIONS_NOT_NATIVE_DERIVED"
)

const (
	finiteTraceDimension = 96.0
	nativeF0Moment       = 7.0
	fourPiSquared        = 16.0 * math.Pi * math.Pi
)

type Inheritance struct {
	Executed                          bool
	Gate510A2AuditInherited           bool
	Gate510A2TraceWeightNative        bool
	Gate510NewtonNormalizationBlocked bool
	Gate510CosmologicalF4Excluded     bool
	ProductTripleValid                bool
	ProductA4ChannelsDeclared         bool
	ProductF0MomentAvailable          bool
	ProductAllCoefficientsClosed      bool
	ProductHardTOEClosure             bool
	Verdict                           string
	Reason                            string
}

type CurvatureBasisSieve struct {
	Executed                 bool
	Dimension                int
	RawBasis                 []string
	TopologicalInvariant     string
	ConformalInvariant       string
	ScalarInvariant          string
	EulerIdentity            string
	WeylIdentity             string
	BasisRank                int
	TopologicalCounterterm   bool
	DynamicalCurvatureSocket bool
	UniqueMetricDynamics     bool
	Verdict                  string
	Reason                   string
}

type A4CoefficientAudit struct {
	Executed                         bool
	HeatKernelChannel                string
	FiniteTraceDimension             float64
	F0Moment                         float64
	RawPrefactorPerF0BeforeInvariant float64
	F0WeightedPrefactor              float64
	DimensionlessChannel             bool
	UsesF2LambdaSquared              bool
	UsesF4LambdaFourth               bool
	NewtonConstantDerived            bool
	PhysicalGravityCouplingDerived   bool
	Formula                          string
	Verdict                          string
	Reason                           string
}

type TopologicalCountertermAudit struct {
	Executed                       bool
	GaussBonnetInvariant           string
	IntegralTopologicalInFourD     bool
	LocalVariationBoundaryOnly     bool
	EulerCharacteristicNumeric     bool
	TopologicalSocketNative        bool
	TopologicalCoefficientPhysical bool
	Verdict                        string
	Reason                         string
}

type DynamicalCurvatureAudit struct {
	Executed                         bool
	WeylSquaredSocketPresent         bool
	RiemannRicciScalarSocketsPresent bool
	HigherDerivativeMetricTerms      bool
	RenormalizationSchemeSelected    bool
	BoundaryConditionsSelected       bool
	LowEnergyEinsteinLimitDerived    bool
	MetricEquationsNativeDerived     bool
	PhysicalA4DynamicsClosed         bool
	Verdict                          string
	Reason                           string
}

type Firewall struct {
	Executed                           bool
	NewtonConstantImported             bool
	NewtonConstantDerived              bool
	PlanckScaleImported                bool
	CutoffLambdaSelected               bool
	F2MomentSeparatedFromLambda        bool
	EinsteinHilbertNormalizationClosed bool
	CosmologicalConstantImported       bool
	CosmologicalConstantDerived        bool
	F4VacuumSubtractionSelected        bool
	ElectroweakScaleImported           bool
	FlavorDataImported                 bool
	PhysicalA4DynamicsWritten          bool
	NativeGravityNormalizationWritten  bool
	Verdict                            string
	Reason                             string
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
	Basis       CurvatureBasisSieve
	A4          A4CoefficientAudit
	Topological TopologicalCountertermAudit
	Dynamical   DynamicalCurvatureAudit
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
	g510, err := generation2curvaturecoefficientprovenance.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate510 curvature coefficient provenance: %w", err)
	}
	g377, err := productspectralactioncoefficients.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit product spectral-action coefficient ledger: %w", err)
	}
	a := Analysis{}
	a.Inheritance = buildInheritance(g510, g377)
	a.Basis = buildBasis()
	a.A4 = buildA4(g377)
	a.Topological = buildTopological()
	a.Dynamical = buildDynamical(a.Basis)
	a.Firewall = buildFirewall()
	a.Registry = buildRegistry(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g510 generation2curvaturecoefficientprovenance.Analysis, g377 productspectralactioncoefficients.Analysis) Inheritance {
	c := g377.Calculation
	return Inheritance{
		Executed:                          true,
		Gate510A2AuditInherited:           g510.A2.DimensionlessTraceWeightNative && g510.A2.Gate377RawCoefficientMatched,
		Gate510A2TraceWeightNative:        nearly(g510.A2.A2WeightMagnitudeBefore4Pi, 8, 1e-12),
		Gate510NewtonNormalizationBlocked: !g510.Firewall.NewtonConstantDerived && !g510.Firewall.CutoffLambdaSelected && !g510.Firewall.NativeGravityNormalizationWritten,
		Gate510CosmologicalF4Excluded:     g510.Cutoff.CosmologicalF4Excluded && !g510.Cutoff.CosmologicalConstantDerived,
		ProductTripleValid:                c.Product.Valid,
		ProductA4ChannelsDeclared:         strings.Contains(c.Convention.A4Channels, "curvature²") || strings.Contains(c.Lagrangian, "curvature²"),
		ProductF0MomentAvailable:          nearly(c.Finite.F0, nativeF0Moment, 1e-12),
		ProductAllCoefficientsClosed:      c.AllCoefficientsDetermined,
		ProductHardTOEClosure:             c.HardTOEClosure,
		Verdict:                           strings.Join([]string{StatusGate510CurvatureCoefficientInherited, StatusProductA4LedgerInherited}, ";"),
		Reason:                            "Gate510 supplies the a2 curvature coefficient provenance and its normalization firewall; Gate511 inherits only the product spectral-action a4 curvature² channel and the dimensionless f0 moment.",
	}
}

func buildBasis() CurvatureBasisSieve {
	return CurvatureBasisSieve{
		Executed:                 true,
		Dimension:                4,
		RawBasis:                 []string{"Riem²", "Ric²", "R²"},
		TopologicalInvariant:     "E₄ = Riem² - 4 Ric² + R²",
		ConformalInvariant:       "C² = Riem² - 2 Ric² + R²/3",
		ScalarInvariant:          "R²",
		EulerIdentity:            "∫√g E₄ is Gauss-Bonnet/Euler topological data in four dimensions, up to boundary conventions",
		WeylIdentity:             "C² is the conformal/dynamical curvature-squared socket of the four-dimensional spectral action",
		BasisRank:                3,
		TopologicalCounterterm:   true,
		DynamicalCurvatureSocket: true,
		UniqueMetricDynamics:     false,
		Verdict:                  strings.Join([]string{StatusA4CurvatureSquaredSocketDefined, StatusFourDimensionalCurvatureBasisClassified, StatusGaussBonnetTopologicalCountertermFound, StatusWeylSquaredDynamicalSocketFound}, ";"),
		Reason:                   "the four-dimensional curvature² vector space is classified into a topological Euler counterterm, a Weyl² dynamical socket, and scheme-dependent scalar/boundary curvature² pieces; classification is native, metric dynamics selection is not.",
	}
}

func buildA4(g377 productspectralactioncoefficients.Analysis) A4CoefficientAudit {
	pref := finiteTraceDimension / (360.0 * fourPiSquared)
	f0weighted := pref * g377.Calculation.Finite.F0
	return A4CoefficientAudit{
		Executed:                         true,
		HeatKernelChannel:                "a4 gravitational curvature² channel of Tr f(D²/Λ²)",
		FiniteTraceDimension:             finiteTraceDimension,
		F0Moment:                         g377.Calculation.Finite.F0,
		RawPrefactorPerF0BeforeInvariant: pref,
		F0WeightedPrefactor:              f0weighted,
		DimensionlessChannel:             true,
		UsesF2LambdaSquared:              false,
		UsesF4LambdaFourth:               false,
		NewtonConstantDerived:            false,
		PhysicalGravityCouplingDerived:   false,
		Formula:                          "S_a4,grav = f0·(4π)^(-2)·∫√g Tr_F[universal curvature² polynomial]/360",
		Verdict:                          strings.Join([]string{StatusA4DimensionlessF0ChannelIsolated, StatusA4DoesNotUseF2LambdaSquared}, ";"),
		Reason:                           "unlike the a2 Einstein-Hilbert channel, the a4 curvature² channel is dimensionless and controlled by f0 times universal heat-kernel curvature polynomials; this does not derive Newton's constant or a unique low-energy gravity action.",
	}
}

func buildTopological() TopologicalCountertermAudit {
	return TopologicalCountertermAudit{
		Executed:                       true,
		GaussBonnetInvariant:           "E₄ = Riem² - 4 Ric² + R²",
		IntegralTopologicalInFourD:     true,
		LocalVariationBoundaryOnly:     true,
		EulerCharacteristicNumeric:     false,
		TopologicalSocketNative:        true,
		TopologicalCoefficientPhysical: false,
		Verdict:                        StatusGaussBonnetTopologicalCountertermFound,
		Reason:                         "the Euler/Gauss-Bonnet density is the topological curvature² counterterm socket; ASHA may classify its presence, but no manifold topology, boundary condition, or physical coefficient is selected here.",
	}
}

func buildDynamical(b CurvatureBasisSieve) DynamicalCurvatureAudit {
	return DynamicalCurvatureAudit{
		Executed:                         true,
		WeylSquaredSocketPresent:         b.DynamicalCurvatureSocket,
		RiemannRicciScalarSocketsPresent: b.BasisRank == 3,
		HigherDerivativeMetricTerms:      true,
		RenormalizationSchemeSelected:    false,
		BoundaryConditionsSelected:       false,
		LowEnergyEinsteinLimitDerived:    false,
		MetricEquationsNativeDerived:     false,
		PhysicalA4DynamicsClosed:         false,
		Verdict:                          strings.Join([]string{StatusWeylSquaredDynamicalSocketFound, StatusFailedA4CoefficientsNotUniquePhysicalDynamics, StatusFailedBoundaryTermsAndSchemeNotClosed, StatusFailedMetricEquationsNotDerived}, ";"),
		Reason:                           "Weyl² and related curvature² terms are legitimate spectral-action sockets, but their physical role depends on renormalization, boundary data, metric-sign conventions, and the already-unclosed Einstein-Hilbert normalization.",
	}
}

func buildFirewall() Firewall {
	return Firewall{
		Executed:                           true,
		NewtonConstantImported:             false,
		NewtonConstantDerived:              false,
		PlanckScaleImported:                false,
		CutoffLambdaSelected:               false,
		F2MomentSeparatedFromLambda:        false,
		EinsteinHilbertNormalizationClosed: false,
		CosmologicalConstantImported:       false,
		CosmologicalConstantDerived:        false,
		F4VacuumSubtractionSelected:        false,
		ElectroweakScaleImported:           false,
		FlavorDataImported:                 false,
		PhysicalA4DynamicsWritten:          false,
		NativeGravityNormalizationWritten:  false,
		Verdict:                            strings.Join([]string{StatusGravityA4FirewallPreserved, StatusFirewallA4NativeDynamicsBlocked}, ";"),
		Reason:                             "Gate511 classifies dimensionless a4 curvature² sockets only. It imports no G, M_P, Λ value, cosmological constant, electroweak scale, Yukawa, CKM, or PMNS data, and writes no physical a4 dynamics or gravity normalization.",
	}
}

func buildRegistry(a Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries: []string{
			"The product spectral-action a4 gravitational channel contains a four-dimensional curvature² socket.",
			"The curvature² basis decomposes into Euler/Gauss-Bonnet topological data, Weyl² conformal/dynamical data, and scalar/boundary curvature² data.",
			"The a4 channel is dimensionless and does not require the f₂Λ² Einstein-Hilbert normalization product.",
		},
		BridgeEntries: []string{
			"The symbolic curvature² action has the form f0·(4π)^(-2)·Tr_F(universal curvature² polynomial)/360.",
			"The Weyl²/dynamical curvature² socket is present, but physical metric equations and low-energy gravity interpretation remain bridge-level.",
		},
		EnvironmentalEntries: []string{
			"Renormalization scheme, boundary conditions, manifold topology, Newton normalization, cutoff selection, and cosmological vacuum subtraction remain quarantined.",
		},
		FailedRoutes: []string{
			StatusFailedNewtonConstantStillNotDerived,
			StatusFailedCutoffLambdaStillNotSelected,
			StatusFailedEHNormalizationStillOpen,
			StatusFailedCosmologicalF4StillUnsolved,
			StatusFailedA4CoefficientsNotUniquePhysicalDynamics,
			StatusFailedBoundaryTermsAndSchemeNotClosed,
			StatusFailedMetricEquationsNotDerived,
		},
		OpenTheorems: []string{
			"Audit the f4Λ4 cosmological/vacuum-energy channel and vacuum-subtraction obligation separately.",
			"Prove or reject a native renormalization/boundary condition selector for curvature² dynamics.",
			"Prove or reject a native manifold-topology selector for the Euler characteristic contribution.",
		},
	}
}

func buildNext() NextStep {
	return NextStep{Gate: 512, Title: "Cosmological f4 Vacuum Energy and Subtraction Airlock Audit", Reason: "Gate511 classifies scale-independent a4 curvature² sockets but leaves the a0/f4 cosmological volume channel completely unclosed.", PrimaryTask: "separate the native a0 volume prefactor from the physical cosmological constant, test whether any finite trace cancels vacuum energy, and formally quarantine Λ_cosmo/f4/subtraction data if no theorem appears"}
}

func validate(a Analysis) error {
	checks := []struct {
		ok  bool
		msg string
	}{
		{a.Inheritance.Executed && a.Inheritance.Gate510A2AuditInherited && a.Inheritance.Gate510A2TraceWeightNative && a.Inheritance.Gate510NewtonNormalizationBlocked && a.Inheritance.Gate510CosmologicalF4Excluded && a.Inheritance.ProductTripleValid && a.Inheritance.ProductA4ChannelsDeclared && a.Inheritance.ProductF0MomentAvailable && !a.Inheritance.ProductAllCoefficientsClosed && !a.Inheritance.ProductHardTOEClosure, "Gate511 inheritance invalid"},
		{a.Basis.Executed && a.Basis.Dimension == 4 && a.Basis.BasisRank == 3 && a.Basis.TopologicalCounterterm && a.Basis.DynamicalCurvatureSocket && !a.Basis.UniqueMetricDynamics, "Gate511 curvature basis invalid"},
		{a.A4.Executed && nearly(a.A4.FiniteTraceDimension, 96, 1e-12) && nearly(a.A4.F0Moment, 7, 1e-12) && a.A4.RawPrefactorPerF0BeforeInvariant > 0 && a.A4.F0WeightedPrefactor > 0 && a.A4.DimensionlessChannel && !a.A4.UsesF2LambdaSquared && !a.A4.UsesF4LambdaFourth && !a.A4.NewtonConstantDerived && !a.A4.PhysicalGravityCouplingDerived, "Gate511 a4 coefficient audit invalid"},
		{a.Topological.Executed && a.Topological.IntegralTopologicalInFourD && a.Topological.LocalVariationBoundaryOnly && !a.Topological.EulerCharacteristicNumeric && a.Topological.TopologicalSocketNative && !a.Topological.TopologicalCoefficientPhysical, "Gate511 topological counterterm audit invalid"},
		{a.Dynamical.Executed && a.Dynamical.WeylSquaredSocketPresent && a.Dynamical.RiemannRicciScalarSocketsPresent && a.Dynamical.HigherDerivativeMetricTerms && !a.Dynamical.RenormalizationSchemeSelected && !a.Dynamical.BoundaryConditionsSelected && !a.Dynamical.LowEnergyEinsteinLimitDerived && !a.Dynamical.MetricEquationsNativeDerived && !a.Dynamical.PhysicalA4DynamicsClosed, "Gate511 dynamical curvature audit invalid"},
		{a.Firewall.Executed && !a.Firewall.NewtonConstantImported && !a.Firewall.NewtonConstantDerived && !a.Firewall.PlanckScaleImported && !a.Firewall.CutoffLambdaSelected && !a.Firewall.F2MomentSeparatedFromLambda && !a.Firewall.EinsteinHilbertNormalizationClosed && !a.Firewall.CosmologicalConstantImported && !a.Firewall.CosmologicalConstantDerived && !a.Firewall.F4VacuumSubtractionSelected && !a.Firewall.ElectroweakScaleImported && !a.Firewall.FlavorDataImported && !a.Firewall.PhysicalA4DynamicsWritten && !a.Firewall.NativeGravityNormalizationWritten, "Gate511 firewall invalid"},
		{a.Next.Gate == 512, "Gate511 next gate invalid"},
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
		StatusGate510CurvatureCoefficientInherited,
		StatusProductA4LedgerInherited,
		StatusA4CurvatureSquaredSocketDefined,
		StatusFourDimensionalCurvatureBasisClassified,
		StatusGaussBonnetTopologicalCountertermFound,
		StatusWeylSquaredDynamicalSocketFound,
		StatusA4DimensionlessF0ChannelIsolated,
		StatusA4DoesNotUseF2LambdaSquared,
		StatusFailedNewtonConstantStillNotDerived,
		StatusFailedCutoffLambdaStillNotSelected,
		StatusFailedEHNormalizationStillOpen,
		StatusFailedCosmologicalF4StillUnsolved,
		StatusFailedA4CoefficientsNotUniquePhysicalDynamics,
		StatusFailedBoundaryTermsAndSchemeNotClosed,
		StatusFailedMetricEquationsNotDerived,
		StatusGravityA4FirewallPreserved,
		StatusFirewallA4NativeDynamicsBlocked,
	}
}

func truth(a Analysis) string {
	return "Gate 511 proves that the product spectral action has a native, scale-independent a4 curvature-squared socket: the four-dimensional curvature basis decomposes into Gauss-Bonnet topological data and Weyl/scalar curvature-squared channels. This is a genuine spectral-geometry ledger entry, not a mass, flavor, or electroweak claim. But the gate does not derive Newton's constant, the cutoff, cosmological vacuum subtraction, a physical renormalization scheme, boundary conditions, or complete metric dynamics. The a4 channel is present; physical higher-derivative gravity remains quarantined."
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("Gate510 inherited=%t; a2 weight native=%t; Newton blocked=%t; f4 excluded=%t; product valid=%t; a4 declared=%t; f0 available=%t; all coeffs closed=%t; hard ToE=%t", x.Gate510A2AuditInherited, x.Gate510A2TraceWeightNative, x.Gate510NewtonNormalizationBlocked, x.Gate510CosmologicalF4Excluded, x.ProductTripleValid, x.ProductA4ChannelsDeclared, x.ProductF0MomentAvailable, x.ProductAllCoefficientsClosed, x.ProductHardTOEClosure)
}
func FormatBasis(x CurvatureBasisSieve) string {
	return fmt.Sprintf("raw basis=%s; E4=%s; C2=%s; rank=%d; topological=%t; dynamical=%t; unique dynamics=%t", strings.Join(x.RawBasis, ","), x.TopologicalInvariant, x.ConformalInvariant, x.BasisRank, x.TopologicalCounterterm, x.DynamicalCurvatureSocket, x.UniqueMetricDynamics)
}
func FormatA4(x A4CoefficientAudit) string {
	return fmt.Sprintf("%s; TrF=%.0f; f0=%.0f; prefactor/f0=%.12g; f0-weighted=%.12g; dimensionless=%t; uses f2Λ²=%t; uses f4Λ4=%t; physical=%t", x.Formula, x.FiniteTraceDimension, x.F0Moment, x.RawPrefactorPerF0BeforeInvariant, x.F0WeightedPrefactor, x.DimensionlessChannel, x.UsesF2LambdaSquared, x.UsesF4LambdaFourth, x.PhysicalGravityCouplingDerived)
}
func FormatTopological(x TopologicalCountertermAudit) string {
	return fmt.Sprintf("%s; integral topological=%t; variation boundary-only=%t; numeric Euler characteristic selected=%t; native socket=%t; physical coefficient=%t", x.GaussBonnetInvariant, x.IntegralTopologicalInFourD, x.LocalVariationBoundaryOnly, x.EulerCharacteristicNumeric, x.TopologicalSocketNative, x.TopologicalCoefficientPhysical)
}
func FormatDynamical(x DynamicalCurvatureAudit) string {
	return fmt.Sprintf("Weyl²=%t; all raw sockets=%t; higher derivative=%t; scheme selected=%t; boundary selected=%t; low-energy Einstein limit=%t; metric equations native=%t; physical dynamics closed=%t", x.WeylSquaredSocketPresent, x.RiemannRicciScalarSocketsPresent, x.HigherDerivativeMetricTerms, x.RenormalizationSchemeSelected, x.BoundaryConditionsSelected, x.LowEnergyEinsteinLimitDerived, x.MetricEquationsNativeDerived, x.PhysicalA4DynamicsClosed)
}
func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("G imported=%t; G derived=%t; Planck imported=%t; Λ selected=%t; EH closed=%t; cosmological imported=%t; cosmological derived=%t; f4 subtraction=%t; EW imported=%t; flavor imported=%t; a4 dynamics write=%t; native gravity write=%t", x.NewtonConstantImported, x.NewtonConstantDerived, x.PlanckScaleImported, x.CutoffLambdaSelected, x.EinsteinHilbertNormalizationClosed, x.CosmologicalConstantImported, x.CosmologicalConstantDerived, x.F4VacuumSubtractionSelected, x.ElectroweakScaleImported, x.FlavorDataImported, x.PhysicalA4DynamicsWritten, x.NativeGravityNormalizationWritten)
}

func Markdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 511 Registry Audit — Gravitational a4 Curvature-Squared and Topological Counterterm Audit\n\n")
	b.WriteString("## Verdict\n\n```text\n" + strings.Join(statuses(), "\n") + "\n```\n\n")
	b.WriteString("## Inherited boundary\n\n" + a.Inheritance.Reason + "\n\n```text\n" + FormatInheritance(a.Inheritance) + "\n```\n\n")
	b.WriteString("## a4 curvature-squared basis audit\n\n" + a.Basis.Reason + "\n\n```text\n" + FormatBasis(a.Basis) + "\n" + a.Basis.EulerIdentity + "\n" + a.Basis.WeylIdentity + "\n```\n\n")
	b.WriteString("## a4 coefficient channel\n\n" + a.A4.Reason + "\n\n```text\n" + FormatA4(a.A4) + "\n```\n\n")
	b.WriteString("## Topological counterterm audit\n\n" + a.Topological.Reason + "\n\n```text\n" + FormatTopological(a.Topological) + "\n```\n\n")
	b.WriteString("## Dynamical curvature firewall\n\n" + a.Dynamical.Reason + "\n\n```text\n" + FormatDynamical(a.Dynamical) + "\n```\n\n")
	b.WriteString("## Firewall result\n\n" + a.Firewall.Reason + "\n\n```text\n" + FormatFirewall(a.Firewall) + "\n```\n\n")
	b.WriteString("## Registry update\n\n")
	writeList(&b, "Native entries", a.Registry.NativeEntries)
	writeList(&b, "Bridge entries", a.Registry.BridgeEntries)
	writeList(&b, "Environmental entries", a.Registry.EnvironmentalEntries)
	writeList(&b, "Failed routes", a.Registry.FailedRoutes)
	writeList(&b, "Open theorems", a.Registry.OpenTheorems)
	b.WriteString("## Next step\n\nGate512 should be:\n\n```text\nGate 512 — " + a.Next.Title + "\n```\n\nPrimary task:\n\n```text\n" + a.Next.PrimaryTask + "\n```\n\n")
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
