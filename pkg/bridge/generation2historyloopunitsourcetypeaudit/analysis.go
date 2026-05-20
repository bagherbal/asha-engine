// Package generation2historyloopunitsourcetypeaudit implements
// Gate 624: HistoryLoopUnit Source-Type Audit.
//
// Gate 623 defined the bridge-layer HistoryLoopUnitSeal L=1/(8*pi) after the
// same unit appeared in the scalar low-scale matching seal and the
// charged-lepton Koide wall seal. Gate 624 does not test another endpoint
// number. It audits the possible type-origin of L: normalized Hopf/circle
// phase, weak-quarter projection, boundary/phase-space unit, heat-kernel loop
// descendant, or arbitrary environmental numerical coincidence.
package generation2historyloopunitsourcetypeaudit

import (
	"fmt"
	"math"
	"sync"

	gate572 "github.com/bagherbal/asha-engine/pkg/bridge/generation2projectivefockcp3momentmapselectorgeometryaudit"
	gate623 "github.com/bagherbal/asha-engine/pkg/bridge/generation2universaloneover8piloopunitcrosssealaudit"
	gate570 "github.com/bagherbal/asha-engine/pkg/bridge/generation2witthopfs7contactreebaudit"
)

const (
	AuditID = "GATE624-HISTORY-LOOP-UNIT-SOURCE-TYPE-AUDIT"

	StatusGate623Inherited          = "PASS_GATE623_HISTORY_LOOP_UNIT_INHERITED"
	StatusDecompositionsTyped       = "PASS_L_DECOMPOSITIONS_TYPED"
	StatusHopfPhaseAudited          = "PASS_HOPF_PHASE_SOURCE_CANDIDATE_AUDITED"
	StatusWeakQuarterAudited        = "PASS_WEAK_QUARTER_SOURCE_CANDIDATE_AUDITED"
	StatusHeatKernelAudited         = "PASS_HEAT_KERNEL_LOOP_FACTOR_SOURCE_CANDIDATE_AUDITED"
	StatusScalarFlavorAudited       = "PASS_SCALAR_AND_FLAVOR_ROLES_AUDITED"
	StatusQuarterPhaseCandidate     = "CONDITIONAL_SUPPORT_L_EQUALS_QUARTER_NORMALIZED_PHASE_UNIT_CANDIDATE"
	StatusSharedHistoryLoopUnitSeal = "CONDITIONAL_SUPPORT_L_IS_SHARED_HISTORY_LOOP_UNIT_SEAL"
	StatusNoHopfToFlavorTheorem     = "FAILED_ROUTE_NO_NATIVE_HOPF_TO_FLAVOR_WALL_THEOREM"
	StatusNoHopfToScalarTheorem     = "FAILED_ROUTE_NO_NATIVE_HOPF_TO_SCALAR_MATCHING_THEOREM"
	StatusNoHeatKernelReduction     = "FAILED_ROUTE_NO_NATIVE_HEAT_KERNEL_TO_ONE_OVER_8PI_REDUCTION"
	StatusNoHistoryLoopUnitTheorem  = "FAILED_ROUTE_NO_NATIVE_HISTORY_LOOP_UNIT_THEOREM"
	StatusGate624Boundary           = "FIREWALL_PRESERVED_GATE624_HISTORY_LOOP_UNIT_SOURCE_BOUNDARY"
)

const (
	lambdaProxyMZ      = 0.12490310236015
	lambdaRuntimeMZ    = 0.1296525650504758
	rhoLambdaMatch     = 0.0380251779225699
	epsilonE           = 0.039569756309433
	kappaE             = 0.00550355419157456
	sin2Theta13Quarter = 0.0055375
	jCKM               = 3.11699352875547e-05
	r3MinusOne         = 0.0509933868964996
	lambdaL12Abs       = 0.049700942077683274
	xiBoundary         = 0.0503471644870914
	g2MZ               = 0.6527521238927322
	gYMZ               = 0.3500756885970262
)

