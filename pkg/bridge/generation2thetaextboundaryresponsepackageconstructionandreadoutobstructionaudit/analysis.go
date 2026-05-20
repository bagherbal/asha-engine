// Package generation2thetaextboundaryresponsepackageconstructionandreadoutobstructionaudit implements
// Gate 785: ThetaExt Boundary Response Package Construction and Readout Obstruction Audit.
//
// Gate 784 identified Theta_ext: M_n -> Lambda^(n-1)B_boundary as the sharp
// missing object behind the boundary cubic-stop candidate. Gate 785 audits the
// stronger requirement: a full boundary response package consisting of a raw
// moment lift, an exterior-degree assignment, a scalar readout, and an
// orientation/sign convention. This gate is forensic only; it does not derive
// scalar runtime lambda, Higgs pole mass, Yukawa operators, PMNS, CKM, flavor
// hierarchy, G_F, VEV, or a native HistoryLoopUnit theorem.
package generation2thetaextboundaryresponsepackageconstructionandreadoutobstructionaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE785-THETAEXT-BOUNDARY-RESPONSE-PACKAGE-CONSTRUCTION-READOUT-OBSTRUCTION-AUDIT"

	StatusGate784BoundaryExteriorDegreeResponseInherited = "PASS_GATE784_BOUNDARY_EXTERIOR_DEGREE_RESPONSE_INHERITED"
	StatusLiftAndReadoutProblemsSeparated                = "PASS_LIFT_AND_READOUT_PROBLEMS_SEPARATED"
	StatusExteriorResponseAlgebraTyped                   = "PASS_EXTERIOR_RESPONSE_ALGEBRA_TYPED"
	StatusConditionalThetaExtResponsePackageConstructed  = "PASS_CONDITIONAL_THETA_EXT_RESPONSE_PACKAGE_CONSTRUCTED"
	StatusNaturalityAuditCompleted                       = "PASS_NATURALITY_AUDIT_COMPLETED"
	StatusMagnitudeAndSignSeparated                      = "PASS_MAGNITUDE_AND_SIGN_SEPARATED"
	StatusCubicStopUnderPackageAudited                   = "PASS_CUBIC_STOP_UNDER_PACKAGE_AUDITED"
	StatusExteriorExponentialShortcutAudited             = "PASS_EXTERIOR_EXPONENTIAL_SHORTCUT_AUDITED"
	StatusPredictionStatusReclassified                   = "PASS_PREDICTION_STATUS_RECLASSIFIED"
	StatusPhysicalFirewallsEnforced                      = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusLabelledBoundaryPairConditionalExteriorAlgebra  = "CONDITIONAL_SUPPORT_LABELLED_BOUNDARY_PAIR_CAN_DEFINE_CONDITIONAL_EXTERIOR_ALGEBRA"
	StatusFWall3RedRepresentedByExteriorResponsePackage   = "CONDITIONAL_SUPPORT_F_WALL_3_RED_CAN_BE_REPRESENTED_BY_EXTERIOR_RESPONSE_PACKAGE"
	Status2PMagnitudeBoundaryPairTimesK7EventSourceType   = "CONDITIONAL_SUPPORT_2P_MAGNITUDE_HAS_BOUNDARY_PAIR_TIMES_K7_EVENT_SOURCE_TYPE"
	StatusM4BlockedIfThetaExtDegreeRuleSupplied           = "CONDITIONAL_SUPPORT_M4_IS_BLOCKED_IF_THETA_EXT_DEGREE_RULE_IS_SUPPLIED"
	StatusFWall3RedConditionalExteriorResponsePackageForm = "CONDITIONAL_SUPPORT_F_WALL_3_RED_HAS_CONDITIONAL_EXTERIOR_RESPONSE_PACKAGE_FORM"

	StatusThetaExtAloneDoesNotDeriveResponsePolynomial                = "FAILED_ROUTE_THETA_EXT_ALONE_DOES_NOT_DERIVE_RESPONSE_POLYNOMIAL"
	StatusNoNativeCanonicalDegreeOneBoundaryAxisOrReadout             = "FAILED_ROUTE_NO_NATIVE_CANONICAL_DEGREE_ONE_BOUNDARY_AXIS_OR_READOUT"
	StatusConditionalPackageRepackagesCoefficientsUnlessReadoutNative = "FAILED_ROUTE_CONDITIONAL_PACKAGE_REPACKAGES_COEFFICIENTS_UNLESS_READOUT_IS_NATIVE"
	StatusDimensionTwoAloneDoesNotDefineBetaBOrChiExt                 = "FAILED_ROUTE_DIMENSION_TWO_ALONE_DOES_NOT_DEFINE_BETA_B_OR_CHI_EXT"
	StatusNoCanonicalNonzeroDegreeOneResponseFromAbstractBBoundary    = "FAILED_ROUTE_NO_CANONICAL_NONZERO_DEGREE_ONE_RESPONSE_FROM_ABSTRACT_B_BOUNDARY"
	StatusNoNativeOrientationSignForNegativeCubicTerm                 = "FAILED_ROUTE_NO_NATIVE_ORIENTATION_SIGN_FOR_NEGATIVE_CUBIC_TERM"
	StatusCubicStopNotDerivedWithoutNativeDegreeRule                  = "FAILED_ROUTE_CUBIC_STOP_IS_NOT_DERIVED_WITHOUT_NATIVE_DEGREE_RULE"
	StatusNoNativeCubicStopTheorem                                    = "FAILED_ROUTE_NO_NATIVE_CUBIC_STOP_THEOREM"
	StatusSingleBoundaryVectorExponentialDoesNotSourceDegreeTwoTerm   = "FAILED_ROUTE_SINGLE_BOUNDARY_VECTOR_EXPONENTIAL_DOES_NOT_SOURCE_DEGREE_TWO_TERM"
	StatusDegreeTwoTermRequiresBoundaryPairProductOrVolumeReadout     = "FAILED_ROUTE_DEGREE_TWO_TERM_REQUIRES_BOUNDARY_PAIR_PRODUCT_OR_VOLUME_READOUT"
	StatusFWall3RedNotLevelCNativeComponent                           = "FAILED_ROUTE_F_WALL_3_RED_NOT_LEVEL_C_NATIVE_COMPONENT"
	StatusCHistoryNotYetFullIndependentPredictionComponent            = "FAILED_ROUTE_C_HISTORY_NOT_YET_FULL_INDEPENDENT_PREDICTION_COMPONENT"
	StatusTreeProxyNotPoleMass                                        = "FAILED_ROUTE_TREE_PROXY_NOT_POLE_MASS"
	StatusNoYukawaOperatorOrEigenvalueTheorem                         = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusFirewallPreservedGate785                                    = "FIREWALL_PRESERVED_GATE785_THETA_EXT_RESPONSE_PACKAGE_BOUNDARY"
)

