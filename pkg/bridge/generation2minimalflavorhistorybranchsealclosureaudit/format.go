package generation2minimalflavorhistorybranchsealclosureaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(a InheritedGate603) string {
	return fmt.Sprintf("electron=%t p3=%t plusJ=%t sigmaGauge=%t optDisc=%t best=%.15g next=%.15g verdict=%q", a.SelectsElectronWall, a.SelectsP3Nu, a.SelectsPositiveJ, a.SigmaGaugeForBFlav, a.OptionalSignedDiscriminantSeal, a.BestResidual, a.NextDistinctResidual, a.Verdict)
}
func FormatBranchStackRow(r BranchStackRow) string {
	return fmt.Sprintf("layer=%q item=%q class=%q needed=%t native=%t role=%q verdict=%q", r.Layer, r.Item, r.Classification, r.NeededForBFlav, r.Native, r.Role, r.Verdict)
}
func FormatBranchStack(rows []BranchStackRow) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatBranchStackRow(r))
	}
	return strings.Join(parts, " | ")
}
func FormatClassificationRow(r ClassificationRow) string {
	return fmt.Sprintf("item=%q native=%t ext=%t seal=%t gauge=%t ledger=%t needed=%t explanation=%q verdict=%q", r.Item, r.Native, r.AlgebraicExtension, r.EnvironmentalBranchSeal, r.GaugeConvention, r.ObservedLedger, r.NeededForBFlav, r.Explanation, r.Verdict)
}
func FormatClassification(rows []ClassificationRow) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatClassificationRow(r))
	}
	return strings.Join(parts, " | ")
}
func FormatMinimalityRow(r MinimalityRow) string {
	return fmt.Sprintf("item=%q requiredB=%t requiredFull=%t reason=%q verdict=%q", r.Item, r.RequiredForBFlav, r.RequiredForFullOrderedHistory, r.Reason, r.Verdict)
}
func FormatMinimality(rows []MinimalityRow) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatMinimalityRow(r))
	}
	return strings.Join(parts, " | ")
}
func FormatMinimalSeal(a MinimalFlavorHistoryBranchSeal) string {
	return fmt.Sprintf("name=%q components=%v selected=%v notIncluded=%v native=%t environmental=%t verdict=%q", a.Name, a.Components, a.SelectedByBFlav, a.NotIncluded, a.IsNative, a.IsEnvironmental, a.Verdict)
}
func FormatOptionalFullOrder(a OptionalFullOrderSeal) string {
	return fmt.Sprintf("name=%q reqB=%t reqFull=%t data=%v native=%t statement=%q verdict=%q", a.Name, a.RequiredForBFlav, a.RequiredForFullOrder, a.Data, a.NativeTheoremPresent, a.Statement, a.Verdict)
}
func FormatFormula(a UpdatedHistoryTransportFlavorFormula) string {
	return fmt.Sprintf("formula=%q Y=%v Omega=%v T=%v remaining=%v verdict=%q", a.Formula, a.YCore, a.OmegaCore, a.TCore, a.RemainingRawInputs, a.Verdict)
}
func FormatFirewalls(a Firewalls) string {
	return fmt.Sprintf("koide=%t masses=%t pmns=%t ckm=%t bflav=%t carrier=%t selector=%t g352=%t g596=%t g599=%t g603=%t verdict=%q", a.DerivesKoide, a.DerivesChargedLeptonMasses, a.DerivesPMNS, a.DerivesCKM, a.DerivesBFlavZero, a.AddsCarrier, a.AddsSelector, a.PreservesGate352, a.PreservesGate596, a.PreservesGate599, a.PreservesGate603, a.Verdict)
}