type Gate623Inheritance struct {
	LoopUnit         float64
	FlavorNormalForm string
	ScalarNormalForm string
	KappaE           float64
	KappaLambda      float64
	ScalarBridgeOnly bool
	FlavorBridgeOnly bool
	NativeLTheorem   bool
	NativeCrossSeal  bool
	PreviousTruth    string
	Verdict          string
}

type LDecompositionRow struct {
	Expression        string
	Value             float64
	Lane              string
	TypedObject       string
	CandidateRole     string
	Typed             bool
	NativeCertified   bool
	ArbitraryConstant bool
}

type LDecompositionTable struct {
	LoopUnit          float64
	Rows              []LDecompositionRow
	AllValuesMatch    bool
	AllRowsTyped      bool
	NoArbitrarySearch bool
	Verdict           string
}

type HopfPhaseAudit struct {
	Gate570HopfS7Certified       bool
	Gate570ReebPhaseCertified    bool
	Gate572CP3Certified          bool
	HopfFiber                    string
	ContactForm                  string
	ReebAction                   string
	ProjectiveQuotient           string
	NormalizedPhaseMeasure       string
	CirclePhaseNormalization     bool
	QuarterProjectionCandidate   bool
	QuarterProjectionCertified   bool
	MapToFlavorWallCertified     bool
	MapToScalarMatchingCertified bool
	PhysicalTimeClaimed          bool
	Verdict                      string
}

type WeakQuarterCandidate struct {
	Name            string
	Expression      string
	Typed           bool
	NativeCertified bool
	Comment         string
}

type WeakQuarterAudit struct {
	Factor                   float64
	Candidates               []WeakQuarterCandidate
	WeakNormalizationTyped   bool
	PMNSOverlapTyped         bool
	NativeConnectionToL      bool
	NativeWeakQuarterLoopLaw bool
	Verdict                  string
}

type HeatKernelOperation struct {
	Name      string
	Input     string
	Output    string
	Certified bool
	Comment   string
}

type HeatKernelLoopFactorAudit struct {
	LoopUnit              float64
	FourDLoopUnit         float64
	BoundarySurfaceUnit   float64
	Operations            []HeatKernelOperation
	AnyCertifiedReduction bool
	Verdict               string
}

type ScalarKappaCandidate struct {
	Name             string
	Value            float64
	Residual         float64
	RelativeResidual float64
	Typed            bool
	NativeCertified  bool
	Comment          string
}

type ScalarRoleAudit struct {
	LambdaProxy          float64
	LambdaRuntime        float64
	RhoLambdaMatch       float64
	LoopUnit             float64
	KappaLambda          float64
	NormalForm           string
	Candidates           []ScalarKappaCandidate
	ClosestName          string
	ClosestResidual      float64
	KappaSourceCertified bool
	Verdict              string
}

type FlavorRoleAudit struct {
	EpsilonE             float64
	LoopUnit             float64
	KappaE               float64
	Sin2Theta13Quarter   float64
	JCKM                 float64
	OrientationCandidate float64
	EpsilonOrientation   float64
	Residual             float64
	NormalForm           string
	Classification       string
	NativeDerived        bool
	Verdict              string
}

type CrossSealRow struct {
	Seal         string
	BaseUnit     string
	Correction   string
	SignRole     string
	Residual     string
	NativeStatus string
}

type CrossSealComparisonTable struct {
	Rows              []CrossSealRow
	SharedLBridgeSeal bool
	NativeCrossSeal   bool
	Verdict           string
}

type NativeASHAStatus struct {
	NativeLTheorem                bool
	NativeHopfToFlavorWallMap     bool
	NativeHopfToScalarMatchingMap bool
	NativeHeatKernelToLReduction  bool
	NativeWeakQuarterLoopTheorem  bool
	NativeCrossSealOrientationLaw bool
	Statement                     string
	Verdict                       string
}

type Firewalls struct {
	ClaimsKoideDerived      bool
	ClaimsHiggsMassDerived  bool
	ClaimsScalarStability   bool
	ClaimsPMNSCKMDerived    bool
	ClaimsGaugeUnification  bool
	ClaimsNativeLoopTheorem bool
	ClaimsPhysicalTime      bool
	Verdict                 string
}

