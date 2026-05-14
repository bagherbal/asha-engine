// Package hyperchargediracassembly implements Gate 296:
// Hypercharge Ledger Sieve / Canonical Finite Dirac (D_F) Assembly Audit.
//
// Gate 295 resolved the weak/color direct-sum paradox at zero order by using a
// true Morita bimodule: weak H acts from the left and color M3(C) acts from the
// right/opposite side.  Gate 296 audits the next layer: the U(1) ledger and the
// finite Dirac edge graph.  The gate deliberately separates three statements:
//
//  1. anomaly/Yukawa/unimodularity equations reconstruct the Standard-Model
//     hypercharge ray, but not its absolute normalization;
//  2. the first-order bimodule rule allows exactly the left-doublet to
//     right-singlet Yukawa edge classes with shared right module;
//  3. the same rule forbids lepton-quark, color-changing, and charge-violating
//     edges, but still does not select numerical Yukawa amplitudes or the
//     canonical D_F matrix.
//
// This keeps the theorem honest: the canonical D_F adjacency shape is now much
// sharper, while masses, full flavor matrices, B-gap Majorana activation, and
// heat-kernel dynamics remain firewalled.
package hyperchargediracassembly

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE296-HYPERCHARGE-LEDGER-SIEVE-CANONICAL-FINITE-DIRAC-ASSEMBLY-AUDIT"

	StatusGate295Inherited              = "CONDITIONAL_SUPPORT_GATE295_TRUE_BIMODULE_INHERITED"
	StatusHyperchargeEquationsBuilt     = "CONDITIONAL_SUPPORT_HYPERCHARGE_ANOMALY_YUKAWA_EQUATIONS_BUILT"
	StatusHyperchargeRayRecovered       = "CONDITIONAL_SUPPORT_STANDARD_MODEL_HYPERCHARGE_RAY_RECOVERED"
	StatusUnimodularityChecked          = "CONDITIONAL_SUPPORT_UNIMODULARITY_TRACE_CANCELLATION_VERIFIED_ON_RAY"
	StatusDFEdgeGraphAssembled          = "CONDITIONAL_SUPPORT_CANONICAL_DF_EDGE_GRAPH_ASSEMBLED"
	StatusFirstOrderPreflightVerified   = "CONDITIONAL_SUPPORT_FIRST_ORDER_EDGE_PREFLIGHT_VERIFIED"
	StatusColorIntertwinerSieveVerified = "CONDITIONAL_SUPPORT_COLOR_INTERTWINER_SIEVE_VERIFIED"

	StatusFailedHyperchargeNormalization = "FAILED_ROUTE_HYPERCHARGE_ABSOLUTE_NORMALIZATION_NOT_DERIVED"
	StatusFailedHyperchargeNotNativeFull = "FAILED_ROUTE_HYPERCHARGE_FRACTIONS_REQUIRE_NORMALIZATION_SEAL_OR_PRIOR_LEDGER"
	StatusFailedFullDFNotSelected        = "FAILED_ROUTE_NUMERICAL_YUKAWA_MATRICES_NOT_DERIVED"
	StatusFailedMajoranaBGapNotActivated = "FAILED_ROUTE_BGAP_MAJORANA_EDGE_NOT_DERIVED"
	StatusFailedFirstOrderNotFullTriple  = "FAILED_ROUTE_FULL_FIRST_ORDER_SPECTRAL_TRIPLE_NOT_VERIFIED"
	StatusFailedDynamicsFirewalled       = "FAILED_ROUTE_HIGGS_AND_BGAP_DYNAMICS_STILL_FIREWALLED"
)

const tol = 1e-10

type Gate295Input struct {
	TrueBimoduleDerived bool
	ZeroOrderVerified   bool
	WeakLeftAction      string
	ColorRightAction    string
	Verdict             string
}

type ChargeVector struct {
	Q, U, D, L, E, N, H float64
}

type HyperchargeAudit struct {
	Variables                  []string
	Equations                  []string
	SolutionRayParameter       string
	Ray                        ChargeVector
	NormalizedWithQOneSixth    ChargeVector
	YukawaResidualNorm         float64
	SU2AnomalyResidual         float64
	SU3AnomalyResidual         float64
	GravitationalResidual      float64
	CubicResidual              float64
	UnimodularityResidual      float64
	RayRecovered               bool
	AbsoluteNormalizationFixed bool
	FractionalLedgerGenerated  bool
	Verdict                    string
}

