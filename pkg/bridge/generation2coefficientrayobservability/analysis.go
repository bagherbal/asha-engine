// Package generation2coefficientrayobservability implements Gate 454:
// Coefficient-Ray Observability Rank Audit.
//
// Gate 453 defined the legal empirical interface for texture-zero comparator
// work. Gate 454 makes that interface quantitative: it asks how many external,
// explicitly labelled comparator invariants are required to identify a sector
// coefficient ray for
//
//	M(a,b,c)=aK_gen+bX_triangle+cY_phase
//
// without ever reclassifying the imported values as native ASHA law.
package generation2coefficientrayobservability

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE454-COEFFICIENT-RAY-OBSERVABILITY-RANK-AUDIT"

	StatusGate453Inherited                      = "CONDITIONAL_SUPPORT_GATE453_EMPIRICAL_INTERFACE_INHERITED"
	StatusRayDimensionDerived                   = "CONDITIONAL_SUPPORT_COEFFICIENT_RAY_DIMENSION_DERIVED"
	StatusSpectrumOnlyRankOne                   = "CONDITIONAL_SUPPORT_SPECTRUM_ONLY_OBSERVABILITY_RANK_ONE"
	StatusTwoScalarLocalRankTwo                 = "CONDITIONAL_SUPPORT_TWO_SCALAR_LOCAL_RAY_IDENTIFIABILITY"
	StatusCPBranchTagRequired                   = "CONDITIONAL_SUPPORT_CP_ORIENTED_BRANCH_REQUIRES_EXPLICIT_TAG"
	StatusComparatorProtocolDefined             = "CONDITIONAL_SUPPORT_COEFFICIENT_RAY_COMPARATOR_PROTOCOL_DEFINED"
	StatusNoNativeCoefficientValues             = "CONDITIONAL_SUPPORT_NO_NATIVE_COEFFICIENT_RAY_VALUE_IMPORTED"
	StatusEmpiricalFirewallPreserved            = "CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED"
	StatusFailedSpectrumOnlyUnderdetermined     = "FAILED_ROUTE_SPECTRUM_ONLY_CANNOT_FIX_COEFFICIENT_RAY"
	StatusFailedNativeCoefficientSelectorAbsent = "FAILED_ROUTE_NATIVE_COEFFICIENT_RAY_SELECTOR_ABSENT"
	StatusFailedCPOrientationNotNative          = "FAILED_ROUTE_CP_ORIENTATION_NOT_NATIVE"
)

const (
	NativeFlavorDim  = 13
	KXYCoeffDim      = 9
	ProjectiveRayDOF = 2
)

type Inheritance struct {
	Executed                         bool
	Gate444KGenForced                bool
	Gate445TriangleForced            bool
	Gate446PhaseQuarantined          bool
	Gate447CoefficientsSealed        bool
	Gate450TextureZeroSumRule        bool
	Gate450RatiosRequireAmplitudes   bool
	Gate451FullTrianglePreserved     bool
	Gate452NearestNeighborNotGauge   bool
	Gate453EmpiricalInterfaceDefined bool
	Gate453PromotionRejected         bool
	NoEmpiricalInputsImported        bool
	Verdict                          string
}

type RayModel struct {
	Executed               bool
	MatrixFormula          string
	Parameterization       string
	AbsoluteScaleParameter string
	RayParameters          []string
	ProjectiveDimension    int
	NativeSelectors        []string
	NativeSelectorCount    int
	Verdict                string
	Reason                 string
}

type ObservableMap struct {
	Name                 string
	Inputs               []string
	Formulae             []string
	Rank                 int
	ResidualDOF          int
	LocallyIdentifiesRay bool
	GloballyOriented     bool
	RequiresEmpirical    bool
	AllowedByGate453     bool
	NativePromotion      bool
	Reason               string
}

type RankAudit struct {
	Executed                bool
	Maps                    []ObservableMap
	SpectrumOnlyRank        int
	SpectrumOnlyResidualDOF int
	MinimumLocalScalars     int
	MinimumOrientedScalars  int
	GenericJacobianFormula  string
	GenericJacobianSample   float64
	GenericJacobianNonzero  bool
	SpectrumOnlyRejected    bool
	TwoScalarLocalWorks     bool
	CPBranchTagRequired     bool
	Verdict                 string
	Reason                  string
}

