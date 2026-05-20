// Package generation2etarecordtransferranktraceobstructionaudit implements Gate 559:
// Eta-Record Transfer Rank/Trace Obstruction Audit.
//
// Gate 558 certified A_eta_rec=span{I_HPhi,eta} inside the sealed scalar H_phi
// carrier and found the idempotent split H_phi=im(P_+)⊕im(P_-) with ranks 2+2.
// Gate 559 audits whether this algebra can lawfully transfer to a 3-dimensional
// spatial or generation carrier.  The answer is deliberately precise: formal
// unital representations of R⊕R on a 3-dimensional vector space exist and are
// classified by complementary idempotents with rank splits 0+3, 1+2, 2+1, 3+0;
// however ASHA supplies no native, basis-independent functor selecting the 2+1
// case, and trace/rank preservation from the sealed 2+2 H_phi carrier to a 3D
// target is impossible.  The weak-plane, Higgs, generation, Yukawa, CKM/PMNS,
// and flavor firewalls remain closed.
package generation2etarecordtransferranktraceobstructionaudit

import (
	"fmt"
	"strings"
	"sync"

	gate558 "github.com/bagherbal/asha-engine/pkg/bridge/generation2etarecordendhphimatrixcertificateaudit"
)

const (
	AuditID = "GATE559-ETA-RECORD-TRANSFER-RANK-TRACE-OBSTRUCTION-AUDIT"

	StatusGate558Inherited                       = "CONDITIONAL_SUPPORT_GATE558_ETA_RECORD_ALGEBRA_INHERITED"
	StatusFormalUnitalRepsClassified             = "PASS_UNITAL_AETA_REC_REPRESENTATIONS_ON_DIM3_CLASSIFIED"
	StatusFormalDim3RepresentationsExist         = "CONDITIONAL_SUPPORT_FORMAL_AETA_REC_TO_END_C3_REPRESENTATIONS_EXIST"
	StatusNoCanonicalTwoPlusOne                  = "FAILED_ROUTE_ETA_TRANSFER_BASIS_DEPENDENT_NO_CANONICAL_2PLUS1"
	StatusTracePreservingTransferObstructed      = "FAILED_ROUTE_ETA_2PLUS2_TO_SPATIAL3_TRACE_PRESERVING_TRANSFER_OBSTRUCTED"
	StatusNormalizedTraceTransferObstructed      = "FAILED_ROUTE_ETA_NORMALIZED_TRACE_TRANSFER_TO_DIM3_OBSTRUCTED"
	StatusFormalTransferCommutesWithBL           = "CONDITIONAL_SUPPORT_FORMAL_TRANSFER_COMMUTES_WITH_B_MINUS_L"
	StatusBLDoesNotCanonicalizeTransfer          = "FAILED_ROUTE_B_MINUS_L_DOES_NOT_CANONICALIZE_ETA_TRANSFER"
	StatusSpectralTripleCompatibilityUnavailable = "FAILED_ROUTE_ETA_TRANSFER_SPECTRAL_TRIPLE_COMPATIBILITY_UNAVAILABLE_NO_CANONICAL_TRANSFER"
	StatusNoNativeGenerationCarrierFunctor       = "FAILED_ROUTE_NO_NATIVE_GENERATION_CARRIER_FUNCTOR"
	StatusNoLawfulEtaRecordTransfer              = "FAILED_ROUTE_NO_LAWFUL_ETA_RECORD_TRANSFER_TO_W_SPATIAL_OR_GENERATION"
	StatusFirewallPreserved                      = "FIREWALL_PRESERVED_GATE559_ETA_RECORD_TRANSFER_BOUNDARY"
)

type InheritedGate558Audit struct {
	AlgebraConstructed        bool
	AlgebraName               string
	SourceCarrier             string
	SourceDimension           int
	SourcePlusRank            int
	SourceMinusRank           int
	SourceSplit               string
	TauEtaTraceValuesOnly     bool
	NoPreviousTransferFunctor bool
	Verdict                   string
}

type RankSplit struct {
	PlusRank           int
	MinusRank          int
	Name               string
	ProducesTwoPlusOne bool
	CanonicalInASHA    bool
}

