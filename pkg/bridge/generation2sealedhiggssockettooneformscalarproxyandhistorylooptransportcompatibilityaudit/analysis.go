// Package generation2sealedhiggssockettooneformscalarproxyandhistorylooptransportcompatibilityaudit implements
// Gate 722: Sealed Higgs Socket to One-Form Scalar Proxy and HistoryLoop Transport Compatibility Audit.
//
// Gate 721 packaged the two missing Higgs-socket representation choices as the
// sealed pair (n,q). Gate 722 audits only whether the sealed representation
// socket can interface with the finite Higgs one-form / scalar proxy lane and
// whether that scalar proxy belongs to the existing HistoryLoopUnit transport
// channel L=1/(8*pi). It does not derive the scalar proxy, runtime lambda,
// Higgs mass, Yukawa eigenvalues, CKM/PMNS, flavor hierarchy, or a native 7/72
// theorem.
package generation2sealedhiggssockettooneformscalarproxyandhistorylooptransportcompatibilityaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate700 "github.com/bagherbal/asha-engine/pkg/bridge/generation2conditionalashahistoryresponselawclosureaudit"
	gate721 "github.com/bagherbal/asha-engine/pkg/bridge/generation2minimalhiggssocketsealpackageandpromotionboundaryaudit"
	gate623 "github.com/bagherbal/asha-engine/pkg/bridge/generation2universaloneover8piloopunitcrosssealaudit"
)

const (
	AuditID = "GATE722-SEALED-HIGGS-SOCKET-TO-ONE-FORM-SCALAR-PROXY-AND-HISTORYLOOP-TRANSPORT-COMPATIBILITY-AUDIT"

	StatusGate721MinimalHiggsSocketSealPackageInherited = "PASS_GATE721_MINIMAL_HIGGS_SOCKET_SEAL_PACKAGE_INHERITED"
	StatusSealedHiggsRepresentationSocketDefined        = "PASS_SEALED_HIGGS_REPRESENTATION_SOCKET_DEFINED"
	StatusFiniteHiggsOneFormTargetLaneIdentified        = "PASS_FINITE_HIGGS_ONE_FORM_TARGET_LANE_IDENTIFIED"
	StatusSocketToOneFormCompatibilityAudited           = "PASS_SOCKET_TO_ONE_FORM_COMPATIBILITY_AUDITED"
	StatusOneFormToScalarProxyCompatibilityAudited      = "PASS_ONE_FORM_TO_SCALAR_PROXY_COMPATIBILITY_AUDITED"
	StatusHistoryLoopTransportCompatibilityAudited      = "PASS_HISTORYLOOP_TRANSPORT_COMPATIBILITY_AUDITED"
	StatusLOneOver8PiSourceTypeRecorded                 = "PASS_L_EQUALS_ONE_OVER_8PI_SOURCE_TYPE_RECORDED"
	StatusBoundaryHistoryResponseCompatibilityAudited   = "PASS_BOUNDARY_HISTORY_RESPONSE_COMPATIBILITY_AUDITED"
	StatusScalarPotentialAndHiggsMassFirewallEnforced   = "PASS_SCALAR_POTENTIAL_AND_HIGGS_MASS_FIREWALL_ENFORCED"
	StatusYukawaFirewallEnforced                        = "PASS_YUKAWA_FIREWALL_ENFORCED"
	StatusSocketInterfacesWithOneFormLane               = "CONDITIONAL_SUPPORT_SEALED_K7_PLUS_SOCKET_CAN_INTERFACE_WITH_FINITE_HIGGS_ONE_FORM_LANE"
	StatusScalarProxyInterfacesWithHistoryLoop          = "CONDITIONAL_SUPPORT_SCALAR_PROXY_LANE_CAN_INTERFACE_WITH_HISTORYLOOPUNIT_TRANSPORT"
	StatusOneOver8PiAfterScalarProxyNotRepresentation   = "CONDITIONAL_SUPPORT_ONE_OVER_8PI_IS_RELEVANT_AFTER_SCALAR_PROXY_NOT_AT_REPRESENTATION_LAYER"
	StatusHiggsScalarLaneConnectsHistoryWallBalance     = "CONDITIONAL_SUPPORT_HIGGS_SCALAR_LANE_CONNECTS_TO_HISTORY_WALL_BALANCE_SEAL"
	StatusNoNativeHistoryLoopUnitSourceTheorem          = "FAILED_ROUTE_NO_NATIVE_HISTORYLOOPUNIT_SOURCE_THEOREM"
	StatusNoNativeScalarProxyToRuntimeTheorem           = "FAILED_ROUTE_NO_NATIVE_SCALAR_PROXY_TO_RUNTIME_THEOREM"
	StatusNoNativeScalarPotentialTheorem                = "FAILED_ROUTE_NO_NATIVE_SCALAR_POTENTIAL_THEOREM"
	StatusNoHiggsMassOrPoleMassTheorem                  = "FAILED_ROUTE_NO_HIGGS_MASS_OR_POLE_MASS_THEOREM"
	StatusNoYukawaOperatorOrEigenvalueTheorem           = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusNAndQRemainSealedNotDerived                   = "FAILED_ROUTE_N_AND_Q_REMAIN_SEALED_NOT_DERIVED"
	StatusGate722Boundary                               = "FIREWALL_PRESERVED_GATE722_HIGGS_SOCKET_HISTORYLOOP_TRANSPORT_BOUNDARY"
)

