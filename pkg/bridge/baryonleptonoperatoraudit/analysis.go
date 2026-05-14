// Package baryonleptonoperatoraudit implements Gate 208: baryon/lepton
// violating operator basis audit / proton-decay channel construction
// obstruction.
//
// Gate 207 showed that the low sealed boundary scale cannot be interpreted with
// naive SU(5) proton-decay formulas unless ASHA actually derives X/Y-like
// leptoquark gauge bosons or equivalent local dimension-six B/L-violating
// operators.  Gate 208 therefore audits only the engine-native inventory: the
// contact-preserving gauge seed, the Fock/u(4) matter-current inventory, and
// the scalar-bundle integration functional.  It does not import SU(5), SO(10),
// proton-lifetime formulas, or observed proton-decay bounds.
package baryonleptonoperatoraudit

import (
	"fmt"
	"strings"
	"sync"
)

const (
	StatusProtonDecayChannelConstructionObstructed = "FAILED_ROUTE_PROTON_DECAY_CHANNEL_CONSTRUCTION"
)

type Gate207Snapshot struct {
	Gate207Inherited                  bool
	UniversalCompletionFailed         bool
	CarrierSealStillConditional       bool
	LowBoundaryScaleWarning           bool
	NaiveSU5ProtonDecayWarning        bool
	EngineXYMediatorsDerived          bool
	EngineDimensionSixOperatorDerived bool
	RecommendedOperatorBasisAudit     bool
	FiniteCoreProtonLifetimeClaimed   bool
	TruthStatement                    string
}

func DefaultGate207Snapshot() Gate207Snapshot {
	return Gate207Snapshot{
		Gate207Inherited:                  true,
		UniversalCompletionFailed:         true,
		CarrierSealStillConditional:       true,
		LowBoundaryScaleWarning:           true,
		NaiveSU5ProtonDecayWarning:        true,
		EngineXYMediatorsDerived:          false,
		EngineDimensionSixOperatorDerived: false,
		RecommendedOperatorBasisAudit:     true,
		FiniteCoreProtonLifetimeClaimed:   false,
		TruthStatement:                    "Gate 207 rejected the external universal beta completion, retained a proton-decay warning for naive unified theories, and required a native B/L-violating operator audit before any proton-lifetime formula is legal.",
	}
}

type MatterCurrentInventory struct {
	FockModes                         int
	U4Dimension                       int
	CentralU1Dimension                int
	ColorSU3Dimension                 int
	BMinusLDimension                  int
	LeptoquarkOffDiagonalDimension    int
	DecompositionComplete             bool
	ContainsQuarkLeptonCurrentSlots   bool
	LeptoquarkSlotsGaugeActivated     bool
	ContactConnectionAlgebra          string
	ContactConnectionHasColor         bool
	ContactConnectionHasLeptoquark    bool
	FullSU5OrSO10GaugeConnection      bool
	BMinusLConservationAloneForbidsPD bool
	Verdict                           string
}

type OperatorTemplate struct {
	Name                          string
	SymbolicForm                  string
	DeltaB                        int
	DeltaL                        int
	DeltaBMinusL                  int
	SMGaugeInvariantTemplate      bool
	RequiresThreeQuarkClosure     bool
	RequiresQuarkLeptonMediator   bool
	RequiresLocalContinuumFields  bool
	RequiresBLCouplingCoefficient bool
	SupportedByContactConnection  bool
	SupportedByU4CurrentAction    bool
	SupportedByTauEta             bool
	ConstructedByFiniteAlgebra    bool
	Obstruction                   string
}

type OperatorSearchAudit struct {
	TemplatesAudited               int
	SMGaugeInvariantTemplatesExist bool
	AnyTemplateConstructed         bool
	AnyContactMediatedChannel      bool
	AnyU4MediatedChannel           bool
	AnyTauEtaSupportedChannel      bool
	SuppressionScaleComputed       bool
	SuppressionScaleExpression     string
	Verdict                        string
	Templates                      []OperatorTemplate
}

type ConservationAudit struct {
	CurrentConnectionProtonStable           bool
	ExactBaryonConservationProved           bool
	ExactLeptonConservationProved           bool
	BMinusLConservedByTemplates             bool
	BMinusLForbidsQQQL                      bool
	ColorTrialityForbidsQQQL                bool
	TrialityOrContactInvariantForbidsAll    bool
	LeptoquarkInventoryPreventsAbsoluteNoGo bool
	ProtonDecayChannelConstructionFailed    bool
	TheoremName                             string
	Verdict                                 string
}

type FirewallAudit struct {
	NoSU5Imported                        bool
	NoSO10Imported                       bool
	NoObservedLifetimeBoundsUsedForProof bool
	NoProtonLifetimeComputed             bool
	NoXYPresumed                         bool
	NoLeptoquarkGaugeActivationPresumed  bool
	NoBaryonConservationOverclaimed      bool
	CurrentConnectionOnly                bool
	FuturePatiSalamDynamicsStillOpen     bool
	RecommendedNextGate                  string
	RemainingUnknowns                    []string
}