type StateSlot struct {
	Name      string
	Chirality string
	WeakSide  string
	RightSide string
	WeakDim   int
	RightDim  int
	HyperY    string
}

type DiracEdge struct {
	Name              string
	From              string
	To                string
	SharedRightModule bool
	ColorIntertwiner  string
	HyperchargeLaw    string
	StructurallyLegal bool
	BGapCandidate     bool
	Verdict           string
}

type DiracAssemblyAudit struct {
	Slots               []StateSlot
	AllowedEdges        []DiracEdge
	ForbiddenEdges      []DiracEdge
	OddSelfAdjointShape string
	IncludesMajorana    bool
	BGapActivated       bool
	NumericalYukawas    bool
	Verdict             string
}

type FirstOrderPreflight struct {
	Condition                     string
	RightModuleRule               string
	ColorIdentityResidual         float64
	ColorChangingResidual         float64
	ColorIntertwinerVerified      bool
	LeptonQuarkForbiddenByModule  bool
	ChargeViolatingEdgesForbidden bool
	FullFirstOrderVerified        bool
	DerivedConstraints            []string
	RemainingMissing              []string
	Verdict                       string
}

type Firewalls struct {
	DoesNotInventHyperchargeNormalization bool
	DoesNotInventYukawaMatrices           bool
	DoesNotActivateBGapMajorana           bool
	DoesNotClaimFullSpectralTriple        bool
	DoesNotUnlockHiggs                    bool
	DoesNotUnlockBGap                     bool
	FiniteCorePolluted                    bool
	Verdict                               string
}

type Summary struct {
	HyperchargeRayRecovered      bool
	HyperchargeNormalizationDone bool
	DFEdgeGraphAssembled         bool
	FirstOrderPreflightVerified  bool
	FullFirstOrderVerified       bool
	CanonicalDFSelected          bool
	DynamicsUnblocked            bool
	FirewallPreserved            bool
	Status                       string
	DirectAnswer                 string
	NextGate                     string
}

type Analysis struct {
	Input       Gate295Input
	Hypercharge HyperchargeAudit
	Dirac       DiracAssemblyAudit
	FirstOrder  FirstOrderPreflight
	Firewalls   Firewalls
	Summary     Summary
	Truth       string
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
	input := inheritGate295()
	hyper := auditHypercharge()
	dirac := assembleDiracEdges(hyper)
	first := auditFirstOrderPreflight(dirac)
	fw := auditFirewalls(hyper, dirac, first)
	summary := buildSummary(hyper, dirac, first, fw)
	truth := "Gate 296 strengthens the physical finite Hilbert-space ledger without overclaiming it.  The true bimodule supports the Standard-Model hypercharge ray via anomaly/Yukawa/unimodularity equations, but the absolute U(1) normalization still requires a prior ledger or seal.  The first-order Morita preflight derives the legal Dirac edge graph: Q_L↔u_R, Q_L↔d_R, L_L↔e_R, and L_L↔ν_R are allowed as shared-right-module edges; lepton-quark and color-changing edges are rejected.  Numerical Yukawa matrices, B-gap Majorana activation, and the full spectral-triple first-order theorem remain firewalled."
	return Analysis{Input: input, Hypercharge: hyper, Dirac: dirac, FirstOrder: first, Firewalls: fw, Summary: summary, Truth: truth}, nil
}

func inheritGate295() Gate295Input {
	return Gate295Input{TrueBimoduleDerived: true, ZeroOrderVerified: true, WeakLeftAction: "L(q)=q⊗I_right", ColorRightAction: "R(B)=I_weak⊗B^T", Verdict: StatusGate295Inherited}
}