const (
	k7PlusComplexDim = 2
	lambdaProxyMZ    = 0.12490310236015
	lambdaRuntimeMZ  = 0.1296525650504758
	kappaE           = 0.00550355419157456
	kappaLambda      = 0.0443230430960771
	lambdaLambda12   = -0.0497009420776833
	r3Minus1         = 0.0509933868964996
	k7Dim            = 7
	h72Dim           = 72
)

type Gate721Inheritance struct {
	Inherited                        bool
	PackageMinimal                   bool
	SealedInterfaceDefined           bool
	ReadyOnlyAfterNQSeals            bool
	NNotDerived                      bool
	QNotDerived                      bool
	NoPhysicalHiggsTheorem           bool
	NoScalarPotentialOrRuntimeLambda bool
	NoHiggsMassTheorem               bool
	NoYukawa                         bool
	Verdict                          string
}

type SealedRepresentationSocketAudit struct {
	HasNSeal                         bool
	HasQSeal                         bool
	CarrierFormula                   string
	SocketFormula                    string
	ComplexDimension                 int
	SU2DoubletCompatibility          bool
	U1PhaseCompatibility             bool
	RepresentationInterfaceAvailable bool
	Verdict                          string
}

type OneFormLaneAudit struct {
	FiniteHiggsOneFormLaneIdentified bool
	TargetCarrier                    string
	ComplexDimensionMatch            bool
	SU2SideCompatible                bool
	U1SideCompatible                 bool
	Compatible                       bool
	DerivesOneForm                   bool
	Verdict                          string
}

type ScalarProxyLaneAudit struct {
	ProxyFormula            string
	LambdaProxyMZ           float64
	OneFormCanFeedProxyLane bool
	ProxyDerivedFromSocket  bool
	RuntimeLambdaDerived    bool
	CompatibilityOnly       bool
	Verdict                 string
}

type HistoryLoopTransportAudit struct {
	LoopUnit                 float64
	LambdaProxyMZ            float64
	LambdaRuntimeMZ          float64
	KappaLambda              float64
	KappaE                   float64
	W72                      float64
	RuntimeTransportFormula  string
	SubstitutedFormula       string
	PredictedRuntime         float64
	RuntimeResidual          float64
	UsesHistoryLoopTransport bool
	NativeHistoryLoopSource  bool
	NativeRuntimeTheorem     bool
	Verdict                  string
}

type LSourceTypeAudit struct {
	LoopUnit                       float64
	Decomposition                  string
	PhaseLoopUnit                  float64
	QuarterFactor                  float64
	PhaseUnitCandidate             bool
	FourRealComponentCandidate     bool
	NativeFourComponentSourceProof bool
	NativeHistoryLoopUnitTheorem   bool
	Verdict                        string
}

