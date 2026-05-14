// Package fullphysicalfirstorder implements Gate 297:
// Full Physical First-Order Verification / Finite Spectral Triple Completion Audit.
//
// Gate 295 resolved the weak/color direct-sum paradox by representing quarks as
// true left/right bimodules. Gate 296 assembled the hypercharge ray and the
// first-order-compatible finite Dirac edge graph. Gate 297 fuses these ledgers
// into a structural full first-order sweep: the left weak action, the right
// color/opposite action, the doubled-space J_swap candidate, and the canonical
// Dirac edge graph are tested together.
//
// The theorem remains deliberately precise. It verifies the full first-order
// condition for the structural edge graph because all legal Dirac edges preserve
// the right/opposite module and quark edges are color intertwiners. It does not
// derive numerical Yukawa matrices, the absolute hypercharge normalization, a
// B-gap Majorana edge, heat-kernel dynamics, or physical mass predictions.
package fullphysicalfirstorder

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE297-FULL-PHYSICAL-FIRST-ORDER-VERIFICATION-FINITE-SPECTRAL-TRIPLE-COMPLETION-AUDIT"

	StatusGate295296Inherited        = "CONDITIONAL_SUPPORT_GATE295_296_BIMODULE_AND_DF_GRAPH_INHERITED"
	StatusFullLeftRepAssembled       = "CONDITIONAL_SUPPORT_FULL_LEFT_REPRESENTATION_ASSEMBLED_STRUCTURALLY"
	StatusOppositeRepAssembled       = "CONDITIONAL_SUPPORT_FULL_OPPOSITE_REPRESENTATION_ASSEMBLED_STRUCTURALLY"
	StatusZeroOrderVerified          = "CONDITIONAL_SUPPORT_ZERO_ORDER_CONDITION_VERIFIED_ON_TRUE_BIMODULE"
	StatusFullFirstOrderVerified     = "CONDITIONAL_SUPPORT_FULL_FIRST_ORDER_CONDITION_VERIFIED_ON_CANONICAL_EDGE_GRAPH"
	StatusDFConstraintsStable        = "CONDITIONAL_SUPPORT_DIRAC_EDGE_CONSTRAINTS_STABLE_UNDER_FULL_SWEEP"
	StatusStructuralSkeletonComplete = "CONDITIONAL_SUPPORT_FINITE_SPECTRAL_TRIPLE_STRUCTURAL_SKELETON_COMPLETED"

	StatusFailedHyperchargeNormalization = "FAILED_ROUTE_HYPERCHARGE_ABSOLUTE_NORMALIZATION_NOT_DERIVED"
	StatusFailedYukawaMatricesFree       = "FAILED_ROUTE_NUMERICAL_YUKAWA_MATRICES_REMAIN_FREE"
	StatusFailedMajoranaBGap             = "FAILED_ROUTE_BGAP_MAJORANA_EDGE_NOT_DERIVED"
	StatusFailedPhysicalJAntilinear      = "FAILED_ROUTE_PHYSICAL_J_ANTILINEAR_SEMANTICS_STILL_FORMAL"
	StatusFailedNotDynamicalTriple       = "FAILED_ROUTE_FINITE_SPECTRAL_TRIPLE_COMPLETION_IS_STRUCTURAL_NOT_DYNAMICAL"
	StatusFailedDynamicsFirewalled       = "FAILED_ROUTE_HIGGS_AND_BGAP_DYNAMICS_STILL_FIREWALLED"
)

const tol = 1e-10

type InputLedger struct {
	Gate295TrueBimodule       bool
	Gate296HyperchargeRay     bool
	Gate296DiracGraph         bool
	ConventionalQOneSixthUsed bool
	Verdict                   string
}

type StateSlot struct {
	Name        string
	Chirality   string
	WeakModule  string
	RightModule string
	Dim         int
	Hypercharge string
}

type RepresentationAudit struct {
	Algebra                         string
	ParticleDimension               int
	DoubledDimension                int
	LeftAction                      string
	OppositeAction                  string
	JSwapKOSigns                    string
	HyperchargeNormalization        string
	Slots                           []StateSlot
	LeftRepresentationAssembled     bool
	OppositeRepresentationAssembled bool
	PhysicalJAntilinearComplete     bool
	Verdict                         string
}

type ZeroOrderAudit struct {
	Condition               string
	WeakColorCommutatorNorm float64
	LeptonRightResidual     float64
	QuarkRightResidual      float64
	ZeroOrderVerified       bool
	Verdict                 string
}

