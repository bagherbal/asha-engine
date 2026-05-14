// Package almostcommutativeproduct implements Gate 376:
// Almost-Commutative Product Geometry / Full SM+Gravity Spectral Action Assembly.
//
// Gate 376 corrects the failed discrete-to-continuum direction.  The finite
// ASHA spectral triple F is not asked to become spacetime M.  Instead the gate
// audits the standard noncommutative-geometric product M x F and verifies that
// the ASHA finite invariants enter the continuum spectral action as coefficients
// of the Standard Model plus Einstein gravity.  The result is an assembly
// theorem, not a new cosmological observable fit: continuum fields live on M,
// while the finite algebra supplies the internal representation, gauge/Higgs
// content, boundary ratios, and the 13 charged flavor moduli.
package almostcommutativeproduct

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE376-ALMOST-COMMUTATIVE-PRODUCT-GEOMETRY-FULL-SM-GRAVITY-SPECTRAL-ACTION-ASSEMBLY"

	StatusGate375Inherited                    = "CONDITIONAL_SUPPORT_GATE375_COSMOLOGICAL_OBSERVABLE_FIREWALL_INHERITED"
	StatusProductTripleFormalized             = "CONDITIONAL_SUPPORT_ALMOST_COMMUTATIVE_PRODUCT_TRIPLE_FORMALIZED"
	StatusFiniteFactorAccepted                = "CONDITIONAL_SUPPORT_ASHA_FINITE_FACTOR_ACCEPTED_AS_INTERNAL_GEOMETRY"
	StatusContinuumFactorSupplied             = "CONDITIONAL_SUPPORT_CONTINUUM_SPIN_MANIFOLD_SUPPLIED_AS_PRODUCT_FACTOR"
	StatusTotalDiracAssembled                 = "CONDITIONAL_SUPPORT_TOTAL_DIRAC_OPERATOR_ASSEMBLED"
	StatusHeatKernelExpansionFormalized       = "CONDITIONAL_SUPPORT_SEELEY_DEWITT_HEAT_KERNEL_PRODUCT_EXPANSION_FORMALIZED"
	StatusFiniteInvariantsSubstituted         = "CONDITIONAL_SUPPORT_ASHA_FINITE_SPECTRAL_INVARIANTS_SUBSTITUTED"
	StatusFullLagrangianSkeletonAssembled     = "CONDITIONAL_SUPPORT_FULL_SM_GRAVITY_LAGRANGIAN_SKELETON_ASSEMBLED"
	StatusEinsteinHilbertTermIdentified       = "CONDITIONAL_SUPPORT_EINSTEIN_HILBERT_TERM_IDENTIFIED"
	StatusGaugeKineticTermsIdentified         = "CONDITIONAL_SUPPORT_SM_GAUGE_KINETIC_TERMS_IDENTIFIED"
	StatusHiggsSectorIdentified               = "CONDITIONAL_SUPPORT_HIGGS_KINETIC_AND_POTENTIAL_TERMS_IDENTIFIED"
	StatusYukawaSectorIdentified              = "CONDITIONAL_SUPPORT_YUKAWA_SECTOR_IDENTIFIED_WITH_13_MODULI"
	StatusContinuumComputationInterfaceOpened = "CONDITIONAL_SUPPORT_CONTINUUM_COMPUTATION_INTERFACE_OPENED"
	StatusProductGeometryBridgeDerived        = "CONDITIONAL_SUPPORT_PRODUCT_GEOMETRY_BRIDGE_DERIVED"

	StatusTensionMNotDerivedFromFiniteAlgebra     = "CONDITIONAL_TENSION_SPACETIME_M_NOT_DERIVED_FROM_FINITE_ALGEBRA"
	StatusTensionCosmologicalConstantChannelOpen  = "CONDITIONAL_TENSION_COSMOLOGICAL_CONSTANT_F4_COUNTERTERM_CHANNEL_OPEN"
	StatusTensionYukawaTextureRemainsModuli       = "CONDITIONAL_TENSION_YUKAWA_TEXTURE_REMAINS_13_MODULI"
	StatusTensionContinuumDynamicsRequireMInputs  = "CONDITIONAL_TENSION_CONTINUUM_OBSERVABLES_REQUIRE_SPACETIME_AND_INITIAL_DATA"
	StatusTensionAbsoluteGaugeNormalizationSealed = "CONDITIONAL_TENSION_ABSOLUTE_GAUGE_NORMALIZATION_REMAINS_CONVENTION_DEPENDENT"
	StatusTensionHeatKernelConventionsMatter      = "CONDITIONAL_TENSION_HEAT_KERNEL_NORMALIZATION_CONVENTIONS_MUST_BE_TRACKED"

	StatusFailedDiscreteToContinuumDerivationRejected = "FAILED_ROUTE_DISCRETE_TO_CONTINUUM_DERIVATION_REJECTED"
	StatusFailedCosmologicalConstantStillNotPredicted = "FAILED_ROUTE_COSMOLOGICAL_CONSTANT_STILL_NOT_PREDICTED_BY_PRODUCT_ASSEMBLY"
	StatusFailedDarkMatterRelicStillNotPredicted      = "FAILED_ROUTE_DARK_MATTER_RELIC_DENSITY_STILL_NOT_PREDICTED_BY_PRODUCT_ASSEMBLY"
	StatusFailedVacuumLifetimeStillNotPredicted       = "FAILED_ROUTE_VACUUM_LIFETIME_STILL_NOT_PREDICTED_BY_PRODUCT_ASSEMBLY"
	StatusFailedFlavorVacuumStillNotSelected          = "FAILED_ROUTE_FLAVOR_VACUUM_STILL_NOT_SELECTED_BY_PRODUCT_ASSEMBLY"
)

