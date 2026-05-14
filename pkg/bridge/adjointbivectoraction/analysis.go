// Package adjointbivectoraction implements Gate 250:
// Adjoint Bivector Action / Explicit Q_8v Matrix Derivation Audit.
//
// Gate 249 identified the coordinate-free neutral-kernel route as the right
// strategy, but blocked it because Q_8v and Z_8v were not derived. Gate 250
// tests whether the missing matrices can be obtained from native Clifford
// bivector commutators on the vector carrier 8_v.
//
// The audit is intentionally strict. It verifies the Clifford commutator action
// of an explicit grade-2 blade on grade-1 vectors, but it does not pretend that
// the scalar-bundle electroweak generators T3 and Y_phi have already been
// retrieved as Cl(1,7) bivectors. It also records a stronger structural warning:
// any real Clifford-bivector adjoint action on 8_v is skew-symmetric, hence has
// even rank and an even-dimensional real kernel. Therefore an exact real
// 3-dimensional neutral kernel cannot be produced by a single real bivector
// adjoint action alone.
package adjointbivectoraction

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/neutraleigenspacekernel"
	"github.com/bagherbal/asha-engine/pkg/clifford"
	"github.com/bagherbal/asha-engine/pkg/linear"
)

const (
	AuditID = "GATE250-ADJOINT-BIVECTOR-ACTION-EXPLICIT-Q8V-MATRIX-DERIVATION-AUDIT"

	StatusCliffordAdjointAvailable       = "CONDITIONAL_SUPPORT_CLIFFORD_BIVECTOR_ADJOINT_ACTION_AVAILABLE"
	StatusCandidateBivectorMatrices      = "CONDITIONAL_SUPPORT_CANDIDATE_BIVECTOR_8V_MATRICES_COMPUTABLE"
	StatusSkewKernelParityObstruction    = "FAILED_ROUTE_REAL_BIVECTOR_ADJOINT_THREE_KERNEL_OBSTRUCTION"
	StatusEWBivectorRetrievalBlocked     = "FAILED_ROUTE_EW_BIVECTOR_RETRIEVAL"
	StatusQ8VMatrixBlocked               = "FAILED_ROUTE_EXPLICIT_Q8V_MATRIX_DERIVATION"
	StatusNeutralKernel3PlaneBlocked     = "FAILED_ROUTE_Q8V_NEUTRAL_3PLANE_DERIVATION"
	StatusScalarToNeutralPlaneStillBlock = "FAILED_ROUTE_SCALAR_TO_NEUTRAL_3PLANE_ISOMORPHISM"
	StatusTrialityStillBlocked           = "FAILED_ROUTE_TRIALITY_PREFLIGHT_WITHOUT_NEUTRAL_8V_VECTOR"
	StatusYukawaStillBlocked             = "FAILED_ROUTE_YUKAWA_TEXTURE_DERIVATION"
)

type InheritedGate249Audit struct {
	Carrier8VKnown            bool
	NeutralKernelStrategy     bool
	EWDerivationActionDerived bool
	NeutralKernelDerived      bool
	NeutralKernelDim3         bool
	ScalarNeutralIsomorphism  bool
	VTauConstructed           bool
	TrialityUnblocked         bool
	YukawaTextureDerived      bool
	TruthStatement            string
}

type CliffordCarrierAudit struct {
	Signature             string
	VectorDimension       int
	MetricDiagonal        []int
	VectorBasis           []string
	Grade1CarrierKnown    bool
	BivectorActionFormula string
	CommutatorActionTyped bool
	Verdict               string
}

type SimpleBivectorMatrixAudit struct {
	Blade               string
	MatrixRows          int
	MatrixCols          int
	SkewSymmetric       bool
	Rank                int
	KernelDimension     int
	KernelDimensionEven bool
	NonzeroEntries      []string
	Verdict             string
}

type GenericBivectorParityAudit struct {
	RealBivectorAdjointMatricesSkew bool
	SkewRankAlwaysEven              bool
	Dimension8KernelAlwaysEven      bool
	Exact3DKernelPossible           bool
	Reason                          string
	Verdict                         string
}

