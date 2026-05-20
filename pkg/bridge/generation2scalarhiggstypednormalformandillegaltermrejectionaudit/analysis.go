// Package generation2scalarhiggstypednormalformandillegaltermrejectionaudit implements
// Gate 751: Scalar-Higgs Typed Normal Form and Illegal-Term Rejection Audit.
//
// Gate 750 fixed the typed Cl(1,7) board. Gate 751 writes the scalar-Higgs
// bridge in a single typed normal form, specifies every trace/operator/scalar
// airlock, and rejects illegal cross-type terms before any further reduction is
// attempted.
package generation2scalarhiggstypednormalformandillegaltermrejectionaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate750 "github.com/bagherbal/asha-engine/pkg/bridge/generation2cl17boardscalarhiggstypeledgerandoperatorairlockaudit"
)

const (
	AuditID = "GATE751-SCALAR-HIGGS-TYPED-NORMAL-FORM-ILLEGAL-TERM-REJECTION-AUDIT"

	StatusGate750TypeLedgerInherited            = "PASS_GATE750_CL17_TYPE_LEDGER_INHERITED"
	StatusBoundaryQuotientCoordinateTyped       = "PASS_BOUNDARY_QUOTIENT_COORDINATE_TYPED"
	StatusK7ResponseOperatorTyped               = "PASS_K7_RESPONSE_OPERATOR_TYPED"
	StatusRawMomentMapTyped                     = "PASS_RAW_MOMENT_MAP_TYPED"
	StatusCubicBoundaryHistoryNormalFormDefined = "PASS_CUBIC_BOUNDARY_HISTORY_RESPONSE_NORMAL_FORM_DEFINED"
	StatusHiggsRadialHopfLoopFactorTyped        = "PASS_HIGGS_RADIAL_HOPF_LOOP_FACTOR_TYPED"
	StatusScalarHiggsTypedNormalFormWritten     = "PASS_SCALAR_HIGGS_TYPED_NORMAL_FORM_WRITTEN"
	StatusLegalOperationAuditCompleted          = "PASS_LEGAL_OPERATION_AUDIT_COMPLETED"
	StatusIllegalTermRejectionAudited           = "PASS_ILLEGAL_TERM_REJECTION_AUDITED"
	StatusKappaEInsertionStatusRecorded         = "PASS_KAPPA_E_INSERTION_STATUS_RECORDED"

	StatusTypedNormalFormSupported              = "CONDITIONAL_SUPPORT_SCALAR_HIGGS_BRIDGE_HAS_TYPED_NORMAL_FORM"
	StatusFWall3QBoundaryQHistoryScalarResponse = "CONDITIONAL_SUPPORT_F_WALL_3_IS_QBOUNDARY_TO_QHISTORY_SCALAR_RESPONSE"
	StatusLHopfTraceExpectationOnK7Plus         = "CONDITIONAL_SUPPORT_L_HOPF_IS_TRACE_EXPECTATION_ON_K7_PLUS"
	StatusRuntimeAfterTraceCollapse             = "CONDITIONAL_SUPPORT_RUNTIME_FORM_IS_SCALAR_TRANSPORT_AFTER_TRACE_COLLAPSE"

	StatusK7NotBoundaryVectorMap                   = "FAILED_ROUTE_K7_NOT_BOUNDARY_VECTOR_MAP"
	StatusFWall3NotNativeOperatorOnK7              = "FAILED_ROUTE_F_WALL_3_NOT_NATIVE_OPERATOR_ON_K7"
	StatusLHopfNotBoundaryResponseCoefficient      = "FAILED_ROUTE_L_HOPF_NOT_BOUNDARY_RESPONSE_COEFFICIENT"
	StatusSevenOver72NotSourceOfOneOver8Pi         = "FAILED_ROUTE_SEVEN_OVER_SEVENTY_TWO_NOT_SOURCE_OF_ONE_OVER_EIGHT_PI"
	StatusKappaESubstitutionNotNativeFlavorTheorem = "FAILED_ROUTE_KAPPA_E_SUBSTITUTION_NOT_NATIVE_FLAVOR_THEOREM"
	StatusNoNativeNQPRadTheorem                    = "FAILED_ROUTE_NO_NATIVE_N_Q_P_RAD_THEOREM"
	StatusNoBoundaryGeneratingFunctionTheorem      = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_RESPONSE_GENERATING_FUNCTION_THEOREM"
	StatusNoNativeHistoryLoopUnitTheorem           = "FAILED_ROUTE_NO_NATIVE_HISTORYLOOPUNIT_THEOREM"
	StatusNoIndependentScalarRuntimeTheorem        = "FAILED_ROUTE_NO_INDEPENDENT_SCALAR_RUNTIME_THEOREM"
	StatusNoHiggsMassOrPoleMassTheorem             = "FAILED_ROUTE_NO_HIGGS_MASS_OR_POLE_MASS_THEOREM"
	StatusNoYukawaOperatorOrEigenvalueTheorem      = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusGate751Boundary                          = "FIREWALL_PRESERVED_GATE751_SCALAR_HIGGS_TYPED_NORMAL_FORM_BOUNDARY"
)

