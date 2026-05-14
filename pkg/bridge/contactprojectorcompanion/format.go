package contactprojectorcompanion

import (
	"fmt"
	"math"
	"strings"
)

func FormatCompanion(a CompanionAudit) string {
	return fmt.Sprintf("poly=%q normalized=%q basis=%q convention=%q C=%s trace=%s det=%s charMatches=%t verdict=%s", a.Polynomial, a.NormalizedPolynomial, a.Basis, a.MatrixConvention, MatrixString(a.Matrix), RatString(a.Trace), RatString(a.Determinant), a.CharacteristicMatches, a.Verdict)
}

func FormatIrreducibility(a IrreducibilityAudit) string {
	return fmt.Sprintf("quarticPrimitive=%t quarticModP=%d quarticMod=%q quarticIrredMod=%t quarticIrredQ=%t resolvent=%q resolventModP=%d resolventMod=%q resolventIrredMod=%t resolventIrredQ=%t verdict=%s", a.QuarticPrimitiveOverZ, a.QuarticModPrime, a.QuarticModPolynomial, a.QuarticIrreducibleOverModP, a.QuarticIrreducibleOverQ, a.ResolventPolynomial, a.ResolventModPrime, a.ResolventModPolynomial, a.ResolventIrreducibleOverModP, a.ResolventIrreducibleOverQ, a.Verdict)
}

func FormatCentralizer(a CentralizerAudit) string {
	return fmt.Sprintf("cyclic=%t dimQ=%d basis={%s} field=%t idempotents=%d trivial={%s} rootProjQ=%t pairProjQ=%t block2x2Q=%t verdict=%s", a.CompanionCyclic, a.CentralizerDimensionOverQ, strings.Join(a.PolynomialBasis, ", "), a.CentralizerIsField, a.NontrivialIdempotentsOverQ, strings.Join(a.TrivialIdempotents, ", "), a.IndividualRootProjectorsOverQ, a.TwoPlusTwoProjectorsOverQ, a.BlockDiagonalizes2x2OverQ, a.Verdict)
}

func FormatNativeAction(c NativeActionCandidate) string {
	residual := "nan"
	if !math.IsNaN(c.Residual) {
		residual = fmt.Sprintf("%.12g", c.Residual)
	}
	return fmt.Sprintf("%s source=%s action=%q acts=%t commutes=%t idem=%t nontrivial=%t selectPair=%t extraMap=%t residual=%s verdict=%s", c.Name, c.SourceGate, c.ProposedAction, c.ActsOnCompanionModule, c.CommutesWithCompanion, c.IsIdempotent, c.IsNontrivial, c.CanSelectRootPair, c.RequiresExternalMap, residual, c.Verdict)
}

func FormatNativeActions(a NativeActionAudit) string {
	parts := make([]string, 0, len(a.Candidates))
	for _, c := range a.Candidates {
		parts = append(parts, FormatNativeAction(c))
	}
	return fmt.Sprintf("legal=%t projector=%t pairSelector=%t candidates={%s} verdict=%s", a.AnyLegalAction, a.AnyCommutingProjector, a.AnyPairSelector, strings.Join(parts, "; "), a.Verdict)
}

func FormatResolvent(a ResolventAdjunctionAudit) string {
	return fmt.Sprintf("pairings={%s} requiresRoot=%t selected=%t wouldSplit=%t nativeAdjunction=%t branches=%d verdict=%s", strings.Join(a.Pairings, ", "), a.PairProjectorRequiresResolventRoot, a.ResolventRootAlreadySelected, a.AdjoiningResolventRootWouldSplit, a.NativeAdjunctionDerived, a.BranchesAfterAdjunction, a.Verdict)
}

func FormatBranch(a BranchAudit) string {
	return fmt.Sprintf("sectorSelected=%t sector=%q projector=%t resolventRoot=%t rootSector=%t rPlus=%.15g rMinus=%.15g rMap=%t selected=%q verdict=%s", a.SectorPairingSelected, a.SectorPairing, a.CompanionProjectorDerived, a.ContactResolventRootSelected, a.RootSectorBijectionDerived, a.RPlus, a.RMinus, a.RBranchMapDerived, a.SelectedBranch, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("ordering=%t masses=%t empiricalY=%t arbitraryResolvent=%t aestheticPair=%t bgapMagnitude=%t higgsNoClaim=%t polluted=%t verdict=%s", a.NoNumericalOrderingPromotion, a.NoObservedMassesUsed, a.NoEmpiricalYukawaInserted, a.NoArbitraryResolventRoot, a.NoAestheticRootPairing, a.NoBGapToRootMagnitudeMap, a.NoHiggsRatioClaimed, a.FiniteCorePolluted, a.Verdict)
}

func FormatFuture(a FutureMap) string {
	parts := make([]string, 0, len(a.Criteria))
	for _, c := range a.Criteria {
		parts = append(parts, fmt.Sprintf("%s[required=%t satisfied=%t detail=%s]", c.Name, c.Required, c.Satisfied, c.Detail))
	}
	return fmt.Sprintf("operator=%t idempotent=%t resolvent=%t rootSector=%t rBranch=%t criteria={%s} next=%s verdict=%s", a.NeedNativeOperatorOnContactModule, a.NeedNontrivialIdempotent, a.NeedResolventRootSelector, a.NeedRootSectorBijection, a.NeedRBranchMap, strings.Join(parts, "; "), a.RecommendedNextGate, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("companion=%t irred=%t projector=%t blockQ=%t resolvent=%t rootSector=%t r=%t higgs=%t firewall=%t next=%s status=%s comment=%q", a.CompanionConstructed, a.IrreducibilityCertified, a.NativeProjectorFound, a.BlockDiagonalizedOverQ, a.ResolventRootSelected, a.RootSectorBijection, a.AmplitudeBranchLocked, a.HiggsRatioDerived, a.FirewallPreserved, a.NextGate, a.Status, a.Comment)
}
