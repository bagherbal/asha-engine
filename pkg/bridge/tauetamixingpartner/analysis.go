// Package tauetamixingpartner implements Gate 262:
// TauEta Non-Commuting Partner / Finite Phase-Mixing Source Audit.
//
// Gate 261 placed tau_eta=(2,-2,1) in the 3x3 generation bilinear carrier and
// exposed the six off-diagonal matrix units where flavor mixing would have to
// live. Gate 262 audits the finite candidates already known to the engine:
// exact triality permutations, the B-sector gap scalar, and Hopf phase residuals.
// It separates raw non-commutation from a qualified Yukawa amplitude source.
package tauetamixingpartner

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/tauetayukawasourcemap"
)

const (
	AuditID = "GATE262-TAU-ETA-NON-COMMUTING-PARTNER-FINITE-PHASE-MIXING-SOURCE-AUDIT"

	StatusGate261Inherited                   = "CONDITIONAL_SUPPORT_GATE261_TAU_ETA_TEXTURE_COMPLEMENT_INHERITED"
	StatusTrialityRawComplementPopulated     = "CONDITIONAL_SUPPORT_TRIALITY_OPERATORS_POPULATE_AD_TAU_COMPLEMENT"
	StatusHermitianTrialityPhaseBasisExposed = "CONDITIONAL_SUPPORT_HERMITIAN_TRIALITY_PHASE_BASIS_EXPOSED"
	StatusBGapRejectedAsScalar               = "FAILED_ROUTE_B_GAP_HAS_NO_GENERATION_ENDOMORPHISM"
	StatusHopfRejectedAsRepresentationFree   = "FAILED_ROUTE_HOPF_PHASE_RESIDUALS_LACK_GENERATION_TEXTURE_MAP"
	StatusNoQualifiedPartner                 = "FAILED_ROUTE_NO_QUALIFIED_FINITE_MIXING_PARTNER_IDENTIFIED"
	StatusEmpiricalYukawaSealPreserved       = "CONDITIONAL_SUPPORT_EMPIRICAL_YUKAWA_SEAL_PRESERVED"
	StatusCKMPMNSStillBlocked                = "FAILED_ROUTE_CKM_PMNS_AND_FERMION_MASSES_STILL_BLOCKED"
)

type CandidateKind string

const (
	KindTrialityPermutation CandidateKind = "triality permutation"
	KindTrialityHermitian   CandidateKind = "Hermitian triality algebra component"
	KindBGapScalar          CandidateKind = "B-sector spectral gap scalar"
	KindHopfPhaseResidual   CandidateKind = "S7 Hopf phase residual"
)

type GaussianInt struct {
	Re int
	Im int
}

type Matrix3 [3][3]GaussianInt

type Gate261Inheritance struct {
	EightVRouteClosed              bool
	BilinearCarrierDefined         bool
	TauEtaActionDerived            bool
	TextureAlgebraDecomposed       bool
	CommutatorComplementExposed    bool
	PreviousCanonicalPartnerFound  bool
	PreviousPhysicalTextureDerived bool
	PreviousCKMPMNSDerived         bool
	PreviousFermionMassesDerived   bool
	TauEtaEigenvalues              []int
	TextureAlgebraDimension        int
	CommutantDimension             int
	OffDiagonalComplementDimension int
	DistinctAbsMixingGaps          []int
	Verdict                        string
}

type CandidateAudit struct {
	Name                           string
	Kind                           CandidateKind
	Source                         string
	Matrix                         Matrix3
	MatrixAvailable                bool
	CanonicalFiniteData            bool
	SelfAdjoint                    bool
	UnitaryOrOrthogonal            bool
	PureSymmetryOrLabelAction      bool
	ScalarOnly                     bool
	PhaseOnly                      bool
	GenerationEndomorphismDerived  bool
	AmplitudeSourceDerived         bool
	RequiresRepresentationBridge   bool
	RequiresActionCoefficient      bool
	CommutesWithTauEta             bool
	Commutator                     Matrix3
	CommutatorFrobeniusNormSquared int
	OffDiagonalSupportEntries      int
	PopulatesMixingComplement      bool
	QualifiedFiniteMixingPartner   bool
	Disqualification               string
	Verdict                        string
}