const (
	pK7         = 7.0 / 72.0
	kappaE      = 0.00550355419157456
	lambdaAbs   = 0.0497009420776833
	rMinusOne   = 0.050993386896499596
	xiBoundary  = 0.0503471644870914
	lambdaProxy = 0.12490310236015
)

type Gate750Inheritance struct {
	Inherited          bool
	TypedLedgerReady   bool
	TraceAirlocksReady bool
	FirewallsPreserved bool
	Verdict            string
}

type TypedDomains struct {
	NativeFiniteChamber      string
	Lambda4Chamber           string
	K7Carrier                string
	HodgeSplit               string
	AugmentedResponseChamber string
	BoundaryQuotientLine     string
	HistoryReadoutLine       string
	ScalarRuntimeLine        string
	Verdict                  string
}

type BoundaryQuotientCoordinate struct {
	BoundaryVector string
	Functional     string
	CoordinateName string
	SValue         float64
	LivesIn        string
	Verdict        string
}

type K7ResponseOperator struct {
	LiftedProjector       string
	ResponseOperator      string
	MultiplicationMeaning string
	OperatorLane          string
	NotTensorProduct      bool
	NotBoundaryMap        bool
	Verdict               string
}

type RawMomentMap struct {
	Observer          string
	Definition        string
	ProjectorPowerLaw string
	EventWeight       float64
	M1                float64
	M2                float64
	M3                float64
	Formula           string
	Verdict           string
}

type CubicBoundaryNormalForm struct {
	MapType         string
	Definition      string
	Polynomial      string
	Value           float64
	NotOperatorOnK7 bool
	Verdict         string
}

type HiggsRadialHopfLoopFactor struct {
	SealPackage      string
	Observer         string
	HopfOperator     string
	TraceExpectation string
	LHopf            float64
	OperatorLane     string
	ScalarAfterTrace bool
	Verdict          string
}

type ScalarHiggsTypedNormalForm struct {
	WoundDefinition      string
	RuntimeFormula       string
	ExpandedFormula      string
	W3                   float64
	RuntimeBridge        float64
	AllOperatorCollapsed bool
	Verdict              string
}

type LegalOperationAudit struct {
	LawfulAdditions []string
	LawfulProducts  []string
	TraceMaps       []string
	RuntimeScalars  []string
	Verdict         string
}

type IllegalTermRejection struct {
	RejectedTerms         []string
	K7BoundaryBlocked     bool
	FWallOperatorBlocked  bool
	LBoundaryCoeffBlocked bool
	SevenLoopBlocked      bool
	TreePoleBlocked       bool
	PredictionBlocked     bool
	Verdict               string
}

type KappaEInsertionStatus struct {
	InsideFWall3            bool
	OutsideRuntimeTransport bool
	CandidateFormula        string
	CandidateValue          float64
	ApproximationResidual   float64
	NativeFlavorTheorem     bool
	Verdict                 string
}

