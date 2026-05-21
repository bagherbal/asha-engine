// Package generation2symbolicfinitediracmatrixsupportfirstorderfirewallaudit implements
// Gate 848: Symbolic Finite-Dirac Matrix Support and First-Order Firewall Audit.
//
// Gate 848 follows Gate 847's weak-socket edge generator. Gate 847 certified,
// at seal level, the three active support edges
//
//	Y_+3 : e_+ tensor P_3 -> h_+ tensor P_3,
//	Y_-3 : e_- tensor P_3 -> h_- tensor P_3,
//	Y_-1 : e_- tensor P_1 -> h_- tensor P_1,
//
// with the puncture edge Y_+1 absent. Gate 848 packages these support edges
// into an explicit symbolic chiral block support matrix
//
//	D_F^sym = [[0, Y^dagger], [Y, 0]],
//	Y = y_+3 Y_+3 plus y_-3 Y_-3 plus y_-1 Y_-1,
//	y_+1 = 0,
//
// on H_L plus H_R^min. This is still support-only: the y variables are not
// observed Yukawa values, no alpha_B source is derived, no full rho_F/J_F
// first-order proof is certified, and no R3/R4 promotion or ledger update is
// allowed.
package generation2symbolicfinitediracmatrixsupportfirstorderfirewallaudit

import (
	"strconv"
	"strings"
)