type EWBivectorRetrievalAudit struct {
	RequestedGenerators      []string
	ScalarT3Available        bool
	ScalarYPhiAvailable      bool
	T3Grade2BladeDerived     bool
	YPhiGrade2BladeDerived   bool
	T3BladeLabel             string
	YPhiBladeLabel           string
	QBladeDerived            bool
	ZBladeDerived            bool
	ManualBladeAssignment    string
	ManualAssignmentRejected bool
	Obstruction              string
	Verdict                  string
}

type ExplicitMatrixAudit struct {
	AdjointActionFormula       string
	CanConstructCandidateBlade bool
	CanConstructT3Matrix       bool
	CanConstructYPhiMatrix     bool
	Q8VConstructed             bool
	Z8VConstructed             bool
	Q8VRows                    int
	Q8VCols                    int
	Q8VSkewIfBivector          bool
	Q8VKernelDimensionKnown    bool
	Q8VKernelDimension         int
	Neutral3PlaneDerived       bool
	BindingFailure             string
	Verdict                    string
}

type ScalarNeutralPlaneAudit struct {
	NeutralTraceSlots       []string
	TauEta                  []int
	NeedsQ8VKernel          bool
	Q8VKernelAvailable      bool
	KernelDimensionExactly3 bool
	CanonicalIsomorphism    bool
	VTauConstructed         bool
	Reason                  string
	Verdict                 string
}

type FirewallAudit struct {
	InventedT3Blade             bool
	InventedYPhiBlade           bool
	AssignedChargesToGammaBasis bool
	ForcedKernelDim3            bool
	ConstructedVTauByHand       bool
	InventedTrialityMap         bool
	InsertedYukawaTexture       bool
	ClaimedCKMPMNS              bool
	PollutedFiniteCore          bool
	Verdict                     string
}

type Summary struct {
	CliffordAdjointAvailable    bool
	CandidateMatricesComputable bool
	EWBivectorsRetrieved        bool
	Q8VMatrixDerived            bool
	NeutralKernelDerived        bool
	NeutralKernelDim3           bool
	RealBivector3KernelPossible bool
	VTauConstructed             bool
	TrialityUnblocked           bool
	YukawaTextureDerived        bool
	Status                      string
	NextGate                    string
	Comment                     string
}

type Analysis struct {
	PreviousGate249 InheritedGate249Audit
	Carrier         CliffordCarrierAudit
	SimpleBlade     SimpleBivectorMatrixAudit
	KernelParity    GenericBivectorParityAudit
	EWBivectors     EWBivectorRetrievalAudit
	Matrices        ExplicitMatrixAudit
	ScalarPlane     ScalarNeutralPlaneAudit
	Firewall        FirewallAudit
	Summary         Summary
	TruthStatement  string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prevRaw, err := neutraleigenspacekernel.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		prev := inheritGate249(prevRaw)
		carrier, err := auditCarrier()
		if err != nil {
			defaultErr = err
			return
		}
		simple := auditSimpleBivector(carrier.MetricDiagonal, 1, 2)
		parity := auditGenericKernelParity()
		ew := auditEWBivectorRetrieval()
		matrices := auditExplicitMatrices(simple, parity, ew)
		scalar := auditScalarNeutralPlane(matrices)
		firewall := auditFirewall()
		summary := summarize(prev, carrier, simple, parity, ew, matrices, scalar)
		truth := buildTruth(prev, carrier, simple, parity, ew, matrices, scalar)
		defaultA = Analysis{PreviousGate249: prev, Carrier: carrier, SimpleBlade: simple, KernelParity: parity, EWBivectors: ew, Matrices: matrices, ScalarPlane: scalar, Firewall: firewall, Summary: summary, TruthStatement: truth}
	})
	return defaultA, defaultErr
}

