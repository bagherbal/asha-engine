package neutrinotextureaudit

import "testing"

func TestGate232BuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Seal.Active || a.Seal.FiniteDerived {
		t.Fatalf("expected phenomenological NeutrinoTextureSeal activation, got %+v", a.Seal)
	}
	if a.Audit.AnySMMassProxySupported {
		t.Fatalf("direct SM mass proxies should fail, got %+v", a.Audit.BestStandardSMMassProxy)
	}
	if !a.Audit.AnyGenerationProxySupported {
		t.Fatalf("expected generation-index support, got %+v", a.Audit)
	}
	if a.Audit.BestGenerationIndexProxy.Name != "generation-index quadratic" {
		t.Fatalf("expected quadratic generation-index best proxy, got %q", a.Audit.BestGenerationIndexProxy.Name)
	}
	if a.Audit.BestGenerationIndexProxy.RatioM2ToM3 < 0.18 || a.Audit.BestGenerationIndexProxy.RatioM2ToM3 > 0.22 {
		t.Fatalf("expected quadratic ratio near 0.20, got %.12g", a.Audit.BestGenerationIndexProxy.RatioM2ToM3)
	}
	if a.Audit.RequiredM2DiracGeV < 2 || a.Audit.RequiredM2DiracGeV > 3 {
		t.Fatalf("expected required mD2 near 2.4 GeV, got %.12g", a.Audit.RequiredM2DiracGeV)
	}
	if a.Matrix.PMNSMatrixDerived || a.Matrix.DiracTextureDerived {
		t.Fatalf("finite texture/PMNS must not be derived: %+v", a.Matrix)
	}
	if a.Firewall.FiniteCorePolluted || a.Firewall.TunesToObservedMixingAngles {
		t.Fatalf("firewall violated: %+v", a.Firewall)
	}
}
