// Package heatkerneldynamicspreflight implements Gate 299:
// Seeley-de Witt Heat-Kernel Formalization / Spectral Action Dynamics Preflight.
//
// Gate 298 recovered the Standard Model kinematic field-content skeleton from
// inner fluctuations of the completed structural finite spectral triple. Gate
// 299 does not attempt to predict masses or couplings. It formalizes the
// heat-kernel projection needed to turn the fluctuated Dirac field content into
// Lagrangian coefficients, and it records the exact remaining normalization
// obligations before dynamics can be claimed.
package heatkerneldynamicspreflight

import (
	"fmt"
	"strings"
	"sync"
)

const (
	AuditID = "GATE299-SEELEY-DE-WITT-HEAT-KERNEL-DYNAMICS-PREFLIGHT"

	StatusGate298Inherited           = "CONDITIONAL_SUPPORT_GATE298_FIELD_CONTENT_INHERITED"
	StatusHeatKernelFormalized       = "CONDITIONAL_SUPPORT_HEAT_KERNEL_EXPANSION_FORMALIZED"
	StatusCoefficientMapBuilt        = "CONDITIONAL_SUPPORT_COEFFICIENT_MAPPING_LEDGER_BUILT"
	StatusNormalizationSieveComplete = "CONDITIONAL_SUPPORT_NORMALIZATION_REQUIREMENT_SIEVE_COMPLETED"
	StatusBGapPreflightComplete      = "CONDITIONAL_SUPPORT_BGAP_MAJORANA_HEAT_KERNEL_PREFLIGHT_COMPLETED"
	StatusFirewallsPreserved         = "CONDITIONAL_SUPPORT_SPECTRAL_DYNAMICS_FIREWALLS_PRESERVED"

	StatusFailedCutoffMomentsNotDerived        = "FAILED_ROUTE_CUTOFF_MOMENTS_NOT_PHYSICALLY_DERIVED_FOR_LAGRANGIAN"
	StatusFailedPhysicalJStillFormal           = "FAILED_ROUTE_PHYSICAL_J_ANTILINEAR_SEMANTICS_STILL_FORMAL"
	StatusFailedScalarGaugeNormMissing         = "FAILED_ROUTE_SCALAR_GAUGE_KINETIC_NORMALIZATION_MISSING"
	StatusFailedHeatKernelSubtractionMissing   = "FAILED_ROUTE_HEAT_KERNEL_SUBTRACTION_SCHEME_MISSING"
	StatusFailedYukawaMatricesFree             = "FAILED_ROUTE_NUMERICAL_YUKAWA_MATRICES_REMAIN_FREE"
	StatusFailedHiggsPotentialNotDerived       = "FAILED_ROUTE_HIGGS_POTENTIAL_COEFFICIENTS_NOT_DERIVED"
	StatusFailedHiggsMassRatioNotDerived       = "FAILED_ROUTE_HIGGS_MASS_RATIO_NOT_DERIVED"
	StatusFailedBGapInstantonNotDerived        = "FAILED_ROUTE_BGAP_INSTANTON_ACTION_NOT_DERIVED"
	StatusFailedMajoranaActivationNotDerived   = "FAILED_ROUTE_BGAP_MAJORANA_EDGE_NOT_DERIVED"
	StatusFailedDynamicalPredictionsFirewalled = "FAILED_ROUTE_DYNAMICAL_MASS_AND_HIERARCHY_PREDICTIONS_STILL_FIREWALLED"
)

type InheritedFieldContent struct {
	Gate298FieldContentDerived bool
	GaugeGroup                 string
	GaugeDirections            int
	ScalarContent              string
	HiggsRealDimension         int
	Sin2ThirdPath              string
	NumericalDynamicsDerived   bool
	Verdict                    string
}

type HeatKernelExpansion struct {
	SpectralAction                   string
	AsymptoticExpansion              string
	CoefficientConvention            string
	A0Role                           string
	A2Role                           string
	A4Role                           string
	RequiresAlmostCommutativeProduct bool
	Formalized                       bool
	Verdict                          string
}

type CoefficientMapping struct {
	FieldName        string
	Source           string
	TargetTerm       string
	Coefficient      string
	DerivedAtGate299 bool
	Requirement      string
}

