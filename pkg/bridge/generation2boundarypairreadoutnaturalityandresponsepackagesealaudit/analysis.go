// Package generation2boundarypairreadoutnaturalityandresponsepackagesealaudit implements
// Gate 786: Boundary Pair Readout Naturality and Response-Package Seal Audit.
//
// Gate 785 showed that F_wall_3_red can be represented by a conditional
// exterior response package only after supplying Theta_ext, chi_ext, a
// degree-one readout, and a degree-two orientation/sign. Gate 786 asks whether
// the already-labelled two-boundary pair sources those missing structures
// natively. It does not derive scalar runtime lambda, Higgs pole mass, Yukawa
// operators, PMNS, CKM, flavor hierarchy, G_F, VEV, or a native
// HistoryLoopUnit theorem.
package generation2boundarypairreadoutnaturalityandresponsepackagesealaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE786-BOUNDARY-PAIR-READOUT-NATURALITY-RESPONSE-PACKAGE-SEAL-AUDIT"

	StatusGate785ThetaExtResponsePackageInherited    = "PASS_GATE785_THETA_EXT_RESPONSE_PACKAGE_INHERITED"
	StatusBoundaryPairDataInventoryRecorded          = "PASS_BOUNDARY_PAIR_DATA_INVENTORY_RECORDED"
	StatusLabelledBasisVersusNativeBasisAudited      = "PASS_LABELLED_BASIS_VERSUS_NATIVE_BASIS_AUDITED"
	StatusDegreeOneAxisReadoutAudited                = "PASS_DEGREE_ONE_AXIS_READOUT_AUDITED"
	StatusDegreeTwoOrientationAndSignAudited         = "PASS_DEGREE_TWO_ORIENTATION_AND_SIGN_AUDITED"
	StatusScalarReadoutChiExtAudited                 = "PASS_SCALAR_READOUT_CHI_EXT_AUDITED"
	StatusBoundarySymmetryNaturalityAudited          = "PASS_BOUNDARY_SYMMETRY_NATURALITY_AUDITED"
	StatusBoundaryExteriorResponsePackageSealDefined = "PASS_BOUNDARY_EXTERIOR_RESPONSE_PACKAGE_SEAL_DEFINED"
	StatusImpactOnFWallStatusAudited                 = "PASS_IMPACT_ON_F_WALL_STATUS_AUDITED"
	StatusStatusPropagationRecorded                  = "PASS_STATUS_PROPAGATION_RECORDED"
	StatusPhysicalFirewallsEnforced                  = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusBoundaryPairHasLabelledBridgeAxesAndScalarReadouts       = "CONDITIONAL_SUPPORT_BOUNDARY_PAIR_HAS_LABELLED_BRIDGE_AXES_AND_SCALAR_READOUTS"
	StatusLabelledBoundaryAxesDefineBridgeBasis                    = "CONDITIONAL_SUPPORT_LABELLED_BOUNDARY_AXES_DEFINE_A_BRIDGE_BASIS"
	StatusSignedSplitAxisDegreeOneBoundaryAxisCandidate            = "CONDITIONAL_SUPPORT_SIGNED_SPLIT_AXIS_IS_DEGREE_ONE_BOUNDARY_AXIS_CANDIDATE"
	StatusMidpointStressAxisBoundaryAxisCandidate                  = "CONDITIONAL_SUPPORT_MIDPOINT_STRESS_AXIS_IS_BOUNDARY_AXIS_CANDIDATE"
	StatusKappaERedFlavorBoundaryReadoutCoefficient                = "CONDITIONAL_SUPPORT_KAPPA_E_RED_IS_FLAVOR_BOUNDARY_READOUT_COEFFICIENT"
	StatusVolumeFormExistsAfterOrderedBoundaryBasis                = "CONDITIONAL_SUPPORT_VOLUME_FORM_EXISTS_AFTER_ORDERED_BOUNDARY_BASIS"
	Status2PMagnitudeBoundaryPairTimesK7EventSourceType            = "CONDITIONAL_SUPPORT_2P_MAGNITUDE_HAS_BOUNDARY_PAIR_TIMES_K7_EVENT_SOURCE_TYPE"
	StatusLabelledBoundaryPairReducesSymmetryForConditionalPackage = "CONDITIONAL_SUPPORT_LABELLED_BOUNDARY_PAIR_REDUCES_SYMMETRY_ENOUGH_FOR_CONDITIONAL_PACKAGE"
	StatusSealMinimalForExteriorResponseRepresentation             = "CONDITIONAL_SUPPORT_THIS_SEAL_IS_MINIMAL_FOR_EXTERIOR_RESPONSE_REPRESENTATION"
	StatusFWall3RedSealedExteriorResponseRepresentable             = "CONDITIONAL_SUPPORT_F_WALL_3_RED_IS_SEALED_EXTERIOR_RESPONSE_REPRESENTABLE"

	StatusBoundaryPairDataDoNotAutomaticallyDefineResponsePackage = "FAILED_ROUTE_BOUNDARY_PAIR_DATA_DO_NOT_AUTOMATICALLY_DEFINE_RESPONSE_PACKAGE"
	StatusLabelledBridgeBasisNotNativeInvariantBasis              = "FAILED_ROUTE_LABELLED_BRIDGE_BASIS_NOT_YET_NATIVE_INVARIANT_BASIS"
	StatusSplitAxisDoesNotSourceKappaECoefficient                 = "FAILED_ROUTE_SPLIT_AXIS_DOES_NOT_SOURCE_KAPPA_E_RED_COEFFICIENT"
	StatusMidpointAxisDoesNotSourceKappaECoefficient              = "FAILED_ROUTE_MIDPOINT_AXIS_DOES_NOT_SOURCE_KAPPA_E_RED_COEFFICIENT"
	StatusKappaENotSourcedByBoundaryPairAlone                     = "FAILED_ROUTE_KAPPA_E_RED_NOT_SOURCED_BY_BOUNDARY_PAIR_ALONE"
	StatusNoNativeDegreeOneReadoutTheorem                         = "FAILED_ROUTE_NO_NATIVE_DEGREE_ONE_READOUT_THEOREM"
	StatusNoNativeOrderedBoundaryOrientationTheorem               = "FAILED_ROUTE_NO_NATIVE_ORDERED_BOUNDARY_ORIENTATION_THEOREM"
	StatusNoNativeNegativeStressPullSignTheorem                   = "FAILED_ROUTE_NO_NATIVE_NEGATIVE_STRESS_PULL_SIGN_THEOREM"
	StatusDegreeOneReadoutRequiresFlavorWallInput                 = "FAILED_ROUTE_DEGREE_ONE_READOUT_REQUIRES_FLAVOR_WALL_INPUT"
	StatusDegreeTwoReadoutRequiresK7WeightPlusSignConvention      = "FAILED_ROUTE_DEGREE_TWO_READOUT_REQUIRES_K7_WEIGHT_PLUS_SIGN_CONVENTION"
	StatusChiExtNotNativeFromBoundaryPairAlone                    = "FAILED_ROUTE_CHI_EXT_NOT_NATIVE_FROM_BOUNDARY_PAIR_ALONE"
	StatusAbstractBBoundaryNoCanonicalResponsePackage             = "FAILED_ROUTE_ABSTRACT_B_BOUNDARY_HAS_NO_CANONICAL_RESPONSE_PACKAGE"
	StatusLabelledPackageRemainsBridgeSealedNotNative             = "FAILED_ROUTE_LABELLED_PACKAGE_REMAINS_BRIDGE_SEALED_NOT_NATIVE"
	StatusResponsePackageSealNotNativeBoundaryResponseTheorem     = "FAILED_ROUTE_RESPONSE_PACKAGE_SEAL_NOT_NATIVE_BOUNDARY_RESPONSE_THEOREM"
	StatusFWall3RedNotNativeBoundaryGeneratingFunction            = "FAILED_ROUTE_F_WALL_3_RED_NOT_NATIVE_BOUNDARY_GENERATING_FUNCTION"
	StatusKappaLambdaRedNotNativeScalarMatchingTheorem            = "FAILED_ROUTE_KAPPA_LAMBDA_RED_NOT_NATIVE_SCALAR_MATCHING_THEOREM"
	StatusCHistoryNotFullIndependentPredictionComponent           = "FAILED_ROUTE_C_HISTORY_NOT_FULL_INDEPENDENT_PREDICTION_COMPONENT"
	StatusCHiggsNotLevelCPrediction                               = "FAILED_ROUTE_C_HIGGS_NOT_LEVEL_C_PREDICTION"
	StatusTreeProxyNotPoleMass                                    = "FAILED_ROUTE_TREE_PROXY_NOT_POLE_MASS"
	StatusNoYukawaOperatorOrEigenvalueTheorem                     = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusFirewallPreservedGate786                                = "FIREWALL_PRESERVED_GATE786_BOUNDARY_PAIR_READOUT_NATURALITY_BOUNDARY"
)

