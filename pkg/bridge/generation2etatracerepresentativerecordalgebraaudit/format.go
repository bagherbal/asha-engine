package generation2etatracerepresentativerecordalgebraaudit

import (
	"fmt"
	"strings"
)

func FormatIntVec(v []int) string {
	parts := make([]string, len(v))
	for i, x := range v {
		parts[i] = fmt.Sprintf("%d", x)
	}
	return "(" + strings.Join(parts, ",") + ")"
}

func FormatInherited(a InheritedGate556Audit) string {
	return fmt.Sprintf("traceOnly=%t nativeAlgebra=%t unitRep=%t canonicalSpatial=%t firewall=%t verdict=%q", a.TauEtaOnlyTraceVector, a.NativeTauSourceAlgebra, a.UnitPreservingTauRep, a.CanonicalSpatialSelector, a.FirewallPreserved, a.Verdict)
}

func FormatEta(a EtaTypeAudit) string {
	return fmt.Sprintf("symbol=%q traceGrading=%t matrix=%t hermitian=%t eta2=%t traceKnown=%t rankKnown=%t signatureKnown=%t spectrumKnown=%t trace=%q rank=%q signature=%q bookkeepingOnly=%t verdict=%q reason=%q", a.EtaSymbol, a.ActsAsTraceGradingOnHPhi, a.NativeEndHPhiMatrixCertified, a.SymmetricOrHermitianCertified, a.InvolutionEtaSquaredIdentityKnown, a.TraceKnown, a.RankKnown, a.SignatureKnown, a.SpectrumKnown, a.Trace, a.Rank, a.Signature, a.BookkeepingOnly, a.Verdict, a.Reason)
}

func FormatTraceRecord(r TraceRecord) string {
	return fmt.Sprintf("%s:%s value=%d Hphi=%t Wspatial=%t gen=%t matrix=%t products=%t verdict=%q", r.Label, r.OperatorExpression, r.EtaTraceValue, r.LivesOnHPhi, r.LivesOnWSpatial, r.LivesOnGeneration, r.NativeMatrixKnown, r.ProductRowsKnown, r.Verdict)
}

func FormatRecordAlgebra(a EtaRecordAlgebraAudit) string {
	rows := make([]string, len(a.Records))
	for i, r := range a.Records {
		rows[i] = FormatTraceRecord(r)
	}
	return fmt.Sprintf("algebra=%q unit=%q unitRep=%t rho1=%t etaMatrix=%t allMatrices=%t closure=%t dimKnown=%t dim=%d comm=%t etaComm=%t commutative=%t idempotents=%t constructed=%t records=[%s] verdict=%q reason=%q", a.RequestedAlgebra, a.UnitSymbol, a.UnitPreservingRepresentation, a.UnitIdentityVerified, a.EtaMatrixKnown, a.AllRecordMatricesKnown, a.ProductClosureKnown, a.DimensionKnown, a.Dimension, a.CommutatorsKnown, a.EtaCommutatorsKnown, a.CommutativeKnown, a.NontrivialIdempotentsKnown, a.ConstructedAsEndHPhiAlgebra, strings.Join(rows, "; "), a.Verdict, a.Reason)
}

func FormatHPhiSplit(a HPhiSplitAudit) string {
	return fmt.Sprintf("algebra=%t 1+3=%t 2+2=%t 2+1+1=%t irreducible4=%t projectors=%t higgs=%t weak=%t flavor=%t verdict=%q reason=%q", a.AlgebraConstructed, a.SplitOnePlusThree, a.SplitTwoPlusTwo, a.SplitTwoPlusOnePlusOne, a.IrreducibleFourCertified, a.ProjectorsAvailable, a.PhysicalHiggsIdentified, a.WeakPlaneIdentified, a.FlavorIdentified, a.Verdict, a.Reason)
}

func FormatTraceSpectrum(a TraceVsSpectrumAudit) string {
	return fmt.Sprintf("tau=%s abs=%s traces=%t spectrum=%t absSpectrum=%t source=%q verdict=%q reason=%q", FormatIntVec(a.TauEtaValues), FormatIntVec(a.AbsTauEtaValues), a.ValuesAreEtaTraces, a.NativeOperatorWithSpectrum, a.NativeOperatorWithAbsSpectrum, a.SpectrumSource, a.Verdict, a.Reason)
}

func FormatEtaGram(a EtaGramAudit) string {
	return fmt.Sprintf("formula=%q products=%t computed=%t rankKnown=%t rank=%d signatureKnown=%t signature=%q eigMultKnown=%t intrinsic2plus1=%t recordOnly=%t verdict=%q reason=%q", a.RequestedFormula, a.ProductTracesAvailable, a.MatrixComputed, a.RankKnown, a.Rank, a.SignatureKnown, a.Signature, a.EigenvalueMultiplicitiesKnown, a.IntrinsicTwoPlusOneSplit, a.RecordSpaceOnlyIfPresent, a.Verdict, a.Reason)
}

func FormatTransfer(a TransferFunctorAudit) string {
	return fmt.Sprintf("algebra=%t toW=%t toGen=%t unit=%t BL=%t spectralTriple=%t allowed=%t verdict=%q reason=%q", a.RecordAlgebraConstructed, a.FunctorToWSpatialExists, a.FunctorToGenerationCarrierExists, a.UnitPreservationVerified, a.BMinusLCompatibilityVerified, a.SpectralTripleCompatibilityVerified, a.NativeTransferAllowed, a.Verdict, a.Reason)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("Wspatial=%t weakIso=%t generation=%t higgs=%t yukawa=%t ckm=%t diagHand=%t traceAsSpectrum=%t polluted=%t verdict=%q", a.PromotedTauEtaToWSpatial, a.PromotedTauEtaToWeakIsospin, a.PromotedTauEtaToGeneration, a.PromotedTauEtaToHiggs, a.PromotedTauEtaToYukawa, a.PromotedTauEtaToCKMPMNS, a.InsertedDiagonalSelectorByHand, a.PromotedTraceValuesToSpectrum, a.PollutedNativeRegistry, a.Verdict)
}

func FormatFinal(a FinalVerdict) string {
	return fmt.Sprintf("etaNativeOnHPhi=%t recordAlgebra=%t splitsHPhi=%t tauSpectrum=%t record2plus1=%t transfer=%t next=%q verdict=%q", a.EtaNativeOnHPhi, a.EtaRecordAlgebraConstructed, a.EtaRecordAlgebraSplitsHPhi, a.TauEtaValuesAreSpectrum, a.NativeTwoPlusOneAtRecordLevel, a.LawfulTransferToWOrGeneration, a.MissingNextTheorem, a.Verdict)
}
