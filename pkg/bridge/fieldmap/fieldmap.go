// Package fieldmap audits the finite-to-continuum status of the scalar/contact
// doublet.
//
// Gate 47 is not allowed to decide that the scalar/contact sector is the
// observed electroweak Higgs field merely because it has four real directions
// and a doublet charge.  It also must not demote the sector to a heavy
// threshold or regulator without a decoupling/matching rule.  The package
// therefore classifies the sector by evidence: what is already derived, what is
// missing for each continuum interpretation, and which interpretations are
// currently forbidden.
package fieldmap

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/scalarscale"
	"github.com/bagherbal/asha-engine/pkg/bridge/thresholdactivation"
	"github.com/bagherbal/asha-engine/pkg/dynamics/scalarpotential"
)

type Classification string

const (
	ContinuumScalarCandidate Classification = "continuum-scalar-candidate"
	LowEnergyFieldOpen       Classification = "low-energy-field-open"
	HighScaleSourceOpen      Classification = "high-scale-source-open"
	FiniteRegulatorRejected  Classification = "finite-regulator-rejected"
)

type EvidenceItem struct {
	Name    string
	Present bool
	Detail  string
}

type RoleRequirement struct {
	Role        Classification
	Description string
	Required    []string
	Missing     []string
	Satisfied   bool
}

type Analysis struct {
	Activation thresholdactivation.Analysis
	Scale      scalarscale.Analysis
	Scalar     scalarpotential.Analysis

	ActiveRealDirections      int
	ComplexDoubletComponents  int
	ProtectedDirections       int
	SectorLevelDoubletDerived bool
	ScalarPotentialDerived    bool
	YukawaChannelsDerived     bool

	PhysicalUnitDerived            bool
	LowEnergyScaleDerived          bool
	BoundaryScaleDerived           bool
	DecouplingRuleDerived          bool
	RegulatorClassificationDerived bool
	ThresholdCorrectionAllowed     bool
	HiddenObservedInput            bool

	PrimaryClassification Classification
	Evidence              []EvidenceItem
	RoleRequirements      []RoleRequirement
	TruthStatement        string
	MinimumMissingData    []string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		act, err := thresholdactivation.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		scale, err := scalarscale.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		scalar, err := scalarpotential.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(act, scale, scalar)
	})
	return defaultValue, defaultErr
}

