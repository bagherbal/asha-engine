// Package connection interprets the failed Boolean compression correctly.
//
// The operators produced by Gate 8 are not a Lie representation after the
// contact complement is projected out. That is not a numerical accident. For a
// projector P and operators A,B, compression A ↦ PAP is a Lie homomorphism only
// when the complementary subspace is invariant. In general,
//
//	P[A,B]P - [PAP,PBP] = PA(1-P)BP - PB(1-P)AP.
//
// The right-hand side is the finite second fundamental curvature: the exact
// contribution of the off-boundary/vacuum-mixing blocks that were discarded by
// strict boundary compression. Physically, this is the place where the finite
// Higgs/Goldstone/vacuum-mixing sector must live, rather than being treated as
// error noise.
package connection

import (
	"sync"

	"fmt"
	"math"

	"github.com/bagherbal/asha-engine/pkg/gauge/lift"
	"github.com/bagherbal/asha-engine/pkg/linear"
)

type Analysis struct {
	Compression lift.Compression

	MaxBlockReconstructionResidual float64
	MaxOffDiagonalNorm             float64
	MaxProjectionDefectNorm        float64
	MaxSecondFundamentalNorm       float64
	MaxCurvatureIdentityResidual   float64
	MaxCurvatureIdentityRelative   float64
}

var (
	connectionDefaultOnce  sync.Once
	connectionDefaultValue Analysis
	connectionDefaultErr   error
)

func BuildDefault() (Analysis, error) {
	connectionDefaultOnce.Do(func() {
		connectionDefaultValue, connectionDefaultErr = buildConnectionDefaultUncached()
	})
	return connectionDefaultValue, connectionDefaultErr
}

func buildConnectionDefaultUncached() (Analysis, error) {
	c, err := lift.BuildDefault()
	if err != nil {
		return Analysis{}, err
	}
	return Build(c)
}

func Build(c lift.Compression) (Analysis, error) {
	p := c.BooleanComplementProjector // finite matter/contact-complement carrier
	q := c.BooleanContactProjector    // protected contact vacuum sector

	maxRecon := 0.0
	maxOffDiag := 0.0
	for _, a := range c.BooleanGenerators {
		pap, paq, qap, qaq, err := blocks(p, q, a)
		if err != nil {
			return Analysis{}, err
		}
		sum, err := pap.Add(paq)
		if err != nil {
			return Analysis{}, err
		}
		sum, err = sum.Add(qap)
		if err != nil {
			return Analysis{}, err
		}
		sum, err = sum.Add(qaq)
		if err != nil {
			return Analysis{}, err
		}
		diff, err := sum.Sub(a)
		if err != nil {
			return Analysis{}, err
		}
		if n := diff.FrobeniusNorm(); n > maxRecon {
			maxRecon = n
		}
		if n := math.Hypot(paq.FrobeniusNorm(), qap.FrobeniusNorm()); n > maxOffDiag {
			maxOffDiag = n
		}
	}

	maxDefect := 0.0
	maxSecond := 0.0
	maxID := 0.0
	maxRel := 0.0
	for i := 0; i < len(c.BooleanGenerators); i++ {
		for j := i + 1; j < len(c.BooleanGenerators); j++ {
			a := c.BooleanGenerators[i]
			b := c.BooleanGenerators[j]

			pap, _, _, _, err := blocks(p, q, a)
			if err != nil {
				return Analysis{}, err
			}
			pbp, _, _, _, err := blocks(p, q, b)
			if err != nil {
				return Analysis{}, err
			}
			compressedBracket, err := commutator(pap, pbp)
			if err != nil {
				return Analysis{}, err
			}

			bracket, err := commutator(a, b)
			if err != nil {
				return Analysis{}, err
			}
			fullProjected, err := sandwich(p, bracket, p)
			if err != nil {
				return Analysis{}, err
			}

			defect, err := fullProjected.Sub(compressedBracket)
			if err != nil {
				return Analysis{}, err
			}
			second, err := secondFundamentalTerm(p, q, a, b)
			if err != nil {
				return Analysis{}, err
			}
			residual, err := defect.Sub(second)
			if err != nil {
				return Analysis{}, err
			}

			defectNorm := defect.FrobeniusNorm()
			secondNorm := second.FrobeniusNorm()
			residualNorm := residual.FrobeniusNorm()
			if defectNorm > maxDefect {
				maxDefect = defectNorm
			}
			if secondNorm > maxSecond {
				maxSecond = secondNorm
			}
			if residualNorm > maxID {
				maxID = residualNorm
			}
			denom := math.Max(defectNorm, secondNorm)
			rel := 0.0
			// Only report a relative identity error when the identity is being
			// tested against a numerically meaningful curvature. For nearly-zero
			// pairs the absolute residual is the reliable diagnostic.
			if denom > 1e-10 {
				rel = residualNorm / denom
			}
			if rel > maxRel {
				maxRel = rel
			}
		}
	}

	return Analysis{
		Compression:                    c,
		MaxBlockReconstructionResidual: maxRecon,
		MaxOffDiagonalNorm:             maxOffDiag,
		MaxProjectionDefectNorm:        maxDefect,
		MaxSecondFundamentalNorm:       maxSecond,
		MaxCurvatureIdentityResidual:   maxID,
		MaxCurvatureIdentityRelative:   maxRel,
	}, nil
}

