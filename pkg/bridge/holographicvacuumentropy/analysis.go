// Package holographicvacuumentropy implements Gate 373:
// Holographic Vacuum Entropy / Gravitational Moduli Constraint Sieve.
//
// The gate asks whether ASHA's macroscopic gravitational data provide actual
// equations on the native finite-Dirac flavor moduli.  It deliberately keeps
// separate:
//  1. scale constraints already derived by the gravitational spectral action;
//  2. symbolic Yukawa invariants entering a possible vacuum-energy functional;
//  3. holographic/Bekenstein/asymptotic-safety inequalities;
//  4. exact coordinate-fixing equations capable of reducing the 13 charged
//     finite-Dirac moduli found in Gate 372.
package holographicvacuumentropy

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate372 "github.com/bagherbal/asha-engine/pkg/bridge/nativemodulispacecensus"
)

const (
	AuditID = "GATE373-HOLOGRAPHIC-VACUUM-ENTROPY-GRAVITATIONAL-MODULI-CONSTRAINT-SIEVE"

	StatusGate372Inherited                 = "CONDITIONAL_SUPPORT_GATE372_NATIVE_MODULI_CENSUS_INHERITED"
	StatusGravitationalBoundaryFormalized  = "CONDITIONAL_SUPPORT_GRAVITATIONAL_BOUNDARY_FORMALIZED"
	StatusPfaffianHierarchyInherited       = "CONDITIONAL_SUPPORT_PFAFFIAN_HIERARCHY_SCALE_INHERITED"
	StatusFlavorInvariantLedgerConstructed = "CONDITIONAL_SUPPORT_FLAVOR_INVARIANT_LEDGER_CONSTRUCTED"
	StatusVacuumEnergyFunctionalFormalized = "CONDITIONAL_SUPPORT_VACUUM_ENERGY_FUNCTIONAL_FORMALIZED_SYMBOLICALLY"
	StatusTraceAnomalySieveExecuted        = "CONDITIONAL_SUPPORT_TRACE_ANOMALY_SIEVE_EXECUTED"
	StatusHolographicBoundAudited          = "CONDITIONAL_SUPPORT_HOLOGRAPHIC_BOUND_AUDITED"
	StatusBekensteinBoundAudited           = "CONDITIONAL_SUPPORT_BEKENSTEIN_BOUND_AUDITED"
	StatusAsymptoticSafetyAudited          = "CONDITIONAL_SUPPORT_ASYMPTOTIC_SAFETY_CONSTRAINT_AUDITED"
	StatusInformationHorizonAudited        = "CONDITIONAL_SUPPORT_INFORMATION_HORIZON_AUDITED"
	StatusModuliCensusUpdated              = "CONDITIONAL_SUPPORT_GRAVITATIONAL_MODULI_CENSUS_UPDATED"
	StatusLandscapePreservationAudited     = "CONDITIONAL_SUPPORT_LANDSCAPE_PRESERVATION_AUDITED"
	StatusEpistemicFirewallAudited         = "CONDITIONAL_SUPPORT_EPISTEMIC_FIREWALL_AUDITED"
	StatusNoHolographicReductionFound      = "CONDITIONAL_SUPPORT_NO_HOLOGRAPHIC_REDUCTION_FOUND_IN_CURRENT_LEDGER"

	StatusTensionGravityFixesScaleNotTexture        = "CONDITIONAL_TENSION_GRAVITY_BOUNDARY_FIXES_SCALE_NOT_FLAVOR_TEXTURE"
	StatusTensionVacuumEnergyNeedsCounterterm       = "CONDITIONAL_TENSION_VACUUM_ENERGY_FUNCTIONAL_NEEDS_RENORMALIZED_COUNTERTERM"
	StatusTensionHolographicBoundIsInequality       = "CONDITIONAL_TENSION_HOLOGRAPHIC_BOUND_IS_AGGREGATE_INEQUALITY_NOT_TEXTURE_EQUATIONS"
	StatusTensionTraceAnomalyHasTooFewInvariants    = "CONDITIONAL_TENSION_TRACE_ANOMALY_SEES_AGGREGATE_TRACES_NOT_FULL_13_COORDINATES"
	StatusTensionHorizonRadiusNotFiniteNative       = "CONDITIONAL_TENSION_HORIZON_RADIUS_NOT_FINITE_NATIVE_INPUT"
	StatusTensionAsymptoticSafetyNeedsBetaLedger    = "CONDITIONAL_TENSION_ASYMPTOTIC_SAFETY_REQUIRES_CONTINUUM_BETA_LEDGER"
	StatusTensionInformationOperatorStillUnselected = "CONDITIONAL_TENSION_INFORMATION_NUMBER_OPERATOR_STILL_NOT_SELECTED"
	StatusTensionNoEqualitySaturationTheorem        = "CONDITIONAL_TENSION_NO_NATIVE_EQUALITY_SATURATION_THEOREM_FOR_HOLOGRAPHIC_BOUND"

	StatusFailedHolographicConstraintNotDerived = "FAILED_ROUTE_HOLOGRAPHIC_MODULI_CONSTRAINT_NOT_DERIVED"
	StatusFailedGravityBoundTooWeak             = "FAILED_ROUTE_GRAVITATIONAL_BOUND_TOO_WEAK_TO_FIX_13_MODULI"
	StatusFailedVacuumEnergyNotUnique           = "FAILED_ROUTE_VACUUM_ENERGY_FUNCTIONAL_NOT_UNIQUELY_DERIVED"
	StatusFailedTraceAnomalyNoTexture           = "FAILED_ROUTE_TRACE_ANOMALY_DOES_NOT_FIX_FLAVOR_TEXTURE"
	StatusFailedInformationBoundaryNotDerived   = "FAILED_ROUTE_INFORMATION_HORIZON_NUMBER_OPERATOR_BOUNDARY_NOT_DERIVED"
	StatusFailedModuliNotReduced                = "FAILED_ROUTE_NATIVE_MODULI_NOT_REDUCED_BY_GRAVITY"
	StatusFailedVacuumPointNotSelected          = "FAILED_ROUTE_PHYSICAL_VACUUM_POINT_STILL_NOT_SELECTED"
	StatusFailedYukawaStillFree                 = "FAILED_ROUTE_YUKAWA_COORDINATES_STILL_FREE_AFTER_HOLOGRAPHIC_AUDIT"
	StatusFailedCKMStillFree                    = "FAILED_ROUTE_CKM_TEXTURE_STILL_FREE_AFTER_HOLOGRAPHIC_AUDIT"
)

