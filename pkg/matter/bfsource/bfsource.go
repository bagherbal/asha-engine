// Package bfsource tests whether a BF/Plebanski-style action contraction can
// turn the finite Maurer-Cartan curvature into a generation texture.
//
// Gate 32 showed that the finite curvature residual is real on the Boolean
// support, but its direct protected and active-to-generation restrictions vanish:
//
//	G^T F G = 0,
//	G^T F H = 0,
//
// while the active Higgs/contact restriction H^T F H is nonzero. This package
// asks whether an action-level contraction can change that conclusion. A finite
// BF term has the schematic form Tr(BF). If the source tensor B is supported on
// the generation carrier, the available response space is exactly G^T F G. If B
// is mixed between generation and active scalar carriers, the response is G^T F H.
// Thus the BF source audit is a clean structural test: only nonzero protected or
// mixed response spaces can produce a canonical generation texture. Active-only
// curvature may contribute to scalar/Higgs dynamics, but it is not a 3x3 flavor
// texture by itself.
package bfsource

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/linear"
	"github.com/bagherbal/asha-engine/pkg/matter/bfcurvature"
)

type QuadraticTexture struct {
	Name            string
	Matrix          linear.Matrix
	Rank            int
	Norm            float64
	Trace           float64
	OffDiagonalNorm float64
	Eigenvalues     []float64
}

type Analysis struct {
	Curvature bfcurvature.Analysis

	FullCurvatureRank int
	FullCurvatureNorm float64

	ProtectedBFResponseRank int
	ProtectedBFMaxNorm      float64
	ProtectedBFMaxOffDiag   float64

	MixedBFResponseRank int
	MixedBFMaxNorm      float64

	ActiveBFResponseRank int
	ActiveBFMaxNorm      float64

	ProtectedQuadratic QuadraticTexture // sum (G^T F G)^T(G^T F G), 3x3.
	MixedQuadratic     QuadraticTexture // sum (G^T F H)(G^T F H)^T, 3x3.
	ActiveQuadratic    QuadraticTexture // sum (H^T F H)^T(H^T F H), 4x4.

	ProtectedBFTextureFound bool
	MixedBFBridgeFound      bool
	ActiveOnlySourceFound   bool
	CanonicalTextureFound   bool

	TruthStatement    string
	RemainingUnknowns []string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		bc, err := bfcurvature.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(bc, 1e-8)
	})
	return defaultValue, defaultErr
}

func Build(bc bfcurvature.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-8
	}

	protMats := matricesFromOperators(bc.ProtectedCurvatures)
	mixedMats := matricesFromOperators(bc.CrossCurvatures)
	activeMats := matricesFromOperators(bc.ActiveCurvatures)

	protRank, err := matrixSpanRank(protMats, eps)
	if err != nil {
		return Analysis{}, err
	}
	mixedRank, err := matrixSpanRank(mixedMats, eps)
	if err != nil {
		return Analysis{}, err
	}
	activeRank, err := matrixSpanRank(activeMats, eps)
	if err != nil {
		return Analysis{}, err
	}

	protectedQ, err := quadratic("protected BF generation texture Σ(GᵀFG)ᵀ(GᵀFG)", protMats, eps, true)
	if err != nil {
		return Analysis{}, err
	}
	mixedQ, err := quadratic("mixed BF generation texture Σ(GᵀFH)(GᵀFH)ᵀ", mixedMats, eps, false)
	if err != nil {
		return Analysis{}, err
	}
	activeQ, err := quadratic("active BF scalar texture Σ(HᵀFH)ᵀ(HᵀFH)", activeMats, eps, true)
	if err != nil {
		return Analysis{}, err
	}

	protectedFound := protRank > 0 && bc.ProtectedMaxNorm > eps
	mixedFound := mixedRank > 0 && bc.CrossMaxNorm > eps
	activeFound := activeRank > 0 && bc.ActiveMaxNorm > eps
	canonical := (protectedFound || mixedFound) && (protectedQ.Rank > 0 || mixedQ.Rank > 0)

	return Analysis{
		Curvature:               bc,
		FullCurvatureRank:       bc.FullCurvatureSpanRank,
		FullCurvatureNorm:       bc.FullMaxNorm,
		ProtectedBFResponseRank: protRank,
		ProtectedBFMaxNorm:      bc.ProtectedMaxNorm,
		ProtectedBFMaxOffDiag:   bc.ProtectedMaxOffDiagonalNorm,
		MixedBFResponseRank:     mixedRank,
		MixedBFMaxNorm:          bc.CrossMaxNorm,
		ActiveBFResponseRank:    activeRank,
		ActiveBFMaxNorm:         bc.ActiveMaxNorm,
		ProtectedQuadratic:      protectedQ,
		MixedQuadratic:          mixedQ,
		ActiveQuadratic:         activeQ,
		ProtectedBFTextureFound: protectedFound,
		MixedBFBridgeFound:      mixedFound,
		ActiveOnlySourceFound:   activeFound,
		CanonicalTextureFound:   canonical,
		TruthStatement:          truth(protectedFound, mixedFound, activeFound),
		RemainingUnknowns: []string{
			"U-17C-BF-SOURCE-SELECTION: derive a nonzero source tensor coupling generations to finite curvature instead of testing only natural support contractions",
			"U-17D-ACTIVE-GENERATION-MAP: derive a canonical nonzero map from active Higgs/contact curvature into the 3D generation carrier",
			"U-16C-NONCOMMUTING-TEXTURES: obtain at least two non-commuting symmetric 3x3 operators before CKM/PMNS claims",
			"U-16D-YUKAWA-SCALE-BRIDGE: normalize any future texture spectra into physical Yukawa strengths without observed-mass fitting",
		},
	}, nil
}