func blocks(p, q, a linear.Matrix) (pap, paq, qap, qaq linear.Matrix, err error) {
	pap, err = sandwich(p, a, p)
	if err != nil {
		return
	}
	paq, err = sandwich(p, a, q)
	if err != nil {
		return
	}
	qap, err = sandwich(q, a, p)
	if err != nil {
		return
	}
	qaq, err = sandwich(q, a, q)
	return
}

func secondFundamentalTerm(p, q, a, b linear.Matrix) (linear.Matrix, error) {
	aqb, err := triple(a, q, b)
	if err != nil {
		return linear.Matrix{}, err
	}
	paqbp, err := sandwich(p, aqb, p)
	if err != nil {
		return linear.Matrix{}, err
	}
	bqa, err := triple(b, q, a)
	if err != nil {
		return linear.Matrix{}, err
	}
	pbqap, err := sandwich(p, bqa, p)
	if err != nil {
		return linear.Matrix{}, err
	}
	return paqbp.Sub(pbqap)
}

func sandwich(left, middle, right linear.Matrix) (linear.Matrix, error) {
	lm, err := left.Mul(middle)
	if err != nil {
		return linear.Matrix{}, err
	}
	return lm.Mul(right)
}

func triple(a, b, c linear.Matrix) (linear.Matrix, error) {
	ab, err := a.Mul(b)
	if err != nil {
		return linear.Matrix{}, err
	}
	return ab.Mul(c)
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

func ValidateProjectionPair(p, q linear.Matrix, eps float64) error {
	if p.Rows() != p.Cols() || q.Rows() != q.Cols() || p.Rows() != q.Rows() {
		return fmt.Errorf("projectors must be square with matching dimensions")
	}
	pp, err := p.Mul(p)
	if err != nil {
		return err
	}
	qq, err := q.Mul(q)
	if err != nil {
		return err
	}
	pq, err := p.Mul(q)
	if err != nil {
		return err
	}
	sum, err := p.Add(q)
	if err != nil {
		return err
	}
	if d, _ := pp.Sub(p); d.FrobeniusNorm() > eps {
		return fmt.Errorf("P is not projective")
	}
	if d, _ := qq.Sub(q); d.FrobeniusNorm() > eps {
		return fmt.Errorf("Q is not projective")
	}
	if pq.FrobeniusNorm() > eps {
		return fmt.Errorf("P and Q are not orthogonal")
	}
	if d, _ := sum.Sub(linear.Identity(p.Rows())); d.FrobeniusNorm() > eps {
		return fmt.Errorf("P+Q is not identity")
	}
	return nil
}