const (
	ChargedFlavorModuli         = 13
	ExternalMinimalVacuumLedger = 15
)

type Inheritance struct {
	Executed             bool
	HighestInheritedGate int
	NativeChargedModuli  int
	ExternalLedger       int
	PreviousTruth        string
	Question             string
	Verdict              string
}

type GravitationalBoundary struct {
	Executed                  bool
	F2LambdaOverPlanckSquared float64
	F2Formula                 string
	VEVOverPlanck             float64
	VEVFormula                string
	FixesAbsoluteScale        bool
	FixesFlavorTexture        bool
	NativeFlavorEquations     int
	Interpretation            string
	Verdict                   string
}

type FlavorInvariant struct {
	Name                       string
	Symbol                     string
	DependsOnCoordinates       []string
	CoordinateCountVisible     int
	FlavorBasisInvariant       bool
	SensitiveToCKMMisalignment bool
	RequiresExtraCoefficient   bool
	Verdict                    string
}

type VacuumEnergyFunctional struct {
	Executed                     bool
	SymbolicForm                 string
	Invariants                   []FlavorInvariant
	CountertermRequired          bool
	RenormalizationScaleRequired bool
	UniqueNativeFunctional       bool
	IndependentFlavorEquations   int
	CKMTextureEquations          int
	CanFixAbsoluteVacuumEnergy   bool
	Verdict                      string
}

type ConstraintLane struct {
	Lane                        string
	Name                        string
	Formula                     string
	InputType                   string
	NativeASHA                  bool
	InequalityOnly              bool
	RequiresContinuumData       bool
	RequiresEmpiricalInput      bool
	RequiresSaturationPostulate bool
	IndependentFlavorEquations  int
	CanReduce13Moduli           bool
	CanSelectVacuumPoint        bool
	Reason                      string
	Verdict                     string
}

type HolographicAudit struct {
	Executed                        bool
	Lanes                           []ConstraintLane
	TotalIndependentFlavorEquations int
	AnyTextureConstraint            bool
	AnyVacuumSelection              bool
	DirectAnswer                    string
	Verdict                         string
}

type InformationAudit struct {
	Executed                        bool
	UsesGate371NumberOperator       bool
	NumberOperatorSelectedByGravity bool
	EntropyFunctional               string
	HorizonActsAsGenerationAddress  bool
	IndependentFlavorEquations      int
	ThermalTimeActivated            bool
	DirectAnswer                    string
	Verdict                         string
}

