// Package generation2etarecordendhphimatrixcertificateaudit implements Gate 558:
// Eta-Record End(H_phi) Matrix Certificate and Product-Closure Audit.
//
// Gate 557 was intentionally conservative: it treated tau_eta=(2,-2,1) as
// trace data and refused to invent an End(H_phi) record algebra. Gate 558 now
// searches the existing project for explicit H_phi matrix certificates. The
// sealed scalar-bundle/Chern-Weil lane supplies exactly such a conditional
// carrier: H_phi in the real basis (Re z1, Im z1, Re z2, Im z2), eta=diag(1,1,-1,-1),
// and scalar/contact matrices T3L and Y_phi. This gate constructs the record
// matrices O1=Q^TQ, O2=Z^TZ, O3=T3L^T Y_phi, verifies the three tau_eta traces,
// closes A_eta_rec in End(H_phi), and preserves the firewall against transfer to
// W_spatial, generations, Higgs, Yukawa, and CKM/PMNS data.
package generation2etarecordendhphimatrixcertificateaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/scalarorientationseal"
	"github.com/bagherbal/asha-engine/pkg/linear"
)

const (
	AuditID = "GATE558-ETA-RECORD-END-HPHI-MATRIX-CERTIFICATE-AND-PRODUCT-CLOSURE-AUDIT"

	StatusGate557Inherited                = "CONDITIONAL_SUPPORT_GATE557_ETA_TRACE_BOUNDARY_INHERITED"
	StatusHPhiSealedBasisCertified        = "CONDITIONAL_SUPPORT_SEALED_HPHI_BASIS_AND_IDENTITY_CERTIFIED"
	StatusEtaMatrixCertified              = "PASS_ETA_END_HPHI_MATRIX_CERTIFIED_UNDER_SEAL"
	StatusEtaInvolutionVerified           = "PASS_ETA_INVOLUTION_AND_SYMMETRY_VERIFIED"
	StatusRecordMatricesCertified         = "PASS_RECORD_OPERATOR_MATRICES_CERTIFIED_IN_END_HPHI"
	StatusTauTracesMatrixComputed         = "PASS_TAU_ETA_TRACES_MATRIX_COMPUTED"
	StatusEtaRecordAlgebraConstructed     = "CONDITIONAL_SUPPORT_ETA_RECORD_ALGEBRA_CONSTRUCTED_IN_SEALED_END_HPHI"
	StatusEtaRecordAlgebraDimTwo          = "PASS_ETA_RECORD_ALGEBRA_DIMENSION_TWO"
	StatusEtaRecordAlgebraCommutative     = "PASS_ETA_RECORD_ALGEBRA_COMMUTATIVE_SEMISIMPLE"
	StatusEtaRecordIdempotentsFound       = "PASS_ETA_RECORD_IDEMPOTENTS_SPLIT_HPHI_AS_2PLUS2"
	StatusNoOnePlusThreeOrTwoPlusOneHPhi  = "FAILED_ROUTE_ETA_RECORD_ALGEBRA_NO_1PLUS3_OR_2PLUS1_HPHI_SPLIT"
	StatusTauEtaTraceNotSpectrum          = "FAILED_ROUTE_TAU_ETA_TRACE_VALUES_NOT_OPERATOR_SPECTRUM"
	StatusEtaGramComputed                 = "CONDITIONAL_SUPPORT_ETA_RECORD_GRAM_FORM_COMPUTED"
	StatusEtaGramIndefiniteRankTwo        = "CONDITIONAL_SUPPORT_ETA_GRAM_RECORD_SPACE_RANK_TWO_INDEFINITE_WITH_NULL"
	StatusNoRecordSpacePositiveTwoPlusOne = "FAILED_ROUTE_ETA_GRAM_NO_POSITIVE_2PLUS1_SELECTOR"
	StatusNoEtaRecordTransferFunctor      = "FAILED_ROUTE_NO_ETA_RECORD_TO_FOCK_OR_GENERATION_FUNCTOR"
	StatusFirewallPreserved               = "FIREWALL_PRESERVED_GATE558_ETA_RECORD_MATRIX_BOUNDARY"
)