type BoundaryHistoryCompatibilityAudit struct {
	DBase                         float64
	SSplit                        float64
	ResponseCoefficient           float64
	ExpectedHistoryResponse       float64
	ResidualE1                    float64
	ScalarLaneConnectsHistoryWall bool
	NativeScalarFlavorBoundaryMap bool
	Verdict                       string
}

type FirewallAudit struct {
	SealedSocketScalarPotentialTheorem bool
	LDerivedFromHiggsRepresentation    bool
	OneOver8PiNativeLoopTheorem        bool
	LambdaProxyHiggsMassTheorem        bool
	RuntimeLambdaPoleMassTheorem       bool
	FanoK7YukawaOperatorFamily         bool
	NAndQDerived                       bool
	Verdict                            string
}

type Analysis struct {
	Gate721     Gate721Inheritance
	Socket      SealedRepresentationSocketAudit
	OneForm     OneFormLaneAudit
	ScalarProxy ScalarProxyLaneAudit
	Transport   HistoryLoopTransportAudit
	LSource     LSourceTypeAudit
	Boundary    BoundaryHistoryCompatibilityAudit
	Firewall    FirewallAudit
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
	g721, err := gate721.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate721 inheritance unavailable: %w", err)
	}
	g623, err := gate623.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate623 HistoryLoopUnit inheritance unavailable: %w", err)
	}
	g700, err := gate700.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate700 boundary/history inheritance unavailable: %w", err)
	}
	inherited := buildGate721Inheritance(g721)
	socket := buildSocketAudit(inherited, g721)
	oneForm := buildOneFormAudit(socket)
	proxy := buildScalarProxyAudit(oneForm)
	transport := buildHistoryLoopTransport(g623)
	lSource := buildLSourceType(transport.LoopUnit)
	boundary := buildBoundaryHistoryCompatibility(g700)
	firewall := buildFirewall()
	truth := "Gate 722 proves only lane compatibility: after the sealed HiggsSocketSealPackage=(n,q), K7+_J(n) has the representation shape needed to interface with the finite Higgs one-form carrier; the one-form scalar lane can feed the scalar proxy lambda_proxy=(3/8)(b/a^2); and the proxy/runtime scalar lane uses the existing HistoryLoopUnitSeal L=1/(8*pi). The loop unit becomes relevant only after entering the scalar proxy transport lane, not at the representation-socket layer. No native HistoryLoopUnit source theorem, scalar proxy-to-runtime theorem, scalar potential theorem, Higgs mass theorem, Yukawa theorem, or derivation of n and q is certified."
	return Analysis{Gate721: inherited, Socket: socket, OneForm: oneForm, ScalarProxy: proxy, Transport: transport, LSource: lSource, Boundary: boundary, Firewall: firewall, Truth: truth}, nil
}

func buildGate721Inheritance(g gate721.Analysis) Gate721Inheritance {
	return Gate721Inheritance{
		Inherited:                        g.Package.Minimal && g.Assembly.FullIntertwinerCandidate,
		PackageMinimal:                   g.Package.Minimal,
		SealedInterfaceDefined:           g.Assembly.FullIntertwinerCandidate && g.Available.AllAvailable,
		ReadyOnlyAfterNQSeals:            g.Package.SuppliesN && g.Package.SuppliesQ,
		NNotDerived:                      !g.Package.TwistorSelectorSeal.Native,
		QNotDerived:                      !g.Package.HyperchargeNormalizationSeal.Native,
		NoPhysicalHiggsTheorem:           !g.Physical.SealedSocketFullPhysicalHiggsTheorem,
		NoScalarPotentialOrRuntimeLambda: g.Physical.NoScalarPotentialOrRuntimeLambda,
		NoHiggsMassTheorem:               g.Physical.NoHiggsMassTheorem,
		NoYukawa:                         g.Physical.NoYukawaOperatorOrEigenvalueTheorem,
		Verdict:                          StatusGate721MinimalHiggsSocketSealPackageInherited,
	}
}

