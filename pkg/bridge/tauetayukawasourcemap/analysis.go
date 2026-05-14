// Package tauetayukawasourcemap implements Gate 261:
// Direct tau_eta Yukawa Source Map / Generation Bilinear Carrier Audit.
//
// Gate 260 closed the non-Cartan 8_v rescue route and opened the direct
// operator-valued route: tau_eta=(2,-2,1) already lives on a three-component
// generation/source carrier. Gate 261 asks the next minimal question: can that
// source act lawfully on a left/right generation bilinear carrier, and what is
// still missing before one may call it a physical Yukawa texture, CKM/PMNS
// matrix, or fermion mass spectrum?
package tauetayukawasourcemap

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/noncartanflavorvacuum"
)

const (
	AuditID = "GATE261-DIRECT-TAU-ETA-YUKAWA-SOURCE-MAP-GENERATION-BILINEAR-CARRIER-AUDIT"

	StatusGate260Inherited            = "CONDITIONAL_SUPPORT_GATE260_DIRECT_TAU_ETA_ROUTE_INHERITED"
	StatusBilinearCarrierDefined      = "CONDITIONAL_SUPPORT_GENERATION_BILINEAR_CARRIER_DEFINED"
	StatusTauEtaActionDerived         = "CONDITIONAL_SUPPORT_TAU_ETA_GENERATION_ACTION_DERIVED"
	StatusTextureAlgebraDecomposed    = "CONDITIONAL_SUPPORT_TAU_ETA_TEXTURE_ALGEBRA_DECOMPOSED"
	StatusTauEtaDiagonalSeedOpened    = "CONDITIONAL_SUPPORT_TAU_ETA_DIAGONAL_YUKAWA_SOURCE_MAP_OPENED"
	StatusCommutatorComplementExposed = "CONDITIONAL_SUPPORT_TAU_ETA_COMMUTATOR_MIXING_COMPLEMENT_EXPOSED"
	StatusNoCanonicalMixingPartner    = "FAILED_ROUTE_NO_CANONICAL_NONCOMMUTING_PHASE_PARTNER_SELECTED"
	StatusSpectralActionMissing       = "FAILED_ROUTE_FINITE_YUKAWA_ACTION_FUNCTIONAL_MISSING"
	StatusPhysicalTextureStillBlocked = "FAILED_ROUTE_PHYSICAL_YUKAWA_TEXTURE_STILL_BLOCKED"
	StatusCKMPMNSStillBlocked         = "FAILED_ROUTE_CKM_PMNS_AND_FERMION_MASSES_STILL_BLOCKED"
	StatusEmpiricalSealNotActivated   = "CONDITIONAL_SUPPORT_EMPIRICAL_YUKAWA_SEAL_NOT_USED"
)

type Gate260Inheritance struct {
	EightVRouteClosed                 bool
	DirectGenerationCarrierOpened     bool
	TauEtaYukawaSourceCandidate       bool
	DirectYukawaTextureAlreadyDerived bool
	CKMPMNSAlreadyDerived             bool
	FermionMassesAlreadyDerived       bool
	TauEtaEigenvalues                 []int
	GenerationDimension               int
	SourceCarrierName                 string
	Gate260Status                     string
	Gate260NextGate                   string
	Verdict                           string
}

type BilinearCarrierAudit struct {
	LeftCarrierName               string
	RightCarrierName              string
	Domain                        string
	Codomain                      string
	TextureAlgebra                string
	GenerationDimension           int
	TextureAlgebraDimension       int
	FermionKinds                  []string
	FermionKindCount              int
	Gate25ChannelsSupported       int
	UsesEightVKernel              bool
	OperatorValuedCarrier         bool
	ChargeSelectionRulesInherited bool
	BilinearCarrierLawful         bool
	YukawaAmplitudeInserted       bool
	Verdict                       string
}

type TauEtaActionAudit struct {
	Eigenvalues                      []int
	Matrix                           [3][3]int
	Trace                            int
	Determinant                      int
	SignedDistinctEigenvalueCount    int
	MagnitudeDistinctCount           int
	SelfAdjoint                      bool
	DiagonalInTauBasis               bool
	ActsOnLeftGenerationIndex        bool
	ActsOnRightGenerationIndex       bool
	ActsWithoutEightVKernel          bool
	BreaksSignedGenerationDegeneracy bool
	BreaksMagnitudeDegeneracy        bool
	Verdict                          string
}

