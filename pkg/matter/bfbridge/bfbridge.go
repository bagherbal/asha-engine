// Package bfbridge tests whether active Higgs/contact curvature can be pushed
// canonically into the protected three-dimensional generation carrier.
//
// Gate 30 found a precise no-go: the contact-side second-fundamental curvature
// is flat on the protected K-residual carrier but nonzero on the active
// Higgs/contact carrier. A tempting next move is to project that active
// curvature down to the protected carrier and call the result generation
// mixing. This package refuses that shortcut. It tests the natural bridge maps
// already present in the finite block connection:
//
//	B_i = G^T A_i H,
//
// where G is the protected 3D carrier, H is the active 4D carrier, and A_i are
// the Boolean-compressed contact-preserving connection generators. If those
// maps vanish, then the active curvature cannot canonically induce a 3x3
// generation operator through the existing connection. In that case the engine
// must wait for a real finite BF curvature operator or another bridge theorem.
package bfbridge

import (
	"fmt"
	"math"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/linear"
	"github.com/bagherbal/asha-engine/pkg/matter/gencurvature"
)

type InducedOperator struct {
	Label           string
	Matrix          linear.Matrix
	Norm            float64
	SymmetricNorm   float64
	SkewNorm        float64
	OffDiagonalNorm float64
}

type Analysis struct {
	Curvature gencurvature.Analysis

	ProtectedDimension int
	ActiveDimension    int
	GeneratorCount     int

	CrossMaps        []linear.Matrix // 3x4 maps G^T A_i H.
	MaxCrossMapNorm  float64
	CrossMapSpanRank int

	ActiveCurvatureNorm float64
	ActiveCurvaturePair string
	ActiveCurvatureRank int

	InducedSkewOperators []InducedOperator // B_i F_active B_j^T.
	MaxInducedSkewNorm   float64
	InducedSkewSpanRank  int

	InducedSymmetricOperators []InducedOperator // B_i F_active^T F_active B_j^T.
	MaxInducedSymmetricNorm   float64
	InducedSymmetricSpanRank  int

	ExistingConnectionBridgeFound    bool
	NonDiagonalGenerationMixingFound bool
	BFCurvatureImplemented           bool
	CanonicalGenerationMixingFound   bool

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
	if gc.CarrierDimension != 3 {
		return Analysis{}, fmt.Errorf("protected generation carrier must be 3D, got %d", gc.CarrierDimension)
	}
	if gc.ActiveDimension != 4 {
		return Analysis{}, fmt.Errorf("active Higgs/contact carrier must be 4D, got %d", gc.ActiveDimension)
	}

	generators := gc.Connection.Compression.BooleanGenerators
	cross := make([]linear.Matrix, 0, len(generators))
	maxCross := 0.0
	for _, gen := range generators {
		b, err := restrict(gc.CarrierFrame, gen, gc.ActiveFrame)
		if err != nil {
			return Analysis{}, err
		}
		cross = append(cross, b)
		if n := b.FrobeniusNorm(); n > maxCross {
			maxCross = n
		}
	}
	crossRank, _, err := matrixSpan(cross, eps)
	if err != nil {
		return Analysis{}, err
	}

	activeF, activePair, activeNorm := strongestActiveCurvature(gc)
	activeRank := 0
	if activeF.Rows() > 0 {
		// Rank of a skew curvature is read from F^T F. This avoids a general
		// nonsymmetric eigensolver and is enough for the finite diagnostic.
		ftf, err := activeF.Transpose().Mul(activeF)
		if err != nil {
			return Analysis{}, err
		}
		eig, err := linear.SymmetricEigenJacobi(ftf, 1e-12, 0)
		if err != nil {
			return Analysis{}, err
		}
		activeRank = linear.RankFromEigenvalues(eig.Values, eps*eps)
	}

	inducedSkew := make([]InducedOperator, 0)
	inducedSym := make([]InducedOperator, 0)
	maxSkew := 0.0
	maxSym := 0.0
	if activeF.Rows() > 0 {
		ftf, err := activeF.Transpose().Mul(activeF)
		if err != nil {
			return Analysis{}, err
		}
		for i := range cross {
			for j := range cross {
				label := fmt.Sprintf("B%d F_%s B%d^T", i, activePair, j)
				m, err := sandwichRect(cross[i], activeF, cross[j].Transpose())
				if err != nil {
					return Analysis{}, err
				}
				op, err := summarize(label, m)
				if err != nil {
					return Analysis{}, err
				}
				inducedSkew = append(inducedSkew, op)
				if op.Norm > maxSkew {
					maxSkew = op.Norm
				}

				labelS := fmt.Sprintf("B%d F_%s^T F_%s B%d^T", i, activePair, activePair, j)
				ms, err := sandwichRect(cross[i], ftf, cross[j].Transpose())
				if err != nil {
					return Analysis{}, err
				}
				ops, err := summarize(labelS, ms)
				if err != nil {
					return Analysis{}, err
				}
				inducedSym = append(inducedSym, ops)
				if ops.Norm > maxSym {
					maxSym = ops.Norm
				}
			}
		}
	}
	skewRank, _, err := operatorSpan(inducedSkew, eps)
	if err != nil {
		return Analysis{}, err
	}
	symRank, _, err := operatorSpan(inducedSym, eps)
	if err != nil {
		return Analysis{}, err
	}

	bridgeFound := maxCross > eps && crossRank > 0
	nonDiag := false
	for _, op := range inducedSym {
		if op.OffDiagonalNorm > eps {
			nonDiag = true
			break
		}
	}
	for _, op := range inducedSkew {
		if op.OffDiagonalNorm > eps {
			nonDiag = true
			break
		}
	}

	canonical := bridgeFound && nonDiag && (skewRank > 0 || symRank > 0)

	return Analysis{
		Curvature:                        gc,
		ProtectedDimension:               gc.CarrierDimension,
		ActiveDimension:                  gc.ActiveDimension,
		GeneratorCount:                   len(generators),
		CrossMaps:                        cross,
		MaxCrossMapNorm:                  maxCross,
		CrossMapSpanRank:                 crossRank,
		ActiveCurvatureNorm:              activeNorm,
		ActiveCurvaturePair:              activePair,
		ActiveCurvatureRank:              activeRank,
		InducedSkewOperators:             inducedSkew,
		MaxInducedSkewNorm:               maxSkew,
		InducedSkewSpanRank:              skewRank,
		InducedSymmetricOperators:        inducedSym,
		MaxInducedSymmetricNorm:          maxSym,
		InducedSymmetricSpanRank:         symRank,
		ExistingConnectionBridgeFound:    bridgeFound,
		NonDiagonalGenerationMixingFound: nonDiag,
		BFCurvatureImplemented:           false,
		CanonicalGenerationMixingFound:   canonical,
		TruthStatement:                   truth(bridgeFound, canonical),
		RemainingUnknowns: []string{
			"U-17-BF-CURVATURE: implement the finite BF/Maurer-Cartan curvature operator on K ⊕ K⊥ rather than reusing only second-fundamental curvature",
			"U-17A-ACTIVE-GENERATION-MAP: derive a nonzero canonical map from the 4D active Higgs/contact carrier to the 3D protected generation carrier",
			"U-16C-NONCOMMUTING-TEXTURES: derive at least two non-commuting finite 3x3 texture operators before CKM/PMNS can be claimed",
			"U-16D-YUKAWA-SCALE-BRIDGE: convert finite dimensionless spectra into physical Yukawa coupling strengths without observed-mass fitting",
		},
	}, nil
}

