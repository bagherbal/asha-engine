// Package generation2k7plusu2higgssocketandquaternioniccommutantaudit implements
// Gate 711: K7+ U(2) Higgs Socket and Quaternionic Commutant Audit.
//
// Gate 710 certified that the Hodge-positive K7+ sector carries an internal
// quaternionic complex-structure triple. Gate 711 audits the next airlock:
// after a choice J_H=J_n, K7+ admits an internal u(2,J_H) socket whose real
// Lie algebra is a U(1) phase line plus a commuting Sp(1)/SU(2)-like socket.
// It preserves the firewall that this internal socket is not yet the physical
// electroweak SU(2)_L x U(1)_Y Higgs representation, hypercharge assignment,
// scalar runtime theorem, Yukawa theorem, or native 7/72 theorem.
package generation2k7plusu2higgssocketandquaternioniccommutantaudit

import (
	"fmt"
	"strings"
	"sync"

	gate710 "github.com/bagherbal/asha-engine/pkg/bridge/generation2k7plusquaternioniccomplexstructureandhiggsdoubletairlockaudit"
)

const (
	AuditID = "GATE711-K7-PLUS-U2-HIGGS-SOCKET-AND-QUATERNIONIC-COMMUTANT-AUDIT"

	StatusGate710QuaternionicK7PlusInherited           = "PASS_GATE710_QUATERNIONIC_K7_PLUS_INHERITED"
	StatusSO4SplitAudited                              = "PASS_SO4_SPLIT_AUDITED"
	StatusQuaternionicCommutantComputed                = "PASS_QUATERNIONIC_COMMUTANT_COMPUTED"
	StatusChosenComplexStructureJHAudited              = "PASS_CHOSEN_COMPLEX_STRUCTURE_JH_AUDITED"
	StatusU2SocketDefinedAfterJHChoice                 = "PASS_U2_SOCKET_DEFINED_AFTER_JH_CHOICE"
	StatusRelationToK7MinusSelectorRecorded            = "PASS_RELATION_TO_K7_MINUS_SELECTOR_RECORDED"
	StatusPhysicalElectroweakFirewallEnforced          = "PASS_PHYSICAL_ELECTROWEAK_FIREWALL_ENFORCED"
	StatusK7PlusInternalU2HiggsSocketAfterChoice       = "CONDITIONAL_SUPPORT_K7_PLUS_HAS_INTERNAL_U2_HIGGS_SOCKET_AFTER_COMPLEX_STRUCTURE_CHOICE"
	StatusCommutantSP1InternalSU2SocketCandidate       = "CONDITIONAL_SUPPORT_COMMUTANT_SP1_SUPPLIES_INTERNAL_SU2_SOCKET_CANDIDATE"
	StatusSpanJHSuppliesInternalU1PhaseSocketCandidate = "CONDITIONAL_SUPPORT_SPAN_JH_SUPPLIES_INTERNAL_U1_PHASE_SOCKET_CANDIDATE"
	StatusK7MinusDirectionCanSelectJHCandidate         = "CONDITIONAL_SUPPORT_K7_MINUS_DIRECTION_CAN_SELECT_JH_CANDIDATE"
	StatusNoCanonicalJHSelected                        = "FAILED_ROUTE_NO_CANONICAL_JH_SELECTED"
	StatusInternalU2SocketNotPhysicalSU2LU1Y           = "FAILED_ROUTE_INTERNAL_U2_SOCKET_NOT_CERTIFIED_AS_PHYSICAL_SU2L_U1Y"
	StatusNoHyperchargeAssignmentOrNormalization       = "FAILED_ROUTE_NO_HYPERCHARGE_ASSIGNMENT_OR_NORMALIZATION"
	StatusNoTypedK7PlusToPhysicalHiggsDoubletMap       = "FAILED_ROUTE_NO_TYPED_K7_PLUS_TO_PHYSICAL_HIGGS_DOUBLET_MAP"
	StatusNoYukawaOperatorOrEigenvalueTheorem          = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusNoHiggsMassOrScalarRuntimeTheorem            = "FAILED_ROUTE_NO_HIGGS_MASS_OR_SCALAR_RUNTIME_THEOREM"
	StatusGate711K7PlusU2HiggsSocketBoundary           = "FIREWALL_PRESERVED_GATE711_K7_PLUS_U2_HIGGS_SOCKET_BOUNDARY"
)

