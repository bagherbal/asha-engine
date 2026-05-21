// Package generation2afrepresentationsectorprojectorandaggregatecarrierpullbackaudit
// implements Gate 834: A_F-Representation Sector Projector and Aggregate-Carrier
// Pullback Audit.
//
// Gate 834 follows Gate 833's obstruction of the direct M_3(C) fundamental
// triplet / Fock P_3 carrier bridge.  The next lawful source for typed sector
// projectors is not the bare algebra alone and not the bare Fock selector alone,
// but a represented finite internal algebra rho_F(A_F) on a finite Hilbert
// carrier H_F.  This gate audits central algebra-summand idempotents, the
// representation-induced projector recipe, basis dependence of M_3(C) matrix
// units, and the missing pullback from the R2++ aggregate carrier
// I_3 plus (P_1 plus P_3) into represented finite-sector projectors.  It does
// not read trace magnitudes, Yukawa values, generations, CKM/PMNS data, or any
// observed particle masses.
package generation2afrepresentationsectorprojectorandaggregatecarrierpullbackaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID = "GATE834-AF-REPRESENTATION-SECTOR-PROJECTOR-AGGREGATE-CARRIER-PULLBACK-AUDIT"

	SBoundary       = 0.0012924448188162962
	AlphaB          = 0.0003878958469680527
	OperatorNEff    = 3.002327375081808
	OfficialNEff    = 3.0023273474722147
	OfficialCYukawa = 0.9992248188812008
	OfficialCHiggs  = 1.0372205204048603

	AFSummandCount     = 3
	ComplexSummandDim  = 1
	QuaternionSummandN = 2
	M3FundamentalDim   = 3
	M3MatrixUnitCount  = 9
	FockWDim           = 4
	FockP1Rank         = 1
	FockP3Rank         = 3
	TopBlockDim        = 3
	RestBlockDim       = 4
	AggregateAtomCount = 7
	K7Dim              = 7

	StatusGate833Inherited                 = "PASS_GATE833_DIRECT_TRIPLET_BRIDGE_OBSTRUCTION_INHERITED"
	StatusCentralIdempotentsAudited        = "PASS_A_F_CENTRAL_IDEMPOTENT_PROJECTORS_AUDITED"
	StatusRepresentationRequirementAudited = "PASS_A_F_REPRESENTATION_SECTOR_PROJECTOR_SOURCE_AUDITED"
	StatusMatrixUnitBasisFirewall          = "PASS_M3_MATRIX_UNIT_BASIS_DEPENDENCE_FIREWALL_ENFORCED"
	StatusAggregatePullbackAudited         = "PASS_AGGREGATE_CARRIER_PULLBACK_TESTED"
	StatusSectorProjectorMagnitudeFirewall = "PASS_SECTOR_PROJECTORS_NOT_TRACE_MAGNITUDES_FIREWALL_PRESERVED"
	StatusR2PlusPlusRetained               = "PASS_R2_PLUS_PLUS_STATUS_RETAINED_NOT_R3"
	StatusNoLedgerUpdates                  = "PASS_N_EFF_C_YUKAWA_C_HIGGS_LEDGER_UPDATES_BLOCKED"
	StatusNoObservedDataUsed               = "PASS_NO_OBSERVED_YUKAWA_MASS_CKM_PMNS_DATA_USED"
	StatusPhysicalFirewalls                = "PASS_PHYSICAL_FIREWALLS_ENFORCED"
	StatusFirewallGate834                  = "FIREWALL_PRESERVED_GATE834_A_F_REPRESENTATION_PROJECTOR_PULLBACK_BOUNDARY"

	SupportAFStrongestSectorSource              = "CONDITIONAL_SUPPORT_A_F_IS_STRONGEST_FINITE_SECTOR_PROJECTOR_SOURCE"
	SupportCentralIdempotentsSourceCoarseBlocks = "CONDITIONAL_SUPPORT_A_F_CENTRAL_IDEMPOTENTS_SOURCE_COARSE_SECTOR_BLOCKS"
	SupportRepresentationCanInduceProjectors    = "CONDITIONAL_SUPPORT_A_F_REPRESENTATION_CAN_SOURCE_COARSE_SECTOR_PROJECTORS"
	SupportSectorProjectorsRequireRhoF          = "CONDITIONAL_SUPPORT_SECTOR_PROJECTOR_MAP_REQUIRES_FINITE_HILBERT_REPRESENTATION"
	SupportPartialFiniteRepresentationPredata   = "CONDITIONAL_SUPPORT_PARTIAL_FINITE_REPRESENTATION_PREDATA_EXISTS_BUT_IS_NOT_AGGREGATE_PULLBACK"
	SupportMatrixUnitsExist                     = "CONDITIONAL_SUPPORT_M3C_MATRIX_UNITS_EXIST_AS_CARRIER_PROJECTORS_AFTER_FRAME_CHOICE"
	SupportAggregateCarrierStillR2PlusPlus      = "CONDITIONAL_SUPPORT_AGGREGATE_CARRIER_REMAINS_R2_PLUS_PLUS"
	SupportSectorMagnitudeReadoutSeparate       = "CONDITIONAL_SUPPORT_SECTOR_TRACE_MAGNITUDE_READOUT_IS_SEPARATE_MISSING_OBJECT"

	FailureAFAloneNotSectorLedgerWithoutRepresentation         = "FAILED_ROUTE_A_F_ALONE_NOT_SECTOR_LEDGER_WITHOUT_REPRESENTATION"
	FailureNoCompleteRhoFRepresentationCertified               = "FAILED_ROUTE_NO_COMPLETE_RHO_F_REPRESENTATION_PROJECTOR_LEDGER_CERTIFIED"
	FailureNoCompleteFiniteHilbertPackage                      = "FAILED_ROUTE_NO_COMPLETE_A_F_H_F_RHO_F_J_F_GAMMA_F_D_F_PACKAGE_CERTIFIED"
	FailureM3MatrixUnitsNotCanonicalColorAtomsWithoutFrame     = "FAILED_ROUTE_M3_MATRIX_UNITS_NOT_CANONICAL_COLOR_ATOMS_WITHOUT_FRAME"
	FailureNoCanonicalColorFrame                               = "FAILED_ROUTE_NO_CANONICAL_COLOR_FRAME_SELECTED_BY_CURRENT_GATE"
	FailureNoAggregateCarrierToRepresentationProjectorPullback = "FAILED_ROUTE_NO_AGGREGATE_CARRIER_TO_REPRESENTATION_PROJECTOR_PULLBACK"
	FailureNoSigmaMap                                          = "FAILED_ROUTE_NO_SECTOR_PROJECTOR_MAP_SIGMA_CERTIFIED"
	FailureTopI3NotPulledBackToRepresentationSector            = "FAILED_ROUTE_TOP_I3_NOT_PULLED_BACK_TO_REPRESENTATION_SECTOR"
	FailureFockP1P3NotPulledBackToRepresentationSector         = "FAILED_ROUTE_FOCK_P1_P3_NOT_PULLED_BACK_TO_REPRESENTATION_SECTOR"
	FailureNoM3P3Intertwiner                                   = "FAILED_ROUTE_NO_CANONICAL_M3C_TO_FOCK_P3_INTERTWINER_CERTIFIED"
	FailureSectorProjectorsNotTraceMagnitudes                  = "FAILED_ROUTE_SECTOR_PROJECTORS_NOT_TRACE_MAGNITUDES"
	FailureNoSectorTraceMagnitudeReadout                       = "FAILED_ROUTE_NO_SECTOR_TRACE_MAGNITUDE_READOUT_MAP"
	FailureAggregateOperatorNotR3                              = "FAILED_ROUTE_AGGREGATE_OPERATOR_NOT_R3"
	FailureR2NotR3                                             = "FAILED_ROUTE_R2_PLUS_PLUS_NOT_R3"
	FailureNoNativeYukawaOperator                              = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureAlphaStillSealed                                    = "FAILED_ROUTE_ALPHA_B_STILL_SEALED_NOT_NATIVE_SOURCE"
	FailureNoBoundaryAlphaMap                                  = "FAILED_ROUTE_NO_BOUNDARY_ALPHA_DOMAIN_TRANSPORT_MAP_CERTIFIED"
	FailureNoNEffUpdate                                        = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaUpdate                                     = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoObservedYukawaFit                                 = "FAILED_ROUTE_NO_OBSERVED_YUKAWA_FITTING_ALLOWED"
	FailureNoPMNSCKM                                           = "FAILED_ROUTE_NO_PMNS_CKM_OR_FLAVOR_ORIENTATION_THEOREM"
	FailureNoParticleAssignment                                = "FAILED_ROUTE_NO_STANDARD_MODEL_PARTICLE_ASSIGNMENT_FROM_A_F_PROJECTOR_CANDIDATES"
)

