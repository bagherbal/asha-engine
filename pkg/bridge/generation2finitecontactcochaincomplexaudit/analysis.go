// Package generation2finitecontactcochaincomplexaudit implements Gate 569:
// Finite Contact Cochain Complex and d²=0 Certificate Audit.
//
// The audit asks whether the project can upgrade existing Boolean incidence,
// G₂ calibration, projector, exterior-language, or contact spectral data into
// a genuine finite cochain/exterior differential on K_7.  It requires the
// minimum structural certificates for contact geometry: typed cochain spaces,
// a signed differential d, d²=0, wedge/Leibniz compatibility, an alpha slot,
// computable d alpha, and only then a possible contact volume/Reeb package.
package generation2finitecontactcochaincomplexaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/exterior"
	"github.com/bagherbal/asha-engine/pkg/geometry/boolean"
	"github.com/bagherbal/asha-engine/pkg/geometry/contact"
)

const (
	AuditID = "GATE569-FINITE-CONTACT-COCHAIN-COMPLEX-D2-ZERO-CERTIFICATE-AUDIT"

	StatusK7CarrierInherited                      = "PASS_K7_CONTACT_CARRIER_INHERITED"
	StatusK7ContainmentCertified                  = "PASS_K7_BOOLEAN_G2_CONTAINMENT_CERTIFIED"
	StatusExteriorR7DimensionsAvailable           = "CONDITIONAL_SUPPORT_FORMAL_R7_EXTERIOR_DIMENSIONS_AVAILABLE"
	StatusNoK7CochainBasis                        = "FAILED_ROUTE_NO_CERTIFIED_K7_COCHAIN_BASIS"
	StatusNoK7WedgeProduct                        = "FAILED_ROUTE_NO_CERTIFIED_WEDGE_PRODUCT_ON_K7_COFORMS"
	StatusUnsignedBooleanConsecutiveMapsAvailable = "CONDITIONAL_SUPPORT_UNSIGNED_BOOLEAN_CONSECUTIVE_INCIDENCE_MAPS_AVAILABLE"
	StatusUnsignedBooleanD2Nonzero                = "FAILED_ROUTE_UNSIGNED_BOOLEAN_INCIDENCE_FAILS_D_SQUARED_ZERO"
	StatusBooleanNotSignedCochainDifferential     = "FAILED_ROUTE_BOOLEAN_INCIDENCE_NOT_SIGNED_COCHAIN_DIFFERENTIAL"
	StatusNoRestrictionToK7Complex                = "FAILED_ROUTE_NO_BOOLEAN_RESTRICTION_TO_K7_COCHAIN_COMPLEX"
	StatusNoD2ZeroCertificate                     = "FAILED_ROUTE_NO_D_SQUARED_ZERO_CERTIFICATE_ON_K7"
	StatusNoLeibnizCertificate                    = "FAILED_ROUTE_NO_GRADED_LEIBNIZ_CERTIFICATE_ON_K7"
	StatusG2NoCochainComplex                      = "FAILED_ROUTE_G2_CALIBRATION_DOES_NOT_SUPPLY_K7_COCHAIN_COMPLEX"
	StatusProjectorNoBoundary                     = "FAILED_ROUTE_PROJECTOR_RELATIVE_POSITION_DOES_NOT_SUPPLY_BOUNDARY_OPERATOR"
	StatusQ4NoCochainComplex                      = "FAILED_ROUTE_Q4_SPECTRAL_DATA_DOES_NOT_SUPPLY_COCHAIN_COMPLEX"
	StatusNoAlphaSlotCompatibility                = "FAILED_ROUTE_NO_ALPHA_SLOT_COMPATIBLE_WITH_FINITE_D"
	StatusNoDAlpha                                = "FAILED_ROUTE_NO_FINITE_DALPHA_COMPUTATION"
	StatusNoContactVolume                         = "FAILED_ROUTE_NO_CONTACT_VOLUME_FROM_COCHAIN_COMPLEX"
	StatusNoReeb                                  = "FAILED_ROUTE_NO_REEB_VECTOR_FROM_COCHAIN_COMPLEX"
	StatusNoPhysicalTime                          = "FAILED_ROUTE_NO_COCHAIN_COMPLEX_TO_PHYSICAL_TIME_AIRLOCK"
	StatusNoRGOSHilbert                           = "FAILED_ROUTE_COCHAIN_COMPLEX_AUDIT_DOES_NOT_OPEN_RG_OS_HILBERT_DYNAMICS"
	StatusGate569Firewall                         = "FIREWALL_PRESERVED_GATE569_FINITE_CONTACT_COCHAIN_COMPLEX_BOUNDARY"
)