const (
	pK7Snapshot        = 7.0 / 72.0
	sSplitSnapshot     = 0.0012924448188162962
	xiBoundarySnapshot = 0.0503471644870914
	kappaERedSnapshot  = 0.005503554218475772
	fWall3Snapshot     = 0.00012565521035653708
)

type Gate785Inheritance struct {
	Inherited          bool
	ConditionalPackage string
	CurrentBottleneck  string
	PriorNative        bool
	Verdict            string
}

type BoundaryPairInventory struct {
	Recorded           bool
	BoundaryCarrier    string
	BoundaryAxes       []string
	BoundaryReadouts   []string
	K7EventWeight      float64
	FlavorScalar       string
	ExteriorPackage    []string
	RolesSeparated     bool
	AutoDefinesPackage bool
	Verdict            string
}

type BasisAudit struct {
	Audited              bool
	LabelledBridgeBasis  bool
	NativeInvariantBasis bool
	SourceTypes          []string
	BridgeBasisVerdict   string
	NativeVerdict        string
}

type DegreeOneAxisReadoutAudit struct {
	Audited                  bool
	SplitAxisCandidate       bool
	SplitAxisSource          string
	SplitAxisSourcesKappaE   bool
	MidpointAxisCandidate    bool
	MidpointAxisSource       string
	MidpointSourcesKappaE    bool
	FlavorReadoutCoefficient bool
	KappaESourcedByBoundary  bool
	NativeReadoutTheorem     bool
	Verdict                  string
}