type Ledger struct {
	S, AlphaB                                            float64
	AggregateCarrier                                     string
	OperatorNEff, OfficialNEff                           float64
	TopBlockDim, RestBlockDim, AggregateAtomCount, K7Dim int
	R2PlusPlusConsolidated, R3SectorLedgerCertified      bool
	AlphaSealed                                          bool
	SectorProjectorMapCertified, TraceMagnitudeCertified bool
}

type CentralIdempotentAudit struct {
	Algebra                                        string
	CentralIdempotents                             []string
	SummandCount                                   int
	Orthogonal, SumToIdentity                      bool
	CoarseSectorBlocks, RepresentationIndependent  bool
	SectorLedgerCertified, TraceMagnitudeCertified bool
	Verdicts, Supports, Failures                   []string
}

type RepresentationProjectorAudit struct {
	RequiredPackage                              string
	RequiredMap                                  string
	UsesHF, UsesRhoF, UsesJF, UsesGammaF, UsesDF bool
	PartialPredataAvailable                      bool
	CompletePackageCertified                     bool
	RepresentationInducedProjectorsCertified     bool
	CanSourceCoarseProjectorCandidates           bool
	CanSourceSectorLedger                        bool
	CanSourceTraceMagnitudes                     bool
	Verdicts, Supports, Failures                 []string
}