type K7CarrierAudit struct {
	Dimension                  int
	ExpectedDimension          int
	AmbientDimension           int
	ContactIndex               float64
	FrameIsometryResidual      float64
	BooleanContainmentResidual float64
	G2ContainmentResidual      float64
	Verdict                    string
}

type FormalR7ExteriorAudit struct {
	VectorDimension               int
	GradeDimensions               []int
	TotalDimension                int
	HasAbstractExteriorDimensions bool
	HasCertifiedK7CochainBasis    bool
	HasCertifiedWedgeProductOnK7  bool
	HasFiniteDOperator            bool
	Verdict                       string
}

type BooleanConsecutiveIncidenceAudit struct {
	M23Rows                       int
	M23Cols                       int
	M34Rows                       int
	M34Cols                       int
	Composition34After23Rows      int
	Composition34After23Cols      int
	Composition34After23Frobenius float64
	Composition34After23MaxAbs    float64
	UnsignedIncidence             bool
	D2ZeroForUnsignedIncidence    bool
	SignedOrientationAvailable    bool
	DefinesK7Differential         bool
	Verdict                       string
}

type K7RestrictionAudit struct {
	BooleanIncidenceDomain                string
	BooleanIncidenceCodomain              string
	K7LivesIn                             string
	K7CochainComplexDefined               bool
	ProjectionFromAmbientFormsToK7Coforms bool
	PullbackDifferentialDefined           bool
	D2ZeroOnRestrictedComplex             bool
	Verdict                               string
}

type SourceAudit struct {
	G2CalibrationSuppliesComplex              bool
	ProjectorRelativePositionSuppliesBoundary bool
	Q4SuppliesComplex                         bool
	Verdict                                   string
}

type DifferentialLawAudit struct {
	HasD0ToD1               bool
	HasD1ToD2               bool
	HasD2ToD3               bool
	HasFullComplex          bool
	HasD2ZeroCertificate    bool
	HasLeibnizCertificate   bool
	HasAlphaCompatibleD     bool
	DAlphaComputable        bool
	ContactVolumeComputable bool
	ReebComputable          bool
	Verdict                 string
}

type ProductTimeFirewallAudit struct {
	CochainToDM                  bool
	CochainToLorentzianTime      bool
	CochainToOSPositivity        bool
	CochainToWickRotation        bool
	CochainToHilbert             bool
	CochainToHamiltonian         bool
	CochainToRGScale             bool
	CochainToArrowOfTime         bool
	ElectroweakBridgeStillSealed bool
	Verdict                      string
}

type FinalVerdict struct {
	K7Certified                     bool
	FormalExteriorDimensionsOnly    bool
	UnsignedBooleanIncidenceFailsD2 bool
	SignedFiniteDifferentialFound   bool
	K7CochainComplexFound           bool
	D2ZeroCertified                 bool
	LeibnizCertified                bool
	DAlphaCertified                 bool
	ContactVolumeCertified          bool
	ReebCertified                   bool
	PhysicalTimeOrRGOpened          bool
	MissingNextTheorem              string
	Verdict                         string
}