type DegreeTwoOrientationSignAudit struct {
	Audited                  bool
	Lambda2Exists            bool
	VolumeForm               string
	RequiresOrderedBasis     bool
	TwoP                     float64
	MagnitudeSource          string
	NegativeSign             string
	NativeOrderedOrientation bool
	NativeNegativeStressPull bool
	Verdict                  string
}

type ChiExtAudit struct {
	Audited                     bool
	DegreeZeroCanonical         bool
	Chi0                        float64
	DegreeOneRequiresFlavor     bool
	Chi1                        float64
	DegreeTwoRequiresK7AndSign  bool
	Chi2                        float64
	NativeFromBoundaryPairAlone bool
	Verdict                     string
}

type BoundarySymmetryNaturalityAudit struct {
	Audited                     bool
	AbstractGL2Freedom          bool
	CanonicalBetaUnderGL2       bool
	CanonicalOmegaSignUnderGL2  bool
	CanonicalChiUnderGL2        bool
	LabelledPairReducesSymmetry bool
	LabelledPackageBridgeSealed bool
	Verdict                     string
}

type ResponsePackageSeal struct {
	Defined            bool
	Name               string
	Components         []string
	Supplies           []string
	Minimal            bool
	Native             bool
	ReconstructedFWall float64
	MatchesFWall       bool
	Verdict            string
}

type ImpactOnFWall struct {
	Audited                  bool
	RepresentableBySeal      bool
	NativeGeneratingFunction bool
	Status                   string
	Verdict                  string
}

type StatusPropagation struct {
	Recorded    bool
	FWall3      string
	KappaLambda string
	CHistory    string
	CHiggs      string
	Verdict     string
}

type Firewalls struct {
	Enforced                 bool
	LabelledAxesNative       bool
	SplitAxisKappaETheorem   bool
	MidpointKappaETheorem    bool
	OmegaSignTheorem         bool
	TwoPMagnitudeFullTheorem bool
	SealNativeGeneratingFunc bool
	FWallNative              bool
	KappaLambdaNative        bool
	CHistoryIndependent      bool
	TreeProxyPoleMass        bool
	YukawaNative             bool
	Verdict                  string
}

