package fockrepresentationtrace

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func FockRepresentationTraceGaugeRatioYukawaAmplitudeSeparationTheorem() theorem.Theorem {
	const id = "BRIDGE-FOCK-REPRESENTATION-TRACE-GAUGE-RATIO-YUKAWA-SEPARATION"
	const name = "Fock representation-trace gauge ratio and Yukawa-amplitude separation theorem"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Fock representation-trace theorem", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 166 unit-incidence boundary reproduction is inherited but reinterpreted", Passed: a.Previous.Gauge.BoundaryDiagMatched && a.Previous.Gauge.WeakAngleSeedMatched && a.Previous.Gauge.RepresentationTraceOnly && a.Previous.AmplitudeSensitivity.ArbitraryAmplitudesAllowedByPriorGates && !a.Previous.AmplitudeSensitivity.BoundaryRatioStable, Detail: fmt.Sprintf("Gate166 unit ratio=%.12g deformed ratio=%.12g unit sin2=%.12g deformed sin2=%.12g", a.Previous.AmplitudeSensitivity.UnitRatioYOverSU2, a.Previous.AmplitudeSensitivity.DeformedRatioYOverSU2, a.Previous.AmplitudeSensitivity.UnitWeakAngle, a.Previous.AmplitudeSensitivity.DeformedWeakAngle)},
			{Name: "correct gauge functional is the amplitude-independent representation trace", Passed: a.TraceAudit.FermionStates == 16 && a.TraceAudit.LeftStates == 8 && a.TraceAudit.RightStates == 8 && a.TraceAudit.SU2Doublets == 4 && a.TraceAudit.KSU2T1.Equal(NewRational(2, 1)) && a.TraceAudit.KSU2T2.Equal(NewRational(2, 1)) && a.TraceAudit.KSU2T3.Equal(NewRational(2, 1)) && a.TraceAudit.KU1Y.Equal(NewRational(10, 3)) && a.TraceAudit.AmplitudeIndependent && !a.TraceAudit.UsesDiracFourthPower && !a.TraceAudit.UsesObservedInput && !a.TraceAudit.UsesContactModeClassifier, Detail: FormatTraceAudit(a.TraceAudit) + " :: sectors " + FormatSectorTraces(a.SectorTraces)},
			{Name: "embedded boundary normalization and weak-angle seed are representation-trace invariants", Passed: a.TraceAudit.BoundaryDiagMatched && a.TraceAudit.WeakAngleSeedMatched && a.TraceAudit.NormalizedY.Equal(NewRational(5, 3)) && a.TraceAudit.WeakAngleSeed.Equal(NewRational(3, 8)), Detail: "normalized diag=(" + a.TraceAudit.NormalizedT1.String() + "," + a.TraceAudit.NormalizedT2.String() + "," + a.TraceAudit.NormalizedT3.String() + "," + a.TraceAudit.NormalizedY.String() + "), sin2=" + a.TraceAudit.WeakAngleSeed.String()},
			{Name: "D_F^4 weighted trace is demoted to a diagnostic, not the gauge kinetic functional", Passed: a.Separation.DWeightedFunctionalAmplitudeDependent && a.Separation.RepresentationTraceAmplitudeIndependent && a.Separation.DWeightedDemotedToDiagnostic && a.Separation.GaugeKineticFunctionalCorrected && a.Separation.RepresentationUnitRatio.Equal(a.Separation.RepresentationDeformedRatio) && a.Separation.RepresentationUnitSin2.Equal(a.Separation.RepresentationDeformedSin2), Detail: FormatSeparation(a.Separation)},
			{Name: "Dirac amplitudes are the finite Yukawa/mass texture variables", Passed: a.YukawaAudit.YukawaChannels == 8 && a.YukawaAudit.OneGenerationAmplitudeSlots == 8 && a.YukawaAudit.FermionKindBlocks == 4 && !a.YukawaAudit.NumericalAmplitudesDerived && !a.YukawaAudit.ColorUniversalAmplitudesDerived && a.YukawaAudit.MassEigenvaluesAreSingularValues && a.YukawaAudit.CKMFromLeftMisalignment && a.YukawaAudit.PMNSFromLeftMisalignment && a.YukawaAudit.ConnectsGate28TextureSearch && !a.YukawaAudit.PhysicalMassesDerived && !a.YukawaAudit.MixingMatricesDerived, Detail: FormatYukawaAudit(a.YukawaAudit) + " :: " + a.YukawaAudit.OneGenerationSpectrumRule + " :: " + a.YukawaAudit.TrialityReplicatedProblem},
			{Name: "gauge side closes while physical mass/running side remains sealed", Passed: a.Firewall.BoundaryGaugeRatioClosed && a.Firewall.BoundaryWeakAngleSeedClosed && !a.Firewall.ContactModeClassificationNeededForBoundary && !a.Firewall.ContactModeClassificationSolved && !a.Firewall.ThresholdCorrectionsDerived && !a.Firewall.RGRunningDerived && !a.Firewall.PhysicalCouplingsDerived && !a.Firewall.PhysicalMassesDerived && !a.Firewall.CKMPMNSDerived && a.Firewall.YukawaTextureProblemOpened && a.Firewall.ResidualNullityBefore == 3 && a.Firewall.ResidualNullityAfter == 3, Detail: FormatFirewall(a.Firewall) + " :: " + a.TruthStatement},
		}, Notes: []string{
			"Gate 167 separates the two meanings that Gate 166 intentionally kept adjacent: representation trace fixes the embedded gauge boundary ratio, while D_F amplitudes are Yukawa/mass data.",
			"The sector ratio diag(1,1,1,5/3) and sin^2_*=3/8 are now amplitude-independent one-generation Fock charge-table theorems, not unit-incidence coincidences.",
			"No physical couplings, RG running, thresholds, fermion masses, CKM, or PMNS entries are derived; the open problem is now a finite Dirac/Yukawa texture eigenvalue problem linked to Gates 28-36.",
		}}
	}}
}