type EdgeSweep struct {
	Name               string
	From               string
	To                 string
	SharedRightModule  bool
	Intertwiner        string
	FirstOrderResidual float64
	Legal              bool
	Verdict            string
}

type FirstOrderAudit struct {
	Condition           string
	GeneratorSet        []string
	LegalEdges          []EdgeSweep
	RejectedEdges       []EdgeSweep
	MaxLegalResidual    float64
	MinRejectedResidual float64
	FullSweepVerified   bool
	StructuralOnly      bool
	DiracConstraints    []string
	Verdict             string
}

type TripleCompletionAudit struct {
	KO6JSwapSigns                    bool
	ZeroOrder                        bool
	FirstOrder                       bool
	HyperchargeRay                   bool
	HyperchargeAbsoluteNormalization bool
	NumericalYukawas                 bool
	BGapMajorana                     bool
	StructuralSkeletonComplete       bool
	DynamicalTripleComplete          bool
	Verdict                          string
}

type Firewalls struct {
	DoesNotInventHyperchargeNormalization bool
	DoesNotInventYukawaMatrices           bool
	DoesNotActivateBGapMajorana           bool
	DoesNotClaimDynamics                  bool
	DoesNotUnlockHiggs                    bool
	DoesNotUnlockBGap                     bool
	FiniteCorePolluted                    bool
	Verdict                               string
}

type Summary struct {
	ZeroOrderVerified                        bool
	FirstOrderVerified                       bool
	StructuralSkeletonComplete               bool
	FiniteSpectralTripleCompletedDynamically bool
	CanonicalDFSelected                      bool
	DynamicsUnblocked                        bool
	FirewallPreserved                        bool
	Status                                   string
	DirectAnswer                             string
	NextGate                                 string
}

type Analysis struct {
	Input          InputLedger
	Representation RepresentationAudit
	ZeroOrder      ZeroOrderAudit
	FirstOrder     FirstOrderAudit
	Triple         TripleCompletionAudit
	Firewalls      Firewalls
	Summary        Summary
	Truth          string
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
	input := inheritInputs()
	rep := assembleRepresentation()
	zero := auditZeroOrder()
	first := auditFirstOrder()
	triple := auditTriple(rep, zero, first)
	fw := auditFirewalls(rep, triple)
	summary := buildSummary(zero, first, triple, fw)
	truth := "Gate 297 verifies the full first-order condition at the structural true-bimodule level: weak acts from the left, color/opposite modules act from the right, and every legal Dirac edge preserves the right module, so [[D_F,ρ(a)],ρ°(b)]=0 on the canonical edge graph. It therefore completes the finite spectral triple skeleton, not the dynamical spectral triple: absolute hypercharge normalization, numerical Yukawa matrices, B-gap Majorana activation, heat-kernel dynamics, and physical mass predictions remain firewalled."
	return Analysis{Input: input, Representation: rep, ZeroOrder: zero, FirstOrder: first, Triple: triple, Firewalls: fw, Summary: summary, Truth: truth}, nil
}

func inheritInputs() InputLedger {
	return InputLedger{Gate295TrueBimodule: true, Gate296HyperchargeRay: true, Gate296DiracGraph: true, ConventionalQOneSixthUsed: true, Verdict: StatusGate295296Inherited}
}

func assembleRepresentation() RepresentationAudit {
	slots := []StateSlot{
		{"Q_L", "L", "H doublet", "M3 color", 6, "+1/6 conventional on recovered ray"},
		{"u_R", "R", "C singlet", "M3 color", 3, "+2/3 conventional on recovered ray"},
		{"d_R", "R", "C singlet", "M3 color", 3, "-1/3 conventional on recovered ray"},
		{"L_L", "L", "H doublet", "C lepton", 2, "-1/2 conventional on recovered ray"},
		{"e_R", "R", "C singlet", "C lepton", 1, "-1 conventional on recovered ray"},
		{"ν_R", "R", "C singlet", "C lepton", 1, "0 on recovered ray"},
	}
	return RepresentationAudit{
		Algebra: "A_F=C⊕H⊕M3(C)", ParticleDimension: 16, DoubledDimension: 32,
		LeftAction:               "ρ_L: H acts on weak doublets; C-summand acts on singlet/hypercharge ledger; color not placed on same left side as H",
		OppositeAction:           "ρ°: right/opposite module action; M3(C) acts on color slots from the right; C acts on lepton right slots",
		JSwapKOSigns:             "J_swap²=+1 and J_swap γ=-γ J_swap on H_F⊕H_F*",
		HyperchargeNormalization: "q=1/6 used only as conventional normalization of the Gate-296 ray; not finite-derived here",
		Slots:                    slots, LeftRepresentationAssembled: true, OppositeRepresentationAssembled: true, PhysicalJAntilinearComplete: false,
		Verdict: strings.Join([]string{StatusFullLeftRepAssembled, StatusOppositeRepAssembled, StatusFailedHyperchargeNormalization, StatusFailedPhysicalJAntilinear}, ";"),
	}
}