type FormalRepresentationClassification struct {
	Algebra                              string
	Carrier                              string
	CarrierDimension                     int
	UnitalRepresentationsExist           bool
	EquivalentToComplementaryIdempotents bool
	RankSplits                           []RankSplit
	Exhaustive                           bool
	AnyTwoPlusOneFormal                  bool
	AnyCanonicalTwoPlusOne               bool
	Verdict                              string
}

type CanonicalChoiceAudit struct {
	WSpatialCarrierAvailable              bool
	WSpatialBasisName                     string
	GenerationCarrierCapacityVisible      bool
	NativeBasisIndependentReasonFor2Plus1 bool
	NativeReasonForU12                    bool
	NativeReasonForGenerationLabels       bool
	Verdict                               string
}

type TraceRankPreservationAudit struct {
	SourceCarrier                   string
	SourceRanks                     []int
	TargetCarrierDimension          int
	OrdinaryTracePreservingPossible bool
	RequiredTargetRanks             []int
	Obstruction                     string
	Verdict                         string
}

type NormalizedTracePreservationAudit struct {
	SourceNormalizedTraces []float64
	TargetDimension        int
	RequiredTargetRanks    []float64
	IntegralRanksPossible  bool
	Obstruction            string
	Verdict                string
}

type BLCompatibilityAudit struct {
	RestrictedBLOnWSpatial          string
	AnyFormalTransferCommutesWithBL bool
	BLSuppliesRankSplit             bool
	BLSuppliesBasisLabels           bool
	BLSuppliesCanonicalU12          bool
	Verdict                         string
}

type SpectralTripleCompatibilityAudit struct {
	CandidateTransferExists  bool
	GradingCheckAvailable    bool
	JCheckAvailable          bool
	DCheckAvailable          bool
	FirstOrderCheckAvailable bool
	CompatibilityPassed      bool
	MissingData              []string
	Verdict                  string
}

type GenerationCarrierAudit struct {
	FormalDim3GenerationCapacityVisible bool
	NativeBasisIndependentLabels        bool
	FunctorFromAetaRec                  bool
	UnitPreservationVerified            bool
	ProducesGenerationHierarchy         bool
	ProducesYukawaOrCKMPMNS             bool
	Verdict                             string
}

type FirewallAudit struct {
	WeakPlaneSelectionClaimed        bool
	WeakIsospinIdentificationClaimed bool
	HiggsRadialGoldstoneClaimed      bool
	GenerationHierarchyClaimed       bool
	YukawaTextureClaimed             bool
	CKMPMNSClaimed                   bool
	ObservedFlavorImported           bool
	Preserved                        bool
	Verdict                          string
}

type FinalVerdict struct {
	FormalRepresentationsExist   bool
	CanonicalInASHA              bool
	TraceRankPreservingTransfer  bool
	BMinusLCanonicalizesTransfer bool
	LawfulTransferAvailable      bool
	MissingNextTheorem           string
	Verdict                      string
}

type Analysis struct {
	Inherited       InheritedGate558Audit
	FormalReps      FormalRepresentationClassification
	CanonicalChoice CanonicalChoiceAudit
	TraceRank       TraceRankPreservationAudit
	NormalizedTrace NormalizedTracePreservationAudit
	BL              BLCompatibilityAudit
	SpectralTriple  SpectralTripleCompatibilityAudit
	Generation      GenerationCarrierAudit
	Firewall        FirewallAudit
	Final           FinalVerdict
	Truth           string
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
	prev, err := gate558.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate 558 eta-record matrix certificate: %w", err)
	}
	inherited := auditInherited(prev)
	formal := auditFormalRepresentations()
	canonical := auditCanonicalChoice()
	traceRank := auditTraceRank(inherited)
	normalized := auditNormalizedTrace()
	bl := auditBLCompatibility()
	spectral := auditSpectralTriple()
	generation := auditGenerationCarrier()
	firewall := auditFirewall()
	final := auditFinal(formal, canonical, traceRank, normalized, bl, generation)
	a := Analysis{Inherited: inherited, FormalReps: formal, CanonicalChoice: canonical, TraceRank: traceRank, NormalizedTrace: normalized, BL: bl, SpectralTriple: spectral, Generation: generation, Firewall: firewall, Final: final}
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func auditInherited(prev gate558.Analysis) InheritedGate558Audit {
	return InheritedGate558Audit{
		AlgebraConstructed:        prev.Closure.Constructed,
		AlgebraName:               prev.Closure.AlgebraName,
		SourceCarrier:             "sealed H_phi scalar-record carrier",
		SourceDimension:           4,
		SourcePlusRank:            2,
		SourceMinusRank:           2,
		SourceSplit:               prev.Split.Split,
		TauEtaTraceValuesOnly:     prev.TraceSpectrum.ValuesAreTraces && !prev.TraceSpectrum.OperatorWithSpectrumSigned && !prev.TraceSpectrum.OperatorWithSpectrumAbs,
		NoPreviousTransferFunctor: !prev.Transfer.TransferAllowed && !prev.Transfer.FunctorToWSpatial && !prev.Transfer.FunctorToGeneration,
		Verdict:                   join(StatusGate558Inherited, "A_eta_rec=span{I,eta} is available only in sealed End(H_phi) with 2+2 idempotent ranks"),
	}
}