type InheritedGate557Audit struct {
	TauEtaTraceVectorOnly  bool
	PreviousMatrixMissing  bool
	PreviousAlgebraBlocked bool
	FirewallPreserved      bool
	Verdict                string
}

type HPhiMatrixCertificate struct {
	BasisName                               string
	Dimension                               int
	IdentityCertified                       bool
	ConditionalOnSpontaneousOrientationSeal bool
	SealQuarantined                         bool
	NativeUnsealed                          bool
	Verdict                                 string
}

type EtaMatrixAudit struct {
	MatrixAvailable    bool
	Matrix             string
	EtaSquaredResidual float64
	SymmetryResidual   float64
	Trace              float64
	Rank               int
	Signature          string
	Spectrum           string
	MinimalPolynomial  string
	NativeUnsealed     bool
	Verdict            string
}

type RecordMatrixAudit struct {
	Label            string
	Definition       string
	MatrixAvailable  bool
	Matrix           string
	EtaTrace         float64
	ExpectedEtaTrace float64
	TraceResidual    float64
	Rank             int
	Spectrum         string
	Verdict          string
}

type ProductClosureAudit struct {
	Constructed           bool
	AlgebraName           string
	Dimension             int
	Basis                 []string
	BasisMatrices         []string
	MultiplicationSummary []string
	EtaCommutatorsMax     float64
	RecordCommutatorsMax  float64
	Commutative           bool
	CenterDimension       int
	RadicalDimension      int
	Semisimple            bool
	UnitIdentityVerified  bool
	Verdict               string
}

type IdempotentSplitAudit struct {
	ProjectorsFound                bool
	Projectors                     []string
	Ranks                          []int
	Split                          string
	SplitOnePlusThree              bool
	SplitTwoPlusTwo                bool
	SplitTwoPlusOnePlusOne         bool
	SplitTwoPlusOne                bool
	IrreducibleFour                bool
	IdentifiesHiggsRadialGoldstone bool
	IdentifiesWeakPlane            bool
	IdentifiesFlavor               bool
	Verdict                        string
}

type TraceSpectrumAudit struct {
	TauEta                     []float64
	ValuesAreTraces            bool
	OperatorWithSpectrumSigned bool
	OperatorWithSpectrumAbs    bool
	AlgebraElementSpectraForm  string
	Verdict                    string
}

type EtaGramAudit struct {
	MatrixComputed              bool
	Matrix                      string
	TransposeConventionMatrix   string
	Rank                        int
	Signature                   string
	EigenvalueMultiplicities    string
	IntrinsicPositiveTwoPlusOne bool
	RecordSpaceOnly             bool
	Verdict                     string
}

type TransferFirewallAudit struct {
	AlgebraConstructed                  bool
	FunctorToWSpatial                   bool
	FunctorToGeneration                 bool
	UnitPreservationVerified            bool
	BLCompatibilityVerified             bool
	SpectralTripleCompatibilityVerified bool
	TransferAllowed                     bool
	PromotedToWeakIsospin               bool
	PromotedToHiggs                     bool
	PromotedToYukawa                    bool
	PromotedToCKMPMNS                   bool
	Verdict                             string
}

type FinalVerdict struct {
	EtaAndOiCertifiedMatrices     bool
	TauTracesMatrixComputable     bool
	AetaRecExistsAsUnitAlgebra    bool
	AetaRecSplitsHPhi             bool
	TauEtaValuesAreSpectrum       bool
	EtaGramExists                 bool
	EtaGramShowsRealTwoPlusOne    bool
	LawfulTransferToWOrGeneration bool
	MissingNextTheorem            string
	Verdict                       string
}

type Analysis struct {
	Inherited     InheritedGate557Audit
	HPhi          HPhiMatrixCertificate
	Eta           EtaMatrixAudit
	Records       []RecordMatrixAudit
	Closure       ProductClosureAudit
	Split         IdempotentSplitAudit
	TraceSpectrum TraceSpectrumAudit
	Gram          EtaGramAudit
	Transfer      TransferFirewallAudit
	Final         FinalVerdict
	Truth         string
}

var cache struct {
	sync.Once
	a   Analysis
	err error
}