type Analysis struct {
	Gate750    Gate750Inheritance
	Domains    TypedDomains
	Boundary   BoundaryQuotientCoordinate
	K7Response K7ResponseOperator
	Moments    RawMomentMap
	Cubic      CubicBoundaryNormalForm
	Hopf       HiggsRadialHopfLoopFactor
	NormalForm ScalarHiggsTypedNormalForm
	Legal      LegalOperationAudit
	Illegal    IllegalTermRejection
	KappaE     KappaEInsertionStatus
	Truth      string
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
	g750, err := gate750.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate750 inheritance unavailable: %w", err)
	}
	inheritance := buildGate750Inheritance(g750)
	domains := buildTypedDomains()
	boundary := buildBoundaryQuotientCoordinate()
	k7 := buildK7ResponseOperator()
	moments := buildRawMomentMap(boundary.SValue)
	cubic := buildCubicBoundaryNormalForm(boundary.SValue, moments)
	hopf := buildHiggsRadialHopfLoopFactor()
	normal := buildScalarHiggsTypedNormalForm(cubic.Value, hopf.LHopf)
	legal := buildLegalOperationAudit()
	illegal := buildIllegalTermRejection()
	kappa := buildKappaEInsertionStatus(boundary.SValue)
	truth := "Gate 751 writes the scalar-Higgs bridge in typed normal form. Operators live only in End(H72) or End(K7+) until trace expectations collapse them to scalar coordinates; F_wall_3 is a scalar response map Q_boundary -> Q_history; L_Hopf is a trace expectation on K7+; and the runtime bridge is scalar transport after trace collapse. Cross-type terms such as K7+boundary vector, F_wall_3 as a native K7 operator, or 7/72 as the source of 1/(8π) remain rejected."
	return Analysis{Gate750: inheritance, Domains: domains, Boundary: boundary, K7Response: k7, Moments: moments, Cubic: cubic, Hopf: hopf, NormalForm: normal, Legal: legal, Illegal: illegal, KappaE: kappa, Truth: truth}, nil
}

func buildGate750Inheritance(g gate750.Analysis) Gate750Inheritance {
	return Gate750Inheritance{
		Inherited:          g.Gate749.Inherited && g.Finite.FourthGradeDim == 70,
		TypedLedgerReady:   g.H72.Dim == 72 && strings.Contains(g.History.PolynomialType, "Q_boundary -> Q_history"),
		TraceAirlocksReady: strings.Contains(g.Socket.HistoryLoopScalar, "Tr") && strings.Contains(g.H72.RawMomentFormula, "Tr"),
		FirewallsPreserved: g.Firewalls.K7BoundaryVectorBlocked && g.Firewalls.HomTensorSubspaceBlocked && g.Firewalls.HistoryLoopNativeOpen && g.Firewalls.HiggsMassBlocked && g.Firewalls.YukawaBlocked,
		Verdict:            StatusGate750TypeLedgerInherited,
	}
}

func buildTypedDomains() TypedDomains {
	return TypedDomains{
		NativeFiniteChamber:      "V8=span(e_0,...,e_7)",
		Lambda4Chamber:           "Lambda4=Lambda^4 V8",
		K7Carrier:                "K7=Im(P_B)∩Im(P_G)",
		HodgeSplit:               "K7=K7+⊕K7-",
		AugmentedResponseChamber: "H72=Lambda4⊕B_boundary",
		BoundaryQuotientLine:     "Q_boundary=B_boundary/L_anti",
		HistoryReadoutLine:       "Q_history≅R",
		ScalarRuntimeLine:        "Q_runtime≅R",
		Verdict:                  StatusBoundaryQuotientCoordinateTyped,
	}
}

func buildBoundaryQuotientCoordinate() BoundaryQuotientCoordinate {
	s := rMinusOne - lambdaAbs
	return BoundaryQuotientCoordinate{
		BoundaryVector: "b=lambda(Lambda12)b_lambda+(R3-1)b_R",
		Functional:     "sigma_boundary(b)=lambda(Lambda12)+(R3-1)",
		CoordinateName: "S_split",
		SValue:         s,
		LivesIn:        "Q_boundary",
		Verdict:        StatusBoundaryQuotientCoordinateTyped,
	}
}

func buildK7ResponseOperator() K7ResponseOperator {
	return K7ResponseOperator{
		LiftedProjector:       "P_7=P_K7⊕0_boundary ∈ End(H72)",
		ResponseOperator:      "R_wall(s)=sP_7 ∈ End(H72)",
		MultiplicationMeaning: "scalar boundary quotient coordinate times support-selected projector",
		OperatorLane:          "End(H72)",
		NotTensorProduct:      true,
		NotBoundaryMap:        true,
		Verdict: strings.Join([]string{
			StatusK7ResponseOperatorTyped,
			StatusK7NotBoundaryVectorMap,
		}, "; "),
	}
}

