package generation2boundaryendpointthresholdtransportspineaudit

import (
	"strings"
	"testing"
)

func TestGate606BoundaryEndpointTables(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.RGTopActionable {
		t.Fatalf("Gate605 inheritance did not select RG transport: %+v", a.Inherited)
	}
	if !containsBoundary(a.NativeBoundaryTable, "k_Y=5/3") || !containsBoundary(a.NativeBoundaryTable, "m_W²/m_Z²=5/8") {
		t.Fatalf("missing native boundary rows: %s", FormatNativeBoundaryTable(a.NativeBoundaryTable))
	}
	if !containsEndpoint(a.EndpointLedger, "g1(M_Z)") || !containsEndpoint(a.EndpointLedger, "m_W") || !containsEndpoint(a.EndpointLedger, "v") {
		t.Fatalf("missing endpoint ledger rows: %s", FormatEndpointLedger(a.EndpointLedger))
	}
}

func TestGate606TransportSlotsAndFirewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !containsGauge(a.GaugeTransport, "Delta_3") || !containsGauge(a.GaugeTransport, "R_3") || !containsGauge(a.GaugeTransport, "Delta_sin²") {
		t.Fatalf("missing gauge transport rows: %s", FormatGaugeTransport(a.GaugeTransport))
	}
	if !containsScalar(a.ScalarTransport, "lambda(Lambda_12)") || !containsThreshold(a.ThresholdSlots, "delta_lambda") || !containsBlocker(a.KineticBlockers, "K_phi") {
		t.Fatalf("missing scalar/threshold/blocker rows")
	}
	if a.ProductTimeFirewall.RGScaleIsProductTime || a.Firewalls.DerivesEndpoint || a.Firewalls.ClaimsFullUnification {
		t.Fatalf("firewall breach: %+v %+v", a.ProductTimeFirewall, a.Firewalls)
	}
}

func TestGate606TheoremStatuses(t *testing.T) {
	res := Generation2BoundaryEndpointThresholdTransportSpineAuditTheorem().Verify()
	if !res.Passed() {
		t.Fatalf("theorem checks failed: %+v", res)
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusGate605Inherited, StatusNativeBoundaryClassified, StatusGaugeRGSlotsDefined, StatusScalarRGSlotsDefined, StatusThresholdLedgerDefined, StatusRGThresholdNextSpine, StatusNoNativeRGThresholdTheorem, StatusNoAbsoluteKineticScale, StatusRGScaleNotProductTime, StatusGate606Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
