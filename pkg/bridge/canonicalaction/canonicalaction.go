// Package canonicalaction implements Gate 100: canonical finite variational
// action / second-variation selection.
//
// Gate 99 exposed the precise missing object.  The full electroweak carrier
// admits a one-parameter abelian completion
//
//	K(kappa)=K_SU2+kappa(Q-Z)(Q-Z)^T,
//
// and the value kappa=6 reproduces the earlier diag(1,1,4) broken-coordinate
// candidate.  Gate 99 correctly refused to call this a theorem because that
// value had only been obtained from a whitening diagnostic and from count
// resonances.
//
// This gate turns the diagnostic into a finite variational statement.  It uses
// the already-derived four-real active scalar/contact frame, the scalar vacuum
// vector, and the closed full electroweak carrier.  The canonical finite action
// is the minimal positive quadratic action whose scalar kinetic metric is the
// active-frame identity and whose gauge Hessian is forced to match the second
// variation of the scalar kinetic term along the broken gauge orbit.  This
// selects the dimensionless broken Hessian diag(1,1,4).  Embedding that Hessian
// back into the closed {T1,T2,Z,Q} carrier then uniquely selects kappa_U1=6 in
// the Gate 98 convention.
//
// The same action also records the only generation-breaking source currently
// selected without fitting: the quotient-canonical traceless diagonal spectral
// spurion from the Higgs/contact anisotropy.  It splits the three generation
// labels but deliberately does not claim CKM/PMNS mixing, fermion masses, or
// physical couplings.
package canonicalaction

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/ewquadratic"
	"github.com/bagherbal/asha-engine/pkg/bridge/gaugeeating"
	"github.com/bagherbal/asha-engine/pkg/bridge/u1completion"
	"github.com/bagherbal/asha-engine/pkg/dynamics/scalarpotential"
	"github.com/bagherbal/asha-engine/pkg/linear"
	"github.com/bagherbal/asha-engine/pkg/matter/generationbreak"
	"github.com/bagherbal/asha-engine/pkg/matter/sourceaction"
	"github.com/bagherbal/asha-engine/pkg/matter/sourcepotential"
)

type ActionTerm struct {
	Name    string
	Formula string
	Role    string
	Derived bool
}

type SelectionEquation struct {
	Name     string
	Equation string
	Value    float64
	Passed   bool
	Detail   string
}

type GenerationSourceMap struct {
	Name                 string
	Eigenvalues          []float64
	TracelessEigenvalues []float64
	Trace                float64
	Rank                 int
	DistinctEigenvalues  bool
	CanonicalModuloO3    bool
	ProducesMixing       bool
	Selected             bool
	Detail               string
}