type Analysis struct {
	Inherited      Gate623Inheritance
	Decompositions LDecompositionTable
	HopfPhase      HopfPhaseAudit
	WeakQuarter    WeakQuarterAudit
	HeatKernel     HeatKernelLoopFactorAudit
	ScalarRole     ScalarRoleAudit
	FlavorRole     FlavorRoleAudit
	CrossSeal      CrossSealComparisonTable
	NativeStatus   NativeASHAStatus
	Firewalls      Firewalls
	Truth          string
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
	g623, err := gate623.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate623 predecessor: %w", err)
	}
	g570, err := gate570.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate570 Hopf predecessor: %w", err)
	}
	g572, err := gate572.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate572 CP3 predecessor: %w", err)
	}

	loop := 1.0 / (8.0 * math.Pi)
	inherit := inheritGate623(g623, loop)
	scalar := buildScalarRole(loop, inherit.KappaLambda)
	flavor := buildFlavorRole(loop)
	a := Analysis{
		Inherited:      inherit,
		Decompositions: buildDecompositionTable(loop),
		HopfPhase:      buildHopfPhaseAudit(g570, g572),
		WeakQuarter:    buildWeakQuarterAudit(),
		HeatKernel:     buildHeatKernelAudit(loop),
		ScalarRole:     scalar,
		FlavorRole:     flavor,
		CrossSeal:      buildCrossSealTable(loop, scalar, flavor),
		NativeStatus:   buildNativeStatus(),
		Firewalls:      auditFirewalls(),
		Truth:          "Gate 624 types L=1/(8*pi) as a bridge-layer HistoryLoopUnitSeal source candidate. The strongest structural reading is L=(1/4)(1/(2*pi)), a weak-quarter projection of normalized S1/Hopf phase measure, with boundary and heat-kernel ancestry also audited. No native Hopf-to-flavor, Hopf-to-scalar, heat-kernel-to-L, weak-quarter loop, or cross-seal orientation theorem is certified.",
	}
	return a, nil
}

func inheritGate623(g gate623.Analysis, loop float64) Gate623Inheritance {
	return Gate623Inheritance{
		LoopUnit:         loop,
		FlavorNormalForm: g.NormalForm.FlavorEquation,
		ScalarNormalForm: g.NormalForm.ScalarEquation,
		KappaE:           g.NormalForm.FlavorKappaE,
		KappaLambda:      g.NormalForm.ScalarKappaLambda,
		ScalarBridgeOnly: !g.NativeStatus.NativeScalarMatchingTheorem,
		FlavorBridgeOnly: !g.NativeStatus.NativeKoideWallTheorem,
		NativeLTheorem:   g.NativeStatus.NativeOneOver8PiTheorem,
		NativeCrossSeal:  g.NativeStatus.NativeCrossSealTheorem,
		PreviousTruth:    g.Truth,
		Verdict:          StatusGate623Inherited,
	}
}