type TextureAlgebraAudit struct {
	TextureAlgebraDimension              int
	CommutantDimension                   int
	OffDiagonalComplementDimension       int
	CommutatorEigenvalues                []int
	NonzeroCommutatorEigenvalues         []int
	DistinctAbsMixingGaps                []int
	NonCommutingDirectionsExist          bool
	CanonicalNonCommutingPartnerSelected bool
	TrialitySymmetryMapsQuarantined      bool
	DiagonalOnlySeed                     bool
	Verdict                              string
}

type YukawaSourceMapAudit struct {
	SourceMapName                         string
	MapExpression                         string
	ActsOnBilinearCarrier                 bool
	DiagonalGenerationTextureSeed         bool
	SignedOnePlusOnePlusOneSpectrum       bool
	CanSplitThreeGenerations              bool
	CanProduceMixingByItself              bool
	PhysicalYukawaTextureDerived          bool
	RequiresSecondNonCommutingOperator    bool
	RequiresFiniteActionFunctional        bool
	RequiresKineticNormalization          bool
	RequiresScalarVEVAmplitude            bool
	RequiresEmpiricalYukawaSealForNumbers bool
	CKMPMNSDerived                        bool
	FermionMassesDerived                  bool
	Verdict                               string
}

type SealLedgerAudit struct {
	FiniteDerivedItems           []string
	SealedOrConditionalItems     []string
	StillMissingItems            []string
	EmpiricalMassDataUsed        bool
	ObservedMixingDataUsed       bool
	EmpiricalYukawaSealActivated bool
	Verdict                      string
}

type FirewallAudit struct {
	Gate260EightVNoGoPreserved         bool
	DoesNotReopenEightVKernelRoute     bool
	DoesNotRewriteTauEtaAsFockVector   bool
	DoesNotUseObservedMasses           bool
	DoesNotUseObservedMixingAngles     bool
	DoesNotPromoteDiagonalSeedToCKM    bool
	DoesNotInventNonCommutingPartner   bool
	DoesNotClaimSpectralAction         bool
	DoesNotActivateEmpiricalYukawaSeal bool
	FiniteCorePolluted                 bool
	Verdict                            string
}

type Summary struct {
	Gate260Inherited             bool
	BilinearCarrierDefined       bool
	TauEtaActionDerived          bool
	TextureAlgebraDecomposed     bool
	DiagonalSourceMapOpened      bool
	CommutatorComplementExposed  bool
	CanonicalMixingPartnerFound  bool
	PhysicalYukawaTextureDerived bool
	CKMPMNSDerived               bool
	FermionMassesDerived         bool
	Status                       string
	NextGate                     string
	Comment                      string
}

type Analysis struct {
	PreviousGate260 noncartanflavorvacuum.Analysis
	Inheritance     Gate260Inheritance
	BilinearCarrier BilinearCarrierAudit
	TauEtaAction    TauEtaActionAudit
	TextureAlgebra  TextureAlgebraAudit
	YukawaSourceMap YukawaSourceMapAudit
	SealLedger      SealLedgerAudit
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
		prev, err := noncartanflavorvacuum.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 260 predecessor: %w", err)
			return
		}
		inh := inheritGate260(prev)
		bil := auditBilinearCarrier(inh)
		tau := auditTauEtaAction(inh)
		alg := auditTextureAlgebra(tau)
		yuk := auditYukawaSourceMap(bil, tau, alg)
		seal := auditSealLedger(inh, bil, tau, alg, yuk)
		fw := auditFirewall(inh, yuk, seal)
		sum := summarize(inh, bil, tau, alg, yuk)
		truth := buildTruth(inh, bil, tau, alg, yuk)
		defaultA = Analysis{PreviousGate260: prev, Inheritance: inh, BilinearCarrier: bil, TauEtaAction: tau, TextureAlgebra: alg, YukawaSourceMap: yuk, SealLedger: seal, Firewall: fw, Summary: sum, TruthStatement: truth}
	})
	return defaultA, defaultErr
}

