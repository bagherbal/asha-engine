package generation2complexd4trilinearinvariantandyukawareadoutobstructionaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "GATE-802-COMPLEX-D4-TRILINEAR-INVARIANT-YUKAWA-READOUT-OBSTRUCTION"
	theoremName = "Gate 802 — Complex D4 Trilinear Invariant and Yukawa Readout Obstruction Audit"
)

func Generation2ComplexD4TrilinearInvariantAndYukawaReadoutObstructionAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 802 analysis", Passed: false, Detail: err.Error()}}, Notes: []string{StatusFirewallPreservedGate802}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 801 complex-airlocked status", Passed: a.Inheritance.ClAlgebra == "Cl(1,7) ≅ Mat(16,R)" && a.Inheritance.VolumeSquare == -1 && !a.Inheritance.RealChiralityCertified && strings.Contains(a.Inheritance.TrialityLevel, "T1") && containsAll(a.Inheritance.Verdicts, []string{StatusGate801Inherited, StatusComplexD4AirlockInherited, StatusT1NotNative, StatusT1NotYukawa}), Detail: a.Inheritance.TrialityLevel},
			{Name: "define three complex triality carriers", Passed: a.Carriers.Defined && a.Carriers.VectorDimC == 8 && a.Carriers.SpinorPlusDimC == 8 && a.Carriers.SpinorMinusDimC == 8 && !a.Carriers.GenerationCopies && !a.Carriers.NativeRealCarriers && containsAll(a.Carriers.Supports, []string{StatusThreeEightDimTypes}) && containsAll(a.Carriers.Failures, []string{StatusComplexTypesNotGenerations, StatusComplexNotNativeCarriers}), Detail: FormatCarriers(a.Carriers)},
			{Name: "define complex D4 trilinear invariant", Passed: a.Trilinear.Defined && strings.Contains(a.Trilinear.Formula, "T_D4") && strings.Contains(a.Trilinear.GammaAction, "S+_C -> S-_C") && a.Trilinear.Equivariant && a.Trilinear.NonZero && containsAll(a.Trilinear.Supports, []string{StatusTD4PreYukawa}) && containsAll(a.Trilinear.Failures, []string{StatusTD4NotSMYukawa, StatusTD4NotTraceLedger}), Detail: FormatTrilinear(a.Trilinear)},
			{Name: "audit invariant multiplicity", Passed: a.Multiplicity.Audited && a.Multiplicity.HomDimension == 1 && a.Multiplicity.CanonicalShapeUpToScale && !a.Multiplicity.DeterminesEigenvalues && !a.Multiplicity.DeterminesHierarchy && containsAll(a.Multiplicity.Supports, []string{StatusUniqueInvariantShape}) && containsAll(a.Multiplicity.Failures, []string{StatusUniqueNoEigenvalues, StatusNormalizationNoHierarchy}), Detail: FormatMultiplicity(a.Multiplicity)},
			{Name: "audit triality covariance", Passed: a.Covariance.Audited && a.Covariance.CyclicStable && !a.Covariance.GenerationTriplication && !a.Covariance.MixingReadout && containsAll(a.Covariance.Supports, []string{StatusTrialityCovariance}) && containsAll(a.Covariance.Failures, []string{StatusCovarianceNotGenerations, StatusCovarianceNotMixing}), Detail: a.Covariance.Verdict},
			{Name: "define TrialityYukawaReadoutPackage requirements", Passed: a.Readout.Defined && containsAll(a.Readout.Items, []string{"operator extraction T_D4 + symmetry-breaking data -> Y_u,Y_d,Y_e,Y_nu", "trace atom map y_i -> x_i=y_i^2", "breaking/deformation explaining hierarchy, top dominance, and N_eff-3"}) && containsAll(a.Readout.Failures, []string{StatusTrilinearNoSectorOps, StatusTrilinearNoAtoms, StatusTrilinearNoNEff}), Detail: FormatReadout(a.Readout)},
			{Name: "audit sector assignment obstruction", Passed: a.Sector.Audited && containsAll(a.Sector.Failures, []string{StatusNoTrialityToSMSectors, StatusThreeFramesNotFourSectors, StatusNoGaugeAssignment}), Detail: FormatObstruction(a.Sector)},
			{Name: "audit generation obstruction", Passed: a.Generation.Audited && containsAll(a.Generation.Failures, []string{StatusTrialityTypesNotGenerations, StatusNoGenerationCarrier, StatusNoPMNSCKMReadout}), Detail: FormatObstruction(a.Generation)},
			{Name: "audit positivity and singular-value obstruction", Passed: a.Positivity.Audited && containsAll(a.Positivity.Failures, []string{StatusComplexAmplitudeNotAtom, StatusNoHermitianSectorOperator, StatusNoSingularValueExtraction}), Detail: FormatObstruction(a.Positivity)},
			{Name: "audit top-dominance and rest-pressure obstruction", Passed: a.TopDominance.Audited && containsAll(a.TopDominance.Failures, []string{StatusTD4NoTopDominance, StatusTD4NoNEffMinusThree, StatusTD4NoScaleStability}), Detail: FormatObstruction(a.TopDominance)},
			{Name: "audit Georgi-Jarlskog readout obstruction", Passed: a.GeorgiJarlskog.Audited && containsAll(a.GeorgiJarlskog.Failures, []string{StatusTD4NoGJ, StatusTD4NoHighScaleClebsch, StatusGJStillNeedsLedger}), Detail: FormatObstruction(a.GeorgiJarlskog)},
			{Name: "preserve real-form obstruction", Passed: a.RealForm.Preserved && containsAll(a.RealForm.Requires, []string{"RealDescentMap from complex triality invariant to real Cl(1,7) typed object", "preserve positivity/reality of trace atoms", "preserve sector readout"}) && containsAll(a.RealForm.Failures, []string{StatusComplexTD4NotRealCL17, StatusNoRealDescentReadout}), Detail: FormatRealForm(a.RealForm)},
			{Name: "record lawful use of T_D4", Passed: a.Lawful.Recorded && containsAll(a.Lawful.Allowed, []string{"airlocked pre-Yukawa coupling shape", "guide representation-theoretic constraint search"}) && containsAll(a.Lawful.Blocked, []string{"derive N_eff", "derive PMNS/CKM", "modify C_Higgs"}) && containsAll(a.Lawful.Supports, []string{StatusTD4UsefulAirlocked}) && containsAll(a.Lawful.Failures, []string{StatusTD4NotPhysicalScalarInput}), Detail: FormatLawfulUse(a.Lawful)},
			{Name: "preserve C_Higgs firewall", Passed: a.CHiggs.Preserved && strings.Contains(a.CHiggs.Formula, "C_Higgs") && containsAll(a.CHiggs.Unchanged, []string{"N_eff", "C_Yukawa", "C_History", "C_Higgs"}) && containsAll(a.CHiggs.Failures, []string{StatusTD4NoCYukawaUpdate, StatusCHiggsStillLevelB}), Detail: FormatCHiggs(a.CHiggs)},
			{Name: "record outcome classification", Passed: a.Outcome.Recorded && len(a.Outcome.Items) == 4 && a.Outcome.Support == StatusComplexBranchInteresting, Detail: strings.Join(a.Outcome.Items, "; ")},
			{Name: "record branch decision", Passed: a.Branch.Recorded && strings.Contains(a.Branch.NextNative, "Gate 803") && strings.Contains(a.Branch.NextNative, "Triality-to-Yukawa") && strings.Contains(a.Branch.Alternative, "External Yukawa"), Detail: a.Branch.NextNative},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.NoYukawa && a.Firewalls.NoEigenvalues && a.Firewalls.NoPMNSCKM && a.Firewalls.NoFlavor && a.Firewalls.NoNEff && a.Firewalls.NoGJ && a.Firewalls.NoScalar && a.Firewalls.NoPoleMass && a.Firewalls.NoVEVGF && a.Firewalls.NoNativeCL17 && a.Firewalls.NoHistoryLoop && a.Firewalls.Verdict == StatusFirewallPreservedGate802, Detail: a.Firewalls.Verdict},
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
		notes := []string{a.Truth, FormatCarriers(a.Carriers), FormatTrilinear(a.Trilinear), FormatMultiplicity(a.Multiplicity), FormatReadout(a.Readout), FormatObstruction(a.Sector), FormatObstruction(a.Generation), FormatObstruction(a.Positivity), FormatObstruction(a.TopDominance), FormatObstruction(a.GeorgiJarlskog), FormatRealForm(a.RealForm), FormatLawfulUse(a.Lawful), FormatCHiggs(a.CHiggs), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