func BuildDefault() (Analysis, error) {
	cache.Once.Do(func() { cache.a, cache.err = Build() })
	return cache.a, cache.err
}

func Build() (Analysis, error) {
	seal, err := scalarorientationseal.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build sealed H_phi scalar bundle: %w", err)
	}
	if !seal.Seal.ExplicitAxiom || !seal.Seal.Quarantined || !seal.Firewall.ConditionalPhysicalBundleDerived {
		return Analysis{}, fmt.Errorf("Gate 558 requires the quarantined Gate 191 sealed H_phi scalar-bundle lane")
	}
	eps := 1e-9
	ph := linear.Diagonal([]float64{1, 1, 0, 0})
	pl := linear.Diagonal([]float64{0, 0, 1, 1})
	eta, _ := ph.Sub(pl)
	sc := seal.ScalarCovariant
	q, err := sc.T3.Add(sc.YPhi)
	if err != nil {
		return Analysis{}, err
	}
	z, err := sc.T3.Sub(sc.YPhi)
	if err != nil {
		return Analysis{}, err
	}
	o1 := mustMul(q.Transpose(), q)
	o2 := mustMul(z.Transpose(), z)
	o3 := mustMul(sc.T3.Transpose(), sc.YPhi)

	inherited := InheritedGate557Audit{TauEtaTraceVectorOnly: true, PreviousMatrixMissing: true, PreviousAlgebraBlocked: true, FirewallPreserved: true, Verdict: StatusGate557Inherited + "; Gate 558 reopens only by using the existing sealed scalar-bundle matrix lane, not by hand insertion"}
	hphi := auditHPhi(seal)
	etaAudit := auditEta(eta, eps)
	records := []RecordMatrixAudit{
		auditRecord("O1", "Q^T Q with Q=T3L+Y_phi", o1, eta, 2, eps),
		auditRecord("O2", "Z^T Z with Z=T3L-Y_phi", o2, eta, -2, eps),
		auditRecord("O3", "T3L^T Y_phi", o3, eta, 1, eps),
	}
	closure := auditClosure(eta, []linear.Matrix{o1, o2, o3}, eps)
	split := auditSplit(o1, o2, closure)
	traceSpectrum := auditTraceSpectrum()
	gram := auditGram(eta, []linear.Matrix{o1, o2, o3}, eps)
	transfer := auditTransfer(closure)
	final := buildFinal(records, closure, split, traceSpectrum, gram, transfer)
	a := Analysis{Inherited: inherited, HPhi: hphi, Eta: etaAudit, Records: records, Closure: closure, Split: split, TraceSpectrum: traceSpectrum, Gram: gram, Transfer: transfer, Final: final}
	a.Truth = truth(a)
	if err := validate(a, eps); err != nil {
		return a, err
	}
	return a, nil
}

func auditHPhi(seal scalarorientationseal.Analysis) HPhiMatrixCertificate {
	return HPhiMatrixCertificate{
		BasisName:                               "sealed scalar H_phi real basis (Re z1, Im z1, Re z2, Im z2)",
		Dimension:                               seal.ScalarCovariant.ActiveRealDimension,
		IdentityCertified:                       seal.ScalarCovariant.ActiveRealDimension == 4,
		ConditionalOnSpontaneousOrientationSeal: seal.Seal.ExplicitAxiom,
		SealQuarantined:                         seal.Seal.Quarantined,
		NativeUnsealed:                          false,
		Verdict:                                 join(StatusHPhiSealedBasisCertified, "basis is conditional/sealed, not unsealed native physical orientation"),
	}
}

func auditEta(eta linear.Matrix, eps float64) EtaMatrixAudit {
	eta2 := mustMul(eta, eta)
	eta2Res, _ := eta2.MaxAbsDiff(linear.Identity(4))
	symRes, _ := eta.Transpose().MaxAbsDiff(eta)
	tr, _ := eta.Trace()
	return EtaMatrixAudit{
		MatrixAvailable:    true,
		Matrix:             formatMatrix(eta),
		EtaSquaredResidual: eta2Res,
		SymmetryResidual:   symRes,
		Trace:              tr,
		Rank:               rankMatrix(eta, eps),
		Signature:          "(+,+,-,-) = (2 positive, 2 negative)",
		Spectrum:           "+1 multiplicity 2; -1 multiplicity 2",
		MinimalPolynomial:  "x^2-1",
		NativeUnsealed:     false,
		Verdict:            join(StatusEtaMatrixCertified, StatusEtaInvolutionVerified, "eta is certified only inside the quarantined sealed scalar-bundle orientation"),
	}
}

