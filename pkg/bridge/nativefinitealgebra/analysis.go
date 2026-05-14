// Package nativefinitealgebra implements Gate 236:
// Native Finite Algebra Derivation / Contact-Preserving Subalgebra Search.
//
// Gate 235 derived the 32-real-dimensional doubled carrier by complexifying the
// native real Cl(1,7) spinor, but it refused to import Connes' finite algebra
// C ⊕ H ⊕ M3(C). Gate 236 asks what can actually be recovered from native
// data: the temporal/spatial 1⊕3 split of the four Fock generators and the
// previously derived contact-preserving su(2)⊕u(1) Lie algebra.
//
// The result is deliberately conservative. The 1⊕3 split on the generator
// carrier supports a mode-level block commutant C ⊕ M3(C), which is the correct
// algebraic shape for a singlet plus color triplet bookkeeping. Complexification
// also gives a natural complex scalar summand for the u(1) preflight. However,
// the quaternionic H summand is not forced: a Lie algebra su(2) is not yet a
// derived left H-module action on the doubled spinor, and the full finite
// algebra representation/opposite action/order-one calculus remain missing.
// Therefore the Standard Model finite algebra is not derived at this gate.
package nativefinitealgebra

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/complexifiedhilbertspace"
	"github.com/bagherbal/asha-engine/pkg/spinor"
)

const (
	AuditID = "GATE236-NATIVE-FINITE-ALGEBRA-SUBALGEBRA-SEARCH"

	StatusSplitSupport      = "CONDITIONAL_SUPPORT_NATIVE_1PLUS3_SPLIT_PREFLIGHT"
	StatusCPlusM3Support    = "CONDITIONAL_SUPPORT_MODE_COMMUTANT_C_PLUS_M3C_PREFLIGHT"
	StatusU1ComplexSupport  = "CONDITIONAL_SUPPORT_U1_COMPLEX_SUMMAND_PREFLIGHT"
	StatusFailedH           = "FAILED_ROUTE_NATIVE_QUATERNIONIC_H_DERIVATION"
	StatusFailedConnes      = "FAILED_ROUTE_EXACT_CONNES_ALGEBRA_DERIVATION"
	StatusFailedOrderOne    = "FAILED_ROUTE_NATIVE_ALGEBRA_ORDER_ONE_READINESS"
	StatusFailedFullAlgebra = "FAILED_ROUTE_FULL_NATIVE_FINITE_ALGEBRA_DERIVATION"
)

type ModeRow struct {
	Index        int
	Name         string
	Kind         string
	Projector    string
	Multiplicity int
}

type SplitAudit struct {
	Carrier                       string
	GeneratorDimension            int
	ComplexifiedCarrierDimensionC int
	ComplexifiedCarrierDimensionR int
	LeptonLikeGeneratorCount      int
	ColorLikeGeneratorCount       int
	ModeLevelProjectionExists     bool
	ProjectionFormula             string
	ExtendsToFockBookkeeping      bool
	FullParticleSpeciesProjection bool
	ColorLeptonBookkeepingNative  bool
	Verdict                       string
}

type CommutantAudit struct {
	SearchSpace                       string
	NaiveFullEndDimensionC            int
	ModeProjectionCommutant           string
	ModeProjectionCommutantDimensionC int
	ColorMatrixAlgebraPreflight       bool
	ComplexSingletPreflight           bool
	M3CDerivedAsPhysicalColorGauge    bool
	LiftToFullExteriorRepresentation  bool
	MaximalAlgebraOnFullSC            bool
	AmbiguityRows                     []string
	Verdict                           string
}

type ContactIntegrationAudit struct {
	DerivedLieInput               string
	U1ComplexSummandPreflight     bool
	SU2LieAlgebraAvailable        bool
	SU2ToQuaternionHModuleDerived bool
	DoubletProjectionDerived      bool
	LeftQuaternionicActionDerived bool
	AssociativeClosureComputed    bool
	HGenerated                    bool
	Verdict                       string
}

type AlgebraVerdictAudit struct {
	ConnesAlgebraImported        bool
	CPlusM3Preflight             bool
	QuaternionicHDerived         bool
	ExactCPlusHPlusM3Derived     bool
	FaithfulRepresentationOnSC   bool
	OppositeAlgebraActionDerived bool
	OrderOneCalculusReady        bool
	MajoranaSieveReady           bool
	Verdict                      string
}

type FirewallAudit struct {
	ImportedConnesAlgebra      bool
	InsertedSMGaugeGroup       bool
	InsertedGaugeMatrices      bool
	InsertedYukawaOrMassData   bool
	BGapPromotedToMass         bool
	ClaimedOrderOne            bool
	ClaimedSMAlgebraDerivation bool
	FiniteCorePolluted         bool
	Verdict                    string
}

