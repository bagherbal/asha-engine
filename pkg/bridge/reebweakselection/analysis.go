// Package reebweakselection implements Gate 241:
// Reeb Vector Spatial Isotropy Break / Contact Geometry Sieve Audit.
//
// Gate 240 sharpened the weak-plane problem: the native u(1) twist rejects the
// three temporal-spatial planes and leaves exactly three pure-spatial planes.
// This gate audits whether the finite contact geometry supplies the missing
// selector through a Reeb vector.  A true contact structure (eta,d eta,R) would
// pick a distinguished direction R satisfying eta(R)=1 and i_R d eta=0.  If a
// derived map from that Reeb direction to the three spatial Fock modes existed,
// it could break the remaining S3 color/spatial degeneracy and select the
// complementary two-plane as the weak plane.
//
// The current finite core contains the contact space K as a seven-dimensional
// projector/intersection inside Λ⁴R⁸.  It does not yet contain an explicit
// one-form eta, a two-form d eta, a Reeb vector field, or a bridge from K to the
// Fock generator carrier W=C^4.  Therefore the contact geometry supplies the
// correct type of future selector, but no native axis selector is derived in
// this gate.
package reebweakselection

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/spinctwistedchirality"
	"github.com/bagherbal/asha-engine/pkg/geometry/contact"
	"github.com/bagherbal/asha-engine/pkg/spinor"
)

const (
	AuditID = "GATE241-REEB-VECTOR-SPATIAL-ISOTROPY-WEAK-PLANE-SIEVE"

	StatusContactKRetrieved        = "CONDITIONAL_SUPPORT_CONTACT_K_RETRIEVED_PREFLIGHT"
	StatusReebSelectorType         = "CONDITIONAL_SUPPORT_REEB_SELECTOR_TYPE_PREFLIGHT"
	StatusFailedReebVector         = "FAILED_ROUTE_NATIVE_REEB_VECTOR_DERIVATION"
	StatusFailedEtaDEta            = "FAILED_ROUTE_CONTACT_FORM_ETA_DETA_DERIVATION"
	StatusFailedKToFockProjection  = "FAILED_ROUTE_CONTACT_TO_FOCK_SPATIAL_PROJECTION"
	StatusFailedSpatialAxis        = "FAILED_ROUTE_SPATIAL_AXIS_TAG_DERIVATION"
	StatusFailedWeakPlaneSelection = "FAILED_ROUTE_REEB_VECTOR_WEAK_PLANE_SELECTION"
	StatusFailedGlobalH            = "FAILED_ROUTE_GLOBAL_H_SUMMAND_STILL_UNSELECTED"
)

type ContactStructureAudit struct {
	ContactDimension        int
	ExpectedContactDim      int
	ContactIndex            float64
	FrameIsometryResidual   float64
	BooleanContainment      float64
	G2Containment           float64
	ContactProjectorExists  bool
	EtaOneFormDerived       bool
	DEtaTwoFormDerived      bool
	ReebVectorDerived       bool
	FiniteContactActionOnly bool
	Verdict                 string
}

type ReebVectorAudit struct {
	Definition                 string
	EtaOfReebCondition         string
	ContractionCondition       string
	CandidateAvailable         bool
	NativeFromVacuumStabilizer bool
	LivesOnContactK            bool
	MappedToFockGeneratorW     bool
	MappedToSpatialFockAxes    bool
	ManualAxisChoice           bool
	Verdict                    string
}

type SpatialProjectionAudit struct {
	SpatialAxes                  []string
	PureSpatialPlanes            []string
	ProjectionMapName            string
	KToWProjectionDerived        bool
	ReebComponentsOnSpatialAxes  []float64
	ComponentSource              string
	ComponentsAreUniformOrAbsent bool
	UniqueSpatialAxisTagged      bool
	TaggedAxis                   string
	S3PermutationBroken          bool
	Verdict                      string
}

type PlaneReebAudit struct {
	Plane                string
	ModeIndices          []int
	ComplementAxis       string
	InheritedFromGate240 bool
	SurvivesU1Twist      bool
	RequiresTaggedAxis   bool
	SelectedByReeb       bool
	SelectionRule        string
	Verdict              string
}

