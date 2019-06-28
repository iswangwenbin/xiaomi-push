package xiaomipush

const (
	SandboxHost          = "https://sandbox.xmpush.xiaomi.com" // sandbox API只提供对IOS支持，不支持Android。
	ProductionHost       = "https://api.xmpush.xiaomi.com"
	GlobalProductionHost = "https://api.xmpush.global.xiaomi.com"
)

const (
	// 1 推送单条消息（支持多包名)
	// https://dev.mi.com/console/doc/detail?pId=1163#_0_0
	RegURL                               = "/v3/message/regid"        // 向某个regid或一组regid列表推送某条消息
	MessageAlisaURL                      = "/v3/message/alias"        // 根据alias，发送消息到指定设备上
	MessageUserAccountURL                = "/v2/message/user_account" // 根据account，发送消息到指定account上
	MessageMultiTopicURL                 = "/v2/message/topic"        // 根据topic，发送消息到指定一组设备上
	MultiPackageNameMessageMultiTopicURL = "/v3/message/multi_topic"  // 根据topic，发送消息到指定一组设备上
	MultiPackageNameMessageAllURL        = "/v3/message/all"          // 向所有设备推送某条消息

	// 2 推送多条消息
	// https://dev.mi.com/console/doc/detail?pId=1163#_1_0
	MultiMessagesRegIDURL       = "/v2/multi_messages/regids"        // 针对不同的regid推送不同的消息
	MultiMessagesAliasURL       = "/v2/multi_messages/aliases"       // 针对不同的aliases推送不同的消息
	MultiMessagesUserAccountURL = "/v2/multi_messages/user_accounts" // 针对不同的accounts推送不同的消息

	MessageAllURL = "/v2/message/all"         // 向所有设备推送某条消息
	MultiTopicURL = "/v3/message/multi_topic" // 向多个topic广播消息

	// 3 获取消息的统计数据
	// https://dev.mi.com/console/doc/detail?pId=1163#_2_0
	StatsURL = "/v1/stats/message/counters" // 统计push

	// 4 追踪消息状态
	// https://dev.mi.com/console/doc/detail?pId=1163#_3_0
	MessageStatusURL  = "/v1/trace/message/status"  // 获取指定ID的消息状态
	MessagesStatusURL = "/v1/trace/messages/status" // 获取某个时间间隔内所有消息的状态

	// 5 订阅/取消订阅标签
	// https://dev.mi.com/console/doc/detail?pId=1163#_4_0
	TopicSubscribeURL          = "/v2/topic/subscribe"         // 订阅RegId的标签
	TopicUnSubscribeURL        = "/v2/topic/unsubscribe"       // 取消订阅RegId的标签
	TopicSubscribeByAliasURL   = "/v2/topic/subscribe/alias"   // 订阅alias的标签
	TopicUnSubscribeByAliasURL = "/v2/topic/unsubscribe/alias" // 取消订阅alias的标签
	TopicUnSubscribeAllURL     = "v2/topic/unsubscribe/all"    // 取消某个标签的所有订阅

	// 6 获取失效的regId列表
	// https://dev.mi.com/console/doc/detail?pId=1163#_6_0
	InvalidRegIDsURL = "https://feedback.xmpush.xiaomi.com/v1/feedback/fetch_invalid_regids"

	// 7 获取一个应用的某个用户目前设置的所有Alias
	// https://dev.mi.com/console/doc/detail?pId=1163#_7_0
	AliasAllURL = "/v1/alias/all" // 获取一个应用的某个用户目前设置的所有Alias

	// 8 获取一个应用的某个用户目前订阅的所有Topic
	// https://dev.mi.com/console/doc/detail?pId=1163#_8_0

	TopicsAllURL = "/v1/topic/all" // 获取一个应用的某个用户的目前订阅的所有Topic
	// 9 获取一个userAccount对应的所有有效regId
	// https://dev.mi.com/console/doc/detail?pId=1163#_9_0
	RegidsAllByUserAccountURL = "v1/useraccount/get_regids_by_useraccount"

	// 10 检测/删除定时任务
	// https://dev.mi.com/console/doc/detail?pId=1163#_10
	ScheduleJobExistURL          = "/v2/schedule_job/exist"  // 检测定时消息的任务是否存在。
	ScheduleJobDeleteURL         = "/v2/schedule_job/delete" // 删除指定的定时消息。
	ScheduleJobDeleteByJobKeyURL = "/v3/schedule_job/delete" // 删除指定的定时消息。
)

var (
	PostRetryTimes = 3
)

// for future targeted push
var (
	BrandsMap = map[string]string{
		"品牌":    "MODEL",
		"小米":    "xiaomi",
		"三星":    "samsung",
		"华为":    "huawei",
		"中兴":    "zte",
		"中兴努比亚": "nubia",
		"酷派":    "coolpad",
		"联想":    "lenovo",
		"魅族":    "meizu",
		"HTC":   "htc",
		"OPPO":  "oppo",
		"VIVO":  "vivo",
		"摩托罗拉":  "motorola",
		"索尼":    "sony",
		"LG":    "lg",
		"金立":    "jinli",
		"天语":    "tianyu",
		"诺基亚":   "nokia",
		"美图秀秀":  "meitu",
		"谷歌":    "google",
		"TCL":   "tcl",
		"锤子手机":  "chuizi",
		"一加手机":  "1+",
		"中国移动":  "chinamobile",
		"昂达":    "angda",
		"邦华":    "banghua",
		"波导":    "bird",
		"长虹":    "changhong",
		"大可乐":   "dakele",
		"朵唯":    "doov",
		"海尔":    "haier",
		"海信":    "hisense",
		"康佳":    "konka",
		"酷比魔方":  "kubimofang",
		"米歌":    "mige",
		"欧博信":   "ouboxin",
		"欧新":    "ouxin",
		"飞利浦":   "philip",
		"维图":    "voto",
		"小辣椒":   "xiaolajiao",
		"夏新":    "xiaxin",
		"亿通":    "yitong",
		"语信":    "yuxin",
	}

	PriceMap = map[string]string{
		"0-999":     "0-999",
		"1000-1999": "1000-1999",
		"2000-3999": "2000-3999",
		"4000+":     "4000+",
	}
)
