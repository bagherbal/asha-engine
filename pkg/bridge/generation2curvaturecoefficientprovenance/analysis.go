// Package generation2curvaturecoefficientprovenance implements Gate 510:
// Curvature Coefficient Provenance and Heat-Kernel Trace Convention Audit.
//
// Gate 509 identified the Einstein-Hilbert scalar-curvature socket but refused
// to normalize Newton's constant. Gate 510 audits the exact symbolic source of
// that boundary.  It separates the native dimensionless trace arithmetic in the
// Dirac-square / heat-kernel a2 channel from the dimensionful bridge product
// f2 Λ² and the still-unselected cutoff/Planck normalization.
package generation2curvaturecoefficientprovenance

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2topologicalgravityredirect"
	"github.com/bagherbal/asha-engine/pkg/bridge/productspectralactioncoefficients"
)

const (
	AuditID = "GATE510-CURVATURE-COEFFICIENT-PROVENANCE-HEAT-KERNEL-TRACE-CONVENTION-AUDIT"

	StatusGate509Inherited                       = "CONDITIONAL_SUPPORT_GATE509_GRAVITY_SOCKET_INHERITED"
	StatusProductHeatKernelConventionInherited   = "CONDITIONAL_SUPPORT_PRODUCT_HEAT_KERNEL_CONVENTION_INHERITED"
	StatusD2EndomorphismSieveExecuted            = "CONDITIONAL_SUPPORT_D2_ENDOMORPHISM_SIEVE_EXECUTED"
	StatusLichnerowiczCurvatureTermAudited       = "CONDITIONAL_SUPPORT_LICHNEROWICZ_CURVATURE_ENDOMORPHISM_AUDITED"
	StatusA2TraceWeightComputed                  = "CONDITIONAL_SUPPORT_A2_TRACE_WEIGHT_COMPUTED"
	StatusFiniteTraceWeightNative                = "CONDITIONAL_SUPPORT_FINITE_TRACE_DIMENSIONLESS_WEIGHT_NATIVE"
	StatusRawGate377CoefficientMatched           = "CONDITIONAL_SUPPORT_GATE377_RAW_A2_CHANNEL_MATCHED"
	StatusF2LambdaObligationIsolated             = "CONDITIONAL_SUPPORT_F2_LAMBDA_SQUARED_OBLIGATION_ISOLATED"
	StatusGravityNormalizationQuarantined        = "CONDITIONAL_SUPPORT_GRAVITY_NORMALIZATION_QUARANTINED"
	StatusNoEmpiricalGravityDataImported         = "CONDITIONAL_SUPPORT_NO_EMPIRICAL_GRAVITY_OR_COSMOLOGY_DATA_IMPORTED"
	StatusNextA4CurvatureGateDefined             = "CONDITIONAL_SUPPORT_GATE511_A4_CURVATURE_SQUARED_AUDIT_DEFINED"
	StatusFirewallPreserved                      = "FIREWALL_PRESERVED_NO_NEWTON_PLANCK_COSMOLOGY_OR_EW_DATA_IMPORTED"
	StatusFirewallNativeGravityNormalizationStop = "FIREWALL_BLOCKED_NEWTON_NORMALIZATION_NATIVE_WRITE"

	StatusFailedNewtonConstantStillNotDerived      = "FAILED_ROUTE_NEWTON_CONSTANT_STILL_NOT_DERIVED"
	StatusFailedCutoffLambdaStillNotSelected       = "FAILED_ROUTE_CUTOFF_SCALE_LAMBDA_STILL_NOT_SELECTED"
	StatusFailedF2MomentStillNotSeparated          = "FAILED_ROUTE_F2_MOMENT_STILL_NOT_SEPARATED_FROM_LAMBDA"
	StatusFailedEHNormalizationStillOpen           = "FAILED_ROUTE_EINSTEIN_HILBERT_NORMALIZATION_STILL_OPEN"
	StatusFailedTraceConventionNotUniquelySelected = "FAILED_ROUTE_HEAT_KERNEL_TRACE_CONVENTION_NOT_UNIQUELY_SELECTED"
	StatusFailedCosmologicalF4Excluded             = "FAILED_ROUTE_COSMOLOGICAL_F4_CHANNEL_EXCLUDED_FROM_GATE510"
	StatusFailedPhysicalMetricDynamicsNotDerived   = "FAILED_ROUTE_PHYSICAL_METRIC_DYNAMICS_NOT_DERIVED"
)

