// Package bfcurvature implements the first genuine finite BF/Maurer-Cartan
// curvature diagnostic on the Boolean-compressed block connection.
//
// Earlier gates showed that the naive compressed generators do not close on the
// contact complement and that the second-fundamental curvature does not act on
// the protected three-dimensional generation carrier. This package tests a
// different object: the Maurer-Cartan residual of the full Boolean-compressed
// connection relative to its seed gauge span.
//
// For Boolean-compressed generators A_i, define
//
//	F_ij = [A_i,A_j] - Pi_seed([A_i,A_j]),
//
// where Pi_seed projects onto span{A_i}. If the compression were an exact flat
// finite gauge representation, these residuals would vanish. Nonzero residuals
// are genuine finite field strength of the compressed connection, not the
// second-fundamental identity alone. We then restrict F_ij to the protected
// generation carrier G and the active Higgs/contact carrier H.
package bfcurvature

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/linear"
	"github.com/bagherbal/asha-engine/pkg/matter/gencurvature"
)

type Operator struct {
	Pair              string
	Matrix            linear.Matrix
	Norm              float64
	OffDiagonalNorm   float64
	SkewResidual      float64
	SymmetricResidual float64
}

type Analysis struct {
	GenerationCurvature gencurvature.Analysis

	GeneratorCount     int
	SeedSpanRank       int
	FullDimension      int
	ProtectedDimension int
	ActiveDimension    int

	FullCurvatures        []Operator
	FullMaxNorm           float64
	FullCurvatureSpanRank int

	ProtectedCurvatures         []Operator // G^T F_ij G.
	ProtectedMaxNorm            float64
	ProtectedMaxOffDiagonalNorm float64
	ProtectedSpanRank           int

	ActiveCurvatures []Operator // H^T F_ij H.
	ActiveMaxNorm    float64
	ActiveSpanRank   int

	CrossCurvatures []Operator // G^T F_ij H.
	CrossMaxNorm    float64
	CrossSpanRank   int

	NonDiagonalGenerationMixingFound bool
	ActiveToGenerationBridgeFound    bool
	CanonicalTextureFound            bool

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
		gc, err := gencurvature.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(gc, 1e-8)
	})
	return defaultValue, defaultErr
}

