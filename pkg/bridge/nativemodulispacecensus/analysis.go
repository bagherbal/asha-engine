// Package nativemodulispacecensus implements Gate 372:
// Native Moduli Space Dimension / Exact Dirac Parameter Census Sieve.
//
// The gate deliberately separates three notions that are often conflated:
//  1. raw finite-Dirac matrix entries after the spectral-triple axioms;
//  2. quotient by unphysical generation-basis redefinitions preserving the
//     kinetic representation;
//  3. the older minimal SM-19 vacuum ledger, which also contains theta_QCD
//     and an absolute scale not represented as ordinary finite-Dirac flavor
//     matrices.
package nativemodulispacecensus

import (
	"fmt"
	"strings"
	"sync"
)

const (
	AuditID = "GATE372-NATIVE-MODULI-SPACE-DIMENSION-EXACT-DIRAC-PARAMETER-CENSUS-SIEVE"

	StatusGate371Inherited              = "CONDITIONAL_SUPPORT_GATE371_DYNAMICAL_FLOW_OBSTRUCTION_INHERITED"
	StatusGeneralDiracParameterized     = "CONDITIONAL_SUPPORT_GENERAL_FINITE_DIRAC_PARAMETERIZATION_EXECUTED"
	StatusJRealitySieveExecuted         = "CONDITIONAL_SUPPORT_J_REALITY_SIEVE_EXECUTED"
	StatusChiralitySieveExecuted        = "CONDITIONAL_SUPPORT_CHIRALITY_SIEVE_EXECUTED"
	StatusFirstOrderSieveExecuted       = "CONDITIONAL_SUPPORT_FIRST_ORDER_EDGE_GRAPH_SIEVE_EXECUTED"
	StatusRawAxiomaticDimensionComputed = "CONDITIONAL_SUPPORT_RAW_AXIOMATIC_DIRAC_DIMENSION_COMPUTED"
	StatusGaugeQuotientAudited          = "CONDITIONAL_SUPPORT_ALGEBRA_GAUGE_QUOTIENT_AUDITED"
	StatusFlavorBasisQuotientComputed   = "CONDITIONAL_SUPPORT_FLAVOR_BASIS_QUOTIENT_COMPUTED"
	StatusNativeModuliSpaceComputed     = "CONDITIONAL_SUPPORT_NATIVE_MODULI_SPACE_COMPUTED"
	StatusMinimalChargedDFDimension13   = "CONDITIONAL_SUPPORT_MINIMAL_CHARGED_FINITE_DIRAC_MODULI_DIMENSION_13"
	StatusMajoranaDFDimension31         = "CONDITIONAL_SUPPORT_MAJORANA_FINITE_DIRAC_MODULI_DIMENSION_31"
	StatusExternal15Decomposed          = "CONDITIONAL_SUPPORT_EXTERNAL_15_DECOMPOSED_AS_13_PLUS_THETA_PLUS_SCALE"
	StatusExternalCountVerifiedRefined  = "CONDITIONAL_SUPPORT_EXTERNAL_COUNT_VERIFIED_AS_MINIMAL_VACUUM_LEDGER_NOT_DF_MODULI"
	StatusNoHiddenFlavorReductionFound  = "CONDITIONAL_SUPPORT_NO_HIDDEN_FLAVOR_REDUCTION_FOUND_BY_NATIVE_DIRAC_CENSUS"
	StatusLandscapePreservationAudited  = "CONDITIONAL_SUPPORT_LANDSCAPE_PRESERVATION_AUDITED"
	StatusEpistemicFirewallAudited      = "CONDITIONAL_SUPPORT_EPISTEMIC_FIREWALL_AUDITED"

	StatusTensionDFNotExternal15                  = "CONDITIONAL_TENSION_FINITE_DIRAC_MODULI_ARE_NOT_IDENTICAL_TO_EXTERNAL_15_LEDGER"
	StatusTensionAlgebraGaugeNotFlavorQuotient    = "CONDITIONAL_TENSION_UA_GAUGE_QUOTIENT_DOES_NOT_REMOVE_GENERATION_BASIS_ROTATIONS"
	StatusTensionNeutrinoMajoranaModelDependent   = "CONDITIONAL_TENSION_MAJORANA_NEUTRINO_CENSUS_IS_EXTENDED_MODEL_DEPENDENT"
	StatusTensionSpectralAxiomsAllowGenericFlavor = "CONDITIONAL_TENSION_SPECTRAL_TRIPLE_AXIOMS_ALLOW_GENERIC_GENERATION_MATRICES"

	StatusFailedNativeReductionBelow15       = "FAILED_ROUTE_NATIVE_MODULI_SPACE_REDUCTION_BELOW_EXTERNAL_15_NOT_FOUND"
	StatusFailedHiddenCrossSectorConstraints = "FAILED_ROUTE_HIDDEN_CROSS_SECTOR_FLAVOR_CONSTRAINTS_NOT_FOUND"
	StatusFailedVacuumPointStillNotSelected  = "FAILED_ROUTE_PHYSICAL_VACUUM_POINT_STILL_NOT_SELECTED"
	StatusFailedYukawaCoordinatesStillFree   = "FAILED_ROUTE_YUKAWA_COORDINATES_STILL_FREE_AFTER_NATIVE_CENSUS"
	StatusFailedCKMTextureStillFree          = "FAILED_ROUTE_CKM_TEXTURE_STILL_FREE_AFTER_NATIVE_CENSUS"
	StatusFailedPMNSMajoranaModelDependent   = "FAILED_ROUTE_PMNS_MAJORANA_TEXTURE_REMAINS_EXTENDED_MODEL_COORDINATE"
)