type Census struct {
	Executed                bool
	StartingChargedModuli   int
	GravitationalEquations  int
	Reduction               int
	RemainingChargedModuli  int
	ExternalLedger          int
	ExternalLedgerReduction int
	FinalStatement          string
	Verdict                 string
}

type Firewall struct {
	Executed                       bool
	NoObservedMassesImported       bool
	NoObservedYukawasImported      bool
	NoCKMValuesImported            bool
	NoPMNSValuesImported           bool
	NoCosmologicalConstantImported bool
	NoHiggsMassTargetImported      bool
	NoSaturationAssumed            bool
	NoContinuumBetaFunctionsFitted bool
	LandscapeRatiosPreserved       bool
	Verdict                        string
}

type Analysis struct {
	Inheritance  Inheritance
	Gravity      GravitationalBoundary
	VacuumEnergy VacuumEnergyFunctional
	Holography   HolographicAudit
	Information  InformationAudit
	Census       Census
	Firewall     Firewall
	Truth        string
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
	prev, err := gate372.BuildDefault()
	if err != nil {
		return Analysis{}, err
	}
	inheritance := inherit(prev)
	gravity := formalizeGravity()
	vacuum := formalizeVacuumEnergy()
	holography := auditHolographicConstraints(gravity, vacuum)
	info := auditInformationHorizon(holography)
	census := updateCensus(holography, info)
	firewall := auditFirewall()
	truth := "Gate 373 tests the grand pivot from finite kinematics to holographic/gravitational thermodynamics.  The ASHA gravitational data f2(Λ/M_P)^2=π/64 and v/M_P=2^(3/2)e^(-4π^2) are real scale constraints, but in the current ledger they do not provide independent equations on the 13 charged finite-Dirac flavor moduli.  A vacuum-energy or trace-anomaly functional can be written symbolically in terms of Yukawa invariants such as T2 and T4, but its cosmological value requires a renormalized counterterm, a continuum scale choice, and at most aggregate trace data; it does not determine the full flavor texture or CKM misalignment.  Bekenstein/holographic bounds are inequalities unless a separate saturation theorem is derived, and asymptotic-safety constraints require a continuum beta-function ledger not present in this finite gate.  The Gate-371 information/number-operator direction remains a capacity witness, not a gravity-selected generation address.  Therefore no holographic reduction of the 13 moduli is derived, and the physical vacuum point remains unselected."
	return Analysis{inheritance, gravity, vacuum, holography, info, census, firewall, truth}, nil
}

func inherit(prev gate372.Analysis) Inheritance {
	return Inheritance{Executed: true, HighestInheritedGate: 372, NativeChargedModuli: prev.Native.MinimalChargedDFDim, ExternalLedger: prev.Native.External15, PreviousTruth: prev.Truth, Question: "do ASHA gravitational/holographic constraints impose native equations on the 13 finite-Dirac charged flavor moduli?", Verdict: join(StatusGate372Inherited)}
}

func formalizeGravity() GravitationalBoundary {
	f2 := math.Pi / 64
	vev := math.Pow(2, 1.5) * math.Exp(-4*math.Pi*math.Pi)
	return GravitationalBoundary{
		Executed:                  true,
		F2LambdaOverPlanckSquared: f2,
		F2Formula:                 "f2 (Lambda/M_P)^2 = pi/64",
		VEVOverPlanck:             vev,
		VEVFormula:                "v/M_P = 2^(3/2) exp(-4 pi^2)",
		FixesAbsoluteScale:        true,
		FixesFlavorTexture:        false,
		NativeFlavorEquations:     0,
		Interpretation:            "the gravitational ledger fixes scale relations and hierarchy normalization, not a generation-dependent Yukawa texture",
		Verdict:                   join(StatusGravitationalBoundaryFormalized, StatusPfaffianHierarchyInherited, StatusTensionGravityFixesScaleNotTexture),
	}
}