func Build(gc gencurvature.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-8
	}
	generators := gc.Connection.Compression.BooleanGenerators
	if len(generators) == 0 {
		return Analysis{}, fmt.Errorf("empty Boolean generator list")
	}
	seedFrame, seedRank, err := orthonormalOperatorFrame(generators, eps)
	if err != nil {
		return Analysis{}, err
	}

	fullOps := make([]Operator, 0)
	protectedOps := make([]Operator, 0)
	activeOps := make([]Operator, 0)
	crossOps := make([]Operator, 0)
	maxFull, maxProt, maxProtOff, maxActive, maxCross := 0.0, 0.0, 0.0, 0.0, 0.0

	for i := 0; i < len(generators); i++ {
		for j := i + 1; j < len(generators); j++ {
			br, err := commutator(generators[i], generators[j])
			if err != nil {
				return Analysis{}, err
			}
			proj, err := projectOperator(br, seedFrame)
			if err != nil {
				return Analysis{}, err
			}
			f, err := br.Sub(proj)
			if err != nil {
				return Analysis{}, err
			}
			full, err := summarize(fmt.Sprintf("[%d,%d]", i, j), f)
			if err != nil {
				return Analysis{}, err
			}
			fullOps = append(fullOps, full)
			if full.Norm > maxFull {
				maxFull = full.Norm
			}

			pg, err := restrict(gc.CarrierFrame, f, gc.CarrierFrame)
			if err != nil {
				return Analysis{}, err
			}
			po, err := summarize(full.Pair, pg)
			if err != nil {
				return Analysis{}, err
			}
			protectedOps = append(protectedOps, po)
			if po.Norm > maxProt {
				maxProt = po.Norm
			}
			if po.OffDiagonalNorm > maxProtOff {
				maxProtOff = po.OffDiagonalNorm
			}

			ah, err := restrict(gc.ActiveFrame, f, gc.ActiveFrame)
			if err != nil {
				return Analysis{}, err
			}
			ao, err := summarize(full.Pair, ah)
			if err != nil {
				return Analysis{}, err
			}
			activeOps = append(activeOps, ao)
			if ao.Norm > maxActive {
				maxActive = ao.Norm
			}

			cr, err := restrict(gc.CarrierFrame, f, gc.ActiveFrame)
			if err != nil {
				return Analysis{}, err
			}
			co, err := summarize(full.Pair, cr)
			if err != nil {
				return Analysis{}, err
			}
			crossOps = append(crossOps, co)
			if co.Norm > maxCross {
				maxCross = co.Norm
			}
		}
	}

	fullRank, err := operatorSpanRank(fullOps, eps)
	if err != nil {
		return Analysis{}, err
	}
	protRank, err := operatorSpanRank(protectedOps, eps)
	if err != nil {
		return Analysis{}, err
	}
	activeRank, err := operatorSpanRank(activeOps, eps)
	if err != nil {
		return Analysis{}, err
	}
	crossRank, err := operatorSpanRank(crossOps, eps)
	if err != nil {
		return Analysis{}, err
	}

	nonDiag := maxProtOff > eps && protRank > 0
	crossFound := maxCross > eps && crossRank > 0
	canonical := nonDiag || crossFound

	return Analysis{
		GenerationCurvature:              gc,
		GeneratorCount:                   len(generators),
		SeedSpanRank:                     seedRank,
		FullDimension:                    generators[0].Rows(),
		ProtectedDimension:               gc.CarrierDimension,
		ActiveDimension:                  gc.ActiveDimension,
		FullCurvatures:                   fullOps,
		FullMaxNorm:                      maxFull,
		FullCurvatureSpanRank:            fullRank,
		ProtectedCurvatures:              protectedOps,
		ProtectedMaxNorm:                 maxProt,
		ProtectedMaxOffDiagonalNorm:      maxProtOff,
		ProtectedSpanRank:                protRank,
		ActiveCurvatures:                 activeOps,
		ActiveMaxNorm:                    maxActive,
		ActiveSpanRank:                   activeRank,
		CrossCurvatures:                  crossOps,
		CrossMaxNorm:                     maxCross,
		CrossSpanRank:                    crossRank,
		NonDiagonalGenerationMixingFound: nonDiag,
		ActiveToGenerationBridgeFound:    crossFound,
		CanonicalTextureFound:            canonical,
		TruthStatement:                   truth(nonDiag, crossFound, maxFull > eps),
		RemainingUnknowns: []string{
			"U-17B-BF-ACTION-SOURCE: couple this finite Maurer-Cartan curvature to a BF/Plebanski-type action instead of treating it only as a residual diagnostic",
			"U-16C-NONCOMMUTING-TEXTURES: if a nonzero protected/cross curvature exists, compare independent pairs for non-commuting 3x3 texture structure",
			"U-16D-YUKAWA-SCALE-BRIDGE: convert dimensionless texture spectra into physical Yukawa strengths without observed-mass fitting",
		},
	}, nil
}

func orthonormalOperatorFrame(ops []linear.Matrix, eps float64) (linear.Matrix, int, error) {
	if len(ops) == 0 {
		return linear.NewMatrix(0, 0), 0, nil
	}
	rows := ops[0].Rows() * ops[0].Cols()
	raw := linear.NewMatrix(rows, len(ops))
	for c, op := range ops {
		idx := 0
		for r := 0; r < op.Rows(); r++ {
			for k := 0; k < op.Cols(); k++ {
				raw.Set(idx, c, op.At(r, k))
				idx++
			}
		}
	}
	gram, err := raw.Transpose().Mul(raw)
	if err != nil {
		return linear.Matrix{}, 0, err
	}
	eig, err := linear.SymmetricEigenJacobi(gram, 1e-12, 0)
	if err != nil {
		return linear.Matrix{}, 0, err
	}
	rank := 0
	for _, v := range eig.Values {
		if math.Abs(v) > eps {
			rank++
		}
	}
	frame := linear.NewMatrix(rows, rank)
	out := 0
	for i, v := range eig.Values {
		if math.Abs(v) <= eps {
			continue
		}
		norm := math.Sqrt(math.Abs(v))
		for r := 0; r < rows; r++ {
			sum := 0.0
			for c := 0; c < len(ops); c++ {
				sum += raw.At(r, c) * eig.Vectors.At(c, i)
			}
			frame.Set(r, out, sum/norm)
		}
		out++
	}
	return frame, rank, nil
}