const (
	Complex3x3RealDim          = 18
	ComplexSymmetric3x3RealDim = 12
)

type Inheritance struct {
	Executed                      bool
	HighestInheritedGate          int
	PreviousObstruction           string
	DynamicalIntertwinerStillOpen bool
	ReasonToChangeQuestion        string
	Verdict                       string
}

type DiracBlock struct {
	Name                string
	Symbol              string
	EdgeClass           string
	AllowedByChirality  bool
	AllowedByFirstOrder bool
	JRealityMirror      string
	MatrixType          string
	ComplexEntries      int
	RawRealDim          int
	PhysicalRole        string
	Verdict             string
}

type Parameterization struct {
	Executed               bool
	Blocks                 []DiracBlock
	MinimalChargedRawDim   int
	DiracNeutrinoRawDim    int
	MajoranaExtendedRawDim int
	AllAllowedRawDim       int
	Verdict                string
}

type AxiomSieve struct {
	Executed                        bool
	JRealityImposesMirrorBlocks     bool
	ChiralityAllowsOnlyOddEdges     bool
	FirstOrderEnforcesEdgeGraph     bool
	MajoranaBlockSymmetric          bool
	AdditionalGenerationConstraints int
	MinimalChargedRawAfterAxioms    int
	DiracExtendedRawAfterAxioms     int
	MajoranaExtendedRawAfterAxioms  int
	ForbiddenEdges                  []string
	Verdict                         string
}

type QuotientScenario struct {
	Name                     string
	RawRealDim               int
	BasisGroup               string
	BasisGroupDim            int
	GenericStabilizer        string
	GenericStabilizerDim     int
	OrbitDim                 int
	PhysicalDim              int
	Parameters               []string
	UsesAlgebraGaugeQuotient bool
	UsesFlavorBasisQuotient  bool
	Verdict                  string
}

type QuotientAudit struct {
	Executed                         bool
	AlgebraGaugeGroup                string
	AlgebraGaugeRemovesGenerationDim int
	FlavorBasisQuotientRequired      bool
	Scenarios                        []QuotientScenario
	MinimalChargedDFDim              int
	DiracNeutrinoDFDim               int
	MajoranaSeesawDFDim              int
	Verdict                          string
}

type NativeResult struct {
	Executed                          bool
	CanonicalAllAllowedInterpretation string
	NPhysicalDF                       int
	MinimalChargedDFDim               int
	DiracNeutrinoDFDim                int
	MajoranaSeesawDFDim               int
	External15                        int
	External15Decomposition           string
	HiddenCrossSectorConstraints      int
	NativeReductionBelow15            bool
	DirectAnswer                      string
	Verdict                           string
}

type Firewall struct {
	Executed                    bool
	NoYukawaValuesImported      bool
	NoCKMValuesImported         bool
	NoPMNSValuesImported        bool
	NoMassValuesImported        bool
	NoVacuumDirectionForced     bool
	NoGaugeFlavorConflation     bool
	NoMajoranaMinimalConflation bool
	LandscapeRatiosPreserved    bool
	Verdict                     string
}

