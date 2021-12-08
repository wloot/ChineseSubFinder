package sub_timeline_fixer

import (
	"github.com/allanpk716/ChineseSubFinder/internal/pkg/config"
	"testing"
)

func TestSubTimelineFixerHelperEx_Check(t *testing.T) {

	if NewSubTimelineFixerHelperEx(config.GetConfig().SubTimelineFixerConfig).Check() == false {
		t.Fatal("Need Install FFMPEG")
	}
}

func TestSubTimelineFixerHelperEx_Process(t *testing.T) {

	type args struct {
		videoFileFullPath string
		srcSubFPath       string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "R&M S05E01", args: args{
				videoFileFullPath: "X:\\连续剧\\瑞克和莫蒂 (2013)\\Season 5\\Rick and Morty - S05E01 - Mort Dinner Rick Andre WEBDL-1080p.mkv",
				srcSubFPath:       "C:\\WorkSpace\\Go2Hell\\src\\github.com\\allanpk716\\ChineseSubFinder\\internal\\logic\\sub_timeline_fixer\\CSF-SubFixCache\\Rick and Morty - S05E01 - Mort Dinner Rick Andre WEBDL-1080p\\R&M S05E01 - 简.ass"},
			wantErr: false,
		},
		{
			name: "Foundation (2021) - S01E09", args: args{
				videoFileFullPath: "X:\\连续剧\\Foundation (2021)\\Season 1\\Foundation (2021) - S01E09 - The First Crisis WEBDL-1080p.mkv",
				srcSubFPath:       "C:\\WorkSpace\\Go2Hell\\src\\github.com\\allanpk716\\ChineseSubFinder\\internal\\logic\\sub_timeline_fixer\\CSF-SubFixCache\\Foundation (2021) - S01E09 - The First Crisis WEBDL-1080p\\chinese(简英,zimuku).default.ass"},
			wantErr: false,
		},
		{
			name: "The Night House (2021)", args: args{
				videoFileFullPath: "X:\\TestMovie\\The Night House (2021)\\The Night House (2021) Bluray-1080p.mkv",
				srcSubFPath:       "X:\\TestMovie\\The Night House (2021)\\The Night House (2021) Bluray-1080p.chinese(简英,zimuku).ass"},
			wantErr: false,
		},
		{
			name: "龙猫 (1988)", args: args{
				videoFileFullPath: "X:\\电影\\龙猫 (1988)\\龙猫 (1988) 1080p DTS.mkv",
				srcSubFPath:       "X:\\电影\\龙猫 (1988)\\龙猫 (1988) 1080p DTS.chinese(简,zimuku).ass"},
			wantErr: false,
		},
		{
			name: "千与千寻 (2001)", args: args{
				videoFileFullPath: "X:\\电影\\千与千寻 (2001)\\千与千寻 (2001) 1080p Opus.mkv",
				srcSubFPath:       "X:\\电影\\千与千寻 (2001)\\千与千寻 (2001) 1080p Opus.chinese(简英,zimuku).ass"},
			wantErr: false,
		},
		{
			name: "Red Notice (2021)", args: args{
				videoFileFullPath: "X:\\电影\\Red Notice (2021)\\Red Notice (2021) WEBRip-1080p.mp4",
				srcSubFPath:       "X:\\电影\\Red Notice (2021)\\Red Notice (2021) WEBRip-1080p.chinese(简,xunlei).default.ass"},
			wantErr: false,
		},
		{
			name: "The Last Duel (2021)", args: args{
				videoFileFullPath: "X:\\电影\\The Last Duel (2021)\\The Last Duel (2021) WEBRip-1080p.mp4",
				srcSubFPath:       "X:\\电影\\The Last Duel (2021)\\The Last Duel (2021) WEBRip-1080p.chinese(简,shooter).default.srt"},
			wantErr: false,
		},
	}

	s := NewSubTimelineFixerHelperEx(config.GetConfig().SubTimelineFixerConfig)
	s.Check()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := s.Process(tt.args.videoFileFullPath, tt.args.srcSubFPath); (err != nil) != tt.wantErr {
				t.Errorf("Process() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}