const (
	k7PlusRealDimension = 4
	so4Dimension        = 6
	sp1Dimension        = 3
	u2Dimension         = 4
	u1Dimension         = 1
	complexDimension    = 2
	k7MinusDimension    = 3
)

type Gate710Inheritance struct {
	QuaternionicK7PlusInherited bool
	K7PlusRealDimension         int
	HasQuaternionicTriple       bool
	HasS2ComplexFamily          bool
	C2PreCarrierAfterChoice     bool
	InternalSU2LikeCandidate    bool
	CanonicalJHSelected         bool
	PhysicalSU2LCertified       bool
	HyperchargeCertified        bool
	PhysicalHiggsDoubletMap     bool
	YukawaOperatorCertified     bool
	HiggsMassCertified          bool
	Verdict                     string
}

type SO4SplitAudit struct {
	Algebra             string
	Dimension           int
	Split               string
	LeftSP1Dimension    int
	RightSP1Dimension   int
	QuaternionicFactor  string
	TripleSelectsFactor bool
	PhysicalGaugeGroup  bool
	Verdict             string
}

type QuaternionicCommutantAudit struct {
	Definition            string
	Dimension             int
	ClosesAsSU2Like       bool
	Commutator            string
	NormalizationRequired bool
	PhysicalSU2LCertified bool
	Verdict               string
}

type ChosenComplexStructureAudit struct {
	Selector                 string
	UnitCondition            string
	JHSquared                string
	ComplexDimension         int
	CanonicalSelected        bool
	SelectedAfterChoice      bool
	PotentialK7MinusSelector bool
	Verdict                  string
}

type U2SocketAudit struct {
	Definition                    string
	Dimension                     int
	Decomposition                 string
	U1Line                        string
	SU2Socket                     string
	SpanJHInternalU1Candidate     bool
	CommutantInternalSU2Candidate bool
	PhysicalElectroweakU2         bool
	Verdict                       string
}

type K7MinusSelectorRelationAudit struct {
	FrameMap                 string
	UnitDirectionRoute       string
	K7MinusDimension         int
	CanSelectJH              bool
	NativeSelectorTheorem    bool
	GenerationTheorem        bool
	FlavorOrientationTheorem bool
	Verdict                  string
}

type PhysicalElectroweakFirewallAudit struct {
	ClaimsInternalU2PhysicalElectroweak bool
	ClaimsCommutantPhysicalSU2L         bool
	ClaimsSpanJHPhysicalHypercharge     bool
	ClaimsHyperchargeNormalization      bool
	ClaimsTypedHiggsDoubletMap          bool
	ClaimsYukawaOperator                bool
	ClaimsYukawaEigenvalues             bool
	ClaimsHiggsMass                     bool
	ClaimsScalarRuntime                 bool
	Verdict                             string
}

type MissingMaps struct {
	ThetaSU2 string
	ThetaY   string
	ThetaH   string
	ThetaJH  string
	Missing  []string
	Verdict  string
}

type SourceTypeClassification struct {
	SO4Role       string
	CommutantRole string
	JHRole        string
	U2SocketRole  string
	K7MinusRole   string
	FirewallRole  string
	Verdict       string
}

