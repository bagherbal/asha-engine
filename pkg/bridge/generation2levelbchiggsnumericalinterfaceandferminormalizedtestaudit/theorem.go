package generation2levelbchiggsnumericalinterfaceandferminormalizedtestaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GATE-791-LEVEL-B-C-HIGGS-NUMERICAL-INTERFACE-FERMI-NORMALIZED-TEST"
	theoremName = "Gate 791 — Level-B C_Higgs Numerical Interface and Fermi-Normalized Test Audit"
)

func Generation2LevelBCHiggsNumericalInterfaceAndFermiNormalizedTestAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := Cached()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build analysis", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate790 Level-B interface", Passed: a.Gate790.Inherited && a.Gate790.FrozenLevelB && a.Gate790.CleanInterface && a.Gate790.DirectRuntimeFree, Detail: a.Gate790.Verdict},
			{Name: "define internal C_Higgs object", Passed: a.Internal.Defined && a.Internal.Dimensionless && !a.Internal.NativeHiggsTheorem && closeAbs(a.Internal.CHiggsValue, cHiggsSnapshot, 1e-15), Detail: a.Internal.CHiggsFormula},
			{Name: "record bridge quartic from C_Higgs", Passed: a.Internal.LambdaHBridgeFormula != "" && closeAbs(a.Internal.LambdaHBridgeValue, lambdaHBridgeSnapshot, 1e-15), Detail: a.Internal.LambdaHBridgeFormula},
			{Name: "define Fermi-normalized tree interface", Passed: a.FermiTree.Defined && a.FermiTree.AtTreeProxyLevel && !a.FermiTree.PoleMassTheorem && a.FermiTree.DimensionlessID == "4 sqrt(2) G_F m_H_tree_proxy^2 = C_Higgs", Detail: a.FermiTree.DimensionlessID},
			{Name: "recompute tree proxy numerical ledger", Passed: a.TreeLedger.Recomputed && closeAbs(a.TreeLedger.SqrtCHiggs, 1.0184402389953279, 1e-15) && closeAbs(a.TreeLedger.MTree, 125.38000000304908, 1e-11) && closeAbs(a.TreeLedger.FermiRatio, cHiggsSnapshot, 1e-15), Detail: FormatTreeLedger(a.TreeLedger)},
			{Name: "define observable-side diagnostic ratio", Passed: a.Diagnostic.Defined && a.Diagnostic.Level1COnly && !a.Diagnostic.NativePoleCorrection && !a.Diagnostic.ExternalPoleDerived && containsAll(a.Diagnostic.GapInterpretation, []string{"tree-to-pole", "scheme", "measurement"}), Detail: a.Diagnostic.DiagnosticGap},
			{Name: "define noncircular Level-B protocol", Passed: a.Protocol.Defined && a.Protocol.ObservedHiggsForbidden && containsAll(a.Protocol.AllowedInputs, []string{"N_eff", "kappa_orient", "boundary coordinates", "p", "L_Hopf"}) && containsAll(a.Protocol.Compute, []string{"C_Higgs", "lambda_H_bridge", "4 sqrt(2)"}), Detail: a.Protocol.ForbiddenCircularInputs[0]},
			{Name: "record correction-factor decomposition", Passed: a.Decomposition.Recorded && a.Decomposition.HistoryDominant && a.Decomposition.YukawaDilution && closeAbs(a.Decomposition.DeltaHiggs, a.Decomposition.DeltaHistory-a.Decomposition.EpsilonYukawa*(1+a.Decomposition.DeltaHistory), 1e-15), Detail: FormatDecomposition(a.Decomposition)},
			{Name: "record Level-B sensitivity formulas", Passed: a.Sensitivity.Recorded && containsAll(a.Sensitivity.KeyChannels, []string{"kappa_orient", "N_eff"}) && a.Sensitivity.CYukawaRelative != "" && a.Sensitivity.CHistoryDifferential != "" && a.Sensitivity.KappaERedDifferential != "", Detail: a.Sensitivity.CHistoryDifferential},
			{Name: "classify test status", Passed: a.Classification.Recorded && a.Classification.CleanObject && !a.Classification.LevelCNative && a.Classification.PoleMass == "m_H_pole not predicted" && a.Classification.PoleDiagnostic != "", Detail: a.Classification.CHiggs},
			{Name: "record source-pressure map", Passed: a.Pressure.Recorded && containsAll(a.Pressure.Pressures, []string{"GenerationMixingOperatorSeal", "Yukawa", "RadialHessianHopfTransportSeal", "BoundaryExteriorResponsePackageSeal", "Electroweak", "Tree-to-pole"}), Detail: FormatPressure(a.Pressure)},
			{Name: "record next branch recommendation", Passed: a.Next.Recorded && a.Next.Recommended == "Gate 792 — Level-B Error Budget and Independent-Input Sensitivity Audit" && a.Next.Alternative != "" && a.Next.Reason != "", Detail: a.Next.Recommended},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Enforced && !a.Firewalls.CHiggsNativeHiggsTheorem && !a.Firewalls.CHiggsPoleMassPrediction && !a.Firewalls.FermiTreeIdentityPoleMass && !a.Firewalls.LambdaHIndependentRuntime && !a.Firewalls.TreeProxyPoleMass && !a.Firewalls.ExternalPoleASHADerived && !a.Firewalls.DeltaRPoleCorrectionTheorem && !a.Firewalls.NEffNativeYukawa && !a.Firewalls.KappaOrientNativePMNSCKM && !a.Firewalls.LHopfNativeHistoryLoop && !a.Firewalls.FWallNativeBoundary && !a.Firewalls.VOrGFNativeScale && a.Firewalls.Verdict == StatusFirewallPreservedGate791, Detail: a.Firewalls.Verdict},
		}
		ok := true
		for _, c := range checks {
			if !c.Passed {
				ok = false
				break
			}
		}
		status := theorem.BridgeRequired
		if !ok {
			status = theorem.FailedRoute
		}
		notes := append([]string{a.Truth, FormatTreeLedger(a.TreeLedger), FormatDecomposition(a.Decomposition), FormatPressure(a.Pressure), a.FinalStatement}, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