func projectOperator(op linear.Matrix, frame linear.Matrix) (linear.Matrix, error) {
	if frame.Cols() == 0 {
		return linear.NewMatrix(op.Rows(), op.Cols()), nil
	}
	v := vectorize(op)
	coeff, err := frame.Transpose().Mul(v)
	if err != nil {
		return linear.Matrix{}, err
	}
	pv, err := frame.Mul(coeff)
	if err != nil {
		return linear.Matrix{}, err
	}
	return unvectorize(pv, op.Rows(), op.Cols()), nil
}

func restrict(leftFrame, op, rightFrame linear.Matrix) (linear.Matrix, error) {
	tmp, err := op.Mul(rightFrame)
	if err != nil {
		return linear.Matrix{}, err
	}
	return leftFrame.Transpose().Mul(tmp)
}

func summarize(pair string, m linear.Matrix) (Operator, error) {
	skewRes := 0.0
	symRes := 0.0
	if m.Rows() == m.Cols() {
		skewResidual, err := m.Add(m.Transpose())
		if err != nil {
			return Operator{}, err
		}
		symResidual, err := m.Sub(m.Transpose())
		if err != nil {
			return Operator{}, err
		}
		skewRes = skewResidual.FrobeniusNorm()
		symRes = symResidual.FrobeniusNorm()
	}
	off := 0.0
	for r := 0; r < m.Rows(); r++ {
		for c := 0; c < m.Cols(); c++ {
			if r != c {
				v := m.At(r, c)
				off += v * v
			}
		}
	}
	return Operator{Pair: pair, Matrix: m, Norm: m.FrobeniusNorm(), OffDiagonalNorm: math.Sqrt(off), SkewResidual: skewRes, SymmetricResidual: symRes}, nil
}

func operatorSpanRank(ops []Operator, eps float64) (int, error) {
	mats := make([]linear.Matrix, len(ops))
	for i, op := range ops {
		mats[i] = op.Matrix
	}
	_, rank, err := orthonormalOperatorFrame(mats, eps)
	return rank, err
}

func vectorize(m linear.Matrix) linear.Matrix {
	v := linear.NewMatrix(m.Rows()*m.Cols(), 1)
	idx := 0
	for r := 0; r < m.Rows(); r++ {
		for c := 0; c < m.Cols(); c++ {
			v.Set(idx, 0, m.At(r, c))
			idx++
		}
	}
	return v
}

func unvectorize(v linear.Matrix, rows, cols int) linear.Matrix {
	m := linear.NewMatrix(rows, cols)
	idx := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			m.Set(r, c, v.At(idx, 0))
			idx++
		}
	}
	return m
}

func commutator(a, b linear.Matrix) (linear.Matrix, error) {
	ab, err := a.Mul(b)
	if err != nil {
		return linear.Matrix{}, err
	}
	ba, err := b.Mul(a)
	if err != nil {
		return linear.Matrix{}, err
	}
	return ab.Sub(ba)
}

func truth(nonDiag, cross, full bool) string {
	switch {
	case nonDiag && cross:
		return "The finite Maurer-Cartan residual supplies both protected generation curvature and an active-to-generation bridge. This is the first genuine candidate source for non-diagonal generation texture; it still needs BF-action and Yukawa-scale bridges."
	case nonDiag:
		return "The finite Maurer-Cartan residual acts non-diagonally on the protected generation carrier. This is a candidate generation texture source, but it still needs independent non-commuting partners and physical normalization."
	case cross:
		return "The finite Maurer-Cartan residual does not act inside the protected carrier, but it creates an active-to-generation bridge. This may source textures only after an additional contraction or action principle is defined."
	case full:
		return "The finite Maurer-Cartan residual is nonzero on the full Boolean support, but it does not project into the protected generation carrier or active-to-generation bridge. It is a real curvature diagnostic, not yet a Yukawa texture source."
	default:
		return "The Boolean-compressed connection is flat relative to its seed span at this diagnostic level. No finite BF curvature source is exposed."
	}
}

func FormatOperators(ops []Operator) string {
	xs := make([]string, 0)
	for _, op := range ops {
		if op.Norm > 1e-10 {
			xs = append(xs, fmt.Sprintf("%s:norm=%.6g,off=%.6g", op.Pair, op.Norm, op.OffDiagonalNorm))
		}
	}
	if len(xs) == 0 {
		return "none"
	}
	return strings.Join(xs, "; ")
}