type Analysis struct {
	Inherited   Gate710Inheritance
	SO4         SO4SplitAudit
	Commutant   QuaternionicCommutantAudit
	ChosenJH    ChosenComplexStructureAudit
	U2Socket    U2SocketAudit
	K7Minus     K7MinusSelectorRelationAudit
	Firewalls   PhysicalElectroweakFirewallAudit
	Missing     MissingMaps
	SourceTypes SourceTypeClassification
	Truth       string
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
	g710, err := gate710.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate710 inheritance unavailable: %w", err)
	}
	inherited := buildGate710Inheritance(g710)
	so4 := buildSO4Split(inherited)
	commutant := buildCommutant(so4)
	chosen := buildChosenJH(inherited)
	socket := buildU2Socket(chosen, commutant)
	k7minus := buildK7MinusSelector(g710)
	firewalls := PhysicalElectroweakFirewallAudit{Verdict: StatusGate711K7PlusU2HiggsSocketBoundary}
	missing := buildMissing()
	sourceTypes := buildSourceTypes()
	truth := "Gate 711 audits the representation socket opened by Gate710.  The real four-dimensional K7+ sector has so(4) dimension six, with an internal sp(1)+sp(1) split; the inherited quaternionic triple occupies one Sp(1)-like factor, while its so(4) commutant supplies a three-dimensional internal SU(2)-socket candidate.  After choosing one complex structure J_H=J_n, the complex-linear orthogonal algebra u(2,J_H) is four-dimensional and decomposes as span{J_H} plus the quaternionic commutant.  This is only an internal U(2)-compatible socket: no canonical J_H, physical SU(2)_L embedding, hypercharge assignment/normalization, typed Higgs doublet map, Yukawa operator/eigenvalue theorem, Higgs mass theorem, scalar runtime theorem, or native 7/72 theorem is certified."
	return Analysis{Inherited: inherited, SO4: so4, Commutant: commutant, ChosenJH: chosen, U2Socket: socket, K7Minus: k7minus, Firewalls: firewalls, Missing: missing, SourceTypes: sourceTypes, Truth: truth}, nil
}

func buildGate710Inheritance(g gate710.Analysis) Gate710Inheritance {
	return Gate710Inheritance{
		QuaternionicK7PlusInherited: g.Inherited.RepresentationAirlockInherited && g.FanoTriple.QuaternionicTriple && g.Endomorphisms.QuaternionicProductCertified,
		K7PlusRealDimension:         g.HiggsCompatibility.K7PlusRealDimension,
		HasQuaternionicTriple:       g.FanoTriple.QuaternionicTriple && g.Endomorphisms.SquaresToMinusIdentity,
		HasS2ComplexFamily:          g.ComplexFamily.S2Family,
		C2PreCarrierAfterChoice:     g.ComplexFamily.C2AfterChoice && g.HiggsCompatibility.CandidateComplexDimension == complexDimension,
		InternalSU2LikeCandidate:    g.SU2LikeAction.Sp1SU2LikeAlgebra,
		CanonicalJHSelected:         g.ComplexFamily.CanonicalSelected,
		PhysicalSU2LCertified:       g.SU2LikeAction.PhysicalSU2LCertified || g.SU2LikeAction.ElectroweakEmbedding,
		HyperchargeCertified:        g.SU2LikeAction.HyperchargeAssignment,
		PhysicalHiggsDoubletMap:     g.HiggsCompatibility.PhysicalHiggsDoubletMap,
		YukawaOperatorCertified:     g.FanoRelation.YukawaOperatorCertified,
		HiggsMassCertified:          g.HiggsCompatibility.ScalarRuntimeTheorem,
		Verdict:                     StatusGate710QuaternionicK7PlusInherited,
	}
}

func buildSO4Split(i Gate710Inheritance) SO4SplitAudit {
	return SO4SplitAudit{
		Algebra:             "so(K7+,g_+)",
		Dimension:           so4Dimension,
		Split:               "so(4) ≅ sp(1)_A ⊕ sp(1)_B",
		LeftSP1Dimension:    sp1Dimension,
		RightSP1Dimension:   sp1Dimension,
		QuaternionicFactor:  "span{J_1,J_2,J_3} occupies one Sp(1)-like factor",
		TripleSelectsFactor: i.HasQuaternionicTriple && i.K7PlusRealDimension == k7PlusRealDimension,
		PhysicalGaugeGroup:  false,
		Verdict: strings.Join([]string{
			StatusSO4SplitAudited,
			"CONDITIONAL_SUPPORT_QUATERNIONIC_TRIPLE_SELECTS_ONE_SP1_FACTOR",
		}, "; "),
	}
}