const (
	pK7Snapshot       = 7.0 / 72.0
	sSplitSnapshot    = 0.0012924448188162962
	kappaERedSnapshot = 0.005503554218475772
	fWall3Snapshot    = 0.00012565521035653708
)

type Gate784Inheritance struct {
	Inherited     bool
	MissingObject string
	PriorLevel    string
	Verdict       string
}

type MomentLedger struct {
	P       float64
	S       float64
	M1      float64
	M2      float64
	M3      float64
	M4      float64
	KappaE  float64
	FWall3  float64
	Matches bool
}

type LiftReadoutSeparation struct {
	Separated                bool
	ThetaExtDomain           string
	ThetaExtCodomain         string
	ChiExtDomain             string
	ChiExtCodomain           string
	PolynomialRepresentation string
	ThetaExtAloneSufficient  bool
	Verdict                  string
}

type ExteriorResponseAlgebra struct {
	Typed                    bool
	Algebra                  string
	Basis                    []string
	ScalarBasis              string
	VolumeForm               string
	RequiresDegreeOneAxis    bool
	HasNativeDegreeOneAxis   bool
	ConditionalLabelledBasis bool
	Verdict                  string
}

type ConditionalPackage struct {
	Constructed         bool
	ThetaM1             string
	ThetaM2             string
	ThetaM3             string
	ThetaHigher         string
	Chi0                float64
	Chi1                float64
	Chi2                float64
	ReconstructedFWall3 float64
	MatchesFWall3       bool
	ReadoutNative       bool
	Verdict             string
}

type NaturalityAudit struct {
	Completed                             bool
	OnlyDimensionTwoNative                bool
	CanonicalNonzeroVectorFromDimension   bool
	CanonicalNonzeroCovectorFromDimension bool
	LabelledBasisConditional              bool
	Verdict                               string
}