type ReebSieveAudit struct {
	InheritedGate240Candidates []string
	CandidatePlaneCount        int
	TaggedAxis                 string
	SelectedPlanes             []string
	HypotheticalRule           string
	S3DegeneracyBroken         bool
	UniqueWeakPlaneSelected    bool
	Verdict                    string
}

type WeakOutcomeAudit struct {
	Gate240ReducedToPureSpatial bool
	ContactGeometryAvailable    bool
	ReebSelectorDerived         bool
	ContactToFockBridgeDerived  bool
	UniqueWeakPlaneSelected     bool
	PhysicalLeftHandedDerived   bool
	GlobalHSummandDerived       bool
	OrderOneReady               bool
	Verdict                     string
}

type FirewallAudit struct {
	ForcedReebAxis               bool
	ImportedContactCoordinates   bool
	ImportedSMWeakPlane          bool
	ImportedElectroweakChirality bool
	PromotedProjectorToReeb      bool
	ClaimedGlobalH               bool
	ClaimedOrderOne              bool
	FiniteCorePolluted           bool
	Verdict                      string
}

type Summary struct {
	ContactKAvailable          bool
	EtaDEtaDerived             bool
	ReebVectorDerived          bool
	ContactToFockProjection    bool
	SpatialAxisTagged          bool
	PureSpatialPlanesInherited int
	UniqueWeakPlaneDerived     bool
	PhysicalChiralityDerived   bool
	GlobalHDerived             bool
	Status                     string
	NextGate                   string
	Comment                    string
}

type Analysis struct {
	Previous       spinctwistedchirality.Analysis
	Contact        ContactStructureAudit
	Reeb           ReebVectorAudit
	Projection     SpatialProjectionAudit
	Planes         []PlaneReebAudit
	Sieve          ReebSieveAudit
	Weak           WeakOutcomeAudit
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
		prev, err := spinctwistedchirality.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 240 predecessor: %w", err)
			return
		}
		space, err := contact.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("construct contact space: %w", err)
			return
		}
		f, err := spinor.NewCovariantPhaseFockSpace(4)
		if err != nil {
			defaultErr = fmt.Errorf("construct Fock space: %w", err)
			return
		}
		defaultA, defaultErr = Build(prev, space, f)
	})
	return defaultA, defaultErr
}

func Build(prev spinctwistedchirality.Analysis, space contact.Space, f spinor.FockSpace) (Analysis, error) {
	if f.ModeCount() != 4 || f.StateCount() != 16 {
		return Analysis{}, fmt.Errorf("Gate 241 requires native four-mode 16-state Fock carrier, got modes=%d states=%d", f.ModeCount(), f.StateCount())
	}
	contactAudit, err := auditContact(space)
	if err != nil {
		return Analysis{}, err
	}
	re := auditReeb(contactAudit)
	proj := auditProjection(f, prev, re)
	planes := auditPlanes(f, prev, proj)
	sieve := auditSieve(planes, proj)
	weak := auditWeak(prev, contactAudit, re, proj, sieve)
	fw := auditFirewall()
	sum := summarize(contactAudit, re, proj, sieve, weak)
	truth := buildTruth(contactAudit, re, proj, sieve, weak)
	return Analysis{Previous: prev, Contact: contactAudit, Reeb: re, Projection: proj, Planes: planes, Sieve: sieve, Weak: weak, Firewall: fw, Summary: sum, TruthStatement: truth}, nil
}

func auditContact(space contact.Space) (ContactStructureAudit, error) {
	frameResidual, err := space.FrameIsometryResidual()
	if err != nil {
		return ContactStructureAudit{}, fmt.Errorf("contact frame residual: %w", err)
	}
	booleanResidual, err := space.BooleanContainmentResidual()
	if err != nil {
		return ContactStructureAudit{}, fmt.Errorf("contact Boolean containment: %w", err)
	}
	g2Residual, err := space.G2ContainmentResidual()
	if err != nil {
		return ContactStructureAudit{}, fmt.Errorf("contact G2 containment: %w", err)
	}
	idx := space.ContactIndex()
	return ContactStructureAudit{
		ContactDimension:        space.Dimension(),
		ExpectedContactDim:      space.ExpectedContactDenominator(),
		ContactIndex:            idx,
		FrameIsometryResidual:   frameResidual,
		BooleanContainment:      booleanResidual,
		G2Containment:           g2Residual,
		ContactProjectorExists:  space.Dimension() == space.ExpectedContactDenominator() && math.Abs(idx-1) < 1e-8,
		EtaOneFormDerived:       false,
		DEtaTwoFormDerived:      false,
		ReebVectorDerived:       false,
		FiniteContactActionOnly: true,
		Verdict:                 StatusContactKRetrieved + "; " + StatusFailedEtaDEta + "; " + StatusFailedReebVector,
	}, nil
}