type Audit struct {
	Gate785        Gate785Inheritance
	Inventory      BoundaryPairInventory
	Basis          BasisAudit
	DegreeOne      DegreeOneAxisReadoutAudit
	DegreeTwo      DegreeTwoOrientationSignAudit
	ChiExt         ChiExtAudit
	Naturality     BoundarySymmetryNaturalityAudit
	Seal           ResponsePackageSeal
	Impact         ImpactOnFWall
	Propagation    StatusPropagation
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
	fwall := m1 + kappaERedSnapshot*m2 - 2*p*m3
	if math.IsNaN(fwall) || math.IsInf(fwall, 0) {
		return Audit{}, fmt.Errorf("non-finite Gate786 F_wall reconstruction")
	}
	matches := closeRel(fwall, fWall3Snapshot, 1e-14)

	a := Audit{
		Gate785: Gate785Inheritance{
			Inherited:          true,
			ConditionalPackage: "Theta_ext + chi_ext + beta_B degree-one readout + omega_B orientation/sign",
			CurrentBottleneck:  "native sourcing of the exterior response package rather than algebraic representation of F_wall_3_red",
			PriorNative:        false,
			Verdict:            StatusGate785ThetaExtResponsePackageInherited,
		},
		Inventory: BoundaryPairInventory{
			Recorded:        true,
			BoundaryCarrier: "B_boundary=span(b_lambda,b_R)",
			BoundaryAxes: []string{
				"b_lambda: scalar wall-depth axis associated to |lambda(Lambda12)|",
				"b_R: gauge/boundary stress axis associated to R3-1",
			},
			BoundaryReadouts: []string{
				"s=lambda(Lambda12)+(R3-1)=(R3-1)-|lambda| when lambda(Lambda12)<0",
				fmt.Sprintf("xi_boundary=0.5(|lambda|+(R3-1))=%.16g", xiBoundarySnapshot),
			},
			K7EventWeight:      p,
			FlavorScalar:       "kappa_e_red is flavor-wall reduced scalar, not a boundary-pair vector by default",
			ExteriorPackage:    []string{"Theta_ext", "chi_ext", "beta_B or degree-one readout", "ordered boundary orientation", "negative stress-pull sign convention"},
			RolesSeparated:     true,
			AutoDefinesPackage: false,
			Verdict:            StatusBoundaryPairHasLabelledBridgeAxesAndScalarReadouts,
		},
		Basis: BasisAudit{
			Audited:              true,
			LabelledBridgeBasis:  true,
			NativeInvariantBasis: false,
			SourceTypes: []string{
				"|lambda(Lambda12)|: scalar wall depth",
				"R3-1: gauge/boundary stress",
			},
			BridgeBasisVerdict: StatusLabelledBoundaryAxesDefineBridgeBasis,
			NativeVerdict:      StatusLabelledBridgeBasisNotNativeInvariantBasis,
		},
		DegreeOne: DegreeOneAxisReadoutAudit{
			Audited:                  true,
			SplitAxisCandidate:       true,
			SplitAxisSource:          "beta_s ~ b_R-b_lambda from signed split s=(R3-1)-|lambda|",
			SplitAxisSourcesKappaE:   false,
			MidpointAxisCandidate:    true,
			MidpointAxisSource:       "beta_xi ~ b_lambda+b_R from midpoint stress xi_boundary",
			MidpointSourcesKappaE:    false,
			FlavorReadoutCoefficient: true,
			KappaESourcedByBoundary:  false,
			NativeReadoutTheorem:     false,
			Verdict:                  StatusKappaERedFlavorBoundaryReadoutCoefficient,
		},
		DegreeTwo: DegreeTwoOrientationSignAudit{
			Audited:                  true,
			Lambda2Exists:            true,
			VolumeForm:               "omega_B=b_lambda∧b_R once an ordered boundary basis is supplied",
			RequiresOrderedBasis:     true,
			TwoP:                     2 * p,
			MagnitudeSource:          "2p=dim(B_boundary)*p_K7=2*(7/72)=7/36",
			NegativeSign:             "negative sign requires stress-pull orientation/readout convention chi_ext(omega_B)=-2p",
			NativeOrderedOrientation: false,
			NativeNegativeStressPull: false,
			Verdict:                  StatusVolumeFormExistsAfterOrderedBoundaryBasis,
		},
		ChiExt: ChiExtAudit{
			Audited:                     true,
			DegreeZeroCanonical:         true,
			Chi0:                        1,
			DegreeOneRequiresFlavor:     true,
			Chi1:                        kappaERedSnapshot,
			DegreeTwoRequiresK7AndSign:  true,
			Chi2:                        -2 * p,
			NativeFromBoundaryPairAlone: false,
			Verdict:                     StatusChiExtNotNativeFromBoundaryPairAlone,
		},
		Naturality: BoundarySymmetryNaturalityAudit{
			Audited:                     true,
			AbstractGL2Freedom:          true,
			CanonicalBetaUnderGL2:       false,
			CanonicalOmegaSignUnderGL2:  false,
			CanonicalChiUnderGL2:        false,
			LabelledPairReducesSymmetry: true,
			LabelledPackageBridgeSealed: true,
			Verdict:                     StatusAbstractBBoundaryNoCanonicalResponsePackage,
		},
		Seal: ResponsePackageSeal{
			Defined: true,
			Name:    "BoundaryExteriorResponsePackageSeal",
			Components: []string{
				"Theta_ext",
				"chi_ext",
				"beta_B or degree-one readout",
				"ordered boundary orientation",
				"negative stress-pull sign convention",
			},
			Supplies: []string{
				"Theta_ext(M_n) in Lambda^(n-1)B_boundary",
				"chi_ext(1_B)=1",
				"chi_ext(beta_B)=kappa_e_red",
				"chi_ext(omega_B)=-2p",
				"Theta_ext(M_n>=4)=0",
			},
			Minimal:            true,
			Native:             false,
			ReconstructedFWall: fwall,
			MatchesFWall:       matches,
			Verdict:            StatusSealMinimalForExteriorResponseRepresentation,
		},
		Impact: ImpactOnFWall{
			Audited:                  true,
			RepresentableBySeal:      true,
			NativeGeneratingFunction: false,
			Status:                   "F_wall_3_red is sealed exterior-response representable, but not a native boundary generating function.",
			Verdict:                  StatusFWall3RedSealedExteriorResponseRepresentable,
		},
		Propagation: StatusPropagation{
			Recorded:    true,
			FWall3:      "Level B+ sealed exterior response candidate",
			KappaLambda: "Level B formula-independent scalar matching complement; not native",
			CHistory:    "Level B semi-independent History correction; not full independent prediction component",
			CHiggs:      "still not Level C prediction",
			Verdict:     StatusFWall3RedNotNativeBoundaryGeneratingFunction,
		},
		Firewalls: Firewalls{
			Enforced:                 true,
			LabelledAxesNative:       false,
			SplitAxisKappaETheorem:   false,
			MidpointKappaETheorem:    false,
			OmegaSignTheorem:         false,
			TwoPMagnitudeFullTheorem: false,
			SealNativeGeneratingFunc: false,
			FWallNative:              false,
			KappaLambdaNative:        false,
			CHistoryIndependent:      false,
			TreeProxyPoleMass:        false,
			YukawaNative:             false,
			Verdict:                  StatusFirewallPreservedGate786,
		},
	}
	a.Truth = "Gate 786 finds that the labelled boundary pair can support a conditional exterior response representation, but it cannot source the response package natively."
	a.FinalStatement = "Gate 786 finds that the existing boundary pair can conditionally support an exterior response representation only after accepting labelled bridge axes, a degree-one flavor readout, an ordered boundary orientation, and a negative stress-pull sign convention. It does not source the package natively. The minimal missing object is BoundaryExteriorResponsePackageSeal: Theta_ext + chi_ext + degree-one readout + ordered boundary orientation + negative stress-pull sign. The next bottleneck is no longer algebraic representation of F_wall_3_red. It is native sourcing of the response-package seal, especially the degree-one flavor readout and the negative degree-two stress-pull sign."
	return a, nil
}