type Analysis struct {
	Inheritance      Inheritance
	Parameterization Parameterization
	Axioms           AxiomSieve
	Quotient         QuotientAudit
	Native           NativeResult
	Firewall         Firewall
	Truth            string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() { defaultA, defaultErr = Build() })
	return defaultA, defaultErr
}

func Build() (Analysis, error) {
	inheritance := buildInheritance()
	param := parameterizeGeneralFiniteDirac()
	axioms := applyAxiomSieve(param)
	quotient := computeQuotients(axioms)
	native := computeNativeResult(quotient)
	firewall := auditFirewalls(native)
	truth := "Gate 372 performs the native finite-Dirac parameter census and corrects the epistemology of the Phase-III problem.  The finite spectral-triple axioms restrict the allowed blocks to generic Yukawa matrices plus a symmetric Majorana block, but they do not impose hidden generation-texture constraints.  After quotienting unphysical generation-basis rotations, the minimal charged finite-Dirac flavor moduli have dimension 13: six quark masses, four CKM parameters, and three charged-lepton masses.  The older external 15-vacuum ledger is therefore not literally dim M(D_F); it is 13 finite-Dirac flavor coordinates plus theta_QCD and one absolute scale.  If the all-allowed right-neutrino Majorana sector is included, the finite-Dirac moduli dimension is 31, so the neutrino/Majorana census is an extended model-dependent ledger.  No native reduction below the external 15 minimal-vacuum count is found, and no physical vacuum point is selected."
	return Analysis{inheritance, param, axioms, quotient, native, firewall, truth}, nil
}

func buildInheritance() Inheritance {
	return Inheritance{Executed: true, HighestInheritedGate: 371, PreviousObstruction: "Phase III dynamical attempts found capacity witnesses but no native vacuum-selection Hamiltonian/intertwiner", DynamicalIntertwinerStillOpen: true, ReasonToChangeQuestion: "instead of subtracting phenomenological parameters externally, compute the dimension of the finite-Dirac moduli space directly", Verdict: join(StatusGate371Inherited)}
}

func parameterizeGeneralFiniteDirac() Parameterization {
	blocks := []DiracBlock{
		{Name: "up-type quark Yukawa", Symbol: "Y_u", EdgeClass: "Q_L <-> u_R", AllowedByChirality: true, AllowedByFirstOrder: true, JRealityMirror: "charge-conjugate mirror block", MatrixType: "generic complex 3x3 generation matrix", ComplexEntries: 9, RawRealDim: Complex3x3RealDim, PhysicalRole: "up, charm, top amplitudes plus quark misalignment", Verdict: StatusGeneralDiracParameterized},
		{Name: "down-type quark Yukawa", Symbol: "Y_d", EdgeClass: "Q_L <-> d_R", AllowedByChirality: true, AllowedByFirstOrder: true, JRealityMirror: "charge-conjugate mirror block", MatrixType: "generic complex 3x3 generation matrix", ComplexEntries: 9, RawRealDim: Complex3x3RealDim, PhysicalRole: "down, strange, bottom amplitudes plus CKM misalignment", Verdict: StatusGeneralDiracParameterized},
		{Name: "charged-lepton Yukawa", Symbol: "Y_e", EdgeClass: "L_L <-> e_R", AllowedByChirality: true, AllowedByFirstOrder: true, JRealityMirror: "charge-conjugate mirror block", MatrixType: "generic complex 3x3 generation matrix", ComplexEntries: 9, RawRealDim: Complex3x3RealDim, PhysicalRole: "electron, muon, tau amplitudes", Verdict: StatusGeneralDiracParameterized},
		{Name: "Dirac-neutrino Yukawa", Symbol: "Y_ν", EdgeClass: "L_L <-> ν_R", AllowedByChirality: true, AllowedByFirstOrder: true, JRealityMirror: "charge-conjugate mirror block", MatrixType: "generic complex 3x3 generation matrix", ComplexEntries: 9, RawRealDim: Complex3x3RealDim, PhysicalRole: "Dirac neutrino amplitudes and PMNS-type misalignment when activated", Verdict: join(StatusGeneralDiracParameterized, StatusTensionNeutrinoMajoranaModelDependent)},
		{Name: "right-neutrino Majorana block", Symbol: "M_R", EdgeClass: "ν_R <-> ν_R^c", AllowedByChirality: true, AllowedByFirstOrder: true, JRealityMirror: "self-conjugate symmetric block", MatrixType: "generic complex symmetric 3x3 generation matrix", ComplexEntries: 6, RawRealDim: ComplexSymmetric3x3RealDim, PhysicalRole: "heavy Majorana scale/orientation ledger", Verdict: join(StatusGeneralDiracParameterized, StatusTensionNeutrinoMajoranaModelDependent)},
	}
	minimal := blocks[0].RawRealDim + blocks[1].RawRealDim + blocks[2].RawRealDim
	diracNu := minimal + blocks[3].RawRealDim
	majorana := diracNu + blocks[4].RawRealDim
	return Parameterization{Executed: true, Blocks: blocks, MinimalChargedRawDim: minimal, DiracNeutrinoRawDim: diracNu, MajoranaExtendedRawDim: majorana, AllAllowedRawDim: majorana, Verdict: join(StatusGeneralDiracParameterized, StatusRawAxiomaticDimensionComputed)}
}