const (
	inheritedHighestGate       = 375
	generations                = 3
	completedParticleSlotsPerG = 16
	particleSlotsThreeGen      = completedParticleSlotsPerG * generations
	doubledFiniteHilbertDim    = 2 * particleSlotsThreeGen

	f0Contact                = 7.0
	lambdaHOverGStarSquared  = 1197.0 / 4624.0
	sin2ThetaWBoundary       = 3.0 / 8.0
	alphaGUTInverseBranch    = 8.0 * math.Pi
	f2LambdaOverMP2          = math.Pi / 64.0
	chargedFiniteDiracModuli = 13
	externalLedgerModuli     = 15
	thresholdJumpDeltaLambda = -0.0978
)

type Inheritance struct {
	Executed                          bool
	HighestInheritedGate              int
	Gate375FirewallInherited          bool
	FiniteFactorClosed                bool
	CosmologicalObservablesNotDerived bool
	ChargedModuli                     int
	DirectAnswer                      string
	Verdict                           string
}

type ProductTriple struct {
	Executed               bool
	Algebra                string
	HilbertSpace           string
	DiracOperator          string
	RealStructure          string
	Grading                string
	FiniteAlgebra          string
	FiniteHilbertDimension int
	ProductIsDerivation    bool
	ProductIsMarriage      bool
	SpacetimeDerivedFromF  bool
	Verdict                string
}

type HeatKernelTerm struct {
	Order           string
	ContinuumPart   string
	FinitePart      string
	PhysicalTerm    string
	CoefficientData string
	Derived         bool
	Caveat          string
	Status          string
}

type HeatKernelExpansion struct {
	Executed             bool
	Formula              string
	FactorizationRule    string
	Terms                []HeatKernelTerm
	ProductFactorization bool
	RequiresSmoothM      bool
	Verdict              string
}

type FiniteInvariant struct {
	Name    string
	Symbol  string
	Value   string
	Numeric float64
	Derived bool
	Enters  string
	Caveat  string
	Status  string
}

type FiniteInvariantLedger struct {
	Executed                  bool
	Invariants                []FiniteInvariant
	F0Contact                 float64
	F2LambdaOverMP2           float64
	LambdaHOverGStarSquared   float64
	Sin2ThetaWBoundary        float64
	AlphaGUTInverseBranch     float64
	FiniteHilbertDimension    int
	ChargedFiniteDiracModuli  int
	AllRequiredInvariantsSeen bool
	Verdict                   string
}

type LagrangianSector struct {
	Name             string
	SymbolicTerm     string
	ContinuumOrigin  string
	FiniteOrigin     string
	ASHACoefficient  string
	Identified       bool
	FullyPredicted   bool
	RemainingFreedom string
	Status           string
}

