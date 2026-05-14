// Package u1completion implements Gate 99: abelian coefficient / U(1)
// completion selection search.
//
// Gate 98 exposed a one-parameter full electroweak quadratic family
//
//	K(kappa)=K_SU2+kappa(Q-Z)(Q-Z)^T
//
// in the closed full carrier [T1,T2,Z,Q].  The earlier broken-coordinate
// metric-whitening candidate diag(1,1,4) appears at kappa=6.  Gate 99 asks
// whether kappa=6 is selected by finite data, rather than merely reverse-
// engineered from the whitening condition.
package u1completion

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/ewquadratic"
)

type Candidate struct {
	Name       string
	Value      float64
	Derivation string
	Selected   bool
	Reason     string
}

type Analysis struct {
	Gate98 ewquadratic.Analysis

	TargetKappa            float64
	TargetKappaSource      string
	AbelianDirection       []float64
	AbelianDirectionNormSq float64
	NullDirectionBasis     string
	CompletionFamilyTyped  bool
	WhiteningSelectsKappa  bool
	ActionSelectsKappa     bool
	FiniteSecondVariation  bool

	CandidateResonances    []Candidate
	CandidateHitCount      int
	UniqueDerivation       bool
	CountResonanceSelected bool

	KappaPhysical             bool
	GaugeKineticHessianFixed  bool
	PhysicalCouplingsOrMasses bool

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
		g98, err := ewquadratic.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(g98)
	})
	return defaultValue, defaultErr
}

func Build(g98 ewquadratic.Analysis) (Analysis, error) {
	if !g98.FullQuadraticActionFamilyTyped || !g98.Diag114ReachableInFamily {
		return Analysis{}, fmt.Errorf("Gate 99 requires Gate 98 full quadratic family with reachable diag(1,1,4)")
	}

	target := g98.Diag114Kappa
	direction := []float64{0, 0, -1, 1}
	normSq := 0.0
	for _, x := range direction {
		normSq += x * x
	}

	// These are deliberately labelled as resonances, not derivations.  Several
	// independent finite count combinations produce the same number 6, which is
	// exactly why the engine refuses to select kappa=6 by count-matching.
	candidates := []Candidate{
		{Name: "Goldstone×pair multiplicity", Value: 3 * 2, Derivation: "3 broken directions × scalar pair multiplicity 2", Selected: false, Reason: "count resonance only; no action variation attaches this product to the abelian quadratic coefficient"},
		{Name: "active scalar + complex components", Value: 4 + 2, Derivation: "4 real active scalar directions + 2 complex doublet components", Selected: false, Reason: "dimension sum, not a kinetic coefficient"},
		{Name: "u(4) abelian complement count", Value: 16 - 10, Derivation: "u(4) current count 16 minus nonabelian/neutral inventory 10", Selected: false, Reason: "bookkeeping identity, not an abelian Hessian second variation"},
		{Name: "protected+broken", Value: 3 + 3, Derivation: "3 protected contact directions + 3 broken generator directions", Selected: false, Reason: "3+3 resonance, not a U(1) propagator coefficient"},
	}
	hits := 0
	for _, c := range candidates {
		if math.Abs(c.Value-target) < 1e-12 {
			hits++
		}
	}

	truth := "Gate 99 confirms that kappa_U1=6 is the unique value required to recover the earlier diag(1,1,4) broken-coordinate whitening candidate within the Gate 98 family. It also finds multiple finite count resonances that equal 6. Because several unrelated counts produce the same number and none comes from a second variation of a finite gauge/scalar action, the gate rejects kappa=6 as an action-selected physical coefficient. The missing object remains an abelian U(1) completion term derived from finite action data, not a numerological selection of the whitening value."

	return Analysis{
		Gate98:                    g98,
		TargetKappa:               target,
		TargetKappaSource:         "diag(1,1,4) whitening condition in Gate 98 broken-coordinate diagnostic",
		AbelianDirection:          direction,
		AbelianDirectionNormSq:    normSq,
		NullDirectionBasis:        "Q-Z=2Y_phi in basis [T1,T2,Z,Q]",
		CompletionFamilyTyped:     g98.FullQuadraticActionFamilyTyped,
		WhiteningSelectsKappa:     true,
		ActionSelectsKappa:        false,
		FiniteSecondVariation:     false,
		CandidateResonances:       candidates,
		CandidateHitCount:         hits,
		UniqueDerivation:          false,
		CountResonanceSelected:    false,
		KappaPhysical:             false,
		GaugeKineticHessianFixed:  false,
		PhysicalCouplingsOrMasses: false,
		TruthStatement:            truth,
		RejectedClaims: []string{
			"kappa_U1=6 is derived because it whitens the broken image metric",
			"finite count resonances equal to 6 select the abelian completion coefficient",
			"the abelian completion family fixes the U(1) gauge coupling",
			"the current gate derives thetaW, alpha, gY, or W/Z masses",
		},
		RemainingUnknowns: []string{
			"U-18C10A-ABELIAN-KAPPA: derive kappa_U1 from a finite action second variation",
			"U-18C8F-SECOND-VARIATION: compute delta^2 S for the full electroweak quadratic action",
			"U-18C6-SCALAR-KINETIC-ACTION: derive scalar/contact kinetic normalization",
			"U-18C9-PHYSICAL-COUPLINGS: bridge action-selected kinetic terms to g2, gY, thetaW, alpha, and physical W/Z masses",
		},
		RecommendedNextGate: "Gate 100 — Abelian Completion Action Source / Second-Variation Search",
	}, nil
}

func CandidateSummary(candidates []Candidate) string {
	parts := make([]string, 0, len(candidates))
	for _, c := range candidates {
		parts = append(parts, fmt.Sprintf("%s=%.10g (%s)", c.Name, c.Value, c.Derivation))
	}
	return strings.Join(parts, "; ")
}

func Join(xs []string) string { return strings.Join(xs, "; ") }