const (
	finiteTraceDimension = 96.0
	universalA2RPart     = 1.0 / 6.0
	diracEndomorphismR   = -1.0 / 4.0
	combinedDiracA2R     = universalA2RPart + diracEndomorphismR // -1/12 in P=-(∇²+E) convention.
)

type Inheritance struct {
	Executed                       bool
	Gate509GravitySocketInherited  bool
	Gate509NormalizationBlocked    bool
	Gate509NoEmpiricalDataImported bool
	ProductTripleValid             bool
	HeatKernelConventionDeclared   bool
	Gate377RawA2ChannelPresent     bool
	Gate377SkeletonChannelPresent  bool
	Gate377AllCoefficientsClosed   bool
	Gate377HardTOEClosure          bool
	Verdict                        string
	Reason                         string
}

type D2EndomorphismSieve struct {
	Executed                      bool
	OperatorConvention            string
	LichnerowiczFormula           string
	UniversalA2RPart              float64
	DiracEndomorphismRPart        float64
	CombinedA2RPart               float64
	CombinedA2RPartMagnitude      float64
	FiniteDiracPartAddsCurvature  bool
	CurvatureEndomorphismAudited  bool
	SignConventionClosed          bool
	PhysicalMetricDynamicsDerived bool
	Verdict                       string
	Reason                        string
}

type A2TraceEvaluation struct {
	Executed                          bool
	FiniteHilbertTraceDimension       float64
	A2WeightMagnitudeBefore4Pi        float64
	RawDensityCoefficientPerF2Lambda2 float64
	Gate377RawDensityCoefficient      float64
	Gate377RawCoefficientMatched      bool
	DimensionlessTraceWeightNative    bool
	IncludesCutoffMoment              bool
	PhysicalCoefficientNative         bool
	Formula                           string
	Verdict                           string
	Reason                            string
}

type TraceConventionAudit struct {
	Executed                        bool
	RawConventionDeclared           bool
	SkeletonConventionDeclared      bool
	RawCoefficientPerF2Lambda2      float64
	SkeletonCoefficientPerF2Lambda2 float64
	RawSkeletonNumericallyDifferent bool
	UniqueTraceConventionSelected   bool
	CanPromoteEitherToNewtonNative  bool
	Verdict                         string
	Reason                          string
}

type CutoffProvenance struct {
	Executed                           bool
	EinsteinHilbertSymbolicCoefficient string
	RequiresF2LambdaSquaredProduct     bool
	F2LambdaProductNativeOnlyAsSymbol  bool
	F2MomentSeparatedFromLambda        bool
	CutoffLambdaSelected               bool
	NewtonConstantDerived              bool
	PlanckScaleImported                bool
	CosmologicalConstantDerived        bool
	CosmologicalF4Excluded             bool
	GravityNormalizationBridgeOnly     bool
	Verdict                            string
	Reason                             string
}