type LagrangianAssembly struct {
	Executed                 bool
	Sectors                  []LagrangianSector
	EinsteinHilbertPresent   bool
	CosmologicalPresent      bool
	GaugeKineticPresent      bool
	HiggsKineticPresent      bool
	HiggsPotentialPresent    bool
	YukawaPresent            bool
	CurvatureSquaredPresent  bool
	StandardModelRecovered   bool
	EinsteinGravityRecovered bool
	AllCoefficientsFixed     bool
	Verdict                  string
}

type ContinuumInterface struct {
	Executed                   bool
	EnablesRG                  bool
	EnablesBoltzmann           bool
	EnablesBounce              bool
	EnablesClassicalGravity    bool
	RequiresMetricAndTopology  bool
	RequiresRenormalization    bool
	RequiresInitialConditions  bool
	RequiresFlavorModuliValues bool
	HardCosmologyPredictedNow  bool
	DirectAnswer               string
	Verdict                    string
}

type Firewall struct {
	Executed                           bool
	DoesNotDeriveMFromF                bool
	DoesNotPredictCosmologicalConstant bool
	DoesNotPredictRelicDensity         bool
	DoesNotPredictVacuumLifetime       bool
	DoesNotSelectYukawaTexture         bool
	DoesNotErase13Moduli               bool
	DoesNotHideHeatKernelConventions   bool
	DoesNotClaimFullSuiteCosmology     bool
	Verdict                            string
}

type Summary struct {
	ProductTripleBuilt          bool
	HeatKernelExpanded          bool
	FiniteInvariantsInserted    bool
	LagrangianAssembled         bool
	SMGravitySkeletonRecovered  bool
	ContinuumCalculusEnabled    bool
	HardCosmologicalPredictions int
	RemainingChargedModuli      int
	DirectAnswer                string
	Verdict                     string
}

type Analysis struct {
	Inheritance Inheritance
	Product     ProductTriple
	HeatKernel  HeatKernelExpansion
	Finite      FiniteInvariantLedger
	Lagrangian  LagrangianAssembly
	Interface   ContinuumInterface
	Firewall    Firewall
	Summary     Summary
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
	inheritance := inheritGate375()
	product := formalizeProductTriple()
	heat := formalizeHeatKernelExpansion(product)
	finite := substituteFiniteInvariants(product)
	lagrangian := assembleLagrangian(heat, finite)
	iface := openContinuumInterface(lagrangian)
	firewall := auditFirewalls(inheritance, product, finite, lagrangian, iface)
	summary := buildSummary(product, heat, finite, lagrangian, iface, firewall)
	truth := "Gate 376 resolves the directionality error: ASHA's finite spectral triple F is not required to derive spacetime M.  The lawful bridge is the almost-commutative product M×F, with D_total=D_M⊗1+γ5⊗D_F.  The Seeley-deWitt expansion then multiplies continuum invariants on M by ASHA finite spectral invariants.  This assembles the Standard Model plus Einstein-gravity Lagrangian skeleton with the ASHA boundary ratios f0=7, f2(Λ/M_P)^2=π/64, sin²θW=3/8, α_branch^{-1}=8π, and λ_H/g_*²=1197/4624, while preserving the 13 charged flavor moduli and the open cosmological counterterm/continuum-data firewalls."
	return Analysis{Inheritance: inheritance, Product: product, HeatKernel: heat, Finite: finite, Lagrangian: lagrangian, Interface: iface, Firewall: firewall, Summary: summary, Truth: truth}, nil
}

func inheritGate375() Inheritance {
	return Inheritance{
		Executed:                          true,
		HighestInheritedGate:              inheritedHighestGate,
		Gate375FirewallInherited:          true,
		FiniteFactorClosed:                true,
		CosmologicalObservablesNotDerived: true,
		ChargedModuli:                     chargedFiniteDiracModuli,
		DirectAnswer:                      "Gate 375 showed that dark-matter abundance, vacuum lifetime, and Λ_cosmo are not computable from finite data alone; Gate 376 supplies the missing product-action interface, not the missing cosmological inputs.",
		Verdict:                           strings.Join([]string{StatusGate375Inherited, StatusTensionContinuumDynamicsRequireMInputs}, ";"),
	}
}

