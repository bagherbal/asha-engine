package generation2gaugemeetingscaletrianglegeometryaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate608LogTriangleGeometry(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.LogGeometry.Ratio13Over12-10.2656454271369) > 1e-9 {
		t.Fatalf("unexpected Lambda13/Lambda12 ratio: %.15g", a.LogGeometry.Ratio13Over12)
	}
	if math.Abs(a.LogGeometry.Ratio23Over13-82.648801838935) > 1e-7 {
		t.Fatalf("unexpected Lambda23/Lambda13 ratio: %.15g", a.LogGeometry.Ratio23Over13)
	}
	if !(a.LogGeometry.SpreadDecades > 2.92 && a.LogGeometry.SpreadDecades < 2.94) {
		t.Fatalf("unexpected spread: %.15g", a.LogGeometry.SpreadDecades)
	}
}

func TestGate608BoundaryChoicesAndBetaDiagnostics(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.BoundaryChoices) != 4 || !containsBoundaryChoice(a.BoundaryChoices, "Lambda_geom") {
		t.Fatalf("bad boundary choices: %s", FormatBoundaryChoices(a.BoundaryChoices))
	}
	if len(a.BetaDeformations) != 3 {
		t.Fatalf("bad beta deformation rows: %s", FormatBetaDeformations(a.BetaDeformations))
	}
	if math.Abs(a.BetaDeformations[0].DeltaB3+0.933360651351616) > 1e-9 {
		t.Fatalf("unexpected strong beta deformation: %.15g", a.BetaDeformations[0].DeltaB3)
	}
	if a.BetaDeformations[2].Norm <= 0 || a.BetaDeformations[2].LambdaUGeV != a.LogGeometry.GeometricMeanGeV {
		t.Fatalf("bad minimal norm row: %+v", a.BetaDeformations[2])
	}
}

func TestGate608TheoremAndFirewalls(t *testing.T) {
	res := Generation2GaugeMeetingScaleTriangleGeometryAuditTheorem().Verify()
	if !res.Passed() {
		t.Fatalf("theorem checks failed: %+v", res)
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusGate607Inherited, StatusLogTriangleComputed, StatusBoundaryChoiceClassified, StatusStructuredTransportLedger, StatusNoSingleOneLoopUnification, StatusNoNativeThresholdTheorem, StatusNoNativeLambdaUSelection, StatusGate608Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