type Firewall struct {
	Executed                           bool
	NewtonConstantImported             bool
	NewtonConstantDerived              bool
	PlanckMassImported                 bool
	CutoffLambdaSelected               bool
	F2MomentSeparatedFromLambda        bool
	EinsteinHilbertNormalizationClosed bool
	CosmologicalConstantImported       bool
	CosmologicalConstantDerived        bool
	ElectroweakScaleImported           bool
	FlavorDataImported                 bool
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
	Inheritance  Inheritance
	Endomorphism D2EndomorphismSieve
	A2           A2TraceEvaluation
	Convention   TraceConventionAudit
	Cutoff       CutoffProvenance
	Firewall     Firewall
	Registry     RegistryUpdate
	Next         NextStep
	Truth        string
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
	g509, err := generation2topologicalgravityredirect.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate509 topological/gravity redirect: %w", err)
	}
	g377, err := productspectralactioncoefficients.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate377 product spectral-action coefficients: %w", err)
	}

	a := Analysis{}
	a.Inheritance = buildInheritance(g509, g377)
	a.Endomorphism = buildEndomorphism()
	a.A2 = buildA2(g377, a.Endomorphism)
	a.Convention = buildConvention(g377, a.A2)
	a.Cutoff = buildCutoff(a.A2)
	a.Firewall = buildFirewall()
	a.Registry = buildRegistry(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g509 generation2topologicalgravityredirect.Analysis, g377 productspectralactioncoefficients.Analysis) Inheritance {
	c := g377.Calculation
	return Inheritance{
		Executed:                       true,
		Gate509GravitySocketInherited:  g509.Gravity.EinsteinHilbertSocketPresent && g509.Gravity.SMGravityStructuralRecovered,
		Gate509NormalizationBlocked:    !g509.Firewall.NewtonConstantDerived && !g509.Firewall.CutoffLambdaSelected && !g509.Firewall.F2MomentSeparatedFromLambda && !g509.Firewall.EinsteinHilbertNormalizationClosed && !g509.Firewall.NativeGravityRegistryWritten,
		Gate509NoEmpiricalDataImported: !g509.Firewall.NewtonConstantImported && !g509.Firewall.PlanckScaleImported && !g509.Firewall.CosmologicalScaleImported,
		ProductTripleValid:             c.Product.Valid,
		HeatKernelConventionDeclared:   c.Convention.Dimension == 4 && c.Convention.Expansion != "" && c.Convention.A2DiracRDensity != "",
		Gate377RawA2ChannelPresent:     c.A2RawEinsteinCoefficientPerMP2.Numeric > 0,
		Gate377SkeletonChannelPresent:  c.A2SkeletonEinsteinCoefficientPerMP2.Numeric > 0,
		Gate377AllCoefficientsClosed:   c.AllCoefficientsDetermined,
		Gate377HardTOEClosure:          c.HardTOEClosure,
		Verdict:                        strings.Join([]string{StatusGate509Inherited, StatusProductHeatKernelConventionInherited}, ";"),
		Reason:                         "Gate509 supplied the structural Einstein-Hilbert socket but blocked normalization; Gate510 inherits only the symbolic product heat-kernel ledger and no empirical gravity/cosmology values.",
	}
}

func buildEndomorphism() D2EndomorphismSieve {
	return D2EndomorphismSieve{
		Executed:                      true,
		OperatorConvention:            "Laplace-type convention P=-(∇²+E); sign flips under alternate Dirac/Lorentzian conventions are tracked, not hidden",
		LichnerowiczFormula:           "D_M²=∇*∇+R/4, equivalently E_R=-R/4 in the P=-(∇²+E) heat-kernel convention",
		UniversalA2RPart:              universalA2RPart,
		DiracEndomorphismRPart:        diracEndomorphismR,
		CombinedA2RPart:               combinedDiracA2R,
		CombinedA2RPartMagnitude:      math.Abs(combinedDiracA2R),
		FiniteDiracPartAddsCurvature:  false,
		CurvatureEndomorphismAudited:  true,
		SignConventionClosed:          false,
		PhysicalMetricDynamicsDerived: false,
		Verdict:                       strings.Join([]string{StatusD2EndomorphismSieveExecuted, StatusLichnerowiczCurvatureTermAudited}, ";"),
		Reason:                        "the continuum Clifford/Lichnerowicz identity supplies the scalar-curvature endomorphism contribution; the finite Dirac operator supplies internal masses/one-forms but no new scalar-curvature sign convention or metric dynamics theorem",
	}
}