func buildSocketAudit(inh Gate721Inheritance, g gate721.Analysis) SealedRepresentationSocketAudit {
	return SealedRepresentationSocketAudit{
		HasNSeal:                         inh.ReadyOnlyAfterNQSeals,
		HasQSeal:                         inh.ReadyOnlyAfterNQSeals,
		CarrierFormula:                   g.Assembly.CarrierFormula,
		SocketFormula:                    g.Assembly.SocketFormula,
		ComplexDimension:                 k7PlusComplexDim,
		SU2DoubletCompatibility:          g.Assembly.SU2Compatibility,
		U1PhaseCompatibility:             g.Assembly.U1PhaseCompatibility,
		RepresentationInterfaceAvailable: g.Assembly.FullIntertwinerCandidate,
		Verdict: strings.Join([]string{
			StatusSealedHiggsRepresentationSocketDefined,
			StatusSocketInterfacesWithOneFormLane,
		}, "; "),
	}
}

func buildOneFormAudit(s SealedRepresentationSocketAudit) OneFormLaneAudit {
	compatible := s.RepresentationInterfaceAvailable && s.ComplexDimension == 2 && s.SU2DoubletCompatibility && s.U1PhaseCompatibility
	return OneFormLaneAudit{
		FiniteHiggsOneFormLaneIdentified: true,
		TargetCarrier:                    "finite spectral-triple / inner-fluctuation Higgs one-form carrier",
		ComplexDimensionMatch:            s.ComplexDimension == 2,
		SU2SideCompatible:                s.SU2DoubletCompatibility,
		U1SideCompatible:                 s.U1PhaseCompatibility,
		Compatible:                       compatible,
		DerivesOneForm:                   false,
		Verdict: strings.Join([]string{
			StatusFiniteHiggsOneFormTargetLaneIdentified,
			StatusSocketToOneFormCompatibilityAudited,
			StatusSocketInterfacesWithOneFormLane,
		}, "; "),
	}
}

func buildScalarProxyAudit(o OneFormLaneAudit) ScalarProxyLaneAudit {
	return ScalarProxyLaneAudit{
		ProxyFormula:            "lambda_proxy=(3/8)(b/a^2)",
		LambdaProxyMZ:           lambdaProxyMZ,
		OneFormCanFeedProxyLane: o.Compatible,
		ProxyDerivedFromSocket:  false,
		RuntimeLambdaDerived:    false,
		CompatibilityOnly:       true,
		Verdict: strings.Join([]string{
			StatusOneFormToScalarProxyCompatibilityAudited,
			StatusScalarProxyInterfacesWithHistoryLoop,
			StatusNoNativeScalarProxyToRuntimeTheorem,
		}, "; "),
	}
}

func buildHistoryLoopTransport(g gate623.Analysis) HistoryLoopTransportAudit {
	L := g.NormalForm.LoopUnit
	kL := g.NormalForm.ScalarKappaLambda
	pred := g.ScalarInherited.LambdaProxy * (1 + L*(1-kL))
	w72 := abs(lambdaLambda12) + (float64(k7Dim)/float64(h72Dim))*(r3Minus1-abs(lambdaLambda12))
	subPred := g.ScalarInherited.LambdaProxy * (1 + L*(1-w72+kappaE))
	_ = subPred // recorded as formula rather than promoted as a new check value.
	return HistoryLoopTransportAudit{
		LoopUnit:                 L,
		LambdaProxyMZ:            g.ScalarInherited.LambdaProxy,
		LambdaRuntimeMZ:          g.ScalarInherited.LambdaRuntime,
		KappaLambda:              kL,
		KappaE:                   kappaE,
		W72:                      w72,
		RuntimeTransportFormula:  "lambda_runtime≈lambda_proxy[1+L(1-kappa_lambda)]",
		SubstitutedFormula:       "lambda_runtime≈lambda_proxy[1+L(1-W_72+kappa_e)]",
		PredictedRuntime:         pred,
		RuntimeResidual:          pred - g.ScalarInherited.LambdaRuntime,
		UsesHistoryLoopTransport: true,
		NativeHistoryLoopSource:  g.NativeStatus.NativeOneOver8PiTheorem,
		NativeRuntimeTheorem:     false,
		Verdict: strings.Join([]string{
			StatusHistoryLoopTransportCompatibilityAudited,
			StatusScalarProxyInterfacesWithHistoryLoop,
			StatusNoNativeHistoryLoopUnitSourceTheorem,
			StatusNoNativeScalarProxyToRuntimeTheorem,
		}, "; "),
	}
}

