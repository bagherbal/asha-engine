package generation2boundaryrawmomentresponsecoordinatenaturalityaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2BoundaryRawMomentResponseCoordinateNaturalityAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 732 — Boundary Raw-Moment Response Coordinate-Naturality Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate732 raw moment coordinate audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate731 cubic coefficient source audit", Passed: a.Gate731.Inherited && a.Gate731.MomentPolynomialAvailable && a.Gate731.NoNativeMomentExpansion && near(a.Gate731.DoubleEventWeight, 2*a.Gate731.P_K7, 1e-18) && strings.Contains(a.Gate731.Verdict, StatusGate731CubicCoefficientSourceInherited), Detail: FormatGate731(a.Gate731)},
			{Name: "rewrite raw moment response function", Passed: a.RawMoment.UsesRawMoments && near(a.RawMoment.M1Wall, a.Gate731.P_K7*a.Gate731.SSplit, 1e-18) && near(a.RawMoment.M2Wall, a.Gate731.P_K7*a.Gate731.SSplit*a.Gate731.SSplit, 1e-18) && near(a.RawMoment.M3Wall, a.Gate731.P_K7*a.Gate731.SSplit*a.Gate731.SSplit*a.Gate731.SSplit, 1e-18) && strings.Contains(a.RawMoment.FactoredFunction, "p_K7 S_split") && strings.Contains(a.RawMoment.Verdict, StatusRawMomentResponseFunctionRewritten), Detail: FormatRawMoment(a.RawMoment)},
			{Name: "record projector-power degeneracy", Passed: a.Degeneracy.AllPowersSupportedOnK7 && !a.Degeneracy.IndependentOperatorDirections && a.Degeneracy.ScalarResponseFunctionOnly && strings.Contains(a.Degeneracy.Verdict, StatusProjectorPowersNoIndependentDirections), Detail: FormatDegeneracy(a.Degeneracy)},
			{Name: "audit variance coordinate", Passed: near(a.Variance.VarianceWall, a.Gate731.P_K7*(1-a.Gate731.P_K7)*a.Gate731.SSplit*a.Gate731.SSplit, 1e-18) && near(a.Variance.CoefficientInVariance, a.Gate731.EWall/a.Variance.VarianceWall, 1e-18) && a.Variance.RawM2CloserToKappaE && a.Variance.TypedButInactive && strings.Contains(a.Variance.Verdict, StatusVarianceCoordinateNotActive), Detail: FormatVariance(a.Variance)},
			{Name: "audit central third moment coordinate", Passed: near(a.CentralM3.Mu3Wall, a.Gate731.P_K7*(1-a.Gate731.P_K7)*(1-2*a.Gate731.P_K7)*a.Gate731.SSplit*a.Gate731.SSplit*a.Gate731.SSplit, 1e-18) && a.CentralM3.RawCompressesBetter && strings.Contains(a.CentralM3.Verdict, StatusCentralMomentNotSelectedOverRawM3), Detail: FormatCentralM3(a.CentralM3)},
			{Name: "compare raw M3 against central moment", Passed: a.Comparison.RawSelectedByCurrentCompression && a.Comparison.CentralResidualAbs > a.Comparison.RawResidualAbs && a.Comparison.ImprovementFactor > 10 && strings.Contains(a.Comparison.Verdict, StatusRawM3CoordinateBestCompressesCurrentResidual), Detail: FormatComparison(a.Comparison)},
			{Name: "record source-type interpretation", Passed: strings.Contains(a.SourceType.Compact, "M1+kappa_e M2-2p_K7 M3") && strings.Contains(a.SourceType.Verdict, StatusSourceTypeInterpretationRecorded), Detail: FormatSourceType(a.SourceType)},
			{Name: "preserve coordinate-naturality firewall", Passed: !a.Firewall.RawMomentsNativelySelected && !a.Firewall.BoundaryMomentTheoremNative && !a.Firewall.ScalarRuntimeTheoremNative && !a.Firewall.HiggsMassTheoremNative && !a.Firewall.YukawaTheoremNative && strings.Contains(a.Firewall.Verdict, StatusNoNativeRawMomentResponseCoordinateTheorem) && strings.Contains(a.Firewall.Verdict, StatusGate732Boundary), Detail: FormatFirewall(a.Firewall)},
		}
		ok := true
		for _, c := range checks {
			if !c.Passed {
				ok = false
				break
			}
		}
		status := theorem.BridgeRequired
		if !ok {
			status = theorem.FailedRoute
		}
		notes := append([]string{a.Truth}, Statuses()...)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