func buildA2(g377 productspectralactioncoefficients.Analysis, e D2EndomorphismSieve) A2TraceEvaluation {
	weight := finiteTraceDimension * e.CombinedA2RPartMagnitude
	raw := weight / (16 * math.Pi * math.Pi)
	gate377Raw := g377.Calculation.A2RawEinsteinCoefficientPerMP2.Numeric / g377.Calculation.Finite.F2LambdaOverMP2
	return A2TraceEvaluation{
		Executed:                          true,
		FiniteHilbertTraceDimension:       finiteTraceDimension,
		A2WeightMagnitudeBefore4Pi:        weight,
		RawDensityCoefficientPerF2Lambda2: raw,
		Gate377RawDensityCoefficient:      gate377Raw,
		Gate377RawCoefficientMatched:      nearly(raw, gate377Raw, 1e-14),
		DimensionlessTraceWeightNative:    true,
		IncludesCutoffMoment:              false,
		PhysicalCoefficientNative:         false,
		Formula:                           "|a2_R| = (4π)^(-2) ∫√g Tr_F(1)·R/12, so |C_R| = f₂Λ²·Tr_F(1)/(192π²)",
		Verdict:                           strings.Join([]string{StatusA2TraceWeightComputed, StatusFiniteTraceWeightNative, StatusRawGate377CoefficientMatched}, ";"),
		Reason:                            "the finite trace dimension fixes the dimensionless curvature weight Tr_F(1)/12=8; the physical Einstein-Hilbert coefficient still requires the external dimensionful product f₂Λ²",
	}
}

func buildConvention(g377 productspectralactioncoefficients.Analysis, a2 A2TraceEvaluation) TraceConventionAudit {
	skeleton := 0.5 * g377.Calculation.Finite.TrOne
	return TraceConventionAudit{
		Executed:                        true,
		RawConventionDeclared:           g377.Calculation.Convention.IncludesRaw16Pi2 && g377.Calculation.Convention.A2DiracRDensity != "",
		SkeletonConventionDeclared:      g377.Calculation.A2SkeletonEinsteinCoefficientPerMP2.Formula != "",
		RawCoefficientPerF2Lambda2:      a2.RawDensityCoefficientPerF2Lambda2,
		SkeletonCoefficientPerF2Lambda2: skeleton,
		RawSkeletonNumericallyDifferent: math.Abs(a2.RawDensityCoefficientPerF2Lambda2-skeleton) > 1,
		UniqueTraceConventionSelected:   false,
		CanPromoteEitherToNewtonNative:  false,
		Verdict:                         StatusFailedTraceConventionNotUniquelySelected,
		Reason:                          "Gate377 deliberately reported both a raw heat-kernel coefficient and a prompt-skeleton coefficient. Gate510 matches the raw coefficient but does not canonically select a universal trace-renormalization convention that would turn either into Newton normalization.",
	}
}

func buildCutoff(a2 A2TraceEvaluation) CutoffProvenance {
	return CutoffProvenance{
		Executed:                           true,
		EinsteinHilbertSymbolicCoefficient: fmt.Sprintf("C_R = ± f₂Λ²·Tr_F(1)/(192π²) = ± f₂Λ²·%.12g", a2.RawDensityCoefficientPerF2Lambda2),
		RequiresF2LambdaSquaredProduct:     true,
		F2LambdaProductNativeOnlyAsSymbol:  true,
		F2MomentSeparatedFromLambda:        false,
		CutoffLambdaSelected:               false,
		NewtonConstantDerived:              false,
		PlanckScaleImported:                false,
		CosmologicalConstantDerived:        false,
		CosmologicalF4Excluded:             true,
		GravityNormalizationBridgeOnly:     true,
		Verdict: strings.Join([]string{
			StatusF2LambdaObligationIsolated,
			StatusGravityNormalizationQuarantined,
			StatusFailedNewtonConstantStillNotDerived,
			StatusFailedCutoffLambdaStillNotSelected,
			StatusFailedF2MomentStillNotSeparated,
			StatusFailedEHNormalizationStillOpen,
			StatusFailedCosmologicalF4Excluded,
		}, ";"),
		Reason: "the spectral action fixes the symbolic slot f₂Λ² times a native dimensionless trace coefficient; it does not choose Λ, derive f₂ as an independent moment, import M_P/G, or solve the f₄ cosmological/vacuum-subtraction channel",
	}
}