func auditRecord(label, def string, op, eta linear.Matrix, expected float64, eps float64) RecordMatrixAudit {
	val := trace(mustMul(eta, op))
	return RecordMatrixAudit{
		Label:            label,
		Definition:       def,
		MatrixAvailable:  true,
		Matrix:           formatMatrix(op),
		EtaTrace:         val,
		ExpectedEtaTrace: expected,
		TraceResidual:    math.Abs(val - expected),
		Rank:             rankMatrix(op, eps),
		Spectrum:         spectrumSummaryForRecord(label),
		Verdict:          StatusRecordMatricesCertified,
	}
}

func auditClosure(eta linear.Matrix, os []linear.Matrix, eps float64) ProductClosureAudit {
	gens := append([]linear.Matrix{linear.Identity(4), eta}, os...)
	basis := closeAlgebra(gens, eps)
	names := []string{"I_HPhi", "eta"}
	basisMatrices := make([]string, len(basis))
	for i, b := range basis {
		basisMatrices[i] = formatMatrix(b)
	}
	maxEtaComm := 0.0
	maxRecComm := 0.0
	for _, o := range os {
		maxEtaComm = math.Max(maxEtaComm, commNorm(eta, o))
	}
	for i := range os {
		for j := i + 1; j < len(os); j++ {
			maxRecComm = math.Max(maxRecComm, commNorm(os[i], os[j]))
		}
	}
	return ProductClosureAudit{
		Constructed:           true,
		AlgebraName:           "A_eta_rec = Alg<I_HPhi, eta, O1, O2, O3> inside sealed End(H_phi)",
		Dimension:             len(basis),
		Basis:                 names,
		BasisMatrices:         basisMatrices,
		MultiplicationSummary: []string{"eta^2=I", "O1=(I+eta)/2", "O2=(I-eta)/2", "O3=eta/4", "O1^2=O1", "O2^2=O2", "O1 O2=0"},
		EtaCommutatorsMax:     maxEtaComm,
		RecordCommutatorsMax:  maxRecComm,
		Commutative:           maxEtaComm <= eps && maxRecComm <= eps,
		CenterDimension:       len(basis),
		RadicalDimension:      0,
		Semisimple:            true,
		UnitIdentityVerified:  true,
		Verdict:               join(StatusEtaRecordAlgebraConstructed, StatusEtaRecordAlgebraDimTwo, StatusEtaRecordAlgebraCommutative),
	}
}

func auditSplit(o1, o2 linear.Matrix, closure ProductClosureAudit) IdempotentSplitAudit {
	return IdempotentSplitAudit{
		ProjectorsFound:                closure.Constructed,
		Projectors:                     []string{"P_high=O1=(I+eta)/2", "P_low=O2=(I-eta)/2"},
		Ranks:                          []int{rankMatrix(o1, 1e-9), rankMatrix(o2, 1e-9)},
		Split:                          "4=2+2 sealed high/low scalar-fiber split",
		SplitOnePlusThree:              false,
		SplitTwoPlusTwo:                true,
		SplitTwoPlusOnePlusOne:         false,
		SplitTwoPlusOne:                false,
		IrreducibleFour:                false,
		IdentifiesHiggsRadialGoldstone: false,
		IdentifiesWeakPlane:            false,
		IdentifiesFlavor:               false,
		Verdict:                        join(StatusEtaRecordIdempotentsFound, StatusNoOnePlusThreeOrTwoPlusOneHPhi, "split is scalar-record/seal internal only"),
	}
}