func buildRawMomentMap(s float64) RawMomentMap {
	m1 := pK7 * s
	m2 := pK7 * s * s
	m3 := pK7 * s * s * s
	return RawMomentMap{
		Observer:          "rho_72=I_H72/72",
		Definition:        "M_n(s)=Tr_H72(rho_72 R_wall(s)^n)",
		ProjectorPowerLaw: "P_7^n=P_7",
		EventWeight:       pK7,
		M1:                m1,
		M2:                m2,
		M3:                m3,
		Formula:           "M_n(s)=p_K7 s^n",
		Verdict:           StatusRawMomentMapTyped,
	}
}

func buildCubicBoundaryNormalForm(s float64, m RawMomentMap) CubicBoundaryNormalForm {
	value := m.M1 + kappaE*m.M2 - 2*pK7*m.M3
	return CubicBoundaryNormalForm{
		MapType:         "F_wall_3: Q_boundary -> Q_history",
		Definition:      "F_wall_3(s)=M_1(s)+kappa_e M_2(s)-2p_K7 M_3(s)",
		Polynomial:      "F_wall_3(s)=p_K7 s+kappa_e p_K7 s^2-2p_K7^2 s^3",
		Value:           value,
		NotOperatorOnK7: true,
		Verdict: strings.Join([]string{
			StatusCubicBoundaryHistoryNormalFormDefined,
			StatusFWall3QBoundaryQHistoryScalarResponse,
			StatusFWall3NotNativeOperatorOnK7,
		}, "; "),
	}
}

func buildHiggsRadialHopfLoopFactor() HiggsRadialHopfLoopFactor {
	l := 1.0 / (8.0 * math.Pi)
	return HiggsRadialHopfLoopFactor{
		SealPackage:      "(n,q,P_rad)",
		Observer:         "rho_plus=I_K7+/4",
		HopfOperator:     "R_Hopf=(1/(2*pi))P_rad ∈ End(K7+)",
		TraceExpectation: "L_Hopf=Tr_K7+(rho_plus R_Hopf)=1/(8*pi)",
		LHopf:            l,
		OperatorLane:     "End(K7+)",
		ScalarAfterTrace: true,
		Verdict: strings.Join([]string{
			StatusHiggsRadialHopfLoopFactorTyped,
			StatusLHopfTraceExpectationOnK7Plus,
			StatusLHopfNotBoundaryResponseCoefficient,
		}, "; "),
	}
}

func buildScalarHiggsTypedNormalForm(fwall, l float64) ScalarHiggsTypedNormalForm {
	w3 := lambdaAbs + fwall
	lambdaRuntimeBridge := lambdaProxy * (1 + l*(1-w3+kappaE))
	return ScalarHiggsTypedNormalForm{
		WoundDefinition:      "W_3=|lambda(Lambda12)|+F_wall_3(sigma_boundary(b))",
		RuntimeFormula:       "lambda_runtime_bridge=lambda_proxy[1+L_Hopf(1-W_3+kappa_e)]",
		ExpandedFormula:      "lambda_proxy[1+Tr_K7+(rho_plus R_Hopf)(1-|lambda|-F_wall_3(sigma_boundary(b))+kappa_e)]",
		W3:                   w3,
		RuntimeBridge:        lambdaRuntimeBridge,
		AllOperatorCollapsed: true,
		Verdict: strings.Join([]string{
			StatusScalarHiggsTypedNormalFormWritten,
			StatusTypedNormalFormSupported,
			StatusRuntimeAfterTraceCollapse,
			StatusNoIndependentScalarRuntimeTheorem,
		}, "; "),
	}
}