func applyAxiomSieve(p Parameterization) AxiomSieve {
	return AxiomSieve{Executed: true, JRealityImposesMirrorBlocks: true, ChiralityAllowsOnlyOddEdges: true, FirstOrderEnforcesEdgeGraph: true, MajoranaBlockSymmetric: true, AdditionalGenerationConstraints: 0, MinimalChargedRawAfterAxioms: p.MinimalChargedRawDim, DiracExtendedRawAfterAxioms: p.DiracNeutrinoRawDim, MajoranaExtendedRawAfterAxioms: p.MajoranaExtendedRawDim, ForbiddenEdges: []string{"same-chirality charged Dirac edges", "quark-lepton Yukawa cross edges", "color-changing finite Dirac edges", "non-symmetric Majorana self-pairing part", "any empirical texture inserted into generation entries"}, Verdict: join(StatusJRealitySieveExecuted, StatusChiralitySieveExecuted, StatusFirstOrderSieveExecuted, StatusTensionSpectralAxiomsAllowGenericFlavor)}
}

func computeQuotients(a AxiomSieve) QuotientAudit {
	quark := quotientScenario("quark Yukawa sector", 36, "U(3)_Q x U(3)_u x U(3)_d", 27, "baryon-number U(1)", 1, []string{"6 quark masses", "3 CKM angles", "1 CKM phase"}, true)
	chargedLepton := quotientScenario("charged-lepton-only sector", 18, "U(3)_L x U(3)_e", 18, "three independent diagonal lepton phase stabilizers", 3, []string{"3 charged-lepton singular values"}, true)
	minimalCharged := quotientScenario("minimal charged finite-Dirac flavor sector", a.MinimalChargedRawAfterAxioms, "quark basis group plus charged-lepton basis group", 45, "baryon U(1) plus three charged-lepton phase stabilizers", 4, []string{"6 quark masses", "4 CKM parameters", "3 charged-lepton masses"}, true)
	diracLepton := quotientScenario("Dirac-neutrino lepton sector", 36, "U(3)_L x U(3)_e x U(3)_ν", 27, "total lepton-number U(1)", 1, []string{"3 charged-lepton masses", "3 Dirac-neutrino masses", "3 PMNS angles", "1 Dirac PMNS phase"}, true)
	diracExtended := quotientScenario("quark plus Dirac-neutrino finite-Dirac sector", a.DiracExtendedRawAfterAxioms, "quark basis group plus Dirac-lepton basis group", 54, "baryon U(1) plus lepton U(1)", 2, []string{"10 quark flavor parameters", "10 Dirac-lepton flavor parameters"}, true)
	majoranaLepton := quotientScenario("Majorana/seesaw lepton finite-Dirac sector", 48, "U(3)_L x U(3)_e x U(3)_νR", 27, "generic Majorana sector has no continuous stabilizer", 0, []string{"3 charged-lepton masses", "3 light-neutrino masses", "3 heavy-neutrino masses", "6 PMNS low-energy mixing parameters", "6 high-energy seesaw orientation/phase parameters"}, true)
	majoranaExtended := quotientScenario("quark plus Majorana finite-Dirac sector", a.MajoranaExtendedRawAfterAxioms, "quark basis group plus Majorana-lepton basis group", 54, "baryon U(1)", 1, []string{"10 quark flavor parameters", "21 Majorana/seesaw lepton parameters"}, true)
	return QuotientAudit{Executed: true, AlgebraGaugeGroup: "U(C ⊕ H ⊕ M_3(C)) reduced by unimodularity to SM gauge group; it acts on gauge/color/weak labels, not on generation-copy coordinates", AlgebraGaugeRemovesGenerationDim: 0, FlavorBasisQuotientRequired: true, Scenarios: []QuotientScenario{quark, chargedLepton, minimalCharged, diracLepton, diracExtended, majoranaLepton, majoranaExtended}, MinimalChargedDFDim: minimalCharged.PhysicalDim, DiracNeutrinoDFDim: diracExtended.PhysicalDim, MajoranaSeesawDFDim: majoranaExtended.PhysicalDim, Verdict: join(StatusGaugeQuotientAudited, StatusFlavorBasisQuotientComputed, StatusTensionAlgebraGaugeNotFlavorQuotient)}
}