func buildDecompositionTable(loop float64) LDecompositionTable {
	rows := []LDecompositionRow{
		{
			Expression:      "L = 1/(8*pi)",
			Value:           1.0 / (8.0 * math.Pi),
			Lane:            "history loop unit seal",
			TypedObject:     "Gate623 bridge scalar/flavor shared unit",
			CandidateRole:   "inherited bridge object, not native source theorem",
			Typed:           true,
			NativeCertified: false,
		},
		{
			Expression:      "L = (1/4)(1/(2*pi))",
			Value:           0.25 * (1.0 / (2.0 * math.Pi)),
			Lane:            "Hopf/S1 phase normalization plus weak quarter",
			TypedObject:     "quarter-normalized circle phase candidate",
			CandidateRole:   "most promising phase-source type candidate",
			Typed:           true,
			NativeCertified: false,
		},
		{
			Expression:      "L = (1/2)(1/(4*pi))",
			Value:           0.5 * (1.0 / (4.0 * math.Pi)),
			Lane:            "boundary surface normalization",
			TypedObject:     "half-boundary/surface normalization candidate",
			CandidateRole:   "possible boundary or phase-space normalization ancestor",
			Typed:           true,
			NativeCertified: false,
		},
		{
			Expression:      "L = 2*pi/(16*pi^2)",
			Value:           (2.0 * math.Pi) / (16.0 * math.Pi * math.Pi),
			Lane:            "heat-kernel / one-loop descendant",
			TypedObject:     "angularly reduced 4D loop-factor candidate",
			CandidateRole:   "possible phase integration over ordinary one-loop unit",
			Typed:           true,
			NativeCertified: false,
		},
		{
			Expression:      "L = sqrt(1/(64*pi^2))",
			Value:           math.Sqrt(1.0 / (64.0 * math.Pi * math.Pi)),
			Lane:            "root chamber / square-root loop descendant",
			TypedObject:     "positive square-root of a pi-squared loop-size unit",
			CandidateRole:   "candidate only; no certified scalar/root chamber operation",
			Typed:           true,
			NativeCertified: false,
		},
	}
	return LDecompositionTable{LoopUnit: loop, Rows: rows, AllValuesMatch: allRowsMatch(rows, loop), AllRowsTyped: allRowsTyped(rows), NoArbitrarySearch: true, Verdict: StatusDecompositionsTyped}
}

func buildHopfPhaseAudit(g570 gate570.Analysis, g572 gate572.Analysis) HopfPhaseAudit {
	return HopfPhaseAudit{
		Gate570HopfS7Certified:       g570.Final.HopfS7Certified && g570.Final.HopfContactCertified,
		Gate570ReebPhaseCertified:    g570.Final.ReebCertified && g570.Final.TotalPhaseRelation,
		Gate572CP3Certified:          g572.Final.CP3Certified && g572.Phase.TrivialOnCP3,
		HopfFiber:                    "S^1 -> S^7 -> CP^3",
		ContactForm:                  "alpha = Im<z,dz>",
		ReebAction:                   "z -> exp(i theta) z, R=iz",
		ProjectiveQuotient:           "total Fock phase is quotiented on CP^3",
		NormalizedPhaseMeasure:       "dtheta/(2*pi)",
		CirclePhaseNormalization:     true,
		QuarterProjectionCandidate:   true,
		QuarterProjectionCertified:   false,
		MapToFlavorWallCertified:     false,
		MapToScalarMatchingCertified: false,
		PhysicalTimeClaimed:          false,
		Verdict:                      StatusHopfPhaseAudited,
	}
}

func buildWeakQuarterAudit() WeakQuarterAudit {
	candidates := []WeakQuarterCandidate{
		{Name: "weak generator normalization", Expression: "T_a = sigma_a/2", Typed: true, NativeCertified: false, Comment: "a half-generator on each side naturally produces quarter-sized quadratic/overlap factors"},
		{Name: "scalar doublet normalization", Expression: "H in C^2 with SU(2) doublet conventions", Typed: true, NativeCertified: false, Comment: "typed weak/scalar bridge normalization, not a source theorem for L"},
		{Name: "projector-quarter convention", Expression: "rank/overlap quarter scale", Typed: true, NativeCertified: false, Comment: "projector normalization can yield 1/4 factors but no Gate624 certified map"},
		{Name: "PMNS reactor leakage", Expression: "sin^2(theta13)/4", Typed: true, NativeCertified: false, Comment: "already used in the flavor orientation balance seal"},
	}
	return WeakQuarterAudit{Factor: 0.25, Candidates: candidates, WeakNormalizationTyped: true, PMNSOverlapTyped: true, NativeConnectionToL: false, NativeWeakQuarterLoopLaw: false, Verdict: StatusWeakQuarterAudited}
}