func buildFirewall() Firewall {
	return Firewall{
		Executed:                           true,
		NewtonConstantImported:             false,
		NewtonConstantDerived:              false,
		PlanckMassImported:                 false,
		CutoffLambdaSelected:               false,
		F2MomentSeparatedFromLambda:        false,
		EinsteinHilbertNormalizationClosed: false,
		CosmologicalConstantImported:       false,
		CosmologicalConstantDerived:        false,
		ElectroweakScaleImported:           false,
		FlavorDataImported:                 false,
		NativeGravityNormalizationWritten:  false,
		Verdict:                            strings.Join([]string{StatusNoEmpiricalGravityDataImported, StatusFirewallPreserved, StatusFirewallNativeGravityNormalizationStop}, ";"),
		Reason:                             "Gate510 computes only dimensionless trace weights and symbolic coefficient slots. It imports no G, M_P, Λ value, cosmological constant, electroweak scale, masses, Yukawas, CKM, or PMNS data, and writes no native gravity normalization.",
	}
}

func buildRegistry(a Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries: []string{
			"The Lichnerowicz/Dirac-square scalar-curvature endomorphism has been audited: in the declared raw convention E_R=-R/4 and E+R/6=-R/12.",
			"The finite Hilbert trace supplies the dimensionless curvature weight Tr_F(1)/12=96/12=8 in the raw a2 channel.",
		},
		BridgeEntries: []string{
			"The raw Einstein-Hilbert spectral socket is C_R=±f₂Λ²·Tr_F(1)/(192π²); the sign and physical normalization remain convention/bridge data.",
			"Gate377's raw a2 coefficient is matched, while its skeleton convention remains an unselected trace-renormalization candidate.",
		},
		EnvironmentalEntries: []string{
			"Newton's constant, Planck/cutoff identification, the independent f₂ moment, cosmological f₄/vacuum subtraction, and physical metric normalization remain quarantined.",
		},
		FailedRoutes: []string{
			StatusFailedNewtonConstantStillNotDerived,
			StatusFailedCutoffLambdaStillNotSelected,
			StatusFailedF2MomentStillNotSeparated,
			StatusFailedEHNormalizationStillOpen,
			StatusFailedTraceConventionNotUniquelySelected,
			StatusFailedCosmologicalF4Excluded,
			StatusFailedPhysicalMetricDynamicsNotDerived,
		},
		OpenTheorems: []string{
			"Prove or reject a native cutoff-scale selector for Λ without importing M_P or G.",
			"Prove or reject a native f₂ moment theorem independent of the cutoff scale.",
			"Audit the gravitational a4 curvature-squared/topological terms separately from Einstein-Hilbert normalization.",
		},
	}
}

func buildNext() NextStep {
	return NextStep{Gate: 511, Title: "Gravitational a4 Curvature-Squared and Topological Counterterm Audit", Reason: "Gate510 closes the a2 coefficient provenance up to the f₂Λ² airlock; the next native gravity lane should audit the scale-independent a4 curvature²/Gauss-Bonnet/Weyl sockets before any cosmological or Newton normalization claim.", PrimaryTask: "classify the spectral a4 gravitational curvature-squared terms, identify topological versus dynamical curvature invariants, and preserve the firewall around f4 vacuum energy and physical gravitational couplings"}
}