type MagnitudeSignAudit struct {
	Separated             bool
	TwoP                  float64
	MagnitudeSource       string
	Sign                  string
	OrientationSignNative bool
	Verdict               string
}

type CubicStopPackageAudit struct {
	Audited             bool
	M4                  float64
	M4Degree            string
	BlockedIfDegreeRule bool
	DegreeRuleNative    bool
	CubicStopNative     bool
	Verdict             string
}

type ExteriorExponentialAudit struct {
	Audited                       bool
	SingleVectorExp               string
	SingleVectorProducesDegreeTwo bool
	DegreeTwoRequirement          string
	Verdict                       string
}

type PredictionStatus struct {
	Reclassified bool
	FWall3Status string
	KappaLambda  string
	CHistory     string
	CHiggs       string
	FWallLevelC  bool
	Verdict      string
}

type Firewalls struct {
	Enforced              bool
	ThetaExtPackageNative bool
	ConditionalBetaNative bool
	ChiExtNative          bool
	DimensionTwoProof     bool
	OmegaBSignDerived     bool
	FWallNative           bool
	KappaLambdaNative     bool
	CHistoryIndependent   bool
	TreeProxyPoleMass     bool
	YukawaNative          bool
	Verdict               string
}

type Audit struct {
	Gate784        Gate784Inheritance
	Ledger         MomentLedger
	Separation     LiftReadoutSeparation
	Algebra        ExteriorResponseAlgebra
	Package        ConditionalPackage
	Naturality     NaturalityAudit
	MagnitudeSign  MagnitudeSignAudit
	CubicStop      CubicStopPackageAudit
	Exponential    ExteriorExponentialAudit
	Prediction     PredictionStatus
	Firewalls      Firewalls
	Truth          string
	FinalStatement string
}

var (
	defaultOnce  sync.Once
	defaultAudit Audit
	defaultErr   error
)

func BuildDefault() (Audit, error) {
	defaultOnce.Do(func() { defaultAudit, defaultErr = build() })
	return defaultAudit, defaultErr
}

