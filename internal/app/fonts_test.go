package app

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

// The app draws in the weights it asks for, in both fonts: registering a
// family as the variable font it is cut from gives one face, which draws
// every weight as its default, and the interface looks like nothing is
// wrong, only flat. Bolder faces cover more pixels.
func TestFontsHaveTheirWeights(t *testing.T) {
	registerFonts()
	registered := map[string]bool{}
	for _, f := range ui.FontFamilies() {
		if f.Registered {
			registered[f.Name] = true
		}
	}
	for _, family := range []string{"Inter", "Noto Sans SC"} {
		if !registered[family] {
			t.Fatalf("%s is not registered", family)
		}
	}
	covered := func(family, text string, weight int) int {
		tt := ui.NewTester(func(c *ui.Context) {
			ui.Box(c).Fill().Background(ui.RGB(0, 0, 0)).Children(func() {
				ui.Text(c, text).Font(family).FontSize(30).FontWeight(weight).TextColor(ui.RGB(255, 255, 255))
			})
		}, 900, 70)
		img := tt.Image()
		n := 0
		for i := 0; i < len(img.Pix); i += 4 {
			if img.Pix[i] > 8 {
				n++
			}
		}
		return n
	}
	for _, c := range []struct {
		family string
		text   string
	}{{"Inter", "Settings and Preferences"}, {"Noto Sans SC", "曾经的那些发生过的开心和难过"}} {
		light := covered(c.family, c.text, 400)
		bold := covered(c.family, c.text, 800)
		if light >= bold {
			t.Errorf("%s: weights 400 and 800 cover %d and %d pixels: one face draws both", c.family, light, bold)
		}
	}
	// A weight between the two must reach the face nearest it, not the
	// lightest: the interface asks for 600, which Noto Sans SC ships as a
	// face of its own.
	light, semi := covered("Noto Sans SC", "曾经的那些发生过的开心和难过", 400), covered("Noto Sans SC", "曾经的那些发生过的开心和难过", 600)
	if semi <= light {
		t.Errorf("Noto Sans SC: weight 600 covers %d pixels, no more than 400's %d: it draws light", semi, light)
	}
}