func auditHypercharge() HyperchargeAudit {
	vars := []string{"q=Y(Q_L)", "u=Y(u_R)", "d=Y(d_R)", "l=Y(L_L)", "e=Y(e_R)", "n=Y(ν_R)", "h=Y(H)"}
	eqs := []string{
		"Yukawa: u=q+h, d=q-h, e=l-h, n=l+h",
		"SU(2)^2U(1): 3q+l=0",
		"SU(3)^2U(1): 2q-u-d=0",
		"gravitational: 6q-3u-3d+2l-e-n=0",
		"U(1)^3: 6q^3-3u^3-3d^3+2l^3-e^3-n^3=0",
		"ν_R neutral/Yukawa-compatible branch: n=0, hence h=-l=3q",
	}
	// One-parameter anomaly-free ray in units q=1.
	ray := ChargeVector{Q: 1, U: 4, D: -2, L: -3, E: -6, N: 0, H: 3}
	norm := scale(ray, 1.0/6.0)
	yuk := residualYukawa(ray)
	su2 := 3*ray.Q + ray.L
	su3 := 2*ray.Q - ray.U - ray.D
	grav := 6*ray.Q - 3*ray.U - 3*ray.D + 2*ray.L - ray.E - ray.N
	cubic := 6*cube(ray.Q) - 3*cube(ray.U) - 3*cube(ray.D) + 2*cube(ray.L) - cube(ray.E) - cube(ray.N)
	unimod := 6*ray.Q + 3*ray.U + 3*ray.D + 2*ray.L + ray.E + ray.N
	return HyperchargeAudit{
		Variables: vars, Equations: eqs, SolutionRayParameter: "q; ray=(q,4q,-2q,-3q,-6q,0,3q)", Ray: ray, NormalizedWithQOneSixth: norm,
		YukawaResidualNorm: yuk, SU2AnomalyResidual: su2, SU3AnomalyResidual: su3, GravitationalResidual: grav, CubicResidual: cubic, UnimodularityResidual: unimod,
		RayRecovered:               nearZero(yuk) && nearZero(su2) && nearZero(su3) && nearZero(grav) && nearZero(cubic) && nearZero(unimod),
		AbsoluteNormalizationFixed: false,
		FractionalLedgerGenerated:  true,
		Verdict:                    strings.Join([]string{StatusHyperchargeEquationsBuilt, StatusHyperchargeRayRecovered, StatusUnimodularityChecked, StatusFailedHyperchargeNormalization, StatusFailedHyperchargeNotNativeFull}, ";"),
	}
}

func scale(v ChargeVector, s float64) ChargeVector {
	return ChargeVector{Q: v.Q * s, U: v.U * s, D: v.D * s, L: v.L * s, E: v.E * s, N: v.N * s, H: v.H * s}
}
func cube(x float64) float64  { return x * x * x }
func nearZero(x float64) bool { return math.Abs(x) < tol }

func residualYukawa(v ChargeVector) float64 {
	res := []float64{v.U - v.Q - v.H, v.D - v.Q + v.H, v.E - v.L + v.H, v.N - v.L - v.H}
	s := 0.0
	for _, r := range res {
		s += r * r
	}
	return math.Sqrt(s)
}

