package generation2etarecordendhphimatrixcertificateaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(a InheritedGate557Audit) string {
	return fmt.Sprintf("traceOnly=%t previousMatrixMissing=%t previousAlgebraBlocked=%t firewall=%t verdict=%q", a.TauEtaTraceVectorOnly, a.PreviousMatrixMissing, a.PreviousAlgebraBlocked, a.FirewallPreserved, a.Verdict)
}

func FormatHPhi(a HPhiMatrixCertificate) string {
	return fmt.Sprintf("basis=%q dim=%d identity=%t conditionalSeal=%t quarantined=%t nativeUnsealed=%t verdict=%q", a.BasisName, a.Dimension, a.IdentityCertified, a.ConditionalOnSpontaneousOrientationSeal, a.SealQuarantined, a.NativeUnsealed, a.Verdict)
}

func FormatEta(a EtaMatrixAudit) string {
	return fmt.Sprintf("matrix=%t eta=%s eta2Residual=%.3g symmetryResidual=%.3g trace=%.3g rank=%d signature=%q spectrum=%q minpoly=%q nativeUnsealed=%t verdict=%q", a.MatrixAvailable, a.Matrix, a.EtaSquaredResidual, a.SymmetryResidual, a.Trace, a.Rank, a.Signature, a.Spectrum, a.MinimalPolynomial, a.NativeUnsealed, a.Verdict)
}

func FormatRecord(a RecordMatrixAudit) string {
	return fmt.Sprintf("%s %s matrix=%t value=%.6g expected=%.6g residual=%.3g rank=%d spectrum=%q M=%s verdict=%q", a.Label, a.Definition, a.MatrixAvailable, a.EtaTrace, a.ExpectedEtaTrace, a.TraceResidual, a.Rank, a.Spectrum, a.Matrix, a.Verdict)
}

func FormatRecords(records []RecordMatrixAudit) string {
	parts := make([]string, len(records))
	for i, r := range records {
		parts[i] = FormatRecord(r)
	}
	return strings.Join(parts, "; ")
}

func FormatClosure(a ProductClosureAudit) string {
	return fmt.Sprintf("constructed=%t algebra=%q dim=%d basis=%v basisMatrices=%v mult=[%s] etaCommMax=%.3g recCommMax=%.3g commutative=%t centerDim=%d radicalDim=%d semisimple=%t unit=%t verdict=%q", a.Constructed, a.AlgebraName, a.Dimension, a.Basis, a.BasisMatrices, strings.Join(a.MultiplicationSummary, "; "), a.EtaCommutatorsMax, a.RecordCommutatorsMax, a.Commutative, a.CenterDimension, a.RadicalDimension, a.Semisimple, a.UnitIdentityVerified, a.Verdict)
}

func FormatSplit(a IdempotentSplitAudit) string {
	return fmt.Sprintf("projectors=%t names=%v ranks=%v split=%q 1+3=%t 2+2=%t 2+1+1=%t 2+1=%t irreducible4=%t higgs=%t weak=%t flavor=%t verdict=%q", a.ProjectorsFound, a.Projectors, a.Ranks, a.Split, a.SplitOnePlusThree, a.SplitTwoPlusTwo, a.SplitTwoPlusOnePlusOne, a.SplitTwoPlusOne, a.IrreducibleFour, a.IdentifiesHiggsRadialGoldstone, a.IdentifiesWeakPlane, a.IdentifiesFlavor, a.Verdict)
}

func FormatTraceSpectrum(a TraceSpectrumAudit) string {
	return fmt.Sprintf("tau=%s traces=%t signedSpectrum=%t absSpectrum=%t elementSpectra=%q verdict=%q", formatFloatVec(a.TauEta), a.ValuesAreTraces, a.OperatorWithSpectrumSigned, a.OperatorWithSpectrumAbs, a.AlgebraElementSpectraForm, a.Verdict)
}

func FormatGram(a EtaGramAudit) string {
	return fmt.Sprintf("computed=%t G=%s GT=%s rank=%d signature=%q eigen=%q positive2plus1=%t recordOnly=%t verdict=%q", a.MatrixComputed, a.Matrix, a.TransposeConventionMatrix, a.Rank, a.Signature, a.EigenvalueMultiplicities, a.IntrinsicPositiveTwoPlusOne, a.RecordSpaceOnly, a.Verdict)
}

func FormatTransfer(a TransferFirewallAudit) string {
	return fmt.Sprintf("algebra=%t toW=%t toGen=%t unit=%t BL=%t spectralTriple=%t allowed=%t weakIso=%t higgs=%t yukawa=%t ckm=%t verdict=%q", a.AlgebraConstructed, a.FunctorToWSpatial, a.FunctorToGeneration, a.UnitPreservationVerified, a.BLCompatibilityVerified, a.SpectralTripleCompatibilityVerified, a.TransferAllowed, a.PromotedToWeakIsospin, a.PromotedToHiggs, a.PromotedToYukawa, a.PromotedToCKMPMNS, a.Verdict)
}

func FormatFinal(a FinalVerdict) string {
	return fmt.Sprintf("matrices=%t traces=%t algebra=%t splits=%t tauSpectrum=%t gram=%t gram2plus1=%t transfer=%t next=%q verdict=%q", a.EtaAndOiCertifiedMatrices, a.TauTracesMatrixComputable, a.AetaRecExistsAsUnitAlgebra, a.AetaRecSplitsHPhi, a.TauEtaValuesAreSpectrum, a.EtaGramExists, a.EtaGramShowsRealTwoPlusOne, a.LawfulTransferToWOrGeneration, a.MissingNextTheorem, a.Verdict)
}