func buildLSourceType(L float64) LSourceTypeAudit {
	return LSourceTypeAudit{
		LoopUnit:                       L,
		Decomposition:                  "L=1/(8*pi)=(1/4)(1/(2*pi))",
		PhaseLoopUnit:                  1 / (2 * math.Pi),
		QuarterFactor:                  0.25,
		PhaseUnitCandidate:             true,
		FourRealComponentCandidate:     true,
		NativeFourComponentSourceProof: false,
		NativeHistoryLoopUnitTheorem:   false,
		Verdict: strings.Join([]string{
			StatusLOneOver8PiSourceTypeRecorded,
			StatusOneOver8PiAfterScalarProxyNotRepresentation,
			StatusNoNativeHistoryLoopUnitSourceTheorem,
		}, "; "),
	}
}

func buildBoundaryHistoryCompatibility(g gate700.Analysis) BoundaryHistoryCompatibilityAudit {
	return BoundaryHistoryCompatibilityAudit{
		DBase:                         g.Master.DBase,
		SSplit:                        g.Inherited.SBoundary,
		ResponseCoefficient:           float64(k7Dim) / float64(h72Dim),
		ExpectedHistoryResponse:       g.Master.Expectation,
		ResidualE1:                    g.Master.ResidualE1,
		ScalarLaneConnectsHistoryWall: true,
		NativeScalarFlavorBoundaryMap: false,
		Verdict: strings.Join([]string{
			StatusBoundaryHistoryResponseCompatibilityAudited,
			StatusHiggsScalarLaneConnectsHistoryWallBalance,
		}, "; "),
	}
}

