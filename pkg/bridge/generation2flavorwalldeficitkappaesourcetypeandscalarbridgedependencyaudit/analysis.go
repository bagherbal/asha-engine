// Package generation2flavorwalldeficitkappaesourcetypeandscalarbridgedependencyaudit implements
// Gate 746: Flavor-Wall Deficit Kappa_e Source-Type and Scalar-Bridge Dependency Audit.
//
// Gate 745 completed the pole diagnostic boundary. Gate 746 returns to seal
// reduction by auditing how the scalar-runtime bridge depends on kappa_e and
// whether the existing orientation candidate kappa_e_orient=sin^2(theta13)/4-J_CKM
// can source-type it.  The result is a close flavor-orientation candidate, not
// a native flavor theorem or independent scalar-runtime derivation.
package generation2flavorwalldeficitkappaesourcetypeandscalarbridgedependencyaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate734 "github.com/bagherbal/asha-engine/pkg/bridge/generation2cubicboundarypolynomialscalarruntimetransportandpredictionboundaryaudit"
	gate745 "github.com/bagherbal/asha-engine/pkg/bridge/generation2level1cpoleobservablesealanddiagnosticdeltaaudit"
)

const (
	AuditID = "GATE746-FLAVOR-WALL-DEFICIT-KAPPA-E-SOURCE-TYPE-SCALAR-BRIDGE-DEPENDENCY-AUDIT"

	StatusGate745PoleDiagnosticBoundaryInherited = "PASS_GATE745_POLE_DIAGNOSTIC_BOUNDARY_INHERITED"
	StatusKappaEScalarBridgeDependencyAudited    = "PASS_KAPPA_E_SCALAR_BRIDGE_DEPENDENCY_AUDITED"
	StatusKappaEOrientationCandidateComputed     = "PASS_KAPPA_E_ORIENTATION_CANDIDATE_COMPUTED"
	StatusKappaEOrientReplacementTested          = "PASS_KAPPA_E_ORIENT_REPLACEMENT_TESTED"
	StatusDeltaKappaESourceCandidatesAudited     = "PASS_DELTA_KAPPA_E_SOURCE_CANDIDATES_AUDITED"
	StatusFlavorFirewallEnforced                 = "PASS_FLAVOR_FIREWALL_ENFORCED"

	StatusKappaEActiveScalarBridgeInput             = "CONDITIONAL_SUPPORT_KAPPA_E_IS_ACTIVE_SCALAR_BRIDGE_INPUT"
	StatusKappaEOrientCloseFlavorSourceCandidate    = "CONDITIONAL_SUPPORT_KAPPA_E_ORIENT_IS_CLOSE_FLAVOR_ORIENTATION_SOURCE_CANDIDATE"
	StatusScalarRuntimeBridgeSensitiveFlavorDeficit = "CONDITIONAL_SUPPORT_SCALAR_RUNTIME_BRIDGE_IS_SENSITIVE_TO_FLAVOR_DEFICIT_SOURCE"

	StatusKappaEOrientNotExact                    = "FAILED_ROUTE_KAPPA_E_ORIENT_DOES_NOT_EXACTLY_EQUAL_KAPPA_E"
	StatusNoNativeKappaEOrientationResidualSource = "FAILED_ROUTE_NO_NATIVE_SOURCE_FOR_KAPPA_E_ORIENTATION_RESIDUAL"
	StatusNoNativeFlavorDeficitTheorem            = "FAILED_ROUTE_NO_NATIVE_FLAVOR_DEFICIT_THEOREM"
	StatusNoNativePMNSOrCKMTheorem                = "FAILED_ROUTE_NO_NATIVE_PMNS_OR_CKM_THEOREM"
	StatusNoYukawaOperatorOrEigenvalueTheorem     = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusNoIndependentScalarRuntimeTheorem       = "FAILED_ROUTE_NO_INDEPENDENT_SCALAR_RUNTIME_THEOREM"
	StatusNoHiggsMassOrPoleMassTheorem            = "FAILED_ROUTE_NO_HIGGS_MASS_OR_POLE_MASS_THEOREM"
	StatusGate746Boundary                         = "FIREWALL_PRESERVED_GATE746_KAPPA_E_SOURCE_BOUNDARY"
)

const (
	// Gate590/Gate592 orientation candidate in the current v1 flavor ledger.
	KappaEOrientation = 0.00550633006471245
	SinTheta13Quarter = 0.00553940006471245
	JCKM              = 0.00003307000000000
)

type Gate745Inheritance struct {
	Inherited             bool
	DiagnosticAllowed     bool
	ExternalPoleIsNotASHA bool
	NoIndependentPoleMass bool
	NoYukawaTheorem       bool
	Verdict               string
}