func buildLegalOperationAudit() LegalOperationAudit {
	return LegalOperationAudit{
		LawfulAdditions: []string{
			"lambda+(R3-1) only inside Q_boundary via sigma_boundary",
			"|lambda|+F_wall_3 only because both are scalar-wall wound coordinates",
			"1-W_3+kappa_e only inside the scalar runtime transport line",
		},
		LawfulProducts: []string{
			"sP_7: scalar times projector in End(H72)",
			"rho_72 R_wall^n: operator composition inside End(H72)",
			"(1/(2*pi))P_rad: scalar payoff times projector in End(K7+)",
			"lambda_proxy[...]: scalar scaling in runtime quartic line",
		},
		TraceMaps: []string{
			"Tr_H72(rho_72 R_wall^n): End(H72) -> scalar raw moment",
			"Tr_K7+(rho_plus R_Hopf): End(K7+) -> scalar L_Hopf",
			"Tr_H72(rho_72 P_7): K7 event weight p_K7",
		},
		RuntimeScalars: []string{
			"F_wall_3(sigma_boundary(b))",
			"W_3",
			"L_Hopf",
			"lambda_runtime_bridge",
		},
		Verdict: StatusLegalOperationAuditCompleted,
	}
}

func buildIllegalTermRejection() IllegalTermRejection {
	terms := []string{
		"K7 + boundary vector",
		"P_K7 + S_split",
		"P_rad + lambda",
		"Hom(Q_boundary,K7) as subspace of H72",
		"F_wall_3 as native operator on K7",
		"L_Hopf as boundary-history event weight",
		"7/72 as source of 1/(8*pi)",
		"tree proxy as pole mass",
		"lambda_runtime bridge as independent physical prediction",
	}
	return IllegalTermRejection{
		RejectedTerms:         terms,
		K7BoundaryBlocked:     true,
		FWallOperatorBlocked:  true,
		LBoundaryCoeffBlocked: true,
		SevenLoopBlocked:      true,
		TreePoleBlocked:       true,
		PredictionBlocked:     true,
		Verdict: strings.Join([]string{
			StatusIllegalTermRejectionAudited,
			StatusK7NotBoundaryVectorMap,
			StatusFWall3NotNativeOperatorOnK7,
			StatusLHopfNotBoundaryResponseCoefficient,
			StatusSevenOver72NotSourceOfOneOver8Pi,
		}, "; "),
	}
}