type CoefficientMapAudit struct {
	Mappings             []CoefficientMapping
	HiggsQuadraticMapped bool
	GaugeKineticMapped   bool
	HiggsQuarticMapped   bool
	OnlyFormalProjection bool
	Verdict              string
}

type NormalizationRequirement struct {
	Name             string
	WhyRequired      string
	Status           string
	BlocksPrediction bool
}

type NormalizationSieve struct {
	Requirements []NormalizationRequirement
	AllCataloged bool
	AnyMissing   bool
	Verdict      string
}

type BGapHeatKernelPreflight struct {
	InsertionHypothesis      string
	A2Scaling                string
	A4Scaling                string
	InverseCouplingGenerated bool
	MajoranaEdgeDerived      bool
	InstantonActionDerived   bool
	Conclusion               string
	Verdict                  string
}

type Firewalls struct {
	DoesNotClaimPhysicalCutoffMoments bool
	DoesNotClaimHiggsPotential        bool
	DoesNotClaimHiggsMassRatio        bool
	DoesNotClaimBGapInstanton         bool
	DoesNotActivateMajorana           bool
	DoesNotInventYukawaMatrices       bool
	FiniteCorePolluted                bool
	Verdict                           string
}

type Summary struct {
	HeatKernelFormalized     bool
	CoefficientLedgerBuilt   bool
	NormalizationObligations bool
	BGapPreflightCompleted   bool
	ActualDynamicsDerived    bool
	FirewallPreserved        bool
	Status                   string
	DirectAnswer             string
	NextGate                 string
}

type Analysis struct {
	Input          InheritedFieldContent
	HeatKernel     HeatKernelExpansion
	CoefficientMap CoefficientMapAudit
	Normalization  NormalizationSieve
	BGap           BGapHeatKernelPreflight
	Firewalls      Firewalls
	Summary        Summary
	Truth          string
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
	input := inheritGate298()
	heat := formalizeHeatKernel()
	coeff := buildCoefficientMap()
	norm := buildNormalizationSieve()
	bgap := auditBGapHeatKernel()
	fw := auditFirewalls(norm, bgap)
	summary := buildSummary(heat, coeff, norm, bgap, fw)
	truth := "Gate 299 formalizes the Seeley-de Witt heat-kernel projection required to turn Gate-298 inner-fluctuation field content into Lagrangian coefficients. It records how a₂ and a₄ would carry Higgs quadratic, gauge kinetic, and Higgs quartic terms, but it does not evaluate physical dynamics. The actual Higgs mass ratio, Higgs potential coefficients, B-gap instanton action, cutoff-moment normalization, scalar/gauge kinetic normalization, and numerical Yukawa data remain firewalled."
	return Analysis{Input: input, HeatKernel: heat, CoefficientMap: coeff, Normalization: norm, BGap: bgap, Firewalls: fw, Summary: summary, Truth: truth}, nil
}

func inheritGate298() InheritedFieldContent {
	return InheritedFieldContent{
		Gate298FieldContentDerived: true,
		GaugeGroup:                 "U(1)_Y×SU(2)_L×SU(3)_C",
		GaugeDirections:            12,
		ScalarContent:              "one complex SU(2)_L Higgs doublet, SU(3)_C singlet",
		HiggsRealDimension:         4,
		Sin2ThirdPath:              "sin²θ_W=3/8 via k_Y=5/3 representation trace normalization",
		NumericalDynamicsDerived:   false,
		Verdict:                    StatusGate298Inherited,
	}
}

func formalizeHeatKernel() HeatKernelExpansion {
	return HeatKernelExpansion{
		SpectralAction:                   "S_B=Tr(f(D_A/Λ))",
		AsymptoticExpansion:              "Tr(f(D_A/Λ)) ~ f_4 Λ^4 a_0(D_A)+f_2 Λ^2 a_2(D_A)+f_0 a_4(D_A)+O(Λ^{-2})",
		CoefficientConvention:            "four-dimensional almost-commutative convention; numerical (4π)^{-2} factors, sign conventions, and cutoff moments must be fixed before physical coefficients are emitted",
		A0Role:                           "cosmological/volume term and finite multiplicity trace",
		A2Role:                           "quadratic scalar/Higgs mass-parameter channel plus curvature-coupling channel in the continuum product geometry",
		A4Role:                           "Yang-Mills kinetic terms, scalar kinetic normalization after rescaling, Higgs quartic channel, and curvature-squared terms",
		RequiresAlmostCommutativeProduct: true,
		Formalized:                       true,
		Verdict:                          StatusHeatKernelFormalized,
	}
}