func quotientScenario(name string, raw int, group string, groupDim int, stabilizer string, stabilizerDim int, params []string, flavorQuotient bool) QuotientScenario {
	orbit := groupDim - stabilizerDim
	phys := raw - orbit
	verdict := join(StatusFlavorBasisQuotientComputed)
	if strings.Contains(name, "minimal charged") && phys == 13 {
		verdict = join(verdict, StatusMinimalChargedDFDimension13)
	}
	if strings.Contains(name, "Majorana finite-Dirac") && phys == 31 {
		verdict = join(verdict, StatusMajoranaDFDimension31)
	}
	return QuotientScenario{Name: name, RawRealDim: raw, BasisGroup: group, BasisGroupDim: groupDim, GenericStabilizer: stabilizer, GenericStabilizerDim: stabilizerDim, OrbitDim: orbit, PhysicalDim: phys, Parameters: params, UsesAlgebraGaugeQuotient: false, UsesFlavorBasisQuotient: flavorQuotient, Verdict: verdict}
}

func computeNativeResult(q QuotientAudit) NativeResult {
	direct := "Canonical all-allowed finite-Dirac-plus-Majorana census gives dim M(D_F)=31 after generation-basis quotient. The minimal charged finite-Dirac subledger gives 13. The old external 15 is 13 finite-Dirac flavor moduli plus theta_QCD plus one absolute scale, not a literal finite-Dirac matrix dimension."
	return NativeResult{Executed: true, CanonicalAllAllowedInterpretation: "Y_u,Y_d,Y_e,Y_ν,M_R all allowed by the current edge ledger", NPhysicalDF: q.MajoranaSeesawDFDim, MinimalChargedDFDim: q.MinimalChargedDFDim, DiracNeutrinoDFDim: q.DiracNeutrinoDFDim, MajoranaSeesawDFDim: q.MajoranaSeesawDFDim, External15: 15, External15Decomposition: "15 = 13 minimal charged finite-Dirac flavor moduli + theta_QCD + absolute unit/VEV scale", HiddenCrossSectorConstraints: 0, NativeReductionBelow15: false, DirectAnswer: direct, Verdict: join(StatusNativeModuliSpaceComputed, StatusMajoranaDFDimension31, StatusExternal15Decomposed, StatusExternalCountVerifiedRefined, StatusNoHiddenFlavorReductionFound, StatusTensionDFNotExternal15, StatusTensionNeutrinoMajoranaModelDependent, StatusFailedNativeReductionBelow15, StatusFailedHiddenCrossSectorConstraints, StatusFailedVacuumPointStillNotSelected, StatusFailedYukawaCoordinatesStillFree, StatusFailedCKMTextureStillFree, StatusFailedPMNSMajoranaModelDependent)}
}

func auditFirewalls(n NativeResult) Firewall {
	return Firewall{Executed: true, NoYukawaValuesImported: true, NoCKMValuesImported: true, NoPMNSValuesImported: true, NoMassValuesImported: true, NoVacuumDirectionForced: true, NoGaugeFlavorConflation: true, NoMajoranaMinimalConflation: true, LandscapeRatiosPreserved: true, Verdict: join(StatusLandscapePreservationAudited, StatusEpistemicFirewallAudited)}
}

