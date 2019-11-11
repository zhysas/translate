package types

import (
	"encoding/json"
	"github.com/pkg/errors"
	"translate/model"
)

type Vmess struct {
	apiBase
	model.VmessSetting
}

func NewVmess(subLinks []string) api {
	return &Vmess{
		apiBase: apiBase{
			SubLinks: subLinks,
		},
	}
}

func (t *Vmess) Run() ([]*model.Setting, error) {
	err := t.getSub()
	if err != nil {
		return nil, errors.Wrap(err, "t.getSub")
	}
	ret := make([]*model.Setting, 0, len(t.Configs))
	for _, value := range t.Configs {
		var v2 *Vmess
		err := json.Unmarshal([]byte(value), &v2)
		if err != nil {
			return nil, errors.Wrap(err, "json.Unmarshal")
		}
		s, err := v2.ToSetting()
		if err != nil {
			return nil, errors.Wrap(err, "v2.ToSetting")
		}
		ret = append(ret, s)
	}
	return ret, nil
}

func (t *Vmess) ToSetting() (s *model.Setting, err error) {
	s = &model.Setting{
		Scheme:       model.Vmess,
		VmessSetting: t.VmessSetting,
	}
	return
}