type Analysis struct {
	K7          K7CarrierAudit
	Exterior    FormalR7ExteriorAudit
	Boolean     BooleanConsecutiveIncidenceAudit
	Restriction K7RestrictionAudit
	Sources     SourceAudit
	Law         DifferentialLawAudit
	Time        ProductTimeFirewallAudit
	Final       FinalVerdict
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
	space, err := contact.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build contact K_7 space: %w", err)
	}
	a := Analysis{}
	a.K7, err = auditK7(space)
	if err != nil {
		return Analysis{}, err
	}
	a.Exterior, err = auditExterior()
	if err != nil {
		return Analysis{}, err
	}
	a.Boolean, err = auditBooleanConsecutiveIncidence()
	if err != nil {
		return Analysis{}, err
	}
	a.Restriction = auditRestriction()
	a.Sources = auditSources()
	a.Law = auditDifferentialLaw()
	a.Time = auditTime()
	a.Final = auditFinal(a)
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func auditK7(space contact.Space) (K7CarrierAudit, error) {
	frame, err := space.FrameIsometryResidual()
	if err != nil {
		return K7CarrierAudit{}, err
	}
	b, err := space.BooleanContainmentResidual()
	if err != nil {
		return K7CarrierAudit{}, err
	}
	g, err := space.G2ContainmentResidual()
	if err != nil {
		return K7CarrierAudit{}, err
	}
	return K7CarrierAudit{Dimension: space.Dimension(), ExpectedDimension: space.ExpectedContactDenominator(), AmbientDimension: space.AmbientDimension(), ContactIndex: space.ContactIndex(), FrameIsometryResidual: frame, BooleanContainmentResidual: b, G2ContainmentResidual: g, Verdict: join(StatusK7CarrierInherited, StatusK7ContainmentCertified)}, nil
}

func auditExterior() (FormalR7ExteriorAudit, error) {
	alg, err := exterior.NewAlgebra(7)
	if err != nil {
		return FormalR7ExteriorAudit{}, err
	}
	dims := make([]int, len(alg.Grades))
	for i, g := range alg.Grades {
		dims[i] = g.Dimension
	}
	return FormalR7ExteriorAudit{VectorDimension: alg.VectorDimension, GradeDimensions: dims, TotalDimension: alg.TotalDimension(), HasAbstractExteriorDimensions: true, HasCertifiedK7CochainBasis: false, HasCertifiedWedgeProductOnK7: false, HasFiniteDOperator: false, Verdict: join(StatusExteriorR7DimensionsAvailable, StatusNoK7CochainBasis, StatusNoK7WedgeProduct)}, nil
}

func auditBooleanConsecutiveIncidence() (BooleanConsecutiveIncidenceAudit, error) {
	b23, err := boolean.BuildIncidenceSupport(8, 2, 3)
	if err != nil {
		return BooleanConsecutiveIncidenceAudit{}, err
	}
	b34, err := boolean.BuildIncidenceSupport(8, 3, 4)
	if err != nil {
		return BooleanConsecutiveIncidenceAudit{}, err
	}
	c3423, err := b34.Incidence.Mul(b23.Incidence)
	if err != nil {
		return BooleanConsecutiveIncidenceAudit{}, err
	}
	return BooleanConsecutiveIncidenceAudit{
		M23Rows: b23.Incidence.Rows(), M23Cols: b23.Incidence.Cols(), M34Rows: b34.Incidence.Rows(), M34Cols: b34.Incidence.Cols(),
		Composition34After23Rows: c3423.Rows(), Composition34After23Cols: c3423.Cols(),
		Composition34After23Frobenius: c3423.FrobeniusNorm(), Composition34After23MaxAbs: c3423.MaxAbs(),
		UnsignedIncidence: true, D2ZeroForUnsignedIncidence: c3423.FrobeniusNorm() < 1e-12, SignedOrientationAvailable: false, DefinesK7Differential: false,
		Verdict: join(StatusUnsignedBooleanConsecutiveMapsAvailable, StatusUnsignedBooleanD2Nonzero, StatusBooleanNotSignedCochainDifferential, StatusNoRestrictionToK7Complex),
	}, nil
}

func auditRestriction() K7RestrictionAudit {
	return K7RestrictionAudit{BooleanIncidenceDomain: "Lambda^3 R^8 and Lambda^4 R^8 ambient Boolean support", BooleanIncidenceCodomain: "ambient middle-grade Boolean support, not K_7 cochains", K7LivesIn: "7-dimensional subspace of Lambda^4 R^8", K7CochainComplexDefined: false, ProjectionFromAmbientFormsToK7Coforms: false, PullbackDifferentialDefined: false, D2ZeroOnRestrictedComplex: false, Verdict: StatusNoRestrictionToK7Complex}
}