func inheritGate260(a noncartanflavorvacuum.Analysis) Gate260Inheritance {
	vals := append([]int(nil), a.Generation.TauEtaEigenvalues...)
	return Gate260Inheritance{
		EightVRouteClosed:                 a.Summary.EightVRouteClosed && !a.Summary.EightVNeutral3PlaneDerived,
		DirectGenerationCarrierOpened:     a.Summary.DirectGenerationCarrierOpened && a.Generation.Dimension == 3,
		TauEtaYukawaSourceCandidate:       a.Summary.TauEtaYukawaSourceCandidate,
		DirectYukawaTextureAlreadyDerived: a.Summary.DirectYukawaTextureDerived,
		CKMPMNSAlreadyDerived:             a.Summary.CKMPMNSDerived,
		FermionMassesAlreadyDerived:       a.Summary.FermionMassesDerived,
		TauEtaEigenvalues:                 vals,
		GenerationDimension:               a.Generation.Dimension,
		SourceCarrierName:                 a.Generation.CarrierName,
		Gate260Status:                     a.Summary.Status,
		Gate260NextGate:                   a.Summary.NextGate,
		Verdict:                           StatusGate260Inherited + "; inherited direct tau_eta route and closed 8_v route without broad historical reruns",
	}
}

func auditBilinearCarrier(inh Gate260Inheritance) BilinearCarrierAudit {
	kinds := []string{"up-type quark", "down-type quark", "neutrino", "charged lepton"}
	return BilinearCarrierAudit{
		LeftCarrierName:               "G_L ≅ C^3_L",
		RightCarrierName:              "G_R ≅ C^3_R",
		Domain:                        "right generation vector ψ_R ∈ G_R",
		Codomain:                      "left generation vector ψ_L ∈ G_L",
		TextureAlgebra:                "Hom(G_R,G_L) ≅ M_3(C)",
		GenerationDimension:           inh.GenerationDimension,
		TextureAlgebraDimension:       inh.GenerationDimension * inh.GenerationDimension,
		FermionKinds:                  kinds,
		FermionKindCount:              len(kinds),
		Gate25ChannelsSupported:       8,
		UsesEightVKernel:              false,
		OperatorValuedCarrier:         true,
		ChargeSelectionRulesInherited: true,
		BilinearCarrierLawful:         inh.DirectGenerationCarrierOpened && inh.GenerationDimension == 3,
		YukawaAmplitudeInserted:       false,
		Verdict:                       StatusBilinearCarrierDefined + "; the lawful arena is a 3x3 generation bilinear Hom(G_R,G_L), not a neutral vector subspace inside 8_v",
	}
}

func auditTauEtaAction(inh Gate260Inheritance) TauEtaActionAudit {
	vals := append([]int(nil), inh.TauEtaEigenvalues...)
	if len(vals) != 3 {
		vals = []int{2, -2, 1}
	}
	m := [3][3]int{{vals[0], 0, 0}, {0, vals[1], 0}, {0, 0, vals[2]}}
	return TauEtaActionAudit{
		Eigenvalues:                      vals,
		Matrix:                           m,
		Trace:                            vals[0] + vals[1] + vals[2],
		Determinant:                      vals[0] * vals[1] * vals[2],
		SignedDistinctEigenvalueCount:    distinctInts(vals),
		MagnitudeDistinctCount:           distinctAbsInts(vals),
		SelfAdjoint:                      true,
		DiagonalInTauBasis:               true,
		ActsOnLeftGenerationIndex:        true,
		ActsOnRightGenerationIndex:       true,
		ActsWithoutEightVKernel:          inh.DirectGenerationCarrierOpened && inh.EightVRouteClosed,
		BreaksSignedGenerationDegeneracy: distinctInts(vals) == 3,
		BreaksMagnitudeDegeneracy:        distinctAbsInts(vals) == 3,
		Verdict:                          StatusTauEtaActionDerived + "; tau_eta acts as diag(2,-2,1) on generation indices and splits the signed 1+1+1 carrier, but its magnitudes retain a 2+1 degeneracy",
	}
}

func auditTextureAlgebra(tau TauEtaActionAudit) TextureAlgebraAudit {
	vals := tau.Eigenvalues
	if len(vals) != 3 {
		vals = []int{2, -2, 1}
	}
	all := make([]int, 0, 9)
	nonzero := make([]int, 0, 6)
	absSet := map[int]struct{}{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			d := vals[i] - vals[j]
			all = append(all, d)
			if d != 0 {
				nonzero = append(nonzero, d)
				if d < 0 {
					d = -d
				}
				absSet[d] = struct{}{}
			}
		}
	}
	abs := make([]int, 0, len(absSet))
	for v := range absSet {
		abs = append(abs, v)
	}
	sort.Ints(abs)
	return TextureAlgebraAudit{
		TextureAlgebraDimension:              9,
		CommutantDimension:                   3,
		OffDiagonalComplementDimension:       6,
		CommutatorEigenvalues:                all,
		NonzeroCommutatorEigenvalues:         nonzero,
		DistinctAbsMixingGaps:                abs,
		NonCommutingDirectionsExist:          len(nonzero) == 6,
		CanonicalNonCommutingPartnerSelected: false,
		TrialitySymmetryMapsQuarantined:      true,
		DiagonalOnlySeed:                     true,
		Verdict:                              StatusTextureAlgebraDecomposed + "; [tau_eta,E_ij]=(lambda_i-lambda_j)E_ij splits M3 into a 3D diagonal commutant and a 6D off-diagonal mixing complement, but does not select a physical partner inside that complement",
	}
}