func buildCoefficientMap() CoefficientMapAudit {
	mappings := []CoefficientMapping{
		{FieldName: "cosmological channel", Source: "a_0(D_A)", TargetTerm: "Λ^4 volume/multiplicity term", Coefficient: "f_4", DerivedAtGate299: false, Requirement: "requires cutoff moment and spacetime volume normalization"},
		{FieldName: "Higgs quadratic channel", Source: "a_2(D_A)", TargetTerm: "-μ_H² |H|² after canonical scalar normalization", Coefficient: "f_2 Λ² times heat-kernel normalization", DerivedAtGate299: false, Requirement: "requires scalar kinetic normalization, sign convention, subtraction scheme, and selected D_F amplitudes"},
		{FieldName: "Yang-Mills gauge kinetic channel", Source: "a_4(D_A)", TargetTerm: "(1/4g_i²) F_i,μν F_i^{μν}", Coefficient: "f_0 times representation trace index", DerivedAtGate299: false, Requirement: "requires gauge field normalization and trace convention; Gate 298 supplies structural indices only"},
		{FieldName: "Higgs quartic channel", Source: "a_4(D_A)", TargetTerm: "λ_H |H|⁴", Coefficient: "f_0 times scalar trace invariant", DerivedAtGate299: false, Requirement: "requires physical Yukawa/Dirac amplitude ledger or sealed proxy plus scalar projection map"},
		{FieldName: "Yukawa interactions", Source: "finite one-forms over D_F", TargetTerm: "ψ_L H ψ_R + h.c.", Coefficient: "Y_f matrices", DerivedAtGate299: false, Requirement: "numerical Yukawa matrices remain empirical/free; only edge graph is structural"},
	}
	return CoefficientMapAudit{Mappings: mappings, HiggsQuadraticMapped: true, GaugeKineticMapped: true, HiggsQuarticMapped: true, OnlyFormalProjection: true, Verdict: strings.Join([]string{StatusCoefficientMapBuilt, StatusFailedHiggsPotentialNotDerived, StatusFailedYukawaMatricesFree}, ";")}
}

func buildNormalizationSieve() NormalizationSieve {
	reqs := []NormalizationRequirement{
		{Name: "physical cutoff moments f_0,f_2,f_4", WhyRequired: "spectral action coefficients multiply a_4,a_2,a_0; contact-spectrum moments are a diagnostic identification, not yet a physical Lagrangian theorem", Status: StatusFailedCutoffMomentsNotDerived, BlocksPrediction: true},
		{Name: "heat-kernel subtraction/renormalization scheme", WhyRequired: "separates finite scalar terms from vacuum/cosmological and regulator-dependent pieces", Status: StatusFailedHeatKernelSubtractionMissing, BlocksPrediction: true},
		{Name: "scalar kinetic normalization", WhyRequired: "converts raw Higgs one-form trace into canonically normalized |D_μH|² and potential coefficients", Status: StatusFailedScalarGaugeNormMissing, BlocksPrediction: true},
		{Name: "gauge kinetic normalization", WhyRequired: "converts representation trace indices into physical 1/g_i² coefficients", Status: StatusFailedScalarGaugeNormMissing, BlocksPrediction: true},
		{Name: "physical anti-linear J semantics", WhyRequired: "J_swap has correct doubled-space KO sign, but full anti-linear particle/antiparticle representation semantics remain formal", Status: StatusFailedPhysicalJStillFormal, BlocksPrediction: true},
		{Name: "numerical Yukawa/Dirac amplitudes", WhyRequired: "a_4 scalar invariants depend on Tr(Y†Y)^2-type data; edge graph alone gives shape, not values", Status: StatusFailedYukawaMatricesFree, BlocksPrediction: true},
		{Name: "B-gap Majorana activation theorem", WhyRequired: "right-handed-neutrino Majorana edge is sealed and cannot be inserted into D_F as native dynamics", Status: StatusFailedMajoranaActivationNotDerived, BlocksPrediction: true},
	}
	return NormalizationSieve{Requirements: reqs, AllCataloged: true, AnyMissing: true, Verdict: strings.Join([]string{StatusNormalizationSieveComplete, StatusFailedCutoffMomentsNotDerived, StatusFailedHeatKernelSubtractionMissing, StatusFailedScalarGaugeNormMissing, StatusFailedPhysicalJStillFormal}, ";")}
}

