// Package looppotential audits the native one-loop effective-potential problem.
//
// The older Higgs prelude identifies the decisive missing computation: derive
// the scalar mass-parameter instability using the Cℓ(1,7) Fock/Yukawa operator
// system itself, instead of importing the Standard Model RGE as a finished
// result.  This package builds the ledger for that computation.
//
// It deliberately does not claim that μ² is negative, does not insert measured
// Yukawa couplings, and does not use observed electroweak data.  It only records
// which sign/multiplicity factors are already forced by the finite matter
// inventory, and which operator-level data are still missing before a native
// one-loop potential can be derived.
package looppotential

import (
	"fmt"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/condensate"
	"github.com/bagherbal/asha-engine/pkg/dynamics/scalarpotential"
	"github.com/bagherbal/asha-engine/pkg/matter/yukawaintertwiner"
	"github.com/bagherbal/asha-engine/pkg/spinor"
)

type ContributionSign string

const (
	Negative ContributionSign = "negative"
	Positive ContributionSign = "positive"
	Unknown  ContributionSign = "unknown"
)

type LedgerTerm struct {
	Name         string
	Sector       string
	Sign         ContributionSign
	Multiplicity int
	Coefficient  string
	Status       string
	Requirement  string
}

type Analysis struct {
	Condensate condensate.Analysis
	Fock       spinor.FockSpace
	Yukawa     yukawaintertwiner.Analysis
	Scalar     scalarpotential.Analysis

	FermionLoopSignAvailable          bool
	AnticommutationSubstrateAvailable bool
	ColorAmplificationFactor          int
	SpinChiralityMultiplicity         int
	TopLikeCoefficientSkeleton        int

	GaugePositiveSectorCount    int
	ScalarSelfSectorAvailable   bool
	ScalarQuarticShapeInvariant float64

	Terms []LedgerTerm

	NativeLoopOperatorDerived          bool
	TopLikeYukawaStrengthDerived       bool
	GaugeCouplingsDerived              bool
	ScalarSelfCouplingScaleDerived     bool
	RegulatorOrCutoffDerived           bool
	RenormalizationPrescriptionDerived bool
	MuSquaredSignDerived               bool
	SymmetricOriginInstabilityDerived  bool

	StructuralInstabilityPressureAvailable bool
	NativeEffectivePotentialComputed       bool
	ImportedSMRGE                          bool
	HiddenObservedCouplingsUsed            bool

	TruthStatement      string
	RecommendedNextGate string
	RemainingUnknowns   []string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		c, err := condensate.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(c)
	})
	return defaultValue, defaultErr
}

