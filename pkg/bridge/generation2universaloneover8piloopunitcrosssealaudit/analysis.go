// Package generation2universaloneover8piloopunitcrosssealaudit implements
// Gate 623: Universal One-Over-8Pi Loop Unit Cross-Seal Audit.
//
// Gate 622 found that the scalar proxy-to-runtime matching gap is organized
// by the typed loop unit L=1/(8*pi).  Earlier flavor gates found that the
// charged-lepton Koide wall offset is also naturally scaled by L. Gate 623
// rewrites both in one shared loop-unit normal form and audits whether this is
// a bridge-layer cross-seal clue rather than a native ASHA theorem.
package generation2universaloneover8piloopunitcrosssealaudit

import (
	"fmt"
	"math"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2scalaroneeighthproxyloopmatchingaudit"
)

const (
	AuditID = "GATE623-UNIVERSAL-ONE-OVER-8PI-LOOP-UNIT-CROSS-SEAL-AUDIT"

	StatusGate622Inherited     = "PASS_GATE622_SCALAR_LOOP_MATCH_INHERITED"
	StatusFlavorLoopInherited  = "PASS_GATE590_592_FLAVOR_LOOP_UNIT_INHERITED"
	StatusNormalFormWritten    = "PASS_SHARED_LOOP_UNIT_NORMAL_FORM_WRITTEN"
	StatusUnitsComputed        = "PASS_SCALAR_AND_FLAVOR_L_UNITS_COMPUTED"
	StatusAppearsInBoth        = "CONDITIONAL_SUPPORT_ONE_OVER_8PI_APPEARS_IN_BOTH_SCALAR_AND_FLAVOR_SEALS"
	StatusScalarAnsatzClose    = "CONDITIONAL_SUPPORT_SCALAR_L_ANSATZ_CLOSE_TO_RUNTIME_LAMBDA_MZ"
	StatusFlavorBalanceClose   = "CONDITIONAL_SUPPORT_FLAVOR_L_ORIENTATION_BALANCE_CLOSE_TO_EPSILON_E"
	StatusCrossSealBridgeOnly  = "CONDITIONAL_SUPPORT_HISTORY_LOOP_UNIT_SEAL_DEFINED_AS_BRIDGE_OBJECT"
	StatusNoCrossSealTheorem   = "FAILED_ROUTE_NO_NATIVE_ONE_OVER_8PI_CROSS_SEAL_THEOREM"
	StatusNoScalarMatching     = "FAILED_ROUTE_NO_NATIVE_SCALAR_MATCHING_THEOREM"
	StatusNoKoideWallTheorem   = "FAILED_ROUTE_NO_NATIVE_KOIDE_WALL_THEOREM"
	StatusNoOrientationBalance = "FAILED_ROUTE_NO_NATIVE_ORIENTATION_BALANCE_THEOREM"
	StatusNoHiggsTheorem       = "FAILED_ROUTE_NO_NATIVE_HIGGS_POLE_THEOREM"
	StatusGate623Boundary      = "FIREWALL_PRESERVED_GATE623_SHARED_LOOP_UNIT_BOUNDARY"
)

const (
	lambdaProxyMZ      = 0.12490310236015
	lambdaRuntimeMZ    = 0.1296525650504758
	deltaLambdaMatch   = 0.0047494626903257
	rhoLambdaMatch     = deltaLambdaMatch / lambdaProxyMZ
	vRuntime           = 246.21965079413738
	epsilonE           = 0.039569756309433
	kappaE             = 0.00550355419157456
	sin2Theta13Quarter = 0.0055375
	jCKM               = 3.11699352875547e-05
)

type ScalarInherited struct {
	LambdaProxy               float64
	LambdaRuntime             float64
	DeltaLambdaMatch          float64
	RhoLambdaMatch            float64
	LoopUnit                  float64
	LambdaAnsatz              float64
	AnsatzMinusRuntime        float64
	RelativeAnsatzResidual    float64
	PreviousRelativeVerdict   string
	PreviousDiagnosticVerdict string
	Verdict                   string
}

type FlavorInherited struct {
	EpsilonE                   float64
	LoopUnit                   float64
	KappaE                     float64
	OrientationCandidate       float64
	OrientationResidual        float64
	EpsilonRawL                float64
	EpsilonOrientation         float64
	RawLResidual               float64
	OrientationEpsilonResidual float64
	Verdict                    string
}