func auditFormalRepresentations() FormalRepresentationClassification {
	splits := []RankSplit{
		{PlusRank: 0, MinusRank: 3, Name: "3=0+3", ProducesTwoPlusOne: false, CanonicalInASHA: false},
		{PlusRank: 1, MinusRank: 2, Name: "3=1+2", ProducesTwoPlusOne: true, CanonicalInASHA: false},
		{PlusRank: 2, MinusRank: 1, Name: "3=2+1", ProducesTwoPlusOne: true, CanonicalInASHA: false},
		{PlusRank: 3, MinusRank: 0, Name: "3=3+0", ProducesTwoPlusOne: false, CanonicalInASHA: false},
	}
	return FormalRepresentationClassification{
		Algebra:                              "A_eta_rec ≅ R⊕R generated by complementary idempotents P_+, P_-",
		Carrier:                              "abstract 3-dimensional vector space V",
		CarrierDimension:                     3,
		UnitalRepresentationsExist:           true,
		EquivalentToComplementaryIdempotents: true,
		RankSplits:                           splits,
		Exhaustive:                           true,
		AnyTwoPlusOneFormal:                  true,
		AnyCanonicalTwoPlusOne:               false,
		Verdict:                              join(StatusFormalUnitalRepsClassified, StatusFormalDim3RepresentationsExist, StatusNoCanonicalTwoPlusOne),
	}
}

func auditCanonicalChoice() CanonicalChoiceAudit {
	return CanonicalChoiceAudit{
		WSpatialCarrierAvailable:              true,
		WSpatialBasisName:                     "W_spatial=span_C{a_1†,a_2†,a_3†}",
		GenerationCarrierCapacityVisible:      true,
		NativeBasisIndependentReasonFor2Plus1: false,
		NativeReasonForU12:                    false,
		NativeReasonForGenerationLabels:       false,
		Verdict:                               join(StatusNoCanonicalTwoPlusOne, "the 2+1 rank choice requires selecting a 2-plane/1-line in a 3D target; ASHA has no current functor that labels which plane is intrinsic"),
	}
}

func auditTraceRank(inherited InheritedGate558Audit) TraceRankPreservationAudit {
	return TraceRankPreservationAudit{
		SourceCarrier:                   inherited.SourceCarrier,
		SourceRanks:                     []int{inherited.SourcePlusRank, inherited.SourceMinusRank},
		TargetCarrierDimension:          3,
		OrdinaryTracePreservingPossible: false,
		RequiredTargetRanks:             []int{2, 2},
		Obstruction:                     "a unital target representation requires rank(rho(P_+))+rank(rho(P_-))=3, but preserving source ranks would require 2+2=4",
		Verdict:                         StatusTracePreservingTransferObstructed,
	}
}

func auditNormalizedTrace() NormalizedTracePreservationAudit {
	return NormalizedTracePreservationAudit{
		SourceNormalizedTraces: []float64{0.5, 0.5},
		TargetDimension:        3,
		RequiredTargetRanks:    []float64{1.5, 1.5},
		IntegralRanksPossible:  false,
		Obstruction:            "normalized trace preservation requires rank(rho(P_+))/3=rank(rho(P_-))/3=1/2, hence ranks 3/2 and 3/2, impossible for idempotents",
		Verdict:                StatusNormalizedTraceTransferObstructed,
	}
}

func auditBLCompatibility() BLCompatibilityAudit {
	return BLCompatibilityAudit{
		RestrictedBLOnWSpatial:          "B-L|W_spatial=(1/3)I_3",
		AnyFormalTransferCommutesWithBL: true,
		BLSuppliesRankSplit:             false,
		BLSuppliesBasisLabels:           false,
		BLSuppliesCanonicalU12:          false,
		Verdict:                         join(StatusFormalTransferCommutesWithBL, StatusBLDoesNotCanonicalizeTransfer),
	}
}

