package bsector

import "testing"

func TestContactVacuumKernel(t *testing.T) {
	vacuum, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error: %v", err)
	}
	if got, want := vacuum.ZeroModeDimension(1e-8), vacuum.Contact.Dimension(); got != want {
		t.Fatalf("zero-mode dimension = %d, want contact dimension %d", got, want)
	}
	residual, err := vacuum.KernelEqualsContactResidual()
	if err != nil {
		t.Fatalf("KernelEqualsContactResidual() error: %v", err)
	}
	if residual > 1e-8 {
		t.Fatalf("kernel/contact residual too large: %.3e", residual)
	}
	if gap := vacuum.FirstPositiveEigenvalue(1e-8); gap <= 1e-8 {
		t.Fatalf("expected positive spectral gap, got %.12g", gap)
	}
}
