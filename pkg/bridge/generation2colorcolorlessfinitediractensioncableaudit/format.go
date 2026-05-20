package generation2colorcolorlessfinitediractensioncableaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(a InheritedGate597) string {
	return fmt.Sprintf("BFlavExpr=%q BFlav=%.15g chargedSeal=%q orientation=%q obstruction=%q environmentalOnly=%t verdict=%q", a.BFlavExpression, a.BFlavValue, a.ChargedLeptonSeal, a.OrientationBalance, a.PrimaryObstruction, a.EnvironmentalOnly, a.Verdict)
}

func FormatDFSectorRow(a DFSectorRow) string {
	return fmt.Sprintf("sector=%q carrier=%q blocks=%q yukawa=%q edges=%q color=%q native=%q verdict=%q", a.Sector, a.Carrier, strings.Join(a.Blocks, ","), strings.Join(a.YukawaBlocks, ","), strings.Join(a.OneFormEdges, ","), a.ColorMultiplicity, a.NativeStatus, a.Verdict)
}

func FormatDFSectorSplit(a DFSectorSplitTable) string {
	rows := make([]string, 0, len(a.Rows))
	for _, row := range a.Rows {
		rows = append(rows, FormatDFSectorRow(row))
	}
	return fmt.Sprintf("algebra=%q decomposition=%q interSector=%t rows=[%s] verdict=%q", a.FiniteAlgebra, a.Decomposition, a.InterSectorDFBlock, strings.Join(rows, "; "), a.Verdict)
}

func FormatOneFormEdge(a OneFormEdge) string {
	return fmt.Sprintf("edge=%q sector=%q yukawa=%q separated=%t crossInvariant=%t verdict=%q", a.Edge, a.Sector, a.YukawaBlock, a.BlockSeparated, a.ProducesCrossSectorInvariant, a.Verdict)
}

func FormatEdges(a EdgeInventory) string {
	rows := make([]string, 0, len(a.Edges))
	for _, edge := range a.Edges {
		rows = append(rows, FormatOneFormEdge(edge))
	}
	return fmt.Sprintf("edges=[%s] verdict=%q", strings.Join(rows, "; "), a.Verdict)
}

func FormatCandidateRow(a CandidateInvariantRow) string {
	return fmt.Sprintf("candidate=%q lane=%q quark=%t leptonProjector=%t rootChamber=%t native=%t verdict=%q reason=%q", a.Candidate, a.NativeLane, a.SeesQuarkOrientation, a.SeesLeptonProjector, a.SeesChargedRootChamber, a.Native, a.Verdict, a.Reason)
}

func FormatCandidateTable(a CandidateInvariantTable) string {
	rows := make([]string, 0, len(a.Rows))
	for _, row := range a.Rows {
		rows = append(rows, FormatCandidateRow(row))
	}
	return fmt.Sprintf("rows=[%s] nativeCable=%t colorVisible=%t verdict=%q", strings.Join(rows, "; "), a.AnyNativeTensionCable, a.ColorColorlessStructureVisible, a.Verdict)
}

func FormatRootRow(a RootObstructionRow) string {
	return fmt.Sprintf("route=%q H14=%t rootTrace=%t epsilon=%t chamber=%t links=%t bflav=%t verdict=%q reason=%q", a.Route, a.ProducesHEOneFourth, a.ProducesRootTrace, a.ProducesEpsilonHE, a.SelectsCanonicalChamber, a.LinksPMNSAndCKM, a.ProvesBFlavZero, a.Verdict, a.Reason)
}

func FormatRootLedger(a RootObstructionLedger) string {
	rows := make([]string, 0, len(a.Rows))
	for _, row := range a.Rows {
		rows = append(rows, FormatRootRow(row))
	}
	return fmt.Sprintf("rows=[%s] gate596Avoided=%t verdict=%q", strings.Join(rows, "; "), a.Gate596Avoided, a.Verdict)
}

func FormatOutcome(a FiniteDiracOutcome) string {
	return fmt.Sprintf("outcome=%q nativeSuccess=%t conditional=%t obstruction=%t missing=%q decision=%q verdict=%q", a.OutcomeName, a.NativeSuccess, a.ConditionalStructure, a.FullObstruction, a.ExactMissingObject, a.Decision, a.Verdict)
}

func FormatFirewalls(a FirewallAudit) string {
	return fmt.Sprintf("koide=%t masses=%t pmns=%t ckm=%t yukawa=%t neutrino=%t texture=%t bflav=%t rootNative=%t carrier=%t selector=%t gate352=%t gate596=%t gate597=%t verdict=%q", a.DerivesKoide, a.DerivesChargedLeptonMasses, a.DerivesPMNS, a.DerivesCKM, a.DerivesYukawaEigenvalues, a.DerivesNeutrinos, a.DerivesFlavorTexture, a.PromotesBFlavZero, a.PromotesRootChamberNative, a.AddsCarrier, a.AddsSelector, a.PreservesGate352, a.PreservesGate596, a.PreservesGate597, a.Verdict)
}

func FormatFinal(a FinalVerdict) string {
	return fmt.Sprintf("sectorSplit=%t colorVisible=%t quarkJ=%t leptonProjector=%t rootNative=%t nativeCable=%t bflavNative=%t missing=%q decision=%q verdict=%q", a.SectorSplitNative, a.ColorColorlessVisible, a.QuarkCommutatorVisible, a.LeptonProjectorVisible, a.RootChamberNative, a.NativeTensionCableFound, a.BFlavNative, a.MissingObject, a.Decision, a.Verdict)
}
