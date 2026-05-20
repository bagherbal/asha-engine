// Package generation2boundaryexteriordegreeresponsemapandcubicstopsourceaudit implements
// Gate 784: Boundary Exterior-Degree Response Map and Cubic Stop Source Audit.
//
// Gate 783 source-typed F_wall_3_red as a cubic raw-moment boundary response
// expectation but did not certify why the series should stop at cubic order.
// Gate 784 audits the strongest current stop candidate: a typed map from raw
// boundary moment layers M_n to exterior boundary response degrees in the
// two-dimensional boundary pair B_boundary. This gate is forensic only: it does
// not derive scalar runtime lambda, Higgs pole mass, Yukawa operators, PMNS,
// CKM, flavor hierarchy, G_F, VEV, or a native HistoryLoopUnit theorem.
package generation2boundaryexteriordegreeresponsemapandcubicstopsourceaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE784-BOUNDARY-EXTERIOR-DEGREE-RESPONSE-MAP-CUBIC-STOP-SOURCE-AUDIT"

	StatusGate783BoundaryRawMomentGeneratingFunctionInherited       = "PASS_GATE783_BOUNDARY_RAW_MOMENT_GENERATING_FUNCTION_INHERITED"
	StatusFWall3RedBoundarySubBottleneckInherited                   = "PASS_F_WALL_3_RED_IDENTIFIED_AS_BOUNDARY_SUB_BOTTLENECK"
	StatusBoundaryPairObjectTyped                                   = "PASS_BOUNDARY_PAIR_OBJECT_TYPED"
	StatusExteriorDegreeCandidateAudited                            = "PASS_EXTERIOR_DEGREE_CANDIDATE_AUDITED"
	StatusRequiredExteriorDegreeMapIdentified                       = "PASS_REQUIRED_EXTERIOR_DEGREE_MAP_IDENTIFIED"
	StatusCubicCoefficientCompatibleWithBoundaryPairDegreeTwoSource = "PASS_CUBIC_COEFFICIENT_COMPATIBLE_WITH_BOUNDARY_PAIR_DEGREE_TWO_SOURCE"
	StatusDegreeByDegreeResponseTableRecorded                       = "PASS_DEGREE_BY_DEGREE_RESPONSE_TABLE_RECORDED"
	StatusProjectorIdempotenceFirewallPreserved                     = "PASS_PROJECTOR_IDEMPOTENCE_FIREWALL_PRESERVED"
	StatusM4RejectionReaudited                                      = "PASS_M4_REJECTION_REAUDITED"
	StatusGeneratingFunctionCandidateRecorded                       = "PASS_GENERATING_FUNCTION_CANDIDATE_RECORDED"
	StatusRelationToKappaLambdaRedRecorded                          = "PASS_RELATION_TO_KAPPA_LAMBDA_RED_RECORDED"
	StatusPredictionLevelClassificationRecorded                     = "PASS_PREDICTION_LEVEL_CLASSIFICATION_RECORDED"
	StatusPhysicalFirewallsEnforced                                 = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusBoundaryPairTwoDimensionalBridgeCarrier                 = "CONDITIONAL_SUPPORT_BOUNDARY_PAIR_IS_TWO_DIMENSIONAL_BRIDGE_CARRIER"
	StatusCubicStopBoundaryExteriorDegreeSourceCandidate          = "CONDITIONAL_SUPPORT_CUBIC_STOP_HAS_BOUNDARY_EXTERIOR_DEGREE_SOURCE_CANDIDATE"
	Status2PBoundaryPairTimesK7EventWeight                        = "CONDITIONAL_SUPPORT_2P_IS_BOUNDARY_PAIR_TIMES_K7_EVENT_WEIGHT"
	StatusKappaERedM2DegreeOneFlavorBoundaryModulation            = "CONDITIONAL_SUPPORT_KAPPA_E_RED_M2_IS_DEGREE_ONE_FLAVOR_BOUNDARY_MODULATION"
	StatusM4BlockedIfThetaExtSupplied                             = "CONDITIONAL_SUPPORT_M4_IS_BLOCKED_IF_RAW_MOMENT_TO_EXTERIOR_DEGREE_MAP_IS_SUPPLIED"
	StatusFWall3RedCubicTruncationExteriorDegreeResponseCandidate = "CONDITIONAL_SUPPORT_F_WALL_3_RED_IS_CUBIC_TRUNCATION_OF_BOUNDARY_EXTERIOR_DEGREE_RESPONSE_CANDIDATE"
	StatusBoundaryExteriorDegreeMapWouldUpgradeFWallSourceStatus  = "CONDITIONAL_SUPPORT_BOUNDARY_EXTERIOR_DEGREE_MAP_WOULD_UPGRADE_F_WALL_SOURCE_STATUS"
	StatusFWall3RedUpgradedToLevelBPlusSourceCandidate            = "CONDITIONAL_SUPPORT_F_WALL_3_RED_UPGRADED_TO_LEVEL_B_PLUS_SOURCE_CANDIDATE"

	StatusBoundaryPairNotNativeSpacetimeOrFlavorCarrier      = "FAILED_ROUTE_BOUNDARY_PAIR_NOT_NATIVE_SPACETIME_OR_FLAVOR_CARRIER"
	StatusNoNativeThetaExtMap                                = "FAILED_ROUTE_NO_NATIVE_THETA_EXT_MAP_FROM_RAW_MOMENTS_TO_BOUNDARY_EXTERIOR_DEGREES"
	StatusDimensionTwoBoundaryPairAloneDoesNotProveCubicStop = "FAILED_ROUTE_DIMENSION_TWO_BOUNDARY_PAIR_ALONE_DOES_NOT_PROVE_CUBIC_STOP"
	StatusNoNativeSignTheoremForNegativeCubicStressPull      = "FAILED_ROUTE_NO_NATIVE_SIGN_THEOREM_FOR_NEGATIVE_CUBIC_STRESS_PULL"
	StatusNoNativeBoundaryPairStressPullTheorem              = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_PAIR_STRESS_PULL_THEOREM"
	StatusNoNativeFlavorBoundaryModulationTheorem            = "FAILED_ROUTE_NO_NATIVE_FLAVOR_BOUNDARY_MODULATION_THEOREM"
	StatusProjectorIdempotenceDoesNotForceCubicStop          = "FAILED_ROUTE_PROJECTOR_IDEMPOTENCE_DOES_NOT_FORCE_CUBIC_STOP"
	StatusNoNativeTypedM4CoefficientSource                   = "FAILED_ROUTE_NO_NATIVE_TYPED_M4_COEFFICIENT_SOURCE"
	StatusNoNativeCubicStopTheorem                           = "FAILED_ROUTE_NO_NATIVE_CUBIC_STOP_THEOREM"
	StatusNoNativeBoundaryResponseGeneratingFunctionTheorem  = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_RESPONSE_GENERATING_FUNCTION_THEOREM"
	StatusFWall3RedNotLevelCNativeComponent                  = "FAILED_ROUTE_F_WALL_3_RED_NOT_LEVEL_C_NATIVE_COMPONENT"
	StatusKappaLambdaRedNotYetNativeScalarMatchingTheorem    = "FAILED_ROUTE_KAPPA_LAMBDA_RED_NOT_YET_NATIVE_SCALAR_MATCHING_THEOREM"
	StatusCHistoryNotYetFullIndependentPredictionComponent   = "FAILED_ROUTE_C_HISTORY_NOT_YET_FULL_INDEPENDENT_PREDICTION_COMPONENT"
	StatusTreeProxyNotPoleMass                               = "FAILED_ROUTE_TREE_PROXY_NOT_POLE_MASS"
	StatusNoYukawaOperatorOrEigenvalueTheorem                = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusFirewallPreservedGate784                           = "FIREWALL_PRESERVED_GATE784_BOUNDARY_EXTERIOR_DEGREE_RESPONSE_BOUNDARY"
)