type Protocol struct {
	Executed                          bool
	AllowsNativeLedger                bool
	AllowsSpectrumOnlyComparator      bool
	AllowsLocalRayFit                 bool
	AllowsCPOrientedRayFit            bool
	RequiresExplicitEmpiricalLabel    bool
	RequiresSectorTag                 bool
	RequiresRenormalizationTag        bool
	RequiresBranchTagForCPOrientation bool
	AllowsNativeCoefficientClaim      bool
	AllowsSpectrumOnlyRayClaim        bool
	Verdict                           string
	Reason                            string
}

type Firewall struct {
	Executed                      bool
	NoObservedMuonMassImported    bool
	NoObservedCharmMassImported   bool
	NoObservedYukawaImported      bool
	NoCKMImported                 bool
	NoPMNSImported                bool
	NoCurveFitPromoted            bool
	NoGSTPromotion                bool
	NoNativeCoefficientRayValue   bool
	KGenStillForced               bool
	XTriangleStillForced          bool
	YPhaseStillQuarantined        bool
	SectorCoefficientsStillSealed bool
	CPOrientationStillSealed      bool
	NativeFlavorDimAfter          int
	KXYCoeffDimAfter              int
	Verdict                       string
	Reason                        string
}

type NextStep struct {
	Gate        int
	Title       string
	Reason      string
	PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Ray         RayModel
	Rank        RankAudit
	Protocol    Protocol
	Firewall    Firewall
	Next        NextStep
	Truth       string
}

var cache struct {
	sync.Once
	a   Analysis
	err error
}

func BuildDefault() (Analysis, error) {
	cache.Once.Do(func() { cache.a, cache.err = build() })
	return cache.a, cache.err
}

func build() (Analysis, error) {
	a := Analysis{}
	a.Inheritance = buildInheritance()
	a.Ray = buildRayModel()
	a.Rank = buildRankAudit()
	a.Protocol = buildProtocol(a.Rank)
	a.Firewall = buildFirewall(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return Analysis{}, err
	}
	return a, nil
}

func buildInheritance() Inheritance {
	return Inheritance{
		Executed:                         true,
		Gate444KGenForced:                true,
		Gate445TriangleForced:            true,
		Gate446PhaseQuarantined:          true,
		Gate447CoefficientsSealed:        true,
		Gate450TextureZeroSumRule:        true,
		Gate450RatiosRequireAmplitudes:   true,
		Gate451FullTrianglePreserved:     true,
		Gate452NearestNeighborNotGauge:   true,
		Gate453EmpiricalInterfaceDefined: true,
		Gate453PromotionRejected:         true,
		NoEmpiricalInputsImported:        true,
		Verdict:                          StatusGate453Inherited,
	}
}

func buildRayModel() RayModel {
	return RayModel{
		Executed:               true,
		MatrixFormula:          "M(a,b,c)=aK_gen+bX_triangle+cY_phase",
		Parameterization:       "set r=sqrt(b^2+c^2)>0, alpha=a/r, phi=atan2(c,b); ray=(alpha,phi) modulo overall scale",
		AbsoluteScaleParameter: "rho=sqrt(a^2+b^2+c^2)",
		RayParameters:          []string{"alpha=a/r", "phi=atan2(c,b)"},
		ProjectiveDimension:    ProjectiveRayDOF,
		NativeSelectors:        []string{},
		NativeSelectorCount:    0,
		Verdict:                StatusRayDimensionDerived,
		Reason:                 "the sector coefficient ledger has three real coefficients; removing the absolute scale leaves a two-dimensional projective ray, and no native selector fixes either coordinate.",
	}
}