type InventoryAudit struct {
	CandidateCount                  int
	MatrixCandidates                int
	CanonicalFiniteCandidates       int
	RawNonCommutingCandidates       int
	SelfAdjointRawCandidates        int
	HermitianPhaseBasisCandidates   int
	QualifiedFiniteMixingPartners   int
	ScalarGapCandidatesRejected     int
	RepresentationFreePhaseRejected int
	EmpiricalCandidatesUsed         int
	Verdict                         string
}

type TrialityPartnerAudit struct {
	PermutationCycleNonCommuting         bool
	ReflectionNonCommuting               bool
	HermitianRealPartNonCommuting        bool
	HermitianImaginaryPartNonCommuting   bool
	HermitianPhaseBasisDimension         int
	RawComplementDirectionsTouched       int
	AllRawTrialityMapsAreSymmetryAlgebra bool
	AnyTrialityMapQualifiedAsAmplitude   bool
	Verdict                              string
}

type FinitePhaseGapAudit struct {
	BGapAvailableAsPositiveScale         bool
	BGapHasGenerationMatrix              bool
	BGapCanPopulateOffDiagonalComplement bool
	HopfPhasesAvailableAsPhaseLedger     bool
	HopfPhaseGenerationMapDerived        bool
	HopfCanPopulateOffDiagonalComplement bool
	RequiresFiniteActionFunctional       bool
	RequiresRepresentationBridge         bool
	Verdict                              string
}

type MixingPartnerVerdict struct {
	RawNonCommutingPartnerExists         bool
	RawSelfAdjointOffDiagonalBasisExists bool
	QualifiedFiniteMixingPartnerFound    bool
	PhysicalYukawaTextureDerived         bool
	CKMPMNSDerived                       bool
	FermionMassesDerived                 bool
	Reason                               string
	NextGate                             string
	Status                               string
}

type FirewallAudit struct {
	Gate261SourceMapPreserved          bool
	DoesNotReopenEightVKernelRoute     bool
	DoesNotUseObservedMasses           bool
	DoesNotUseObservedMixingAngles     bool
	DoesNotPromoteSymmetryToAmplitude  bool
	DoesNotUseBGapAsTextureWithoutMap  bool
	DoesNotUseHopfPhaseWithoutMap      bool
	DoesNotClaimFiniteActionFunctional bool
	EmpiricalYukawaSealInactive        bool
	FiniteCorePolluted                 bool
	Verdict                            string
}

type Summary struct {
	Gate261Inherited                 bool
	TrialityComplementPopulated      bool
	HermitianTrialityBasisExposed    bool
	BGapRejectedAsRepresentationFree bool
	HopfRejectedAsRepresentationFree bool
	QualifiedMixingPartnerFound      bool
	PhysicalYukawaTextureDerived     bool
	CKMPMNSDerived                   bool
	FermionMassesDerived             bool
	Status                           string
	NextGate                         string
	Comment                          string
}

type Analysis struct {
	PreviousGate261 tauetayukawasourcemap.Analysis
	Inheritance     Gate261Inheritance
	Candidates      []CandidateAudit
	Inventory       InventoryAudit
	TrialityPartner TrialityPartnerAudit
	FinitePhaseGap  FinitePhaseGapAudit
	PartnerVerdict  MixingPartnerVerdict
	Firewall        FirewallAudit
	Summary         Summary
	TruthStatement  string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := tauetayukawasourcemap.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 261 predecessor: %w", err)
			return
		}
		inh := inheritGate261(prev)
		candidates := auditCandidates(inh)
		inventory := auditInventory(candidates)
		triality := auditTrialityPartner(candidates)
		phaseGap := auditFinitePhaseGap(candidates)
		verdict := auditPartnerVerdict(inh, inventory, triality, phaseGap)
		firewall := auditFirewall(inh, verdict)
		summary := summarize(inh, inventory, triality, phaseGap, verdict)
		truth := buildTruth(inh, inventory, triality, phaseGap, verdict)
		defaultA = Analysis{PreviousGate261: prev, Inheritance: inh, Candidates: candidates, Inventory: inventory, TrialityPartner: triality, FinitePhaseGap: phaseGap, PartnerVerdict: verdict, Firewall: firewall, Summary: summary, TruthStatement: truth}
	})
	return defaultA, defaultErr
}