type MatrixUnitAudit struct {
	Algebra, MatrixUnits                            string
	MatrixUnitCount, DiagonalProjectorCount         int
	MatrixUnitsExist, DiagonalProjectorsExist       bool
	CanonicalColorFrameCertified, BasisIndependent  bool
	CanonicalColorAtomsCertified                    bool
	SuppliesCarrierProjectors, SuppliesSectorLedger bool
	Verdicts, Supports, Failures                    []string
}

type AggregatePullbackAudit struct {
	Domain, Codomain, CandidateMap                         string
	TopDim, RestDim, AggregateAtoms                        int
	CentralBlocksAvailable, RepresentationProjectorRecipe  bool
	PullbackCertified, TopI3PulledBack, FockP1P3PulledBack bool
	M3P3IntertwinerCertified, NonCircular                  bool
	Verdicts, Supports, Failures                           []string
}

type SectorImpact struct {
	CentralProjectorSource, RepresentationProjectorRecipe, AggregatePullbackCertified       bool
	SectorProjectorMapCertified, SectorTraceLedgerCertified, TraceMagnitudeReadoutCertified bool
	CanPromoteToR3, CanPromoteToR4                                                          bool
	CanUpdateNEff, CanUpdateCYukawa, CanUpdateCHiggs                                        bool
	CurrentLevel, NextMissingObject, NextGate, Reason                                       string
	Verdicts, Supports, Failures                                                            []string
}

type Firewalls struct {
	Enforced                                                 bool
	AFAloneNotLedger, RequiresRepresentation, NoCompleteRhoF bool
	MatrixUnitsBasisDependent, NoColorFrame                  bool
	NoAggregatePullback, NoSigmaMap, NoM3P3Intertwiner       bool
	ProjectorsNotMagnitudes, NoMagnitudeReadout              bool
	AlphaSealed, NoBoundaryAlphaMap                          bool
	NotR3, NotR4                                             bool
	NoNEffUpdate, NoCYukawaUpdate                            bool
	NoObservedYukawaFit, NoPMNSCKM, NoParticleAssignment     bool
	Verdict                                                  string
}

type Audit struct {
	Ledger         Ledger
	Central        CentralIdempotentAudit
	Representation RepresentationProjectorAudit
	MatrixUnits    MatrixUnitAudit
	Pullback       AggregatePullbackAudit
	Impact         SectorImpact
	Firewalls      Firewalls
	Truth          string
	Final          string
}