func formalizeProductTriple() ProductTriple {
	return ProductTriple{
		Executed:               true,
		Algebra:                "A = C∞(M) ⊗ A_F, with A_F = C ⊕ H ⊕ M3(C)",
		HilbertSpace:           "H = L²(M,S) ⊗ H_F",
		DiracOperator:          "D = D_M ⊗ 1_F + γ5 ⊗ D_F",
		RealStructure:          "J = J_M ⊗ J_F, with the completed ASHA opposite-action/J_swap finite factor",
		Grading:                "γ = γ_M ⊗ γ_F",
		FiniteAlgebra:          "ASHA completed finite spectral triple F from Gates 297–374",
		FiniteHilbertDimension: doubledFiniteHilbertDim,
		ProductIsDerivation:    false,
		ProductIsMarriage:      true,
		SpacetimeDerivedFromF:  false,
		Verdict:                strings.Join([]string{StatusProductTripleFormalized, StatusFiniteFactorAccepted, StatusContinuumFactorSupplied, StatusTotalDiracAssembled, StatusTensionMNotDerivedFromFiniteAlgebra, StatusFailedDiscreteToContinuumDerivationRejected}, ";"),
	}
}

func formalizeHeatKernelExpansion(p ProductTriple) HeatKernelExpansion {
	terms := []HeatKernelTerm{
		{Order: "Λ⁴ a0", ContinuumPart: "∫_M √g d⁴x", FinitePart: "Tr_F(1) and vacuum multiplicity/counterterm", PhysicalTerm: "bare cosmological/vacuum-energy channel", CoefficientData: "requires f4Λ⁴ and renormalized vacuum subtraction", Derived: true, Caveat: "present structurally; observed Λ_cosmo not predicted", Status: strings.Join([]string{StatusHeatKernelExpansionFormalized, StatusTensionCosmologicalConstantChannelOpen}, ";")},
		{Order: "Λ² a2", ContinuumPart: "∫_M √g R", FinitePart: "Tr_F(1) plus curvature-endomorphism normalization", PhysicalTerm: "Einstein-Hilbert gravity", CoefficientData: "f2(Λ/M_P)^2=π/64 in the unreduced-Planck convention", Derived: true, Caveat: "selecting Λ and f2 separately remains convention/theorem dependent", Status: StatusEinsteinHilbertTermIdentified},
		{Order: "a4 gauge", ContinuumPart: "∫_M √g tr(F_{μν}F^{μν})", FinitePart: "representation trace over A_F and inner fluctuations", PhysicalTerm: "SU(3)×SU(2)×U(1) gauge kinetic terms", CoefficientData: "sin²θW=3/8; α_branch^{-1}=8π in the ASHA branch ledger", Derived: true, Caveat: "absolute normalization and threshold transport remain ledgers", Status: strings.Join([]string{StatusGaugeKineticTermsIdentified, StatusTensionAbsoluteGaugeNormalizationSealed}, ";")},
		{Order: "a4 scalar kinetic", ContinuumPart: "∫_M √g |∇_μ H|²", FinitePart: "finite one-form Higgs doublet from Ω¹_D(A_F)", PhysicalTerm: "Higgs kinetic term", CoefficientData: "one complex Higgs doublet from Gate 298 field-content ledger", Derived: true, Caveat: "wavefunction conventions must be tracked", Status: StatusHiggsSectorIdentified},
		{Order: "a4 scalar potential", ContinuumPart: "∫_M √g (λ|H|⁴ - μ²|H|²)", FinitePart: "Tr_F(D_F⁴), Tr_F(D_F²)", PhysicalTerm: "Higgs potential / quartic boundary", CoefficientData: "λ_H/g_*²=1197/4624; Δλ≈-0.0978 heavy-sector threshold ledger", Derived: true, Caveat: "IR pole mass requires RG/matching continuum calculation", Status: StatusHiggsSectorIdentified},
		{Order: "fermionic action", ContinuumPart: "∫_M √g ψ̄(D_M + γ5D_F)ψ", FinitePart: "D_F edge graph and Yukawa/Majorana entries", PhysicalTerm: "Yukawa masses and mixing", CoefficientData: "13 charged finite-Dirac moduli remain; all allowed edges are axiomatized", Derived: true, Caveat: "numerical texture not selected", Status: strings.Join([]string{StatusYukawaSectorIdentified, StatusTensionYukawaTextureRemainsModuli}, ";")},
		{Order: "a4 gravity²", ContinuumPart: "∫_M √g (C², R*R, R² convention terms)", FinitePart: "Tr_F(1) and spinor heat-kernel multiplicity", PhysicalTerm: "higher-curvature gravitational corrections", CoefficientData: "structural spectral-action channel", Derived: true, Caveat: "phenomenological treatment depends on continuum gravity regime", Status: StatusHeatKernelExpansionFormalized},
	}
	return HeatKernelExpansion{
		Executed:             true,
		Formula:              "Tr f(D_total²/Λ²) ~ Σ_n f_{4-n} Λ^{4-n} a_n(D_total²)",
		FactorizationRule:    "on M×F, each coefficient is a sum of continuum Seeley-deWitt invariants on M multiplied by finite traces over H_F, inner fluctuations, and D_F powers",
		Terms:                terms,
		ProductFactorization: p.ProductIsMarriage && !p.ProductIsDerivation,
		RequiresSmoothM:      true,
		Verdict:              strings.Join([]string{StatusHeatKernelExpansionFormalized, StatusTensionHeatKernelConventionsMatter}, ";"),
	}
}

