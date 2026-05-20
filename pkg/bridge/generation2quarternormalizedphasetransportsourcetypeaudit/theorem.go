package generation2quarternormalizedphasetransportsourcetypeaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2QuarterNormalizedPhaseTransportSourceTypeAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 723 — Quarter-Normalized Phase Transport Source-Type Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate723 quarter phase transport audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate722 Higgs socket HistoryLoop transport", Passed: a.Gate722.Inherited && a.Gate722.SocketInterfacesWithOneFormLane && a.Gate722.ScalarProxyInterfacesWithHistoryLoop && a.Gate722.OneOver8PiAfterScalarProxyNotRepresentation && a.Gate722.BoundaryHistoryCompatible && a.Gate722.NoNativeHistoryLoopUnit && a.Gate722.NoNativeScalarProxyRuntime && a.Gate722.NoHiggsMassOrPoleMass && a.Gate722.NoYukawa && a.Gate722.NAndQSealedNotDerived && a.Gate722.Verdict == StatusGate722HiggsSocketHistoryLoopTransportInherited, Detail: FormatGate722(a.Gate722)},
			{Name: "audit phase-loop source candidate", Passed: strings.Contains(a.PhaseLoop.PhaseLine, "J_H") && strings.Contains(a.PhaseLoop.PhaseAction, "exp") && near(a.PhaseLoop.PhaseLoopUnit, 1/(2*math.Pi), 1e-18) && a.PhaseLoop.Candidate && !a.PhaseLoop.NativeHistoryTransportUsesMeasure && strings.Contains(a.PhaseLoop.Verdict, StatusPhaseLoopSourceCandidateAudited), Detail: FormatPhaseLoop(a.PhaseLoop)},
			{Name: "audit quarter normalization candidate", Passed: a.Quarter.RealCarrierDimension == 4 && a.Quarter.ComplexCarrierDimension == 2 && near(a.Quarter.QuarterFactor, 0.25, 1e-18) && near(a.Quarter.CandidateValue, 1/(8*math.Pi), 1e-18) && a.Quarter.EqualsHistoryLoopUnit && !a.Quarter.ScalarTransportAveragesOverFourComponents && strings.Contains(a.Quarter.Verdict, StatusQuarterNormalizationCandidateAudited), Detail: FormatQuarter(a.Quarter)},
			{Name: "reconstruct L as quarter phase unit", Passed: near(a.Quarter.CandidateValue, a.Ledger.LCandidate, 1e-18) && strings.Contains(a.Quarter.CandidateFormula, "1/(8*pi)") && strings.Contains(a.Quarter.Verdict, StatusLEqualsOneOver8PiReconstructedAsQuarterPhase), Detail: FormatQuarter(a.Quarter)},
			{Name: "audit scalar proxy transport placement", Passed: a.Placement.BelongsAfterScalarProxy && !a.Placement.DerivedFromRepresentationSocketAlone && strings.Contains(a.Placement.Chain, "lambda_proxy") && strings.Contains(a.Placement.Verdict, StatusScalarProxyTransportRoleAudited), Detail: FormatPlacement(a.Placement)},
			{Name: "audit q normalization firewall", Passed: a.QFirewall.QRescalesPhysicalChargeGenerator && a.QFirewall.GeometricCircleUnitIndependentOfQ && !a.QFirewall.NativeRelationQToL && strings.Contains(a.QFirewall.Verdict, StatusQDoesNotSourceL), Detail: FormatQFirewall(a.QFirewall)},
			{Name: "audit n selector firewall", Passed: a.NFirewall.PhaseLineDependsOnN && a.NFirewall.LoopMeasureUniformOverTwistorLines && !a.NFirewall.LSelectsN && strings.Contains(a.NFirewall.Verdict, StatusLDoesNotSelectN), Detail: FormatNFirewall(a.NFirewall)},
			{Name: "audit 7 over 72 firewall", Passed: near(a.SevenFirewall.EventProbability, float64(k7Dim)/float64(h72Dim), 1e-18) && near(a.SevenFirewall.LoopUnit, 1/(8*math.Pi), 1e-18) && !a.SevenFirewall.SameObject && !a.SevenFirewall.SevenOver72SourcesOneOver8Pi && strings.Contains(a.SevenFirewall.Verdict, Status7Over72DoesNotSourceOneOver8Pi), Detail: FormatSevenFirewall(a.SevenFirewall)},
			{Name: "record numerical scalar matching ledger", Passed: near(a.Ledger.LCandidate, 1/(8*math.Pi), 1e-18) && a.Ledger.RhoLambdaMatch > 0.038 && a.Ledger.RhoLambdaMatch < 0.039 && a.Ledger.KappaLambda > 0.044 && a.Ledger.KappaLambda < 0.045 && math.Abs(a.Ledger.TransportResidual) < 1e-15 && strings.Contains(a.Ledger.ScalarMatchingDeficitForm, "L(1-kappa_lambda)"), Detail: FormatLedger(a.Ledger)},
			{Name: "preserve source-type firewalls", Passed: !a.Firewall.NativeHistoryLoopUnitSourceTheorem && !a.Firewall.NativeScalarProxyToRuntimeTheorem && !a.Firewall.HiggsMassOrPoleMassTheorem && !a.Firewall.YukawaOperatorOrEigenvalueTheorem && !a.Firewall.NativeQSource && !a.Firewall.NativeNSelector && !a.Firewall.Native7Over72ToLTheorem && strings.Contains(a.Firewall.Verdict, StatusGate723Boundary), Detail: FormatFirewall(a.Firewall)},
		}
		notes := append(Statuses(), a.Truth)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