type Analysis struct {
	Gate207        Gate207Snapshot
	Inventory      MatterCurrentInventory
	OperatorSearch OperatorSearchAudit
	Conservation   ConservationAudit
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
		defaultValue, defaultErr = Build(DefaultGate207Snapshot())
	})
	return defaultValue, defaultErr
}

func Build(prev Gate207Snapshot) (Analysis, error) {
	if !prev.Gate207Inherited || !prev.RecommendedOperatorBasisAudit {
		return Analysis{}, fmt.Errorf("Gate 208 requires the Gate 207 operator-basis audit obligation")
	}

	inventory := buildInventory()
	opSearch := buildOperatorSearch(inventory)
	cons := buildConservationAudit(inventory, opSearch)
	firewall := buildFirewall(cons)

	status := StatusProtonDecayChannelConstructionObstructed
	truth := "Gate 208 proves a current-engine proton-decay channel construction obstruction: the contact-preserving gauge connection has no X/Y or B/L-violating curvature, the scalar-bundle functional tau_eta does not supply a local four-fermion B/L operator, and the u(4) leptoquark matter-current slots remain unactivated without a finite action, propagator, local-field map, and coefficient. Therefore ASHA cannot legally compute a proton lifetime or import SU(5) formulas. This is a current-connection algebraic proton-stability theorem, not an absolute proof that all future Pati-Salam-sector dynamics forbid B/L violation."

	return Analysis{
		Gate207:        prev,
		Inventory:      inventory,
		OperatorSearch: opSearch,
		Conservation:   cons,
		Firewall:       firewall,
		Status:         status,
		TruthStatement: truth,
	}, nil
}

func buildInventory() MatterCurrentInventory {
	fockModes := 4
	central := 1
	color := 8
	bl := 1
	leptoquark := 6
	u4 := fockModes * fockModes
	return MatterCurrentInventory{
		FockModes:                         fockModes,
		U4Dimension:                       u4,
		CentralU1Dimension:                central,
		ColorSU3Dimension:                 color,
		BMinusLDimension:                  bl,
		LeptoquarkOffDiagonalDimension:    leptoquark,
		DecompositionComplete:             central+color+bl+leptoquark == u4,
		ContainsQuarkLeptonCurrentSlots:   leptoquark > 0,
		LeptoquarkSlotsGaugeActivated:     false,
		ContactConnectionAlgebra:          "contact-preserving su(2)+u(1) seed on the Boolean/contact carrier",
		ContactConnectionHasColor:         false,
		ContactConnectionHasLeptoquark:    false,
		FullSU5OrSO10GaugeConnection:      false,
		BMinusLConservationAloneForbidsPD: false,
		Verdict:                           "u(4) contains six quark-lepton current slots as inventory, but the derived contact gauge connection does not contain X/Y-like leptoquark curvature and no finite action activates those slots as proton-decay mediators",
	}
}

func buildOperatorSearch(inv MatterCurrentInventory) OperatorSearchAudit {
	templates := []OperatorTemplate{
		{
			Name:                          "left-handed QQQL",
			SymbolicForm:                  "epsilon_color epsilon_weak (Q Q)(Q L) / Lambda^2",
			DeltaB:                        1,
			DeltaL:                        1,
			DeltaBMinusL:                  0,
			SMGaugeInvariantTemplate:      true,
			RequiresThreeQuarkClosure:     true,
			RequiresQuarkLeptonMediator:   true,
			RequiresLocalContinuumFields:  true,
			RequiresBLCouplingCoefficient: true,
			SupportedByContactConnection:  false,
			SupportedByU4CurrentAction:    inv.LeptoquarkSlotsGaugeActivated,
			SupportedByTauEta:             false,
			ConstructedByFiniteAlgebra:    false,
			Obstruction:                   "B-L does not forbid QQQL, but ASHA has not derived a local four-Weyl product, X/Y mediator, finite leptoquark propagator, or coefficient that would instantiate this template",
		},
		{
			Name:                          "right/conjugate UUD E",
			SymbolicForm:                  "epsilon_color (u^c u^c)(d^c e^c) / Lambda^2",
			DeltaB:                        -1,
			DeltaL:                        -1,
			DeltaBMinusL:                  0,
			SMGaugeInvariantTemplate:      true,
			RequiresThreeQuarkClosure:     true,
			RequiresQuarkLeptonMediator:   true,
			RequiresLocalContinuumFields:  true,
			RequiresBLCouplingCoefficient: true,
			SupportedByContactConnection:  false,
			SupportedByU4CurrentAction:    inv.LeptoquarkSlotsGaugeActivated,
			SupportedByTauEta:             false,
			ConstructedByFiniteAlgebra:    false,
			Obstruction:                   "the charge template is gauge-neutral and B-L preserving, but no finite ASHA connection/current action produces its local dimension-six operator or suppression scale",
		},
		{
			Name:                          "mixed QQL d-like channel",
			SymbolicForm:                  "schematic Q Q L d^c / Lambda^2",
			DeltaB:                        1,
			DeltaL:                        1,
			DeltaBMinusL:                  0,
			SMGaugeInvariantTemplate:      true,
			RequiresThreeQuarkClosure:     true,
			RequiresQuarkLeptonMediator:   true,
			RequiresLocalContinuumFields:  true,
			RequiresBLCouplingCoefficient: true,
			SupportedByContactConnection:  false,
			SupportedByU4CurrentAction:    inv.LeptoquarkSlotsGaugeActivated,
			SupportedByTauEta:             false,
			ConstructedByFiniteAlgebra:    false,
			Obstruction:                   "no finite operator product, Fierz channel, or Pati-Salam leptoquark exchange action has been derived for this schematic B/L-violating class",
		},
	}

	constructed := false
	contact := false
	u4 := false
	tau := false
	gaugeTemplate := false
	for _, t := range templates {
		constructed = constructed || t.ConstructedByFiniteAlgebra
		contact = contact || t.SupportedByContactConnection
		u4 = u4 || t.SupportedByU4CurrentAction
		tau = tau || t.SupportedByTauEta
		gaugeTemplate = gaugeTemplate || t.SMGaugeInvariantTemplate
	}
	return OperatorSearchAudit{
		TemplatesAudited:               len(templates),
		SMGaugeInvariantTemplatesExist: gaugeTemplate,
		AnyTemplateConstructed:         constructed,
		AnyContactMediatedChannel:      contact,
		AnyU4MediatedChannel:           u4,
		AnyTauEtaSupportedChannel:      tau,
		SuppressionScaleComputed:       false,
		SuppressionScaleExpression:     "not legal: no B/L-violating local operator or mediator coefficient was constructed",
		Verdict:                        "FAILED_ROUTE: SM-gauge-neutral B/L-violating templates exist as external QFT patterns, but ASHA currently constructs none of them from the contact connection, tau_eta, or activated u(4) current dynamics",
		Templates:                      templates,
	}
}