const (
	pK7Snapshot       = 7.0 / 72.0
	sSplitSnapshot    = 0.0012924448188162962
	kappaERedSnapshot = 0.005503554218475772
	fWall3Snapshot    = 0.00012565521035653708
)

type Gate783Inheritance struct {
	Inherited             bool
	ResponseFunction      string
	BoundarySubBottleneck string
	Verdict               string
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

type BoundaryPairAudit struct {
	Typed                bool
	Carrier              string
	Basis                []string
	ScalarSplit          string
	MidpointStress       string
	Dim                  int
	ExteriorDimensions   []int
	IsSpacetimeCarrier   bool
	IsFlavorCarrier      bool
	IsK7Carrier          bool
	TwoDimensionalBridge bool
	Verdict              string
}

type ExteriorDegreeCandidateAudit struct {
	Audited                    bool
	DegreeAssignments          map[string]string
	Lambda3Zero                bool
	CubicStopCandidate         bool
	DimensionAloneProvesStop   bool
	NativeThetaExtMapCertified bool
	Verdict                    string
}

type RequiredMapAudit struct {
	Identified                bool
	MapName                   string
	Domain                    string
	Codomain                  string
	Assignments               []string
	Certified                 bool
	DimensionShortcutRejected bool
	Verdict                   string
}

type CubicCoefficientAudit struct {
	Audited               bool
	TwoP                  float64
	MagnitudeSource       string
	DegreeTwoLayer        string
	SignSourceCertified   bool
	StressPullTheorem     bool
	CompatibleWithDegree2 bool
	Verdict               string
}

type DegreeResponseTable struct {
	Recorded                       bool
	Degree0                        string
	Degree1                        string
	Degree2                        string
	Degree3                        string
	KappaERedNaturallyDegree1      bool
	FlavorBoundaryModulationNative bool
	Verdict                        string
}

type ProjectorFirewallAudit struct {
	Preserved               bool
	PowerLaw                string
	ProjectorStopsExpansion bool
	NeedsExteriorMap        bool
	Verdict                 string
}

type M4RejectionAudit struct {
	Audited                 bool
	M4                      float64
	M4Degree                string
	BlockedIfThetaExt       bool
	NativeThetaExtCertified bool
	TypedCoefficientSource  bool
	NativeCubicStopTheorem  bool
	UntypedM4FitRejected    bool
	Verdict                 string
}

type GeneratingFunctionAudit struct {
	Recorded                bool
	Form                    string
	Truncation              string
	ExteriorDegreeReading   string
	NativeFunctionTheorem   bool
	HigherTermsProvedVanish bool
	Verdict                 string
}

type RelationAudit struct {
	Recorded            bool
	KappaLambdaRelation string
	FWallStatus         string
	KappaLambdaStatus   string
	CHistoryStatus      string
	CHiggsStatus        string
	WouldUpgradeFWall   bool
	KappaLambdaNative   bool
	CHistoryIndependent bool
	Verdict             string
}

type PredictionLevelClassification struct {
	Recorded         bool
	FWall3Level      string
	KappaLambdaLevel string
	CHistoryLevel    string
	CHiggsLevel      string
	FWallLevelC      bool
	Verdict          string
}

type Firewalls struct {
	Enforced                     bool
	BoundaryExteriorDegreeNative bool
	DimensionTwoProof            bool
	TwoPCoefficientNative        bool
	NegativeSignNative           bool
	FWallNative                  bool
	KappaLambdaNative            bool
	CHistoryIndependent          bool
	TreeProxyPoleMass            bool
	YukawaNative                 bool
	Verdict                      string
}

type Audit struct {
	Gate783        Gate783Inheritance
	Ledger         MomentLedger
	BoundaryPair   BoundaryPairAudit
	Exterior       ExteriorDegreeCandidateAudit
	ThetaExt       RequiredMapAudit
	Cubic          CubicCoefficientAudit
	Table          DegreeResponseTable
	Projector      ProjectorFirewallAudit
	M4             M4RejectionAudit
	Generating     GeneratingFunctionAudit
	Relation       RelationAudit
	Prediction     PredictionLevelClassification
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
		return Audit{}, fmt.Errorf("non-finite raw-moment ledger")
	}
	matches := closeRel(fwall, fWall3Snapshot, 1e-14)

	a := Audit{
		Gate783: Gate783Inheritance{
			Inherited:             true,
			ResponseFunction:      "F_wall_3_red=Tr(rho_72 f_3(R_wall)), f_3(x)=x+kappa_e_red x^2-2p x^3",
			BoundarySubBottleneck: "F_wall_3_red was identified as the boundary sub-bottleneck inside kappa_lambda_red.",
			Verdict:               StatusGate783BoundaryRawMomentGeneratingFunctionInherited,
		},
		Ledger: MomentLedger{P: p, S: s, M1: m1, M2: m2, M3: m3, M4: m4, KappaE: kappaERedSnapshot, FWall3: fwall, Matches: matches},
		BoundaryPair: BoundaryPairAudit{
			Typed:                true,
			Carrier:              "B_boundary=span(b_lambda,b_R)",
			Basis:                []string{"b_lambda <-> |lambda(Lambda12)|", "b_R <-> R3-1"},
			ScalarSplit:          "s=lambda(Lambda12)+(R3-1)",
			MidpointStress:       "xi_boundary=0.5(|lambda(Lambda12)|+(R3-1))",
			Dim:                  2,
			ExteriorDimensions:   []int{1, 2, 1, 0},
			IsSpacetimeCarrier:   false,
			IsFlavorCarrier:      false,
			IsK7Carrier:          false,
			TwoDimensionalBridge: true,
			Verdict:              StatusBoundaryPairTwoDimensionalBridgeCarrier,
		},
		Exterior: ExteriorDegreeCandidateAudit{
			Audited: true,
			DegreeAssignments: map[string]string{
				"M1": "degree 0 leading K7 event response in Lambda^0 B_boundary",
				"M2": "degree 1 boundary modulation response in Lambda^1 B_boundary",
				"M3": "degree 2 boundary-pair stress-pull response in Lambda^2 B_boundary",
				"M4": "degree 3 boundary stress response, blocked if Lambda^3 B_boundary=0 through Theta_ext",
			},
			Lambda3Zero:                true,
			CubicStopCandidate:         true,
			DimensionAloneProvesStop:   false,
			NativeThetaExtMapCertified: false,
			Verdict:                    StatusCubicStopBoundaryExteriorDegreeSourceCandidate,
		},
		ThetaExt: RequiredMapAudit{
			Identified: true,
			MapName:    "Theta_ext",
			Domain:     "raw moment layer M_n",
			Codomain:   "Lambda^(n-1) B_boundary response degree",
			Assignments: []string{
				"Theta_ext(M1) in Lambda^0 B_boundary",
				"Theta_ext(M2) in Lambda^1 B_boundary",
				"Theta_ext(M3) in Lambda^2 B_boundary",
				"Theta_ext(M4) in Lambda^3 B_boundary=0",
			},
			Certified:                 false,
			DimensionShortcutRejected: true,
			Verdict:                   StatusNoNativeThetaExtMap,
		},
		Cubic: CubicCoefficientAudit{
			Audited:               true,
			TwoP:                  2 * p,
			MagnitudeSource:       "2p=dim(B_boundary)*p_K7=2*(7/72)=7/36",
			DegreeTwoLayer:        "-2p M3 is the negative degree-2 boundary-pair stress-pull correction candidate.",
			SignSourceCertified:   false,
			StressPullTheorem:     false,
			CompatibleWithDegree2: true,
			Verdict:               Status2PBoundaryPairTimesK7EventWeight,
		},
		Table: DegreeResponseTable{
			Recorded:                       true,
			Degree0:                        "degree 0: M1=p s, leading K7 event response",
			Degree1:                        "degree 1: kappa_e_red M2=kappa_e_red p s^2, flavor-wall modulation of first boundary stress degree",
			Degree2:                        "degree 2: -2p M3=-2p^2 s^3, boundary-pair stress-pull correction",
			Degree3:                        "degree 3: blocked candidate because Lambda^3 B_boundary=0 if Theta_ext is supplied",
			KappaERedNaturallyDegree1:      true,
			FlavorBoundaryModulationNative: false,
			Verdict:                        StatusKappaERedM2DegreeOneFlavorBoundaryModulation,
		},
		Projector: ProjectorFirewallAudit{
			Preserved:               true,
			PowerLaw:                "R_wall(s)^n=s^n P_7 for all n>=1; projector idempotence scalarizes powers but does not stop them.",
			ProjectorStopsExpansion: false,
			NeedsExteriorMap:        true,
			Verdict:                 StatusProjectorIdempotenceDoesNotForceCubicStop,
		},
		M4: M4RejectionAudit{
			Audited:                 true,
			M4:                      m4,
			M4Degree:                "M4 would correspond to degree 3 boundary stress under Theta_ext.",
			BlockedIfThetaExt:       true,
			NativeThetaExtCertified: false,
			TypedCoefficientSource:  false,
			NativeCubicStopTheorem:  false,
			UntypedM4FitRejected:    true,
			Verdict:                 StatusM4BlockedIfThetaExtSupplied,
		},
		Generating: GeneratingFunctionAudit{
			Recorded:                true,
			Form:                    "F_wall(s)=Tr(rho_72 f(R_wall(s)))",
			Truncation:              "f_{<=3}(x)=x+kappa_e_red x^2-2p x^3",
			ExteriorDegreeReading:   "degree-0 leading response + degree-1 flavor-modulated response + degree-2 boundary-pair stress-pull response",
			NativeFunctionTheorem:   false,
			HigherTermsProvedVanish: false,
			Verdict:                 StatusFWall3RedCubicTruncationExteriorDegreeResponseCandidate,
		},
		Relation: RelationAudit{
			Recorded:            true,
			KappaLambdaRelation: "kappa_lambda_red=|lambda|+F_wall_3_red-kappa_e_red",
			FWallStatus:         "Level B+ source-typed boundary response candidate; formula-independent but not native.",
			KappaLambdaStatus:   "Level B formula-independent scalar complement; not native.",
			CHistoryStatus:      "Level B semi-independent History correction; not full independent prediction component.",
			CHiggsStatus:        "still not Level C.",
			WouldUpgradeFWall:   true,
			KappaLambdaNative:   false,
			CHistoryIndependent: false,
			Verdict:             StatusBoundaryExteriorDegreeMapWouldUpgradeFWallSourceStatus,
		},
		Prediction: PredictionLevelClassification{
			Recorded:         true,
			FWall3Level:      "Level B+ source-typed boundary response candidate; not Level C native component.",
			KappaLambdaLevel: "Level B formula-independent scalar complement; not native.",
			CHistoryLevel:    "Level B semi-independent History correction; not native.",
			CHiggsLevel:      "still not Level C.",
			FWallLevelC:      false,
			Verdict:          StatusFWall3RedUpgradedToLevelBPlusSourceCandidate,
		},
		Firewalls: Firewalls{
			Enforced:                     true,
			BoundaryExteriorDegreeNative: false,
			DimensionTwoProof:            false,
			TwoPCoefficientNative:        false,
			NegativeSignNative:           false,
			FWallNative:                  false,
			KappaLambdaNative:            false,
			CHistoryIndependent:          false,
			TreeProxyPoleMass:            false,
			YukawaNative:                 false,
			Verdict:                      StatusFirewallPreservedGate784,
		},
	}
	a.Truth = "Gate 784 identifies Theta_ext: M_n -> Lambda^(n-1)B_boundary as the sharp missing object for a native cubic-stop theorem."
	a.FinalStatement = "Gate 784 does not prove the cubic stop. It improves the status by identifying the sharp missing object: a native map from raw boundary moments to exterior boundary response degree. Under that missing map, the cubic stop would be explained by dim(B_boundary)=2 and Lambda^3 B_boundary=0, with M3 as the boundary-pair stress-pull layer. The next bottleneck is the construction or rejection of Theta_ext: raw moment layer -> exterior boundary degree."
	return a, nil
}