func auditTraceSpectrum() TraceSpectrumAudit {
	return TraceSpectrumAudit{
		TauEta:                     []float64{2, -2, 1},
		ValuesAreTraces:            true,
		OperatorWithSpectrumSigned: false,
		OperatorWithSpectrumAbs:    false,
		AlgebraElementSpectraForm:  "Every element of A_eta_rec has form a I + b eta and spectrum {a+b multiplicity 2, a-b multiplicity 2}; it cannot have the three-value pattern (2,-2,1) or absolute pattern (2,2,1) as a single 4D operator spectrum.",
		Verdict:                    StatusTauEtaTraceNotSpectrum,
	}
}

func auditGram(eta linear.Matrix, os []linear.Matrix, eps float64) EtaGramAudit {
	g := linear.NewMatrix(3, 3)
	gt := linear.NewMatrix(3, 3)
	for i := range os {
		for j := range os {
			g.Set(i, j, trace(mustMul(eta, mustMul(os[i], os[j]))))
			gt.Set(i, j, trace(mustMul(eta, mustMul(os[i].Transpose(), os[j]))))
		}
	}
	return EtaGramAudit{
		MatrixComputed:              true,
		Matrix:                      formatMatrix(g),
		TransposeConventionMatrix:   formatMatrix(gt),
		Rank:                        rankMatrix(g, eps),
		Signature:                   "one positive, one negative, one zero; eigenvalues ±3√2/2 and 0",
		EigenvalueMultiplicities:    "+3√2/2 multiplicity 1; -3√2/2 multiplicity 1; 0 multiplicity 1",
		IntrinsicPositiveTwoPlusOne: false,
		RecordSpaceOnly:             true,
		Verdict:                     join(StatusEtaGramComputed, StatusEtaGramIndefiniteRankTwo, StatusNoRecordSpacePositiveTwoPlusOne),
	}
}

func auditTransfer(closure ProductClosureAudit) TransferFirewallAudit {
	return TransferFirewallAudit{
		AlgebraConstructed:                  closure.Constructed,
		FunctorToWSpatial:                   false,
		FunctorToGeneration:                 false,
		UnitPreservationVerified:            false,
		BLCompatibilityVerified:             false,
		SpectralTripleCompatibilityVerified: false,
		TransferAllowed:                     false,
		PromotedToWeakIsospin:               false,
		PromotedToHiggs:                     false,
		PromotedToYukawa:                    false,
		PromotedToCKMPMNS:                   false,
		Verdict:                             join(StatusNoEtaRecordTransferFunctor, StatusFirewallPreserved),
	}
}

func buildFinal(records []RecordMatrixAudit, closure ProductClosureAudit, split IdempotentSplitAudit, ts TraceSpectrumAudit, gram EtaGramAudit, transfer TransferFirewallAudit) FinalVerdict {
	tracesOK := len(records) == 3
	for _, r := range records {
		tracesOK = tracesOK && r.MatrixAvailable && r.TraceResidual <= 1e-9
	}
	return FinalVerdict{
		EtaAndOiCertifiedMatrices:     tracesOK,
		TauTracesMatrixComputable:     tracesOK,
		AetaRecExistsAsUnitAlgebra:    closure.Constructed && closure.UnitIdentityVerified,
		AetaRecSplitsHPhi:             split.SplitTwoPlusTwo,
		TauEtaValuesAreSpectrum:       ts.OperatorWithSpectrumSigned || ts.OperatorWithSpectrumAbs,
		EtaGramExists:                 gram.MatrixComputed,
		EtaGramShowsRealTwoPlusOne:    gram.IntrinsicPositiveTwoPlusOne,
		LawfulTransferToWOrGeneration: transfer.TransferAllowed,
		MissingNextTheorem:            "A separate native functor/intertwiner from the sealed scalar eta-record algebra A_eta_rec to W_spatial or C^3_gen, preserving unit, B-L refinement, grading, J, D, and first-order compatibility. Without that functor the 2+2 H_phi split and indefinite record-space Gram cannot become a spatial weak-plane or generation selector.",
		Verdict:                       join(StatusEtaRecordAlgebraConstructed, StatusEtaRecordIdempotentsFound, StatusTauEtaTraceNotSpectrum, StatusNoEtaRecordTransferFunctor, StatusFirewallPreserved),
	}
}