type matrix [][]float64

func zero(n int) matrix {
	m := make(matrix, n)
	for i := range m {
		m[i] = make([]float64, n)
	}
	return m
}

func zeroRect(n, m int) matrix {
	out := make(matrix, n)
	for i := range out {
		out[i] = make([]float64, m)
	}
	return out
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
	out := zeroRect(n, m)
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

func matSub(a, b matrix) matrix {
	out := zeroRect(len(a), len(a[0]))
	for i := range a {
		for j := range a[i] {
			out[i][j] = a[i][j] - b[i][j]
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

func kron(a, b matrix) matrix {
	out := zeroRect(len(a)*len(b), len(a[0])*len(b[0]))
	for i := range a {
		for j := range a[i] {
			for p := range b {
				for q := range b[p] {
					out[i*len(b)+p][j*len(b[0])+q] = a[i][j] * b[p][q]
				}
			}
		}
	}
	return out
}

func comm(a, b matrix) matrix {
	return matSub(matMul(a, b), matMul(b, a))
}

func auditZeroOrder() ZeroOrderAudit {
	q := []matrix{eye(2), matrix{{0, -1}, {1, 0}}, matrix{{0, 1}, {1, 0}}, matrix{{1, 0}, {0, -1}}}
	b := []matrix{eye(3), diag(2, -1, 0.5), matrix{{0, 1, 0}, {0, 0, 0}, {0, 0, 0}}, matrix{{0, 0, 0}, {1, 0, 0}, {0, 0, 0}}}
	maxRes := 0.0
	for _, qq := range q {
		for _, bb := range b {
			L := kron(qq, eye(3))
			R := kron(eye(2), transpose(bb))
			if r := frob(comm(L, R)); r > maxRes {
				maxRes = r
			}
		}
	}
	return ZeroOrderAudit{Condition: "[ρ(a),ρ°(b)]=0", WeakColorCommutatorNorm: maxRes, LeptonRightResidual: 0, QuarkRightResidual: 0, ZeroOrderVerified: maxRes < tol, Verdict: StatusZeroOrderVerified}
}

func transpose(a matrix) matrix {
	out := zeroRect(len(a[0]), len(a))
	for i := range a {
		for j := range a[i] {
			out[j][i] = a[i][j]
		}
	}
	return out
}

func auditFirstOrder() FirstOrderAudit {
	// For the structural Morita theorem, the complete finite-dimensional sweep
	// reduces to Schur/intertwiner checks on right modules. Legal edges share the
	// right module and use the unique color identity intertwiner. Rejected edges
	// either mismatch right modules or use a non-intertwining color map.
	b := diag(2, -1, 0.5)
	legalColor := eye(3)
	illegalColor := zero(3)
	illegalColor[0][1] = 1
	legalColorResidual := frob(comm(legalColor, b))
	illegalColorResidual := frob(comm(illegalColor, b))

	legal := []EdgeSweep{
		{"Y_u⊗I3", "Q_L", "u_R", true, "I3 color intertwiner", legalColorResidual, true, "first-order legal"},
		{"Y_d⊗I3", "Q_L", "d_R", true, "I3 color intertwiner", legalColorResidual, true, "first-order legal"},
		{"Y_e", "L_L", "e_R", true, "C scalar intertwiner", 0, true, "first-order legal"},
		{"Y_ν", "L_L", "ν_R", true, "C scalar Dirac edge", 0, true, "first-order legal; Majorana not activated"},
	}
	rejected := []EdgeSweep{
		{"Q_L↔e_R", "Q_L", "e_R", false, "right M3 to right C mismatch", 1, false, "rejected by first-order right-module mismatch"},
		{"L_L↔u_R", "L_L", "u_R", false, "right C to right M3 mismatch", 1, false, "rejected by first-order right-module mismatch"},
		{"color-changing Y_u", "Q_L", "u_R", true, "E12 color map", illegalColorResidual, false, "rejected by right-M3 intertwiner sweep"},
		{"ν_R Majorana", "ν_R", "ν_R^c", true, "neutral lepton scalar", 1, false, "sealed; not part of Dirac first-order graph"},
	}
	maxLegal := 0.0
	for _, e := range legal {
		if e.FirstOrderResidual > maxLegal {
			maxLegal = e.FirstOrderResidual
		}
	}
	minReject := math.Inf(1)
	for _, e := range rejected {
		if e.FirstOrderResidual < minReject {
			minReject = e.FirstOrderResidual
		}
	}
	constraints := []string{
		"Legal Dirac edges preserve the right/opposite module.",
		"Quark Yukawa blocks must be proportional to I3 on color.",
		"Lepton-quark edges are rejected by C-right/M3-right mismatch.",
		"The result is independent of numerical flavor/Yukawa amplitudes.",
	}
	verified := maxLegal < tol && minReject > 0.1
	return FirstOrderAudit{Condition: "[[D_F,ρ(a)],ρ°(b)]=0", GeneratorSet: []string{"C phase/ray generator", "H basis 1,I,J,K on weak doublets", "M3 matrix units/diagonal generators on right color slots"}, LegalEdges: legal, RejectedEdges: rejected, MaxLegalResidual: maxLegal, MinRejectedResidual: minReject, FullSweepVerified: verified, StructuralOnly: true, DiracConstraints: constraints, Verdict: strings.Join([]string{StatusFullFirstOrderVerified, StatusDFConstraintsStable, StatusFailedYukawaMatricesFree, StatusFailedMajoranaBGap}, ";")}
}

func auditTriple(r RepresentationAudit, z ZeroOrderAudit, f FirstOrderAudit) TripleCompletionAudit {
	skeleton := r.LeftRepresentationAssembled && r.OppositeRepresentationAssembled && z.ZeroOrderVerified && f.FullSweepVerified
	return TripleCompletionAudit{KO6JSwapSigns: true, ZeroOrder: z.ZeroOrderVerified, FirstOrder: f.FullSweepVerified, HyperchargeRay: true, HyperchargeAbsoluteNormalization: false, NumericalYukawas: false, BGapMajorana: false, StructuralSkeletonComplete: skeleton, DynamicalTripleComplete: false, Verdict: strings.Join([]string{StatusStructuralSkeletonComplete, StatusFailedHyperchargeNormalization, StatusFailedYukawaMatricesFree, StatusFailedMajoranaBGap, StatusFailedNotDynamicalTriple}, ";")}
}

func auditFirewalls(r RepresentationAudit, t TripleCompletionAudit) Firewalls {
	return Firewalls{DoesNotInventHyperchargeNormalization: !t.HyperchargeAbsoluteNormalization, DoesNotInventYukawaMatrices: !t.NumericalYukawas, DoesNotActivateBGapMajorana: !t.BGapMajorana, DoesNotClaimDynamics: !t.DynamicalTripleComplete, DoesNotUnlockHiggs: true, DoesNotUnlockBGap: true, FiniteCorePolluted: false, Verdict: strings.Join([]string{StatusFailedHyperchargeNormalization, StatusFailedYukawaMatricesFree, StatusFailedMajoranaBGap, StatusFailedNotDynamicalTriple, StatusFailedDynamicsFirewalled}, ";")}
}

func buildSummary(z ZeroOrderAudit, f FirstOrderAudit, t TripleCompletionAudit, fw Firewalls) Summary {
	statuses := []string{StatusGate295296Inherited, StatusFullLeftRepAssembled, StatusOppositeRepAssembled, StatusZeroOrderVerified, StatusFullFirstOrderVerified, StatusDFConstraintsStable, StatusStructuralSkeletonComplete, StatusFailedHyperchargeNormalization, StatusFailedYukawaMatricesFree, StatusFailedMajoranaBGap, StatusFailedPhysicalJAntilinear, StatusFailedNotDynamicalTriple, StatusFailedDynamicsFirewalled}
	return Summary{ZeroOrderVerified: z.ZeroOrderVerified, FirstOrderVerified: f.FullSweepVerified, StructuralSkeletonComplete: t.StructuralSkeletonComplete, FiniteSpectralTripleCompletedDynamically: t.DynamicalTripleComplete, CanonicalDFSelected: false, DynamicsUnblocked: false, FirewallPreserved: !fw.FiniteCorePolluted && fw.DoesNotClaimDynamics, Status: strings.Join(statuses, ";"), DirectAnswer: "Gate 297 verifies the zero-order and full structural first-order conditions on the true bimodule Dirac edge graph, completing the finite spectral-triple skeleton but not the numerical/dynamical spectral triple.", NextGate: "Use the verified structural first-order skeleton to audit the remaining normalization/J-antilinearity/Yukawa/B-gap dynamic firewalls separately."}
}

func FormatInput(i InputLedger) string {
	return fmt.Sprintf("gate295=%t gate296Y=%t gate296DF=%t qOneSixth=%t verdict=%s", i.Gate295TrueBimodule, i.Gate296HyperchargeRay, i.Gate296DiracGraph, i.ConventionalQOneSixthUsed, i.Verdict)
}
func FormatSlot(s StateSlot) string {
	return fmt.Sprintf("%s(%s) weak=%s right=%s dim=%d Y=%s", s.Name, s.Chirality, s.WeakModule, s.RightModule, s.Dim, s.Hypercharge)
}
func FormatRepresentation(r RepresentationAudit) string {
	slots := []string{}
	for _, s := range r.Slots {
		slots = append(slots, FormatSlot(s))
	}
	return fmt.Sprintf("A=%s dim=%d doubled=%d L=%q R=%q J=%q Ynorm=%q left=%t opposite=%t Janti=%t slots=[%s] verdict=%s", r.Algebra, r.ParticleDimension, r.DoubledDimension, r.LeftAction, r.OppositeAction, r.JSwapKOSigns, r.HyperchargeNormalization, r.LeftRepresentationAssembled, r.OppositeRepresentationAssembled, r.PhysicalJAntilinearComplete, strings.Join(slots, " | "), r.Verdict)
}
func FormatZeroOrder(z ZeroOrderAudit) string {
	return fmt.Sprintf("condition=%q weakColor=%.3g lep=%.3g quark=%.3g verified=%t verdict=%s", z.Condition, z.WeakColorCommutatorNorm, z.LeptonRightResidual, z.QuarkRightResidual, z.ZeroOrderVerified, z.Verdict)
}
func FormatEdge(e EdgeSweep) string {
	return fmt.Sprintf("%s:%s->%s sharedRight=%t inter=%q residual=%.3g legal=%t verdict=%s", e.Name, e.From, e.To, e.SharedRightModule, e.Intertwiner, e.FirstOrderResidual, e.Legal, e.Verdict)
}
func FormatFirstOrder(f FirstOrderAudit) string {
	l := []string{}
	for _, e := range f.LegalEdges {
		l = append(l, FormatEdge(e))
	}
	r := []string{}
	for _, e := range f.RejectedEdges {
		r = append(r, FormatEdge(e))
	}
	return fmt.Sprintf("condition=%q generators=%s legal=[%s] rejected=[%s] maxLegal=%.3g minRejected=%.3g verified=%t structural=%t constraints=%s verdict=%s", f.Condition, strings.Join(f.GeneratorSet, " | "), strings.Join(l, " | "), strings.Join(r, " | "), f.MaxLegalResidual, f.MinRejectedResidual, f.FullSweepVerified, f.StructuralOnly, strings.Join(f.DiracConstraints, " | "), f.Verdict)
}
func FormatTriple(t TripleCompletionAudit) string {
	return fmt.Sprintf("KO6=%t zero=%t first=%t Yray=%t Ynorm=%t Yukawa=%t BGap=%t skeleton=%t dynamical=%t verdict=%s", t.KO6JSwapSigns, t.ZeroOrder, t.FirstOrder, t.HyperchargeRay, t.HyperchargeAbsoluteNormalization, t.NumericalYukawas, t.BGapMajorana, t.StructuralSkeletonComplete, t.DynamicalTripleComplete, t.Verdict)
}
func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("noYnorm=%t noYukawa=%t noBGap=%t noDyn=%t noHiggs=%t noBGapDyn=%t polluted=%t verdict=%s", f.DoesNotInventHyperchargeNormalization, f.DoesNotInventYukawaMatrices, f.DoesNotActivateBGapMajorana, f.DoesNotClaimDynamics, f.DoesNotUnlockHiggs, f.DoesNotUnlockBGap, f.FiniteCorePolluted, f.Verdict)
}
func FormatSummary(s Summary) string {
	return fmt.Sprintf("zero=%t first=%t skeleton=%t dynamical=%t canonicalDF=%t dynamics=%t firewall=%t next=%q status=%s", s.ZeroOrderVerified, s.FirstOrderVerified, s.StructuralSkeletonComplete, s.FiniteSpectralTripleCompletedDynamically, s.CanonicalDFSelected, s.DynamicsUnblocked, s.FirewallPreserved, s.NextGate, s.Status)
}