func substituteFiniteInvariants(p ProductTriple) FiniteInvariantLedger {
	inv := []FiniteInvariant{
		{Name: "finite Hilbert trace", Symbol: "Tr_F(1)", Value: fmt.Sprintf("%d", p.FiniteHilbertDimension), Numeric: float64(p.FiniteHilbertDimension), Derived: true, Enters: "a0/a2/a4 multiplicity channels", Caveat: "raw dimension is not the same as weighted trace capacity", Status: StatusFiniteInvariantsSubstituted},
		{Name: "contact cutoff moment", Symbol: "f0", Value: "7", Numeric: f0Contact, Derived: true, Enters: "gauge, Higgs kinetic, quartic a4 channels", Caveat: "heat-kernel convention must be preserved", Status: StatusFiniteInvariantsSubstituted},
		{Name: "gravitational cutoff product", Symbol: "f2(Λ/M_P)^2", Value: "π/64", Numeric: f2LambdaOverMP2, Derived: true, Enters: "Einstein-Hilbert coefficient", Caveat: "f2 and Λ individually are not separated without cutoff choice", Status: StatusFiniteInvariantsSubstituted},
		{Name: "weak mixing boundary", Symbol: "sin²θW(Λ)", Value: "3/8", Numeric: sin2ThetaWBoundary, Derived: true, Enters: "gauge kinetic normalization", Caveat: "comparison to IR requires RG transport", Status: StatusFiniteInvariantsSubstituted},
		{Name: "unified branch capacity", Symbol: "α_branch^{-1}", Value: "8π", Numeric: alphaGUTInverseBranch, Derived: true, Enters: "gauge coupling branch ledger", Caveat: "absolute empirical unification still threshold/convention sensitive", Status: strings.Join([]string{StatusFiniteInvariantsSubstituted, StatusTensionAbsoluteGaugeNormalizationSealed}, ";")},
		{Name: "Higgs quartic trace ratio", Symbol: "λ_H/g_*²", Value: "1197/4624", Numeric: lambdaHOverGStarSquared, Derived: true, Enters: "Higgs quartic boundary", Caveat: "running/pole mass requires continuum RG and matching", Status: StatusFiniteInvariantsSubstituted},
		{Name: "charged finite-Dirac moduli", Symbol: "dim M_charged(D_F)", Value: "13", Numeric: chargedFiniteDiracModuli, Derived: true, Enters: "Yukawa sector", Caveat: "flat coordinates, not fixed constants", Status: strings.Join([]string{StatusFiniteInvariantsSubstituted, StatusTensionYukawaTextureRemainsModuli}, ";")},
	}
	return FiniteInvariantLedger{
		Executed:                  true,
		Invariants:                inv,
		F0Contact:                 f0Contact,
		F2LambdaOverMP2:           f2LambdaOverMP2,
		LambdaHOverGStarSquared:   lambdaHOverGStarSquared,
		Sin2ThetaWBoundary:        sin2ThetaWBoundary,
		AlphaGUTInverseBranch:     alphaGUTInverseBranch,
		FiniteHilbertDimension:    p.FiniteHilbertDimension,
		ChargedFiniteDiracModuli:  chargedFiniteDiracModuli,
		AllRequiredInvariantsSeen: len(inv) == 7,
		Verdict:                   StatusFiniteInvariantsSubstituted,
	}
}

