package ui

import (
	"slices"

	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/gopxl/pixel/v2/ext/text"
	"github.com/wirdos/util"
	"golang.org/x/image/font/basicfont"
)

// UI isn't rendered through the camera, this is to avoid it being transformed
// based on the camera's matrix, it should be simply directly rendered to the
// view based on fixed coordinates
type UI struct {
	components []UIComponent
	window *opengl.Window

	theme *Theme
}

func (ui *UI) Update() {
	i := 0
	for i < len(ui.components) {
		if ui.components[i].IsDestroyed() {
			ui.components = slices.Delete(ui.components, i, i+1)
			continue
		}

		ui.components[i].Update()
		i++
	}
}

// TODO: look into whether it's possible to batch here and return a list of batches
// instead of individual components which would require separate calls. It would be
// easy if the components only used sprites from the ui-elements files but they're
// also going to commonly write text. The problem there is that the text can't be
// written to the batch as it's not the same picture, if it's not part of the same
// batch it's likely there could be layering issues. To come back to after looking
// more into font atlases
func (ui *UI) Components() []UIComponent {
	return ui.components
}

func (ui *UI) AddComponent(c UIComponent) {
	ui.components = append(ui.components, c)
}

// TODO: theme is probably the wrong word right now but I like the idea of choosing
// different ui sprites based on a theme selected by the player
func (ui *UI) Theme() *Theme {
	return ui.theme
}

func NewUI(window *opengl.Window) (*UI, error) {
	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	palette, err := util.NewPalette("ui-elements")

	if err != nil {
		return nil, err
	}

	components := make([]UIComponent, 0)

	return &UI{
		window: window,
		theme: &Theme{
			Palette: palette,
			TextAtlas: atlas,
		},
		components: components,
	}, nil
}

type UIComponent interface {
	Update()
	Draw(pixel.Target)
	IsDestroyed() bool
}