func inheritGate249(a neutraleigenspacekernel.Analysis) InheritedGate249Audit {
	return InheritedGate249Audit{
		Carrier8VKnown:            a.Summary.Basis8VKnown,
		NeutralKernelStrategy:     a.Summary.Basis8VKnown && !a.Summary.EWDerivationActionDerived,
		EWDerivationActionDerived: a.Summary.EWDerivationActionDerived,
		NeutralKernelDerived:      a.Summary.NeutralKernelDerived,
		NeutralKernelDim3:         a.Summary.NeutralKernelDim3,
		ScalarNeutralIsomorphism:  a.Summary.ScalarNeutralIsomorphism,
		VTauConstructed:           a.Summary.VTauConstructed,
		TrialityUnblocked:         a.Summary.TrialityUnblocked,
		YukawaTextureDerived:      a.Summary.YukawaTextureDerived,
		TruthStatement:            a.TruthStatement,
	}
}

func auditCarrier() (CliffordCarrierAudit, error) {
	alg, err := clifford.New(clifford.Signature{Positive: 1, Negative: 7})
	if err != nil {
		return CliffordCarrierAudit{}, err
	}
	labels := make([]string, alg.VectorDimension())
	for i := range labels {
		labels[i] = fmt.Sprintf("e%d", i)
	}
	return CliffordCarrierAudit{
		Signature:             "Cl(1,7)",
		VectorDimension:       alg.VectorDimension(),
		MetricDiagonal:        alg.MetricDiagonal(),
		VectorBasis:           labels,
		Grade1CarrierKnown:    alg.VectorDimension() == 8,
		BivectorActionFormula: "R(e_i e_j)e_k = [e_i e_j,e_k] = 2(η_jk e_i - η_ik e_j)",
		CommutatorActionTyped: true,
		Verdict:               "the Clifford algebra supplies a lawful commutator action of explicit grade-2 blades on the grade-1 vector carrier 8_v",
	}, nil
}

func auditSimpleBivector(metric []int, i, j int) SimpleBivectorMatrixAudit {
	m := bivectorAdjointMatrix(metric, i, j)
	skew := isSkew(m, 1e-12)
	rank := rank(m, 1e-12)
	kernel := m.Rows() - rank
	entries := nonzeroEntries(m)
	return SimpleBivectorMatrixAudit{
		Blade:               fmt.Sprintf("e%d∧e%d", i, j),
		MatrixRows:          m.Rows(),
		MatrixCols:          m.Cols(),
		SkewSymmetric:       skew,
		Rank:                rank,
		KernelDimension:     kernel,
		KernelDimensionEven: kernel%2 == 0,
		NonzeroEntries:      entries,
		Verdict:             "a simple Clifford bivector gives an explicit real skew 8v matrix; its kernel is six-dimensional, not three-dimensional",
	}
}

func bivectorAdjointMatrix(metric []int, i, j int) linear.Matrix {
	n := len(metric)
	m := linear.NewMatrix(n, n)
	// Column k is R(e_i e_j)e_k = 2(η_jk e_i - η_ik e_j).
	for k := 0; k < n; k++ {
		if k == j {
			m.Set(i, k, 2*float64(metric[j]))
		}
		if k == i {
			m.Set(j, k, -2*float64(metric[i]))
		}
	}
	return m
}

func auditGenericKernelParity() GenericBivectorParityAudit {
	return GenericBivectorParityAudit{
		RealBivectorAdjointMatricesSkew: true,
		SkewRankAlwaysEven:              true,
		Dimension8KernelAlwaysEven:      true,
		Exact3DKernelPossible:           false,
		Reason:                          "the adjoint action of any real Clifford bivector on 8_v lies in so(1,7) / skew-adjoint form; real skew matrices have even rank, so in dimension 8 the kernel dimension is even, never exactly 3",
		Verdict:                         "the proposed neutral 3-plane cannot be obtained as the real kernel of a single bivector-adjoint Q_8v matrix; a different representation functor or complex/weight decomposition would be required",
	}
}

