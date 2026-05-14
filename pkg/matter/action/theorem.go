package action

import (
	"fmt"
	"math"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func RepresentationActionTheorem() theorem.Theorem {
	const id = "MATTER-FOCK-REPRESENTATION-ACTION"
	const name = "second-quantized action of finite Higgs/contact spectrum on the Fock basis"
	return theorem.Theorem{
		ID:     id,
		Name:   name,
		Layer:  theorem.LayerMatter,
		Status: theorem.BridgeRequired,
		Verify: func() theorem.Result {
			const eps = 1e-8
			a, err := BuildDefault()
			if err != nil {
				return theorem.Result{ID: id, Name: name, Layer: theorem.LayerMatter, Status: theorem.FailedRoute,
					Checks: []theorem.Check{{Name: "construct representation-action bridge", Passed: false, Detail: err.Error()}},
				}
			}

			oneParticleOK := sameMultiset(a.OneParticleResponses, a.OneParticleWeights, eps)
			vacuumOK := math.Abs(a.VacuumResponse) < eps
			traceOK := a.TraceResidual < eps
			maxOK := math.Abs(a.MaxResponse-a.ExpectedMaxResponse) < eps
			rankOK := a.Rank == a.Bridge.FockStateCount-1
			parityOK := a.ParityTraceResidual < eps

			checks := []theorem.Check{
				{
					Name:   "one-particle datum",
					Passed: len(a.OneParticleWeights) == 4,
					Detail: fmt.Sprintf("active contact-Higgs spectrum used as one-particle weights %s", FormatFloatSlice(a.OneParticleWeights)),
				},
				{
					Name:   "second-quantized Fock operator",
					Passed: a.Operator.Rows() == 16 && a.Operator.Cols() == 16,
					Detail: fmt.Sprintf("H_F is %dx%d diagonal on the 16 occupation states", a.Operator.Rows(), a.Operator.Cols()),
				},
				{
					Name:   "vacuum remains sterile at this level",
					Passed: vacuumOK,
					Detail: fmt.Sprintf("|Ω⟩ index=%d has response %.3e", a.VacuumIndex, a.VacuumResponse),
				},
				{
					Name:   "one-particle spectrum is preserved",
					Passed: oneParticleOK,
					Detail: fmt.Sprintf("one-particle responses %s match active finite spectrum", FormatFloatSlice(a.OneParticleResponses)),
				},
				{
					Name:   "pair-degenerate action seed",
					Passed: a.PairDegenerate,
					Detail: fmt.Sprintf("pair residual %.3e in the one-particle response spectrum", a.PairResidual),
				},
				{
					Name:   "finite Fock trace identity",
					Passed: traceOK,
					Detail: fmt.Sprintf("Tr(H_F)=%.10f, expected 2^(4-1)Σλ=%.10f, residual %.3e", a.Trace, a.ExpectedTrace, a.TraceResidual),
				},
				{
					Name:   "full-occupation response",
					Passed: maxOK,
					Detail: fmt.Sprintf("max response %.10f equals Σλ=%.10f", a.MaxResponse, a.ExpectedMaxResponse),
				},
				{
					Name:   "kernel size",
					Passed: rankOK,
					Detail: fmt.Sprintf("rank(H_F)=%d, kernel dimension=%d; only the Fock vacuum has zero response", a.Rank, a.Bridge.FockStateCount-a.Rank),
				},
				{
					Name:   "fermion-parity trace balance",
					Passed: parityOK,
					Detail: fmt.Sprintf("even trace %.10f, odd trace %.10f, residual %.3e", a.EvenParityTrace, a.OddParityTrace, a.ParityTraceResidual),
				},
				{
					Name:   "canonical embedding discipline",
					Passed: !a.CanonicalEigenvectorEmbeddingConstructed,
					Detail: "OPEN U-05: the spectral action is built, but no canonical K₇ eigenvector → Fock-mode map is claimed",
				},
				{
					Name:   "Yukawa discipline",
					Passed: !a.YukawaTextureConstructed,
					Detail: "OPEN U-07: this is a number-operator response, not a Standard Model Yukawa texture or mass matrix",
				},
			}
			notes := []string{
				"This gate creates the first actual action of the finite Higgs/contact data on the 16-state matter basis: H_F|n⟩=(Σλ_μ n_μ)|n⟩.",
				"The construction is standard second quantization of a four-dimensional one-particle spectrum; it does not use measured particle masses or fitted constants.",
				"The next hard problem is no longer dimension matching. It is constructing the canonical eigenvector/charge representation that turns this spectral action into a Yukawa texture.",
			}
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerMatter, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
		},
	}
}

func sameMultiset(a, b []float64, eps float64) bool {
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