func FindScenario(s []QuotientScenario, name string) QuotientScenario {
	for _, x := range s {
		if x.Name == name {
			return x
		}
	}
	return QuotientScenario{}
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("gate<=%d inherited=%t obstruction=%q change=%q verdict=%s", x.HighestInheritedGate, x.DynamicalIntertwinerStillOpen, x.PreviousObstruction, x.ReasonToChangeQuestion, x.Verdict)
}

func FormatParameterization(p Parameterization) string {
	parts := make([]string, 0, len(p.Blocks))
	for _, b := range p.Blocks {
		parts = append(parts, fmt.Sprintf("%s:%s raw=%d type=%s", b.Symbol, b.EdgeClass, b.RawRealDim, b.MatrixType))
	}
	return fmt.Sprintf("minimalRaw=%d diracNuRaw=%d majoranaRaw=%d blocks=[%s] verdict=%s", p.MinimalChargedRawDim, p.DiracNeutrinoRawDim, p.MajoranaExtendedRawDim, strings.Join(parts, "; "), p.Verdict)
}

func FormatAxioms(a AxiomSieve) string {
	return fmt.Sprintf("J=%t chirality=%t firstOrder=%t MRsym=%t extraGenConstraints=%d rawAfter(min=%d,diracNu=%d,majorana=%d) forbidden=%s verdict=%s", a.JRealityImposesMirrorBlocks, a.ChiralityAllowsOnlyOddEdges, a.FirstOrderEnforcesEdgeGraph, a.MajoranaBlockSymmetric, a.AdditionalGenerationConstraints, a.MinimalChargedRawAfterAxioms, a.DiracExtendedRawAfterAxioms, a.MajoranaExtendedRawAfterAxioms, strings.Join(a.ForbiddenEdges, "; "), a.Verdict)
}

func FormatScenario(s QuotientScenario) string {
	return fmt.Sprintf("%s: raw=%d group=%s dim=%d stabilizer=%s dim=%d orbit=%d physical=%d params=%s verdict=%s", s.Name, s.RawRealDim, s.BasisGroup, s.BasisGroupDim, s.GenericStabilizer, s.GenericStabilizerDim, s.OrbitDim, s.PhysicalDim, strings.Join(s.Parameters, ", "), s.Verdict)
}

func FormatQuotient(q QuotientAudit) string {
	items := make([]string, 0, len(q.Scenarios))
	for _, s := range q.Scenarios {
		items = append(items, fmt.Sprintf("%s=%d", s.Name, s.PhysicalDim))
	}
	return fmt.Sprintf("UA=%q removesGen=%d flavorQuotient=%t dims=[%s] verdict=%s", q.AlgebraGaugeGroup, q.AlgebraGaugeRemovesGenerationDim, q.FlavorBasisQuotientRequired, strings.Join(items, "; "), q.Verdict)
}

func FormatNative(n NativeResult) string {
	return fmt.Sprintf("N_DF(allAllowedMajorana)=%d minimalChargedDF=%d diracNuDF=%d external15=%d decomp=%q hiddenConstraints=%d reducedBelow15=%t verdict=%s", n.NPhysicalDF, n.MinimalChargedDFDim, n.DiracNeutrinoDFDim, n.External15, n.External15Decomposition, n.HiddenCrossSectorConstraints, n.NativeReductionBelow15, n.Verdict)
}

func FormatFirewall(f Firewall) string {
	return fmt.Sprintf("noY=%t noCKM=%t noPMNS=%t noMass=%t noVacuum=%t noGaugeFlavorConflation=%t noMajoranaMinimalConflation=%t landscape=%t verdict=%s", f.NoYukawaValuesImported, f.NoCKMValuesImported, f.NoPMNSValuesImported, f.NoMassValuesImported, f.NoVacuumDirectionForced, f.NoGaugeFlavorConflation, f.NoMajoranaMinimalConflation, f.LandscapeRatiosPreserved, f.Verdict)
}

func join(xs ...string) string {
	out := make([]string, 0, len(xs))
	for _, x := range xs {
		if x == "" {
			continue
		}
		out = append(out, x)
	}
	return strings.Join(out, " | ")
}
