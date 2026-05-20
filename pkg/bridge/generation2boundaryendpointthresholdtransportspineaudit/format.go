package generation2boundaryendpointthresholdtransportspineaudit

import (
	"fmt"
	"math"
	"strings"
)

func f64(x float64) string {
	if math.IsNaN(x) {
		return "symbolic"
	}
	return fmt.Sprintf("%.15g", x)
}

func FormatInherited(a InheritedGate605) string {
	return fmt.Sprintf("masterVector=%t topRecommendation=%q rgTop=%t verdict=%q", a.MasterVectorBuilt, a.TopRecommendation, a.RGTopActionable, a.Verdict)
}

func FormatNativeBoundaryRow(r NativeBoundaryConditionRow) string {
	return fmt.Sprintf("symbol=%q class=%q source=%q normalizations=%v blocker=%q meaning=%q verdict=%q", r.Symbol, r.Classification, r.GateSource, r.RequiredNormalizations, r.PhysicalPromotionBlock, r.Meaning, r.Verdict)
}

func FormatNativeBoundaryTable(rows []NativeBoundaryConditionRow) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatNativeBoundaryRow(r))
	}
	return strings.Join(parts, " | ")
}

func FormatEndpointRow(r EndpointObservedLedgerRow) string {
	return fmt.Sprintf("symbol=%q value=%s unit=%q scale=%q scheme=%q source=%q role=%q class=%q verdict=%q", r.Symbol, f64(r.Value), r.Unit, r.Scale, r.Scheme, r.Source, r.Role, r.Classification, r.Verdict)
}

func FormatEndpointLedger(rows []EndpointObservedLedgerRow) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatEndpointRow(r))
	}
	return strings.Join(parts, " | ")
}

func FormatGaugeRow(r GaugeTransportRow) string {
	return fmt.Sprintf("quantity=%q boundaryOrFlow=%q value=%s formula=%q approx=%q interpretation=%q verdict=%q", r.Quantity, r.BoundaryOrFlow, f64(r.RuntimeValue), r.Formula, r.Approximation, r.Interpretation, r.Verdict)
}

func FormatGaugeTransport(rows []GaugeTransportRow) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatGaugeRow(r))
	}
	return strings.Join(parts, " | ")
}

func FormatScalarRow(r ScalarTransportRow) string {
	return fmt.Sprintf("quantity=%q value=%s unit=%q formula=%q approx=%q interpretation=%q verdict=%q", r.Quantity, f64(r.RuntimeValue), r.Unit, r.Formula, r.Approximation, r.Interpretation, r.Verdict)
}

func FormatScalarTransport(rows []ScalarTransportRow) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatScalarRow(r))
	}
	return strings.Join(parts, " | ")
}

func FormatThresholdRow(r ThresholdCorrectionSlotRow) string {
	return fmt.Sprintf("slot=%q sector=%q purpose=%q status=%q residual=%q verdict=%q", r.Slot, r.Sector, r.Purpose, r.CurrentStatus, r.RuntimeResidual, r.Verdict)
}

func FormatThresholdSlots(rows []ThresholdCorrectionSlotRow) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatThresholdRow(r))
	}
	return strings.Join(parts, " | ")
}

func FormatKineticBlockerRow(r KineticNormalizationBlockerRow) string {
	return fmt.Sprintf("blocker=%q sector=%q why=%q status=%q verdict=%q", r.Blocker, r.Sector, r.WhyRequired, r.CurrentStatus, r.Verdict)
}

func FormatKineticBlockers(rows []KineticNormalizationBlockerRow) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatKineticBlockerRow(r))
	}
	return strings.Join(parts, " | ")
}

func FormatFlavorRelationRow(r FlavorSealRGRelationRow) string {
	return fmt.Sprintf("item=%q role=%q native=%q firewall=%q verdict=%q", r.Item, r.RoleInTransport, r.NativeStatus, r.Firewall, r.Verdict)
}

func FormatFlavorRelation(rows []FlavorSealRGRelationRow) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatFlavorRelationRow(r))
	}
	return strings.Join(parts, " | ")
}

func FormatProductTimeFirewall(p ProductTimeFirewall) string {
	return fmt.Sprintf("rgProductTime=%t rgOSHilbert=%t rgCosmoTime=%t statement=%q verdict=%q", p.RGScaleIsProductTime, p.RGScaleIsOSHilbert, p.RGScaleIsCosmoTime, p.Statement, p.Verdict)
}

func FormatFormula(f UpdatedFormula) string {
	return fmt.Sprintf("formula=%q native=%v transport=%v thresholds=%v endpoints=%v blocked=%v verdict=%q", f.Formula, f.NativeBoundary, f.TransportSpine, f.ThresholdSlots, f.EndpointLedgers, f.BlockedPromotions, f.Verdict)
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("fullUnification=%t endpoint=%t kinetic=%t vev=%t flavor=%t productTime=%t thresholdsExplicit=%t verdict=%q", f.ClaimsFullUnification, f.DerivesEndpoint, f.DerivesKineticScale, f.DerivesVEV, f.DerivesFlavor, f.DerivesProductTime, f.ThresholdsExplicit, f.Verdict)
}
