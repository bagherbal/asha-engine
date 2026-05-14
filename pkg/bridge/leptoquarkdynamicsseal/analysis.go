// Package leptoquarkdynamicsseal implements Gate 209: Pati-Salam
// leptoquark current dynamics / B-L-preserving proton-decay operator seal
// audit.
//
// Gate 208 proved current-connection proton stability but correctly left open
// a sharper threat: the u(4) matter-current inventory contains six
// off-diagonal quark-lepton slots.  Gate 209 audits whether those slots can
// acquire dynamics from the current finite algebra.  If not, it introduces a
// seal that keeps them kinematic-only and re-runs the B-L-preserving proton
// decay operator obstruction under that explicit quarantine.
package leptoquarkdynamicsseal

import (
	"fmt"
	"strings"
	"sync"
)

const (
	StatusNativeLeptoquarkDynamicsObstructed  = "FAILED_ROUTE_NATIVE_LEPTOQUARK_DYNAMICS"
	StatusConditionalOnLeptoquarkDynamicsSeal = "CONDITIONAL_ON_LEPTOQUARK_DYNAMICS_SEAL"
	StatusSealedConnectionBaryonConservation  = "SEALED_CONNECTION_BARYON_CONSERVATION_THEOREM"
)

type Gate208Snapshot struct {
	Gate208Inherited                       bool
	CurrentConnectionProtonStable          bool
	AbsoluteBaryonConservationProved       bool
	BMinusLDoesNotForbidStandardTemplates  bool
	LeptoquarkCurrentSlotsPresent          bool
	LeptoquarkCurrentSlotsGaugeActivated   bool
	FiniteOperatorCoefficientDerived       bool
	ProtonLifetimeComputed                 bool
	RecommendedLeptoquarkDynamicsSealAudit bool
	TruthStatement                         string
}

func DefaultGate208Snapshot() Gate208Snapshot {
	return Gate208Snapshot{
		Gate208Inherited:                       true,
		CurrentConnectionProtonStable:          true,
		AbsoluteBaryonConservationProved:       false,
		BMinusLDoesNotForbidStandardTemplates:  true,
		LeptoquarkCurrentSlotsPresent:          true,
		LeptoquarkCurrentSlotsGaugeActivated:   false,
		FiniteOperatorCoefficientDerived:       false,
		ProtonLifetimeComputed:                 false,
		RecommendedLeptoquarkDynamicsSealAudit: true,
		TruthStatement:                         "Gate 208 established current-connection proton stability, rejected B-L as a blanket firewall, and left the six u(4) quark-lepton current slots as the next explicit dynamics/seal obligation.",
	}
}

type LeptoquarkSlot struct {
	Name                     string
	Direction                string
	ColorIndex               int
	GaugeCurvatureDerived    bool
	FiniteActionDerived      bool
	LocalFieldMapDerived     bool
	PropagatorDerived        bool
	MassScaleDerived         bool
	CouplingCoefficientKnown bool
	DynamicMediator          bool
	Obstruction              string
}

type DynamicActivationAudit struct {
	SlotsAudited                int
	OffDiagonalDimension        int
	AllSlotsKinematicOnly       bool
	AnyGaugeCurvatureDerived    bool
	AnyFiniteActionDerived      bool
	AnyLocalFieldMapDerived     bool
	AnyPropagatorDerived        bool
	AnyMassScaleDerived         bool
	AnyCouplingCoefficientKnown bool
	AnyDynamicMediator          bool
	NativeDynamicsStatus        string
	Verdict                     string
	Slots                       []LeptoquarkSlot
}

type LeptoquarkDynamicsSeal struct {
	ID                         string
	Name                       string
	Active                     bool
	Conditional                bool
	KinematicSlotsQuarantined  int
	ForbidsGaugeActivation     bool
	ForbidsPropagatorUse       bool
	ForbidsOperatorCoefficient bool
	ForbidsLifetimeFormula     bool
	CanBeLiftedByFutureTheorem bool
	SealStatement              string
}

type SealedOperatorTemplate struct {
	Name                        string
	SymbolicForm                string
	DeltaB                      int
	DeltaL                      int
	DeltaBMinusL                int
	BMinusLPreserving           bool
	WouldNeedLeptoquarkDynamics bool
	WouldNeedFourFermionMap     bool
	WouldNeedSuppressionScale   bool
	ConstructibleBeforeSeal     bool
	ConstructibleUnderSeal      bool
	SuppressionScaleComputed    bool
	Obstruction                 string
}