func Statuses() []string {
	return []string{
		StatusGate783BoundaryRawMomentGeneratingFunctionInherited,
		StatusFWall3RedBoundarySubBottleneckInherited,
		StatusBoundaryPairObjectTyped,
		StatusExteriorDegreeCandidateAudited,
		StatusRequiredExteriorDegreeMapIdentified,
		StatusCubicCoefficientCompatibleWithBoundaryPairDegreeTwoSource,
		StatusDegreeByDegreeResponseTableRecorded,
		StatusProjectorIdempotenceFirewallPreserved,
		StatusM4RejectionReaudited,
		StatusGeneratingFunctionCandidateRecorded,
		StatusRelationToKappaLambdaRedRecorded,
		StatusPredictionLevelClassificationRecorded,
		StatusPhysicalFirewallsEnforced,
		StatusBoundaryPairTwoDimensionalBridgeCarrier,
		StatusCubicStopBoundaryExteriorDegreeSourceCandidate,
		Status2PBoundaryPairTimesK7EventWeight,
		StatusKappaERedM2DegreeOneFlavorBoundaryModulation,
		StatusM4BlockedIfThetaExtSupplied,
		StatusFWall3RedCubicTruncationExteriorDegreeResponseCandidate,
		StatusBoundaryExteriorDegreeMapWouldUpgradeFWallSourceStatus,
		StatusFWall3RedUpgradedToLevelBPlusSourceCandidate,
		StatusBoundaryPairNotNativeSpacetimeOrFlavorCarrier,
		StatusNoNativeThetaExtMap,
		StatusDimensionTwoBoundaryPairAloneDoesNotProveCubicStop,
		StatusNoNativeSignTheoremForNegativeCubicStressPull,
		StatusNoNativeBoundaryPairStressPullTheorem,
		StatusNoNativeFlavorBoundaryModulationTheorem,
		StatusProjectorIdempotenceDoesNotForceCubicStop,
		StatusNoNativeTypedM4CoefficientSource,
		StatusNoNativeCubicStopTheorem,
		StatusNoNativeBoundaryResponseGeneratingFunctionTheorem,
		StatusFWall3RedNotLevelCNativeComponent,
		StatusKappaLambdaRedNotYetNativeScalarMatchingTheorem,
		StatusCHistoryNotYetFullIndependentPredictionComponent,
		StatusTreeProxyNotPoleMass,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusFirewallPreservedGate784,
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

func FormatBoundaryPair(b BoundaryPairAudit) string {
	return fmt.Sprintf("%s dim=%d exterior_dims=%v split=%s midpoint=%s", b.Carrier, b.Dim, b.ExteriorDimensions, b.ScalarSplit, b.MidpointStress)
}

func FormatExterior(e ExteriorDegreeCandidateAudit) string {
	keys := []string{"M1", "M2", "M3", "M4"}
	parts := make([]string, 0, len(keys))
	for _, k := range keys {
		parts = append(parts, k+":"+e.DegreeAssignments[k])
	}
	return fmt.Sprintf("lambda3_zero=%v candidate=%v dimension_proves=%v map_certified=%v; %s", e.Lambda3Zero, e.CubicStopCandidate, e.DimensionAloneProvesStop, e.NativeThetaExtMapCertified, strings.Join(parts, "; "))
}

func FormatTheta(t RequiredMapAudit) string {
	return fmt.Sprintf("%s: %s -> %s; certified=%v shortcut_rejected=%v", t.MapName, t.Domain, t.Codomain, t.Certified, t.DimensionShortcutRejected)
}

func FormatCubic(c CubicCoefficientAudit) string {
	return fmt.Sprintf("2p=%.17g %s degree2=%s sign_native=%v stress_theorem=%v", c.TwoP, c.MagnitudeSource, c.DegreeTwoLayer, c.SignSourceCertified, c.StressPullTheorem)
}

func FormatRelation(r RelationAudit) string {
	return fmt.Sprintf("F_wall=%s; kappa_lambda=%s; C_History=%s; C_Higgs=%s", r.FWallStatus, r.KappaLambdaStatus, r.CHistoryStatus, r.CHiggsStatus)
}