func auditReeb(c ContactStructureAudit) ReebVectorAudit {
	return ReebVectorAudit{
		Definition:                 "Reeb vector R of a contact form eta satisfies eta(R)=1 and i_R d eta=0",
		EtaOfReebCondition:         "blocked: eta is not constructed as a one-form on the finite carrier",
		ContractionCondition:       "blocked: d eta is not constructed as a two-form with contraction operation",
		CandidateAvailable:         false,
		NativeFromVacuumStabilizer: false,
		LivesOnContactK:            c.ContactProjectorExists && false,
		MappedToFockGeneratorW:     false,
		MappedToSpatialFockAxes:    false,
		ManualAxisChoice:           false,
		Verdict:                    StatusReebSelectorType + "; " + StatusFailedReebVector,
	}
}

func auditProjection(f spinor.FockSpace, prev spinctwistedchirality.Analysis, re ReebVectorAudit) SpatialProjectionAudit {
	axes := []string{}
	for _, m := range f.Modes {
		if m.Kind == spinor.SpatialMode {
			axes = append(axes, m.Name)
		}
	}
	planes := append([]string(nil), prev.Sieve.U1PreservingPlanes...)
	sort.Strings(planes)
	return SpatialProjectionAudit{
		SpatialAxes:                  axes,
		PureSpatialPlanes:            planes,
		ProjectionMapName:            "π_K→W_spatial",
		KToWProjectionDerived:        false,
		ReebComponentsOnSpatialAxes:  nil,
		ComponentSource:              "absent: no derived Reeb vector and no K-to-Fock projection map",
		ComponentsAreUniformOrAbsent: true,
		UniqueSpatialAxisTagged:      false,
		TaggedAxis:                   "",
		S3PermutationBroken:          false,
		Verdict:                      StatusFailedKToFockProjection + "; " + StatusFailedSpatialAxis,
	}
}

func auditPlanes(f spinor.FockSpace, prev spinctwistedchirality.Analysis, proj SpatialProjectionAudit) []PlaneReebAudit {
	preserve := map[string]bool{}
	for _, p := range prev.Sieve.U1PreservingPlanes {
		preserve[p] = true
	}
	out := []PlaneReebAudit{}
	for i := 0; i < f.ModeCount(); i++ {
		for j := i + 1; j < f.ModeCount(); j++ {
			if f.Modes[i].Kind != spinor.SpatialMode || f.Modes[j].Kind != spinor.SpatialMode {
				continue
			}
			plane := fmt.Sprintf("U={%s,%s}", f.Modes[i].Name, f.Modes[j].Name)
			comp := complementSpatialAxis(f, i, j)
			selected := proj.UniqueSpatialAxisTagged && comp == proj.TaggedAxis
			verdict := StatusFailedWeakPlaneSelection
			if selected {
				verdict = "CONDITIONAL_SUPPORT_REEB_VECTOR_WEAK_PLANE_SELECTION"
			}
			out = append(out, PlaneReebAudit{
				Plane:                plane,
				ModeIndices:          []int{i, j},
				ComplementAxis:       comp,
				InheritedFromGate240: preserve[plane],
				SurvivesU1Twist:      preserve[plane],
				RequiresTaggedAxis:   true,
				SelectedByReeb:       selected,
				SelectionRule:        "if Reeb tags spatial axis e_k, the weak plane would be the complementary spatial two-plane",
				Verdict:              verdict,
			})
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Plane < out[j].Plane })
	return out
}

func complementSpatialAxis(f spinor.FockSpace, i, j int) string {
	for _, m := range f.Modes {
		if m.Kind == spinor.SpatialMode && m.Index != i && m.Index != j {
			return m.Name
		}
	}
	return ""
}