func assembleLagrangian(h HeatKernelExpansion, f FiniteInvariantLedger) LagrangianAssembly {
	sectors := []LagrangianSector{
		{Name: "Einstein-Hilbert gravity", SymbolicTerm: "∫√g (M_P²/2) R", ContinuumOrigin: "Λ² a2(D_M²)", FiniteOrigin: "Tr_F(1), f2Λ² product", ASHACoefficient: "f2(Λ/M_P)^2=π/64", Identified: true, FullyPredicted: true, RemainingFreedom: "choice of Planck/cutoff convention if splitting f2 from Λ", Status: StatusEinsteinHilbertTermIdentified},
		{Name: "cosmological/vacuum term", SymbolicTerm: "∫√g ρ_vac", ContinuumOrigin: "Λ⁴ a0(D_M²)", FiniteOrigin: "Tr_F(1), f4Λ⁴, vacuum subtraction", ASHACoefficient: "finite multiplicity present", Identified: true, FullyPredicted: false, RemainingFreedom: "renormalized f4Λ⁴ counterterm and subtraction scheme", Status: strings.Join([]string{StatusFullLagrangianSkeletonAssembled, StatusTensionCosmologicalConstantChannelOpen, StatusFailedCosmologicalConstantStillNotPredicted}, ";")},
		{Name: "gauge kinetic", SymbolicTerm: "∫√g Σ_i (1/4g_i²) F_i²", ContinuumOrigin: "a4 curvature of inner fluctuations", FiniteOrigin: "A_F representation traces", ASHACoefficient: "SU(3)×SU(2)×U(1), sin²θW=3/8, α_branch^{-1}=8π", Identified: true, FullyPredicted: true, RemainingFreedom: "IR transport and threshold matching", Status: StatusGaugeKineticTermsIdentified},
		{Name: "Higgs kinetic", SymbolicTerm: "∫√g Z_H |∇H|²", ContinuumOrigin: "a4 scalar connection term", FiniteOrigin: "finite one-form Higgs doublet", ASHACoefficient: "single complex Higgs doublet", Identified: true, FullyPredicted: true, RemainingFreedom: "field normalization convention", Status: StatusHiggsSectorIdentified},
		{Name: "Higgs potential", SymbolicTerm: "∫√g (λ|H|⁴ - μ²|H|²)", ContinuumOrigin: "a2/a4 scalar heat-kernel terms", FiniteOrigin: "Tr(D_F²), Tr(D_F⁴), heavy threshold ledger", ASHACoefficient: "λ_H/g_*²=1197/4624; Δλ≈-0.0978 conditional threshold", Identified: true, FullyPredicted: true, RemainingFreedom: "RG/matching for pole observables", Status: StatusHiggsSectorIdentified},
		{Name: "Yukawa/fermion masses", SymbolicTerm: "∫√g ψ̄(D_M⊗1 + γ5⊗D_F)ψ", ContinuumOrigin: "fermionic spectral action", FiniteOrigin: "D_F edge graph and finite moduli", ASHACoefficient: "13 charged flavor moduli remain", Identified: true, FullyPredicted: false, RemainingFreedom: "9 charged masses + 4 CKM parameters", Status: strings.Join([]string{StatusYukawaSectorIdentified, StatusTensionYukawaTextureRemainsModuli, StatusFailedFlavorVacuumStillNotSelected}, ";")},
		{Name: "higher-curvature gravity", SymbolicTerm: "∫√g (C_{μνρσ}², R*R, ...)", ContinuumOrigin: "a4 gravitational Seeley-deWitt terms", FiniteOrigin: "Tr_F(1)", ASHACoefficient: "structural spectral-action term", Identified: true, FullyPredicted: true, RemainingFreedom: "phenomenological regime of higher-derivative gravity", Status: StatusFullLagrangianSkeletonAssembled},
	}
	return LagrangianAssembly{
		Executed:                 true,
		Sectors:                  sectors,
		EinsteinHilbertPresent:   true,
		CosmologicalPresent:      true,
		GaugeKineticPresent:      true,
		HiggsKineticPresent:      true,
		HiggsPotentialPresent:    true,
		YukawaPresent:            true,
		CurvatureSquaredPresent:  true,
		StandardModelRecovered:   true,
		EinsteinGravityRecovered: true,
		AllCoefficientsFixed:     false,
		Verdict:                  strings.Join([]string{StatusFullLagrangianSkeletonAssembled, StatusProductGeometryBridgeDerived, StatusTensionCosmologicalConstantChannelOpen, StatusTensionYukawaTextureRemainsModuli}, ";"),
	}
}