func validate(a Analysis) error {
	checks := []struct {
		ok  bool
		msg string
	}{
		{a.Inheritance.Executed && a.Inheritance.Gate509GravitySocketInherited && a.Inheritance.Gate509NormalizationBlocked && a.Inheritance.Gate509NoEmpiricalDataImported && a.Inheritance.ProductTripleValid && a.Inheritance.HeatKernelConventionDeclared && a.Inheritance.Gate377RawA2ChannelPresent && a.Inheritance.Gate377SkeletonChannelPresent && !a.Inheritance.Gate377AllCoefficientsClosed && !a.Inheritance.Gate377HardTOEClosure, "Gate510 inheritance invalid"},
		{a.Endomorphism.Executed && a.Endomorphism.CurvatureEndomorphismAudited && nearly(a.Endomorphism.CombinedA2RPart, -1.0/12.0, 1e-15) && nearly(a.Endomorphism.CombinedA2RPartMagnitude, 1.0/12.0, 1e-15) && !a.Endomorphism.FiniteDiracPartAddsCurvature && !a.Endomorphism.SignConventionClosed && !a.Endomorphism.PhysicalMetricDynamicsDerived, "Gate510 D2 endomorphism sieve invalid"},
		{a.A2.Executed && nearly(a.A2.FiniteHilbertTraceDimension, 96, 1e-15) && nearly(a.A2.A2WeightMagnitudeBefore4Pi, 8, 1e-15) && a.A2.RawDensityCoefficientPerF2Lambda2 > 0 && a.A2.Gate377RawCoefficientMatched && a.A2.DimensionlessTraceWeightNative && !a.A2.IncludesCutoffMoment && !a.A2.PhysicalCoefficientNative, "Gate510 a2 trace evaluation invalid"},
		{a.Convention.Executed && a.Convention.RawConventionDeclared && a.Convention.SkeletonConventionDeclared && a.Convention.RawSkeletonNumericallyDifferent && !a.Convention.UniqueTraceConventionSelected && !a.Convention.CanPromoteEitherToNewtonNative, "Gate510 convention audit invalid"},
		{a.Cutoff.Executed && a.Cutoff.RequiresF2LambdaSquaredProduct && a.Cutoff.F2LambdaProductNativeOnlyAsSymbol && !a.Cutoff.F2MomentSeparatedFromLambda && !a.Cutoff.CutoffLambdaSelected && !a.Cutoff.NewtonConstantDerived && !a.Cutoff.PlanckScaleImported && !a.Cutoff.CosmologicalConstantDerived && a.Cutoff.CosmologicalF4Excluded && a.Cutoff.GravityNormalizationBridgeOnly, "Gate510 cutoff provenance firewall invalid"},
		{a.Firewall.Executed && !a.Firewall.NewtonConstantImported && !a.Firewall.NewtonConstantDerived && !a.Firewall.PlanckMassImported && !a.Firewall.CutoffLambdaSelected && !a.Firewall.F2MomentSeparatedFromLambda && !a.Firewall.EinsteinHilbertNormalizationClosed && !a.Firewall.CosmologicalConstantImported && !a.Firewall.CosmologicalConstantDerived && !a.Firewall.ElectroweakScaleImported && !a.Firewall.FlavorDataImported && !a.Firewall.NativeGravityNormalizationWritten, "Gate510 gravity firewall violated"},
		{a.Next.Gate == 511, "Gate511 next step missing"},
	}
	for _, c := range checks {
		if !c.ok {
			return fmt.Errorf(c.msg)
		}
	}
	return nil
}

func truth(a Analysis) string {
	if a.A2.Gate377RawCoefficientMatched && a.Cutoff.GravityNormalizationBridgeOnly && !a.Firewall.NativeGravityNormalizationWritten {
		return "Gate 510 proves exactly where ASHA has native gravitational leverage and where it stops. The D²/Lichnerowicz heat-kernel channel fixes the dimensionless scalar-curvature trace weight: E_R=-R/4, E+R/6=-R/12, and Tr_F(1)/12=8. This matches the raw Gate377 Einstein-Hilbert a2 socket. But physical gravity requires the dimensionful product f₂Λ² plus a selected trace/sign convention and matching to Newton normalization. None of those are native-closed here, so G, M_P, Λ, the cosmological constant, and physical metric normalization remain quarantined."
	}
	return "Gate 510 failed before separating dimensionless curvature trace provenance from Newton normalization."
}

