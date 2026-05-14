package familycoefficientselector

import (
	"strings"
	"testing"
)

func TestGate414ArenaAndFunctionals(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Arena.NoncommutingCapacity || a.Arena.GeneratedAlgebraDimension != 9 || a.Arena.CoefficientsNative {
		t.Fatalf("bad arena: %s", FormatArena(a.Arena))
	}
	if len(a.Functionals) < 4 {
		t.Fatalf("missing functionals")
	}
	var traceOK, curvatureOK, sourceQuarantined bool
	for _, f := range a.Functionals {
		if f.Name == "quadratic trace/norm" && !f.UniqueCoefficientRay && !f.SelectorNative {
			traceOK = true
		}
		if f.Name == "adjoint-curvature relative to K" && f.SelectsNoncommutingTexture && !f.SelectsPhysicalSectorWeights && !f.SelectorNative {
			curvatureOK = true
		}
		if f.Name == "sector-split source functional" && f.UniqueCoefficientRay && !f.EmpiricalIndependent && !f.SelectorNative {
			sourceQuarantined = true
		}
	}
	if !traceOK || !curvatureOK || !sourceQuarantined {
		t.Fatalf("bad functionals:\n%s", RenderFunctionalSummary(a.Functionals))
	}
}

func TestGate414ConnectionCoefficientsFirewall(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Connection.YangMillsMinimizerFlat || !a.Connection.FlatMinimizerCommutes || !a.Connection.NonzeroCurvatureRequiresSource || a.Connection.ConnectionNativeInCurrentAsha || a.Connection.CoefficientsFixedByCurvature || a.Connection.CKMAnglePredicted {
		t.Fatalf("bad connection: %s", FormatConnection(a.Connection))
	}
	if a.Coefficients.TopologicalCoefficientValuesFound != 0 || a.Coefficients.RootsOfUnityFixCoefficients || a.Coefficients.TraceFixesCoefficients || a.Coefficients.CurvatureFixesCoefficients || a.Coefficients.SectorSplittingNative || a.Coefficients.YukawaDataImported {
		t.Fatalf("bad coefficients: %s", FormatCoefficients(a.Coefficients))
	}
	if !a.Firewall.NoObservedMassesImported || !a.Firewall.NoCKMImported || !a.Firewall.NoPMNSImported || !a.Firewall.NoYukawaMatricesInserted || !a.Firewall.AxiomStatusPreserved || !a.Firewall.NoNativeDerivationClaimed {
		t.Fatalf("bad firewall: %s", FormatFirewall(a.Firewall))
	}
}

func TestGate414ModuliStatusesAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Moduli.BestNativeDim != Gate372ChargedFlavorModuliDim || a.Moduli.NativeReductionBelow13 || !a.Moduli.ConditionalMixingCapacity || !a.Moduli.CoefficientsRemainFree || !a.Moduli.FirewallPreserved {
		t.Fatalf("bad moduli: %s", FormatModuli(a.Moduli))
	}
	joined := strings.Join(Statuses(a), "\n")
	for _, needle := range []string{StatusSelectorArenaFormalized, StatusTraceFunctionalAudited, StatusCurvatureFunctionalAudited, StatusFailedNoNativeSelector, StatusFailedSectorWeightsFree, StatusFailedNoCKMAnglePrediction, StatusFirewallPreserved13Moduli} {
		if !strings.Contains(joined, needle) {
			t.Fatalf("missing %q in\n%s", needle, joined)
		}
	}
	res := FamilyCoefficientSelectorConstrainedConnectionCurvatureSieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem failed: %+v", res)
	}
}

func TestGate414Markdown(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderMarkdown(a)
	for _, needle := range []string{"Gate 414 Registry Audit", "Functional table", "Constrained family connection", StatusFailedNoNativeSelector, StatusFailedSectorWeightsFree, StatusFirewallPreserved13Moduli, "gate=415"} {
		if !strings.Contains(md, needle) {
			t.Fatalf("markdown missing %q\n%s", needle, md)
		}
	}
}