func auditEWBivectorRetrieval() EWBivectorRetrievalAudit {
	return EWBivectorRetrievalAudit{
		RequestedGenerators:      []string{"T3L", "Y_phi", "Q=T3L+Y_phi", "Z=T3L-Y_phi"},
		ScalarT3Available:        true,
		ScalarYPhiAvailable:      true,
		T3Grade2BladeDerived:     false,
		YPhiGrade2BladeDerived:   false,
		T3BladeLabel:             "",
		YPhiBladeLabel:           "",
		QBladeDerived:            false,
		ZBladeDerived:            false,
		ManualBladeAssignment:    "T3L ?= e_i∧e_j, Y_phi ?= e_k∧e_l",
		ManualAssignmentRejected: true,
		Obstruction:              "the project has scalar/contact 4x4 bridge matrices for T3L and Y_phi, but not their native Cl(1,7) grade-2 blade representatives on the 8_v carrier",
		Verdict:                  "electroweak generators are not yet retrievable as Clifford bivectors, so Q_8v and Z_8v cannot be constructed as derived matrices",
	}
}

func auditExplicitMatrices(simple SimpleBivectorMatrixAudit, parity GenericBivectorParityAudit, ew EWBivectorRetrievalAudit) ExplicitMatrixAudit {
	return ExplicitMatrixAudit{
		AdjointActionFormula:       "R(B)v=[B,v] for explicit B∈Λ²R⁸⊂Cl(1,7)",
		CanConstructCandidateBlade: simple.MatrixRows == 8 && simple.MatrixCols == 8 && simple.SkewSymmetric,
		CanConstructT3Matrix:       ew.T3Grade2BladeDerived,
		CanConstructYPhiMatrix:     ew.YPhiGrade2BladeDerived,
		Q8VConstructed:             ew.QBladeDerived,
		Z8VConstructed:             ew.ZBladeDerived,
		Q8VRows:                    0,
		Q8VCols:                    0,
		Q8VSkewIfBivector:          true,
		Q8VKernelDimensionKnown:    false,
		Q8VKernelDimension:         0,
		Neutral3PlaneDerived:       false,
		BindingFailure:             fmt.Sprintf("no T3/Y_phi grade-2 blades are derived; moreover %s", parity.Reason),
		Verdict:                    "candidate Clifford-bivector matrices are computable, but the electroweak Q_8v matrix is not derived and a real-bivector kernel cannot be exactly three-dimensional",
	}
}

func auditScalarNeutralPlane(m ExplicitMatrixAudit) ScalarNeutralPlaneAudit {
	return ScalarNeutralPlaneAudit{
		NeutralTraceSlots:       []string{"Q^TQ", "Z^TZ", "T3L^T Y_phi"},
		TauEta:                  []int{2, -2, 1},
		NeedsQ8VKernel:          true,
		Q8VKernelAvailable:      m.Q8VConstructed && m.Q8VKernelDimensionKnown,
		KernelDimensionExactly3: false,
		CanonicalIsomorphism:    false,
		VTauConstructed:         false,
		Reason:                  "without a derived Q_8v neutral kernel, there is no coordinate-free neutral 3-plane and no lawful host for tau_eta as a vector representative",
		Verdict:                 "the Gate 249 neutral-kernel route remains blocked after the Clifford-adjoint audit",
	}
}

func auditFirewall() FirewallAudit {
	return FirewallAudit{Verdict: "firewall preserved: no electroweak bivectors, charges, kernels, v_tau, triality maps, or Yukawa textures were invented"}
}

