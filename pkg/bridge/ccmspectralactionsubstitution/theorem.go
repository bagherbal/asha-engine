package ccmspectralactionsubstitution

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func CCMSpectralActionDirectSubstitutionCompleteCoefficientLedgerTheorem() theorem.Theorem {
	const id = "BRIDGE-CCM-SPECTRAL-ACTION-DIRECT-SUBSTITUTION"
	const name = "CCM Spectral Action Direct Substitution / Complete Coefficient Ledger"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build CCM substitution ledger", Passed: false, Detail: err.Error()}}}
		}
		c := a.Calculation
		checks := []theorem.Check{
			{Name: "CCM almost-commutative formula is installed directly", Passed: c.Executed && strings.Contains(c.Formula.EinsteinHilbert, "96 f₂Λ²") && strings.Contains(c.Formula.HiggsQuartic, "f₀/2"), Detail: c.Formula.Source},
			{Name: "product geometry context matches the CCM theorem", Passed: c.StructuralClosure && strings.Contains(c.Lagrangian, "D_M⊗1") && strings.Contains(c.Lagrangian, "γ₅⊗D_F"), Detail: c.Lagrangian},
			{Name: "Einstein-Hilbert coefficient supersedes Gate 378 generic arithmetic", Passed: math.Abs(c.Einstein.CoefficientWithPreviousF2NoC-(1.0/(16.0*math.Pi))) < 1e-15 && math.Abs(c.Einstein.GapFactor-8.0*math.Pi) < 1e-12, Detail: c.Einstein.Verdict},
			{Name: "canonical leading cutoff moment is π²/8, not π/64", Passed: math.Abs(c.Einstein.RequiredF2LambdaMP2Leading-math.Pi*math.Pi/8.0) < 1e-15 && c.Einstein.RequiredF2LambdaMP2Leading > PreviousF2LambdaMP2, Detail: c.Einstein.RequiredFormula},
			{Name: "Higgs finite trace ratio is not automatically the canonical quartic", Passed: math.Abs(c.Higgs.EOverA2-LambdaTraceRatio) < 1e-15 && c.Higgs.QuarticNoOuterPiConvention < c.Higgs.PreviousAshaRatio && c.Higgs.QuarticCanonicalOuterPiConvention != c.Higgs.PreviousAshaRatio, Detail: c.Higgs.Verdict},
			{Name: "gauge absolute normalization remains a representation-trace channel", Passed: c.Gauge.RequiresRepresentationTrace && !c.Gauge.AbsoluteClosed, Detail: c.Gauge.Verdict},
			{Name: "cosmological channel is symbolic, not a dark-energy prediction", Passed: c.Cosmological.NeedsF4 && c.Cosmological.NeedsVacuumSubtraction, Detail: c.Cosmological.Verdict},
			{Name: "full numerical ToE closure is not claimed", Passed: !c.FullNumericalTOEClosure && strings.Contains(StatusLine(c), StatusFailedFullNumericalTOENotClosed), Detail: c.Truth},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Calculation.Truth}}
	}}
}