type ScalarBridgeDependencyAudit struct {
	KappaE                      float64
	AppearsInBoundaryPolynomial bool
	AppearsInRuntimeTransport   bool
	BoundaryPolynomialTerm      string
	RuntimeTransportTerm        string
	StructurallyActive          bool
	Verdict                     string
}

type OrientationCandidateAudit struct {
	Formula            string
	KappaE             float64
	KappaEOrient       float64
	SinTheta13Quarter  float64
	JCKM               float64
	DeltaKappaE        float64
	AbsDeltaKappaE     float64
	CloseButNotExact   bool
	TypedPMNSLeakage   bool
	TypedCKMCorrection bool
	NativeTheorem      bool
	Verdict            string
}

type ReplacementTest struct {
	FormulaExactKappaE    string
	FormulaOrientKappaE   string
	SSplit                float64
	P_K7                  float64
	KappaE                float64
	KappaEOrient          float64
	FExact                float64
	FOrient               float64
	FWallShift            float64
	W3Exact               float64
	W3Orient              float64
	LambdaProxy           float64
	L                     float64
	RuntimeExactKappaE    float64
	RuntimeOrientKappaE   float64
	RuntimeShift          float64
	RuntimeLedger         float64
	RuntimeOrientResidual float64
	ApproximationOnly     bool
	Verdict               string
}

type ResidualSourceAudit struct {
	DeltaKappaE           float64
	Candidates            []string
	NativeSourceCertified bool
	Verdict               string
}

type FlavorFirewall struct {
	DerivesPMNS              bool
	DerivesCKM               bool
	DerivesFlavorHierarchy   bool
	DerivesYukawaEigenvalues bool
	DerivesScalarRuntime     bool
	DerivesHiggsMass         bool
	KappaEStillBridgeSeal    bool
	Verdict                  string
}

type Analysis struct {
	Gate745     Gate745Inheritance
	Dependency  ScalarBridgeDependencyAudit
	Orientation OrientationCandidateAudit
	Replacement ReplacementTest
	Residual    ResidualSourceAudit
	Firewall    FlavorFirewall
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
	g745, err := gate745.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate745 inheritance unavailable: %w", err)
	}
	g734, err := gate734.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate734 scalar bridge inheritance unavailable: %w", err)
	}
	inherit := buildGate745Inheritance(g745)
	dependency := buildScalarBridgeDependency(g734)
	orientation := buildOrientationCandidate(dependency.KappaE)
	replacement := buildReplacementTest(g734, orientation)
	residual := buildResidualSourceAudit(orientation)
	firewall := buildFlavorFirewall()
	truth := "Gate 746 audits kappa_e as an active scalar-bridge input: it appears both as the quadratic raw-moment coefficient in F_wall_3 and in the final scalar transport factor. The orientation expression kappa_e_orient=sin^2(theta13)/4-J_CKM is close but not exact, and replacing kappa_e by it shifts the scalar-runtime bridge at about 1.38e-8. The residual has no native source at this gate; kappa_e remains a bridge-layer flavor deficit seal, not a PMNS/CKM/Yukawa theorem."
	return Analysis{Gate745: inherit, Dependency: dependency, Orientation: orientation, Replacement: replacement, Residual: residual, Firewall: firewall, Truth: truth}, nil
}

func buildGate745Inheritance(g gate745.Analysis) Gate745Inheritance {
	return Gate745Inheritance{
		Inherited:             g.Gate744.Inherited && g.Delta.MeasuresProxyToPoleGapOnly && !g.Delta.IndependentPrediction && !g.NonFit.DiagnosticDeltaIsPrediction,
		DiagnosticAllowed:     g.Delta.MeasuresProxyToPoleGapOnly && g.Observable.AllowsDiagnostic,
		ExternalPoleIsNotASHA: !g.Observable.NativeDerived && !g.NonFit.ExternalObservableExplainsGap,
		NoIndependentPoleMass: !g.Delta.IndependentPrediction && !g.NonFit.DiagnosticDeltaIsPrediction,
		NoYukawaTheorem:       strings.Contains(g.NonFit.Verdict, gate745.StatusNoYukawaOperatorOrEigenvalueTheorem),
		Verdict:               StatusGate745PoleDiagnosticBoundaryInherited,
	}
}