func formalizeVacuumEnergy() VacuumEnergyFunctional {
	invariants := []FlavorInvariant{
		{Name: "quadratic Yukawa trace", Symbol: "T2=Tr(Y_u†Y_u+Y_d†Y_d+Y_e†Y_e)", DependsOnCoordinates: []string{"9 charged singular values"}, CoordinateCountVisible: 9, FlavorBasisInvariant: true, SensitiveToCKMMisalignment: false, RequiresExtraCoefficient: false, Verdict: StatusFlavorInvariantLedgerConstructed},
		{Name: "quartic Yukawa trace", Symbol: "T4=Tr((Y_u†Y_u)^2+(Y_d†Y_d)^2+(Y_e†Y_e)^2)", DependsOnCoordinates: []string{"9 charged singular values"}, CoordinateCountVisible: 9, FlavorBasisInvariant: true, SensitiveToCKMMisalignment: false, RequiresExtraCoefficient: false, Verdict: StatusFlavorInvariantLedgerConstructed},
		{Name: "quark misalignment commutator", Symbol: "C_ud=Tr([Y_uY_u†,Y_dY_d†]^2)", DependsOnCoordinates: []string{"6 quark singular values", "4 CKM parameters"}, CoordinateCountVisible: 10, FlavorBasisInvariant: true, SensitiveToCKMMisalignment: true, RequiresExtraCoefficient: true, Verdict: join(StatusFlavorInvariantLedgerConstructed, StatusTensionTraceAnomalyHasTooFewInvariants)},
		{Name: "cosmological counterterm", Symbol: "rho_0", DependsOnCoordinates: []string{"renormalized vacuum subtraction"}, CoordinateCountVisible: 0, FlavorBasisInvariant: true, SensitiveToCKMMisalignment: false, RequiresExtraCoefficient: true, Verdict: StatusTensionVacuumEnergyNeedsCounterterm},
	}
	return VacuumEnergyFunctional{
		Executed:                     true,
		SymbolicForm:                 "rho_vac(Y)=rho_0 + A v^2 T2 + B v^4 T4 + C v^4 C_ud + ...",
		Invariants:                   invariants,
		CountertermRequired:          true,
		RenormalizationScaleRequired: true,
		UniqueNativeFunctional:       false,
		IndependentFlavorEquations:   0,
		CKMTextureEquations:          0,
		CanFixAbsoluteVacuumEnergy:   false,
		Verdict:                      join(StatusVacuumEnergyFunctionalFormalized, StatusTraceAnomalySieveExecuted, StatusTensionVacuumEnergyNeedsCounterterm, StatusTensionTraceAnomalyHasTooFewInvariants, StatusFailedVacuumEnergyNotUnique, StatusFailedTraceAnomalyNoTexture),
	}
}

func auditHolographicConstraints(g GravitationalBoundary, v VacuumEnergyFunctional) HolographicAudit {
	lanes := []ConstraintLane{
		lane("A", "ASHA cutoff moment", g.F2Formula, "native gravitational scale", true, false, false, false, false, g.NativeFlavorEquations, false, false, "fixes the cutoff/Planck normalization; it contains no Yukawa-coordinate equation", join(StatusGravitationalBoundaryFormalized, StatusTensionGravityFixesScaleNotTexture)),
		lane("B", "Pfaffian hierarchy scale", g.VEVFormula, "native hierarchy scale", true, false, false, false, false, 0, false, false, "fixes v/M_P but not the dimensionless flavor ratios inside D_F", join(StatusPfaffianHierarchyInherited, StatusTensionGravityFixesScaleNotTexture)),
		lane("C", "vacuum-energy trace functional", v.SymbolicForm, "symbolic finite-to-continuum energy", false, false, true, false, false, v.IndependentFlavorEquations, false, false, "requires rho_0, renormalization scheme, and coefficients before becoming equations; trace aggregates cannot determine all 13 coordinates", join(StatusVacuumEnergyFunctionalFormalized, StatusTraceAnomalySieveExecuted, StatusTensionVacuumEnergyNeedsCounterterm, StatusFailedVacuumEnergyNotUnique)),
		lane("D", "Bekenstein entropy bound", "S <= 2 pi E R", "holographic inequality", false, true, true, false, true, 0, false, false, "without a native horizon radius and saturation theorem, this is an inequality on aggregate energy/information rather than a texture equation", join(StatusBekensteinBoundAudited, StatusHolographicBoundAudited, StatusTensionHolographicBoundIsInequality, StatusTensionHorizonRadiusNotFiniteNative, StatusTensionNoEqualitySaturationTheorem)),
		lane("E", "covariant holographic area bound", "S <= A/(4G)", "gravitational entropy inequality", false, true, true, false, true, 0, false, false, "bounds total entropy; it does not assign generation weights or CKM angles", join(StatusHolographicBoundAudited, StatusTensionHolographicBoundIsInequality, StatusFailedGravityBoundTooWeak)),
		lane("F", "asymptotic-safety fixed-point condition", "beta_Y(Y,g,G_N)=0", "continuum RG equation family", false, false, true, false, false, 0, false, false, "could become dynamical in a later continuum ledger, but this gate has no native gravitational beta functions or threshold data", join(StatusAsymptoticSafetyAudited, StatusTensionAsymptoticSafetyNeedsBetaLedger)),
		lane("G", "de Sitter/AdS stability pressure", "delta rho_vac(Y) compatible with stable background", "stability criterion", false, true, true, false, true, 0, false, false, "stability pressure is not an equality selecting the flavor point unless a native extremization principle is proven", join(StatusHolographicBoundAudited, StatusTensionHolographicBoundIsInequality, StatusTensionNoEqualitySaturationTheorem)),
	}
	total := 0
	anyTexture := false
	anyVacuum := false
	for _, l := range lanes {
		total += l.IndependentFlavorEquations
		anyTexture = anyTexture || l.CanReduce13Moduli
		anyVacuum = anyVacuum || l.CanSelectVacuumPoint
	}
	return HolographicAudit{Executed: true, Lanes: lanes, TotalIndependentFlavorEquations: total, AnyTextureConstraint: anyTexture, AnyVacuumSelection: anyVacuum, DirectAnswer: "no current holographic/gravitational lane supplies a native equality constraint on the 13 charged finite-Dirac moduli", Verdict: join(StatusHolographicBoundAudited, StatusBekensteinBoundAudited, StatusAsymptoticSafetyAudited, StatusNoHolographicReductionFound, StatusFailedHolographicConstraintNotDerived, StatusFailedGravityBoundTooWeak)}
}