func statuses() []string {
	return []string{
		StatusGate509Inherited,
		StatusProductHeatKernelConventionInherited,
		StatusD2EndomorphismSieveExecuted,
		StatusLichnerowiczCurvatureTermAudited,
		StatusA2TraceWeightComputed,
		StatusFiniteTraceWeightNative,
		StatusRawGate377CoefficientMatched,
		StatusF2LambdaObligationIsolated,
		StatusGravityNormalizationQuarantined,
		StatusNoEmpiricalGravityDataImported,
		StatusNextA4CurvatureGateDefined,
		StatusFailedNewtonConstantStillNotDerived,
		StatusFailedCutoffLambdaStillNotSelected,
		StatusFailedF2MomentStillNotSeparated,
		StatusFailedEHNormalizationStillOpen,
		StatusFailedTraceConventionNotUniquelySelected,
		StatusFailedCosmologicalF4Excluded,
		StatusFailedPhysicalMetricDynamicsNotDerived,
		StatusFirewallPreserved,
		StatusFirewallNativeGravityNormalizationStop,
	}
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%t gate509_socket=%t gate509_norm_blocked=%t no_empirical=%t product=%t heat_kernel=%t raw_a2=%t skeleton=%t all_coeffs=%t TOE=%t verdict=%s reason=%s", x.Executed, x.Gate509GravitySocketInherited, x.Gate509NormalizationBlocked, x.Gate509NoEmpiricalDataImported, x.ProductTripleValid, x.HeatKernelConventionDeclared, x.Gate377RawA2ChannelPresent, x.Gate377SkeletonChannelPresent, x.Gate377AllCoefficientsClosed, x.Gate377HardTOEClosure, x.Verdict, x.Reason)
}

func FormatEndomorphism(x D2EndomorphismSieve) string {
	return fmt.Sprintf("executed=%t convention=%q lichnerowicz=%q universal=%.12g E_R=%.12g combined=%.12g |combined|=%.12g finite_DF_curvature=%t E_audited=%t sign_closed=%t metric_dynamics=%t verdict=%s reason=%s", x.Executed, x.OperatorConvention, x.LichnerowiczFormula, x.UniversalA2RPart, x.DiracEndomorphismRPart, x.CombinedA2RPart, x.CombinedA2RPartMagnitude, x.FiniteDiracPartAddsCurvature, x.CurvatureEndomorphismAudited, x.SignConventionClosed, x.PhysicalMetricDynamicsDerived, x.Verdict, x.Reason)
}

func FormatA2(x A2TraceEvaluation) string {
	return fmt.Sprintf("executed=%t TrF=%.12g weight_before_4pi=%.12g raw_density_per_f2Lambda2=%.12g gate377_raw_density=%.12g matched=%t native_weight=%t includes_cutoff=%t physical=%t formula=%s verdict=%s reason=%s", x.Executed, x.FiniteHilbertTraceDimension, x.A2WeightMagnitudeBefore4Pi, x.RawDensityCoefficientPerF2Lambda2, x.Gate377RawDensityCoefficient, x.Gate377RawCoefficientMatched, x.DimensionlessTraceWeightNative, x.IncludesCutoffMoment, x.PhysicalCoefficientNative, x.Formula, x.Verdict, x.Reason)
}

func FormatConvention(x TraceConventionAudit) string {
	return fmt.Sprintf("executed=%t raw_declared=%t skeleton_declared=%t raw_per_f2Lambda2=%.12g skeleton_per_f2Lambda2=%.12g different=%t unique=%t promotable=%t verdict=%s reason=%s", x.Executed, x.RawConventionDeclared, x.SkeletonConventionDeclared, x.RawCoefficientPerF2Lambda2, x.SkeletonCoefficientPerF2Lambda2, x.RawSkeletonNumericallyDifferent, x.UniqueTraceConventionSelected, x.CanPromoteEitherToNewtonNative, x.Verdict, x.Reason)
}