func matricesFromOperators(ops []bfcurvature.Operator) []linear.Matrix {
	out := make([]linear.Matrix, len(ops))
	for i, op := range ops {
		out[i] = op.Matrix
	}
	return out
}

func quadratic(name string, mats []linear.Matrix, eps float64, transposeLeft bool) (QuadraticTexture, error) {
	if len(mats) == 0 {
		return QuadraticTexture{Name: name, Matrix: linear.NewMatrix(0, 0)}, nil
	}
	rows := mats[0].Rows()
	cols := mats[0].Cols()
	size := cols
	if !transposeLeft {
		size = rows
	}
	q := linear.NewMatrix(size, size)
	for _, m := range mats {
		var term linear.Matrix
		var err error
		if transposeLeft {
			term, err = m.Transpose().Mul(m)
		} else {
			term, err = m.Mul(m.Transpose())
		}
		if err != nil {
			return QuadraticTexture{}, err
		}
		q, err = q.Add(term)
		if err != nil {
			return QuadraticTexture{}, err
		}
	}
	tr, err := q.Trace()
	if err != nil {
		return QuadraticTexture{}, err
	}
	off := offDiagonalNorm(q)
	eigValues := []float64(nil)
	rank := 0
	if q.Rows() == q.Cols() && q.Rows() > 0 {
		eig, err := linear.SymmetricEigenJacobi(q, 1e-12, 0)
		if err != nil {
			return QuadraticTexture{}, err
		}
		eigValues = append(eigValues, eig.Values...)
		rank = linear.RankFromEigenvalues(eig.Values, eps*eps)
	}
	return QuadraticTexture{Name: name, Matrix: q, Rank: rank, Norm: q.FrobeniusNorm(), Trace: tr, OffDiagonalNorm: off, Eigenvalues: eigValues}, nil
}

func matrixSpanRank(mats []linear.Matrix, eps float64) (int, error) {
	if len(mats) == 0 {
		return 0, nil
	}
	rows := mats[0].Rows()
	cols := mats[0].Cols()
	raw := linear.NewMatrix(rows*cols, len(mats))
	for c, m := range mats {
		if m.Rows() != rows || m.Cols() != cols {
			return 0, fmt.Errorf("matrix span requires uniform dimensions")
		}
		idx := 0
		for r := 0; r < rows; r++ {
			for k := 0; k < cols; k++ {
				raw.Set(idx, c, m.At(r, k))
				idx++
			}
		}
	}
	gram, err := raw.Transpose().Mul(raw)
	if err != nil {
		return 0, err
	}
	eig, err := linear.SymmetricEigenJacobi(gram, 1e-12, 0)
	if err != nil {
		return 0, err
	}
	return linear.RankFromEigenvalues(eig.Values, eps*eps), nil
}

func offDiagonalNorm(m linear.Matrix) float64 {
	sum := 0.0
	for r := 0; r < m.Rows(); r++ {
		for c := 0; c < m.Cols(); c++ {
			if r == c {
				continue
			}
			v := m.At(r, c)
			sum += v * v
		}
	}
	return math.Sqrt(sum)
}

func truth(protected, mixed, active bool) string {
	switch {
	case protected || mixed:
		return "The BF action contraction exposes a nonzero generation response. This is a candidate texture source, but it still needs symmetry, normalization, and non-commuting partner tests."
	case active:
		return "The finite BF/Maurer-Cartan curvature has a nonzero active Higgs/contact action response, but all natural protected and mixed BF source contractions vanish. The action-level contraction still does not induce a 3x3 generation texture."
	default:
		return "No nonzero BF source contraction is exposed by the current finite curvature data."
	}
}

func FormatTexture(q QuadraticTexture) string {
	if q.Matrix.Rows() == 0 {
		return "empty"
	}
	xs := make([]string, 0, len(q.Eigenvalues))
	for _, v := range q.Eigenvalues {
		if math.Abs(v) > 1e-12 {
			xs = append(xs, fmt.Sprintf("%.6g", v))
		}
	}
	if len(xs) == 0 {
		xs = append(xs, "zero")
	}
	return fmt.Sprintf("rank=%d,norm=%.6g,trace=%.6g,off=%.6g,eig=[%s]", q.Rank, q.Norm, q.Trace, q.OffDiagonalNorm, strings.Join(xs, ","))
}