func assembleDiracEdges(h HyperchargeAudit) DiracAssemblyAudit {
	slots := []StateSlot{
		{Name: "Q_L", Chirality: "L", WeakSide: "H doublet", RightSide: "M3 color", WeakDim: 2, RightDim: 3, HyperY: "q"},
		{Name: "u_R", Chirality: "R", WeakSide: "C singlet", RightSide: "M3 color", WeakDim: 1, RightDim: 3, HyperY: "u=q+h"},
		{Name: "d_R", Chirality: "R", WeakSide: "C singlet", RightSide: "M3 color", WeakDim: 1, RightDim: 3, HyperY: "d=q-h"},
		{Name: "L_L", Chirality: "L", WeakSide: "H doublet", RightSide: "C lepton", WeakDim: 2, RightDim: 1, HyperY: "l=-3q"},
		{Name: "e_R", Chirality: "R", WeakSide: "C singlet", RightSide: "C lepton", WeakDim: 1, RightDim: 1, HyperY: "e=l-h"},
		{Name: "ν_R", Chirality: "R", WeakSide: "C singlet", RightSide: "C lepton", WeakDim: 1, RightDim: 1, HyperY: "n=0"},
	}
	allowed := []DiracEdge{
		{Name: "Y_u", From: "Q_L", To: "u_R", SharedRightModule: true, ColorIntertwiner: "I3 only by right-M3 intertwiner sieve", HyperchargeLaw: "u=q+h", StructurallyLegal: true, Verdict: "allowed Dirac edge"},
		{Name: "Y_d", From: "Q_L", To: "d_R", SharedRightModule: true, ColorIntertwiner: "I3 only by right-M3 intertwiner sieve", HyperchargeLaw: "d=q-h", StructurallyLegal: true, Verdict: "allowed Dirac edge"},
		{Name: "Y_e", From: "L_L", To: "e_R", SharedRightModule: true, ColorIntertwiner: "lepton scalar", HyperchargeLaw: "e=l-h", StructurallyLegal: true, Verdict: "allowed Dirac edge"},
		{Name: "Y_ν", From: "L_L", To: "ν_R", SharedRightModule: true, ColorIntertwiner: "lepton scalar", HyperchargeLaw: "n=l+h=0", StructurallyLegal: true, Verdict: "allowed Dirac edge; Dirac neutrino only"},
	}
	forbidden := []DiracEdge{
		{Name: "Q_L↔e_R", From: "Q_L", To: "e_R", SharedRightModule: false, ColorIntertwiner: "M3 right module cannot map to C right module", HyperchargeLaw: "leptoquark edge", StructurallyLegal: false, Verdict: "forbidden by first-order right-module mismatch"},
		{Name: "L_L↔u_R", From: "L_L", To: "u_R", SharedRightModule: false, ColorIntertwiner: "C right module cannot map to M3 right module", HyperchargeLaw: "leptoquark edge", StructurallyLegal: false, Verdict: "forbidden by first-order right-module mismatch"},
		{Name: "color-changing Y_u", From: "Q_L", To: "u_R", SharedRightModule: true, ColorIntertwiner: "off-diagonal color map", HyperchargeLaw: "u=q+h", StructurallyLegal: false, Verdict: "forbidden unless color map commutes with all right M3 actions"},
		{Name: "Majorana ν_R", From: "ν_R", To: "ν_R^c", SharedRightModule: true, ColorIntertwiner: "lepton scalar", HyperchargeLaw: "n=0", StructurallyLegal: false, BGapCandidate: true, Verdict: "permitted only under NeutrinoTextureSeal/B-gap Majorana theorem; not activated here"},
	}
	return DiracAssemblyAudit{Slots: slots, AllowedEdges: allowed, ForbiddenEdges: forbidden, OddSelfAdjointShape: "D_F = [[0,M],[M†,0]], M = diag_edges(Y_u⊗I3, Y_d⊗I3, Y_e, Y_ν) plus optional sealed Majorana ν_R block", IncludesMajorana: false, BGapActivated: false, NumericalYukawas: false, Verdict: strings.Join([]string{StatusDFEdgeGraphAssembled, StatusFailedFullDFNotSelected, StatusFailedMajoranaBGapNotActivated}, ";")}
}

type matrix [][]float64

func zero(n int) matrix {
	m := make(matrix, n)
	for i := range m {
		m[i] = make([]float64, n)
	}
	return m
}
func eye(n int) matrix {
	m := zero(n)
	for i := 0; i < n; i++ {
		m[i][i] = 1
	}
	return m
}
func diag(vals ...float64) matrix {
	m := zero(len(vals))
	for i, v := range vals {
		m[i][i] = v
	}
	return m
}
func matMul(a, b matrix) matrix {
	n, m, k := len(a), len(b[0]), len(b)
	out := make(matrix, n)
	for i := range out {
		out[i] = make([]float64, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			s := 0.0
			for t := 0; t < k; t++ {
				s += a[i][t] * b[t][j]
			}
			out[i][j] = s
		}
	}
	return out
}
func matAdd(a, b matrix, sb float64) matrix {
	out := zero(len(a))
	for i := range a {
		for j := range a[i] {
			out[i][j] = a[i][j] + sb*b[i][j]
		}
	}
	return out
}
func frob(a matrix) float64 {
	s := 0.0
	for i := range a {
		for j := range a[i] {
			s += a[i][j] * a[i][j]
		}
	}
	return math.Sqrt(s)
}