func lane(id, name, formula, inputType string, native, inequality, continuum, empirical, saturation bool, equations int, reduce, selectVacuum bool, reason, verdict string) ConstraintLane {
	return ConstraintLane{Lane: id, Name: name, Formula: formula, InputType: inputType, NativeASHA: native, InequalityOnly: inequality, RequiresContinuumData: continuum, RequiresEmpiricalInput: empirical, RequiresSaturationPostulate: saturation, IndependentFlavorEquations: equations, CanReduce13Moduli: reduce, CanSelectVacuumPoint: selectVacuum, Reason: reason, Verdict: verdict}
}

func auditInformationHorizon(h HolographicAudit) InformationAudit {
	return InformationAudit{
		Executed:                        true,
		UsesGate371NumberOperator:       true,
		NumberOperatorSelectedByGravity: false,
		EntropyFunctional:               "S_N(rho)=-Tr(rho log rho), with rho_N=exp(-beta N)/Z only after choosing N",
		HorizonActsAsGenerationAddress:  false,
		IndependentFlavorEquations:      0,
		ThermalTimeActivated:            false,
		DirectAnswer:                    "the horizon can bound aggregate entropy, but the current ledger does not derive N or a generation-address operator from gravity",
		Verdict:                         join(StatusInformationHorizonAudited, StatusTensionInformationOperatorStillUnselected, StatusFailedInformationBoundaryNotDerived),
	}
}

func updateCensus(h HolographicAudit, info InformationAudit) Census {
	reduction := h.TotalIndependentFlavorEquations + info.IndependentFlavorEquations
	remaining := ChargedFlavorModuli - reduction
	if remaining < 0 {
		remaining = 0
	}
	return Census{Executed: true, StartingChargedModuli: ChargedFlavorModuli, GravitationalEquations: h.TotalIndependentFlavorEquations + info.IndependentFlavorEquations, Reduction: reduction, RemainingChargedModuli: remaining, ExternalLedger: ExternalMinimalVacuumLedger, ExternalLedgerReduction: reduction, FinalStatement: "Gate 373 reduces zero charged finite-Dirac flavor moduli; the minimal 13+theta+scale ledger is not compressed by the current gravitational/holographic audit", Verdict: join(StatusModuliCensusUpdated, StatusNoHolographicReductionFound, StatusFailedModuliNotReduced, StatusFailedVacuumPointNotSelected, StatusFailedYukawaStillFree, StatusFailedCKMStillFree)}
}