func auditBGapHeatKernel() BGapHeatKernelPreflight {
	return BGapHeatKernelPreflight{
		InsertionHypothesis:      "If B_gap were sealed as a right-handed neutrino Majorana entry M_R in D_F, it would enter a_2/a_4 through positive even powers of M_R in finite spectral traces.",
		A2Scaling:                "schematically contributes like +|M_R|² to Tr(D_F²)-type channels after projection",
		A4Scaling:                "schematically contributes like +|M_R|⁴ and mixed Yukawa-Majorana invariants to Tr(D_F⁴)-type channels",
		InverseCouplingGenerated: false,
		MajoranaEdgeDerived:      false,
		InstantonActionDerived:   false,
		Conclusion:               "heat-kernel polynomial channels do not by themselves generate S_inst=(4/π)/B_gap; an inverse-coupling/nonperturbative determinant or a native Majorana-action theorem remains required",
		Verdict:                  strings.Join([]string{StatusBGapPreflightComplete, StatusFailedMajoranaActivationNotDerived, StatusFailedBGapInstantonNotDerived}, ";"),
	}
}

func auditFirewalls(n NormalizationSieve, b BGapHeatKernelPreflight) Firewalls {
	return Firewalls{
		DoesNotClaimPhysicalCutoffMoments: true,
		DoesNotClaimHiggsPotential:        true,
		DoesNotClaimHiggsMassRatio:        true,
		DoesNotClaimBGapInstanton:         !b.InstantonActionDerived,
		DoesNotActivateMajorana:           !b.MajoranaEdgeDerived,
		DoesNotInventYukawaMatrices:       true,
		FiniteCorePolluted:                false,
		Verdict:                           strings.Join([]string{StatusFirewallsPreserved, StatusFailedHiggsPotentialNotDerived, StatusFailedHiggsMassRatioNotDerived, StatusFailedBGapInstantonNotDerived, StatusFailedDynamicalPredictionsFirewalled}, ";"),
	}
}

func buildSummary(h HeatKernelExpansion, c CoefficientMapAudit, n NormalizationSieve, b BGapHeatKernelPreflight, f Firewalls) Summary {
	statuses := []string{StatusGate298Inherited, StatusHeatKernelFormalized, StatusCoefficientMapBuilt, StatusNormalizationSieveComplete, StatusBGapPreflightComplete, StatusFirewallsPreserved, StatusFailedCutoffMomentsNotDerived, StatusFailedPhysicalJStillFormal, StatusFailedScalarGaugeNormMissing, StatusFailedHeatKernelSubtractionMissing, StatusFailedYukawaMatricesFree, StatusFailedHiggsPotentialNotDerived, StatusFailedHiggsMassRatioNotDerived, StatusFailedBGapInstantonNotDerived, StatusFailedMajoranaActivationNotDerived, StatusFailedDynamicalPredictionsFirewalled}
	return Summary{
		HeatKernelFormalized:     h.Formalized,
		CoefficientLedgerBuilt:   c.HiggsQuadraticMapped && c.GaugeKineticMapped && c.HiggsQuarticMapped,
		NormalizationObligations: n.AllCataloged && n.AnyMissing,
		BGapPreflightCompleted:   !b.InstantonActionDerived && !b.MajoranaEdgeDerived,
		ActualDynamicsDerived:    false,
		FirewallPreserved:        !f.FiniteCorePolluted && f.DoesNotClaimHiggsMassRatio && f.DoesNotClaimBGapInstanton,
		Status:                   strings.Join(statuses, ";"),
		DirectAnswer:             "Gate 299 formalizes the Seeley-de Witt expansion and maps Gate-298 field content to the Lagrangian coefficient channels, but derives no physical masses, Higgs potential coefficients, or B-gap instanton action.",
		NextGate:                 "A future gate must derive or seal the physical cutoff moments, heat-kernel subtraction scheme, scalar/gauge normalizations, and full anti-linear J semantics before spectral dynamics can be evaluated.",
	}
}