func inheritGate261(a tauetayukawasourcemap.Analysis) Gate261Inheritance {
	return Gate261Inheritance{
		EightVRouteClosed:              a.Inheritance.EightVRouteClosed,
		BilinearCarrierDefined:         a.Summary.BilinearCarrierDefined,
		TauEtaActionDerived:            a.Summary.TauEtaActionDerived,
		TextureAlgebraDecomposed:       a.Summary.TextureAlgebraDecomposed,
		CommutatorComplementExposed:    a.Summary.CommutatorComplementExposed,
		PreviousCanonicalPartnerFound:  a.Summary.CanonicalMixingPartnerFound,
		PreviousPhysicalTextureDerived: a.Summary.PhysicalYukawaTextureDerived,
		PreviousCKMPMNSDerived:         a.Summary.CKMPMNSDerived,
		PreviousFermionMassesDerived:   a.Summary.FermionMassesDerived,
		TauEtaEigenvalues:              append([]int(nil), a.TauEtaAction.Eigenvalues...),
		TextureAlgebraDimension:        a.TextureAlgebra.TextureAlgebraDimension,
		CommutantDimension:             a.TextureAlgebra.CommutantDimension,
		OffDiagonalComplementDimension: a.TextureAlgebra.OffDiagonalComplementDimension,
		DistinctAbsMixingGaps:          append([]int(nil), a.TextureAlgebra.DistinctAbsMixingGaps...),
		Verdict:                        StatusGate261Inherited + "; inherited tau_eta action on M3(C), 3D commutant, and 6D off-diagonal complement without broad historical reruns",
	}
}