func buildConservationAudit(inv MatterCurrentInventory, ops OperatorSearchAudit) ConservationAudit {
	return ConservationAudit{
		CurrentConnectionProtonStable:           !inv.ContactConnectionHasLeptoquark && !ops.AnyContactMediatedChannel,
		ExactBaryonConservationProved:           false,
		ExactLeptonConservationProved:           false,
		BMinusLConservedByTemplates:             true,
		BMinusLForbidsQQQL:                      false,
		ColorTrialityForbidsQQQL:                false,
		TrialityOrContactInvariantForbidsAll:    false,
		LeptoquarkInventoryPreventsAbsoluteNoGo: inv.ContainsQuarkLeptonCurrentSlots,
		ProtonDecayChannelConstructionFailed:    !ops.AnyTemplateConstructed,
		TheoremName:                             "Algebraic Proton Stability Theorem, current-connection version",
		Verdict:                                 "Current contact connection is proton-stable by mediator/operator absence; absolute baryon conservation is not proven because u(4) contains unactivated leptoquark current slots and B-L/color triality do not forbid standard dimension-six templates",
	}
}

func buildFirewall(cons ConservationAudit) FirewallAudit {
	return FirewallAudit{
		NoSU5Imported:                        true,
		NoSO10Imported:                       true,
		NoObservedLifetimeBoundsUsedForProof: true,
		NoProtonLifetimeComputed:             true,
		NoXYPresumed:                         true,
		NoLeptoquarkGaugeActivationPresumed:  true,
		NoBaryonConservationOverclaimed:      !cons.ExactBaryonConservationProved && cons.LeptoquarkInventoryPreventsAbsoluteNoGo,
		CurrentConnectionOnly:                true,
		FuturePatiSalamDynamicsStillOpen:     true,
		RecommendedNextGate:                  "Gate 209 — Pati-Salam leptoquark current dynamics / B-L-preserving proton-decay operator seal audit",
		RemainingUnknowns: []string{
			"derive or reject a finite kinetic/action term for the six u(4) leptoquark current slots",
			"derive a local continuum four-Weyl operator product map before writing QQQL or UUD E as an ASHA operator",
			"derive Fierz coefficients and propagator denominators for leptoquark current exchange",
			"derive a baryon-number grading if exact proton stability is to become stronger than current-connection stability",
			"keep proton lifetime formulas sealed until a concrete B/L-violating operator coefficient exists",
		},
	}
}

func FormatInventory(inv MatterCurrentInventory) string {
	return fmt.Sprintf("u(4)=%d = central:%d + su3:%d + B-L:%d + leptoquark:%d; contact=%s; LQ activated=%v", inv.U4Dimension, inv.CentralU1Dimension, inv.ColorSU3Dimension, inv.BMinusLDimension, inv.LeptoquarkOffDiagonalDimension, inv.ContactConnectionAlgebra, inv.LeptoquarkSlotsGaugeActivated)
}

func FormatTemplates(xs []OperatorTemplate) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		state := "blocked"
		if x.ConstructedByFiniteAlgebra {
			state = "constructed"
		}
		parts = append(parts, fmt.Sprintf("%s[%s, Δ(B-L)=%d]: %s", x.Name, state, x.DeltaBMinusL, x.Obstruction))
	}
	return strings.Join(parts, " | ")
}

func FormatUnknowns(xs []string) string { return strings.Join(xs, "; ") }