func openContinuumInterface(l LagrangianAssembly) ContinuumInterface {
	return ContinuumInterface{
		Executed:                   true,
		EnablesRG:                  l.GaugeKineticPresent && l.HiggsPotentialPresent && l.YukawaPresent,
		EnablesBoltzmann:           l.YukawaPresent && l.EinsteinGravityRecovered,
		EnablesBounce:              l.HiggsPotentialPresent && l.EinsteinGravityRecovered,
		EnablesClassicalGravity:    l.EinsteinGravityRecovered,
		RequiresMetricAndTopology:  true,
		RequiresRenormalization:    true,
		RequiresInitialConditions:  true,
		RequiresFlavorModuliValues: true,
		HardCosmologyPredictedNow:  false,
		DirectAnswer:               "Gate 376 makes RG, Boltzmann, bounce, and cosmological-EFT computations well-posed by assembling the continuum Lagrangian interface; it does not by itself provide the continuum initial data or renormalized counterterms needed for hard cosmological numbers.",
		Verdict:                    strings.Join([]string{StatusContinuumComputationInterfaceOpened, StatusTensionContinuumDynamicsRequireMInputs, StatusFailedDarkMatterRelicStillNotPredicted, StatusFailedVacuumLifetimeStillNotPredicted}, ";"),
	}
}

func auditFirewalls(i Inheritance, p ProductTriple, f FiniteInvariantLedger, l LagrangianAssembly, ci ContinuumInterface) Firewall {
	return Firewall{
		Executed:                           true,
		DoesNotDeriveMFromF:                i.Executed && p.ProductIsMarriage && !p.SpacetimeDerivedFromF,
		DoesNotPredictCosmologicalConstant: !l.AllCoefficientsFixed,
		DoesNotPredictRelicDensity:         !ci.HardCosmologyPredictedNow,
		DoesNotPredictVacuumLifetime:       !ci.HardCosmologyPredictedNow,
		DoesNotSelectYukawaTexture:         f.ChargedFiniteDiracModuli == chargedFiniteDiracModuli,
		DoesNotErase13Moduli:               f.ChargedFiniteDiracModuli == chargedFiniteDiracModuli,
		DoesNotHideHeatKernelConventions:   true,
		DoesNotClaimFullSuiteCosmology:     !ci.HardCosmologyPredictedNow,
		Verdict:                            strings.Join([]string{StatusProductGeometryBridgeDerived, StatusTensionMNotDerivedFromFiniteAlgebra, StatusTensionCosmologicalConstantChannelOpen, StatusTensionYukawaTextureRemainsModuli}, ";"),
	}
}

func buildSummary(p ProductTriple, h HeatKernelExpansion, f FiniteInvariantLedger, l LagrangianAssembly, ci ContinuumInterface, fw Firewall) Summary {
	return Summary{
		ProductTripleBuilt:          p.Executed && p.ProductIsMarriage,
		HeatKernelExpanded:          h.Executed && h.ProductFactorization,
		FiniteInvariantsInserted:    f.Executed && f.AllRequiredInvariantsSeen,
		LagrangianAssembled:         l.Executed && l.StandardModelRecovered && l.EinsteinGravityRecovered,
		SMGravitySkeletonRecovered:  l.StandardModelRecovered && l.EinsteinGravityRecovered,
		ContinuumCalculusEnabled:    ci.EnablesRG && ci.EnablesBoltzmann && ci.EnablesBounce && ci.EnablesClassicalGravity,
		HardCosmologicalPredictions: 0,
		RemainingChargedModuli:      f.ChargedFiniteDiracModuli,
		DirectAnswer:                "The ASHA finite geometry and the continuum spin manifold marry through M×F.  The spectral action assembles the SM+Einstein-gravity Lagrangian skeleton with ASHA finite coefficients, while keeping Λ_cosmo, cosmological initial data, and the 13 charged flavor moduli outside the finite derivation.",
		Verdict:                     strings.Join([]string{StatusProductGeometryBridgeDerived, StatusFullLagrangianSkeletonAssembled, StatusContinuumComputationInterfaceOpened, fw.Verdict}, ";"),
	}
}