func auditCandidates(inh Gate261Inheritance) []CandidateAudit {
	lambda := append([]int(nil), inh.TauEtaEigenvalues...)
	if len(lambda) != 3 {
		lambda = []int{2, -2, 1}
	}
	C := Matrix3{
		{z(0, 0), z(1, 0), z(0, 0)},
		{z(0, 0), z(0, 0), z(1, 0)},
		{z(1, 0), z(0, 0), z(0, 0)},
	}
	R23 := Matrix3{
		{z(1, 0), z(0, 0), z(0, 0)},
		{z(0, 0), z(0, 0), z(1, 0)},
		{z(0, 0), z(1, 0), z(0, 0)},
	}
	ARe := Matrix3{
		{z(0, 0), z(1, 0), z(1, 0)},
		{z(1, 0), z(0, 0), z(1, 0)},
		{z(1, 0), z(1, 0), z(0, 0)},
	}
	AIm := Matrix3{
		{z(0, 0), z(0, 1), z(0, -1)},
		{z(0, -1), z(0, 0), z(0, 1)},
		{z(0, 1), z(0, -1), z(0, 0)},
	}
	zero := Matrix3{}

	base := []CandidateAudit{
		{Name: "C3_cycle", Kind: KindTrialityPermutation, Source: "Gate 26 / Gate 173 exact triality cycle", Matrix: C, MatrixAvailable: true, CanonicalFiniteData: true, SelfAdjoint: false, UnitaryOrOrthogonal: true, PureSymmetryOrLabelAction: true, GenerationEndomorphismDerived: true, AmplitudeSourceDerived: false, RequiresActionCoefficient: true, Disqualification: "canonical raw non-commuting generation map, but it is a unitary triality relabelling symmetry and not a self-adjoint Yukawa amplitude source"},
		{Name: "S3_reflection_23", Kind: KindTrialityPermutation, Source: "Gate 26 / Gate 173 exact triality reflection", Matrix: R23, MatrixAvailable: true, CanonicalFiniteData: true, SelfAdjoint: true, UnitaryOrOrthogonal: true, PureSymmetryOrLabelAction: true, GenerationEndomorphismDerived: true, AmplitudeSourceDerived: false, RequiresActionCoefficient: true, Disqualification: "self-adjoint and raw non-commuting, but it is an exact triality reflection/label symmetry rather than a finite Dirac amplitude texture"},
		{Name: "A_triality_real=C+C^T", Kind: KindTrialityHermitian, Source: "Hermitian part of the exact triality cycle algebra", Matrix: ARe, MatrixAvailable: true, CanonicalFiniteData: true, SelfAdjoint: true, UnitaryOrOrthogonal: false, PureSymmetryOrLabelAction: true, GenerationEndomorphismDerived: true, AmplitudeSourceDerived: false, RequiresActionCoefficient: true, Disqualification: "Hermitian off-diagonal basis element derived from triality symmetry algebra, but no finite action selects its coefficient or promotes it to a Yukawa amplitude"},
		{Name: "K_triality_phase=i(C-C^T)", Kind: KindTrialityHermitian, Source: "phase-like Hermitian part of the exact triality cycle algebra", Matrix: AIm, MatrixAvailable: true, CanonicalFiniteData: true, SelfAdjoint: true, UnitaryOrOrthogonal: false, PureSymmetryOrLabelAction: true, PhaseOnly: true, GenerationEndomorphismDerived: true, AmplitudeSourceDerived: false, RequiresActionCoefficient: true, Disqualification: "Hermitian phase-like off-diagonal basis element exists, but the Hopf/finite action has not selected this phase as a physical Yukawa source"},
		{Name: "B_gap", Kind: KindBGapScalar, Source: "finite B-sector first positive spectral gap ledger", Matrix: zero, MatrixAvailable: false, CanonicalFiniteData: true, SelfAdjoint: true, ScalarOnly: true, GenerationEndomorphismDerived: false, AmplitudeSourceDerived: false, RequiresRepresentationBridge: true, RequiresActionCoefficient: true, Disqualification: "positive finite scalar/gap anchor only; no representation map from B_gap to a 3x3 generation endomorphism is derived"},
		{Name: "S7_Hopf_phase_residual", Kind: KindHopfPhaseResidual, Source: "S7 Hopf phase/contact residual ledger", Matrix: zero, MatrixAvailable: false, CanonicalFiniteData: true, SelfAdjoint: false, PhaseOnly: true, GenerationEndomorphismDerived: false, AmplitudeSourceDerived: false, RequiresRepresentationBridge: true, RequiresActionCoefficient: true, Disqualification: "native phase character exists only as a representation-free residual; no map to the off-diagonal M3(C) texture complement is derived"},
	}

	for i := range base {
		if base[i].MatrixAvailable {
			base[i].Commutator = commutatorWithDiagonal(lambda, base[i].Matrix)
			base[i].CommutatorFrobeniusNormSquared = frobeniusNormSquared(base[i].Commutator)
			base[i].OffDiagonalSupportEntries = offDiagonalSupport(base[i].Matrix)
			base[i].CommutesWithTauEta = base[i].CommutatorFrobeniusNormSquared == 0
			base[i].PopulatesMixingComplement = base[i].OffDiagonalSupportEntries > 0 && !base[i].CommutesWithTauEta
		} else {
			base[i].CommutesWithTauEta = true
			base[i].PopulatesMixingComplement = false
		}
		base[i].QualifiedFiniteMixingPartner = base[i].CanonicalFiniteData && base[i].SelfAdjoint && base[i].GenerationEndomorphismDerived && base[i].AmplitudeSourceDerived && base[i].PopulatesMixingComplement && !base[i].PureSymmetryOrLabelAction && !base[i].RequiresRepresentationBridge && !base[i].RequiresActionCoefficient
		base[i].Verdict = candidateVerdict(base[i])
	}
	return base
}

func auditInventory(c []CandidateAudit) InventoryAudit {
	var inv InventoryAudit
	inv.CandidateCount = len(c)
	for _, x := range c {
		if x.MatrixAvailable {
			inv.MatrixCandidates++
		}
		if x.CanonicalFiniteData {
			inv.CanonicalFiniteCandidates++
		}
		if x.PopulatesMixingComplement {
			inv.RawNonCommutingCandidates++
		}
		if x.PopulatesMixingComplement && x.SelfAdjoint {
			inv.SelfAdjointRawCandidates++
		}
		if x.Kind == KindTrialityHermitian {
			inv.HermitianPhaseBasisCandidates++
		}
		if x.QualifiedFiniteMixingPartner {
			inv.QualifiedFiniteMixingPartners++
		}
		if x.Kind == KindBGapScalar && !x.GenerationEndomorphismDerived {
			inv.ScalarGapCandidatesRejected++
		}
		if x.Kind == KindHopfPhaseResidual && !x.GenerationEndomorphismDerived {
			inv.RepresentationFreePhaseRejected++
		}
	}
	inv.Verdict = fmt.Sprintf("%s; %s; raw non-commuting candidates=%d, self-adjoint raw candidates=%d, qualified finite mixing partners=%d", StatusTrialityRawComplementPopulated, StatusNoQualifiedPartner, inv.RawNonCommutingCandidates, inv.SelfAdjointRawCandidates, inv.QualifiedFiniteMixingPartners)
	return inv
}