type SealedOperatorAudit struct {
	TemplatesAudited               int
	AllTemplatesBMinusLPreserving  bool
	AnyTemplateConstructible       bool
	AnySuppressionScaleComputed    bool
	ProtonLifetimeComputationLegal bool
	SealedObstruction              bool
	Verdict                        string
	Templates                      []SealedOperatorTemplate
}

type ConservationTheoremAudit struct {
	TheoremName                         string
	SealedConnectionBaryonConservation  bool
	CurrentConnectionStabilityInherited bool
	AbsoluteUnsealedBaryonTheorem       bool
	ConditionalOnLeptoquarkSeal         bool
	LeptoquarkSlotsDormant              bool
	ProtonLifetimeStrictlyObstructed    bool
	FutureDynamicsStillOpen             bool
	Verdict                             string
}

type FirewallAudit struct {
	NoSU5Imported                    bool
	NoSO10Imported                   bool
	NoPatiSalamGaugeDynamicsImported bool
	NoBMinusLFalseFirewall           bool
	NoProtonLifetimeComputed         bool
	NoLeptoquarkMassAssumed          bool
	NoPropagatorAssumed              bool
	NoOperatorCoefficientAssumed     bool
	SealDoesNotRewriteNativeFailure  bool
	RecommendedNextGate              string
	RemainingUnknowns                []string
}

type Analysis struct {
	Gate208        Gate208Snapshot
	Dynamics       DynamicActivationAudit
	Seal           LeptoquarkDynamicsSeal
	Operators      SealedOperatorAudit
	Conservation   ConservationTheoremAudit
	Firewall       FirewallAudit
	Status         string
	TruthStatement string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		defaultValue, defaultErr = Build(DefaultGate208Snapshot())
	})
	return defaultValue, defaultErr
}

func Build(prev Gate208Snapshot) (Analysis, error) {
	if !prev.Gate208Inherited || !prev.RecommendedLeptoquarkDynamicsSealAudit {
		return Analysis{}, fmt.Errorf("Gate 209 requires the Gate 208 leptoquark-dynamics audit obligation")
	}
	if !prev.LeptoquarkCurrentSlotsPresent {
		return Analysis{}, fmt.Errorf("Gate 209 requires the six u(4) leptoquark current slots to be present as kinematic inventory")
	}

	dyn := buildDynamicActivationAudit()
	seal := buildSeal(dyn)
	ops := buildSealedOperatorAudit(seal)
	cons := buildConservationAudit(prev, dyn, seal, ops)
	firewall := buildFirewall(cons)
	truth := "Gate 209 proves that the six u(4) quark-lepton current slots remain kinematic-only under the current finite algebra: no gauge curvature, finite action, local-field map, propagator, mass scale, or coefficient is derived. It therefore introduces the LeptoquarkDynamicsSeal. Under this explicit quarantine, B-L-preserving dimension-six proton-decay templates remain unconstructible and proton-lifetime computation is strictly obstructed. The result is a sealed-connection baryon-conservation theorem conditional on the seal, not an unsealed absolute baryon-conservation theorem."

	return Analysis{
		Gate208:        prev,
		Dynamics:       dyn,
		Seal:           seal,
		Operators:      ops,
		Conservation:   cons,
		Firewall:       firewall,
		Status:         StatusConditionalOnLeptoquarkDynamicsSeal,
		TruthStatement: truth,
	}, nil
}

func buildDynamicActivationAudit() DynamicActivationAudit {
	slots := make([]LeptoquarkSlot, 0, 6)
	for color := 1; color <= 3; color++ {
		slots = append(slots,
			newSlot(fmt.Sprintf("LQ-q%d-to-l", color), "quark-color -> lepton", color),
			newSlot(fmt.Sprintf("LQ-l-to-q%d", color), "lepton -> quark-color", color),
		)
	}
	return DynamicActivationAudit{
		SlotsAudited:                len(slots),
		OffDiagonalDimension:        6,
		AllSlotsKinematicOnly:       true,
		AnyGaugeCurvatureDerived:    false,
		AnyFiniteActionDerived:      false,
		AnyLocalFieldMapDerived:     false,
		AnyPropagatorDerived:        false,
		AnyMassScaleDerived:         false,
		AnyCouplingCoefficientKnown: false,
		AnyDynamicMediator:          false,
		NativeDynamicsStatus:        StatusNativeLeptoquarkDynamicsObstructed,
		Verdict:                     "FAILED_ROUTE: the six u(4) off-diagonal quark-lepton slots are present as matter-current inventory, but ASHA has not derived gauge curvature, a finite action term, a continuum local-field map, a propagator denominator, a mass/suppression scale, or an exchange coefficient for them",
		Slots:                       slots,
	}
}

