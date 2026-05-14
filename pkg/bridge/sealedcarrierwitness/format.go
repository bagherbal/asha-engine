package sealedcarrierwitness

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

func FormatCharges(a ChargeExtractionAudit) string {
	parts := make([]string, 0, len(a.Sources))
	for _, s := range a.Sources {
		vec := "∅"
		if len(s.CoefficientVector) > 0 {
			vec = FormatVec(s.CoefficientVector)
		}
		parts = append(parts, fmt.Sprintf("%s source=%q carrier=%q expr=%q eigen=%t coeff=%s role=%q seal=%t observed=%t verdict=%q", s.Name, s.Source, s.Carrier, s.Expression, s.EigenvaluesDerived, vec, s.CoefficientVectorRole, s.RequiresEmbeddingSeal, s.UsesObservedInput, s.Verdict))
	}
	return fmt.Sprintf("B-L=%t scalarY=%t T3L=%t table=%t directT3=%t directY=%t external=%t verdict=%q sources=[%s]", a.BMinusLFockLedgerDerived, a.ScalarYphiEigenvaluesDerived, a.T3LLeftDoubletEigenvaluesDerived, a.ChargeEigenvalueTableDerived, a.PhysicalT3LDirectSCVector, a.PhysicalYPhiDirectSCVector, a.ExternalChargeInputUsed, a.Verdict, strings.Join(parts, "; "))
}

func FormatEmbedding(a EmbeddingWitnessAudit) string {
	return fmt.Sprintf("seal=%q weak=%d scalar=%d combined=%d nativeWeak=%t nativeScalar=%t allSealed=%t masses=%t yukawas=%t verdict=%q", a.SealName, a.WeakFrameCount, a.ScalarEmbeddingCount, a.TotalCombinedWitnesses, a.NativeWeakFrameSelected, a.NativeScalarEmbeddingSelected, a.AllWitnessesSealed, a.UsesObservedMasses, a.UsesObservedYukawas, a.Verdict)
}

func FormatSO8(a SO8WitnessAudit) string {
	return fmt.Sprintf("witt=%t witnesses=%d translated=%t bivectors=%v verdict=%q", a.WittDictionaryInherited, a.WitnessCount, a.AllTranslated, a.CartanBivectors, a.Verdict)
}

func FormatBranches(branches []TrialityBranch) string {
	parts := make([]string, 0, len(branches))
	for _, b := range branches {
		parts = append(parts, fmt.Sprintf("%s orth=%t det=%t admissible=%t", b.Name, b.Orthogonal, b.DetAbsOne, b.AdmissibleCartan))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatTrialityScan(a TrialityScanAudit) string {
	return fmt.Sprintf("branches=%d results=%d exactPolarized3=%d exactFull3=%d maxPolarized=%d maxFull=%d unique=%t selected=%q byOutcome=%t allScanned=%t yOnly={%s} verdict=%q branchHits=%v", a.BranchCount, a.ResultCount, a.ExactPolarized3PlaneResults, a.ExactFull3KernelResults, a.MaxPolarizedKernelComplexDim, a.MaxFullQ8vCKernelComplexDim, a.UniqueBranchForPolarized3Plane, a.SelectedBranch, a.SelectedByKernelOutcome, a.AllBranchesScanned, FormatYOnly(a.YOnly), a.Verdict, sortedBranchNames(a.Results))
}

func FormatYOnly(y YOnlyDiagnostic) string {
	return fmt.Sprintf("run=%t scalar=%q branch=%q transformed=%s polarized=%d would3=%t rejected=%t verdict=%q", y.Run, y.ScalarEmbeddingName, y.BranchName, FormatVec(y.TransformedCoefficients), y.PolarizedKernelComplexDim, y.WouldGivePolarizedThreeSlot, y.RejectedBecauseMissingT3L, y.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("nativeNoGo=%t derivedCharges=%t sealedEmbedding=%t obsCharges=%t masses=%t yukawas=%t weakPlane=%t branchHand=%t branchKernel=%t forced3=%t yOnlyQ=%t sealAsFinite=%t vtau=%t yukawa=%t polluted=%t verdict=%q", a.Gate256NativeNoGoPreserved, a.ChargeEigenvaluesTreatedAsDerived, a.EmbeddingOrientationTreatedAsSealed, a.ImportedObservedChargeTable, a.ImportedObservedMasses, a.ImportedObservedYukawas, a.ForcedWeakPlane, a.SelectedTrialityByHand, a.SelectedTrialityByDesiredKernel, a.ForcedKernelDim3, a.AcceptedYOnlyAsQ, a.TreatedSealAsFiniteDerivation, a.ConstructedVTauByHand, a.InsertedYukawaTexture, a.PollutedFiniteCore, a.Verdict)
}

func FormatDownstream(a DownstreamAudit) string {
	return fmt.Sprintf("plane=%t full3=%t vtau=%t texture=%t yukawa=%t CKM=%t masses=%t verdict=%q", a.Neutral3PlaneAvailable, a.FullQ8vCKernelDimThree, a.VTauConstructed, a.TrialityTextureOpened, a.YukawaTextureDerived, a.CKMPMNSDerived, a.FermionMassesDerived, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("gate256=%t charges=%t external=%t witnesses=%t so8=%t branches=%t uniqueBranch=%t plane=%t full3=%t yukawa=%t status=%q next=%q comment=%q", a.Gate256SealInherited, a.NativeChargeEigenvaluesExtracted, a.ChargeCoefficientsExternal, a.EmbeddingWitnessesScanned, a.SO8WitnessesTranslated, a.AllTrialityBranchesScanned, a.UniqueTrialityBranchSelected, a.NeutralPolarized3PlaneDerived, a.FullQ8vCKernelDimThree, a.YukawaTextureDerived, a.Status, a.NextGate, a.Comment)
}
