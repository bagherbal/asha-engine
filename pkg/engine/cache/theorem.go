// Package cache verifies the engine-level finite-core fixture layer.
//
// This gate is deliberately engineering-facing: it does not introduce a new
// physical theorem. It verifies that the heavy finite objects are stable across
// repeated default construction calls so the theorem ladder can be run as an
// instrument instead of as a one-off script.
package cache

import (
	"fmt"
	"math"

	"github.com/bagherbal/asha-engine/pkg/dynamics/bsector"
	"github.com/bagherbal/asha-engine/pkg/dynamics/higgspotential"
	"github.com/bagherbal/asha-engine/pkg/gauge/boundary"
	"github.com/bagherbal/asha-engine/pkg/gauge/higgs"
	"github.com/bagherbal/asha-engine/pkg/gauge/lift"
	"github.com/bagherbal/asha-engine/pkg/geometry/contact"
	"github.com/bagherbal/asha-engine/pkg/matter/trialityyukawa"
	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func RuntimeFixtureCacheTheorem() theorem.Theorem {
	const id = "ENGINE-FINITE-CORE-FIXTURE-CACHE"
	const name = "finite-core cache and deterministic fixture layer"
	return theorem.Theorem{
		ID: id, Name: name, Layer: theorem.LayerValidation, Status: theorem.VerifiedNumeric,
		Verify: func() theorem.Result {
			checks := make([]theorem.Check, 0)

			c1, err := contact.BuildDefault()
			if err != nil {
				return failed(id, name, "contact fixture", err)
			}
			c2, err := contact.BuildDefault()
			if err != nil {
				return failed(id, name, "contact fixture repeat", err)
			}
			checks = append(checks, theorem.Check{Name: "contact fixture stable", Passed: c1.Dimension() == c2.Dimension() && near(c1.ContactIndex(), c2.ContactIndex()), Detail: fmt.Sprintf("dim K=%d, I_BG=%.10f", c2.Dimension(), c2.ContactIndex())})

			b1, err := bsector.BuildDefault()
			if err != nil {
				return failed(id, name, "B-sector fixture", err)
			}
			b2, err := bsector.BuildDefault()
			if err != nil {
				return failed(id, name, "B-sector fixture repeat", err)
			}
			checks = append(checks, theorem.Check{Name: "B-sector fixture stable", Passed: b1.ZeroModeDimension(1e-8) == b2.ZeroModeDimension(1e-8) && near(b1.FirstPositiveEigenvalue(1e-8), b2.FirstPositiveEigenvalue(1e-8)), Detail: fmt.Sprintf("zero modes=%d, first gap=%.10f", b2.ZeroModeDimension(1e-8), b2.FirstPositiveEigenvalue(1e-8))})

			l1, err := lift.BuildDefault()
			if err != nil {
				return failed(id, name, "Boolean lift fixture", err)
			}
			l2, err := lift.BuildDefault()
			if err != nil {
				return failed(id, name, "Boolean lift fixture repeat", err)
			}
			checks = append(checks, theorem.Check{Name: "Boolean lift fixture stable", Passed: l1.CompressedFrameRank == l2.CompressedFrameRank && near(l1.ClosureRelativeResidual, l2.ClosureRelativeResidual), Detail: fmt.Sprintf("restricted rank=%d, naive closure residual=%.6e", l2.CompressedFrameRank, l2.ClosureRelativeResidual)})

			bd, err := boundary.BuildDefault()
			if err != nil {
				return failed(id, name, "boundary closure fixture", err)
			}
			checks = append(checks, theorem.Check{Name: "boundary closure capped for runtime", Passed: bd.MaxDimension <= 8 && bd.CutoffReached, Detail: fmt.Sprintf("cap=%d, observed dim=%d, cutoff=%t", bd.MaxDimension, bd.ClosureDimension, bd.CutoffReached)})

			h, err := higgs.BuildDefault()
			if err != nil {
				return failed(id, name, "Higgs mixing fixture", err)
			}
			hp, err := higgspotential.BuildDefault()
			if err != nil {
				return failed(id, name, "Higgs potential fixture", err)
			}
			checks = append(checks, theorem.Check{Name: "Higgs/contact fixture stable", Passed: h.HiggsSpanRank == 2 && hp.ActiveContactDimension == 4 && hp.ProtectedContactDimension == 3, Detail: fmt.Sprintf("Φ span=%d, active=%d, protected=%d", h.HiggsSpanRank, hp.ActiveContactDimension, hp.ProtectedContactDimension)})

			ty, err := trialityyukawa.BuildDefault()
			if err != nil {
				return failed(id, name, "triality/Yukawa fixture", err)
			}
			checks = append(checks, theorem.Check{Name: "bridge fixture stable", Passed: ty.GenerationCount == 3 && ty.DiagonalChannelCount == 24 && ty.FullMixingMapCount == 72, Detail: fmt.Sprintf("generations=%d, diagonal channels=%d, full-mixing channels=%d", ty.GenerationCount, ty.DiagonalChannelCount, ty.FullMixingMapCount)})

			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerValidation, Status: theorem.VerifiedNumeric, Checks: checks, Notes: []string{
				"This gate changes no mathematics. It makes the heavy finite-core objects reusable so the full theorem ladder is an executable instrument.",
				"The boundary closure diagnostic is intentionally capped: its purpose is to expose dimension growth, not to spend runtime generating a large algebra at every run.",
			}}
		},
	}
}

func failed(id, name, step string, err error) theorem.Result {
	return theorem.Result{ID: id, Name: name, Layer: theorem.LayerValidation, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: step, Passed: false, Detail: err.Error()}}}
}

func near(a, b float64) bool { return math.Abs(a-b) < 1e-10 }