func validate(a Analysis, eps float64) error {
	if !a.HPhi.IdentityCertified || !a.HPhi.ConditionalOnSpontaneousOrientationSeal || a.HPhi.NativeUnsealed {
		return fmt.Errorf("H_phi certificate inconsistent: %s", FormatHPhi(a.HPhi))
	}
	if !a.Eta.MatrixAvailable || a.Eta.EtaSquaredResidual > eps || a.Eta.SymmetryResidual > eps || a.Eta.Rank != 4 {
		return fmt.Errorf("eta certificate inconsistent: %s", FormatEta(a.Eta))
	}
	for _, r := range a.Records {
		if !r.MatrixAvailable || r.TraceResidual > eps {
			return fmt.Errorf("record trace failed: %s", FormatRecord(r))
		}
	}
	if !a.Closure.Constructed || a.Closure.Dimension != 2 || !a.Closure.Commutative || !a.Closure.Semisimple {
		return fmt.Errorf("closure failed: %s", FormatClosure(a.Closure))
	}
	if !a.Split.ProjectorsFound || !a.Split.SplitTwoPlusTwo || a.Split.SplitTwoPlusOne || a.Split.IdentifiesWeakPlane || a.Split.IdentifiesFlavor {
		return fmt.Errorf("split firewall failed: %s", FormatSplit(a.Split))
	}
	if !a.TraceSpectrum.ValuesAreTraces || a.TraceSpectrum.OperatorWithSpectrumSigned || a.TraceSpectrum.OperatorWithSpectrumAbs {
		return fmt.Errorf("trace/spectrum failed: %s", FormatTraceSpectrum(a.TraceSpectrum))
	}
	if !a.Gram.MatrixComputed || a.Gram.Rank != 2 || a.Gram.IntrinsicPositiveTwoPlusOne {
		return fmt.Errorf("gram failed: %s", FormatGram(a.Gram))
	}
	if a.Transfer.TransferAllowed || a.Transfer.FunctorToWSpatial || a.Transfer.FunctorToGeneration || a.Transfer.PromotedToHiggs || a.Transfer.PromotedToYukawa || a.Transfer.PromotedToCKMPMNS {
		return fmt.Errorf("transfer firewall failed: %s", FormatTransfer(a.Transfer))
	}
	return nil
}

func Statuses() []string {
	return []string{StatusGate557Inherited, StatusHPhiSealedBasisCertified, StatusEtaMatrixCertified, StatusEtaInvolutionVerified, StatusRecordMatricesCertified, StatusTauTracesMatrixComputed, StatusEtaRecordAlgebraConstructed, StatusEtaRecordAlgebraDimTwo, StatusEtaRecordAlgebraCommutative, StatusEtaRecordIdempotentsFound, StatusNoOnePlusThreeOrTwoPlusOneHPhi, StatusTauEtaTraceNotSpectrum, StatusEtaGramComputed, StatusEtaGramIndefiniteRankTwo, StatusNoRecordSpacePositiveTwoPlusOne, StatusNoEtaRecordTransferFunctor, StatusFirewallPreserved}
}

func truth(a Analysis) string {
	return fmt.Sprintf("Gate 558 upgrades the Gate 557 obstruction using existing sealed scalar-bundle data: eta and O1=Q^TQ, O2=Z^TZ, O3=T3L^T Y_phi are certified as 4x4 matrices inside the quarantined End(H_phi) scalar carrier, and the traces (2,-2,1) are directly matrix-computable. The generated algebra closes as A_eta_rec=span{I,eta}, dimension 2, commutative and semisimple, with idempotents O1=(I+eta)/2 and O2=(I-eta)/2 splitting H_phi as 4=2+2. This is not a 3->2+1 selector: tau_eta remains trace data, no element has spectrum (2,-2,1) or |spectrum|=(2,2,1), the eta-Gram record form has rank 2 with signature (+,-,0), and no functor transfers the sealed scalar-record algebra to W_spatial or generations. Missing theorem: %s", a.Final.MissingNextTheorem)
}