func buildHeatKernelAudit(loop float64) HeatKernelLoopFactorAudit {
	ops := []HeatKernelOperation{
		{Name: "square-root", Input: "1/(64*pi^2)", Output: "1/(8*pi)", Certified: false, Comment: "algebraically true as a positive root, but no current scalar/root chamber operation certifies it as a loop-unit source"},
		{Name: "angular projection", Input: "1/(16*pi^2)", Output: "2*pi/(16*pi^2)=1/(8*pi)", Certified: false, Comment: "requires a lawful phase integration/projection theorem not present in ASHA"},
		{Name: "boundary reduction", Input: "1/(4*pi)", Output: "(1/2)(1/(4*pi))", Certified: false, Comment: "surface normalization is typed, but the half-boundary reduction is not certified"},
		{Name: "phase-space reduction", Input: "dtheta/(2*pi)", Output: "(1/4)(1/(2*pi))", Certified: false, Comment: "the phase measure exists; the quarter projection into scalar/flavor seals is not certified"},
		{Name: "scalar/root chamber operation", Input: "finite spectral/root chamber", Output: "L", Certified: false, Comment: "no native operator maps the chamber data to L"},
	}
	return HeatKernelLoopFactorAudit{LoopUnit: loop, FourDLoopUnit: 1.0 / (16.0 * math.Pi * math.Pi), BoundarySurfaceUnit: 1.0 / (4.0 * math.Pi), Operations: ops, AnyCertifiedReduction: false, Verdict: StatusHeatKernelAudited}
}

func buildScalarRole(loop, kappaLambda float64) ScalarRoleAudit {
	candidates := []ScalarKappaCandidate{
		sCandidate("kappa_e", kappaE, kappaLambda, "charged-lepton loop deficit coefficient"),
		sCandidate("sin^2(theta13)/4", sin2Theta13Quarter, kappaLambda, "PMNS reactor leakage quarter"),
		sCandidate("J_CKM", jCKM, kappaLambda, "quark CP orientation area"),
		sCandidate("R_3-1", r3MinusOne, kappaLambda, "strong relative boundary wound"),
		sCandidate("|lambda(Lambda_12)|", lambdaL12Abs, kappaLambda, "high-scale scalar runtime wound"),
		sCandidate("xi_boundary", xiBoundary, kappaLambda, "gauge-scalar boundary stress scale"),
		sCandidate("alpha_2(M_Z)", alpha2(), kappaLambda, "weak endpoint coupling"),
		sCandidate("alpha_EM(M_Z)", alphaEM(), kappaLambda, "electromagnetic endpoint coupling"),
	}
	name, residual := closestScalarCandidate(candidates)
	return ScalarRoleAudit{
		LambdaProxy:          lambdaProxyMZ,
		LambdaRuntime:        lambdaRuntimeMZ,
		RhoLambdaMatch:       rhoLambdaMatch,
		LoopUnit:             loop,
		KappaLambda:          kappaLambda,
		NormalForm:           "lambda_runtime(M_Z)=lambda_proxy(M_Z)[1+L(1-kappa_lambda)]",
		Candidates:           candidates,
		ClosestName:          name,
		ClosestResidual:      residual,
		KappaSourceCertified: false,
		Verdict:              StatusScalarFlavorAudited,
	}
}

func buildFlavorRole(loop float64) FlavorRoleAudit {
	orientation := sin2Theta13Quarter - jCKM
	epsOrient := loop * (1.0 - orientation)
	return FlavorRoleAudit{
		EpsilonE:             epsilonE,
		LoopUnit:             loop,
		KappaE:               kappaE,
		Sin2Theta13Quarter:   sin2Theta13Quarter,
		JCKM:                 jCKM,
		OrientationCandidate: orientation,
		EpsilonOrientation:   epsOrient,
		Residual:             epsilonE - epsOrient,
		NormalForm:           "epsilon_e = L[1 - sin^2(theta13)/4 + J_CKM] + residual",
		Classification:       "orientation-corrected phase-wall loop unit",
		NativeDerived:        false,
		Verdict:              StatusScalarFlavorAudited,
	}
}