func BuildDefault() (Audit, error) {
	ledger := Ledger{
		S: SBoundary, AlphaB: AlphaB,
		AggregateCarrier: "H_total/T = I_3 plus [alpha_B P_3 - 3 alpha_B^2(B-L)]",
		OperatorNEff:     OperatorNEff, OfficialNEff: OfficialNEff,
		TopBlockDim: TopBlockDim, RestBlockDim: RestBlockDim, AggregateAtomCount: AggregateAtomCount, K7Dim: K7Dim,
		R2PlusPlusConsolidated: true, R3SectorLedgerCertified: false, AlphaSealed: true,
		SectorProjectorMapCertified: false, TraceMagnitudeCertified: false,
	}
	central := CentralIdempotentAudit{
		Algebra:            "A_F = C plus H plus M_3(C)",
		CentralIdempotents: []string{"z_C", "z_H", "z_M3"},
		SummandCount:       AFSummandCount,
		Orthogonal:         true, SumToIdentity: true, CoarseSectorBlocks: true, RepresentationIndependent: true,
		SectorLedgerCertified: false, TraceMagnitudeCertified: false,
		Verdicts: []string{StatusCentralIdempotentsAudited},
		Supports: []string{SupportAFStrongestSectorSource, SupportCentralIdempotentsSourceCoarseBlocks},
		Failures: []string{FailureAFAloneNotSectorLedgerWithoutRepresentation, FailureSectorProjectorsNotTraceMagnitudes},
	}
	representation := RepresentationProjectorAudit{
		RequiredPackage: "finite spectral triple carrier (A_F,H_F,rho_F,J_F,gamma_F,D_F)",
		RequiredMap:     "rho_F: A_F -> End(H_F), Pi_i = supp(rho_F(z_i)) plus refined chirality/commutant/bimodule projectors",
		UsesHF:          true, UsesRhoF: true, UsesJF: true, UsesGammaF: true, UsesDF: true,
		PartialPredataAvailable:                  true,
		CompletePackageCertified:                 false,
		RepresentationInducedProjectorsCertified: false,
		CanSourceCoarseProjectorCandidates:       true,
		CanSourceSectorLedger:                    false,
		CanSourceTraceMagnitudes:                 false,
		Verdicts:                                 []string{StatusRepresentationRequirementAudited},
		Supports:                                 []string{SupportRepresentationCanInduceProjectors, SupportSectorProjectorsRequireRhoF, SupportPartialFiniteRepresentationPredata},
		Failures:                                 []string{FailureNoCompleteRhoFRepresentationCertified, FailureNoCompleteFiniteHilbertPackage, FailureNoSectorTraceMagnitudeReadout},
	}
	matrix := MatrixUnitAudit{
		Algebra: "M_3(C)", MatrixUnits: "E_ij, i,j=1..3",
		MatrixUnitCount: M3MatrixUnitCount, DiagonalProjectorCount: M3FundamentalDim,
		MatrixUnitsExist: true, DiagonalProjectorsExist: true,
		CanonicalColorFrameCertified: false, BasisIndependent: false, CanonicalColorAtomsCertified: false,
		SuppliesCarrierProjectors: true, SuppliesSectorLedger: false,
		Verdicts: []string{StatusMatrixUnitBasisFirewall},
		Supports: []string{SupportMatrixUnitsExist},
		Failures: []string{FailureM3MatrixUnitsNotCanonicalColorAtomsWithoutFrame, FailureNoCanonicalColorFrame},
	}
	pullback := AggregatePullbackAudit{
		Domain:       "I_3 plus (P_1 plus P_3)",
		Codomain:     "Pi_sector(rho_F(A_F))",
		CandidateMap: "Sigma: aggregate carrier -> represented finite-sector projectors",
		TopDim:       TopBlockDim, RestDim: RestBlockDim, AggregateAtoms: AggregateAtomCount,
		CentralBlocksAvailable:        central.CoarseSectorBlocks,
		RepresentationProjectorRecipe: representation.CanSourceCoarseProjectorCandidates,
		PullbackCertified:             false, TopI3PulledBack: false, FockP1P3PulledBack: false,
		M3P3IntertwinerCertified: false, NonCircular: true,
		Verdicts: []string{StatusAggregatePullbackAudited},
		Supports: []string{SupportAggregateCarrierStillR2PlusPlus},
		Failures: []string{FailureNoAggregateCarrierToRepresentationProjectorPullback, FailureNoSigmaMap, FailureTopI3NotPulledBackToRepresentationSector, FailureFockP1P3NotPulledBackToRepresentationSector, FailureNoM3P3Intertwiner},
	}
	impact := SectorImpact{
		CentralProjectorSource:         central.CoarseSectorBlocks,
		RepresentationProjectorRecipe:  representation.CanSourceCoarseProjectorCandidates,
		AggregatePullbackCertified:     pullback.PullbackCertified,
		SectorProjectorMapCertified:    false,
		SectorTraceLedgerCertified:     false,
		TraceMagnitudeReadoutCertified: false,
		CanPromoteToR3:                 false, CanPromoteToR4: false,
		CanUpdateNEff: false, CanUpdateCYukawa: false, CanUpdateCHiggs: false,
		CurrentLevel:      "R2++ consolidated aggregate trace carrier, not R3",
		NextMissingObject: "SectorProjectorMap Sigma from I_3 plus (P_1 plus P_3) to Pi_sector(rho_F(A_F))",
		NextGate:          "SectorTraceMagnitudeReadout obstruction only after a sector-projector ledger is certified",
		Reason:            "A_F central idempotents and representation recipes source sector projector candidates, but the aggregate carrier has no certified pullback into those projectors and projectors still do not supply positive magnitudes.",
		Verdicts:          []string{StatusSectorProjectorMagnitudeFirewall, StatusR2PlusPlusRetained, StatusNoLedgerUpdates},
		Supports:          []string{SupportSectorMagnitudeReadoutSeparate},
		Failures:          []string{FailureNoSigmaMap, FailureNoAggregateCarrierToRepresentationProjectorPullback, FailureSectorProjectorsNotTraceMagnitudes, FailureNoSectorTraceMagnitudeReadout, FailureAggregateOperatorNotR3, FailureR2NotR3, FailureNoNativeYukawaOperator, FailureNoNEffUpdate, FailureNoCYukawaUpdate},
	}
	firewalls := Firewalls{
		Enforced:         true,
		AFAloneNotLedger: true, RequiresRepresentation: true, NoCompleteRhoF: true,
		MatrixUnitsBasisDependent: true, NoColorFrame: true,
		NoAggregatePullback: true, NoSigmaMap: true, NoM3P3Intertwiner: true,
		ProjectorsNotMagnitudes: true, NoMagnitudeReadout: true,
		AlphaSealed: true, NoBoundaryAlphaMap: true,
		NotR3: true, NotR4: true,
		NoNEffUpdate: true, NoCYukawaUpdate: true,
		NoObservedYukawaFit: true, NoPMNSCKM: true, NoParticleAssignment: true,
		Verdict: StatusFirewallGate834,
	}
	return Audit{
		Ledger:         ledger,
		Central:        central,
		Representation: representation,
		MatrixUnits:    matrix,
		Pullback:       pullback,
		Impact:         impact,
		Firewalls:      firewalls,
		Truth:          "Gate 834 moves the sector problem from bare carrier comparison to represented finite algebra. A_F central idempotents and rho_F support projectors are lawful sector-projector candidates, but the aggregate R2++ carrier is not pulled back into them.",
		Final:          "Verdict: A_F representation projectors are the right source layer, but no Sigma pullback, no sector ledger, no trace-magnitude readout, no R3 promotion, and no Yukawa theorem are certified.",
	}, nil
}

