package action

import "github.com/sanity32/colleutil/rect"

type MatchesOpts struct {
	Selector string `json:"sel"`
}

var (
	AddElement       Action[AddElementOpts, Mark]        = "addElement"
	Blur             Action[any, bool]                   = "blur"
	GetClassList     Action[any, []string]               = "getClassList"
	UpdateClass      Action[UpdateClasslistOpts, bool]   = "updateClassList"
	Click            Action[any, bool]                   = "click"
	Count            Action[any, int]                    = "count"
	Focus            Action[any, bool]                   = "focus"
	GetStyle         Action[any, string]                 = "getStyle"
	GetValue         Action[GetValueOpts, any]           = "getValue"
	SetStyle         Action[SetStyleRequest, bool]       = "setStyle"
	SetValue         Action[SetValueOpts, bool]          = "setValue"
	SetMark          Action[Mark, Mark]                  = "setMark"
	GetDataset       Action[GetValueOpts, *string]       = "getDataset"
	GetDatasetAll    Action[any, map[string]string]      = "getDatasetAll"
	SetDataset       Action[SetValueOpts, bool]          = "setDataset"
	DelDataset       Action[SetValueOpts, bool]          = "delDataset"
	LocationInfo     Action[any, LocationInfoResp]       = "locationInfo"
	Rect             Action[any, rect.Rect]              = "rect"
	Text             Action[any, string]                 = "text"
	ToViewport       Action[any, bool]                   = "toViewport"
	TimeoutVisualise Action[TimeoutVisualizerOpts, bool] = "visualizeTimeout"
	Matches          Action[MatchesOpts, bool]           = "matches"
)