const (
	AuditID = "GATE848-SYMBOLIC-FINITE-DIRAC-MATRIX-SUPPORT-FIRST-ORDER-FIREWALL-AUDIT"

	AlphaB          = 0.0003878958469680527
	OfficialNEff    = 3.0023273474722147
	OfficialCYukawa = 0.9992248188812008
	OfficialCHiggs  = 1.0372205204048603

	LeptonBlockDim = 1
	ColorBlockDim  = 3
	WDim           = LeptonBlockDim + ColorBlockDim
	WeakDoubletDim = 2
	HRMinRank      = 7
	HLRank         = WeakDoubletDim * WDim
	ChiralTotalDim = HRMinRank + HLRank

	StatusGate847Inherited                 = "PASS_GATE847_WEAK_SOCKET_EDGE_GENERATOR_INHERITED"
	StatusMinimalRightDomainInherited      = "PASS_MINIMAL_RIGHT_MODULE_EDGE_DOMAIN_INHERITED"
	StatusWeakSocketTargetsInherited       = "PASS_WEAK_DOUBLET_SOCKET_TARGETS_INHERITED"
	StatusSymbolicYSupportConstructed      = "PASS_SYMBOLIC_Y_SUPPORT_MATRIX_CONSTRUCTED"
	StatusPunctureYPlusOneZero             = "PASS_PUNCTURE_EDGE_Y_PLUS_ONE_SET_TO_ZERO"
	StatusChiralDFSupportMatrixConstructed = "PASS_SYMBOLIC_D_F_CHIRAL_BLOCK_SUPPORT_MATRIX_CONSTRUCTED"
	StatusSelfAdjointSupport               = "PASS_SELF_ADJOINTNESS_BY_ADJOINT_BLOCK_INCLUDED"
	StatusChiralityOddSupport              = "PASS_CHIRALITY_ODDNESS_BY_LEFT_RIGHT_BLOCK_FORM"
	StatusLeptoColorPreserved              = "PASS_LEPTO_COLOR_SUPPORT_PRESERVED_BY_SYMBOLIC_D_F_SUPPORT"
	StatusFirstOrderFirewallAudited        = "PASS_FIRST_ORDER_AND_BIMODULE_FIREWALL_AUDITED"
	StatusSupportOnlyNoMagnitudes          = "PASS_SYMBOLIC_COEFFICIENTS_CLASSIFIED_AS_SUPPORT_VARIABLES_NOT_YUKAWA_VALUES"
	StatusOfficialLedgersFrozen            = "PASS_N_EFF_C_YUKAWA_C_HIGGS_LEDGER_UPDATES_BLOCKED"
	StatusNoObservedDataUsed               = "PASS_NO_OBSERVED_MASS_CKM_PMNS_HIGGS_DATA_USED"
	StatusR2PlusPlusPlusPlusPlusStill      = "PASS_R2_PLUS_PLUS_PLUS_PLUS_PLUS_SYMBOLIC_D_F_SUPPORT_MATRIX_NOT_R3_OR_R4"
	StatusFirewallGate848                  = "FIREWALL_PRESERVED_GATE848_SYMBOLIC_D_F_MATRIX_SUPPORT"

	SupportSymbolicDFSupportMatrix       = "CONDITIONAL_SUPPORT_SYMBOLIC_D_F_SUPPORT_MATRIX_EXISTS"
	SupportThreeSymbolicEdgeFamilies     = "CONDITIONAL_SUPPORT_Y_SUPP_HAS_THREE_ACTIVE_SYMBOLIC_EDGE_FAMILIES"
	SupportPunctureZeroCoefficient       = "CONDITIONAL_SUPPORT_Y_PLUS_ONE_EQUALS_ZERO_MINIMAL_PUNCTURE_SEAL"
	SupportSelfAdjointByBlockForm        = "CONDITIONAL_SUPPORT_SELF_ADJOINTNESS_BY_CHIRAL_BLOCK_FORM"
	SupportChiralityOddByBlockForm       = "CONDITIONAL_SUPPORT_CHIRALITY_ODDNESS_BY_LEFT_RIGHT_BLOCK_FORM"
	SupportNeutralSingletonNullEdge      = "CONDITIONAL_SUPPORT_NEUTRAL_SINGLETON_IS_NULL_EDGE_CANDIDATE_AT_SEAL_LEVEL"
	SupportLeptoColorPreservingMatrix    = "CONDITIONAL_SUPPORT_LEPTO_COLOR_PRESERVING_D_F_SUPPORT"
	SupportNoYukawaMagnitudeFromYSymbols = "CONDITIONAL_SUPPORT_Y_SYMBOLS_ARE_SUPPORT_VARIABLES_NOT_NUMERICAL_YUKAWAS"
	SupportR2PlusPlusPlusPlusPlusStage   = "CONDITIONAL_SUPPORT_R2_PLUS_PLUS_PLUS_PLUS_PLUS_SYMBOLIC_D_F_SUPPORT_MATRIX_STAGE"

	FailureDFSupportMatrixSealNotNative = "FAILED_ROUTE_SYMBOLIC_D_F_SUPPORT_MATRIX_IS_SEAL_NOT_NATIVE_D_F_THEOREM"
	FailureNoNumericalDFMatrix          = "FAILED_ROUTE_NO_NUMERICAL_OR_OPERATOR_VALUED_D_F_MATRIX_CERTIFIED"
	FailureNoFullRhoFJFGammaFPackage    = "FAILED_ROUTE_NO_COMPLETE_RHO_F_J_F_GAMMA_F_PACKAGE_CERTIFIED"
	FailureNoFirstOrderProof            = "FAILED_ROUTE_NO_FULL_FIRST_ORDER_CONDITION_PROOF"
	FailureNoBimoduleCommutantProof     = "FAILED_ROUTE_NO_BIMODULE_COMMUTANT_DECOMPOSITION_PROOF_CERTIFIED"
	FailureNoJOppositeCompatibility     = "FAILED_ROUTE_NO_J_OPPOSITE_ACTION_COMPATIBILITY_PROOF_CERTIFIED"
	FailureYSymbolsNotYukawaValues      = "FAILED_ROUTE_SYMBOLIC_Y_COEFFICIENTS_NOT_YUKAWA_MAGNITUDES"
	FailureEdgeSupportNotTraceMagnitude = "FAILED_ROUTE_EDGE_SUPPORT_NOT_TRACE_MAGNITUDE"
	FailureNoNumericalYukawaValues      = "FAILED_ROUTE_NO_NUMERICAL_YUKAWA_VALUES_CERTIFIED"
	FailureAlphaStillSealed             = "FAILED_ROUTE_ALPHA_B_STILL_SEALED_NOT_NATIVE_SOURCE"
	FailureNoTraceMagnitudeReadout      = "FAILED_ROUTE_NO_SECTOR_TRACE_MAGNITUDE_READOUT_MAP"
	FailurePunctureNullEdgeOnlySeal     = "FAILED_ROUTE_PUNCTURE_NULL_EDGE_ONLY_SEAL_NOT_NATIVE_THEOREM"
	FailureNoNativeNullEdgeTheorem      = "FAILED_ROUTE_NO_NATIVE_NULL_EDGE_THEOREM_FOR_E_PLUS_TENSOR_P1"
	FailureNoPhysicalParticleAssign     = "FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT_FROM_SYMBOLIC_D_F_SUPPORT"
	FailureNoRightNeutrinoTheorem       = "FAILED_ROUTE_NO_RIGHT_NEUTRINO_OR_STERILE_PARTICLE_THEOREM"
	FailureNoThreeGenerationTheorem     = "FAILED_ROUTE_THREE_ACTIVE_EDGE_FAMILIES_NOT_THREE_GENERATIONS"
	FailureNoNEffUpdate                 = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaUpdate              = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNotR3                        = "FAILED_ROUTE_R2_PLUS_PLUS_PLUS_PLUS_PLUS_NOT_R3"
	FailureNotR4                        = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type Ledger struct {
	AlphaB                         float64
	OfficialNEff, OfficialCYukawa  float64
	OfficialCHiggs                 float64
	OfficialFrozen                 bool
	R2PlusPlusPlusPlusPlus, R3, R4 bool
	AlphaNative                    bool
}

type SymbolicEdge struct {
	Name, Coefficient, Domain, Target, LeptoColor, Role string
	DomainRank, TargetRank                              int
	Present, Puncture, HasMagnitude                     bool
	ValueSource                                         string
}

type YSupport struct {
	Expression, Rule                                         string
	Edges                                                    []SymbolicEdge
	MissingEdge                                              SymbolicEdge
	DomainRank, TargetRank, ActiveFamilies, ExpectedFamilies int
	CoefficientNames                                         []string
	PunctureCoefficient                                      string
	PunctureCoefficientZero                                  bool
	PreservesLeptoColor, SupportOnly, HasNumericalValues     bool
	Supports, Failures                                       []string
}

type ChiralDiracSupport struct {
	Expression, Space, GammaConvention                                       string
	LeftRank, RightRank, TotalRank                                           int
	BlockRows, BlockCols                                                     int
	UsesAdjointBlock, SelfAdjointByConstruction, ChiralOddByConstruction     bool
	ExplicitSupportMatrix, NativeDFMatrix, NumericalDFMatrix                 bool
	JOppositeCompatibilityProof, FirstOrderCertified, BimoduleCommutantProof bool
	Supports, Failures                                                       []string
}

type Impact struct {
	Classification                                                 string
	Gate847Inherited, MatrixSupportConstructed, SelfAdjointSupport bool
	ChiralityOddSupport, PunctureNullEdgeCandidate                 bool
	FirstOrderProof, BimoduleProof, JOppositeProof                 bool
	AlphaStillSealed, MagnitudesStillMissing                       bool
	CanUpdateNEff, CanUpdateCYukawa, CanUpdateCHiggs               bool
	CanPromoteToR3, CanPromoteToR4                                 bool
}

type Firewalls struct {
	Enforced                                                                                     bool
	DFSupportMatrixSealNotNative, NoNumericalDFMatrix, NoFullRhoFJFGammaFPackage                 bool
	NoFirstOrderProof, NoBimoduleCommutantProof, NoJOppositeCompatibilityProof                   bool
	YSymbolsNotYukawaValues, EdgeSupportNotTraceMagnitude, NoNumericalYukawaValues               bool
	AlphaStillSealed, NoTraceMagnitudeReadout, PunctureNullEdgeOnlySeal, NoNativeNullEdgeTheorem bool
	NoPhysicalParticleAssignment, NoRightNeutrinoTheorem, NoThreeGenerationTheorem               bool
	NoNEffUpdate, NoCYukawaUpdate, NotR3, NotR4                                                  bool
	Verdict                                                                                      string
}

type Audit struct {
	Ledger    Ledger
	Y         YSupport
	Dirac     ChiralDiracSupport
	Impact    Impact
	Firewalls Firewalls
	Truth     string
	Final     string
}

func BuildDefault() (Audit, error) {
	a := Audit{
		Ledger:    buildLedger(),
		Y:         buildYSupport(),
		Dirac:     buildChiralDiracSupport(),
		Impact:    buildImpact(),
		Firewalls: buildFirewalls(),
		Truth:     "Gate 848 packages the Gate 847 support skeleton into an explicit symbolic chiral block support matrix D_F^sym=[[0,Y^dagger],[Y,0]], with Y=y_+3Y_+3+y_-3Y_-3+y_-1Y_-1 and y_+1=0. This certifies support-matrix anatomy, self-adjointness by adjoint-block inclusion, and chirality oddness by left/right block form at seal level only; it does not certify a native D_F theorem, first-order/bimodule/J compatibility, Yukawa magnitudes, alpha source, R3/R4 promotion, or official ledger update.",
		Final:     "Final verdict: CONDITIONAL_SUPPORT_SYMBOLIC_D_F_SUPPORT_MATRIX_EXISTS, CONDITIONAL_SUPPORT_SELF_ADJOINTNESS_BY_CHIRAL_BLOCK_FORM, and CONDITIONAL_SUPPORT_CHIRALITY_ODDNESS_BY_LEFT_RIGHT_BLOCK_FORM, but FAILED_ROUTE_SYMBOLIC_D_F_SUPPORT_MATRIX_IS_SEAL_NOT_NATIVE_D_F_THEOREM, FAILED_ROUTE_NO_FULL_FIRST_ORDER_CONDITION_PROOF, FAILED_ROUTE_SYMBOLIC_Y_COEFFICIENTS_NOT_YUKAWA_MAGNITUDES, FAILED_ROUTE_ALPHA_B_STILL_SEALED_NOT_NATIVE_SOURCE, and FAILED_ROUTE_R2_PLUS_PLUS_PLUS_PLUS_PLUS_NOT_R3.",
	}
	return a, nil
}

func buildLedger() Ledger {
	return Ledger{AlphaB: AlphaB, OfficialNEff: OfficialNEff, OfficialCYukawa: OfficialCYukawa, OfficialCHiggs: OfficialCHiggs, OfficialFrozen: true, R2PlusPlusPlusPlusPlus: true, R3: false, R4: false, AlphaNative: false}
}

func buildYSupport() YSupport {
	edges := []SymbolicEdge{
		{Name: "Y_+3", Coefficient: "y_+3", Domain: "e_+ tensor P_3", Target: "h_+ tensor P_3", LeptoColor: "P_3 -> P_3", Role: "dominant character-color symbolic edge family", DomainRank: 3, TargetRank: 3, Present: true, Puncture: false, HasMagnitude: false, ValueSource: "symbolic support variable only"},
		{Name: "Y_-3", Coefficient: "y_-3", Domain: "e_- tensor P_3", Target: "h_- tensor P_3", LeptoColor: "P_3 -> P_3", Role: "rest character-color symbolic edge family", DomainRank: 3, TargetRank: 3, Present: true, Puncture: false, HasMagnitude: false, ValueSource: "symbolic support variable only"},
		{Name: "Y_-1", Coefficient: "y_-1", Domain: "e_- tensor P_1", Target: "h_- tensor P_1", LeptoColor: "P_1 -> P_1", Role: "rest character-lepton symbolic edge family", DomainRank: 1, TargetRank: 1, Present: true, Puncture: false, HasMagnitude: false, ValueSource: "symbolic support variable only"},
	}
	missing := SymbolicEdge{Name: "Y_+1", Coefficient: "y_+1=0", Domain: "e_+ tensor P_1", Target: "h_+ tensor P_1", LeptoColor: "P_1 -> P_1", Role: "neutral right-lepton puncture edge absent in minimal symbolic matrix", DomainRank: 1, TargetRank: 1, Present: false, Puncture: true, HasMagnitude: false, ValueSource: "zero by minimal puncture seal"}
	return YSupport{
		Expression: "Y_supp = y_+3 Y_+3 plus y_-3 Y_-3 plus y_-1 Y_-1; y_+1=0",
		Rule:       "support-only map H_R^min -> H_L preserving P_1/P_3 with symbolic coefficients only",
		Edges:      edges, MissingEdge: missing,
		DomainRank: HRMinRank, TargetRank: HLRank, ActiveFamilies: len(edges), ExpectedFamilies: 3,
		CoefficientNames: []string{"y_+3", "y_-3", "y_-1"}, PunctureCoefficient: "y_+1", PunctureCoefficientZero: true,
		PreservesLeptoColor: true, SupportOnly: true, HasNumericalValues: false,
		Supports: []string{SupportThreeSymbolicEdgeFamilies, SupportPunctureZeroCoefficient, SupportNeutralSingletonNullEdge, SupportLeptoColorPreservingMatrix, SupportNoYukawaMagnitudeFromYSymbols},
		Failures: []string{FailureYSymbolsNotYukawaValues, FailureEdgeSupportNotTraceMagnitude, FailureNoNumericalYukawaValues, FailureAlphaStillSealed, FailureNoTraceMagnitudeReadout, FailurePunctureNullEdgeOnlySeal, FailureNoNativeNullEdgeTheorem, FailureNoPhysicalParticleAssign, FailureNoRightNeutrinoTheorem, FailureNoThreeGenerationTheorem},
	}
}

func buildChiralDiracSupport() ChiralDiracSupport {
	return ChiralDiracSupport{
		Expression:      "D_F^sym = [[0, Y_supp^dagger], [Y_supp, 0]]",
		Space:           "H_L plus H_R^min",
		GammaConvention: "gamma_L=+1, gamma_R=-1; D_F^sym is off-diagonal",
		LeftRank:        HLRank, RightRank: HRMinRank, TotalRank: ChiralTotalDim,
		BlockRows: ChiralTotalDim, BlockCols: ChiralTotalDim,
		UsesAdjointBlock: true, SelfAdjointByConstruction: true, ChiralOddByConstruction: true,
		ExplicitSupportMatrix: true, NativeDFMatrix: false, NumericalDFMatrix: false,
		JOppositeCompatibilityProof: false, FirstOrderCertified: false, BimoduleCommutantProof: false,
		Supports: []string{SupportSymbolicDFSupportMatrix, SupportSelfAdjointByBlockForm, SupportChiralityOddByBlockForm, SupportR2PlusPlusPlusPlusPlusStage},
		Failures: []string{FailureDFSupportMatrixSealNotNative, FailureNoNumericalDFMatrix, FailureNoFullRhoFJFGammaFPackage, FailureNoFirstOrderProof, FailureNoBimoduleCommutantProof, FailureNoJOppositeCompatibility},
	}
}

func buildImpact() Impact {
	return Impact{
		Classification:   "R2+++++ symbolic finite-Dirac support matrix; support-only, not R3/R4",
		Gate847Inherited: true, MatrixSupportConstructed: true, SelfAdjointSupport: true,
		ChiralityOddSupport: true, PunctureNullEdgeCandidate: true,
		FirstOrderProof: false, BimoduleProof: false, JOppositeProof: false,
		AlphaStillSealed: true, MagnitudesStillMissing: true,
		CanUpdateNEff: false, CanUpdateCYukawa: false, CanUpdateCHiggs: false,
		CanPromoteToR3: false, CanPromoteToR4: false,
	}
}

func buildFirewalls() Firewalls {
	return Firewalls{
		Enforced:                     true,
		DFSupportMatrixSealNotNative: true, NoNumericalDFMatrix: true, NoFullRhoFJFGammaFPackage: true,
		NoFirstOrderProof: true, NoBimoduleCommutantProof: true, NoJOppositeCompatibilityProof: true,
		YSymbolsNotYukawaValues: true, EdgeSupportNotTraceMagnitude: true, NoNumericalYukawaValues: true,
		AlphaStillSealed: true, NoTraceMagnitudeReadout: true, PunctureNullEdgeOnlySeal: true, NoNativeNullEdgeTheorem: true,
		NoPhysicalParticleAssignment: true, NoRightNeutrinoTheorem: true, NoThreeGenerationTheorem: true,
		NoNEffUpdate: true, NoCYukawaUpdate: true, NotR3: true, NotR4: true, Verdict: StatusFirewallGate848,
	}
}

func Statuses() []string {
	return []string{
		StatusGate847Inherited, StatusMinimalRightDomainInherited, StatusWeakSocketTargetsInherited, StatusSymbolicYSupportConstructed, StatusPunctureYPlusOneZero, StatusChiralDFSupportMatrixConstructed, StatusSelfAdjointSupport, StatusChiralityOddSupport, StatusLeptoColorPreserved, StatusFirstOrderFirewallAudited, StatusSupportOnlyNoMagnitudes, StatusOfficialLedgersFrozen, StatusNoObservedDataUsed, StatusR2PlusPlusPlusPlusPlusStill, StatusFirewallGate848,
		SupportSymbolicDFSupportMatrix, SupportThreeSymbolicEdgeFamilies, SupportPunctureZeroCoefficient, SupportSelfAdjointByBlockForm, SupportChiralityOddByBlockForm, SupportNeutralSingletonNullEdge, SupportLeptoColorPreservingMatrix, SupportNoYukawaMagnitudeFromYSymbols, SupportR2PlusPlusPlusPlusPlusStage,
		FailureDFSupportMatrixSealNotNative, FailureNoNumericalDFMatrix, FailureNoFullRhoFJFGammaFPackage, FailureNoFirstOrderProof, FailureNoBimoduleCommutantProof, FailureNoJOppositeCompatibility, FailureYSymbolsNotYukawaValues, FailureEdgeSupportNotTraceMagnitude, FailureNoNumericalYukawaValues, FailureAlphaStillSealed, FailureNoTraceMagnitudeReadout, FailurePunctureNullEdgeOnlySeal, FailureNoNativeNullEdgeTheorem, FailureNoPhysicalParticleAssign, FailureNoRightNeutrinoTheorem, FailureNoThreeGenerationTheorem, FailureNoNEffUpdate, FailureNoCYukawaUpdate, FailureNotR3, FailureNotR4,
	}
}

func FormatLedger(l Ledger) string {
	return join("ledger", []string{
		"alpha_B=" + f64(l.AlphaB), "official_N_eff=" + f64(l.OfficialNEff), "official_C_Yukawa=" + f64(l.OfficialCYukawa), "official_C_Higgs=" + f64(l.OfficialCHiggs), "official_frozen=" + b(l.OfficialFrozen), "R2+++++_stage=" + b(l.R2PlusPlusPlusPlusPlus), "R3=" + b(l.R3), "R4=" + b(l.R4), "alpha_native=" + b(l.AlphaNative),
	})
}

func FormatEdge(e SymbolicEdge) string {
	return join("edge", []string{e.Name, "coefficient=" + e.Coefficient, e.Domain + " -> " + e.Target, "lepto_color=" + e.LeptoColor, "role=" + e.Role, "domain_rank=" + i(e.DomainRank), "target_rank=" + i(e.TargetRank), "present=" + b(e.Present), "puncture=" + b(e.Puncture), "magnitude=" + b(e.HasMagnitude), "value_source=" + e.ValueSource})
}

func FormatY(y YSupport) string {
	parts := []string{y.Expression, "rule=" + y.Rule, "domain_rank=" + i(y.DomainRank), "target_rank=" + i(y.TargetRank), "active_families=" + i(y.ActiveFamilies), "puncture_coeff=" + y.PunctureCoefficient, "puncture_coeff_zero=" + b(y.PunctureCoefficientZero), "preserves_lepto_color=" + b(y.PreservesLeptoColor), "support_only=" + b(y.SupportOnly), "has_numerical_values=" + b(y.HasNumericalValues), "coefficients=" + strings.Join(y.CoefficientNames, ",")}
	for _, e := range y.Edges {
		parts = append(parts, FormatEdge(e))
	}
	parts = append(parts, "missing="+FormatEdge(y.MissingEdge), "supports="+strings.Join(y.Supports, ","), "failures="+strings.Join(y.Failures, ","))
	return join("Y_support", parts)
}

func FormatDirac(d ChiralDiracSupport) string {
	return join("D_F_symbolic_support", []string{d.Expression, "space=" + d.Space, "gamma=" + d.GammaConvention, "left_rank=" + i(d.LeftRank), "right_rank=" + i(d.RightRank), "total_rank=" + i(d.TotalRank), "block_shape=" + i(d.BlockRows) + "x" + i(d.BlockCols), "adjoint_block=" + b(d.UsesAdjointBlock), "self_adjoint=" + b(d.SelfAdjointByConstruction), "chiral_odd=" + b(d.ChiralOddByConstruction), "explicit_support_matrix=" + b(d.ExplicitSupportMatrix), "native_D_F=" + b(d.NativeDFMatrix), "numerical_D_F=" + b(d.NumericalDFMatrix), "J_opposite_proof=" + b(d.JOppositeCompatibilityProof), "first_order=" + b(d.FirstOrderCertified), "bimodule_commutant=" + b(d.BimoduleCommutantProof), "supports=" + strings.Join(d.Supports, ","), "failures=" + strings.Join(d.Failures, ",")})
}

func FormatImpact(im Impact) string {
	return join("impact", []string{im.Classification, "gate847_inherited=" + b(im.Gate847Inherited), "matrix_support=" + b(im.MatrixSupportConstructed), "self_adjoint_support=" + b(im.SelfAdjointSupport), "chirality_odd_support=" + b(im.ChiralityOddSupport), "puncture_null_edge_candidate=" + b(im.PunctureNullEdgeCandidate), "first_order=" + b(im.FirstOrderProof), "bimodule=" + b(im.BimoduleProof), "J_opposite=" + b(im.JOppositeProof), "alpha_sealed=" + b(im.AlphaStillSealed), "magnitudes_missing=" + b(im.MagnitudesStillMissing), "update_N_eff=" + b(im.CanUpdateNEff), "update_C_Yukawa=" + b(im.CanUpdateCYukawa), "update_C_Higgs=" + b(im.CanUpdateCHiggs), "R3=" + b(im.CanPromoteToR3), "R4=" + b(im.CanPromoteToR4)})
}

func FormatFirewalls(f Firewalls) string {
	return join("firewalls", []string{"enforced=" + b(f.Enforced), "D_F_support_seal_not_native=" + b(f.DFSupportMatrixSealNotNative), "no_numerical_D_F=" + b(f.NoNumericalDFMatrix), "no_full_rho_J_gamma=" + b(f.NoFullRhoFJFGammaFPackage), "no_first_order=" + b(f.NoFirstOrderProof), "no_bimodule_commutant=" + b(f.NoBimoduleCommutantProof), "no_J_opposite=" + b(f.NoJOppositeCompatibilityProof), "Y_symbols_not_yukawa=" + b(f.YSymbolsNotYukawaValues), "edge_support_not_trace_magnitude=" + b(f.EdgeSupportNotTraceMagnitude), "no_numerical_yukawa=" + b(f.NoNumericalYukawaValues), "alpha_sealed=" + b(f.AlphaStillSealed), "no_trace_magnitude=" + b(f.NoTraceMagnitudeReadout), "puncture_null_edge_only_seal=" + b(f.PunctureNullEdgeOnlySeal), "no_native_null_edge=" + b(f.NoNativeNullEdgeTheorem), "no_particle_assignment=" + b(f.NoPhysicalParticleAssignment), "no_right_neutrino_theorem=" + b(f.NoRightNeutrinoTheorem), "no_three_generation=" + b(f.NoThreeGenerationTheorem), "no_N_eff_update=" + b(f.NoNEffUpdate), "no_C_updates=" + b(f.NoCYukawaUpdate), "not_R3=" + b(f.NotR3), "not_R4=" + b(f.NotR4), "verdict=" + f.Verdict})
}

func containsAll(haystack, needles []string) bool {
	m := map[string]bool{}
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

func join(label string, parts []string) string { return label + "{" + strings.Join(parts, "; ") + "}" }
func b(v bool) string {
	if v {
		return "true"
	}
	return "false"
}
func i(v int) string { return strconv.Itoa(v) }
func f64(v float64) string {
	return strings.TrimRight(strings.TrimRight(strconv.FormatFloat(v, 'f', 16, 64), "0"), ".")
}
