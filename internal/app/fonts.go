package app

import (
	_ "embed"
	"log"

	"github.com/egoist/mygo/ui"
)

// The app draws in fonts of its own: Inter for the interface and the Latin
// of the lyrics, Noto Sans SC for Chinese. Neither ships with every
// system, macOS drawing Chinese in PingFang SC and Windows in Microsoft
// YaHei, so a subset of each comes with the app; what the subsets lack,
// rare characters, other scripts and emoji, the system's fonts draw, as
// Font's families do, and a weight between the faces draws in the one
// below.

//go:embed fonts/Inter-Regular.ttf
var interRegular []byte

//go:embed fonts/Inter-SemiBold.ttf
var interSemiBold []byte

//go:embed fonts/Inter-ExtraBold.ttf
var interExtraBold []byte

//go:embed fonts/NotoSansSC-Regular.ttf
var notoRegular []byte

//go:embed fonts/NotoSansSC-SemiBold.ttf
var notoSemiBold []byte

//go:embed fonts/NotoSansSC-ExtraBold.ttf
var notoExtraBold []byte

// uiFont is the interface's family: Inter, and Noto Sans SC for the Chinese
// names and titles it shows.
const uiFont = "Inter, Noto Sans SC"

// lyricsFont is the family the lyrics draw with: the same list, so that a
// line of Chinese set among English words keeps Inter's Latin.
const lyricsFont = uiFont

// lyricsTracking is how much room the lyrics leave between characters, in
// ems. Noto Sans SC draws its hanzi filling the whole em, and Inter's Latin
// keeps its own spacing, so the lyrics want none added.
const lyricsTracking = 0

// registerFonts adds the fonts the app ships to those text can use. It runs
// before any window, so that every text is laid out with them there.
func registerFonts() {
	for _, f := range []struct {
		data   []byte
		family string
	}{
		{interRegular, "Inter"},
		{interSemiBold, "Inter"},
		{interExtraBold, "Inter"},
		{notoRegular, "Noto Sans SC"},
		{notoSemiBold, "Noto Sans SC"},
		{notoExtraBold, "Noto Sans SC"},
	} {
		if err := ui.RegisterFont(f.data, f.family); err != nil {
			// The texts fall back to the system's fonts, as without the
			// subsets: the app is as it was before it had them.
			log.Printf("netease-music-mygo: the font of %s: %v", f.family, err)
			return
		}
	}
}