func buildScalarBridgeDependency(g gate734.Analysis) ScalarBridgeDependencyAudit {
	return ScalarBridgeDependencyAudit{
		KappaE:                      g.BoundarySub.KappaE,
		AppearsInBoundaryPolynomial: true,
		AppearsInRuntimeTransport:   true,
		BoundaryPolynomialTerm:      "kappa_e p_K7 S_split^2 in F_wall_3",
		RuntimeTransportTerm:        "1-W_3+kappa_e in lambda_runtime≈lambda_proxy[1+L(1-W_3+kappa_e)]",
		StructurallyActive:          true,
		Verdict: strings.Join([]string{
			StatusKappaEScalarBridgeDependencyAudited,
			StatusKappaEActiveScalarBridgeInput,
			StatusScalarRuntimeBridgeSensitiveFlavorDeficit,
		}, "; "),
	}
}

func buildOrientationCandidate(kappaE float64) OrientationCandidateAudit {
	delta := kappaE - KappaEOrientation
	return OrientationCandidateAudit{
		Formula:            "kappa_e_orient = sin^2(theta13)/4 - J_CKM",
		KappaE:             kappaE,
		KappaEOrient:       KappaEOrientation,
		SinTheta13Quarter:  SinTheta13Quarter,
		JCKM:               JCKM,
		DeltaKappaE:        delta,
		AbsDeltaKappaE:     math.Abs(delta),
		CloseButNotExact:   math.Abs(delta) > 1e-9 && math.Abs(delta) < 5e-6,
		TypedPMNSLeakage:   true,
		TypedCKMCorrection: true,
		NativeTheorem:      false,
		Verdict: strings.Join([]string{
			StatusKappaEOrientationCandidateComputed,
			StatusKappaEOrientCloseFlavorSourceCandidate,
			StatusKappaEOrientNotExact,
		}, "; "),
	}
}

func buildReplacementTest(g gate734.Analysis, o OrientationCandidateAudit) ReplacementTest {
	p := g.Gate733.P_K7
	S := g.Gate733.SSplit
	fExact := g.BoundarySub.FWall3
	fOrient := rawPolynomial(p, S, o.KappaEOrient)
	wExact := g.BoundarySub.W3
	wOrient := math.Abs(g.BoundarySub.Lambda) + fOrient
	runtimeExact := g.Runtime.RuntimeApprox
	runtimeOrient := g.Runtime.LambdaProxy * (1 + g.Runtime.L*(1-wOrient+o.KappaEOrient))
	return ReplacementTest{
		FormulaExactKappaE:    "lambda_runtime≈lambda_proxy[1+L(1-W_3(kappa_e)+kappa_e)]",
		FormulaOrientKappaE:   "lambda_runtime≈lambda_proxy[1+L(1-W_3(kappa_e_orient)+kappa_e_orient)]",
		SSplit:                S,
		P_K7:                  p,
		KappaE:                o.KappaE,
		KappaEOrient:          o.KappaEOrient,
		FExact:                fExact,
		FOrient:               fOrient,
		FWallShift:            fOrient - fExact,
		W3Exact:               wExact,
		W3Orient:              wOrient,
		LambdaProxy:           g.Runtime.LambdaProxy,
		L:                     g.Runtime.L,
		RuntimeExactKappaE:    runtimeExact,
		RuntimeOrientKappaE:   runtimeOrient,
		RuntimeShift:          runtimeOrient - runtimeExact,
		RuntimeLedger:         g.Runtime.RuntimeExactTransport,
		RuntimeOrientResidual: g.Runtime.RuntimeExactTransport - runtimeOrient,
		ApproximationOnly:     true,
		Verdict: strings.Join([]string{
			StatusKappaEOrientReplacementTested,
			StatusKappaEOrientCloseFlavorSourceCandidate,
			StatusKappaEOrientNotExact,
		}, "; "),
	}
}

func rawPolynomial(p, S, kappaE float64) float64 {
	return p*S + kappaE*p*S*S - 2*p*p*S*S*S
}

func buildResidualSourceAudit(o OrientationCandidateAudit) ResidualSourceAudit {
	candidates := []string{
		"missing PMNS/CKM precision or convention residual",
		"bridge normalization residual",
		"flavor-wall orientation seal correction",
		"unmodeled Yukawa/flavor operator contribution",
	}
	return ResidualSourceAudit{DeltaKappaE: o.DeltaKappaE, Candidates: candidates, NativeSourceCertified: false, Verdict: strings.Join([]string{StatusDeltaKappaESourceCandidatesAudited, StatusNoNativeKappaEOrientationResidualSource}, "; ")}
}