type Summary struct {
	SplitDerived          bool
	CPlusM3Preflight      bool
	U1ComplexPreflight    bool
	QuaternionicHDerived  bool
	ExactSMAlgebraDerived bool
	OrderOneReady         bool
	Status                string
	NextGate              string
	Comment               string
}

type Analysis struct {
	Previous       complexifiedhilbertspace.Analysis
	Modes          []ModeRow
	Split          SplitAudit
	Commutant      CommutantAudit
	Contact        ContactIntegrationAudit
	Algebra        AlgebraVerdictAudit
	Firewall       FirewallAudit
	Summary        Summary
	TruthStatement string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := complexifiedhilbertspace.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 235 predecessor: %w", err)
			return
		}
		f, err := spinor.NewCovariantPhaseFockSpace(4)
		if err != nil {
			defaultErr = fmt.Errorf("construct Fock space: %w", err)
			return
		}
		defaultA, defaultErr = Build(prev, f)
	})
	return defaultA, defaultErr
}

func Build(prev complexifiedhilbertspace.Analysis, f spinor.FockSpace) (Analysis, error) {
	if f.ModeCount() != 4 || f.StateCount() != 16 {
		return Analysis{}, fmt.Errorf("Gate 236 requires native four-mode 16-state Fock carrier, got modes=%d states=%d", f.ModeCount(), f.StateCount())
	}
	modes := buildModeRows(f)
	split := auditSplit(f)
	comm := auditCommutant(split)
	contact := auditContactIntegration(comm)
	alg := auditAlgebraVerdict(comm, contact)
	fw := auditFirewall()
	sum := summarize(split, comm, contact, alg)
	truth := buildTruth(split, comm, contact, alg)
	return Analysis{Previous: prev, Modes: modes, Split: split, Commutant: comm, Contact: contact, Algebra: alg, Firewall: fw, Summary: sum, TruthStatement: truth}, nil
}

func buildModeRows(f spinor.FockSpace) []ModeRow {
	out := make([]ModeRow, 0, len(f.Modes))
	for _, m := range f.Modes {
		projector := "P_color"
		if m.Kind == spinor.TemporalMode {
			projector = "P_lepton"
		}
		out = append(out, ModeRow{Index: m.Index, Name: m.Name, Kind: string(m.Kind), Projector: projector, Multiplicity: 1})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Index < out[j].Index })
	return out
}

func auditSplit(f spinor.FockSpace) SplitAudit {
	lepton := f.TemporalModeCount()
	color := f.SpatialModeCount()
	return SplitAudit{
		Carrier:                       "complexified native spinor S_C = Λ*(C·e0 ⊕ C^3_spatial)",
		GeneratorDimension:            f.ModeCount(),
		ComplexifiedCarrierDimensionC: f.StateCount(),
		ComplexifiedCarrierDimensionR: 2 * f.StateCount(),
		LeptonLikeGeneratorCount:      lepton,
		ColorLikeGeneratorCount:       color,
		ModeLevelProjectionExists:     lepton == 1 && color == 3,
		ProjectionFormula:             "W = C·e0 ⊕ C·{e1,e2,e3}; P_lepton(e0)=e0, P_color(ei)=ei for i=1,2,3",
		ExtendsToFockBookkeeping:      true,
		FullParticleSpeciesProjection: false,
		ColorLeptonBookkeepingNative:  lepton == 1 && color == 3,
		Verdict:                       StatusSplitSupport,
	}
}

func auditCommutant(split SplitAudit) CommutantAudit {
	rows := []string{
		"The 1⊕3 projector on the generator carrier W has block-preserving algebra End(C)⊕End(C³)=C⊕M3(C).",
		"This is a mode-level commutant/preflight, not yet a faithful Standard Model algebra on the full exterior spinor S_C.",
		"Without explicit lifted matrices on Λ*W, the maximal associative subalgebra of End(S_C) is not uniquely fixed by the split alone.",
		"The split can support color bookkeeping, but it does not by itself derive gluon curvature, the opposite algebra, or order-one calculus.",
	}
	return CommutantAudit{
		SearchSpace:                       "End_C(S_C), constrained first by the native 1⊕3 mode projection",
		NaiveFullEndDimensionC:            16 * 16,
		ModeProjectionCommutant:           "C ⊕ M3(C) on the four generator modes W",
		ModeProjectionCommutantDimensionC: 1 + 9,
		ColorMatrixAlgebraPreflight:       split.ModeLevelProjectionExists && split.ColorLikeGeneratorCount == 3,
		ComplexSingletPreflight:           split.ModeLevelProjectionExists && split.LeptonLikeGeneratorCount == 1,
		M3CDerivedAsPhysicalColorGauge:    false,
		LiftToFullExteriorRepresentation:  false,
		MaximalAlgebraOnFullSC:            false,
		AmbiguityRows:                     rows,
		Verdict:                           StatusCPlusM3Support,
	}
}