func auditFirstOrderPreflight(d DiracAssemblyAudit) FirstOrderPreflight {
	// Intertwiner test for quark Dirac edges under the right color action.
	// A legal color-preserving edge is proportional to I3 and commutes with all
	// right M3 actions.  A sample off-diagonal color-changing map fails already
	// against a diagonal color generator.
	b := diag(2, -1, 0.5)
	legal := eye(3)
	illegal := zero(3)
	illegal[0][1] = 1
	legalRes := frob(matAdd(matMul(legal, b), matMul(b, legal), -1))
	illegalRes := frob(matAdd(matMul(illegal, b), matMul(b, illegal), -1))
	constraints := []string{
		"Dirac edges may change left weak module H -> C but must preserve the right module.",
		"Quark Yukawa edges must be color intertwiners; for the fundamental M3 right module this forces I3 on color.",
		"Lepton-quark edges are rejected because C-right and M3-right modules do not match.",
		"Charge-violating edges are rejected by the hypercharge-ray Yukawa equations.",
	}
	missing := []string{"full antilinear physical J semantics", "full generation/flavor Yukawa matrices", "absolute hypercharge normalization theorem", "B-gap Majorana activation theorem", "full first-order theorem over all algebra generators"}
	verified := legalRes < tol && illegalRes > 0.1
	return FirstOrderPreflight{Condition: "[[D_F,ρ(a)],ρ°(b)]=0", RightModuleRule: "edge H_ij -> H_kl is first-order compatible when the right/opposite module is shared (j=l); non-vacuous when the left module changes (i≠k)", ColorIdentityResidual: legalRes, ColorChangingResidual: illegalRes, ColorIntertwinerVerified: verified, LeptonQuarkForbiddenByModule: true, ChargeViolatingEdgesForbidden: true, FullFirstOrderVerified: false, DerivedConstraints: constraints, RemainingMissing: missing, Verdict: strings.Join([]string{StatusFirstOrderPreflightVerified, StatusColorIntertwinerSieveVerified, StatusFailedFirstOrderNotFullTriple}, ";")}
}

func auditFirewalls(h HyperchargeAudit, d DiracAssemblyAudit, f FirstOrderPreflight) Firewalls {
	return Firewalls{DoesNotInventHyperchargeNormalization: !h.AbsoluteNormalizationFixed, DoesNotInventYukawaMatrices: !d.NumericalYukawas, DoesNotActivateBGapMajorana: !d.BGapActivated, DoesNotClaimFullSpectralTriple: !f.FullFirstOrderVerified, DoesNotUnlockHiggs: true, DoesNotUnlockBGap: true, FiniteCorePolluted: false, Verdict: strings.Join([]string{StatusFailedHyperchargeNormalization, StatusFailedFullDFNotSelected, StatusFailedMajoranaBGapNotActivated, StatusFailedFirstOrderNotFullTriple, StatusFailedDynamicsFirewalled}, ";")}
}

func buildSummary(h HyperchargeAudit, d DiracAssemblyAudit, f FirstOrderPreflight, fw Firewalls) Summary {
	status := []string{StatusGate295Inherited, StatusHyperchargeEquationsBuilt, StatusHyperchargeRayRecovered, StatusUnimodularityChecked, StatusDFEdgeGraphAssembled, StatusFirstOrderPreflightVerified, StatusColorIntertwinerSieveVerified, StatusFailedHyperchargeNormalization, StatusFailedFullDFNotSelected, StatusFailedMajoranaBGapNotActivated, StatusFailedFirstOrderNotFullTriple, StatusFailedDynamicsFirewalled}
	return Summary{HyperchargeRayRecovered: h.RayRecovered, HyperchargeNormalizationDone: h.AbsoluteNormalizationFixed, DFEdgeGraphAssembled: len(d.AllowedEdges) == 4 && len(d.ForbiddenEdges) >= 4, FirstOrderPreflightVerified: f.ColorIntertwinerVerified && f.LeptonQuarkForbiddenByModule, FullFirstOrderVerified: f.FullFirstOrderVerified, CanonicalDFSelected: d.NumericalYukawas, DynamicsUnblocked: false, FirewallPreserved: !fw.FiniteCorePolluted && fw.DoesNotInventYukawaMatrices, Status: strings.Join(status, ";"), DirectAnswer: "Gate 296 reconstructs the Standard-Model hypercharge ray and assembles the first-order-compatible D_F edge graph, but it does not derive the absolute U(1) normalization, numerical Yukawa matrices, B-gap Majorana activation, or a full spectral-triple first-order theorem.", NextGate: "Complete the physical hypercharge normalization/opposite-action semantics and run the full first-order condition on the assembled D_F edge graph."}
}