func build() (Audit, error) {
	p := pK7Snapshot
	s := sSplitSnapshot
	m1 := p * s
	m2 := p * s * s
	m3 := p * s * s * s
	m4 := p * math.Pow(s, 4)
	fwall := m1 + kappaERedSnapshot*m2 - 2*p*m3
	if math.IsNaN(fwall) || math.IsInf(fwall, 0) || math.IsNaN(m4) || math.IsInf(m4, 0) {
		return Audit{}, fmt.Errorf("non-finite Gate785 raw-moment ledger")
	}
	matches := closeRel(fwall, fWall3Snapshot, 1e-14)

	a := Audit{
		Gate784: Gate784Inheritance{
			Inherited:     true,
			MissingObject: "Theta_ext: raw moment layer -> exterior boundary degree",
			PriorLevel:    "F_wall_3_red was Level B+ source-typed boundary response candidate, not native.",
			Verdict:       StatusGate784BoundaryExteriorDegreeResponseInherited,
		},
		Ledger: MomentLedger{P: p, S: s, M1: m1, M2: m2, M3: m3, M4: m4, KappaE: kappaERedSnapshot, FWall3: fwall, Matches: matches},
		Separation: LiftReadoutSeparation{
			Separated:                true,
			ThetaExtDomain:           "raw moment layers M_n",
			ThetaExtCodomain:         "Lambda^(n-1) B_boundary",
			ChiExtDomain:             "Lambda^0 B_boundary ⊕ Lambda^1 B_boundary ⊕ Lambda^2 B_boundary",
			ChiExtCodomain:           "R",
			PolynomialRepresentation: "F_wall_3_red=chi_ext[Theta_ext(M1)+Theta_ext(M2)+Theta_ext(M3)]",
			ThetaExtAloneSufficient:  false,
			Verdict:                  StatusThetaExtAloneDoesNotDeriveResponsePolynomial,
		},
		Algebra: ExteriorResponseAlgebra{
			Typed:                    true,
			Algebra:                  "E_boundary=Lambda^0 B_boundary ⊕ Lambda^1 B_boundary ⊕ Lambda^2 B_boundary",
			Basis:                    []string{"b_lambda <-> |lambda(Lambda12)|", "b_R <-> R3-1"},
			ScalarBasis:              "1_B in Lambda^0 B_boundary",
			VolumeForm:               "omega_B=b_lambda∧b_R in Lambda^2 B_boundary after labelled/ordered boundary basis",
			RequiresDegreeOneAxis:    true,
			HasNativeDegreeOneAxis:   false,
			ConditionalLabelledBasis: true,
			Verdict:                  StatusLabelledBoundaryPairConditionalExteriorAlgebra,
		},
		Package: ConditionalPackage{
			Constructed:         true,
			ThetaM1:             "Theta_ext(M1)=M1·1_B",
			ThetaM2:             "Theta_ext(M2)=M2·beta_B",
			ThetaM3:             "Theta_ext(M3)=M3·omega_B",
			ThetaHigher:         "Theta_ext(M_n>=4)=0 if the degree rule M_n -> Lambda^(n-1)B_boundary is supplied",
			Chi0:                1,
			Chi1:                kappaERedSnapshot,
			Chi2:                -2 * p,
			ReconstructedFWall3: fwall,
			MatchesFWall3:       matches,
			ReadoutNative:       false,
			Verdict:             StatusFWall3RedRepresentedByExteriorResponsePackage,
		},
		Naturality: NaturalityAudit{
			Completed:                             true,
			OnlyDimensionTwoNative:                true,
			CanonicalNonzeroVectorFromDimension:   false,
			CanonicalNonzeroCovectorFromDimension: false,
			LabelledBasisConditional:              true,
			Verdict:                               StatusNoCanonicalNonzeroDegreeOneResponseFromAbstractBBoundary,
		},
		MagnitudeSign: MagnitudeSignAudit{
			Separated:             true,
			TwoP:                  2 * p,
			MagnitudeSource:       "2p=dim(B_boundary)*p_K7=2*(7/72)=7/36",
			Sign:                  "negative sign requires chi_ext(omega_B)=-2p, i.e. an orientation/stress-pull convention",
			OrientationSignNative: false,
			Verdict:               Status2PMagnitudeBoundaryPairTimesK7EventSourceType,
		},
		CubicStop: CubicStopPackageAudit{
			Audited:             true,
			M4:                  m4,
			M4Degree:            "M4 would map to Lambda^3 B_boundary=0 under the supplied degree rule",
			BlockedIfDegreeRule: true,
			DegreeRuleNative:    false,
			CubicStopNative:     false,
			Verdict:             StatusM4BlockedIfThetaExtDegreeRuleSupplied,
		},
		Exponential: ExteriorExponentialAudit{
			Audited:                       true,
			SingleVectorExp:               "exp(beta)=1+beta because beta∧beta=0 for a single boundary vector",
			SingleVectorProducesDegreeTwo: false,
			DegreeTwoRequirement:          "degree-two term requires two distinct boundary legs, an explicitly supplied volume form omega_B, or a boundary-pair product/readout",
			Verdict:                       StatusSingleBoundaryVectorExponentialDoesNotSourceDegreeTwoTerm,
		},
		Prediction: PredictionStatus{
			Reclassified: true,
			FWall3Status: "Level B+ exterior-response package representation; formula-independent but not Level C native component.",
			KappaLambda:  "Level B formula-independent scalar complement; not native.",
			CHistory:     "Level B semi-independent History correction; not full independent prediction component.",
			CHiggs:       "still not Level C.",
			FWallLevelC:  false,
			Verdict:      StatusFWall3RedConditionalExteriorResponsePackageForm,
		},
		Firewalls: Firewalls{
			Enforced:              true,
			ThetaExtPackageNative: false,
			ConditionalBetaNative: false,
			ChiExtNative:          false,
			DimensionTwoProof:     false,
			OmegaBSignDerived:     false,
			FWallNative:           false,
			KappaLambdaNative:     false,
			CHistoryIndependent:   false,
			TreeProxyPoleMass:     false,
			YukawaNative:          false,
			Verdict:               StatusFirewallPreservedGate785,
		},
	}
	a.Truth = "Gate 785 shows that Theta_ext alone is insufficient; F_wall_3_red needs a full exterior response package Theta_ext + chi_ext + degree-one axis/readout + orientation/sign."
	a.FinalStatement = "Gate 785 does not construct Theta_ext natively. It conditionally constructs the full exterior response package needed to represent F_wall_3_red, and shows that the missing object is not only Theta_ext but also the scalar readout chi_ext, the degree-one boundary axis/readout, and the orientation/sign of the degree-two stress-pull term. The next bottleneck is native sourcing of the exterior response package: Theta_ext + chi_ext + boundary orientation/sign + degree-one modulation axis."
	return a, nil
}