func buildCommutant(s SO4SplitAudit) QuaternionicCommutantAudit {
	ok := s.TripleSelectsFactor && s.Dimension == so4Dimension && s.LeftSP1Dimension == sp1Dimension && s.RightSP1Dimension == sp1Dimension
	return QuaternionicCommutantAudit{
		Definition:            "Comm_so4(J_1,J_2,J_3)={X in so(4): [X,J_a]=0 for all a}",
		Dimension:             sp1Dimension,
		ClosesAsSU2Like:       ok,
		Commutator:            "[X_i,X_j]=2 epsilon_ijk X_k after basis normalization",
		NormalizationRequired: true,
		PhysicalSU2LCertified: false,
		Verdict: strings.Join([]string{
			StatusQuaternionicCommutantComputed,
			StatusCommutantSP1InternalSU2SocketCandidate,
			StatusInternalU2SocketNotPhysicalSU2LU1Y,
		}, "; "),
	}
}

func buildChosenJH(i Gate710Inheritance) ChosenComplexStructureAudit {
	return ChosenComplexStructureAudit{
		Selector:                 "J_H=J_n=n_1J_1+n_2J_2+n_3J_3",
		UnitCondition:            "n_1^2+n_2^2+n_3^2=1",
		JHSquared:                "J_H^2=-I",
		ComplexDimension:         complexDimension,
		CanonicalSelected:        false,
		SelectedAfterChoice:      i.HasS2ComplexFamily && i.C2PreCarrierAfterChoice,
		PotentialK7MinusSelector: true,
		Verdict: strings.Join([]string{
			StatusChosenComplexStructureJHAudited,
			StatusNoCanonicalJHSelected,
			StatusK7MinusDirectionCanSelectJHCandidate,
		}, "; "),
	}
}

func buildU2Socket(j ChosenComplexStructureAudit, c QuaternionicCommutantAudit) U2SocketAudit {
	ok := j.SelectedAfterChoice && c.ClosesAsSU2Like && c.Dimension == sp1Dimension
	return U2SocketAudit{
		Definition:                    "u(2,J_H)={X in so(4): [X,J_H]=0}",
		Dimension:                     u2Dimension,
		Decomposition:                 "u(2,J_H)=span{J_H} ⊕ Comm_so4(J_1,J_2,J_3)",
		U1Line:                        "span{J_H}",
		SU2Socket:                     "Comm_so4(J_1,J_2,J_3)",
		SpanJHInternalU1Candidate:     ok,
		CommutantInternalSU2Candidate: ok,
		PhysicalElectroweakU2:         false,
		Verdict: strings.Join([]string{
			StatusU2SocketDefinedAfterJHChoice,
			StatusK7PlusInternalU2HiggsSocketAfterChoice,
			StatusCommutantSP1InternalSU2SocketCandidate,
			StatusSpanJHSuppliesInternalU1PhaseSocketCandidate,
			StatusInternalU2SocketNotPhysicalSU2LU1Y,
		}, "; "),
	}
}

func buildK7MinusSelector(g gate710.Analysis) K7MinusSelectorRelationAudit {
	canSelect := g.FanoRelation.K7MinusIndexesTriple && g.FanoRelation.K7MinusChannels == k7MinusDimension && g.FanoRelation.TwoFormsOnK7Plus
	return K7MinusSelectorRelationAudit{
		FrameMap:                 "F_A:K7- -> Lambda^2(K7+)^*, eta_a -> omega_a -> J_a",
		UnitDirectionRoute:       "unit n in K7- selects J_H=n_a J_a",
		K7MinusDimension:         k7MinusDimension,
		CanSelectJH:              canSelect,
		NativeSelectorTheorem:    false,
		GenerationTheorem:        false,
		FlavorOrientationTheorem: false,
		Verdict: strings.Join([]string{
			StatusRelationToK7MinusSelectorRecorded,
			StatusK7MinusDirectionCanSelectJHCandidate,
			StatusNoCanonicalJHSelected,
		}, "; "),
	}
}