func FormatCutoff(x CutoffProvenance) string {
	return fmt.Sprintf("executed=%t CR=%q requires_f2Lambda2=%t symbolic_only=%t f2_separated=%t Lambda_selected=%t G=%t Planck_imported=%t cosmo_const=%t f4_excluded=%t bridge_only=%t verdict=%s reason=%s", x.Executed, x.EinsteinHilbertSymbolicCoefficient, x.RequiresF2LambdaSquaredProduct, x.F2LambdaProductNativeOnlyAsSymbol, x.F2MomentSeparatedFromLambda, x.CutoffLambdaSelected, x.NewtonConstantDerived, x.PlanckScaleImported, x.CosmologicalConstantDerived, x.CosmologicalF4Excluded, x.GravityNormalizationBridgeOnly, x.Verdict, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("executed=%t G_imported=%t G_derived=%t Planck_imported=%t Lambda_selected=%t f2_separated=%t EH_norm_closed=%t cosmo_imported=%t cosmo_derived=%t EW_imported=%t flavor_imported=%t native_write=%t verdict=%s reason=%s", x.Executed, x.NewtonConstantImported, x.NewtonConstantDerived, x.PlanckMassImported, x.CutoffLambdaSelected, x.F2MomentSeparatedFromLambda, x.EinsteinHilbertNormalizationClosed, x.CosmologicalConstantImported, x.CosmologicalConstantDerived, x.ElectroweakScaleImported, x.FlavorDataImported, x.NativeGravityNormalizationWritten, x.Verdict, x.Reason)
}

func Markdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 510 Registry Audit — Curvature Coefficient Provenance and Heat-Kernel Trace Convention Audit\n\n")
	b.WriteString("## Verdict\n\n")
	for _, s := range statuses() {
		b.WriteString("- `" + s + "`\n")
	}
	b.WriteString("\n## Inherited boundary\n\n")
	b.WriteString("Gate509 accepted the anomaly theorem and the structural Einstein-Hilbert socket, but blocked Newton normalization, cutoff selection, f2 separation, cosmological normalization, electroweak scales, and flavor data. Gate510 audits only the symbolic `D²`/`a2` provenance of the scalar-curvature coefficient.\n\n")
	b.WriteString("```text\n" + FormatInheritance(a.Inheritance) + "\n```\n\n")
	b.WriteString("## D² endomorphism sieve\n\n")
	b.WriteString("```text\n" + FormatEndomorphism(a.Endomorphism) + "\n```\n\n")
	b.WriteString("The raw convention gives `E_R=-R/4`, so the universal heat-kernel term `E+R/6` contributes `-R/12`. Gate510 records the magnitude and keeps the sign convention explicit.\n\n")
	b.WriteString("## a2 trace evaluation\n\n")
	b.WriteString("```text\n" + FormatA2(a.A2) + "\n```\n\n")
	b.WriteString("The native finite part fixes the dimensionless trace weight `Tr_F(1)/12 = 96/12 = 8`. This is the exact provenance of the raw scalar-curvature socket. It is not yet Newton's constant because the coefficient still carries `f₂Λ²`.\n\n")
	b.WriteString("## Trace convention audit\n\n")
	b.WriteString("```text\n" + FormatConvention(a.Convention) + "\n```\n\n")
	b.WriteString("## Cutoff provenance\n\n")
	b.WriteString("```text\n" + FormatCutoff(a.Cutoff) + "\n```\n\n")
	b.WriteString("## Firewall result\n\n")
	b.WriteString("```text\n" + FormatFirewall(a.Firewall) + "\n```\n\n")
	b.WriteString("No empirical value of `G`, `M_P`, `Λ`, the cosmological constant, electroweak scales, or flavor data was imported. No physical gravitational normalization was written natively.\n\n")
	b.WriteString("## Registry update\n\n")
	writeList(&b, "Native", a.Registry.NativeEntries)
	writeList(&b, "Bridge", a.Registry.BridgeEntries)
	writeList(&b, "Environmental", a.Registry.EnvironmentalEntries)
	writeList(&b, "Failed routes", a.Registry.FailedRoutes)
	writeList(&b, "Open theorems", a.Registry.OpenTheorems)
	b.WriteString("## Next step\n\nGate511 should be:\n\n```text\nGate 511 — " + a.Next.Title + "\n```\n\nPrimary task:\n\n```text\n" + a.Next.PrimaryTask + "\n```\n\n")
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