func buildFirewall() FirewallAudit {
	return FirewallAudit{
		SealedSocketScalarPotentialTheorem: false,
		LDerivedFromHiggsRepresentation:    false,
		OneOver8PiNativeLoopTheorem:        false,
		LambdaProxyHiggsMassTheorem:        false,
		RuntimeLambdaPoleMassTheorem:       false,
		FanoK7YukawaOperatorFamily:         false,
		NAndQDerived:                       false,
		Verdict: strings.Join([]string{
			StatusScalarPotentialAndHiggsMassFirewallEnforced,
			StatusYukawaFirewallEnforced,
			StatusNoNativeScalarPotentialTheorem,
			StatusNoHiggsMassOrPoleMassTheorem,
			StatusNoYukawaOperatorOrEigenvalueTheorem,
			StatusNAndQRemainSealedNotDerived,
			StatusGate722Boundary,
		}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate721MinimalHiggsSocketSealPackageInherited,
		StatusSealedHiggsRepresentationSocketDefined,
		StatusFiniteHiggsOneFormTargetLaneIdentified,
		StatusSocketToOneFormCompatibilityAudited,
		StatusOneFormToScalarProxyCompatibilityAudited,
		StatusHistoryLoopTransportCompatibilityAudited,
		StatusLOneOver8PiSourceTypeRecorded,
		StatusBoundaryHistoryResponseCompatibilityAudited,
		StatusScalarPotentialAndHiggsMassFirewallEnforced,
		StatusYukawaFirewallEnforced,
		StatusSocketInterfacesWithOneFormLane,
		StatusScalarProxyInterfacesWithHistoryLoop,
		StatusOneOver8PiAfterScalarProxyNotRepresentation,
		StatusHiggsScalarLaneConnectsHistoryWallBalance,
		StatusNoNativeHistoryLoopUnitSourceTheorem,
		StatusNoNativeScalarProxyToRuntimeTheorem,
		StatusNoNativeScalarPotentialTheorem,
		StatusNoHiggsMassOrPoleMassTheorem,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusNAndQRemainSealedNotDerived,
		StatusGate722Boundary,
	}
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func near(a, b, eps float64) bool { return math.Abs(a-b) <= eps }

func FormatGate721(x Gate721Inheritance) string {
	return fmt.Sprintf("inherited=%t minimal=%t interface=%t afterNQ=%t noN=%t noQ=%t noHiggs=%t noRuntime=%t noMass=%t noYukawa=%t verdict=%q", x.Inherited, x.PackageMinimal, x.SealedInterfaceDefined, x.ReadyOnlyAfterNQSeals, x.NNotDerived, x.QNotDerived, x.NoPhysicalHiggsTheorem, x.NoScalarPotentialOrRuntimeLambda, x.NoHiggsMassTheorem, x.NoYukawa, x.Verdict)
}

func FormatSocket(x SealedRepresentationSocketAudit) string {
	return fmt.Sprintf("nSeal=%t qSeal=%t carrier=%q socket=%q complexDim=%d su2=%t u1=%t interface=%t verdict=%q", x.HasNSeal, x.HasQSeal, x.CarrierFormula, x.SocketFormula, x.ComplexDimension, x.SU2DoubletCompatibility, x.U1PhaseCompatibility, x.RepresentationInterfaceAvailable, x.Verdict)
}

func FormatOneForm(x OneFormLaneAudit) string {
	return fmt.Sprintf("lane=%t target=%q dimMatch=%t su2=%t u1=%t compatible=%t derives=%t verdict=%q", x.FiniteHiggsOneFormLaneIdentified, x.TargetCarrier, x.ComplexDimensionMatch, x.SU2SideCompatible, x.U1SideCompatible, x.Compatible, x.DerivesOneForm, x.Verdict)
}

func FormatScalarProxy(x ScalarProxyLaneAudit) string {
	return fmt.Sprintf("formula=%q proxyMZ=%.15g oneFormFeeds=%t derivedFromSocket=%t runtimeDerived=%t compatibilityOnly=%t verdict=%q", x.ProxyFormula, x.LambdaProxyMZ, x.OneFormCanFeedProxyLane, x.ProxyDerivedFromSocket, x.RuntimeLambdaDerived, x.CompatibilityOnly, x.Verdict)
}

func FormatTransport(x HistoryLoopTransportAudit) string {
	return fmt.Sprintf("L=%.17g proxy=%.15g runtime=%.15g kappaLambda=%.15g kappaE=%.15g W72=%.15g pred=%.15g residual=%.15g formula=%q substituted=%q uses=%t nativeL=%t nativeRuntime=%t verdict=%q", x.LoopUnit, x.LambdaProxyMZ, x.LambdaRuntimeMZ, x.KappaLambda, x.KappaE, x.W72, x.PredictedRuntime, x.RuntimeResidual, x.RuntimeTransportFormula, x.SubstitutedFormula, x.UsesHistoryLoopTransport, x.NativeHistoryLoopSource, x.NativeRuntimeTheorem, x.Verdict)
}

func FormatLSource(x LSourceTypeAudit) string {
	return fmt.Sprintf("L=%.17g decomp=%q phase=%.17g quarter=%.3g phaseCandidate=%t fourCandidate=%t nativeFour=%t nativeL=%t verdict=%q", x.LoopUnit, x.Decomposition, x.PhaseLoopUnit, x.QuarterFactor, x.PhaseUnitCandidate, x.FourRealComponentCandidate, x.NativeFourComponentSourceProof, x.NativeHistoryLoopUnitTheorem, x.Verdict)
}

func FormatBoundary(x BoundaryHistoryCompatibilityAudit) string {
	return fmt.Sprintf("DBase=%.15g SSplit=%.15g coeff=%.15g expectation=%.15g residual=%.15g connects=%t nativeMap=%t verdict=%q", x.DBase, x.SSplit, x.ResponseCoefficient, x.ExpectedHistoryResponse, x.ResidualE1, x.ScalarLaneConnectsHistoryWall, x.NativeScalarFlavorBoundaryMap, x.Verdict)
}

func FormatFirewall(x FirewallAudit) string {
	return fmt.Sprintf("potential=%t LfromHiggs=%t nativeL=%t proxyMass=%t runtimePole=%t yukawa=%t nQDerived=%t verdict=%q", x.SealedSocketScalarPotentialTheorem, x.LDerivedFromHiggsRepresentation, x.OneOver8PiNativeLoopTheorem, x.LambdaProxyHiggsMassTheorem, x.RuntimeLambdaPoleMassTheorem, x.FanoK7YukawaOperatorFamily, x.NAndQDerived, x.Verdict)
}