func auditSieve(planes []PlaneReebAudit, proj SpatialProjectionAudit) ReebSieveAudit {
	selected := []string{}
	for _, p := range planes {
		if p.SelectedByReeb {
			selected = append(selected, p.Plane)
		}
	}
	return ReebSieveAudit{
		InheritedGate240Candidates: proj.PureSpatialPlanes,
		CandidatePlaneCount:        len(planes),
		TaggedAxis:                 proj.TaggedAxis,
		SelectedPlanes:             selected,
		HypotheticalRule:           "tag one spatial axis by a native Reeb vector, select the complementary spatial two-plane",
		S3DegeneracyBroken:         proj.S3PermutationBroken,
		UniqueWeakPlaneSelected:    len(selected) == 1,
		Verdict:                    StatusFailedWeakPlaneSelection,
	}
}

func auditWeak(prev spinctwistedchirality.Analysis, c ContactStructureAudit, re ReebVectorAudit, proj SpatialProjectionAudit, sieve ReebSieveAudit) WeakOutcomeAudit {
	return WeakOutcomeAudit{
		Gate240ReducedToPureSpatial: prev.Summary.U1RejectsTemporalPlanes && prev.Summary.PureSpatialPlanesRemain == 3,
		ContactGeometryAvailable:    c.ContactProjectorExists,
		ReebSelectorDerived:         re.CandidateAvailable && re.MappedToSpatialFockAxes,
		ContactToFockBridgeDerived:  proj.KToWProjectionDerived,
		UniqueWeakPlaneSelected:     sieve.UniqueWeakPlaneSelected,
		PhysicalLeftHandedDerived:   false,
		GlobalHSummandDerived:       false,
		OrderOneReady:               false,
		Verdict:                     StatusFailedWeakPlaneSelection + "; " + StatusFailedGlobalH,
	}
}

func auditFirewall() FirewallAudit {
	return FirewallAudit{
		ForcedReebAxis:               false,
		ImportedContactCoordinates:   false,
		ImportedSMWeakPlane:          false,
		ImportedElectroweakChirality: false,
		PromotedProjectorToReeb:      false,
		ClaimedGlobalH:               false,
		ClaimedOrderOne:              false,
		FiniteCorePolluted:           false,
		Verdict:                      "FIREWALL_PRESERVED_NO_REEB_AXIS_FORCED",
	}
}

func summarize(c ContactStructureAudit, re ReebVectorAudit, proj SpatialProjectionAudit, sieve ReebSieveAudit, weak WeakOutcomeAudit) Summary {
	status := strings.Join([]string{StatusContactKRetrieved, StatusReebSelectorType, StatusFailedEtaDEta, StatusFailedReebVector, StatusFailedKToFockProjection, StatusFailedSpatialAxis, StatusFailedWeakPlaneSelection, StatusFailedGlobalH}, ";")
	return Summary{
		ContactKAvailable:          c.ContactProjectorExists,
		EtaDEtaDerived:             c.EtaOneFormDerived && c.DEtaTwoFormDerived,
		ReebVectorDerived:          re.CandidateAvailable,
		ContactToFockProjection:    proj.KToWProjectionDerived,
		SpatialAxisTagged:          proj.UniqueSpatialAxisTagged,
		PureSpatialPlanesInherited: sieve.CandidatePlaneCount,
		UniqueWeakPlaneDerived:     weak.UniqueWeakPlaneSelected,
		PhysicalChiralityDerived:   weak.PhysicalLeftHandedDerived,
		GlobalHDerived:             weak.GlobalHSummandDerived,
		Status:                     status,
		NextGate:                   "derive an explicit contact one-form/Reeb vector and a natural projection from K⊂Λ⁴R⁸ to the Fock generator carrier W, or prove the weak-plane selector must be externally sealed",
		Comment:                    "The contact projector K is exact finite geometry, but no eta,deta,Reeb vector, or K-to-Fock spatial projection is derived. Therefore the three pure-spatial weak planes remain S3-degenerate.",
	}
}

func buildTruth(c ContactStructureAudit, re ReebVectorAudit, proj SpatialProjectionAudit, sieve ReebSieveAudit, weak WeakOutcomeAudit) string {
	return fmt.Sprintf("Gate 241 inherits Gate 240's three pure-spatial candidates and retrieves contact K with dim=%d,index=%.6g. A Reeb vector would be the right kind of selector, but eta,deta,R and the K→W_spatial projection are not derived; no spatial axis is tagged and no weak plane is selected.", c.ContactDimension, c.ContactIndex)
}