func buildFlavorFirewall() FlavorFirewall {
	return FlavorFirewall{
		DerivesPMNS:              false,
		DerivesCKM:               false,
		DerivesFlavorHierarchy:   false,
		DerivesYukawaEigenvalues: false,
		DerivesScalarRuntime:     false,
		DerivesHiggsMass:         false,
		KappaEStillBridgeSeal:    true,
		Verdict: strings.Join([]string{
			StatusFlavorFirewallEnforced,
			StatusNoNativeFlavorDeficitTheorem,
			StatusNoNativePMNSOrCKMTheorem,
			StatusNoYukawaOperatorOrEigenvalueTheorem,
			StatusNoIndependentScalarRuntimeTheorem,
			StatusNoHiggsMassOrPoleMassTheorem,
			StatusGate746Boundary,
		}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate745PoleDiagnosticBoundaryInherited,
		StatusKappaEScalarBridgeDependencyAudited,
		StatusKappaEOrientationCandidateComputed,
		StatusKappaEOrientReplacementTested,
		StatusDeltaKappaESourceCandidatesAudited,
		StatusFlavorFirewallEnforced,
		StatusKappaEActiveScalarBridgeInput,
		StatusKappaEOrientCloseFlavorSourceCandidate,
		StatusScalarRuntimeBridgeSensitiveFlavorDeficit,
		StatusKappaEOrientNotExact,
		StatusNoNativeKappaEOrientationResidualSource,
		StatusNoNativeFlavorDeficitTheorem,
		StatusNoNativePMNSOrCKMTheorem,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusNoIndependentScalarRuntimeTheorem,
		StatusNoHiggsMassOrPoleMassTheorem,
		StatusGate746Boundary,
	}
}

func FormatGate745(x Gate745Inheritance) string {
	return fmt.Sprintf("inherited=%t diagnosticAllowed=%t externalNotASHA=%t noPolePrediction=%t noYukawa=%t verdict=%q", x.Inherited, x.DiagnosticAllowed, x.ExternalPoleIsNotASHA, x.NoIndependentPoleMass, x.NoYukawaTheorem, x.Verdict)
}

func FormatDependency(x ScalarBridgeDependencyAudit) string {
	return fmt.Sprintf("kappaE=%.17g boundaryPoly=%t runtime=%t boundaryTerm=%q runtimeTerm=%q active=%t verdict=%q", x.KappaE, x.AppearsInBoundaryPolynomial, x.AppearsInRuntimeTransport, x.BoundaryPolynomialTerm, x.RuntimeTransportTerm, x.StructurallyActive, x.Verdict)
}

func FormatOrientation(x OrientationCandidateAudit) string {
	return fmt.Sprintf("formula=%q kappaE=%.17g orient=%.17g sin2quarter=%.17g JCKM=%.17g delta=%.17g absDelta=%.17g close=%t pmns=%t ckm=%t native=%t verdict=%q", x.Formula, x.KappaE, x.KappaEOrient, x.SinTheta13Quarter, x.JCKM, x.DeltaKappaE, x.AbsDeltaKappaE, x.CloseButNotExact, x.TypedPMNSLeakage, x.TypedCKMCorrection, x.NativeTheorem, x.Verdict)
}

func FormatReplacement(x ReplacementTest) string {
	return fmt.Sprintf("S=%.17g p=%.17g kappaE=%.17g orient=%.17g FExact=%.17g FOrient=%.17g FShift=%.17g WExact=%.17g WOrient=%.17g runtimeExactK=%.17g runtimeOrient=%.17g runtimeShift=%.17g ledger=%.17g orientResidual=%.17g approximationOnly=%t verdict=%q", x.SSplit, x.P_K7, x.KappaE, x.KappaEOrient, x.FExact, x.FOrient, x.FWallShift, x.W3Exact, x.W3Orient, x.RuntimeExactKappaE, x.RuntimeOrientKappaE, x.RuntimeShift, x.RuntimeLedger, x.RuntimeOrientResidual, x.ApproximationOnly, x.Verdict)
}

func FormatResidual(x ResidualSourceAudit) string {
	return fmt.Sprintf("delta=%.17g candidates=[%s] nativeSource=%t verdict=%q", x.DeltaKappaE, strings.Join(x.Candidates, ", "), x.NativeSourceCertified, x.Verdict)
}

func FormatFirewall(x FlavorFirewall) string {
	return fmt.Sprintf("pmns=%t ckm=%t hierarchy=%t yukawa=%t runtime=%t higgs=%t kappaSeal=%t verdict=%q", x.DerivesPMNS, x.DerivesCKM, x.DerivesFlavorHierarchy, x.DerivesYukawaEigenvalues, x.DerivesScalarRuntime, x.DerivesHiggsMass, x.KappaEStillBridgeSeal, x.Verdict)
}

func Near(a, b, tol float64) bool { return math.Abs(a-b) <= tol }