func summarize(prev InheritedGate249Audit, carrier CliffordCarrierAudit, simple SimpleBivectorMatrixAudit, parity GenericBivectorParityAudit, ew EWBivectorRetrievalAudit, matrices ExplicitMatrixAudit, scalar ScalarNeutralPlaneAudit) Summary {
	status := strings.Join([]string{
		StatusCliffordAdjointAvailable,
		StatusCandidateBivectorMatrices,
		StatusSkewKernelParityObstruction,
		StatusEWBivectorRetrievalBlocked,
		StatusQ8VMatrixBlocked,
		StatusNeutralKernel3PlaneBlocked,
		StatusScalarToNeutralPlaneStillBlock,
		StatusTrialityStillBlocked,
		StatusYukawaStillBlocked,
	}, "\n")
	return Summary{
		CliffordAdjointAvailable:    carrier.Grade1CarrierKnown && carrier.CommutatorActionTyped,
		CandidateMatricesComputable: simple.MatrixRows == 8 && simple.MatrixCols == 8,
		EWBivectorsRetrieved:        ew.T3Grade2BladeDerived && ew.YPhiGrade2BladeDerived,
		Q8VMatrixDerived:            matrices.Q8VConstructed,
		NeutralKernelDerived:        scalar.Q8VKernelAvailable,
		NeutralKernelDim3:           scalar.KernelDimensionExactly3,
		RealBivector3KernelPossible: parity.Exact3DKernelPossible,
		VTauConstructed:             scalar.VTauConstructed,
		TrialityUnblocked:           prev.TrialityUnblocked && scalar.VTauConstructed,
		YukawaTextureDerived:        false,
		Status:                      status,
		NextGate:                    "Gate 251 — EW derivation representation functor audit / complex weight-space route beyond real-bivector kernels",
		Comment:                     "Gate 250 gives the Clifford commutator matrix machinery, but it also shows that the requested 3D real neutral kernel is not available from a real bivector adjoint action; the EW action on 8_v remains un-derived.",
	}
}

func buildTruth(prev InheritedGate249Audit, carrier CliffordCarrierAudit, simple SimpleBivectorMatrixAudit, parity GenericBivectorParityAudit, ew EWBivectorRetrievalAudit, matrices ExplicitMatrixAudit, scalar ScalarNeutralPlaneAudit) string {
	return fmt.Sprintf("Gate 250 verifies the lawful Clifford commutator action R(B)v=[B,v] for explicit grade-2 blades on 8_v and constructs a diagnostic simple-bivector 8x8 matrix with rank %d and kernel %d. However, T3L and Y_phi are still only scalar/contact bridge generators, not derived Cl(1,7) bivectors, so Q_8v and Z_8v are not constructed. More strongly, a real Clifford-bivector adjoint action has even-dimensional kernel on 8_v, so the exact 3D neutral-kernel strategy cannot be completed by this real-bivector route alone. The scalar-to-neutral-plane, v_tau, Spin(8) triality, and Yukawa texture derivations remain blocked.", simple.Rank, simple.KernelDimension)
}

func isSkew(m linear.Matrix, eps float64) bool {
	if m.Rows() != m.Cols() {
		return false
	}
	for r := 0; r < m.Rows(); r++ {
		for c := 0; c < m.Cols(); c++ {
			if math.Abs(m.At(r, c)+m.At(c, r)) > eps {
				return false
			}
		}
	}
	return true
}

func rank(m linear.Matrix, eps float64) int {
	rows, cols := m.Rows(), m.Cols()
	a := make([][]float64, rows)
	for r := 0; r < rows; r++ {
		a[r] = make([]float64, cols)
		for c := 0; c < cols; c++ {
			a[r][c] = m.At(r, c)
		}
	}
	rank := 0
	for col := 0; col < cols && rank < rows; col++ {
		pivot := rank
		for r := rank + 1; r < rows; r++ {
			if math.Abs(a[r][col]) > math.Abs(a[pivot][col]) {
				pivot = r
			}
		}
		if math.Abs(a[pivot][col]) <= eps {
			continue
		}
		a[rank], a[pivot] = a[pivot], a[rank]
		pv := a[rank][col]
		for c := col; c < cols; c++ {
			a[rank][c] /= pv
		}
		for r := 0; r < rows; r++ {
			if r == rank {
				continue
			}
			factor := a[r][col]
			if math.Abs(factor) <= eps {
				continue
			}
			for c := col; c < cols; c++ {
				a[r][c] -= factor * a[rank][c]
			}
		}
		rank++
	}
	return rank
}

func nonzeroEntries(m linear.Matrix) []string {
	out := []string{}
	for r := 0; r < m.Rows(); r++ {
		for c := 0; c < m.Cols(); c++ {
			v := m.At(r, c)
			if math.Abs(v) > 1e-12 {
				out = append(out, fmt.Sprintf("M[%d,%d]=%g", r, c, v))
			}
		}
	}
	return out
}