func Build(act thresholdactivation.Analysis, scale scalarscale.Analysis, scalar scalarpotential.Analysis) (Analysis, error) {
	if scalar.ActiveRealDimension != 4 {
		return Analysis{}, fmt.Errorf("field-map audit expects four active scalar/contact directions, got %d", scalar.ActiveRealDimension)
	}
	if !act.ScalarSectorRemainsContinuumCandidate {
		return Analysis{}, fmt.Errorf("threshold activation audit did not expose scalar/contact sector as continuum-field candidate")
	}

	physicalUnit := scale.HasDimensionfulAnchor
	decoupling := act.DecouplingRuleDerived
	thresholdAllowed := act.BetaCorrectionAllowedCount > 0

	evidence := []EvidenceItem{
		{Name: "four real scalar/contact directions", Present: scalar.ActiveRealDimension == 4, Detail: "4 real active modes = 2 complex scalar-doublet components"},
		{Name: "sector-level scalar doublet representation", Present: act.ScalarSectorRemainsContinuumCandidate, Detail: "threshold activation keeps the scalar/contact sector as a field candidate, not a heavy threshold"},
		{Name: "finite scalar potential normal form", Present: scalar.ShiftedNormalFormAvailable, Detail: fmt.Sprintf("r0²=%.10f, lambda_shape=%.10f", scalar.VacuumRadiusSquared, scalar.LambdaShape)},
		{Name: "protected-direction resonance", Present: scalar.ProtectedDirectionCount == 3, Detail: "three protected contact directions resonate with would-be broken gauge directions, but gauge-eating remains open"},
		{Name: "physical scale bridge", Present: physicalUnit, Detail: "not derived; the scalar sector remains dimensionless"},
		{Name: "decoupling/regulator rule", Present: decoupling, Detail: "not derived; no rule classifies the scalar/contact doublet as integrated-out regulator"},
		{Name: "threshold correction permission", Present: thresholdAllowed, Detail: "not allowed; scalar/contact sector is not a beta-correcting heavy threshold"},
	}

	roles := []RoleRequirement{
		{
			Role:        ContinuumScalarCandidate,
			Description: "finite scalar/contact doublet candidate",
			Required: []string{
				"four real active directions",
				"sector-level (1,2) scalar doublet representation",
				"finite scalar potential shape",
			},
			Missing:   nil,
			Satisfied: true,
		},
		{
			Role:        LowEnergyFieldOpen,
			Description: "observed low-energy electroweak Higgs field",
			Required: []string{
				"non-fitted physical unit or electroweak scale bridge",
				"gauge-eating theorem for three protected directions",
				"kinetic normalization and RG matching",
			},
			Missing: []string{
				"physical unit mu is free",
				"v=246 GeV is not derived",
				"Higgs mass bridge is not derived",
				"gauge-eating theorem is not derived",
			},
			Satisfied: false,
		},
		{
			Role:        HighScaleSourceOpen,
			Description: "high-scale finite source seed for a continuum scalar",
			Required: []string{
				"boundary scale M*",
				"finite-to-continuum matching rule",
				"RG evolution and threshold activation map",
			},
			Missing: []string{
				"boundary scale M* is not derived",
				"finite-to-continuum field normalization is not derived",
				"threshold activation map is open",
			},
			Satisfied: false,
		},
		{
			Role:        FiniteRegulatorRejected,
			Description: "integrated-out finite regulator object",
			Required: []string{
				"decoupling prescription",
				"regulator cancellation or matching role",
				"proof that scalar/contact doublet is not continuum-active",
			},
			Missing: []string{
				"no decoupling prescription",
				"no regulator cancellation theorem",
				"derived doublet representation points toward a field candidate instead",
			},
			Satisfied: false,
		},
	}

	return Analysis{
		Activation:                     act,
		Scale:                          scale,
		Scalar:                         scalar,
		ActiveRealDirections:           scalar.ActiveRealDimension,
		ComplexDoubletComponents:       scalar.ActiveRealDimension / 2,
		ProtectedDirections:            scalar.ProtectedDirectionCount,
		SectorLevelDoubletDerived:      act.ScalarSectorRemainsContinuumCandidate,
		ScalarPotentialDerived:         scalar.ShiftedNormalFormAvailable,
		YukawaChannelsDerived:          true,
		PhysicalUnitDerived:            physicalUnit,
		LowEnergyScaleDerived:          false,
		BoundaryScaleDerived:           false,
		DecouplingRuleDerived:          decoupling,
		RegulatorClassificationDerived: false,
		ThresholdCorrectionAllowed:     thresholdAllowed,
		HiddenObservedInput:            false,
		PrimaryClassification:          ContinuumScalarCandidate,
		Evidence:                       evidence,
		RoleRequirements:               roles,
		TruthStatement:                 "The scalar/contact active sector is best classified as a finite continuum-scalar candidate: it has four real active directions, a sector-level doublet representation, and a finite potential shape. It is not yet the observed low-energy Higgs field, not yet a high-scale source seed with matching data, and not a derived finite regulator or heavy threshold.",
		MinimumMissingData: []string{
			"derive a non-fitted physical unit or electroweak scale bridge before identifying the candidate with the observed Higgs field",
			"derive the gauge-eating theorem linking three protected contact directions to broken SU(2)_L directions",
			"derive finite-to-continuum kinetic normalization and RG matching for the scalar field",
			"derive a boundary scale M* before treating the scalar/contact sector as a high-scale source seed",
			"derive a decoupling prescription before treating any scalar/contact mode as a regulator or heavy threshold",
		},
	}, nil
}

func FormatEvidence(items []EvidenceItem) string {
	parts := make([]string, 0, len(items))
	for _, e := range items {
		mark := "no"
		if e.Present {
			mark = "yes"
		}
		parts = append(parts, fmt.Sprintf("%s:%s", e.Name, mark))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatRoleRequirements(rs []RoleRequirement) string {
	parts := make([]string, 0, len(rs))
	for _, r := range rs {
		state := "open"
		if r.Satisfied {
			state = "satisfied"
		}
		parts = append(parts, fmt.Sprintf("%s:%s", r.Role, state))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}