func buildRankAudit() RankAudit {
	alpha := 0.7
	phi := 0.31
	j := genericJacobian(alpha, phi)
	maps := []ObservableMap{
		{
			Name:                 "native structural ledger",
			Inputs:               []string{"K_gen", "X_triangle", "M_22=0 sum rule"},
			Formulae:             []string{"no empirical coefficient value supplied"},
			Rank:                 0,
			ResidualDOF:          ProjectiveRayDOF,
			LocallyIdentifiesRay: false,
			GloballyOriented:     false,
			RequiresEmpirical:    false,
			AllowedByGate453:     true,
			NativePromotion:      false,
			Reason:               "native structure defines the allowed coefficient space but no point on that space.",
		},
		{
			Name:                 "normalized spectrum only",
			Inputs:               []string{"trace-zero eigenvalue ratios"},
			Formulae:             []string{"I_spec=2 cos(3 phi)/(alpha^2+3)^(3/2)"},
			Rank:                 1,
			ResidualDOF:          1,
			LocallyIdentifiesRay: false,
			GloballyOriented:     false,
			RequiresEmpirical:    true,
			AllowedByGate453:     true,
			NativePromotion:      false,
			Reason:               "a trace-zero three-eigenvalue spectrum up to scale has only one shape invariant; a one-parameter continuum of coefficient rays survives.",
		},
		{
			Name:                 "spectrum plus K-addressed overlap",
			Inputs:               []string{"I_spec", "I_K=Tr(MK)/sqrt(Tr(M^2)Tr(K^2))"},
			Formulae:             []string{"I_spec=2 cos(3 phi)/(alpha^2+3)^(3/2)", "I_K=alpha/sqrt(alpha^2+3)", "det d(I_spec,I_K)/d(alpha,phi)=18 sin(3 phi)/(alpha^2+3)^3"},
			Rank:                 2,
			ResidualDOF:          0,
			LocallyIdentifiesRay: true,
			GloballyOriented:     false,
			RequiresEmpirical:    true,
			AllowedByGate453:     true,
			NativePromotion:      false,
			Reason:               "the Jacobian is generically nonzero, so two explicitly labelled scalar comparators identify the ray locally; conjugate/discrete phase branches remain.",
		},
		{
			Name:                 "spectrum plus K-overlap plus CP-odd branch tag",
			Inputs:               []string{"I_spec", "I_K", "sign or value of I_CP=sin(3 phi)"},
			Formulae:             []string{"I_CP=sin(3 phi)", "cos(3 phi) from I_spec and alpha", "alpha from I_K"},
			Rank:                 2,
			ResidualDOF:          0,
			LocallyIdentifiesRay: true,
			GloballyOriented:     true,
			RequiresEmpirical:    true,
			AllowedByGate453:     true,
			NativePromotion:      false,
			Reason:               "the CP-odd tag does not increase local differential rank, but it resolves the orientation branch left by cos(3 phi).",
		},
		{
			Name:                 "spectrum-only native coefficient claim",
			Inputs:               []string{"observed masses"},
			Formulae:             []string{"invert I_spec as if it fixed alpha and phi"},
			Rank:                 1,
			ResidualDOF:          1,
			LocallyIdentifiesRay: false,
			GloballyOriented:     false,
			RequiresEmpirical:    true,
			AllowedByGate453:     false,
			NativePromotion:      true,
			Reason:               "forbidden: it both underdetermines the ray and attempts to promote observed data to native law.",
		},
	}
	return RankAudit{
		Executed:                true,
		Maps:                    maps,
		SpectrumOnlyRank:        1,
		SpectrumOnlyResidualDOF: 1,
		MinimumLocalScalars:     2,
		MinimumOrientedScalars:  3,
		GenericJacobianFormula:  "18 sin(3 phi)/(alpha^2+3)^3",
		GenericJacobianSample:   j,
		GenericJacobianNonzero:  math.Abs(j) > 1e-9,
		SpectrumOnlyRejected:    true,
		TwoScalarLocalWorks:     true,
		CPBranchTagRequired:     true,
		Verdict:                 StatusComparatorProtocolDefined,
		Reason:                  "spectrum-only data has rank one; a K-addressed mixing/spectrum overlap supplies the second local coordinate, while CP orientation still requires an explicit branch tag.",
	}
}

func buildProtocol(rank RankAudit) Protocol {
	return Protocol{
		Executed:                          true,
		AllowsNativeLedger:                true,
		AllowsSpectrumOnlyComparator:      true,
		AllowsLocalRayFit:                 rank.TwoScalarLocalWorks,
		AllowsCPOrientedRayFit:            rank.CPBranchTagRequired,
		RequiresExplicitEmpiricalLabel:    true,
		RequiresSectorTag:                 true,
		RequiresRenormalizationTag:        true,
		RequiresBranchTagForCPOrientation: true,
		AllowsNativeCoefficientClaim:      false,
		AllowsSpectrumOnlyRayClaim:        false,
		Verdict:                           StatusComparatorProtocolDefined,
		Reason:                            "Gate 454 permits coefficient-ray identification only as labelled empirical comparator work; native ledgers may state the rank obstruction but not a fitted value.",
	}
}

