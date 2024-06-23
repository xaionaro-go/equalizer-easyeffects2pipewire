package easyeffects

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/xaionaro-go/equalizer-easyeffects2pipewire/pkg/equalizer"
)

type EqualizerPreset = equalizer.Preset

func ParseEqualizerPreset(r io.Reader) (*EqualizerPreset, error) {
	preset := EqualizerPreset{}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		words := strings.Split(scanner.Text(), " ")
		if len(words) < 2 {
			continue
		}
		switch words[0] {
		case "Preamp:":
			gain, err := strconv.ParseFloat(words[1], 64)
			if err != nil {
				return nil, fmt.Errorf("unable to parse preamp gain '%s': %w", words[1], err)
			}
			preset.Bands = append(preset.Bands, equalizer.Band{
				Type:      equalizer.BandTypeHighShelf,
				Frequency: 0,
				Q:         1,
				Gain:      gain,
			})
		case "Filter":
			bandTypeStr, frequencyStr, gainStr, qStr := words[3], words[5], words[8], words[11]
			bandType, err := parseBandType(bandTypeStr)
			if err != nil {
				return nil, fmt.Errorf("unable to parse band type '%s': %w", words[1], err)
			}

			frequency, err := strconv.ParseFloat(frequencyStr, 64)
			if err != nil {
				return nil, fmt.Errorf("unable to parse frequency '%s': %w", words[2], err)
			}

			gain, err := strconv.ParseFloat(gainStr, 64)
			if err != nil {
				return nil, fmt.Errorf("unable to parse gain '%s': %w", words[3], err)
			}

			q, err := strconv.ParseFloat(qStr, 64)
			if err != nil {
				return nil, fmt.Errorf("unable to parse q '%s': %w", words[4], err)
			}

			preset.Bands = append(preset.Bands, equalizer.Band{
				Type:      bandType,
				Frequency: frequency,
				Q:         q,
				Gain:      gain,
			})
		default:
			return nil, fmt.Errorf("unknown statement '%s'", words[0])
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("unable to get a line: %w", err)
	}

	return &preset, nil
}

func parseBandType(s string) (equalizer.BandType, error) {
	switch s {
	case "LSC":
		return equalizer.BandTypeLowShelf, nil
	case "HSC":
		return equalizer.BandTypeHighShelf, nil
	case "PK":
		return equalizer.BandTypePeaking, nil
	}
	return equalizer.UndefinedBandType, fmt.Errorf("unknown band-type: '%s'", s)
}