func auditSpectralTriple() SpectralTripleCompatibilityAudit {
	return SpectralTripleCompatibilityAudit{
		CandidateTransferExists:  false,
		GradingCheckAvailable:    false,
		JCheckAvailable:          false,
		DCheckAvailable:          false,
		FirstOrderCheckAvailable: false,
		CompatibilityPassed:      false,
		MissingData:              []string{"canonical rho_eta_rec:A_eta_rec->End(W_spatial) or End(C^3_gen)", "basis-independent image projectors", "gamma action on target", "J action on target", "D/finite Dirac action on target", "first-order test rows for the transferred representation"},
		Verdict:                  StatusSpectralTripleCompatibilityUnavailable,
	}
}

func auditGenerationCarrier() GenerationCarrierAudit {
	return GenerationCarrierAudit{
		FormalDim3GenerationCapacityVisible: true,
		NativeBasisIndependentLabels:        false,
		FunctorFromAetaRec:                  false,
		UnitPreservationVerified:            false,
		ProducesGenerationHierarchy:         false,
		ProducesYukawaOrCKMPMNS:             false,
		Verdict:                             StatusNoNativeGenerationCarrierFunctor,
	}
}

func auditFirewall() FirewallAudit {
	return FirewallAudit{
		WeakPlaneSelectionClaimed:        false,
		WeakIsospinIdentificationClaimed: false,
		HiggsRadialGoldstoneClaimed:      false,
		GenerationHierarchyClaimed:       false,
		YukawaTextureClaimed:             false,
		CKMPMNSClaimed:                   false,
		ObservedFlavorImported:           false,
		Preserved:                        true,
		Verdict:                          StatusFirewallPreserved,
	}
}

func auditFinal(formal FormalRepresentationClassification, canonical CanonicalChoiceAudit, tr TraceRankPreservationAudit, nt NormalizedTracePreservationAudit, bl BLCompatibilityAudit, gen GenerationCarrierAudit) FinalVerdict {
	lawful := canonical.NativeBasisIndependentReasonFor2Plus1 && tr.OrdinaryTracePreservingPossible && nt.IntegralRanksPossible && gen.FunctorFromAetaRec
	return FinalVerdict{
		FormalRepresentationsExist:   formal.UnitalRepresentationsExist,
		CanonicalInASHA:              canonical.NativeBasisIndependentReasonFor2Plus1 || formal.AnyCanonicalTwoPlusOne,
		TraceRankPreservingTransfer:  tr.OrdinaryTracePreservingPossible || nt.IntegralRanksPossible,
		BMinusLCanonicalizesTransfer: bl.BLSuppliesRankSplit || bl.BLSuppliesBasisLabels || bl.BLSuppliesCanonicalU12,
		LawfulTransferAvailable:      lawful,
		MissingNextTheorem:           "A native, basis-independent functor/intertwiner F:A_eta_rec->End(W_spatial) or End(C^3_gen), with F(1)=I, a canonical rank split, non-arbitrary target labels, B-L refinement data, and gamma/J/D/first-order compatibility. If trace preservation is required, the current sealed 2+2 source-to-dim3 target route is rank/trace obstructed.",
		Verdict:                      join(StatusFormalUnitalRepsClassified, StatusNoCanonicalTwoPlusOne, StatusTracePreservingTransferObstructed, StatusNormalizedTraceTransferObstructed, StatusBLDoesNotCanonicalizeTransfer, StatusNoLawfulEtaRecordTransfer, StatusFirewallPreserved),
	}
}