func FormatInput(i InheritedFieldContent) string {
	return fmt.Sprintf("fieldContent=%t group=%s directions=%d scalar=%s realDim=%d sin2=%s dynamics=%t verdict=%s", i.Gate298FieldContentDerived, i.GaugeGroup, i.GaugeDirections, i.ScalarContent, i.HiggsRealDimension, i.Sin2ThirdPath, i.NumericalDynamicsDerived, i.Verdict)
}

func FormatHeatKernel(h HeatKernelExpansion) string {
	return fmt.Sprintf("S=%q expansion=%q convention=%q a0=%q a2=%q a4=%q product=%t formalized=%t verdict=%s", h.SpectralAction, h.AsymptoticExpansion, h.CoefficientConvention, h.A0Role, h.A2Role, h.A4Role, h.RequiresAlmostCommutativeProduct, h.Formalized, h.Verdict)
}

func FormatCoefficientMapping(m CoefficientMapping) string {
	return fmt.Sprintf("%s source=%s target=%s coeff=%s derived=%t req=%s", m.FieldName, m.Source, m.TargetTerm, m.Coefficient, m.DerivedAtGate299, m.Requirement)
}

func FormatCoefficientMap(c CoefficientMapAudit) string {
	parts := []string{}
	for _, m := range c.Mappings {
		parts = append(parts, FormatCoefficientMapping(m))
	}
	return fmt.Sprintf("mappings=[%s] Higgs2=%t gauge=%t Higgs4=%t formalOnly=%t verdict=%s", strings.Join(parts, " | "), c.HiggsQuadraticMapped, c.GaugeKineticMapped, c.HiggsQuarticMapped, c.OnlyFormalProjection, c.Verdict)
}

func FormatRequirement(r NormalizationRequirement) string {
	return fmt.Sprintf("%s required=%s status=%s blocks=%t", r.Name, r.WhyRequired, r.Status, r.BlocksPrediction)
}

func FormatNormalization(n NormalizationSieve) string {
	parts := []string{}
	for _, r := range n.Requirements {
		parts = append(parts, FormatRequirement(r))
	}
	return fmt.Sprintf("requirements=[%s] cataloged=%t missing=%t verdict=%s", strings.Join(parts, " | "), n.AllCataloged, n.AnyMissing, n.Verdict)
}

func FormatBGap(b BGapHeatKernelPreflight) string {
	return fmt.Sprintf("hypothesis=%q a2=%q a4=%q inverse=%t majorana=%t instanton=%t conclusion=%q verdict=%s", b.InsertionHypothesis, b.A2Scaling, b.A4Scaling, b.InverseCouplingGenerated, b.MajoranaEdgeDerived, b.InstantonActionDerived, b.Conclusion, b.Verdict)
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("noCutoff=%t noPotential=%t noHiggsRatio=%t noBGap=%t noMajorana=%t noYukawa=%t polluted=%t verdict=%s", f.DoesNotClaimPhysicalCutoffMoments, f.DoesNotClaimHiggsPotential, f.DoesNotClaimHiggsMassRatio, f.DoesNotClaimBGapInstanton, f.DoesNotActivateMajorana, f.DoesNotInventYukawaMatrices, f.FiniteCorePolluted, f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("heat=%t coeff=%t obligations=%t bgap=%t dynamics=%t firewall=%t status=%s answer=%q next=%q", s.HeatKernelFormalized, s.CoefficientLedgerBuilt, s.NormalizationObligations, s.BGapPreflightCompleted, s.ActualDynamicsDerived, s.FirewallPreserved, s.Status, s.DirectAnswer, s.NextGate)
}