func auditFirewall() Firewall {
	return Firewall{Executed: true, NoObservedMassesImported: true, NoObservedYukawasImported: true, NoCKMValuesImported: true, NoPMNSValuesImported: true, NoCosmologicalConstantImported: true, NoHiggsMassTargetImported: true, NoSaturationAssumed: true, NoContinuumBetaFunctionsFitted: true, LandscapeRatiosPreserved: true, Verdict: join(StatusLandscapePreservationAudited, StatusEpistemicFirewallAudited)}
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("highest_gate=%d native_charged_moduli=%d external_ledger=%d question=%s verdict=%s", x.HighestInheritedGate, x.NativeChargedModuli, x.ExternalLedger, x.Question, x.Verdict)
}

func FormatGravity(x GravitationalBoundary) string {
	return fmt.Sprintf("%s %.12g; %s %.12g; fixes_scale=%t fixes_texture=%t flavor_equations=%d verdict=%s", x.F2Formula, x.F2LambdaOverPlanckSquared, x.VEVFormula, x.VEVOverPlanck, x.FixesAbsoluteScale, x.FixesFlavorTexture, x.NativeFlavorEquations, x.Verdict)
}

func FormatVacuumEnergy(x VacuumEnergyFunctional) string {
	parts := []string{fmt.Sprintf("form=%s", x.SymbolicForm)}
	for _, inv := range x.Invariants {
		parts = append(parts, fmt.Sprintf("%s[%s visible=%d ckm=%t extra_coeff=%t]", inv.Name, inv.Symbol, inv.CoordinateCountVisible, inv.SensitiveToCKMMisalignment, inv.RequiresExtraCoefficient))
	}
	parts = append(parts, fmt.Sprintf("unique=%t counterterm=%t scale=%t flavor_eqs=%d ckm_eqs=%d verdict=%s", x.UniqueNativeFunctional, x.CountertermRequired, x.RenormalizationScaleRequired, x.IndependentFlavorEquations, x.CKMTextureEquations, x.Verdict))
	return strings.Join(parts, "; ")
}

func FormatLane(x ConstraintLane) string {
	return fmt.Sprintf("lane=%s name=%s formula=%s native=%t inequality=%t continuum=%t saturation=%t flavor_eqs=%d reduce=%t select=%t reason=%s verdict=%s", x.Lane, x.Name, x.Formula, x.NativeASHA, x.InequalityOnly, x.RequiresContinuumData, x.RequiresSaturationPostulate, x.IndependentFlavorEquations, x.CanReduce13Moduli, x.CanSelectVacuumPoint, x.Reason, x.Verdict)
}

func FormatHolography(x HolographicAudit) string {
	parts := []string{fmt.Sprintf("total_flavor_equations=%d any_texture=%t any_vacuum=%t", x.TotalIndependentFlavorEquations, x.AnyTextureConstraint, x.AnyVacuumSelection)}
	for _, lane := range x.Lanes {
		parts = append(parts, FormatLane(lane))
	}
	parts = append(parts, "answer="+x.DirectAnswer, "verdict="+x.Verdict)
	return strings.Join(parts, "\n")
}

func FormatInformation(x InformationAudit) string {
	return fmt.Sprintf("uses_gate371_N=%t gravity_selects_N=%t horizon_generation_address=%t flavor_eqs=%d thermal_time=%t entropy=%s answer=%s verdict=%s", x.UsesGate371NumberOperator, x.NumberOperatorSelectedByGravity, x.HorizonActsAsGenerationAddress, x.IndependentFlavorEquations, x.ThermalTimeActivated, x.EntropyFunctional, x.DirectAnswer, x.Verdict)
}

func FormatCensus(x Census) string {
	return fmt.Sprintf("start_13=%d gravitational_eqs=%d reduction=%d remaining=%d external_ledger=%d external_reduction=%d final=%s verdict=%s", x.StartingChargedModuli, x.GravitationalEquations, x.Reduction, x.RemainingChargedModuli, x.ExternalLedger, x.ExternalLedgerReduction, x.FinalStatement, x.Verdict)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("no_masses=%t no_yukawas=%t no_ckm=%t no_pmns=%t no_lambda_cosmo=%t no_higgs_target=%t no_saturation=%t no_beta_fit=%t landscape=%t verdict=%s", x.NoObservedMassesImported, x.NoObservedYukawasImported, x.NoCKMValuesImported, x.NoPMNSValuesImported, x.NoCosmologicalConstantImported, x.NoHiggsMassTargetImported, x.NoSaturationAssumed, x.NoContinuumBetaFunctionsFitted, x.LandscapeRatiosPreserved, x.Verdict)
}

func join(parts ...string) string { return strings.Join(parts, "; ") }