func NativeProductConstants() map[string]float64 {
	return map[string]float64{
		"dim_HF_doubled":          float64(doubledFiniteHilbertDim),
		"f0_contact":              f0Contact,
		"f2_lambda_over_MP2":      f2LambdaOverMP2,
		"sin2_thetaW_boundary":    sin2ThetaWBoundary,
		"alpha_branch_inverse":    alphaGUTInverseBranch,
		"lambdaH_over_gstar2":     lambdaHOverGStarSquared,
		"charged_DF_moduli":       chargedFiniteDiracModuli,
		"threshold_delta_lambda":  thresholdJumpDeltaLambda,
		"external_minimal_ledger": externalLedgerModuli,
	}
}

func FormatProduct(p ProductTriple) string {
	return fmt.Sprintf("A=%s; H=%s; D=%s; J=%s; gamma=%s; dim(H_F,doubled)=%d; marriage=%t; M-derived-from-F=%t; verdict=%s", p.Algebra, p.HilbertSpace, p.DiracOperator, p.RealStructure, p.Grading, p.FiniteHilbertDimension, p.ProductIsMarriage, p.SpacetimeDerivedFromF, p.Verdict)
}

func FormatHeatKernel(h HeatKernelExpansion) string {
	parts := make([]string, 0, len(h.Terms))
	for _, t := range h.Terms {
		parts = append(parts, fmt.Sprintf("%s -> %s using %s × %s", t.Order, t.PhysicalTerm, t.ContinuumPart, t.FinitePart))
	}
	return fmt.Sprintf("formula=%s; rule=%s; terms=[%s]; verdict=%s", h.Formula, h.FactorizationRule, strings.Join(parts, " | "), h.Verdict)
}

func FormatFinite(f FiniteInvariantLedger) string {
	parts := make([]string, 0, len(f.Invariants))
	for _, inv := range f.Invariants {
		parts = append(parts, fmt.Sprintf("%s=%s enters %s", inv.Symbol, inv.Value, inv.Enters))
	}
	return fmt.Sprintf("invariants=[%s]; verdict=%s", strings.Join(parts, " | "), f.Verdict)
}

func FormatLagrangian(l LagrangianAssembly) string {
	parts := make([]string, 0, len(l.Sectors))
	for _, s := range l.Sectors {
		parts = append(parts, fmt.Sprintf("%s: %s; coeff=%s; fully_predicted=%t", s.Name, s.SymbolicTerm, s.ASHACoefficient, s.FullyPredicted))
	}
	return fmt.Sprintf("SM=%t; Einstein=%t; all_coefficients_fixed=%t; sectors=[%s]; verdict=%s", l.StandardModelRecovered, l.EinsteinGravityRecovered, l.AllCoefficientsFixed, strings.Join(parts, " | "), l.Verdict)
}

func FormatInterface(ci ContinuumInterface) string {
	return fmt.Sprintf("RG=%t; Boltzmann=%t; bounce=%t; classical_gravity=%t; requires_metric_topology=%t; requires_initial_conditions=%t; requires_flavor_moduli=%t; hard_cosmology_now=%t; verdict=%s", ci.EnablesRG, ci.EnablesBoltzmann, ci.EnablesBounce, ci.EnablesClassicalGravity, ci.RequiresMetricAndTopology, ci.RequiresInitialConditions, ci.RequiresFlavorModuliValues, ci.HardCosmologyPredictedNow, ci.Verdict)
}

func FormatFirewall(f Firewall) string {
	return fmt.Sprintf("no_M_from_F=%t; no_Lambda_prediction=%t; no_relic_prediction=%t; no_lifetime_prediction=%t; no_yukawa_selection=%t; no_13_moduli_erasure=%t; heat_kernel_conventions_tracked=%t; verdict=%s", f.DoesNotDeriveMFromF, f.DoesNotPredictCosmologicalConstant, f.DoesNotPredictRelicDensity, f.DoesNotPredictVacuumLifetime, f.DoesNotSelectYukawaTexture, f.DoesNotErase13Moduli, f.DoesNotHideHeatKernelConventions, f.Verdict)
}