func Build(c condensate.Analysis) (Analysis, error) {
	fock := c.Fock
	y := c.Yukawa
	sp := c.ScalarPotential
	if fock.ModeCount() != 4 || fock.StateCount() != 16 {
		return Analysis{}, fmt.Errorf("one-loop ledger expects 4-mode/16-state Fock space")
	}
	if y.UpChannels != 3 || y.DownChannels != 3 {
		return Analysis{}, fmt.Errorf("one-loop ledger expects three color channels in the Yukawa audit")
	}
	if sp.ActiveRealDimension != 4 {
		return Analysis{}, fmt.Errorf("one-loop ledger expects four active scalar directions")
	}

	color := y.UpChannels
	spinChiral := 2 // algebraic loop degeneracy placeholder: two chiral/spin contractions in a Dirac-like loop.
	skeleton := -color * spinChiral
	gaugePositive := 4 // SU(2)_L has three generators and U(1)_Y has one.

	terms := []LedgerTerm{
		{
			Name:         "closed fermion-loop sign",
			Sector:       "Fock/Yukawa",
			Sign:         Negative,
			Multiplicity: 1,
			Coefficient:  "−1",
			Status:       "STRUCTURAL",
			Requirement:  "native trace over finite anticommuting Fock operators still required for the actual loop integral",
		},
		{
			Name:         "three-color top-like amplification",
			Sector:       "spatial/color Fock modes",
			Sign:         Negative,
			Multiplicity: color,
			Coefficient:  "−Nc",
			Status:       "STRUCTURAL",
			Requirement:  "derive the top-like Yukawa strength from finite overlap/condensate dynamics rather than inserting y_t",
		},
		{
			Name:         "Dirac/chiral contraction multiplicity",
			Sector:       "left-right bilinear",
			Sign:         Negative,
			Multiplicity: spinChiral,
			Coefficient:  "×2",
			Status:       "PLACEHOLDER",
			Requirement:  "derive this multiplicity from the native Cℓ(1,7) spinor trace, not from the continuum RGE table",
		},
		{
			Name:         "top-like negative skeleton",
			Sector:       "fermion condensate channel",
			Sign:         Negative,
			Multiplicity: -skeleton,
			Coefficient:  fmt.Sprintf("%d · y_top-like²", skeleton),
			Status:       "SKELETON_ONLY",
			Requirement:  "the coefficient skeleton is available, but the native Yukawa strength and loop integral are not computed",
		},
		{
			Name:         "electroweak gauge positive sector",
			Sector:       "SU(2)_L × U(1)_Y",
			Sign:         Positive,
			Multiplicity: gaugePositive,
			Coefficient:  "+ gauge-coupling terms",
			Status:       "OPEN",
			Requirement:  "requires gauge kinetic normalization, couplings, and native loop trace",
		},
		{
			Name:         "scalar self positive sector",
			Sector:       "finite scalar/contact potential",
			Sign:         Positive,
			Multiplicity: sp.ActiveRealDimension,
			Coefficient:  "+ scalar self-coupling term",
			Status:       "BRIDGE_ONLY",
			Requirement:  "finite quartic shape exists, but physical scalar coupling normalization is still open",
		},
	}

	structuralPressure := color == 3 && spinChiral == 2 && skeleton == -6 && c.CompositeHiggsDirectionPreferred

	return Analysis{
		Condensate:                             c,
		Fock:                                   fock,
		Yukawa:                                 y,
		Scalar:                                 sp,
		FermionLoopSignAvailable:               true,
		AnticommutationSubstrateAvailable:      true,
		ColorAmplificationFactor:               color,
		SpinChiralityMultiplicity:              spinChiral,
		TopLikeCoefficientSkeleton:             skeleton,
		GaugePositiveSectorCount:               gaugePositive,
		ScalarSelfSectorAvailable:              true,
		ScalarQuarticShapeInvariant:            sp.LambdaShape,
		Terms:                                  terms,
		NativeLoopOperatorDerived:              false,
		TopLikeYukawaStrengthDerived:           false,
		GaugeCouplingsDerived:                  false,
		ScalarSelfCouplingScaleDerived:         false,
		RegulatorOrCutoffDerived:               false,
		RenormalizationPrescriptionDerived:     false,
		MuSquaredSignDerived:                   false,
		SymmetricOriginInstabilityDerived:      false,
		StructuralInstabilityPressureAvailable: structuralPressure,
		NativeEffectivePotentialComputed:       false,
		ImportedSMRGE:                          false,
		HiddenObservedCouplingsUsed:            false,
		TruthStatement:                         truth(structuralPressure),
		RecommendedNextGate:                    "Gate 53 — Finite Loop Operator Construction",
		RemainingUnknowns: []string{
			"U-20A1-NATIVE-LOOP-TRACE: construct the finite one-loop trace over Fock/Yukawa operators",
			"U-20A2-TOP-LIKE-OVERLAP: derive the top-like Yukawa/overlap strength instead of inserting y_t",
			"U-20A3-BOSONIC-COUNTERWEIGHTS: derive finite gauge/scalar positive loop contributions with kinetic normalization",
			"U-20A4-REGULATOR-CUTOFF: derive the finite cutoff or spectral regulator used by the loop integral",
			"U-20A5-MU-SIGN: compute the native scalar mass-parameter sign rather than importing the SM RGE",
		},
	}, nil
}

func truth(structural bool) string {
	if structural {
		return "The finite engine now reproduces the structural ledger behind radiative Higgs instability: anticommuting Fock matter supplies the fermion-loop sign, the 1+3 Fock split supplies three-color amplification, and the left-right scalar channels supply the candidate top-like condensate path. The native one-loop operator, Yukawa strength, and regulator are still missing, so μ²<0 is not yet derived."
	}
	return "The finite engine does not yet expose the full structural ledger needed for the native one-loop potential."
}

func FormatTerms(terms []LedgerTerm) string {
	if len(terms) == 0 {
		return "none"
	}
	out := ""
	for i, t := range terms {
		if i > 0 {
			out += "; "
		}
		out += fmt.Sprintf("%s[%s,%s,%s]", t.Name, t.Sector, t.Sign, t.Status)
	}
	return out
}