func Statuses() []string {
	return []string{
		StatusGate784BoundaryExteriorDegreeResponseInherited,
		StatusLiftAndReadoutProblemsSeparated,
		StatusExteriorResponseAlgebraTyped,
		StatusConditionalThetaExtResponsePackageConstructed,
		StatusNaturalityAuditCompleted,
		StatusMagnitudeAndSignSeparated,
		StatusCubicStopUnderPackageAudited,
		StatusExteriorExponentialShortcutAudited,
		StatusPredictionStatusReclassified,
		StatusPhysicalFirewallsEnforced,
		StatusLabelledBoundaryPairConditionalExteriorAlgebra,
		StatusFWall3RedRepresentedByExteriorResponsePackage,
		Status2PMagnitudeBoundaryPairTimesK7EventSourceType,
		StatusM4BlockedIfThetaExtDegreeRuleSupplied,
		StatusFWall3RedConditionalExteriorResponsePackageForm,
		StatusThetaExtAloneDoesNotDeriveResponsePolynomial,
		StatusNoNativeCanonicalDegreeOneBoundaryAxisOrReadout,
		StatusConditionalPackageRepackagesCoefficientsUnlessReadoutNative,
		StatusDimensionTwoAloneDoesNotDefineBetaBOrChiExt,
		StatusNoCanonicalNonzeroDegreeOneResponseFromAbstractBBoundary,
		StatusNoNativeOrientationSignForNegativeCubicTerm,
		StatusCubicStopNotDerivedWithoutNativeDegreeRule,
		StatusNoNativeCubicStopTheorem,
		StatusSingleBoundaryVectorExponentialDoesNotSourceDegreeTwoTerm,
		StatusDegreeTwoTermRequiresBoundaryPairProductOrVolumeReadout,
		StatusFWall3RedNotLevelCNativeComponent,
		StatusCHistoryNotYetFullIndependentPredictionComponent,
		StatusTreeProxyNotPoleMass,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusFirewallPreservedGate785,
	}
}

func closeRel(got, want, tol float64) bool {
	if math.IsNaN(got) || math.IsInf(got, 0) || math.IsNaN(want) || math.IsInf(want, 0) {
		return false
	}
	d := math.Abs(got - want)
	if d <= tol {
		return true
	}
	den := math.Max(1, math.Abs(want))
	return d/den <= tol
}

func containsAll(haystack []string, needles []string) bool {
	joined := strings.Join(haystack, "\n")
	for _, n := range needles {
		if !strings.Contains(joined, n) {
			return false
		}
	}
	return true
}

func FormatLedger(l MomentLedger) string {
	return fmt.Sprintf("p=%.17g s=%.17g M1=%.17g M2=%.17g M3=%.17g M4=%.17g F_wall_3_red=%.17g matches=%v", l.P, l.S, l.M1, l.M2, l.M3, l.M4, l.FWall3, l.Matches)
}

func FormatSeparation(s LiftReadoutSeparation) string {
	return fmt.Sprintf("Theta_ext: %s -> %s; chi_ext: %s -> %s; theta_alone_sufficient=%v", s.ThetaExtDomain, s.ThetaExtCodomain, s.ChiExtDomain, s.ChiExtCodomain, s.ThetaExtAloneSufficient)
}

func FormatPackage(p ConditionalPackage) string {
	return fmt.Sprintf("%s; %s; %s; chi=(%.17g, %.17g, %.17g); F=%.17g native_readout=%v", p.ThetaM1, p.ThetaM2, p.ThetaM3, p.Chi0, p.Chi1, p.Chi2, p.ReconstructedFWall3, p.ReadoutNative)
}

func FormatPrediction(p PredictionStatus) string {
	return fmt.Sprintf("F_wall=%s; kappa_lambda=%s; C_History=%s; C_Higgs=%s", p.FWall3Status, p.KappaLambda, p.CHistory, p.CHiggs)
}
