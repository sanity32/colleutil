package colleaction

type Action[TOpts, TResp any] string

type Request struct {
	Action    string `json:"action"`
	Options   any    `json:"options"`
	ColleData any    `json:"data"`
}

func (a Action[TOpts, TResp]) Str() string {
	return string(a)
}

// func (a Action[TOpts, TResp]) Waltz(c any, l waltz.Leader, opts TOpts) (colleresponse.Response[TResp], error) {
// 	// fn := collerequestwaltz.Adapter(l)
// 	req := Request{
// 		ColleData: c,
// 		Action:    a.Str(),
// 		Options:   opts,
// 	}
// 	if resp, err := waltz.LeaderAction(l, "colle", req, 0); err != nil {
// 		return nil, err
// 	} else {
// 		return colleresponse.Parse[TResp](resp)
// 	}
// }
