package generation2masterenvironmentalhistorysealvectoraudit

import (
	"fmt"
	"strings"
)

func FormatInherited(a InheritedGate604) string {
	return fmt.Sprintf("minimalFlavor=%t sigmaGauge=%t optionalFullOrder=%t formula=%q verdict=%q", a.MinimalFlavorSealDefined, a.SigmaGaugeForBFlav, a.OptionalFullOrderSeal, a.FlavorFormula, a.Verdict)
}

func FormatMasterSealRow(r MasterSealRow) string {
	return fmt.Sprintf("symbol=%q sector=%q class=%q source=%q obstruction=%q meaning=%q verdict=%q", r.Symbol, r.Sector, r.Classification, r.GateSource, r.NativeObstruction, r.Meaning, r.Verdict)
}

func FormatMasterSealTable(rows []MasterSealRow) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatMasterSealRow(r))
	}
	return strings.Join(parts, " | ")
}

func FormatSummary(s ClassificationSummary) string {
	return fmt.Sprintf("native=%d algebraic=%d bridge=%d environmental=%d observed=%d gauge=%d clear=%t verdict=%q", s.NativeCount, s.AlgebraicExtensionCount, s.BridgeSealCount, s.EnvironmentalSealCount, s.ObservedLedgerCount, s.GaugeConventionCount, s.BoundaryClear, s.Verdict)
}

func FormatFormula(f MasterFormula) string {
	return fmt.Sprintf("formula=%q native=%v algebraic=%v seals=%v bridge=%v observed=%v verdict=%q", f.Formula, f.NativeLawInputs, f.AlgebraicExtensions, f.HistorySeals, f.BridgeNormalizations, f.ObservedEndpointLedgers, f.Verdict)
}

func FormatSolvedUnsolvedRow(r SolvedUnsolvedRow) string {
	return fmt.Sprintf("item=%q sector=%q status=%q reason=%q verdict=%q", r.Item, r.Sector, r.Status, r.Reason, r.Verdict)
}

func FormatSolvedUnsolved(rows []SolvedUnsolvedRow) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatSolvedUnsolvedRow(r))
	}
	return strings.Join(parts, " | ")
}

func FormatRankingRow(r NextTargetRankingRow) string {
	return fmt.Sprintf("rank=%d path=%q value=%q status=%q rationale=%q recommendation=%q verdict=%q", r.Rank, r.Path, r.Value, r.CurrentStatus, r.Rationale, r.Recommendation, r.Verdict)
}

func FormatRanking(rows []NextTargetRankingRow) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatRankingRow(r))
	}
	return strings.Join(parts, " | ")
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("koide=%t flavor=%t ew=%t cosmo=%t endpoint=%t constants=%t g352=%t g596=%t g604=%t verdict=%q", f.DerivesKoide, f.DerivesFlavor, f.DerivesEWMasses, f.DerivesCosmology, f.DerivesObservedEndpoint, f.SearchesNewConstants, f.PreservesGate352, f.PreservesGate596, f.PreservesGate604, f.Verdict)
}