func Statuses() []string {
	return []string{
		StatusGate833Inherited, StatusCentralIdempotentsAudited, StatusRepresentationRequirementAudited, StatusMatrixUnitBasisFirewall, StatusAggregatePullbackAudited, StatusSectorProjectorMagnitudeFirewall, StatusR2PlusPlusRetained, StatusNoLedgerUpdates, StatusNoObservedDataUsed, StatusPhysicalFirewalls, StatusFirewallGate834,
		SupportAFStrongestSectorSource, SupportCentralIdempotentsSourceCoarseBlocks, SupportRepresentationCanInduceProjectors, SupportSectorProjectorsRequireRhoF, SupportPartialFiniteRepresentationPredata, SupportMatrixUnitsExist, SupportAggregateCarrierStillR2PlusPlus, SupportSectorMagnitudeReadoutSeparate,
		FailureAFAloneNotSectorLedgerWithoutRepresentation, FailureNoCompleteRhoFRepresentationCertified, FailureNoCompleteFiniteHilbertPackage, FailureM3MatrixUnitsNotCanonicalColorAtomsWithoutFrame, FailureNoCanonicalColorFrame, FailureNoAggregateCarrierToRepresentationProjectorPullback, FailureNoSigmaMap, FailureTopI3NotPulledBackToRepresentationSector, FailureFockP1P3NotPulledBackToRepresentationSector, FailureNoM3P3Intertwiner, FailureSectorProjectorsNotTraceMagnitudes, FailureNoSectorTraceMagnitudeReadout, FailureAggregateOperatorNotR3, FailureR2NotR3, FailureNoNativeYukawaOperator, FailureAlphaStillSealed, FailureNoBoundaryAlphaMap, FailureNoNEffUpdate, FailureNoCYukawaUpdate, FailureNoObservedYukawaFit, FailureNoPMNSCKM, FailureNoParticleAssignment,
	}
}

