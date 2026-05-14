// Package chiraltrace constructs the first finite chiral scalar-bilinear metric
// needed by the native Fierz/NJL program.
//
// Gate 57 exposed the missing tensor behind the symbolic expression
//
//	c_A = <scalar LR projector, J_A⊗J_A>_finite.
//
// This package builds the part that can be constructed without importing
// continuum Fierz tables: the finite scalar left-right bilinear projector
// induced by the already-derived Yukawa-incidence operator
//
//	Y : H_left ⊗ H_scalar -> H_right.
//
// Because the current engine has only selection-rule incidence, the resulting
// metric is a normalized finite Fock/Yukawa trace object, not a full Clifford
// Lorentz trace.  It upgrades the Fierz audit from "no scalar projector" to
// "scalar LR projector exists; signed current-current coefficients still open".
package chiraltrace

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/loopoperator"
	"github.com/bagherbal/asha-engine/pkg/linear"
)

type TraceIdentity struct {
	Name    string
	Formula string
	Proven  bool
}

type Requirement struct {
	Name        string
	Available   bool
	MissingPart string
}

type Analysis struct {
	Loop loopoperator.Analysis

	DomainDimension      int
	RightDimension       int
	ScalarFiberDimension int
	AllowedFiberEntries  int

	NormalizedBilinearMap linear.Matrix // U : H_right -> H_left⊗H_scalar, stored as domain x right
	ScalarLRProjector     linear.Matrix // P_LR = U U^T on H_left⊗H_scalar
	RightMetric           linear.Matrix // U^T U on H_right

	CommonRightRowNormSquared float64
	NormalizationFactor       float64

	ProjectorTrace            float64
	ProjectorRank             int
	ProjectorSymResidual      float64
	ProjectorIdemResidual     float64
	RightMetricResidual       float64
	DomainComplementDimension int

	BilinearMetricConstructed                bool
	ScalarLRProjectorConstructed             bool
	FiniteFockTraceRulesConstructed          bool
	FullCliffordTraceRulesDerived            bool
	CurrentScalarProjectionCoefficientsKnown bool
	GeneratorNormalizationDerived            bool
	AttractiveSignDerived                    bool
	UpDownTieResolved                        bool
	HiddenObservedInputUsed                  bool

	TraceIdentities     []TraceIdentity
	Requirements        []Requirement
	TruthStatement      string
	RecommendedNextGate string
	RemainingUnknowns   []string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		l, err := loopoperator.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(l, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(l loopoperator.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if !l.FiniteYukawaIncidenceOperatorDerived || !l.NativeLoopTraceSkeletonDerived {
		return Analysis{}, fmt.Errorf("Gate 58 requires Gate 53 finite Yukawa incidence and loop trace skeleton")
	}
	if !l.RowNormsEqual || l.MinRightRowNormSquared <= eps {
		return Analysis{}, fmt.Errorf("Gate 58 requires a nonzero common right-row norm; got min=%g max=%g", l.MinRightRowNormSquared, l.MaxRightRowNormSquared)
	}

	rowNorm := 0.5 * (l.MinRightRowNormSquared + l.MaxRightRowNormSquared)
	normFactor := 1.0 / math.Sqrt(rowNorm)
	u := l.Operator.Transpose().Scale(normFactor) // domain x right
	rightMetric, err := u.Transpose().Mul(u)
	if err != nil {
		return Analysis{}, fmt.Errorf("right metric: %w", err)
	}
	p, err := u.Mul(u.Transpose())
	if err != nil {
		return Analysis{}, fmt.Errorf("scalar LR projector: %w", err)
	}

	idR := linear.Identity(l.RightDimension)
	rmDiff, err := rightMetric.Sub(idR)
	if err != nil {
		return Analysis{}, err
	}
	rightMetricResidual := rmDiff.FrobeniusNorm()

	pp, err := p.Mul(p)
	if err != nil {
		return Analysis{}, err
	}
	idem, err := pp.Sub(p)
	if err != nil {
		return Analysis{}, err
	}
	sym, err := p.Sub(p.Transpose())
	if err != nil {
		return Analysis{}, err
	}
	tr, err := p.Trace()
	if err != nil {
		return Analysis{}, err
	}
	eig, err := linear.SymmetricEigenJacobi(p, eps, 0)
	if err != nil {
		return Analysis{}, fmt.Errorf("scalar LR projector spectrum: %w", err)
	}
	rank := linear.RankFromEigenvalues(eig.Values, eps)

	constructed := rightMetricResidual < 1e-8 && idem.FrobeniusNorm() < 1e-8 && sym.FrobeniusNorm() < 1e-8 && rank == l.RightDimension
	identities := []TraceIdentity{
		{Name: "normalized bilinear isometry", Formula: "U^T U = I_R", Proven: rightMetricResidual < 1e-8},
		{Name: "scalar LR projector", Formula: "P_LR = U U^T, P_LR^2=P_LR=P_LR^T", Proven: constructed},
		{Name: "finite trace rank", Formula: "Tr(P_LR)=dim(H_right)=8", Proven: math.Abs(tr-float64(l.RightDimension)) < 1e-8},
		{Name: "loop trace skeleton", Formula: "Tr(Y Y^T)=16 and P_LR normalizes each right channel by 1/sqrt(2)", Proven: math.Abs(rowNorm-2) < 1e-8 && math.Abs(l.RightTrace-16) < 1e-8},
	}
	reqs := []Requirement{
		{Name: "finite scalar LR bilinear projector", Available: constructed, MissingPart: ""},
		{Name: "ordinary finite Fock/Yukawa trace", Available: constructed, MissingPart: ""},
		{Name: "full Clifford/Lorentz trace", Available: false, MissingPart: "gamma/pseudoscalar trace identities and Lorentz scalar contraction are not yet represented"},
		{Name: "u(4) current-generator normalization", Available: false, MissingPart: "relative kinetic trace weights for central, color, B-L, and leptoquark currents remain open"},
		{Name: "current-current scalar projection coefficients", Available: false, MissingPart: "c_A coefficients require applying current generators to the scalar LR projector"},
		{Name: "attractive sign", Available: false, MissingPart: "finite action/propagator sign for scalar channel is not derived"},
	}

	truth := "The finite scalar LR bilinear metric is now constructed at the Fock/Yukawa incidence level: allowed left×scalar states are normalized into an 8-dimensional scalar-bilinear image. This supplies the target projector needed by the Fierz program, but it is not yet the full Clifford/Lorentz trace and it does not compute signed current-current coefficients."

	return Analysis{
		Loop:                                     l,
		DomainDimension:                          l.DomainDimension,
		RightDimension:                           l.RightDimension,
		ScalarFiberDimension:                     l.ScalarFiberDimension,
		AllowedFiberEntries:                      l.AllowedFiberEntries,
		NormalizedBilinearMap:                    u,
		ScalarLRProjector:                        p,
		RightMetric:                              rightMetric,
		CommonRightRowNormSquared:                rowNorm,
		NormalizationFactor:                      normFactor,
		ProjectorTrace:                           tr,
		ProjectorRank:                            rank,
		ProjectorSymResidual:                     sym.FrobeniusNorm(),
		ProjectorIdemResidual:                    idem.FrobeniusNorm(),
		RightMetricResidual:                      rightMetricResidual,
		DomainComplementDimension:                l.DomainDimension - rank,
		BilinearMetricConstructed:                constructed,
		ScalarLRProjectorConstructed:             constructed,
		FiniteFockTraceRulesConstructed:          constructed,
		FullCliffordTraceRulesDerived:            false,
		CurrentScalarProjectionCoefficientsKnown: false,
		GeneratorNormalizationDerived:            false,
		AttractiveSignDerived:                    false,
		UpDownTieResolved:                        false,
		HiddenObservedInputUsed:                  false,
		TraceIdentities:                          identities,
		Requirements:                             reqs,
		TruthStatement:                           truth,
		RecommendedNextGate:                      "Gate 59 — Current Action on Scalar LR Projector / Coefficient Audit",
		RemainingUnknowns: []string{
			"U-20D1B-CLIFFORD-TRACE-RULES: extend the Fock/Yukawa trace to native Clifford/Lorentz trace identities",
			"U-20D1C-GENERATOR-NORMALIZATION: normalize x∧p/u(4) current generators by finite kinetic trace",
			"U-20D1D-CURRENT-ACTION-ON-P_LR: compute <P_LR,J_A⊗J_A> for each current sector",
			"U-20D2-ATTRACTIVE-SIGN: derive scalar-channel sign from finite action/propagator rule",
			"U-20D4-UP-DOWN-SPLITTING: break up/down tie without observed Yukawa input",
		},
	}, nil
}

func FormatTraceIdentities(xs []TraceIdentity) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		state := "open"
		if x.Proven {
			state = "proven"
		}
		parts = append(parts, fmt.Sprintf("%s[%s]: %s", x.Name, state, x.Formula))
	}
	return strings.Join(parts, "; ")
}

func FormatRequirements(xs []Requirement) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		state := "open"
		if x.Available {
			state = "available"
		}
		detail := x.MissingPart
		if detail == "" {
			detail = "constructed"
		}
		parts = append(parts, fmt.Sprintf("%s[%s]: %s", x.Name, state, detail))
	}
	return strings.Join(parts, "; ")
}

func FormatUnknowns(xs []string) string { return strings.Join(xs, "; ") }
