package bminuslweakselector

import (
	"fmt"
	"strings"
)

func FormatVec(v []float64) string {
	parts := make([]string, len(v))
	for i, x := range v {
		parts[i] = fmt.Sprintf("%.6g", x)
	}
	return "(" + strings.Join(parts, ",") + ")"
}

func FormatInheritance(a Gate257Inheritance) string {
	return fmt.Sprintf("charges=%t witnesses=%t branches=%t gate257Plane=%t status=%q verdict=%q", a.NativeChargeEigenvaluesExtracted, a.EmbeddingWitnessesScanned, a.AllTrialityBranchesScanned, a.Gate257Neutral3PlaneDerived, a.Gate257Status, a.Verdict)
}

func FormatBMinusL(a BMinusLLedger) string {
	return fmt.Sprintf("expr=%q coeff=%s temporal=%d spatial=%v iso=%t split=%t derived=%t observed=%t verdict=%q", a.Expression, FormatVec(a.Coefficients), a.TemporalMode, a.SpatialModes, a.SpatialIsotropy, a.OnePlusThreeSplit, a.DerivedFiniteFockLedger, a.UsesObservedInput, a.Verdict)
}

func FormatScalarSieve(a ScalarSieveAudit) string {
	return fmt.Sprintf("input=%d survivors=%d reduced=%t unique=%t signDeg=%t survivorNames=%v rejected=%v verdict=%q", a.InputCount, a.SurvivorCount, a.Reduced, a.UniqueSelected, a.SignDegeneracyRemains, a.SurvivorNames, a.RejectedNames, a.Verdict)
}

func FormatWeakSieve(a WeakFrameSieveAudit) string {
	return fmt.Sprintf("input=%d survivors=%d reduced=%t unique=%t spatialDeg=%t survivorNames=%v rejected=%v verdict=%q", a.InputCount, a.SurvivorCount, a.Reduced, a.UniqueSelected, a.SpatialPlaneDegeneracyLeft, a.SurvivorNames, a.RejectedNames, a.Verdict)
}

func FormatCombined(a CombinedWitnessSieveAudit) string {
	return fmt.Sprintf("input=%d survivors=%d reduced=%t unique=%t verdict=%q", a.InputWitnessCount, a.SurvivingWitnessCount, a.Reduced, a.UniqueOrientation, a.Verdict)
}

func FormatRestrictedScan(a RestrictedTrialityAudit) string {
	return fmt.Sprintf("branches=%d results=%d exactPolarized3=%d exactFull3=%d maxPolarized=%d maxFull=%d uniqueBranch=%t selected=%q allScanned=%t afterSelector=%t verdict=%q", a.BranchCount, a.ResultCount, a.ExactPolarized3PlaneResults, a.ExactFull3KernelResults, a.MaxPolarizedKernelComplexDim, a.MaxFullQ8vCKernelComplexDim, a.UniqueBranchForPolarized3Plane, a.SelectedBranch, a.AllSurvivorsScanned, a.ScannedAfterSelector, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("gate257NoGo=%t beforeKernel=%t selectorNotOutcome=%t obsCharges=%t masses=%t yukawas=%t forcedWeak=%t forcedScalar=%t branchHand=%t forced3=%t yOnlyQ=%t sealFinite=%t vtau=%t yukawa=%t polluted=%t verdict=%q", a.Gate257NoGoPreserved, a.BMinusLAppliedBeforeKernel, a.BMinusLUsedAsSelectorNotOutcome, a.ImportedObservedChargeTable, a.ImportedObservedMasses, a.ImportedObservedYukawas, a.ForcedWeakPlane, a.ForcedScalarOrientation, a.SelectedTrialityByHand, a.ForcedKernelDim3, a.AcceptedYOnlyAsQ, a.TreatedSealAsFiniteDerivation, a.ConstructedVTauByHand, a.InsertedYukawaTexture, a.PollutedFiniteCore, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("gate257=%t bminusl=%t scalarReduced=%t weakReduced=%t combinedReduced=%t rescanned=%t uniqueEW=%t plane=%t branch=%t yukawa=%t status=%q next=%q comment=%q", a.Gate257Inherited, a.BMinusLLedgerRetrieved, a.ScalarSieveReduced, a.WeakFrameSieveReduced, a.CombinedWitnessSpaceReduced, a.RestrictedTrialityRescanned, a.UniqueEWOrientationSelected, a.Neutral3PlaneDerived, a.TrialityBranchSelected, a.YukawaTextureDerived, a.Status, a.NextGate, a.Comment)
}