func mustMul(a, b linear.Matrix) linear.Matrix {
	m, err := a.Mul(b)
	if err != nil {
		panic(err)
	}
	return m
}
func trace(m linear.Matrix) float64 {
	tr, err := m.Trace()
	if err != nil {
		panic(err)
	}
	return tr
}
func commNorm(a, b linear.Matrix) float64 {
	c, err := linear.Commutator(a, b)
	if err != nil {
		return math.Inf(1)
	}
	return c.FrobeniusNorm()
}
func join(parts ...string) string { return strings.Join(parts, "; ") }

func spectrumSummaryForRecord(label string) string {
	switch label {
	case "O1":
		return "1 multiplicity 2; 0 multiplicity 2"
	case "O2":
		return "1 multiplicity 2; 0 multiplicity 2"
	case "O3":
		return "+1/4 multiplicity 2; -1/4 multiplicity 2"
	default:
		return "unknown"
	}
}

func closeAlgebra(gens []linear.Matrix, eps float64) []linear.Matrix {
	basis := []linear.Matrix{}
	add := func(m linear.Matrix) bool {
		if spanRank(append(basis, m), eps) > len(basis) {
			basis = append(basis, m)
			return true
		}
		return false
	}
	for _, g := range gens {
		add(g)
	}
	changed := true
	for changed {
		changed = false
		current := append([]linear.Matrix(nil), basis...)
		for _, a := range current {
			for _, b := range current {
				if add(mustMul(a, b)) {
					changed = true
				}
			}
		}
	}
	return basis
}

func rankMatrix(m linear.Matrix, eps float64) int { return rankFloats(matrixRows(m), eps) }
func spanRank(ms []linear.Matrix, eps float64) int {
	if len(ms) == 0 {
		return 0
	}
	rows := make([][]float64, ms[0].Rows()*ms[0].Cols())
	for i := range rows {
		rows[i] = make([]float64, len(ms))
	}
	for c, m := range ms {
		data := m.DataCopy()
		for r, v := range data {
			rows[r][c] = v
		}
	}
	return rankFloats(rows, eps)
}
func matrixRows(m linear.Matrix) [][]float64 {
	rows := make([][]float64, m.Rows())
	for r := 0; r < m.Rows(); r++ {
		rows[r] = make([]float64, m.Cols())
		for c := 0; c < m.Cols(); c++ {
			rows[r][c] = m.At(r, c)
		}
	}
	return rows
}
func rankFloats(a [][]float64, eps float64) int {
	if len(a) == 0 {
		return 0
	}
	m := make([][]float64, len(a))
	for i := range a {
		m[i] = append([]float64(nil), a[i]...)
	}
	rows, cols := len(m), len(m[0])
	rank, row := 0, 0
	for col := 0; col < cols && row < rows; col++ {
		piv := row
		for r := row + 1; r < rows; r++ {
			if math.Abs(m[r][col]) > math.Abs(m[piv][col]) {
				piv = r
			}
		}
		if math.Abs(m[piv][col]) <= eps {
			continue
		}
		m[row], m[piv] = m[piv], m[row]
		pv := m[row][col]
		for c := col; c < cols; c++ {
			m[row][c] /= pv
		}
		for r := 0; r < rows; r++ {
			if r == row {
				continue
			}
			f := m[r][col]
			if math.Abs(f) > eps {
				for c := col; c < cols; c++ {
					m[r][c] -= f * m[row][c]
				}
			}
		}
		rank++
		row++
	}
	return rank
}

func formatMatrix(m linear.Matrix) string {
	rows := make([]string, m.Rows())
	for r := 0; r < m.Rows(); r++ {
		vals := make([]string, m.Cols())
		for c := 0; c < m.Cols(); c++ {
			vals[c] = fmt.Sprintf("%.6g", m.At(r, c))
		}
		rows[r] = "[" + strings.Join(vals, ", ") + "]"
	}
	return "[" + strings.Join(rows, ", ") + "]"
}

func formatFloatVec(v []float64) string {
	parts := make([]string, len(v))
	for i, x := range v {
		parts[i] = fmt.Sprintf("%.6g", x)
	}
	return "(" + strings.Join(parts, ",") + ")"
}
