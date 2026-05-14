package ashafinalarchitectureledger

import (
	"math"
	"strconv"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func AshaFrameworkFinalArchitectureLedgerEpistemologicalSealTheorem() theorem.Theorem {
	const id = "BRIDGE-ASHA-FRAMEWORK-FINAL-ARCHITECTURE-LEDGER-EPISTEMOLOGICAL-SEAL"
	const name = "ASHA Framework Final Architecture Ledger & Epistemological Seal"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 387 final architecture ledger", Passed: false, Detail: err.Error()}}}
		}
		status := StatusLine(a)
		checks := []theorem.Check{
			{Name: "absolute geometric predictions are catalogued", Passed: a.Absolute.Executed && a.Absolute.GaugeGroup == "SU(3) × SU(2) × U(1)" && a.Absolute.Generations == 3 && a.Absolute.HiggsDoublets == 1 && a.Absolute.GaugeBosons == 12 && a.Absolute.ParameterFreeDerivations >= 10, Detail: a.Absolute.Verdict},
			{Name: "almost-commutative product bridge is sealed structurally", Passed: a.Product.Executed && a.Product.CCMInstalled && strings.Contains(a.Product.DiracOperator, "D_M") && strings.Contains(a.Product.DiracOperator, "D_F") && strings.Contains(a.Product.Geometry, "M × F"), Detail: a.Product.Action},
			{Name: "Gate 385 Higgs tree proxy is inherited without pole overclaim", Passed: a.Higgs.Executed && a.Higgs.SourceGate == 385 && a.Higgs.TreeProxySealed && a.Higgs.EdgeMeasureSelected && !a.Higgs.PoleMassDerived && a.Higgs.LambdaEW > 0.12 && a.Higgs.LambdaEW < 0.13 && a.Higgs.MassPfaffianGeV > 124 && a.Higgs.MassPfaffianGeV < 126, Detail: a.Higgs.Verdict},
			{Name: "Gate 386 cosmology structures are catalogued but not promoted to predictions", Passed: a.Cosmology.Executed && len(a.Cosmology.Structures) == 3 && a.Cosmology.ConditionalTargetsOpened == 2 && a.Cosmology.HardPredictionsDerived == 0, Detail: a.Cosmology.Verdict},
			{Name: "Gate 372 thirteen-moduli environmental quarantine is preserved", Passed: a.Moduli.Executed && a.Moduli.SourceGate == 372 && a.Moduli.MinimalChargedFiniteDiracDim == 13 && a.Moduli.ExternalMinimalLedger == 15 && !a.Moduli.NativeReductionBelow13 && a.Moduli.HiddenFlavorConstraints == 0, Detail: a.Moduli.Decomposition},
			{Name: "epistemic boundary forbids overclaiming numerical ToE closure", Passed: a.Boundary.Executed && a.Boundary.NativePredictionsAreNotInputs && a.Boundary.EnvironmentalInputsAreQuarantined && a.Boundary.NoFlavorFitting && a.Boundary.NoCosmologyFitting && a.Boundary.NoPoleMassOverclaim, Detail: a.Boundary.NotCompleteAs},
			{Name: "final ledger is sealed with explicit failure routes", Passed: a.Final.Executed && a.Final.ProjectSealed && a.Final.EnvironmentalModuliCount == 13 && a.Final.CosmologicalHardCount == 0 && strings.Contains(status, StatusFinalLedgerCompiled) && strings.Contains(status, StatusFailedDarkMatterNotPredicted) && strings.Contains(status, StatusFailedUniverseLifetimeNotDerived), Detail: a.Final.FinalStatement},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth, constantsSummary(a)}}
	}}
}

func constantsSummary(a Analysis) string {
	m := FinalArchitectureConstants()
	return strings.Join([]string{
		"sin²θ_W=" + format(m["sin2_theta_w_boundary"]),
		"α_GUT^-1=" + format(m["alpha_gut_inverse_branch"]),
		"e/a²(node)=" + format(m["trace_ratio_e_over_a2_node"]),
		"e/a²(edge)=" + format(m["trace_ratio_e_over_a2_edge"]),
		"λ_edge=" + format(m["lambda_higgs_edge"]),
		"m_H(tree proxy)=" + format(m["higgs_tree_proxy_GeV"]),
	}, "; ")
}

func format(v float64) string {
	if math.IsNaN(v) {
		return "NaN"
	}
	return strings.TrimRight(strings.TrimRight(fmtFloat(v), "0"), ".")
}

func fmtFloat(v float64) string { return strconv.FormatFloat(v, 'f', 12, 64) }
