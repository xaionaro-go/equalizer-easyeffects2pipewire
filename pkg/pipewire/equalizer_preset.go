package pipewire

import (
	"fmt"
	"io"
	"text/template"

	"github.com/xaionaro-go/equalizer-easyeffects2pipewire/pkg/equalizer"
)

type EqualizerPreset = equalizer.Preset

func WriteEqualizerPreset(w io.Writer, preset *EqualizerPreset) error {
	templateString := `
context.modules = [
    {
        name = libpipewire-module-filter-chain
        args = {
            node.description = "equalizer"
            media.name       = "equalizer"
            filter.graph = {
                nodes = [
{{- range .Nodes}}
                    {
                        type  = builtin
                        name  = {{ .Name }}
                        label = {{ .Label }}
                        control = { "Freq" = {{ .Frequency }} "Q" = {{ .Q }} "Gain" = {{ .Gain }} }
                    }
{{- end}}
                ]
                links = [
{{- range .Links}}
                    { output =  "{{ .NameOut }}:Out" input = "{{ .NameIn }}:In"  }
{{- end}}
                ]
            }
	    audio.channels = 2
	    audio.position = [ FL FR ]
            capture.props = {
                node.name   = "effect_input.eq"
                media.class = Audio/Sink
            }
            playback.props = {
                node.name   = "effect_output.eq"
                node.passive = true
            }
        }
    }
]
`
	type node struct {
		Name      string
		Label     string
		Frequency float64
		Q         float64
		Gain      float64
	}
	type link struct {
		NameOut string
		NameIn  string
	}
	type dataT struct {
		Nodes []node
		Links []link
	}

	data := dataT{}
	for idx, band := range preset.Bands {
		label := bandTypeToLabel(band.Type)
		if label == "" {
			return fmt.Errorf("unknown band type: %d", band.Type)
		}
		node := node{
			Name:      fmt.Sprintf("eq_band_%d", idx),
			Label:     label,
			Frequency: band.Frequency,
			Q:         band.Q,
			Gain:      band.Gain,
		}
		data.Nodes = append(data.Nodes, node)
		if len(data.Nodes) >= 2 {
			data.Links = append(data.Links, link{
				NameOut: data.Nodes[len(data.Nodes)-2].Name,
				NameIn:  node.Name,
			})
		}
	}

	template, err := template.New("preset-config").Parse(templateString)
	if err != nil {
		return fmt.Errorf("internal error: unable to parse the template: %w", err)
	}

	err = template.Execute(w, data)
	if err != nil {
		return fmt.Errorf("unable to execute the template: %w", err)
	}

	return nil
}

func bandTypeToLabel(bt equalizer.BandType) string {
	switch bt {
	case equalizer.BandTypeLowShelf:
		return "bq_lowshelf"
	case equalizer.BandTypeHighShelf:
		return "bq_highshelf"
	case equalizer.BandTypePeaking:
		return "bq_peaking"
	}
	return ""
}
