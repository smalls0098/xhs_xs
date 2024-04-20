package xhs

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func Payload(str string) string {
	keys := []int32{187050025, 472920585, 186915882, 876157969, 255199502, 806945584, 220596020, 958210835, 757275681, 940378667, 489892883, 705504304, 354103316, 688857884, 890312192, 219096591, 622400037, 254088489, 907618332, 52759587, 907877143, 53870614, 839463457, 389417746, 975774727, 372382245, 437136414, 909246726, 168694017, 473575703, 52697872, 1010440969}
	enc := Des([]byte(base64.StdEncoding.EncodeToString([]byte(str))), keys, true)
	return hex.EncodeToString(enc)
}

func ProfileData(ts string) string {
	keys := []int32{187567141, 875696391, 170266120, 876222754, 188089115, 1010309137, 187054378, 957950720, 758514978, 941162813, 221382708, 990709537, 758848528, 688730163, 890444313, 722272792, 890962233, 252521496, 890843430, 185009704, 874317360, 119997734, 907612693, 119932961, 841824786, 120993794, 839716879, 909248796, 439099654, 372901635, 439091750, 1009915397}
	str := fmt.Sprintf(`{"x1":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36","x2":"false","x3":"zh-CN","x4":"30","x5":"8","x6":"30","x7":"Google Inc. (Intel Inc.),ANGLE (Intel Inc., Intel(R) UHD Graphics 630, OpenGL 4.1)","x8":"12","x9":"1792;1120","x10":"1792;1001","x11":"-480","x12":"Asia/Shanghai","x13":"true","x14":"true","x15":"true","x16":"false","x17":"true","x18":"un","x19":"MacIntel","x20":"1","x21":"Chrome PDF Plugin,Chrome PDF Viewer,Native Client","x22":"c2c1c8b9cdb6330e0fffb01097f01936","x23":"false","x24":"false","x25":"false","x26":"false","x27":"false","x28":"0,false,false","x29":"2,3,6,7,8","x30":"swf object not loaded","x33":"0","x34":"0","x35":"0","x36":"2","x37":"0|0|0|0|0|0|0|0|0|0|0|0|0","x38":"0|0|0|0|1|0|0|0|0|0|1|0|1|0|1|0","x39":"0","x40":"0","x41":"0","x42":"3.2.1","x43":"bd2c602c","x44":"%s","x45":"connecterror","x46":"false","x31":"124.04347657808103"}`, ts)
	enc := Des([]byte(base64.StdEncoding.EncodeToString([]byte(str))), keys, true)
	return hex.EncodeToString(enc)
}
