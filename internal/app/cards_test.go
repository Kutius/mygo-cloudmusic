package app

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"

	"netease-music-mygo/internal/netease"
)

// A card whose name takes two lines does not push the subtitles of the cards
// beside it down: a grid stretches the cells of a row as tall as the tallest,
// and a name left to its own height moved the line under it whenever it
// wrapped.
func TestCardSubtitlesShareOneLine(t *testing.T) {
	a := newTestApp()
	a.settings = defaultSettings()
	names := []string{"短名", "一个相当长到需要折行的歌单名字在这里", "短名"}
	for i, name := range names {
		a.recPlaylists = append(a.recPlaylists, netease.Playlist{ID: int64(100 + i), Name: name,
			UpdateFreq: "更新" + string(rune('1'+i))})
	}
	a.charts = []netease.Playlist{{ID: 9, Name: "飙升榜", UpdateFreq: "每天更新"}}
	a.homeLoading = false
	tt := ui.NewTester(a.view, 900, 800)
	tt.Advance(time.Second)
	tt.Frame()

	first := float32(-1)
	for i := range names {
		sub, ok := tt.Find("更新" + string(rune('1'+i)))
		if !ok {
			t.Fatalf("card %d: no subtitle", i)
		}
		if first < 0 {
			first = sub.Y
			continue
		}
		if sub.Y != first {
			t.Errorf("card %d: its subtitle is at %v, the first card's at %v: a wrapped name moved it", i, sub.Y, first)
		}
	}
}