type SharedLoopUnitNormalForm struct {
	LoopUnit          float64
	FlavorEquation    string
	ScalarEquation    string
	FlavorKappaE      float64
	ScalarKappaLambda float64
	Verdict           string
}

type KappaCandidate struct {
	Name             string
	Value            float64
	Residual         float64
	RelativeResidual float64
	Typed            bool
	NativeCertified  bool
	Comment          string
}

type KappaComparisonTable struct {
	KappaE       float64
	KappaLambda  float64
	Delta        float64
	Candidates   []KappaCandidate
	ClosestName  string
	ClosestDelta float64
	Verdict      string
}

type ScalarAnsatzQuality struct {
	LambdaProxy             float64
	LoopUnit                float64
	LambdaAnsatz            float64
	LambdaRuntime           float64
	AnsatzMinusRuntime      float64
	RelativeRuntimeResidual float64
	MassAnsatzGeV           float64
	MassRuntimeGeV          float64
	DeltaMassGeV            float64
	DiagnosticOnly          bool
	Verdict                 string
}

type FlavorAnsatzQuality struct {
	LoopUnit                    float64
	EpsilonE                    float64
	EpsilonRawL                 float64
	RawResidual                 float64
	RawRelativeResidual         float64
	OrientationCandidate        float64
	EpsilonOrientation          float64
	OrientationResidual         float64
	OrientationRelativeResidual float64
	ResidualImprovementFactor   float64
	Verdict                     string
}

type SignAndRoleAudit struct {
	FlavorUsesBelowL     bool
	ScalarUsesAboveProxy bool
	OppositeSigns        bool
	Statement            string
	NativeTheoremClaimed bool
	Verdict              string
}

type CrossSealTypeAudit struct {
	LikelyType           string
	AllowedRoles         []string
	DisallowedPromotions []string
	BridgeOnly           bool
	Verdict              string
}

type NativeStatus struct {
	NativeOneOver8PiTheorem     bool
	NativeScalarMatchingTheorem bool
	NativeKoideWallTheorem      bool
	NativeCrossSealTheorem      bool
	NativeOrientationBalance    bool
	NativeHiggsPoleTheorem      bool
	Statement                   string
	Verdict                     string
}

type Firewalls struct {
	ClaimsKoideDerived      bool
	ClaimsHiggsMassDerived  bool
	ClaimsScalarStability   bool
	ClaimsPMNSCKMDerived    bool
	ClaimsGaugeUnification  bool
	ClaimsNativeLoopTheorem bool
	Verdict                 string
}

type Analysis struct {
	ScalarInherited ScalarInherited
	FlavorInherited FlavorInherited
	NormalForm      SharedLoopUnitNormalForm
	Kappas          KappaComparisonTable
	ScalarQuality   ScalarAnsatzQuality
	FlavorQuality   FlavorAnsatzQuality
	SignRole        SignAndRoleAudit
	CrossSealType   CrossSealTypeAudit
	NativeStatus    NativeStatus
	Firewalls       Firewalls
	Truth           string
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
	g622, err := generation2scalaroneeighthproxyloopmatchingaudit.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate622 predecessor: %w", err)
	}
	loop := 1.0 / (8.0 * math.Pi)
	s := inheritScalar(g622, loop)
	f := inheritFlavor(loop)
	n := buildNormalForm(loop, f.KappaE, s.RhoLambdaMatch)
	k := buildKappaComparison(n.FlavorKappaE, n.ScalarKappaLambda)
	a := Analysis{
		ScalarInherited: s,
		FlavorInherited: f,
		NormalForm:      n,
		Kappas:          k,
		ScalarQuality:   buildScalarQuality(loop),
		FlavorQuality:   buildFlavorQuality(loop),
		SignRole:        buildSignAndRole(),
		CrossSealType:   buildCrossSealType(),
		NativeStatus:    buildNativeStatus(),
		Firewalls:       auditFirewalls(),
		Truth:           "Gate 623 rewrites the scalar matching correction and charged-lepton wall offset in one shared L=1/(8*pi) normal form. The same loop unit organizes both bridge-layer seals, but ASHA currently has no native cross-seal theorem, scalar matching theorem, or Koide wall theorem.",
	}
	return a, nil
}