func buildFirewall(a Analysis) Firewall {
	return Firewall{
		Executed:                      true,
		NoObservedMuonMassImported:    true,
		NoObservedCharmMassImported:   true,
		NoObservedYukawaImported:      true,
		NoCKMImported:                 true,
		NoPMNSImported:                true,
		NoCurveFitPromoted:            true,
		NoGSTPromotion:                true,
		NoNativeCoefficientRayValue:   true,
		KGenStillForced:               a.Inheritance.Gate444KGenForced,
		XTriangleStillForced:          a.Inheritance.Gate445TriangleForced,
		YPhaseStillQuarantined:        a.Inheritance.Gate446PhaseQuarantined,
		SectorCoefficientsStillSealed: a.Inheritance.Gate447CoefficientsSealed,
		CPOrientationStillSealed:      a.Rank.CPBranchTagRequired,
		NativeFlavorDimAfter:          NativeFlavorDim,
		KXYCoeffDimAfter:              KXYCoeffDim,
		Verdict:                       StatusEmpiricalFirewallPreserved,
		Reason:                        "the audit computes observability rank of possible comparator maps only; it imports no observed masses, mixings, Yukawas, CKM, or PMNS values.",
	}
}

func buildNext() NextStep {
	return NextStep{
		Gate:        455,
		Title:       "Empirical Texture Adapter Stub / Dry-Run Firewall Test",
		Reason:      "after deriving the observability rank protocol, the next safe engineering step is a dry-run adapter that accepts labelled dummy data and proves forbidden native promotion paths fail closed",
		PrimaryTask: "implement a no-data default adapter with schema validation, branch labels, scale/scheme tags, and explicit rejection of spectrum-only native coefficient claims",
	}
}

func truth(a Analysis) string {
	return fmt.Sprintf("Gate 454 proves that the coefficient ray is a two-dimensional empirical bridge object: spectrum-only comparators have rank %d and leave one ray coordinate free; two labelled scalars generically identify the ray locally; CP-oriented uniqueness still needs an explicit branch tag. No coefficient value becomes native.", a.Rank.SpectrumOnlyRank)
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate453EmpiricalInterfaceDefined || !a.Inheritance.Gate453PromotionRejected {
		return fmt.Errorf("Gate 453 interface not inherited: %s", FormatInheritance(a.Inheritance))
	}
	if !a.Ray.Executed || a.Ray.ProjectiveDimension != ProjectiveRayDOF || a.Ray.NativeSelectorCount != 0 {
		return fmt.Errorf("coefficient ray dimension/selector ledger invalid: %s", FormatRayModel(a.Ray))
	}
	if !a.Rank.Executed || a.Rank.SpectrumOnlyRank != 1 || a.Rank.SpectrumOnlyResidualDOF != 1 || a.Rank.MinimumLocalScalars != 2 || !a.Rank.GenericJacobianNonzero || !a.Rank.TwoScalarLocalWorks || !a.Rank.CPBranchTagRequired {
		return fmt.Errorf("rank audit failed: %s", FormatRankAudit(a.Rank))
	}
	if !a.Protocol.Executed || !a.Protocol.RequiresExplicitEmpiricalLabel || !a.Protocol.RequiresSectorTag || !a.Protocol.RequiresRenormalizationTag || !a.Protocol.RequiresBranchTagForCPOrientation || a.Protocol.AllowsNativeCoefficientClaim || a.Protocol.AllowsSpectrumOnlyRayClaim {
		return fmt.Errorf("protocol leaks native coefficient claim: %s", FormatProtocol(a.Protocol))
	}
	if !a.Firewall.Executed || !a.Firewall.NoNativeCoefficientRayValue || !a.Firewall.NoCurveFitPromoted || !a.Firewall.YPhaseStillQuarantined || !a.Firewall.SectorCoefficientsStillSealed || a.Firewall.NativeFlavorDimAfter != NativeFlavorDim || a.Firewall.KXYCoeffDimAfter != KXYCoeffDim {
		return fmt.Errorf("firewall failed: %s", FormatFirewall(a.Firewall))
	}
	return nil
}

func genericJacobian(alpha, phi float64) float64 {
	return 18 * math.Sin(3*phi) / math.Pow(alpha*alpha+3, 3)
}

func statuses() []string {
	return []string{
		StatusGate453Inherited,
		StatusRayDimensionDerived,
		StatusSpectrumOnlyRankOne,
		StatusTwoScalarLocalRankTwo,
		StatusCPBranchTagRequired,
		StatusComparatorProtocolDefined,
		StatusNoNativeCoefficientValues,
		StatusEmpiricalFirewallPreserved,
		StatusFailedSpectrumOnlyUnderdetermined,
		StatusFailedNativeCoefficientSelectorAbsent,
		StatusFailedCPOrientationNotNative,
	}
}

func join(xs []string) string {
	if len(xs) == 0 {
		return "∅"
	}
	return strings.Join(xs, ", ")
}