func validate(a Analysis) error {
	if !a.Inherited.AlgebraConstructed || a.Inherited.SourcePlusRank != 2 || a.Inherited.SourceMinusRank != 2 || !a.Inherited.TauEtaTraceValuesOnly || !a.Inherited.NoPreviousTransferFunctor {
		return fmt.Errorf("inherited Gate 558 boundary inconsistent: %s", FormatInherited(a.Inherited))
	}
	if !a.FormalReps.UnitalRepresentationsExist || !a.FormalReps.EquivalentToComplementaryIdempotents || !a.FormalReps.Exhaustive || len(a.FormalReps.RankSplits) != 4 || !a.FormalReps.AnyTwoPlusOneFormal || a.FormalReps.AnyCanonicalTwoPlusOne {
		return fmt.Errorf("formal representation classification failed: %s", FormatFormalReps(a.FormalReps))
	}
	if a.CanonicalChoice.NativeBasisIndependentReasonFor2Plus1 || a.CanonicalChoice.NativeReasonForU12 || a.CanonicalChoice.NativeReasonForGenerationLabels {
		return fmt.Errorf("canonical choice firewall failed: %s", FormatCanonicalChoice(a.CanonicalChoice))
	}
	if a.TraceRank.OrdinaryTracePreservingPossible || a.NormalizedTrace.IntegralRanksPossible {
		return fmt.Errorf("trace/rank obstruction failed: %s | %s", FormatTraceRank(a.TraceRank), FormatNormalizedTrace(a.NormalizedTrace))
	}
	if !a.BL.AnyFormalTransferCommutesWithBL || a.BL.BLSuppliesRankSplit || a.BL.BLSuppliesBasisLabels || a.BL.BLSuppliesCanonicalU12 {
		return fmt.Errorf("B-L compatibility/canonicalization audit failed: %s", FormatBL(a.BL))
	}
	if a.SpectralTriple.CandidateTransferExists || a.SpectralTriple.CompatibilityPassed || a.SpectralTriple.GradingCheckAvailable || a.SpectralTriple.JCheckAvailable || a.SpectralTriple.DCheckAvailable || a.SpectralTriple.FirstOrderCheckAvailable {
		return fmt.Errorf("spectral triple compatibility should be unavailable without transfer: %s", FormatSpectralTriple(a.SpectralTriple))
	}
	if a.Generation.FunctorFromAetaRec || a.Generation.NativeBasisIndependentLabels || a.Generation.ProducesGenerationHierarchy || a.Generation.ProducesYukawaOrCKMPMNS {
		return fmt.Errorf("generation firewall failed: %s", FormatGeneration(a.Generation))
	}
	if !a.Firewall.Preserved || a.Firewall.WeakPlaneSelectionClaimed || a.Firewall.WeakIsospinIdentificationClaimed || a.Firewall.HiggsRadialGoldstoneClaimed || a.Firewall.GenerationHierarchyClaimed || a.Firewall.YukawaTextureClaimed || a.Firewall.CKMPMNSClaimed || a.Firewall.ObservedFlavorImported {
		return fmt.Errorf("physical firewall failed: %s", FormatFirewall(a.Firewall))
	}
	if !a.Final.FormalRepresentationsExist || a.Final.CanonicalInASHA || a.Final.TraceRankPreservingTransfer || a.Final.BMinusLCanonicalizesTransfer || a.Final.LawfulTransferAvailable {
		return fmt.Errorf("final verdict inconsistent: %s", FormatFinal(a.Final))
	}
	return nil
}

func Statuses() []string {
	return []string{
		StatusGate558Inherited,
		StatusFormalUnitalRepsClassified,
		StatusFormalDim3RepresentationsExist,
		StatusNoCanonicalTwoPlusOne,
		StatusTracePreservingTransferObstructed,
		StatusNormalizedTraceTransferObstructed,
		StatusFormalTransferCommutesWithBL,
		StatusBLDoesNotCanonicalizeTransfer,
		StatusSpectralTripleCompatibilityUnavailable,
		StatusNoNativeGenerationCarrierFunctor,
		StatusNoLawfulEtaRecordTransfer,
		StatusFirewallPreserved,
	}
}

func truth(a Analysis) string {
	return fmt.Sprintf("Gate 559 proves that A_eta_rec has formal unital 3D representations with rank splits %s, including formal 2+1 choices. But the sealed H_phi source split is 2+2, trace and normalized-trace preservation to dimension 3 are impossible, B-L is scalar on W_spatial and therefore cannot choose a plane, and no native functor to W_spatial or a generation carrier is present. The eta-record transfer remains arbitrary/basis-dependent and firewalled.", FormatRankSplitList(a.FormalReps.RankSplits))
}

func join(parts ...string) string {
	clean := make([]string, 0, len(parts))
	for _, p := range parts {
		if strings.TrimSpace(p) != "" {
			clean = append(clean, p)
		}
	}
	return strings.Join(clean, "; ")
}