func strongestActiveCurvature(gc gencurvature.Analysis) (linear.Matrix, string, float64) {
	best := linear.NewMatrix(0, 0)
	pair := "none"
	norm := 0.0
	for _, op := range gc.ActiveOperators {
		if op.Norm > norm {
			best = op.Matrix
			pair = op.Pair
			norm = op.Norm
		}
	}
	return best, pair, norm
}

func restrict(leftFrame, op, rightFrame linear.Matrix) (linear.Matrix, error) {
	tmp, err := op.Mul(rightFrame)
	if err != nil {
		return linear.Matrix{}, err
	}
	return leftFrame.Transpose().Mul(tmp)
}

func sandwichRect(a, b, c linear.Matrix) (linear.Matrix, error) {
	ab, err := a.Mul(b)
	if err != nil {
		return linear.Matrix{}, err
	}
	return ab.Mul(c)
}

func summarize(label string, m linear.Matrix) (InducedOperator, error) {
	sym, err := m.Sub(m.Transpose()) // zero iff symmetric
	if err != nil {
		return InducedOperator{}, err
	}
	skew, err := m.Add(m.Transpose()) // zero iff skew
	if err != nil {
		return InducedOperator{}, err
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
	return InducedOperator{Label: label, Matrix: m, Norm: m.FrobeniusNorm(), SymmetricNorm: sym.FrobeniusNorm(), SkewNorm: skew.FrobeniusNorm(), OffDiagonalNorm: math.Sqrt(off)}, nil
}

func matrixSpan(ms []linear.Matrix, eps float64) (int, linear.Matrix, error) {
	if len(ms) == 0 {
		return 0, linear.NewMatrix(0, 0), nil
	}
	ops := make([]InducedOperator, len(ms))
	for i, m := range ms {
		ops[i] = InducedOperator{Matrix: m}
	}
	return operatorSpan(ops, eps)
}

func operatorSpan(ops []InducedOperator, eps float64) (int, linear.Matrix, error) {
	if len(ops) == 0 {
		return 0, linear.NewMatrix(0, 0), nil
	}
	rows := ops[0].Matrix.Rows() * ops[0].Matrix.Cols()
	cols := len(ops)
	raw := linear.NewMatrix(rows, cols)
	for c, op := range ops {
		idx := 0
		for r := 0; r < op.Matrix.Rows(); r++ {
			for k := 0; k < op.Matrix.Cols(); k++ {
				raw.Set(idx, c, op.Matrix.At(r, k))
				idx++
			}
		}
	}
	gram, err := raw.Transpose().Mul(raw)
	if err != nil {
		return 0, linear.Matrix{}, err
	}
	eig, err := linear.SymmetricEigenJacobi(gram, 1e-12, 0)
	if err != nil {
		return 0, linear.Matrix{}, err
	}
	rank := linear.RankFromEigenvalues(eig.Values, eps*eps)
	return rank, raw, nil
}

func truth(bridgeFound, canonical bool) string {
	switch {
	case canonical:
		return "The existing finite block connection supplies a nonzero active-to-protected bridge and induces a non-diagonal generation operator. This is a candidate finite generation-mixing source, still awaiting a Yukawa-scale bridge."
	case bridgeFound:
		return "The existing finite block connection supplies an active-to-protected bridge, but the induced generation operators are still too small or diagonal to select CKM/PMNS-type mixing."
	default:
		return "The active Higgs/contact curvature is real, but the existing compressed connection has no canonical active-to-protected bridge. Therefore it cannot yet induce a 3x3 generation-mixing texture; a genuine finite BF curvature or new projection theorem is still missing."
	}
}