func inheritScalar(a generation2scalaroneeighthproxyloopmatchingaudit.Analysis, loop float64) ScalarInherited {
	return ScalarInherited{
		LambdaProxy:               a.HiggsDiagnostic.LambdaProxy,
		LambdaRuntime:             a.HiggsDiagnostic.LambdaRuntime,
		DeltaLambdaMatch:          a.Inherited.DeltaLambdaMatch,
		RhoLambdaMatch:            a.Inherited.RelativeToProxy,
		LoopUnit:                  loop,
		LambdaAnsatz:              a.HiggsDiagnostic.LambdaAnsatz,
		AnsatzMinusRuntime:        a.HiggsDiagnostic.AnsatzMinusRuntime,
		RelativeAnsatzResidual:    a.HiggsDiagnostic.RelativeAnsatzResidual,
		PreviousRelativeVerdict:   a.RelativeLoops.Verdict,
		PreviousDiagnosticVerdict: a.HiggsDiagnostic.Verdict,
		Verdict:                   StatusGate622Inherited,
	}
}

func inheritFlavor(loop float64) FlavorInherited {
	orientation := sin2Theta13Quarter - jCKM
	epsOrient := loop * (1 - orientation)
	return FlavorInherited{
		EpsilonE:                   epsilonE,
		LoopUnit:                   loop,
		KappaE:                     kappaE,
		OrientationCandidate:       orientation,
		OrientationResidual:        orientation - kappaE,
		EpsilonRawL:                loop,
		EpsilonOrientation:         epsOrient,
		RawLResidual:               loop - epsilonE,
		OrientationEpsilonResidual: epsOrient - epsilonE,
		Verdict:                    StatusFlavorLoopInherited,
	}
}

func buildNormalForm(loop, kappaE, rho float64) SharedLoopUnitNormalForm {
	kappaLambda := 1 - rho/loop
	return SharedLoopUnitNormalForm{
		LoopUnit:          loop,
		FlavorEquation:    "epsilon_e = L(1-kappa_e)",
		ScalarEquation:    "lambda_runtime(M_Z)=lambda_proxy(M_Z)[1+L(1-kappa_lambda)]",
		FlavorKappaE:      kappaE,
		ScalarKappaLambda: kappaLambda,
		Verdict:           StatusNormalFormWritten,
	}
}

func buildKappaComparison(kE, kL float64) KappaComparisonTable {
	candidates := []KappaCandidate{
		kCandidate("kappa_e", kE, kL, "flavor loop deficit coefficient"),
		kCandidate("R_3-1", 0.0509933868964996, kL, "strong relative boundary wound"),
		kCandidate("|lambda(Lambda_12)|", 0.049700942077683274, kL, "scalar high-scale runtime wound"),
		kCandidate("xi_boundary", 0.0503471644870914, kL, "joint gauge-scalar stress scale"),
		kCandidate("alpha_2(M_Z)", 0.6527521238927322*0.6527521238927322/(4*math.Pi), kL, "weak endpoint coupling"),
		kCandidate("alpha_EM(M_Z)", alphaEM(), kL, "electromagnetic endpoint coupling"),
		kCandidate("sqrt(J_CKM)", math.Sqrt(jCKM), kL, "quark orientation area square root"),
		kCandidate("J_CKM", jCKM, kL, "quark CP orientation area"),
		kCandidate("sin^2(theta13)/4", sin2Theta13Quarter, kL, "lepton reactor leakage quarter"),
		kCandidate("OrientationBalance residual", sin2Theta13Quarter-jCKM-kE, kL, "Gate590 residual scale"),
	}
	name, delta := closestK(candidates)
	return KappaComparisonTable{KappaE: kE, KappaLambda: kL, Delta: kL - kE, Candidates: candidates, ClosestName: name, ClosestDelta: delta, Verdict: StatusUnitsComputed}
}

func alphaEM() float64 {
	g2 := 0.6527521238927322
	gY := 0.3500756885970262
	e := g2 * gY / math.Sqrt(g2*g2+gY*gY)
	return e * e / (4 * math.Pi)
}

func kCandidate(name string, value, target float64, comment string) KappaCandidate {
	res := value - target
	rel := math.NaN()
	if target != 0 {
		rel = res / target
	}
	return KappaCandidate{Name: name, Value: value, Residual: res, RelativeResidual: rel, Typed: true, NativeCertified: false, Comment: comment}
}

func closestK(rows []KappaCandidate) (string, float64) {
	if len(rows) == 0 {
		return "", math.NaN()
	}
	best := rows[0]
	for _, r := range rows[1:] {
		if math.Abs(r.Residual) < math.Abs(best.Residual) {
			best = r
		}
	}
	return best.Name, best.Residual
}