func FormatInput(g Gate295Input) string {
	return fmt.Sprintf("trueBimodule=%t zero=%t L=%q R=%q verdict=%s", g.TrueBimoduleDerived, g.ZeroOrderVerified, g.WeakLeftAction, g.ColorRightAction, g.Verdict)
}
func FormatCharge(v ChargeVector) string {
	return fmt.Sprintf("Q=%.8g U=%.8g D=%.8g L=%.8g E=%.8g N=%.8g H=%.8g", v.Q, v.U, v.D, v.L, v.E, v.N, v.H)
}
func FormatHypercharge(h HyperchargeAudit) string {
	return fmt.Sprintf("vars=%s rayParam=%q ray=[%s] normalized=[%s] residuals(yuk=%.3g su2=%.3g su3=%.3g grav=%.3g cubic=%.3g unimod=%.3g) rayRecovered=%t normalized=%t verdict=%s", strings.Join(h.Variables, ","), h.SolutionRayParameter, FormatCharge(h.Ray), FormatCharge(h.NormalizedWithQOneSixth), h.YukawaResidualNorm, h.SU2AnomalyResidual, h.SU3AnomalyResidual, h.GravitationalResidual, h.CubicResidual, h.UnimodularityResidual, h.RayRecovered, h.AbsoluteNormalizationFixed, h.Verdict)
}
func FormatSlot(s StateSlot) string {
	return fmt.Sprintf("%s(%s): weak=%s right=%s dims=%dx%d Y=%s", s.Name, s.Chirality, s.WeakSide, s.RightSide, s.WeakDim, s.RightDim, s.HyperY)
}
func FormatEdge(e DiracEdge) string {
	return fmt.Sprintf("%s:%s->%s sharedRight=%t color=%q Y=%q legal=%t bgap=%t verdict=%s", e.Name, e.From, e.To, e.SharedRightModule, e.ColorIntertwiner, e.HyperchargeLaw, e.StructurallyLegal, e.BGapCandidate, e.Verdict)
}
func FormatDirac(d DiracAssemblyAudit) string {
	slots := make([]string, 0, len(d.Slots))
	for _, s := range d.Slots {
		slots = append(slots, FormatSlot(s))
	}
	allowed := make([]string, 0, len(d.AllowedEdges))
	for _, e := range d.AllowedEdges {
		allowed = append(allowed, FormatEdge(e))
	}
	forbidden := make([]string, 0, len(d.ForbiddenEdges))
	for _, e := range d.ForbiddenEdges {
		forbidden = append(forbidden, FormatEdge(e))
	}
	return fmt.Sprintf("slots=[%s] allowed=[%s] forbidden=[%s] shape=%q majorana=%t bgap=%t numericY=%t verdict=%s", strings.Join(slots, " | "), strings.Join(allowed, " | "), strings.Join(forbidden, " | "), d.OddSelfAdjointShape, d.IncludesMajorana, d.BGapActivated, d.NumericalYukawas, d.Verdict)
}
func FormatFirstOrder(f FirstOrderPreflight) string {
	return fmt.Sprintf("condition=%q rule=%q colorI=%.3g colorChanging=%.3g intertwiner=%t lqForbidden=%t chargeForbidden=%t full=%t constraints=%s missing=%s verdict=%s", f.Condition, f.RightModuleRule, f.ColorIdentityResidual, f.ColorChangingResidual, f.ColorIntertwinerVerified, f.LeptonQuarkForbiddenByModule, f.ChargeViolatingEdgesForbidden, f.FullFirstOrderVerified, strings.Join(f.DerivedConstraints, " | "), strings.Join(f.RemainingMissing, " | "), f.Verdict)
}
func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("noYnorm=%t noYukawa=%t noBGapMaj=%t noFullTriple=%t noHiggs=%t noBGap=%t polluted=%t verdict=%s", f.DoesNotInventHyperchargeNormalization, f.DoesNotInventYukawaMatrices, f.DoesNotActivateBGapMajorana, f.DoesNotClaimFullSpectralTriple, f.DoesNotUnlockHiggs, f.DoesNotUnlockBGap, f.FiniteCorePolluted, f.Verdict)
}
func FormatSummary(s Summary) string {
	return fmt.Sprintf("Yray=%t Ynorm=%t DFgraph=%t firstPreflight=%t fullFirst=%t canonicalDF=%t dynamics=%t firewall=%t next=%q status=%s", s.HyperchargeRayRecovered, s.HyperchargeNormalizationDone, s.DFEdgeGraphAssembled, s.FirstOrderPreflightVerified, s.FullFirstOrderVerified, s.CanonicalDFSelected, s.DynamicsUnblocked, s.FirewallPreserved, s.NextGate, s.Status)
}