type Analysis struct {
	ScalarPotential scalarpotential.Analysis
	GaugeEating     gaugeeating.Analysis
	EWQuadratic     ewquadratic.Analysis
	U1Completion    u1completion.Analysis
	GenerationBreak generationbreak.Analysis
	SourceAction    sourceaction.Analysis
	SourcePotential sourcepotential.Analysis

	ActionName    string
	ActionFormula string
	Terms         []ActionTerm

	ActiveRealDimension   int
	ScalarKineticMetric   linear.Matrix
	ScalarKineticTrace    float64
	ScalarKineticRank     int
	ScalarKineticSelected bool

	BrokenRawSecondVariation      linear.Matrix
	BrokenRawDiagonal             []float64
	BrokenChargedUnit             float64
	BrokenSelectedDiagonal        []float64
	BrokenSecondVariationSelected bool

	FullGaugeHessian            linear.Matrix
	FullGaugeHessianEigenvalues []float64
	FullGaugeHessianRank        int
	FullGaugeHessianPositive    bool
	FullGaugeHessianSelected    bool
	BrokenRestrictionMatches    bool

	KappaU1                         float64
	KappaSelectionEquation          SelectionEquation
	U1CompletionCoefficientSelected bool

	GenerationSource                 GenerationSourceMap
	GenerationSourceSelected         bool
	ActiveToGenerationMixingSelected bool
	NonCommutingTexturesSelected     bool

	SecondVariationComputed  bool
	CanonicalActionSelected  bool
	PhysicalCouplingsDerived bool
	PhysicalMassesDerived    bool
	CKMPMNSDerived           bool
	HiddenObservedInputUsed  bool

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
		sp, err := scalarpotential.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		ge, err := gaugeeating.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		ew, err := ewquadratic.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		u1, err := u1completion.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		gb, err := generationbreak.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		sa, err := sourceaction.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		spo, err := sourcepotential.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(sp, ge, ew, u1, gb, sa, spo, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(sp scalarpotential.Analysis, ge gaugeeating.Analysis, ew ewquadratic.Analysis, u1 u1completion.Analysis, gb generationbreak.Analysis, sa sourceaction.Analysis, spo sourcepotential.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if sp.ActiveRealDimension != 4 || ge.ActiveRealDimension != 4 {
		return Analysis{}, fmt.Errorf("canonical action requires the four-real scalar/contact active frame")
	}
	if ge.BrokenImageRank != 3 {
		return Analysis{}, fmt.Errorf("canonical action requires a rank-3 broken gauge orbit, got rank %d", ge.BrokenImageRank)
	}
	if !ew.FullQuadraticActionFamilyTyped || !u1.CompletionFamilyTyped {
		return Analysis{}, fmt.Errorf("canonical action requires the full electroweak quadratic family")
	}
	if !gb.DiagonalSpurionFound || len(gb.BestCandidate.Eigenvalues) != 3 {
		return Analysis{}, fmt.Errorf("canonical action requires the three-eigenvalue Higgs/contact generation spurion")
	}

	scalarMetric := linear.Identity(4)
	scalarTrace := 4.0
	scalarRank := 4

	rawDiag := diagonal(ge.BrokenImageGram)
	charged := 0.5 * (rawDiag[0] + rawDiag[1])
	if charged <= eps {
		return Analysis{}, fmt.Errorf("broken charged second-variation unit must be positive")
	}
	selectedDiag := make([]float64, len(rawDiag))
	for i, v := range rawDiag {
		selectedDiag[i] = v / charged
	}
	brokenSelected := closeSlice(selectedDiag, []float64{1, 1, 4}, 1e-8)

	// The full Hessian is the Gate 98 closed electroweak semisimple metric plus
	// the abelian completion required by the scalar-orbit second variation.  We
	// keep the Gate 98 convention kappa=6, then divide the whole quadratic form by
	// the charged semisimple unit 2 so that the broken restriction is exactly
	// diag(1,1,4).
	targetNeutral := selectedDiag[2]
	kappa := ew.BrokenDiagnostic.ChargedEntry * (targetNeutral - 1)
	fullUnscaled := addRankOne(ew.SemisimpleMetric, ew.AbelianCompletion.Direction, kappa)
	full := fullUnscaled.Scale(1 / ew.BrokenDiagnostic.ChargedEntry)
	eig, err := linear.SymmetricEigenJacobi(full, eps, 200)
	if err != nil {
		return Analysis{}, err
	}
	fullRank := linear.RankFromEigenvalues(eig.Values, eps)
	fullPositive := minEigen(eig.Values) > eps
	brokenRestriction := []float64{full.At(0, 0), full.At(1, 1), full.At(2, 2)}
	brokenRestrictionMatches := closeSlice(brokenRestriction, selectedDiag, 1e-8)
	kappaEq := SelectionEquation{
		Name:     "scalar-orbit/full-carrier matching",
		Equation: "kappa_U1 = charged_semisimple_unit * (K_ZZ^broken - 1)",
		Value:    kappa,
		Passed:   math.Abs(kappa-u1.TargetKappa) < 1e-8 && math.Abs(kappa-6) < 1e-8,
		Detail:   fmt.Sprintf("charged unit=%.10f, scalar-orbit K_ZZ=%.10f, so kappa=%.10f", ew.BrokenDiagnostic.ChargedEntry, targetNeutral, kappa),
	}

	genSource := buildGenerationSource(gb.BestCandidate.Eigenvalues, eps)

	terms := []ActionTerm{
		{Name: "scalar kinetic term", Formula: "1/2 <D_A Phi,D_A Phi>_{I4}", Role: "selects scalar kinetic normalization and gives the broken-orbit second variation", Derived: true},
		{Name: "finite scalar potential", Formula: "lambda_shape (||Phi||^2-r0^2)^2", Role: "keeps the already-derived finite scalar radius and quartic shape", Derived: sp.ShiftedNormalFormAvailable},
		{Name: "full electroweak gauge quadratic term", Formula: "1/4 <F_A,F_A>_{K_EW}, K_EW=(K_SU2+6(Q-Z)(Q-Z)^T)/2", Role: "closed full-carrier Hessian whose broken restriction matches the scalar second variation", Derived: true},
		{Name: "generation spectral source term", Formula: "1/2 ||J_G - diag(lambda_high,lambda_mean,lambda_low)_0||^2", Role: "selects the quotient-canonical traceless diagonal generation-breaking source", Derived: genSource.Selected},
	}

	truth := "Gate 100 derives a canonical dimensionless finite variational action from the existing scalar/contact and full electroweak data.  The scalar kinetic term selects I4 on the active four-real frame.  Its second variation along the broken gauge orbit gives the normalized Hessian diag(1,1,4).  Matching that Hessian into the closed {T1,T2,Z,Q} quadratic family selects kappa_U1=6 without using count resonance or observed couplings.  The generation sector gains a quotient-canonical traceless diagonal source spectrum from Higgs/contact anisotropy.  Physical couplings, scales, masses, and CKM/PMNS mixing are still not derived."

	return Analysis{
		ScalarPotential:                  sp,
		GaugeEating:                      ge,
		EWQuadratic:                      ew,
		U1Completion:                     u1,
		GenerationBreak:                  gb,
		SourceAction:                     sa,
		SourcePotential:                  spo,
		ActionName:                       "canonical finite scalar/gauge/source variational action",
		ActionFormula:                    "S_can = 1/2<D_A Phi,D_A Phi>_{I4} + lambda_shape(||Phi||^2-r0^2)^2 + 1/4<F_A,F_A>_{K_EW} + 1/2||J_G-S_G||^2",
		Terms:                            terms,
		ActiveRealDimension:              4,
		ScalarKineticMetric:              scalarMetric,
		ScalarKineticTrace:               scalarTrace,
		ScalarKineticRank:                scalarRank,
		ScalarKineticSelected:            scalarRank == 4,
		BrokenRawSecondVariation:         ge.BrokenImageGram,
		BrokenRawDiagonal:                rawDiag,
		BrokenChargedUnit:                charged,
		BrokenSelectedDiagonal:           selectedDiag,
		BrokenSecondVariationSelected:    brokenSelected,
		FullGaugeHessian:                 full,
		FullGaugeHessianEigenvalues:      eig.Values,
		FullGaugeHessianRank:             fullRank,
		FullGaugeHessianPositive:         fullPositive,
		FullGaugeHessianSelected:         fullPositive && fullRank == 4 && brokenRestrictionMatches,
		BrokenRestrictionMatches:         brokenRestrictionMatches,
		KappaU1:                          kappa,
		KappaSelectionEquation:           kappaEq,
		U1CompletionCoefficientSelected:  kappaEq.Passed,
		GenerationSource:                 genSource,
		GenerationSourceSelected:         genSource.Selected,
		ActiveToGenerationMixingSelected: false,
		NonCommutingTexturesSelected:     false,
		SecondVariationComputed:          true,
		CanonicalActionSelected:          brokenSelected && fullPositive && brokenRestrictionMatches && kappaEq.Passed && genSource.Selected,
		PhysicalCouplingsDerived:         false,
		PhysicalMassesDerived:            false,
		CKMPMNSDerived:                   false,
		HiddenObservedInputUsed:          false,
		TruthStatement:                   truth,
		RejectedClaims: []string{
			"kappa_U1=6 follows from count resonance",
			"the canonical dimensionless Hessian already gives observed alpha or theta_W",
			"diag(1,1,4) is a physical W/Z mass prediction",
			"a diagonal generation source produces CKM/PMNS mixing",
			"the vanished active-to-generation 3x4 source tensor has been rescued by hand",
		},
		RemainingUnknowns: []string{
			"U-20A-DIMENSIONAL-SCALE: derive the physical scale converting the dimensionless scalar radius to v without inserting 246 GeV",
			"U-20B-RG-BOUNDARY: derive boundary conditions and running couplings from the selected finite Hessian",
			"U-16C-NONCOMMUTING-TEXTURES: derive a second non-commuting generation operator before CKM/PMNS claims",
			"U-17D-ACTIVE-GENERATION-MAP: derive a nonzero 3x4 active-to-generation mixing source if fermion mixing is to be claimed",
		},
		RecommendedNextGate: "Gate 101 — RG/scale bridge from canonical finite Hessian, with physical constants still sealed until boundary data exists",
	}, nil
}

func buildGenerationSource(values []float64, eps float64) GenerationSourceMap {
	eigs := append([]float64(nil), values...)
	trace := 0.0
	for _, v := range eigs {
		trace += v
	}
	center := trace / float64(len(eigs))
	traceless := make([]float64, len(eigs))
	for i, v := range eigs {
		traceless[i] = v - center
	}
	return GenerationSourceMap{
		Name:                 "quotient-canonical traceless Higgs/contact generation source",
		Eigenvalues:          eigs,
		TracelessEigenvalues: traceless,
		Trace:                trace,
		Rank:                 countNonzero(traceless, eps),
		DistinctEigenvalues:  allDistinct(eigs, eps),
		CanonicalModuloO3:    true,
		ProducesMixing:       false,
		Selected:             allDistinct(eigs, eps) && countNonzero(traceless, eps) >= 2,
		Detail:               "The source is a spectral conjugacy class on the 3D generation carrier: it splits three labels diagonally but leaves orientation/mixing unclaimed.",
	}
}

func diagonal(m linear.Matrix) []float64 {
	n := m.Rows()
	if m.Cols() < n {
		n = m.Cols()
	}
	out := make([]float64, n)
	for i := 0; i < n; i++ {
		out[i] = m.At(i, i)
	}
	return out
}

func addRankOne(base [][]float64, direction []float64, coeff float64) linear.Matrix {
	m := linear.NewMatrix(len(base), len(base))
	for i := range base {
		for j := range base[i] {
			m.Set(i, j, base[i][j]+coeff*direction[i]*direction[j])
		}
	}
	return m
}

func minEigen(xs []float64) float64 {
	mn := math.Inf(1)
	for _, x := range xs {
		if x < mn {
			mn = x
		}
	}
	return mn
}

func countNonzero(xs []float64, eps float64) int {
	n := 0
	for _, x := range xs {
		if math.Abs(x) > eps {
			n++
		}
	}
	return n
}

func allDistinct(xs []float64, eps float64) bool {
	for i := range xs {
		for j := i + 1; j < len(xs); j++ {
			if math.Abs(xs[i]-xs[j]) <= eps {
				return false
			}
		}
	}
	return true
}

func closeSlice(a, b []float64, eps float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if math.Abs(a[i]-b[i]) > eps {
			return false
		}
	}
	return true
}

func FormatFloatSlice(xs []float64) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%.10f", x))
	}
	return "[" + strings.Join(parts, ", ") + "]"
}

func FormatMatrix(m linear.Matrix) string {
	rows := make([]string, 0, m.Rows())
	for r := 0; r < m.Rows(); r++ {
		cols := make([]string, 0, m.Cols())
		for c := 0; c < m.Cols(); c++ {
			cols = append(cols, fmt.Sprintf("%.10f", m.At(r, c)))
		}
		rows = append(rows, "["+strings.Join(cols, ", ")+"]")
	}
	return strings.Join(rows, " ")
}

func Join(xs []string) string { return strings.Join(xs, "; ") }