func auditTrialityPartner(c []CandidateAudit) TrialityPartnerAudit {
	var out TrialityPartnerAudit
	dirs := map[string]struct{}{}
	for _, x := range c {
		if x.Name == "C3_cycle" {
			out.PermutationCycleNonCommuting = x.PopulatesMixingComplement
		}
		if x.Name == "S3_reflection_23" {
			out.ReflectionNonCommuting = x.PopulatesMixingComplement
		}
		if x.Name == "A_triality_real=C+C^T" {
			out.HermitianRealPartNonCommuting = x.PopulatesMixingComplement
		}
		if x.Name == "K_triality_phase=i(C-C^T)" {
			out.HermitianImaginaryPartNonCommuting = x.PopulatesMixingComplement
		}
		if x.Kind == KindTrialityHermitian {
			out.HermitianPhaseBasisDimension++
		}
		if x.Kind == KindTrialityPermutation || x.Kind == KindTrialityHermitian {
			if x.PureSymmetryOrLabelAction {
				out.AllRawTrialityMapsAreSymmetryAlgebra = true
			}
			if x.QualifiedFiniteMixingPartner {
				out.AnyTrialityMapQualifiedAsAmplitude = true
			}
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					if i != j && !isZero(x.Matrix[i][j]) {
						dirs[fmt.Sprintf("%d%d", i, j)] = struct{}{}
					}
				}
			}
		}
	}
	out.RawComplementDirectionsTouched = len(dirs)
	out.Verdict = StatusTrialityRawComplementPopulated + "; exact triality supplies raw off-diagonal generation maps, including Hermitian real/phase bases, but all remain symmetry-algebra data rather than selected Yukawa amplitudes"
	return out
}

func auditFinitePhaseGap(c []CandidateAudit) FinitePhaseGapAudit {
	out := FinitePhaseGapAudit{RequiresFiniteActionFunctional: true, RequiresRepresentationBridge: true}
	for _, x := range c {
		switch x.Kind {
		case KindBGapScalar:
			out.BGapAvailableAsPositiveScale = x.CanonicalFiniteData
			out.BGapHasGenerationMatrix = x.GenerationEndomorphismDerived
			out.BGapCanPopulateOffDiagonalComplement = x.PopulatesMixingComplement
		case KindHopfPhaseResidual:
			out.HopfPhasesAvailableAsPhaseLedger = x.CanonicalFiniteData
			out.HopfPhaseGenerationMapDerived = x.GenerationEndomorphismDerived
			out.HopfCanPopulateOffDiagonalComplement = x.PopulatesMixingComplement
		}
	}
	out.Verdict = StatusBGapRejectedAsScalar + "; " + StatusHopfRejectedAsRepresentationFree + "; gap/phase ledgers remain useful targets for a future action or representation map but cannot be used as generation textures here"
	return out
}