func auditSources() SourceAudit {
	return SourceAudit{G2CalibrationSuppliesComplex: false, ProjectorRelativePositionSuppliesBoundary: false, Q4SuppliesComplex: false, Verdict: join(StatusG2NoCochainComplex, StatusProjectorNoBoundary, StatusQ4NoCochainComplex)}
}

func auditDifferentialLaw() DifferentialLawAudit {
	return DifferentialLawAudit{HasD0ToD1: false, HasD1ToD2: false, HasD2ToD3: false, HasFullComplex: false, HasD2ZeroCertificate: false, HasLeibnizCertificate: false, HasAlphaCompatibleD: false, DAlphaComputable: false, ContactVolumeComputable: false, ReebComputable: false, Verdict: join(StatusNoD2ZeroCertificate, StatusNoLeibnizCertificate, StatusNoAlphaSlotCompatibility, StatusNoDAlpha, StatusNoContactVolume, StatusNoReeb)}
}

func auditTime() ProductTimeFirewallAudit {
	return ProductTimeFirewallAudit{CochainToDM: false, CochainToLorentzianTime: false, CochainToOSPositivity: false, CochainToWickRotation: false, CochainToHilbert: false, CochainToHamiltonian: false, CochainToRGScale: false, CochainToArrowOfTime: false, ElectroweakBridgeStillSealed: true, Verdict: join(StatusNoPhysicalTime, StatusNoRGOSHilbert)}
}

func auditFinal(a Analysis) FinalVerdict {
	return FinalVerdict{K7Certified: a.K7.Dimension == 7, FormalExteriorDimensionsOnly: a.Exterior.HasAbstractExteriorDimensions && !a.Exterior.HasFiniteDOperator, UnsignedBooleanIncidenceFailsD2: !a.Boolean.D2ZeroForUnsignedIncidence, SignedFiniteDifferentialFound: false, K7CochainComplexFound: false, D2ZeroCertified: false, LeibnizCertified: false, DAlphaCertified: false, ContactVolumeCertified: false, ReebCertified: false, PhysicalTimeOrRGOpened: false, MissingNextTheorem: "Construct a native signed finite contact cochain complex over K_7 with explicit C^k(K_7), wedge product, differential d_k:C^k->C^{k+1}, d_{k+1}d_k=0, graded Leibniz rule, alpha in C^1(K_7), d alpha in C^2(K_7), and only then alpha wedge (d alpha)^3 and Reeb equations", Verdict: join(StatusUnsignedBooleanD2Nonzero, StatusNoD2ZeroCertificate, StatusNoLeibnizCertificate, StatusNoDAlpha, StatusGate569Firewall)}
}

