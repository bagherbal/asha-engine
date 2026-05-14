package cosmologicalobservablesdarksector

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func CosmologicalObservablesDarkSectorPredictionAfterHiggsSealSieveTheorem() theorem.Theorem {
	const id = "BRIDGE-COSMOLOGICAL-OBSERVABLES-DARK-SECTOR-PREDICTION-AFTER-HIGGS-SEAL-SIEVE"
	const name = "Cosmological Observables & Dark Sector Prediction Sieve"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 386 audit", Passed: false, Detail: err.Error()}}}
		}
		status := StatusLine(a)
		checks := []theorem.Check{
			{Name: "Gate 385 sealed Higgs proxy is inherited", Passed: a.Higgs.Executed && a.Higgs.SourceGate == 385 && a.Higgs.TreeProxySealed && a.Higgs.UsesEdgeMeasure && !a.Higgs.PoleMassDerived && a.Higgs.LambdaEW > 0.12 && a.Higgs.LambdaEW < 0.13 && a.Higgs.HiggsMassProxyGeV > 124 && a.Higgs.HiggsMassProxyGeV < 126, Detail: a.Higgs.DirectAnswer},
			{Name: "B-gap / heavy-sector ledger is inherited without overpromoting dark stability", Passed: a.Heavy.Executed && a.Heavy.GeometricallyMandated && a.Heavy.BGapScaleGeV > 1e6 && a.Heavy.ThresholdJumpDeltaLambda < 0 && !a.Heavy.StableRelicTheorem && !a.Heavy.DecayWidthDerived && !a.Heavy.AnnihilationRateDerived, Detail: a.Heavy.DirectAnswer},
			{Name: "Boltzmann relic-density problem is formalized but not closed", Passed: a.DarkMatter.Executed && a.DarkMatter.NativeMassScaleAvailable && len(a.DarkMatter.Inputs) >= 6 && !a.DarkMatter.StableCandidateDerived && !a.DarkMatter.BoltzmannKernelClosed && !a.DarkMatter.OmegaH2Derived && math.IsNaN(a.DarkMatter.OmegaH2), Detail: a.DarkMatter.DirectAnswer},
			{Name: "vacuum RG/bounce problem is formalized but not solved", Passed: a.Vacuum.Executed && a.Vacuum.LambdaEW > 0 && a.Vacuum.ThresholdDeltaLambda < 0 && strings.Contains(a.Vacuum.OneLoopBetaLambda, "y_t") && !a.Vacuum.FullRGTrajectoryDerived && !a.Vacuum.ThresholdMatchingDerived && !a.Vacuum.LambdaMinimumDerived && !a.Vacuum.LifetimeDerived && math.IsNaN(a.Vacuum.EuclideanBounceAction) && math.IsNaN(a.Vacuum.LifetimeYears), Detail: a.Vacuum.DirectAnswer},
			{Name: "observable census records zero native hard cosmological predictions", Passed: a.Census.Executed && a.Census.RequestedObservables == 2 && a.Census.HardPredictionsDerived == 0 && a.Census.ConditionalTargetsOpened == 2 && !a.Census.DarkMatterDerived && !a.Census.VacuumFateDerived && !a.Census.FullNumericalTOEClosed, Detail: a.Census.FinalStatement},
			{Name: "firewalls prevent fitting cosmology/top/RG data", Passed: a.Firewall.Executed && a.Firewall.NoObservedOmegaFitted && a.Firewall.NoReheatingTemperatureInserted && a.Firewall.NoCrossSectionInserted && a.Firewall.NoDecayWidthInserted && a.Firewall.NoTopYukawaInserted && a.Firewall.NoGaugeTrajectoryInserted && a.Firewall.NoThresholdSignAssumed && a.Firewall.NoBounceMinimumInserted && a.Firewall.NoLifetimeTargetFitted && a.Firewall.NoClaimBeyondNativeInputs, Detail: a.Firewall.Verdict},
			{Name: "correct status ledger is emitted", Passed: strings.Contains(status, StatusFailedDarkMatterRelicNotDerived) && strings.Contains(status, StatusFailedVacuumStabilityNotDerived) && strings.Contains(status, StatusFailedFullNumericalTOENotClosed), Detail: status},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth}}
	}}
}