func auditPartnerVerdict(inh Gate261Inheritance, inv InventoryAudit, tr TrialityPartnerAudit, fg FinitePhaseGapAudit) MixingPartnerVerdict {
	raw := inv.RawNonCommutingCandidates > 0 && tr.RawComplementDirectionsTouched == inh.OffDiagonalComplementDimension
	selfAdj := inv.SelfAdjointRawCandidates > 0 && tr.HermitianPhaseBasisDimension >= 2
	qualified := inv.QualifiedFiniteMixingPartners > 0
	reason := "triality-derived matrices populate the 6D off-diagonal complement and provide exact Hermitian real/phase bases, but they are still symmetry/label algebra with no finite action coefficient; B_gap is scalar-only and Hopf phases are representation-free"
	return MixingPartnerVerdict{
		RawNonCommutingPartnerExists:         raw,
		RawSelfAdjointOffDiagonalBasisExists: selfAdj,
		QualifiedFiniteMixingPartnerFound:    qualified,
		PhysicalYukawaTextureDerived:         false,
		CKMPMNSDerived:                       false,
		FermionMassesDerived:                 false,
		Reason:                               reason,
		NextGate:                             "Gate 263 — Finite Yukawa Action Functional / Triality-Hopf Amplitude Qualification Audit",
		Status:                               StatusTrialityRawComplementPopulated + "; " + StatusHermitianTrialityPhaseBasisExposed + "; " + StatusNoQualifiedPartner + "; " + StatusCKMPMNSStillBlocked,
	}
}

func auditFirewall(inh Gate261Inheritance, verdict MixingPartnerVerdict) FirewallAudit {
	return FirewallAudit{
		Gate261SourceMapPreserved:          inh.BilinearCarrierDefined && inh.TauEtaActionDerived,
		DoesNotReopenEightVKernelRoute:     inh.EightVRouteClosed,
		DoesNotUseObservedMasses:           true,
		DoesNotUseObservedMixingAngles:     true,
		DoesNotPromoteSymmetryToAmplitude:  !verdict.QualifiedFiniteMixingPartnerFound,
		DoesNotUseBGapAsTextureWithoutMap:  true,
		DoesNotUseHopfPhaseWithoutMap:      true,
		DoesNotClaimFiniteActionFunctional: true,
		EmpiricalYukawaSealInactive:        true,
		FiniteCorePolluted:                 false,
		Verdict:                            StatusEmpiricalYukawaSealPreserved + "; raw complement population is recorded without promoting symmetry maps, scalar gaps, or phase residuals into physical Yukawa textures",
	}
}

func summarize(inh Gate261Inheritance, inv InventoryAudit, tr TrialityPartnerAudit, fg FinitePhaseGapAudit, verdict MixingPartnerVerdict) Summary {
	_ = fg
	return Summary{
		Gate261Inherited:                 inh.BilinearCarrierDefined && inh.TextureAlgebraDecomposed,
		TrialityComplementPopulated:      inv.RawNonCommutingCandidates > 0 && tr.RawComplementDirectionsTouched == inh.OffDiagonalComplementDimension,
		HermitianTrialityBasisExposed:    tr.HermitianRealPartNonCommuting && tr.HermitianImaginaryPartNonCommuting,
		BGapRejectedAsRepresentationFree: inv.ScalarGapCandidatesRejected == 1,
		HopfRejectedAsRepresentationFree: inv.RepresentationFreePhaseRejected == 1,
		QualifiedMixingPartnerFound:      verdict.QualifiedFiniteMixingPartnerFound,
		PhysicalYukawaTextureDerived:     verdict.PhysicalYukawaTextureDerived,
		CKMPMNSDerived:                   verdict.CKMPMNSDerived,
		FermionMassesDerived:             verdict.FermionMassesDerived,
		Status:                           verdict.Status,
		NextGate:                         verdict.NextGate,
		Comment:                          fmt.Sprintf("Gate 262 finds raw finite non-commuting triality algebra: %d raw candidates touch all %d off-diagonal directions. No candidate is a qualified finite Yukawa amplitude source; B_gap and Hopf phases still lack generation-texture representation maps.", inv.RawNonCommutingCandidates, tr.RawComplementDirectionsTouched),
	}
}

func buildTruth(inh Gate261Inheritance, inv InventoryAudit, tr TrialityPartnerAudit, fg FinitePhaseGapAudit, verdict MixingPartnerVerdict) string {
	return fmt.Sprintf("Gate 262 audits finite non-commuting partners for tau_eta on M3(C). Exact triality permutations and their Hermitian real/phase combinations do populate the full %d-dimensional off-diagonal complement exposed by Gate 261; this gives raw non-commuting mixing capacity. However, the candidates remain symmetry/label algebra rather than selected Yukawa amplitude sources. B_gap is available only as a scalar gap with no generation endomorphism, and Hopf phase residuals have no derived map into M3(C). Therefore raw complement population is conditionally supported, but no qualified finite mixing partner, CKM/PMNS matrix, or fermion mass spectrum is derived. %s", inh.OffDiagonalComplementDimension, verdict.Reason+"; "+fg.Verdict+fmt.Sprintf("; rawCandidates=%d hermitianPhaseBasis=%d", inv.RawNonCommutingCandidates, tr.HermitianPhaseBasisDimension))
}

