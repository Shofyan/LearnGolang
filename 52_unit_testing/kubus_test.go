package main

import (
	//"testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	Kubus               kubus   = kubus{4}
	VolumeSeharusnya    float64 = 64
	LuasSeharusnya      float64 = 96
	KelilingiSeharusnya float64 = 48
)

//
//func TestHitungVolume(t *testing.T) {
//	t.Logf("Volume : %.2f", Kubus.Volume())
//
//	if Kubus.Volume() != VolumeSeharusnya {
//		t.Errorf("Salah harusnya %.2f", VolumeSeharusnya)
//	}
//}
//
//func TestHitungLuas(t *testing.T) {
//	t.Logf("Luas : %.2f", Kubus.Luas())
//
//	if Kubus.Luas() != LuasSeharusnya {
//		t.Errorf("Salah harusnya %.2f", LuasSeharusnya)
//	}
//}
//
//func TestHitungKelilings(t *testing.T) {
//	t.Logf("Luas : %.2f", Kubus.keliling())
//
//	if Kubus.keliling() != KelilingiSeharusnya {
//		t.Errorf("Salah harusnya %.2f", KelilingiSeharusnya)
//	}
//}

func TesthitungVolume2(t *testing.T) {
	assert.Equal(t, Kubus.Volume(), VolumeSeharusnya, "perhitungannya salah")
}

func TesthitungLuas2(t *testing.T) {
	assert.Equal(t, Kubus.Luas(), LuasSeharusnya, "perhitungannya salah")
}
func TesthitungKeliling2(t *testing.T) {
	assert.Equal(t, Kubus.keliling(), KelilingiSeharusnya, "perhitungannya salah")
}
