package generation2environmentalflavorsealintegrationhistorytransportaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(a InheritedGate596) string {
	return fmt.Sprintf("BFlavExpr=%q BFlav=%.15g seal=%q parts=%q primary=%q closest=%q epsilonNative=%t bFlavNative=%t verdict=%q", a.BFlavExpression, a.BFlavValue, a.ChargedLeptonSealName, strings.Join(a.ChargedLeptonSealParts, ","), a.PrimaryObstruction, a.ClosestRoute, a.EpsilonNative, a.BFlavNative, a.Verdict)
}

func FormatFlavorSealRow(a FlavorSealRow) string {
	return fmt.Sprintf("seal=%q variable=%q object=%q equation=%q role=%q compressed=%q native=%q verdict=%q", a.Seal, a.HistoryVariable, a.Object, a.Equation, a.Role, a.CompressedQuantity, a.NativeStatus, a.Verdict)
}

func FormatSealTable(a IntegratedFlavorSealTable) string {
	items := make([]string, 0, len(a.Rows))
	for _, row := range a.Rows {
		items = append(items, FormatFlavorSealRow(row))
	}
	return fmt.Sprintf("rows=[%s] verdict=%q", strings.Join(items, "; "), a.Verdict)
}

func FormatEmbedding(a HistoryVariableEmbedding) string {
	return fmt.Sprintf("YCore=%q OmegaCore=%q TCore=%q YNative=%t OmegaNative=%t TNative=%t verdict=%q", strings.Join(a.YCore, ";"), strings.Join(a.OmegaCore, ";"), strings.Join(a.TCore, ";"), a.YCoreNative, a.OmegaNative, a.TNative, a.Verdict)
}

func FormatEndMap(a FlavorEndMap) string {
	return fmt.Sprintf("equation=%q inputs=%q compressed=%q raw=%q bridgeOnly=%t native=%t verdict=%q", a.Equation, strings.Join(a.Inputs, ";"), strings.Join(a.CompressedQuantities, ";"), strings.Join(a.RawEnvironmentalInputs, ";"), a.BridgeOnly, a.NativeDerivation, a.Verdict)
}

func FormatCompression(a CompressionLedger) string {
	return fmt.Sprintf("before=%q after=%q compressed=%q stillRaw=%q nativeCompression=%t verdict=%q", strings.Join(a.Before, ";"), strings.Join(a.After, ";"), strings.Join(a.CompressedBySeals, ";"), strings.Join(a.StillRaw, ";"), a.NativeCompression, a.Verdict)
}

func FormatMissingTheorem(a MissingTheorem) string {
	return fmt.Sprintf("name=%q requirements=%q reason=%q present=%t verdict=%q", a.Name, strings.Join(a.Requirements, ";"), a.Reason, a.Present, a.Verdict)
}

func FormatFirewalls(a FirewallAudit) string {
	return fmt.Sprintf("koide=%t pmns=%t ckm=%t chargedMasses=%t yukawaEigen=%t texture=%t bFlavNative=%t observedNative=%t carrier=%t selector=%t gate352=%t verdict=%q", a.DerivesKoide, a.DerivesPMNS, a.DerivesCKM, a.DerivesChargedLeptonMasses, a.DerivesYukawaEigenvalues, a.DerivesFlavorTexture, a.PromotesBFlavZero, a.PromotesObservedAsNative, a.AddsCarrier, a.AddsSelector, a.PreservesGate352, a.Verdict)
}

func FormatFinal(a FinalVerdict) string {
	return fmt.Sprintf("integrated=%t YCore=%t OmegaCore=%t TBridge=%t nativeFourthRoot=%t nativeBFlav=%t theorem=%q decision=%q verdict=%q", a.FlavorSealIntegrated, a.YCoreSharpened, a.OmegaCoreSharpened, a.TCoreBridgeOnly, a.NativeFourthRootTheorem, a.NativeBFlavZeroTheorem, a.ExactMissingTheorem, a.Decision, a.Verdict)
}