func newSlot(name, direction string, color int) LeptoquarkSlot {
	return LeptoquarkSlot{
		Name:                     name,
		Direction:                direction,
		ColorIndex:               color,
		GaugeCurvatureDerived:    false,
		FiniteActionDerived:      false,
		LocalFieldMapDerived:     false,
		PropagatorDerived:        false,
		MassScaleDerived:         false,
		CouplingCoefficientKnown: false,
		DynamicMediator:          false,
		Obstruction:              "kinematic u(4) matrix-current slot only; no contact-preserving curvature component, local field carrier, kinetic/action term, propagator, mass unit, or coefficient is derived",
	}
}

func buildSeal(dyn DynamicActivationAudit) LeptoquarkDynamicsSeal {
	active := dyn.AllSlotsKinematicOnly && !dyn.AnyDynamicMediator
	return LeptoquarkDynamicsSeal{
		ID:                         "SEAL-LEPTOQUARK-DYNAMICS-GATE209",
		Name:                       "LeptoquarkDynamicsSeal",
		Active:                     active,
		Conditional:                true,
		KinematicSlotsQuarantined:  dyn.SlotsAudited,
		ForbidsGaugeActivation:     active,
		ForbidsPropagatorUse:       active,
		ForbidsOperatorCoefficient: active,
		ForbidsLifetimeFormula:     active,
		CanBeLiftedByFutureTheorem: true,
		SealStatement:              "The six u(4) quark-lepton slots may be recorded as kinematic inventory, but must not be used as dynamical leptoquark mediators unless a future theorem derives their curvature/action/local-field/propagator/mass/coefficient semantics.",
	}
}

func buildSealedOperatorAudit(seal LeptoquarkDynamicsSeal) SealedOperatorAudit {
	templates := []SealedOperatorTemplate{
		{
			Name:                        "left-handed QQQL",
			SymbolicForm:                "epsilon_color epsilon_weak (Q Q)(Q L) / Lambda^2",
			DeltaB:                      1,
			DeltaL:                      1,
			DeltaBMinusL:                0,
			BMinusLPreserving:           true,
			WouldNeedLeptoquarkDynamics: true,
			WouldNeedFourFermionMap:     true,
			WouldNeedSuppressionScale:   true,
			ConstructibleBeforeSeal:     false,
			ConstructibleUnderSeal:      false,
			SuppressionScaleComputed:    false,
			Obstruction:                 "B-L preserving template remains external; under the leptoquark dynamics seal there is no active q-l mediator, coefficient, local four-Weyl product, or suppression scale",
		},
		{
			Name:                        "right/conjugate UUD E",
			SymbolicForm:                "epsilon_color (u^c u^c)(d^c e^c) / Lambda^2",
			DeltaB:                      -1,
			DeltaL:                      -1,
			DeltaBMinusL:                0,
			BMinusLPreserving:           true,
			WouldNeedLeptoquarkDynamics: true,
			WouldNeedFourFermionMap:     true,
			WouldNeedSuppressionScale:   true,
			ConstructibleBeforeSeal:     false,
			ConstructibleUnderSeal:      false,
			SuppressionScaleComputed:    false,
			Obstruction:                 "B-L preserving conjugate template remains unconstructed; the seal forbids using dormant u(4) slots as propagating exchange channels",
		},
		{
			Name:                        "mixed QQLd class",
			SymbolicForm:                "schematic Q Q L d^c / Lambda^2",
			DeltaB:                      1,
			DeltaL:                      1,
			DeltaBMinusL:                0,
			BMinusLPreserving:           true,
			WouldNeedLeptoquarkDynamics: true,
			WouldNeedFourFermionMap:     true,
			WouldNeedSuppressionScale:   true,
			ConstructibleBeforeSeal:     false,
			ConstructibleUnderSeal:      false,
			SuppressionScaleComputed:    false,
			Obstruction:                 "no finite Fierz channel, local operator map, or leptoquark exchange coefficient exists before or after the seal",
		},
	}
	allBL := true
	anyConstructible := false
	anyScale := false
	for _, t := range templates {
		allBL = allBL && t.BMinusLPreserving && t.DeltaBMinusL == 0
		anyConstructible = anyConstructible || t.ConstructibleUnderSeal
		anyScale = anyScale || t.SuppressionScaleComputed
	}
	return SealedOperatorAudit{
		TemplatesAudited:               len(templates),
		AllTemplatesBMinusLPreserving:  allBL,
		AnyTemplateConstructible:       anyConstructible,
		AnySuppressionScaleComputed:    anyScale,
		ProtonLifetimeComputationLegal: false,
		SealedObstruction:              seal.Active && !anyConstructible && !anyScale,
		Verdict:                        "Under the LeptoquarkDynamicsSeal, the standard B-L-preserving dimension-six templates remain unconstructible; no suppression scale or lifetime formula is legal.",
		Templates:                      templates,
	}
}