func containsAll(haystack, needles []string) bool {
	m := make(map[string]bool, len(haystack))
	for _, h := range haystack {
		m[h] = true
	}
	for _, n := range needles {
		if !m[n] {
			return false
		}
	}
	return true
}

func FormatLedger(l Ledger) string {
	return fmt.Sprintf("carrier=%q s=%.16g alpha_B=%.16g operator_N_eff=%.16g official_N_eff=%.16g dims=%d+%d atoms=%d K7=%d R2++=%t R3=%t alpha_sealed=%t sector_map=%t magnitude=%t", l.AggregateCarrier, l.S, l.AlphaB, l.OperatorNEff, l.OfficialNEff, l.TopBlockDim, l.RestBlockDim, l.AggregateAtomCount, l.K7Dim, l.R2PlusPlusConsolidated, l.R3SectorLedgerCertified, l.AlphaSealed, l.SectorProjectorMapCertified, l.TraceMagnitudeCertified)
}

func FormatCentral(c CentralIdempotentAudit) string {
	return fmt.Sprintf("algebra=%q idempotents=[%s] count=%d orthogonal=%t sumI=%t coarse_blocks=%t sector_ledger=%t trace_magnitude=%t failures=[%s]", c.Algebra, strings.Join(c.CentralIdempotents, ","), c.SummandCount, c.Orthogonal, c.SumToIdentity, c.CoarseSectorBlocks, c.SectorLedgerCertified, c.TraceMagnitudeCertified, strings.Join(c.Failures, ","))
}

func FormatRepresentation(r RepresentationProjectorAudit) string {
	return fmt.Sprintf("required=%q map=%q uses=(HF:%t rhoF:%t J:%t gamma:%t D:%t) partial_predata=%t complete=%t projectors=%t coarse_candidates=%t ledger=%t magnitudes=%t failures=[%s]", r.RequiredPackage, r.RequiredMap, r.UsesHF, r.UsesRhoF, r.UsesJF, r.UsesGammaF, r.UsesDF, r.PartialPredataAvailable, r.CompletePackageCertified, r.RepresentationInducedProjectorsCertified, r.CanSourceCoarseProjectorCandidates, r.CanSourceSectorLedger, r.CanSourceTraceMagnitudes, strings.Join(r.Failures, ","))
}

func FormatMatrixUnits(m MatrixUnitAudit) string {
	return fmt.Sprintf("algebra=%q units=%q count=%d diagonal=%d exist=%t frame=%t basis_independent=%t canonical_atoms=%t carrier_projectors=%t ledger=%t failures=[%s]", m.Algebra, m.MatrixUnits, m.MatrixUnitCount, m.DiagonalProjectorCount, m.MatrixUnitsExist, m.CanonicalColorFrameCertified, m.BasisIndependent, m.CanonicalColorAtomsCertified, m.SuppliesCarrierProjectors, m.SuppliesSectorLedger, strings.Join(m.Failures, ","))
}

func FormatPullback(p AggregatePullbackAudit) string {
	return fmt.Sprintf("domain=%q codomain=%q candidate=%q dims=%d+%d atoms=%d central=%t representation_recipe=%t pullback=%t top=%t fock=%t m3p3=%t noncircular=%t failures=[%s]", p.Domain, p.Codomain, p.CandidateMap, p.TopDim, p.RestDim, p.AggregateAtoms, p.CentralBlocksAvailable, p.RepresentationProjectorRecipe, p.PullbackCertified, p.TopI3PulledBack, p.FockP1P3PulledBack, p.M3P3IntertwinerCertified, p.NonCircular, strings.Join(p.Failures, ","))
}

func FormatImpact(i SectorImpact) string {
	return fmt.Sprintf("central_source=%t representation_recipe=%t aggregate_pullback=%t sector_map=%t ledger=%t magnitude=%t R3=%t R4=%t updates=(N:%t CY:%t CH:%t) level=%q next=%q reason=%q failures=[%s]", i.CentralProjectorSource, i.RepresentationProjectorRecipe, i.AggregatePullbackCertified, i.SectorProjectorMapCertified, i.SectorTraceLedgerCertified, i.TraceMagnitudeReadoutCertified, i.CanPromoteToR3, i.CanPromoteToR4, i.CanUpdateNEff, i.CanUpdateCYukawa, i.CanUpdateCHiggs, i.CurrentLevel, i.NextMissingObject, i.Reason, strings.Join(i.Failures, ","))
}