func validate(a Analysis) error {
	failures := []string{}
	if a.K7.Dimension != 7 || a.K7.ExpectedDimension != 7 || math.Abs(a.K7.ContactIndex-1) > 1e-8 || a.K7.FrameIsometryResidual > 1e-8 || a.K7.BooleanContainmentResidual > 1e-8 || a.K7.G2ContainmentResidual > 1e-8 {
		failures = append(failures, "K7 certificate failed")
	}
	expected := []int{1, 7, 21, 35, 35, 21, 7, 1}
	if a.Exterior.VectorDimension != 7 || len(a.Exterior.GradeDimensions) != len(expected) || a.Exterior.TotalDimension != 128 || !a.Exterior.HasAbstractExteriorDimensions || a.Exterior.HasCertifiedK7CochainBasis || a.Exterior.HasCertifiedWedgeProductOnK7 || a.Exterior.HasFiniteDOperator {
		failures = append(failures, "formal exterior audit failed")
	} else {
		for i := range expected {
			if a.Exterior.GradeDimensions[i] != expected[i] {
				failures = append(failures, "formal R7 grade dimensions failed")
				break
			}
		}
	}
	if a.Boolean.M23Rows != 56 || a.Boolean.M23Cols != 28 || a.Boolean.M34Rows != 70 || a.Boolean.M34Cols != 56 || a.Boolean.Composition34After23Rows != 70 || a.Boolean.Composition34After23Cols != 28 || a.Boolean.Composition34After23Frobenius <= 0 || a.Boolean.Composition34After23MaxAbs <= 0 || !a.Boolean.UnsignedIncidence || a.Boolean.D2ZeroForUnsignedIncidence || a.Boolean.SignedOrientationAvailable || a.Boolean.DefinesK7Differential {
		failures = append(failures, "Boolean consecutive incidence audit failed")
	}
	if a.Restriction.K7CochainComplexDefined || a.Restriction.ProjectionFromAmbientFormsToK7Coforms || a.Restriction.PullbackDifferentialDefined || a.Restriction.D2ZeroOnRestrictedComplex {
		failures = append(failures, "restriction audit failed")
	}
	if a.Sources.G2CalibrationSuppliesComplex || a.Sources.ProjectorRelativePositionSuppliesBoundary || a.Sources.Q4SuppliesComplex {
		failures = append(failures, "source audit failed")
	}
	if a.Law.HasD0ToD1 || a.Law.HasD1ToD2 || a.Law.HasD2ToD3 || a.Law.HasFullComplex || a.Law.HasD2ZeroCertificate || a.Law.HasLeibnizCertificate || a.Law.HasAlphaCompatibleD || a.Law.DAlphaComputable || a.Law.ContactVolumeComputable || a.Law.ReebComputable {
		failures = append(failures, "differential law audit failed")
	}
	if a.Time.CochainToDM || a.Time.CochainToLorentzianTime || a.Time.CochainToOSPositivity || a.Time.CochainToWickRotation || a.Time.CochainToHilbert || a.Time.CochainToHamiltonian || a.Time.CochainToRGScale || a.Time.CochainToArrowOfTime || !a.Time.ElectroweakBridgeStillSealed {
		failures = append(failures, "time firewall failed")
	}
	if !a.Final.K7Certified || !a.Final.FormalExteriorDimensionsOnly || !a.Final.UnsignedBooleanIncidenceFailsD2 || a.Final.SignedFiniteDifferentialFound || a.Final.K7CochainComplexFound || a.Final.D2ZeroCertified || a.Final.LeibnizCertified || a.Final.DAlphaCertified || a.Final.ContactVolumeCertified || a.Final.ReebCertified || a.Final.PhysicalTimeOrRGOpened {
		failures = append(failures, "final verdict failed")
	}
	if len(failures) > 0 {
		return fmt.Errorf(strings.Join(failures, "; "))
	}
	return nil
}

func Statuses() []string {
	return []string{StatusK7CarrierInherited, StatusK7ContainmentCertified, StatusExteriorR7DimensionsAvailable, StatusNoK7CochainBasis, StatusNoK7WedgeProduct, StatusUnsignedBooleanConsecutiveMapsAvailable, StatusUnsignedBooleanD2Nonzero, StatusBooleanNotSignedCochainDifferential, StatusNoRestrictionToK7Complex, StatusNoD2ZeroCertificate, StatusNoLeibnizCertificate, StatusG2NoCochainComplex, StatusProjectorNoBoundary, StatusQ4NoCochainComplex, StatusNoAlphaSlotCompatibility, StatusNoDAlpha, StatusNoContactVolume, StatusNoReeb, StatusNoPhysicalTime, StatusNoRGOSHilbert, StatusGate569Firewall}
}

func truth(a Analysis) string {
	return join("Gate 569 attempts to upgrade the finite contact-differential search into the stricter cochain-complex test", "the abstract exterior dimensions for a 7-dimensional carrier are available, but no certified K_7 cochain basis, wedge product, or finite d operator is present", "consecutive unsigned Boolean incidence maps exist in the ambient R^8 grade ladder, but their compositions have nonzero Frobenius norm and therefore fail the d squared equals zero test", "no signed orientation, K_7 restriction, d^2 certificate, Leibniz certificate, alpha-compatible d, d alpha, contact volume, Reeb vector, product-time, RG, OS/Wick/Hilbert, or electroweak physical dynamics are opened")
}
func join(parts ...string) string { return strings.Join(parts, "; ") }