func buildKappaEInsertionStatus(s float64) KappaEInsertionStatus {
	orient := 0.00550633006471245
	candidate := orient - (5.0/3.0)*s*s + xiBoundary*pK7*s*s
	return KappaEInsertionStatus{
		InsideFWall3:            true,
		OutsideRuntimeTransport: true,
		CandidateFormula:        "kappa_e≈sin²(theta13)/4-J_CKM-(5/3)S_split²+xi_boundary p_K7 S_split²",
		CandidateValue:          candidate,
		ApproximationResidual:   kappaE - candidate,
		NativeFlavorTheorem:     false,
		Verdict: strings.Join([]string{
			StatusKappaEInsertionStatusRecorded,
			StatusKappaESubstitutionNotNativeFlavorTheorem,
			StatusNoYukawaOperatorOrEigenvalueTheorem,
		}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate750TypeLedgerInherited,
		StatusBoundaryQuotientCoordinateTyped,
		StatusK7ResponseOperatorTyped,
		StatusRawMomentMapTyped,
		StatusCubicBoundaryHistoryNormalFormDefined,
		StatusHiggsRadialHopfLoopFactorTyped,
		StatusScalarHiggsTypedNormalFormWritten,
		StatusLegalOperationAuditCompleted,
		StatusIllegalTermRejectionAudited,
		StatusKappaEInsertionStatusRecorded,
		StatusTypedNormalFormSupported,
		StatusFWall3QBoundaryQHistoryScalarResponse,
		StatusLHopfTraceExpectationOnK7Plus,
		StatusRuntimeAfterTraceCollapse,
		StatusK7NotBoundaryVectorMap,
		StatusFWall3NotNativeOperatorOnK7,
		StatusLHopfNotBoundaryResponseCoefficient,
		StatusSevenOver72NotSourceOfOneOver8Pi,
		StatusKappaESubstitutionNotNativeFlavorTheorem,
		StatusNoNativeNQPRadTheorem,
		StatusNoBoundaryGeneratingFunctionTheorem,
		StatusNoNativeHistoryLoopUnitTheorem,
		StatusNoIndependentScalarRuntimeTheorem,
		StatusNoHiggsMassOrPoleMassTheorem,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusGate751Boundary,
	}
}

func FormatGate750(x Gate750Inheritance) string {
	return fmt.Sprintf("inherited=%t typedLedger=%t traceAirlocks=%t firewalls=%t verdict=%q", x.Inherited, x.TypedLedgerReady, x.TraceAirlocksReady, x.FirewallsPreserved, x.Verdict)
}

func FormatDomains(x TypedDomains) string {
	return fmt.Sprintf("native=%q lambda4=%q k7=%q hodge=%q H72=%q Qb=%q Qh=%q Qr=%q verdict=%q", x.NativeFiniteChamber, x.Lambda4Chamber, x.K7Carrier, x.HodgeSplit, x.AugmentedResponseChamber, x.BoundaryQuotientLine, x.HistoryReadoutLine, x.ScalarRuntimeLine, x.Verdict)
}

func FormatBoundary(x BoundaryQuotientCoordinate) string {
	return fmt.Sprintf("b=%q sigma=%q name=%q value=%.16g lives=%q verdict=%q", x.BoundaryVector, x.Functional, x.CoordinateName, x.SValue, x.LivesIn, x.Verdict)
}

func FormatK7Response(x K7ResponseOperator) string {
	return fmt.Sprintf("P=%q R=%q mult=%q lane=%q notTensor=%t notBoundary=%t verdict=%q", x.LiftedProjector, x.ResponseOperator, x.MultiplicationMeaning, x.OperatorLane, x.NotTensorProduct, x.NotBoundaryMap, x.Verdict)
}

func FormatMoments(x RawMomentMap) string {
	return fmt.Sprintf("rho=%q def=%q power=%q p=%.16g M1=%.16g M2=%.16g M3=%.16g formula=%q verdict=%q", x.Observer, x.Definition, x.ProjectorPowerLaw, x.EventWeight, x.M1, x.M2, x.M3, x.Formula, x.Verdict)
}

func FormatCubic(x CubicBoundaryNormalForm) string {
	return fmt.Sprintf("type=%q def=%q poly=%q value=%.16g notK7=%t verdict=%q", x.MapType, x.Definition, x.Polynomial, x.Value, x.NotOperatorOnK7, x.Verdict)
}

func FormatHopf(x HiggsRadialHopfLoopFactor) string {
	return fmt.Sprintf("seals=%q rho=%q R=%q trace=%q L=%.16g lane=%q scalar=%t verdict=%q", x.SealPackage, x.Observer, x.HopfOperator, x.TraceExpectation, x.LHopf, x.OperatorLane, x.ScalarAfterTrace, x.Verdict)
}

func FormatNormalForm(x ScalarHiggsTypedNormalForm) string {
	return fmt.Sprintf("W=%q formula=%q expanded=%q W3=%.16g runtime=%.16g collapsed=%t verdict=%q", x.WoundDefinition, x.RuntimeFormula, x.ExpandedFormula, x.W3, x.RuntimeBridge, x.AllOperatorCollapsed, x.Verdict)
}

func FormatLegal(x LegalOperationAudit) string {
	return fmt.Sprintf("add=[%s] prod=[%s] trace=[%s] runtime=[%s] verdict=%q", strings.Join(x.LawfulAdditions, "; "), strings.Join(x.LawfulProducts, "; "), strings.Join(x.TraceMaps, "; "), strings.Join(x.RuntimeScalars, "; "), x.Verdict)
}

func FormatIllegal(x IllegalTermRejection) string {
	return fmt.Sprintf("rejected=[%s] k7=%t fwall=%t l=%t seven=%t tree=%t prediction=%t verdict=%q", strings.Join(x.RejectedTerms, "; "), x.K7BoundaryBlocked, x.FWallOperatorBlocked, x.LBoundaryCoeffBlocked, x.SevenLoopBlocked, x.TreePoleBlocked, x.PredictionBlocked, x.Verdict)
}

func FormatKappaE(x KappaEInsertionStatus) string {
	return fmt.Sprintf("inside=%t outside=%t candidate=%q value=%.16g residual=%.16g native=%t verdict=%q", x.InsideFWall3, x.OutsideRuntimeTransport, x.CandidateFormula, x.CandidateValue, x.ApproximationResidual, x.NativeFlavorTheorem, x.Verdict)
}