func Statuses() []string {
	return []string{
		StatusGate785ThetaExtResponsePackageInherited,
		StatusBoundaryPairDataInventoryRecorded,
		StatusLabelledBasisVersusNativeBasisAudited,
		StatusDegreeOneAxisReadoutAudited,
		StatusDegreeTwoOrientationAndSignAudited,
		StatusScalarReadoutChiExtAudited,
		StatusBoundarySymmetryNaturalityAudited,
		StatusBoundaryExteriorResponsePackageSealDefined,
		StatusImpactOnFWallStatusAudited,
		StatusStatusPropagationRecorded,
		StatusPhysicalFirewallsEnforced,
		StatusBoundaryPairHasLabelledBridgeAxesAndScalarReadouts,
		StatusLabelledBoundaryAxesDefineBridgeBasis,
		StatusSignedSplitAxisDegreeOneBoundaryAxisCandidate,
		StatusMidpointStressAxisBoundaryAxisCandidate,
		StatusKappaERedFlavorBoundaryReadoutCoefficient,
		StatusVolumeFormExistsAfterOrderedBoundaryBasis,
		Status2PMagnitudeBoundaryPairTimesK7EventSourceType,
		StatusLabelledBoundaryPairReducesSymmetryForConditionalPackage,
		StatusSealMinimalForExteriorResponseRepresentation,
		StatusFWall3RedSealedExteriorResponseRepresentable,
		StatusBoundaryPairDataDoNotAutomaticallyDefineResponsePackage,
		StatusLabelledBridgeBasisNotNativeInvariantBasis,
		StatusSplitAxisDoesNotSourceKappaECoefficient,
		StatusMidpointAxisDoesNotSourceKappaECoefficient,
		StatusKappaENotSourcedByBoundaryPairAlone,
		StatusNoNativeDegreeOneReadoutTheorem,
		StatusNoNativeOrderedBoundaryOrientationTheorem,
		StatusNoNativeNegativeStressPullSignTheorem,
		StatusDegreeOneReadoutRequiresFlavorWallInput,
		StatusDegreeTwoReadoutRequiresK7WeightPlusSignConvention,
		StatusChiExtNotNativeFromBoundaryPairAlone,
		StatusAbstractBBoundaryNoCanonicalResponsePackage,
		StatusLabelledPackageRemainsBridgeSealedNotNative,
		StatusResponsePackageSealNotNativeBoundaryResponseTheorem,
		StatusFWall3RedNotNativeBoundaryGeneratingFunction,
		StatusKappaLambdaRedNotNativeScalarMatchingTheorem,
		StatusCHistoryNotFullIndependentPredictionComponent,
		StatusCHiggsNotLevelCPrediction,
		StatusTreeProxyNotPoleMass,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusFirewallPreservedGate786,
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

func FormatInventory(i BoundaryPairInventory) string {
	return fmt.Sprintf("%s axes=%s readouts=%s p=%.17g roles_separated=%v auto_package=%v", i.BoundaryCarrier, strings.Join(i.BoundaryAxes, "; "), strings.Join(i.BoundaryReadouts, "; "), i.K7EventWeight, i.RolesSeparated, i.AutoDefinesPackage)
}

func FormatDegreeOne(d DegreeOneAxisReadoutAudit) string {
	return fmt.Sprintf("split=%v sources_kappa=%v midpoint=%v sources_kappa=%v flavor_readout=%v boundary_sourced=%v native=%v", d.SplitAxisCandidate, d.SplitAxisSourcesKappaE, d.MidpointAxisCandidate, d.MidpointSourcesKappaE, d.FlavorReadoutCoefficient, d.KappaESourcedByBoundary, d.NativeReadoutTheorem)
}

func FormatSeal(s ResponsePackageSeal) string {
	return fmt.Sprintf("%s components=[%s] supplies=[%s] minimal=%v native=%v F=%.17g matches=%v", s.Name, strings.Join(s.Components, "; "), strings.Join(s.Supplies, "; "), s.Minimal, s.Native, s.ReconstructedFWall, s.MatchesFWall)
}

func FormatPropagation(p StatusPropagation) string {
	return fmt.Sprintf("F_wall=%s; kappa_lambda=%s; C_History=%s; C_Higgs=%s", p.FWall3, p.KappaLambda, p.CHistory, p.CHiggs)
}
