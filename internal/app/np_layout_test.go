package app

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"

	"netease-music-mygo/internal/lyrics"
	"netease-music-mygo/internal/netease"
)

// The cover, the title and the progress bar of Now Playing begin on one edge,
// at every size of the window: the cover used to be centered in a column it
// did not fill, which moved it in from the title under it whenever the window
// was shorter than the cover's column.
func TestNowPlayingColumnsShareOneEdge(t *testing.T) {
	for _, size := range [][2]int{{880, 580}, {1180, 780}, {1440, 900}} {
		a := newTestApp()
		a.settings = defaultSettings()
		tr := netease.Track{ID: 7, Name: "春日影", Artists: []netease.Artist{{Name: "MyGO!!!!!"}}, Album: netease.Album{Name: "迷跡波"}, Duration: 199000}
		a.pl.Cue([]netease.Track{tr})
		a.pl.SetLyrics(lyrics.Build("[00:01.00]悴んだ心 ふるえる眼差し\n[00:05.00]世界で 僕は ひとりぼっちだった", "", ""))
		a.pl.Seek(6 * time.Second)
		a.nowPlaying, a.npPane = true, paneLyrics
		tt := ui.NewTester(a.view, size[0], size[1])
		tt.Advance(time.Second)
		tt.Frame()

		left := func(what string) float32 {
			r, ok := tt.Find(what)
			if !ok {
				t.Fatalf("%dx%d: no %q", size[0], size[1], what)
			}
			return r.X
		}
		cover, title, bar := left("专辑封面"), left("春日影"), left("播放进度")
		if cover != title || cover != bar {
			t.Errorf("%dx%d: cover at %v, title at %v, progress at %v", size[0], size[1], cover, title, bar)
		}
	}
}