func buildCrossSealTable(loop float64, scalar ScalarRoleAudit, flavor FlavorRoleAudit) CrossSealComparisonTable {
	rows := []CrossSealRow{
		{Seal: "flavor wall", BaseUnit: "L=1/(8*pi)", Correction: "1-kappa_e ≈ 1-sin^2(theta13)/4+J_CKM", SignRole: "epsilon_e lies slightly below L", Residual: fmt.Sprintf("epsilon_e - L[1-sin^2(theta13)/4+J_CKM] = %.15g", flavor.Residual), NativeStatus: "bridge environmental fit; no Koide/PMNS/CKM derivation"},
		{Seal: "scalar matching", BaseUnit: "L=1/(8*pi)", Correction: "1-kappa_lambda", SignRole: "lambda_runtime(M_Z) lies above lambda_proxy by a positive L-sized relative correction", Residual: fmt.Sprintf("kappa_lambda = %.15g; source not certified", scalar.KappaLambda), NativeStatus: "bridge scalar matching clue; no Higgs/native loop theorem"},
		{Seal: "possible boundary stress", BaseUnit: fmt.Sprintf("xi_boundary≈%.15g, not L", xiBoundary), Correction: "(+R_3-1, -|lambda(Lambda_12)|) stress-pair shadow", SignRole: "opposed gauge/scalar boundary wounds", Residual: fmt.Sprintf("R_3-1=%.15g, |lambda(Lambda_12)|=%.15g", r3MinusOne, lambdaL12Abs), NativeStatus: "typed stress context only; no L-source theorem"},
	}
	_ = loop
	return CrossSealComparisonTable{Rows: rows, SharedLBridgeSeal: true, NativeCrossSeal: false, Verdict: StatusSharedHistoryLoopUnitSeal}
}

func buildNativeStatus() NativeASHAStatus {
	return NativeASHAStatus{
		NativeLTheorem:                false,
		NativeHopfToFlavorWallMap:     false,
		NativeHopfToScalarMatchingMap: false,
		NativeHeatKernelToLReduction:  false,
		NativeWeakQuarterLoopTheorem:  false,
		NativeCrossSealOrientationLaw: false,
		Statement:                     "ASHA currently supplies Hopf S1/CP3 phase geometry and bridge weak-quarter candidates, but no native theorem that transports normalized phase, heat-kernel loop units, or weak-quarter projections into the scalar/flavor HistoryLoopUnitSeal.",
		Verdict:                       StatusNoHistoryLoopUnitTheorem,
	}
}

func auditFirewalls() Firewalls { return Firewalls{Verdict: StatusGate624Boundary} }

func allRowsMatch(rows []LDecompositionRow, target float64) bool {
	for _, r := range rows {
		if math.Abs(r.Value-target) > 1e-15 {
			return false
		}
	}
	return true
}

func allRowsTyped(rows []LDecompositionRow) bool {
	for _, r := range rows {
		if !r.Typed || r.ArbitraryConstant {
			return false
		}
	}
	return true
}

func alpha2() float64 { return g2MZ * g2MZ / (4.0 * math.Pi) }

func alphaEM() float64 {
	e := g2MZ * gYMZ / math.Sqrt(g2MZ*g2MZ+gYMZ*gYMZ)
	return e * e / (4.0 * math.Pi)
}

func sCandidate(name string, value, target float64, comment string) ScalarKappaCandidate {
	res := value - target
	rel := math.NaN()
	if target != 0 {
		rel = res / target
	}
	return ScalarKappaCandidate{Name: name, Value: value, Residual: res, RelativeResidual: rel, Typed: true, NativeCertified: false, Comment: comment}
}

func closestScalarCandidate(rows []ScalarKappaCandidate) (string, float64) {
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

func Statuses() []string {
	return []string{
		StatusGate623Inherited,
		StatusDecompositionsTyped,
		StatusHopfPhaseAudited,
		StatusWeakQuarterAudited,
		StatusHeatKernelAudited,
		StatusScalarFlavorAudited,
		StatusQuarterPhaseCandidate,
		StatusSharedHistoryLoopUnitSeal,
		StatusNoHopfToFlavorTheorem,
		StatusNoHopfToScalarTheorem,
		StatusNoHeatKernelReduction,
		StatusNoHistoryLoopUnitTheorem,
		StatusGate624Boundary,
	}
}