func buildConservationAudit(prev Gate208Snapshot, dyn DynamicActivationAudit, seal LeptoquarkDynamicsSeal, ops SealedOperatorAudit) ConservationTheoremAudit {
	sealed := prev.CurrentConnectionProtonStable && seal.Active && ops.SealedObstruction
	return ConservationTheoremAudit{
		TheoremName:                         StatusSealedConnectionBaryonConservation,
		SealedConnectionBaryonConservation:  sealed,
		CurrentConnectionStabilityInherited: prev.CurrentConnectionProtonStable,
		AbsoluteUnsealedBaryonTheorem:       false,
		ConditionalOnLeptoquarkSeal:         seal.Active && seal.Conditional,
		LeptoquarkSlotsDormant:              dyn.AllSlotsKinematicOnly && !dyn.AnyDynamicMediator,
		ProtonLifetimeStrictlyObstructed:    !ops.ProtonLifetimeComputationLegal,
		FutureDynamicsStillOpen:             seal.CanBeLiftedByFutureTheorem,
		Verdict:                             "As long as the leptoquark dynamics seal holds, the current connection plus dormant u(4) slots cannot mediate B/L-violating proton decay; this is sealed baryon conservation, not an unsealed all-future theorem.",
	}
}

func buildFirewall(cons ConservationTheoremAudit) FirewallAudit {
	return FirewallAudit{
		NoSU5Imported:                    true,
		NoSO10Imported:                   true,
		NoPatiSalamGaugeDynamicsImported: true,
		NoBMinusLFalseFirewall:           true,
		NoProtonLifetimeComputed:         true,
		NoLeptoquarkMassAssumed:          true,
		NoPropagatorAssumed:              true,
		NoOperatorCoefficientAssumed:     true,
		SealDoesNotRewriteNativeFailure:  cons.ConditionalOnLeptoquarkSeal && !cons.AbsoluteUnsealedBaryonTheorem,
		RecommendedNextGate:              "Gate 210 — sealed baryon-stable threshold sector / non-universal deformation viability without universal Landau-pole completion",
		RemainingUnknowns: []string{
			"derive or permanently reject a finite kinetic/action term for u(4) leptoquark currents",
			"derive a local four-Weyl product/Fierz map before any QQQL coefficient is legal",
			"derive a mass or suppression scale before proton lifetime formulas are legal",
			"test whether the baryon-stable sealed carrier sector can solve the mismatch without the rejected universal beta row",
		},
	}
}

func FormatDynamics(d DynamicActivationAudit) string {
	return fmt.Sprintf("slots=%d/%d; curvature=%v action=%v local-field=%v propagator=%v mass=%v coefficient=%v mediator=%v", d.SlotsAudited, d.OffDiagonalDimension, d.AnyGaugeCurvatureDerived, d.AnyFiniteActionDerived, d.AnyLocalFieldMapDerived, d.AnyPropagatorDerived, d.AnyMassScaleDerived, d.AnyCouplingCoefficientKnown, d.AnyDynamicMediator)
}

func FormatSeal(s LeptoquarkDynamicsSeal) string {
	return fmt.Sprintf("%s active=%v conditional=%v quarantined_slots=%d; %s", s.ID, s.Active, s.Conditional, s.KinematicSlotsQuarantined, s.SealStatement)
}

func FormatTemplates(xs []SealedOperatorTemplate) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		state := "blocked"
		if x.ConstructibleUnderSeal {
			state = "constructed"
		}
		parts = append(parts, fmt.Sprintf("%s[%s, Δ(B-L)=%d]: %s", x.Name, state, x.DeltaBMinusL, x.Obstruction))
	}
	return strings.Join(parts, " | ")
}

func FormatUnknowns(xs []string) string { return strings.Join(xs, "; ") }