func auditContactIntegration(comm CommutantAudit) ContactIntegrationAudit {
	return ContactIntegrationAudit{
		DerivedLieInput:               "contact-preserving su(2)⊕u(1) from earlier centralizer/gauge gates",
		U1ComplexSummandPreflight:     comm.ComplexSingletPreflight,
		SU2LieAlgebraAvailable:        true,
		SU2ToQuaternionHModuleDerived: false,
		DoubletProjectionDerived:      false,
		LeftQuaternionicActionDerived: false,
		AssociativeClosureComputed:    false,
		HGenerated:                    false,
		Verdict:                       strings.Join([]string{StatusU1ComplexSupport, StatusFailedH}, ";"),
	}
}

func auditAlgebraVerdict(comm CommutantAudit, contact ContactIntegrationAudit) AlgebraVerdictAudit {
	exact := comm.ColorMatrixAlgebraPreflight && contact.HGenerated && contact.U1ComplexSummandPreflight && comm.LiftToFullExteriorRepresentation
	return AlgebraVerdictAudit{
		ConnesAlgebraImported:        false,
		CPlusM3Preflight:             comm.ColorMatrixAlgebraPreflight && comm.ComplexSingletPreflight,
		QuaternionicHDerived:         contact.HGenerated,
		ExactCPlusHPlusM3Derived:     exact,
		FaithfulRepresentationOnSC:   false,
		OppositeAlgebraActionDerived: false,
		OrderOneCalculusReady:        false,
		MajoranaSieveReady:           false,
		Verdict:                      strings.Join([]string{StatusFailedConnes, StatusFailedOrderOne, StatusFailedFullAlgebra}, ";"),
	}
}

func auditFirewall() FirewallAudit {
	return FirewallAudit{
		ImportedConnesAlgebra:      false,
		InsertedSMGaugeGroup:       false,
		InsertedGaugeMatrices:      false,
		InsertedYukawaOrMassData:   false,
		BGapPromotedToMass:         false,
		ClaimedOrderOne:            false,
		ClaimedSMAlgebraDerivation: false,
		FiniteCorePolluted:         false,
		Verdict:                    "FIREWALL_PRESERVED_NATIVE_ALGEBRA_SEARCH_ONLY",
	}
}

func summarize(split SplitAudit, comm CommutantAudit, contact ContactIntegrationAudit, alg AlgebraVerdictAudit) Summary {
	status := strings.Join([]string{StatusSplitSupport, StatusCPlusM3Support, StatusU1ComplexSupport, StatusFailedH, StatusFailedConnes, StatusFailedFullAlgebra}, ";")
	return Summary{
		SplitDerived:          split.ModeLevelProjectionExists && split.ColorLeptonBookkeepingNative,
		CPlusM3Preflight:      alg.CPlusM3Preflight,
		U1ComplexPreflight:    contact.U1ComplexSummandPreflight,
		QuaternionicHDerived:  alg.QuaternionicHDerived,
		ExactSMAlgebraDerived: alg.ExactCPlusHPlusM3Derived,
		OrderOneReady:         alg.OrderOneCalculusReady,
		Status:                status,
		NextGate:              "derive explicit contact-preserving su(2) matrices on S_C and test whether their associative closure supplies a quaternionic left module",
		Comment:               "Gate 236 derives a native C⊕M3(C) mode-level preflight from the 1⊕3 split, but it does not derive H or the exact Connes algebra on S_C.",
	}
}

func buildTruth(split SplitAudit, comm CommutantAudit, contact ContactIntegrationAudit, alg AlgebraVerdictAudit) string {
	return fmt.Sprintf("The native generator split %d⊕%d supports a mode-level commutant %s of complex dimension %d, giving C⊕M3(C) preflight. The u(1) complex summand is plausible, but su(2) has not been promoted to a quaternionic H-module; exact C⊕H⊕M3(C), faithful S_C representation, opposite action, and order-one readiness remain un-derived.", split.LeptonLikeGeneratorCount, split.ColorLikeGeneratorCount, comm.ModeProjectionCommutant, comm.ModeProjectionCommutantDimensionC)
}