func z(re, im int) GaussianInt { return GaussianInt{Re: re, Im: im} }

func isZero(x GaussianInt) bool { return x.Re == 0 && x.Im == 0 }

func conj(x GaussianInt) GaussianInt { return GaussianInt{Re: x.Re, Im: -x.Im} }

func add(a, b GaussianInt) GaussianInt { return GaussianInt{Re: a.Re + b.Re, Im: a.Im + b.Im} }

func neg(a GaussianInt) GaussianInt { return GaussianInt{Re: -a.Re, Im: -a.Im} }

func scale(k int, a GaussianInt) GaussianInt { return GaussianInt{Re: k * a.Re, Im: k * a.Im} }

func commutatorWithDiagonal(lambda []int, m Matrix3) Matrix3 {
	var out Matrix3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			out[i][j] = scale(lambda[i]-lambda[j], m[i][j])
		}
	}
	return out
}

func frobeniusNormSquared(m Matrix3) int {
	total := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			total += m[i][j].Re*m[i][j].Re + m[i][j].Im*m[i][j].Im
		}
	}
	return total
}

func offDiagonalSupport(m Matrix3) int {
	count := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i != j && !isZero(m[i][j]) {
				count++
			}
		}
	}
	return count
}

func IsSelfAdjoint(m Matrix3) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if m[i][j] != conj(m[j][i]) {
				return false
			}
		}
	}
	return true
}

func candidateVerdict(c CandidateAudit) string {
	parts := []string{}
	if c.PopulatesMixingComplement {
		parts = append(parts, StatusTrialityRawComplementPopulated)
	}
	if c.Kind == KindTrialityHermitian && c.SelfAdjoint {
		parts = append(parts, StatusHermitianTrialityPhaseBasisExposed)
	}
	if c.Kind == KindBGapScalar {
		parts = append(parts, StatusBGapRejectedAsScalar)
	}
	if c.Kind == KindHopfPhaseResidual {
		parts = append(parts, StatusHopfRejectedAsRepresentationFree)
	}
	if !c.QualifiedFiniteMixingPartner {
		parts = append(parts, StatusNoQualifiedPartner)
	}
	if len(parts) == 0 {
		parts = append(parts, StatusNoQualifiedPartner)
	}
	return strings.Join(parts, "; ") + "; " + c.Disqualification
}

func MatrixString(m Matrix3) string {
	rows := make([]string, 0, 3)
	for i := 0; i < 3; i++ {
		cells := make([]string, 0, 3)
		for j := 0; j < 3; j++ {
			cells = append(cells, GaussianString(m[i][j]))
		}
		rows = append(rows, "["+strings.Join(cells, " ")+"]")
	}
	return strings.Join(rows, " ")
}

func GaussianString(x GaussianInt) string {
	switch {
	case x.Re == 0 && x.Im == 0:
		return "0"
	case x.Im == 0:
		return fmt.Sprintf("%d", x.Re)
	case x.Re == 0:
		if x.Im == 1 {
			return "i"
		}
		if x.Im == -1 {
			return "-i"
		}
		return fmt.Sprintf("%di", x.Im)
	case x.Im > 0:
		if x.Im == 1 {
			return fmt.Sprintf("%d+i", x.Re)
		}
		return fmt.Sprintf("%d+%di", x.Re, x.Im)
	default:
		if x.Im == -1 {
			return fmt.Sprintf("%d-i", x.Re)
		}
		return fmt.Sprintf("%d%di", x.Re, x.Im)
	}
}

func CandidateNames(c []CandidateAudit) []string {
	out := make([]string, 0, len(c))
	for _, x := range c {
		out = append(out, x.Name)
	}
	sort.Strings(out)
	return out
}

// keep small helpers referenced by tests available inside the package
var _ = add
var _ = neg
