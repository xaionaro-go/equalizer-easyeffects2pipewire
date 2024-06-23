package easyeffects

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xaionaro-go/equalizer-easyeffects2pipewire/pkg/equalizer"
)

func TestParseEqualizerPreset(t *testing.T) {
	presetStr := `Preamp: -6.98 dB
Filter 1: ON LSC Fc 105.0 Hz Gain 2.1 dB Q 0.70
Filter 2: ON PK Fc 161.3 Hz Gain 0.7 dB Q 1.34
Filter 3: ON PK Fc 353.7 Hz Gain 1.1 dB Q 0.67
Filter 4: ON PK Fc 463.7 Hz Gain 0.6 dB Q 3.21
Filter 5: ON PK Fc 1205.3 Hz Gain -0.2 dB Q 2.29
Filter 6: ON PK Fc 2538.5 Hz Gain -3.8 dB Q 2.09
Filter 7: ON PK Fc 3193.8 Hz Gain -1.3 dB Q 4.67
Filter 8: ON PK Fc 3255.2 Hz Gain 1.3 dB Q 5.99

Filter 9: ON PK Fc 5470.7 Hz Gain -5.8 dB Q 3.55
Filter 10: ON HSC Fc 10000.0 Hz Gain 7.4 dB Q 0.70`

	preset, err := ParseEqualizerPreset(bytes.NewReader([]byte(presetStr)))
	require.NoError(t, err)
	require.NotNil(t, preset)
	require.Equal(t, &equalizer.Preset{Bands: []equalizer.Band{
		{Type: equalizer.BandTypeLowShelf, Frequency: 0, Q: 1, Gain: -6.98},
		{Type: equalizer.BandTypeLowShelf, Frequency: 105, Q: 0.7, Gain: 2.1},
		{Type: equalizer.BandTypePeaking, Frequency: 161.3, Q: 1.34, Gain: 0.7},
		{Type: equalizer.BandTypePeaking, Frequency: 353.7, Q: 0.67, Gain: 1.1},
		{Type: equalizer.BandTypePeaking, Frequency: 463.7, Q: 3.21, Gain: 0.6},
		{Type: equalizer.BandTypePeaking, Frequency: 1205.3, Q: 2.29, Gain: -0.2},
		{Type: equalizer.BandTypePeaking, Frequency: 2538.5, Q: 2.09, Gain: -3.8},
		{Type: equalizer.BandTypePeaking, Frequency: 3193.8, Q: 4.67, Gain: -1.3},
		{Type: equalizer.BandTypePeaking, Frequency: 3255.2, Q: 5.99, Gain: 1.3},
		{Type: equalizer.BandTypePeaking, Frequency: 5470.7, Q: 3.55, Gain: -5.8},
		{Type: equalizer.BandTypeHighShelf, Frequency: 10000, Q: 0.7, Gain: 7.4},
	}}, preset)
}
