package generation2levelbchiggsnumericalinterfaceandferminormalizedtestaudit

import (
	"strings"
	"testing"
)

func TestGate791InternalObjectAndFermiInterface(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate790.Inherited || !a.Gate790.FrozenLevelB || !a.Gate790.CleanInterface || !a.Gate790.DirectRuntimeFree {
		t.Fatalf("bad Gate790 inheritance: %+v", a.Gate790)
	}
	if !a.Internal.Defined || !a.Internal.Dimensionless || a.Internal.NativeHiggsTheorem || !strings.Contains(a.Internal.CHiggsFormula, "3/N_eff") || !closeAbs(a.Internal.CHiggsValue, cHiggsSnapshot, 1e-15) || !closeAbs(a.Internal.LambdaHBridgeValue, lambdaHBridgeSnapshot, 1e-15) {
		t.Fatalf("bad internal C_Higgs object: %+v", a.Internal)
	}
	if !a.FermiTree.Defined || !a.FermiTree.AtTreeProxyLevel || a.FermiTree.PoleMassTheorem || !strings.Contains(a.FermiTree.DimensionlessID, "4 sqrt(2)") || !strings.Contains(a.FermiTree.TreeProxyFormula, "sqrt(C_Higgs)") {
		t.Fatalf("bad Fermi tree interface: %+v", a.FermiTree)
	}
}

func TestGate791TreeProxyLedgerAndDecomposition(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.TreeLedger.Recomputed || !closeAbs(a.TreeLedger.GF, gFEquivalentSnapshot, 1e-18) || !closeAbs(a.TreeLedger.SqrtCHiggs, 1.0184402389953279, 1e-15) || !closeAbs(a.TreeLedger.VHalf, 123.1098254, 1e-12) || !closeAbs(a.TreeLedger.MTree, 125.38000000304908, 1e-11) || !closeAbs(a.TreeLedger.MTreeSquared, 15720.144400764586, 1e-9) || !closeAbs(a.TreeLedger.FermiRatio, cHiggsSnapshot, 1e-15) || !closeAbs(a.TreeLedger.ReducedQuarter, cHiggsSnapshot/4, 1e-15) {
		t.Fatalf("bad tree proxy ledger: %s", FormatTreeLedger(a.TreeLedger))
	}
	if !a.Decomposition.Recorded || !a.Decomposition.HistoryDominant || !a.Decomposition.YukawaDilution || !closeAbs(a.Decomposition.EpsilonYukawa, 0.0007751811187991509, 1e-18) || !closeAbs(a.Decomposition.DeltaHistory, 0.03802517792362492, 1e-16) || !closeAbs(a.Decomposition.DeltaHiggs, 0.03722052040486035, 1e-16) {
		t.Fatalf("bad decomposition: %s", FormatDecomposition(a.Decomposition))
	}
}

func TestGate791DiagnosticProtocolSensitivityAndStatus(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Diagnostic.Defined || !a.Diagnostic.Level1COnly || a.Diagnostic.NativePoleCorrection || a.Diagnostic.ExternalPoleDerived || !strings.Contains(a.Diagnostic.ExternalRatio, "m_H_pole_external") || !strings.Contains(a.Diagnostic.DiagnosticGap, "R_pole_external") {
		t.Fatalf("bad diagnostic: %+v", a.Diagnostic)
	}
	if !a.Protocol.Defined || !a.Protocol.ObservedHiggsForbidden || !containsAll(a.Protocol.AllowedInputs, []string{"N_eff", "kappa_orient", "boundary coordinates", "p", "L_Hopf"}) || !containsAll(a.Protocol.Compute, []string{"C_Higgs", "lambda_H_bridge", "4 sqrt(2)"}) {
		t.Fatalf("bad protocol: %+v", a.Protocol)
	}
	if !a.Sensitivity.Recorded || !containsAll(a.Sensitivity.KeyChannels, []string{"kappa_orient", "N_eff"}) || a.Sensitivity.CYukawaRelative == "" || a.Sensitivity.CHistoryDifferential == "" || a.Sensitivity.KappaLambdaRedDifferential == "" || a.Sensitivity.KappaERedDifferential == "" {
		t.Fatalf("bad sensitivity: %+v", a.Sensitivity)
	}
	if !a.Classification.Recorded || !a.Classification.CleanObject || a.Classification.LevelCNative || a.Classification.PoleMass != "m_H_pole not predicted" || !strings.Contains(a.Classification.CHiggs, "Level-B") || !strings.Contains(a.Classification.PoleDiagnostic, "Level-1C") {
		t.Fatalf("bad classification: %+v", a.Classification)
	}
}

func TestGate791PressureBranchFirewallsAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Pressure.Recorded || !containsAll(a.Pressure.Pressures, []string{"GenerationMixingOperatorSeal", "Yukawa", "RadialHessianHopfTransportSeal", "BoundaryExteriorResponsePackageSeal", "Electroweak", "Tree-to-pole"}) {
		t.Fatalf("bad pressure map: %+v", a.Pressure)
	}
	if !a.Next.Recorded || a.Next.Recommended != "Gate 792 — Level-B Error Budget and Independent-Input Sensitivity Audit" || a.Next.Alternative == "" || a.Next.Reason == "" {
		t.Fatalf("bad next recommendation: %+v", a.Next)
	}
	if !a.Firewalls.Enforced || a.Firewalls.CHiggsNativeHiggsTheorem || a.Firewalls.CHiggsPoleMassPrediction || a.Firewalls.FermiTreeIdentityPoleMass || a.Firewalls.LambdaHIndependentRuntime || a.Firewalls.TreeProxyPoleMass || a.Firewalls.ExternalPoleASHADerived || a.Firewalls.DeltaRPoleCorrectionTheorem || a.Firewalls.NEffNativeYukawa || a.Firewalls.KappaOrientNativePMNSCKM || a.Firewalls.LHopfNativeHistoryLoop || a.Firewalls.FWallNativeBoundary || a.Firewalls.VOrGFNativeScale {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
	if !strings.Contains(a.FinalStatement, "does not turn C_Higgs into a native prediction") || !strings.Contains(a.FinalStatement, "Fermi-normalized tree identity") || !strings.Contains(a.FinalStatement, "Gate 792") {
		t.Fatalf("bad final statement: %s", a.FinalStatement)
	}
	res := Generation2LevelBCHiggsNumericalInterfaceAndFermiNormalizedTestAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status note %s", want)
		}
	}
}
