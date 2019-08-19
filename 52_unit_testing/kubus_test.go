package main

import (
	"testing"
)

var (
	Kubus               kubus   = kubus{4}
	VolumeSeharusnya    float64 = 64
	LuasSeharusnya      float64 = 96
	KelilingiSeharusnya float64 = 48
)

func TestHitungVolume(t *testing.T) {
	t.Logf("Volume : %.2f", Kubus.Volume())

	if Kubus.Volume() != VolumeSeharusnya {
		t.Errorf("Salah harusnya %.2f", VolumeSeharusnya)
	}
}

func TestHitungLuas(t *testing.T) {
	t.Logf("Luas : %.2f", Kubus.Luas())

	if Kubus.Luas() != LuasSeharusnya {
		t.Errorf("Salah harusnya %.2f", LuasSeharusnya)
	}
}

func TestHitungKelilings(t *testing.T) {
	t.Logf("Luas : %.2f", Kubus.keliling())

	if Kubus.keliling() != KelilingiSeharusnya {
		t.Errorf("Salah harusnya %.2f", KelilingiSeharusnya)
	}
}