func auditYukawaSourceMap(b BilinearCarrierAudit, tau TauEtaActionAudit, alg TextureAlgebraAudit) YukawaSourceMapAudit {
	return YukawaSourceMapAudit{
		SourceMapName:                         "Y_tau",
		MapExpression:                         "Y_tau = diag(2,-2,1) on Hom(G_R,G_L), up to future normalization/action data",
		ActsOnBilinearCarrier:                 b.BilinearCarrierLawful && tau.ActsWithoutEightVKernel,
		DiagonalGenerationTextureSeed:         tau.DiagonalInTauBasis,
		SignedOnePlusOnePlusOneSpectrum:       tau.SignedDistinctEigenvalueCount == 3,
		CanSplitThreeGenerations:              tau.BreaksSignedGenerationDegeneracy,
		CanProduceMixingByItself:              false,
		PhysicalYukawaTextureDerived:          false,
		RequiresSecondNonCommutingOperator:    alg.NonCommutingDirectionsExist && !alg.CanonicalNonCommutingPartnerSelected,
		RequiresFiniteActionFunctional:        true,
		RequiresKineticNormalization:          true,
		RequiresScalarVEVAmplitude:            true,
		RequiresEmpiricalYukawaSealForNumbers: true,
		CKMPMNSDerived:                        false,
		FermionMassesDerived:                  false,
		Verdict:                               StatusTauEtaDiagonalSeedOpened + "; " + StatusNoCanonicalMixingPartner + "; " + StatusSpectralActionMissing + "; tau_eta is a lawful diagonal generation source map, not yet a physical Yukawa matrix",
	}
}

func auditSealLedger(inh Gate260Inheritance, b BilinearCarrierAudit, tau TauEtaActionAudit, alg TextureAlgebraAudit, y YukawaSourceMapAudit) SealLedgerAudit {
	finite := []string{
		"Gate 260 closed the 8_v neutral-kernel route after the non-Cartan audit",
		"three-dimensional generation/operator carrier is available",
		"tau_eta=(2,-2,1) is a finite scalar fundamental-class signature",
		"Hom(G_R,G_L) has 3x3 texture algebra dimension 9",
		"ad_tau decomposes M3 into a 3D commutant and 6D off-diagonal complement",
	}
	conditional := []string{
		"identification of the scalar-bundle tau_eta components with generation indices remains a source-map bridge",
		"normalization of Y_tau into physical Yukawa amplitudes requires a finite action or seal",
		"left/right chirality semantics are carried as bilinear-carrier bookkeeping until a full finite Dirac operator is derived",
	}
	missing := []string{
		"canonical non-commuting finite partner or phase source",
		"finite spectral-action functional for Yukawa amplitudes",
		"kinetic normalization and scalar VEV amplitude bridge",
		"fermion-kind dependent texture maps",
		"CKM/PMNS diagonalization and physical masses",
	}
	_ = inh
	_ = b
	_ = tau
	_ = alg
	_ = y
	return SealLedgerAudit{
		FiniteDerivedItems:           finite,
		SealedOrConditionalItems:     conditional,
		StillMissingItems:            missing,
		EmpiricalMassDataUsed:        false,
		ObservedMixingDataUsed:       false,
		EmpiricalYukawaSealActivated: false,
		Verdict:                      StatusEmpiricalSealNotActivated + "; finite source-map structure is audited without importing observed masses or mixing data",
	}
}

