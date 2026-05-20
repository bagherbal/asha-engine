package generation2trialitytoyukawareadoutminimalityandnogoaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "GATE-803-TRIALITY-TO-YUKAWA-READOUT-MINIMALITY-NO-GO"
	theoremName = "Gate 803 — Triality-to-Yukawa Readout Package Minimality and No-Go Audit"
)

func Generation2TrialityToYukawaReadoutMinimalityAndNoGoAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Notes: []string{err.Error()}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 802 trilinear obstruction", Passed: a.Inheritance.AirlockLevel == "T1 — complex D4 triality only" && !a.Inheritance.HasNativeRealDescent && !a.Inheritance.HasSectorOperators && !a.Inheritance.HasTraceAtoms && containsAll(a.Inheritance.Verdicts, []string{StatusGate802Inherited, StatusTD4OnlyPreYukawa, StatusTD4NotTraceLedger}), Detail: a.Inheritance.TD4Status},
			{Name: "define TrialityYukawaReadoutPackage", Passed: a.Package.Defined && len(a.Package.Seals) == 10 && containsAll(a.Package.Seals, []string{"RealDescentSeal", "GaugeRepresentationAssignmentSeal", "SectorAssignmentSeal", "GenerationCarrierSeal", "HermitianOperatorSeal", "SymmetryBreakingHierarchySeal", "TraceAtomExtractionSeal", "ColorMultiplicitySeal", "ScaleSchemeSeal", "NonCircularitySeal"}) && containsAll(a.Package.TargetChain, []string{"Y_u,Y_d,Y_e,Y_nu", "a,b,N_eff"}) && containsAll(a.Package.Supports, []string{StatusNeedsExtraSeals}) && containsAll(a.Package.Failures, []string{StatusTrilinearAloneNoPackage}), Detail: FormatPackage(a.Package)},
			{Name: "audit RealDescentSeal", Passed: a.RealDescent.Audited && containsAll(a.RealDescent.Required, []string{"real structure", "positivity/reality of later trace atoms"}) && containsAll(a.RealDescent.Failures, []string{StatusNoRealDescent, StatusComplexNotNativeYukawa}), Detail: FormatSeal(a.RealDescent)},
			{Name: "audit gauge representation assignment", Passed: a.GaugeAssignment.Audited && containsAll(a.GaugeAssignment.Required, []string{"Q_L -> u_R", "SU(3)c/SU(2)L/U(1)Y compatibility", "Higgs doublet interface"}) && containsAll(a.GaugeAssignment.Failures, []string{StatusNoGaugeAssignment, StatusTrialityFramesNoSMEdges}), Detail: FormatSeal(a.GaugeAssignment)},
			{Name: "audit sector assignment minimality", Passed: a.SectorAssignment.Audited && containsAll(a.SectorAssignment.Required, []string{"Y_u,Y_d,Y_e,Y_nu", "three triality types versus four Yukawa sectors"}) && containsAll(a.SectorAssignment.Failures, []string{StatusThreeTypesNotFourSectors, StatusNoSMSectorAssignment}), Detail: FormatSeal(a.SectorAssignment)},
			{Name: "audit generation carrier requirement", Passed: a.GenerationCarrier.Audited && containsAll(a.GenerationCarrier.Required, []string{"G_gen", "cardinality three", "frame comparison rules"}) && containsAll(a.GenerationCarrier.Failures, []string{StatusTrialityTypesNotGenerations, StatusNoGenerationCarrier, StatusNoPMNSCKMFrames}), Detail: FormatSeal(a.GenerationCarrier)},
			{Name: "audit HermitianOperatorSeal", Passed: a.HermitianOperator.Audited && containsAll(a.HermitianOperator.Required, []string{"Y_f†Y_f positive", "singular-value extraction"}) && containsAll(a.HermitianOperator.Failures, []string{StatusComplexNotHermitian, StatusNoSVDTheorem}), Detail: FormatSeal(a.HermitianOperator)},
			{Name: "audit symmetry breaking and hierarchy", Passed: a.Hierarchy.Audited && containsAll(a.Hierarchy.Required, []string{"top-dominance mechanism", "rest-pressure mechanism"}) && containsAll(a.Hierarchy.Failures, []string{StatusUniqueNoHierarchy, StatusTD4NoTopDominance, StatusTD4NoNEffMinusThree}), Detail: FormatSeal(a.Hierarchy)},
			{Name: "audit trace atom extraction", Passed: a.TraceAtom.Audited && containsAll(a.TraceAtom.Required, []string{"trace atoms x_i=y_i^2", "no backward solving"}) && containsAll(a.TraceAtom.Failures, []string{StatusNoPositiveAtoms, StatusNoBackwardAtoms}), Detail: FormatSeal(a.TraceAtom)},
			{Name: "audit color multiplicity", Passed: a.Color.Audited && containsAll(a.Color.Required, []string{"color factor 3", "no double counting"}) && containsAll(a.Color.Supports, []string{StatusColorSU3TraceMultiplicity}) && containsAll(a.Color.Failures, []string{StatusD4DoesNotReplaceColor, StatusColorDoubleCount}), Detail: FormatSeal(a.Color)},
			{Name: "audit scale scheme", Passed: a.Scale.Audited && containsAll(a.Scale.Required, []string{"scale_mu", "renormalization scheme", "neutrino convention"}) && containsAll(a.Scale.Failures, []string{StatusNoScaleLedger, StatusNoScaleStability}), Detail: FormatSeal(a.Scale)},
			{Name: "audit noncircularity", Passed: a.NonCircularity.Audited && containsAll(a.NonCircularity.Required, []string{"N_eff", "C_Higgs", "observed Higgs mass"}) && containsAll(a.NonCircularity.Failures, []string{StatusNoTargetTuning}), Detail: FormatSeal(a.NonCircularity)},
			{Name: "audit minimality removal failures", Passed: a.Minimality.Audited && len(a.Minimality.RemovalFailures) == 10 && containsAll(a.Minimality.Supports, []string{StatusAllSealsNonCosmetic}) && containsAll(a.Minimality.Failures, []string{StatusCannotCompressToTD4}), Detail: FormatMinimality(a.Minimality)},
			{Name: "define triality-to-Yukawa no-go", Passed: a.NoGo.Defined && containsAll(a.NoGo.GivenOnly, []string{"ComplexD4TrialityAirlock", "T_D4"}) && containsAll(a.NoGo.CannotConstruct, []string{"Y_u,Y_d,Y_e,Y_nu", "positive trace atoms x_i", "a,b,N_eff", "PMNS/CKM", "Georgi-Jarlskog ratios", "C_Higgs update"}) && strings.Contains(a.NoGo.Reason, "not a sector-labeled") && containsAll(a.NoGo.Failures, []string{StatusTD4AloneNoLedger, StatusTD4AloneNoNEffFlavor}), Detail: FormatNoGo(a.NoGo)},
			{Name: "audit current ASHA readout subobjects", Passed: a.Current.Audited && containsAll(a.Current.Supplies, []string{"finite spectral triple allowed chiral edge shapes and trace templates", "color SU(3) trace multiplicity"}) && containsAll(a.Current.DoesNotSupply, []string{"triality Yukawa readout package", "trace atoms from T_D4"}) && containsAll(a.Current.Supports, []string{StatusFSTEdgeTemplate, StatusColorSU3TraceMultiplicity}) && containsAll(a.Current.Failures, []string{StatusNoCurrentReadoutPackage}), Detail: FormatCurrent(a.Current)},
			{Name: "separate empirical and native Yukawa paths", Passed: a.Paths.Recorded && containsAll(a.Paths.EmpiricalPath, []string{"ExternalYukawaLedgerSeal", "N_eff audit"}) && containsAll(a.Paths.NativePath, []string{"T_D4", "TrialityYukawaReadoutPackage"}) && containsAll(a.Paths.Supports, []string{StatusExternalFastest}) && containsAll(a.Paths.Failures, []string{StatusExternalNotNative, StatusTrialityNotReadyExternal}), Detail: FormatPaths(a.Paths)},
			{Name: "preserve C_Higgs firewall", Passed: a.CHiggs.Preserved && strings.Contains(a.CHiggs.Formula, "C_Higgs") && containsAll(a.CHiggs.Unchanged, []string{"N_eff", "C_Yukawa", "C_History", "C_Higgs"}) && containsAll(a.CHiggs.Failures, []string{StatusNoCYukawaUpdate, StatusCHiggsLevelB}), Detail: FormatCHiggs(a.CHiggs)},
			{Name: "record branch decision", Passed: a.Branch.Recorded && strings.Contains(a.Branch.NextNative, "Gate 804") && strings.Contains(a.Branch.NextNative, "Finite Spectral Triple") && containsAll(a.Branch.Supports, []string{StatusNextFSTCompatibility}), Detail: a.Branch.NextNative},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.NoYukawa && a.Firewalls.NoEigenvalues && a.Firewalls.NoPMNSCKM && a.Firewalls.NoFlavor && a.Firewalls.NoNEff && a.Firewalls.NoGJ && a.Firewalls.NoScalar && a.Firewalls.NoPoleMass && a.Firewalls.NoVEVGF && a.Firewalls.NoNativeTriality && a.Firewalls.NoHistoryLoop && a.Firewalls.Verdict == StatusFirewallGate803, Detail: a.Firewalls.Verdict},
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
		notes := []string{a.Truth, FormatPackage(a.Package), FormatSeal(a.RealDescent), FormatSeal(a.GaugeAssignment), FormatSeal(a.SectorAssignment), FormatSeal(a.GenerationCarrier), FormatSeal(a.HermitianOperator), FormatSeal(a.Hierarchy), FormatSeal(a.TraceAtom), FormatSeal(a.Color), FormatSeal(a.Scale), FormatSeal(a.NonCircularity), FormatMinimality(a.Minimality), FormatNoGo(a.NoGo), FormatCurrent(a.Current), FormatPaths(a.Paths), FormatCHiggs(a.CHiggs), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