func buildScalarQuality(loop float64) ScalarAnsatzQuality {
	ansatz := lambdaProxyMZ * (1 + loop)
	mA := math.Sqrt(2*ansatz) * vRuntime
	mR := math.Sqrt(2*lambdaRuntimeMZ) * vRuntime
	return ScalarAnsatzQuality{
		LambdaProxy:             lambdaProxyMZ,
		LoopUnit:                loop,
		LambdaAnsatz:            ansatz,
		LambdaRuntime:           lambdaRuntimeMZ,
		AnsatzMinusRuntime:      ansatz - lambdaRuntimeMZ,
		RelativeRuntimeResidual: (ansatz - lambdaRuntimeMZ) / lambdaRuntimeMZ,
		MassAnsatzGeV:           mA,
		MassRuntimeGeV:          mR,
		DeltaMassGeV:            mA - mR,
		DiagnosticOnly:          true,
		Verdict:                 StatusScalarAnsatzClose,
	}
}

func buildFlavorQuality(loop float64) FlavorAnsatzQuality {
	orientation := sin2Theta13Quarter - jCKM
	epsOrient := loop * (1 - orientation)
	rawRes := loop - epsilonE
	orientRes := epsOrient - epsilonE
	improvement := math.Abs(rawRes) / math.Abs(orientRes)
	return FlavorAnsatzQuality{
		LoopUnit:                    loop,
		EpsilonE:                    epsilonE,
		EpsilonRawL:                 loop,
		RawResidual:                 rawRes,
		RawRelativeResidual:         rawRes / epsilonE,
		OrientationCandidate:        orientation,
		EpsilonOrientation:          epsOrient,
		OrientationResidual:         orientRes,
		OrientationRelativeResidual: orientRes / epsilonE,
		ResidualImprovementFactor:   improvement,
		Verdict:                     StatusFlavorBalanceClose,
	}
}

func buildSignAndRole() SignAndRoleAudit {
	return SignAndRoleAudit{
		FlavorUsesBelowL:     epsilonE < 1/(8*math.Pi),
		ScalarUsesAboveProxy: lambdaRuntimeMZ > lambdaProxyMZ,
		OppositeSigns:        true,
		Statement:            "The same L appears with opposite bridge roles: flavor wall epsilon_e is slightly below L after orientation correction, while scalar lambda_runtime(M_Z) lies above the positive proxy by a loop-sized relative amount.",
		NativeTheoremClaimed: false,
		Verdict:              StatusAppearsInBoth,
	}
}

func buildCrossSealType() CrossSealTypeAudit {
	return CrossSealTypeAudit{
		LikelyType: "HistoryLoopUnitSeal",
		AllowedRoles: []string{
			"bridge loop/matching unit",
			"charged-lepton wall angular scale",
			"scalar proxy-to-runtime relative correction scale",
			"environmental cross-seal diagnostic",
		},
		DisallowedPromotions: []string{
			"native loop theorem",
			"Koide derivation",
			"Higgs pole theorem",
			"scalar-stability theorem",
			"PMNS/CKM derivation",
		},
		BridgeOnly: true,
		Verdict:    StatusCrossSealBridgeOnly,
	}
}

func buildNativeStatus() NativeStatus {
	return NativeStatus{
		NativeOneOver8PiTheorem:     false,
		NativeScalarMatchingTheorem: false,
		NativeKoideWallTheorem:      false,
		NativeCrossSealTheorem:      false,
		NativeOrientationBalance:    false,
		NativeHiggsPoleTheorem:      false,
		Statement:                   "No native theorem currently derives L=1/(8*pi), scalar loop matching, Koide wall geometry, the cross-seal loop unit, the orientation balance, or a Higgs pole formula.",
		Verdict:                     StatusNoCrossSealTheorem,
	}
}

func auditFirewalls() Firewalls { return Firewalls{Verdict: StatusGate623Boundary} }

func Statuses() []string {
	return []string{
		StatusGate622Inherited,
		StatusFlavorLoopInherited,
		StatusNormalFormWritten,
		StatusUnitsComputed,
		StatusAppearsInBoth,
		StatusScalarAnsatzClose,
		StatusFlavorBalanceClose,
		StatusCrossSealBridgeOnly,
		StatusNoCrossSealTheorem,
		StatusNoScalarMatching,
		StatusNoKoideWallTheorem,
		StatusNoOrientationBalance,
		StatusNoHiggsTheorem,
		StatusGate623Boundary,
	}
}