func buildMissing() MissingMaps {
	missing := []string{
		StatusNoCanonicalJHSelected,
		StatusInternalU2SocketNotPhysicalSU2LU1Y,
		StatusNoHyperchargeAssignmentOrNormalization,
		StatusNoTypedK7PlusToPhysicalHiggsDoubletMap,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusNoHiggsMassOrScalarRuntimeTheorem,
	}
	return MissingMaps{
		ThetaSU2: "Theta_SU2: internal commutant sp(1) -> already-derived electroweak SU(2)_L action",
		ThetaY:   "Theta_Y: span{J_H} -> U(1)_Y hypercharge with correct Higgs charge/normalization",
		ThetaH:   "Theta_H: K7+_J -> physical Higgs doublet representation",
		ThetaJH:  "Theta_JH: native selector for the physical complex structure J_H",
		Missing:  missing,
		Verdict:  strings.Join(missing, "; "),
	}
}

func buildSourceTypes() SourceTypeClassification {
	return SourceTypeClassification{
		SO4Role:       "real K7+ rotation algebra so(4), internally split as two Sp(1)-like factors",
		CommutantRole: "three-dimensional quaternionic commutant; internal SU(2)-socket candidate only",
		JHRole:        "chosen complex structure J_H supplies a U(1)-phase socket candidate but is not canonical",
		U2SocketRole:  "u(2,J_H) internal representation socket after a complex-structure choice, not physical electroweak U(2)",
		K7MinusRole:   "unit direction in K7- can select J_H through the Fano frame; no selector/flavor theorem yet",
		FirewallRole:  "socket remains internal until electroweak embedding, hypercharge normalization, and Higgs representation maps are typed",
		Verdict: strings.Join([]string{
			StatusK7PlusInternalU2HiggsSocketAfterChoice,
			StatusCommutantSP1InternalSU2SocketCandidate,
			StatusSpanJHSuppliesInternalU1PhaseSocketCandidate,
			StatusGate711K7PlusU2HiggsSocketBoundary,
		}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate710QuaternionicK7PlusInherited,
		StatusSO4SplitAudited,
		StatusQuaternionicCommutantComputed,
		StatusChosenComplexStructureJHAudited,
		StatusU2SocketDefinedAfterJHChoice,
		StatusRelationToK7MinusSelectorRecorded,
		StatusPhysicalElectroweakFirewallEnforced,
		StatusK7PlusInternalU2HiggsSocketAfterChoice,
		StatusCommutantSP1InternalSU2SocketCandidate,
		StatusSpanJHSuppliesInternalU1PhaseSocketCandidate,
		StatusK7MinusDirectionCanSelectJHCandidate,
		StatusNoCanonicalJHSelected,
		StatusInternalU2SocketNotPhysicalSU2LU1Y,
		StatusNoHyperchargeAssignmentOrNormalization,
		StatusNoTypedK7PlusToPhysicalHiggsDoubletMap,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusNoHiggsMassOrScalarRuntimeTheorem,
		StatusGate711K7PlusU2HiggsSocketBoundary,
	}
}

func FormatInherited(x Gate710Inheritance) string {
	return fmt.Sprintf("quatInherited=%t plusDim=%d quatTriple=%t s2=%t c2=%t su2like=%t canonical=%t physicalSU2L=%t hypercharge=%t physicalHiggs=%t yukawa=%t higgsMass=%t verdict=%q", x.QuaternionicK7PlusInherited, x.K7PlusRealDimension, x.HasQuaternionicTriple, x.HasS2ComplexFamily, x.C2PreCarrierAfterChoice, x.InternalSU2LikeCandidate, x.CanonicalJHSelected, x.PhysicalSU2LCertified, x.HyperchargeCertified, x.PhysicalHiggsDoubletMap, x.YukawaOperatorCertified, x.HiggsMassCertified, x.Verdict)
}

func FormatSO4(x SO4SplitAudit) string {
	return fmt.Sprintf("alg=%q dim=%d split=%q left=%d right=%d factor=%q selects=%t physical=%t verdict=%q", x.Algebra, x.Dimension, x.Split, x.LeftSP1Dimension, x.RightSP1Dimension, x.QuaternionicFactor, x.TripleSelectsFactor, x.PhysicalGaugeGroup, x.Verdict)
}

func FormatCommutant(x QuaternionicCommutantAudit) string {
	return fmt.Sprintf("def=%q dim=%d closes=%t comm=%q norm=%t physicalSU2L=%t verdict=%q", x.Definition, x.Dimension, x.ClosesAsSU2Like, x.Commutator, x.NormalizationRequired, x.PhysicalSU2LCertified, x.Verdict)
}

func FormatChosenJH(x ChosenComplexStructureAudit) string {
	return fmt.Sprintf("selector=%q unit=%q squared=%q complexDim=%d canonical=%t afterChoice=%t k7minus=%t verdict=%q", x.Selector, x.UnitCondition, x.JHSquared, x.ComplexDimension, x.CanonicalSelected, x.SelectedAfterChoice, x.PotentialK7MinusSelector, x.Verdict)
}

func FormatU2(x U2SocketAudit) string {
	return fmt.Sprintf("def=%q dim=%d decomposition=%q u1=%q su2=%q spanJH=%t commSU2=%t physicalU2=%t verdict=%q", x.Definition, x.Dimension, x.Decomposition, x.U1Line, x.SU2Socket, x.SpanJHInternalU1Candidate, x.CommutantInternalSU2Candidate, x.PhysicalElectroweakU2, x.Verdict)
}

func FormatK7Minus(x K7MinusSelectorRelationAudit) string {
	return fmt.Sprintf("frame=%q route=%q dim=%d canSelect=%t nativeSelector=%t generation=%t flavor=%t verdict=%q", x.FrameMap, x.UnitDirectionRoute, x.K7MinusDimension, x.CanSelectJH, x.NativeSelectorTheorem, x.GenerationTheorem, x.FlavorOrientationTheorem, x.Verdict)
}

func FormatFirewalls(x PhysicalElectroweakFirewallAudit) string {
	return fmt.Sprintf("physicalU2=%t physicalSU2L=%t hypercharge=%t hyperNorm=%t higgsMap=%t yukawa=%t eig=%t higgsMass=%t runtime=%t verdict=%q", x.ClaimsInternalU2PhysicalElectroweak, x.ClaimsCommutantPhysicalSU2L, x.ClaimsSpanJHPhysicalHypercharge, x.ClaimsHyperchargeNormalization, x.ClaimsTypedHiggsDoubletMap, x.ClaimsYukawaOperator, x.ClaimsYukawaEigenvalues, x.ClaimsHiggsMass, x.ClaimsScalarRuntime, x.Verdict)
}

func FormatMissing(x MissingMaps) string {
	return fmt.Sprintf("thetaSU2=%q thetaY=%q thetaH=%q thetaJH=%q missing=%d verdict=%q", x.ThetaSU2, x.ThetaY, x.ThetaH, x.ThetaJH, len(x.Missing), x.Verdict)
}

func FormatSourceTypes(x SourceTypeClassification) string {
	return fmt.Sprintf("so4=%q commutant=%q jh=%q u2=%q k7minus=%q firewall=%q verdict=%q", x.SO4Role, x.CommutantRole, x.JHRole, x.U2SocketRole, x.K7MinusRole, x.FirewallRole, x.Verdict)
}