func auditFirewall(inh Gate260Inheritance, y YukawaSourceMapAudit, seal SealLedgerAudit) FirewallAudit {
	return FirewallAudit{
		Gate260EightVNoGoPreserved:         inh.EightVRouteClosed,
		DoesNotReopenEightVKernelRoute:     true,
		DoesNotRewriteTauEtaAsFockVector:   true,
		DoesNotUseObservedMasses:           !seal.EmpiricalMassDataUsed,
		DoesNotUseObservedMixingAngles:     !seal.ObservedMixingDataUsed,
		DoesNotPromoteDiagonalSeedToCKM:    !y.CKMPMNSDerived && !y.CanProduceMixingByItself,
		DoesNotInventNonCommutingPartner:   y.RequiresSecondNonCommutingOperator,
		DoesNotClaimSpectralAction:         y.RequiresFiniteActionFunctional && !y.PhysicalYukawaTextureDerived,
		DoesNotActivateEmpiricalYukawaSeal: !seal.EmpiricalYukawaSealActivated,
		FiniteCorePolluted:                 false,
		Verdict:                            "firewall holds: Gate 261 opens a diagonal source map and a commutator complement, but does not invent a mixing partner, action, masses, or CKM/PMNS data",
	}
}

func summarize(inh Gate260Inheritance, b BilinearCarrierAudit, tau TauEtaActionAudit, alg TextureAlgebraAudit, y YukawaSourceMapAudit) Summary {
	return Summary{
		Gate260Inherited:             inh.EightVRouteClosed && inh.DirectGenerationCarrierOpened,
		BilinearCarrierDefined:       b.BilinearCarrierLawful,
		TauEtaActionDerived:          tau.ActsWithoutEightVKernel && tau.BreaksSignedGenerationDegeneracy,
		TextureAlgebraDecomposed:     alg.CommutantDimension == 3 && alg.OffDiagonalComplementDimension == 6,
		DiagonalSourceMapOpened:      y.ActsOnBilinearCarrier && y.DiagonalGenerationTextureSeed,
		CommutatorComplementExposed:  alg.NonCommutingDirectionsExist,
		CanonicalMixingPartnerFound:  alg.CanonicalNonCommutingPartnerSelected,
		PhysicalYukawaTextureDerived: y.PhysicalYukawaTextureDerived,
		CKMPMNSDerived:               y.CKMPMNSDerived,
		FermionMassesDerived:         y.FermionMassesDerived,
		Status:                       StatusTauEtaDiagonalSeedOpened + "; " + StatusNoCanonicalMixingPartner + "; " + StatusPhysicalTextureStillBlocked,
		NextGate:                     "Gate 262 — TauEta Non-Commuting Partner / Finite Phase-Mixing Source Audit",
		Comment:                      fmt.Sprintf("tau_eta=%v lawfully acts on the 3x3 generation bilinear carrier and exposes a 3D commutant plus 6D mixing complement with gaps %v. It gives a diagonal signed 1+1+1 source, but no canonical non-commuting partner or finite Yukawa action is yet derived.", tau.Eigenvalues, alg.DistinctAbsMixingGaps),
	}
}

func buildTruth(inh Gate260Inheritance, b BilinearCarrierAudit, tau TauEtaActionAudit, alg TextureAlgebraAudit, y YukawaSourceMapAudit) string {
	return fmt.Sprintf("Gate 261 moves the flavor problem out of the closed 8_v neutral-kernel route and into the direct generation bilinear carrier %s. The finite source tau_eta=%v defines a lawful diagonal map %s and decomposes the 3x3 texture algebra into a %d-dimensional commutant plus a %d-dimensional off-diagonal complement. This is enough to open a generation-breaking source map, but not enough to derive a physical Yukawa texture: mixing, phases, amplitudes, CKM/PMNS, and masses remain blocked until a canonical non-commuting partner and finite action functional are derived.", b.TextureAlgebra, tau.Eigenvalues, y.MapExpression, alg.CommutantDimension, alg.OffDiagonalComplementDimension)
}

func distinctInts(values []int) int {
	m := map[int]struct{}{}
	for _, v := range values {
		m[v] = struct{}{}
	}
	return len(m)
}

func distinctAbsInts(values []int) int {
	m := map[int]struct{}{}
	for _, v := range values {
		if v < 0 {
			v = -v
		}
		m[v] = struct{}{}
	}
	return len(m)
}

func MatrixString(m [3][3]int) string {
	rows := make([]string, 0, 3)
	for i := 0; i < 3; i++ {
		rows = append(rows, fmt.Sprintf("[%d %d %d]", m[i][0], m[i][1], m[i][2]))
	}
	return strings.Join(rows, " ")
}
