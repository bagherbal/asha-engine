// Package ewquadratic implements Gate 98: full electroweak quadratic action /
// abelian completion search.
//
// Gate 97 showed that the closed full electroweak carrier {T1,T2,Z,Q} has a
// semisimple adjoint/Killing diagnostic of rank three.  Its null direction is
// Q-Z, i.e. the pure scalar/contact abelian direction 2Y_phi.  Gate 98 adds the
// missing abelian completion as an explicit one-parameter positive term and
// asks whether the finite action selects its coefficient.
//
// The result is intentionally disciplined: a family of positive quadratic
// actions exists.  In the convention K = K_SU2 + kappa (Q-Z)(Q-Z)^T, the broken
// raw-coordinate diagnostic diag(1,1,4) appears at kappa=6 after normalizing the
// charged directions.  But kappa=6 is not selected by the finite action in this
// gate; it is only the value required to reproduce the metric-whitening
// candidate exposed earlier.
package ewquadratic

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/ewcurvature"
)

type QuadraticFamily struct {
	Convention  string
	Direction   []float64
	KappaName   string
	PositiveFor string
}

type BrokenFamily struct {
	ChargedEntry      float64
	NeutralFormula    string
	NormalizedFormula string
	KappaForDiag114   float64
	Diag114Reachable  bool
}

type Analysis struct {
	Gate97 ewcurvature.Analysis

	Variables                      []string
	SemisimpleMetric               [][]float64
	SemisimpleRank                 int
	SemisimpleNull                 []float64
	SemisimplePositiveSemidefinite bool

	AbelianCompletion          QuadraticFamily
	AbelianCompletionTyped     bool
	AbelianCoefficientSelected bool

	BrokenDiagnostic         BrokenFamily
	Diag114ReachableInFamily bool
	Diag114Kappa             float64
	Diag114SelectedByAction  bool

	FullQuadraticActionFamilyTyped bool
	PositiveQuadraticFamilyExists  bool
	GaugeKineticHessianSelected    bool
	PhysicalCouplingsOrMasses      bool

	TruthStatement      string
	RejectedClaims      []string
	RemainingUnknowns   []string
	RecommendedNextGate string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		g97, err := ewcurvature.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(g97)
	})
	return defaultValue, defaultErr
}

func Build(g97 ewcurvature.Analysis) (Analysis, error) {
	if !g97.Closed || !g97.FullFieldStrengthTyped {
		return Analysis{}, fmt.Errorf("Gate 98 requires Gate 97 closed full electroweak field-strength carrier")
	}

	// Flip the sign of Gate 97's compact adjoint/Killing diagnostic to expose a
	// positive-semidefinite semisimple curvature metric.  It remains rank three;
	// its null direction is the pure abelian Q-Z direction.
	ksu2 := make([][]float64, len(g97.AdjointMetric))
	for i := range g97.AdjointMetric {
		ksu2[i] = make([]float64, len(g97.AdjointMetric[i]))
		for j, x := range g97.AdjointMetric[i] {
			ksu2[i][j] = -x
		}
	}
	null := []float64{0, 0, -1, 1}

	// Completion convention: add kappa*(Q-Z)(Q-Z)^T in the basis
	// [T1,T2,Z,Q].  With photon coordinate set to zero for a broken-coordinate
	// diagnostic, the broken entries are [2,2,2+kappa].  Dividing by the charged
	// entry 2 gives [1,1,1+kappa/2].  The previous diag(1,1,4) candidate is
	// therefore recovered at kappa=6, but this gate does not select that value.
	kappaForDiag := 6.0
	broken := BrokenFamily{
		ChargedEntry:      2,
		NeutralFormula:    "2 + kappa",
		NormalizedFormula: "diag(1,1,1+kappa/2)",
		KappaForDiag114:   kappaForDiag,
		Diag114Reachable:  true,
	}

	truth := "Gate 98 completes the Gate 97 semisimple curvature diagnostic by adding the missing abelian quadratic term as a one-parameter family K(kappa)=K_SU2+kappa(Q-Z)(Q-Z)^T. This produces a mathematically valid full electroweak quadratic-action family. In this convention the earlier diag(1,1,4) broken-coordinate candidate is reachable at kappa=6, but the finite action still does not select kappa=6 or any other abelian kinetic coefficient. Therefore the gate exposes the correct abelian-completion problem without deriving physical g2, gY, thetaW, alpha, or W/Z masses."

	return Analysis{
		Gate97:                         g97,
		Variables:                      []string{"T1", "T2", "Z=T3-Y_phi", "Q=T3+Y_phi"},
		SemisimpleMetric:               ksu2,
		SemisimpleRank:                 g97.AdjointRank,
		SemisimpleNull:                 null,
		SemisimplePositiveSemidefinite: true,
		AbelianCompletion: QuadraticFamily{
			Convention:  "K(kappa)=K_SU2+kappa(Q-Z)(Q-Z)^T in basis [T1,T2,Z,Q]",
			Direction:   null,
			KappaName:   "kappa_U1",
			PositiveFor: "kappa_U1 > 0",
		},
		AbelianCompletionTyped:         true,
		AbelianCoefficientSelected:     false,
		BrokenDiagnostic:               broken,
		Diag114ReachableInFamily:       true,
		Diag114Kappa:                   kappaForDiag,
		Diag114SelectedByAction:        false,
		FullQuadraticActionFamilyTyped: true,
		PositiveQuadraticFamilyExists:  true,
		GaugeKineticHessianSelected:    false,
		PhysicalCouplingsOrMasses:      false,
		TruthStatement:                 truth,
		RejectedClaims: []string{
			"the semisimple adjoint metric alone normalizes U(1)",
			"the abelian completion coefficient is selected by closure",
			"diag(1,1,4) is action-derived because it is reachable in the family",
			"the full electroweak quadratic family fixes physical couplings or masses",
		},
		RemainingUnknowns: []string{
			"U-18C10A-ABELIAN-KAPPA: derive kappa_U1 from finite action data rather than metric whitening",
			"U-18C8F-SECOND-VARIATION: compute the second variation of the full electroweak action",
			"U-18C6-SCALAR-KINETIC-ACTION: derive scalar/contact kinetic normalization",
			"U-18C9-PHYSICAL-COUPLINGS: derive g2, gY, thetaW, alpha, and physical W/Z masses through action-selected kinetic terms and scale/RG bridges",
		},
		RecommendedNextGate: "Gate 99 — Abelian Coefficient / U(1) Completion Selection Search",
	}, nil
}

func MatrixString(m [][]float64) string {
	rows := make([]string, 0, len(m))
	for _, r := range m {
		parts := make([]string, 0, len(r))
		for _, x := range r {
			if math.Abs(x) < 1e-12 {
				x = 0
			}
			parts = append(parts, fmt.Sprintf("%.3g", x))
		}
		rows = append(rows, "["+strings.Join(parts, ",")+"]")
	}
	return strings.Join(rows, " ")
}

func Join(xs []string) string { return strings.Join(xs, "; ") }
