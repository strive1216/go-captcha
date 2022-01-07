/**
 * @Author Awen
 * @Description Captcha
 * @Date 2021/7/18
 * @Email wengaolng@gmail.com
 **/

package captcha

import (
	"github.com/wenlng/go-captcha/captcha/assets"
	"golang.org/x/image/font"
)

// RangeVal is a type
/**
 * @Description: 范围值
 * @Example: {min: 0, max: 45} 从0-45中取任意值
 */
type RangeVal struct {
	Min, Max int
}

// Size is a type
/**
 * @Description: 尺寸
 * @Example: {width: 0, height: 45} 从0-45中取任意值
 */
type Size struct {
	Width, Height int
}

/**
 * @Description: 扭曲程度
 */
const (
	// 无扭曲
	DistortNone = iota
	// 扭曲程度 1-5级别
	DistortLevel1
	DistortLevel2
	DistortLevel3
	DistortLevel4
	DistortLevel5
)

// Config is a type
/**
* @Description: 验证码配置
 */
type Config struct {
	// 随机字符串长度范围
	rangTextLen 		RangeVal
	// 随机验证字符串长度范围, 注意：RangCheckTextLen < RangTextLen
	rangCheckTextLen 	RangeVal
	// 随机文本角度范围集合
	rangTexAnglePos 	[]RangeVal
	// 随机文本尺寸范围集合
	rangFontSize 		RangeVal
	// 随机缩略文本尺寸范围集合
	rangCheckFontSize 	RangeVal
	// 随机文本颜色	格式："#541245"
	rangFontColors 		[]string
	// 文本阴影偏移位置
	showTextShadow 		bool
	// 文本阴影颜色
	textShadowColor 	string
	// 文本阴影偏移位置
	textShadowPoint 	Point
	// 缩略图随机文本颜色	格式："#541245"
	rangThumbFontColors []string
	// 随机字体	格式：字体绝对路径字符串, /home/..../xxx.ttf
	rangFont 			[]string
	// 屏幕每英寸的分辨率
	fontDPI 			int
	// 随机验证码背景图		格式：图片绝对路径字符串, /home/..../xxx.png
	rangBackground 		[]string
	// 验证码尺寸, 注意：高度 > RangFontSize.max , 长度 > RangFontSize.max * RangFontSize.max
	imageSize 			Size
	// 图片清晰度 1-101
	imageQuality 		int
	// 验证码文本扭曲程度
	imageFontDistort 	int
	// 验证码文本透明度 0-1
	imageFontAlpha 		float64
	// 缩略图尺寸, 注意：高度 > RangCheckFontSize.max , 长度 > RangCheckFontSize.max * RangFontSize.max
	thumbnailSize  		Size
	// 字体Hinting
	fontHinting 		font.Hinting
	// 随机缩略背景图		格式：图片绝对路径字符串, /home/..../xxx.png
	rangThumbBackground []string
	// 缩略图背景随机色	格式："#541245"
	rangThumbBgColors 	[]string
	// 缩略图扭曲程度，值为 Distort...,
	thumbBgDistort 		int
	// 缩略图文字扭曲程度，值为 Distort...,
	thumbFontDistort 	int
	// 缩略图小圆点数量
	thumbBgCirclesNum 	int
	// 缩略图线条数量
	thumbBgSlimLineNum 	int
}


var chars = []string{"龙", "龟", "鼠", "龄", "齿", "齐", "鼻", "鼓", "鼎", "默", "黔", "黑", "黎", "黍", "黄", "麻", "麸", "麦", "鹿", "鹰", "鹦", "鹤", "鹏", "鹊", "鹉", "鹅", "鹃", "鸿", "鸽", "鸵", "鸳", "鸯", "鸭", "鸦", "鸥", "鸣", "鸡", "鸠", "鸟", "鳞", "鳖", "鳍", "鳄", "鲸", "鲫", "鲤", "鲜", "鲁", "鱼", "魔", "魏", "魄", "魂", "魁", "鬼", "鬓", "高", "髓", "骨", "骤", "骡", "骚", "骗", "骑", "骏", "验", "骇", "骆", "骄", "骂", "驾", "驼", "驻", "驹", "驶", "驴", "驳", "驱", "驰", "驯", "驮", "马", "香", "首", "馒", "馏", "馍", "馋", "馆", "馅", "馁", "饿", "饼", "饺", "饶", "饵", "饲", "饱", "饰", "饮", "饭", "饥", "餐", "食", "飞", "飘", "飒", "风", "颤", "颠", "额", "颜", "题", "颗", "颖", "颓", "频", "颊", "颈", "颇", "领", "颅", "预", "颂", "颁", "顿", "顾", "顽", "须", "顺", "项", "顷", "顶", "页", "韵", "音", "韭", "韩", "韧", "鞠", "鞍", "鞋", "靶", "靴", "革", "面", "靡", "靠", "非", "静", "靖", "青", "霹", "霸", "露", "霞", "霜", "霎", "霍", "霉", "震", "需", "雾", "雹", "雷", "零", "雳", "雪", "雨", "雕", "雏", "雌", "雇", "集", "雅", "雄", "雁", "雀", "难", "隶", "隧", "障", "隙", "隘", "隔", "隐", "随", "隆", "隅", "陷", "陶", "陵", "陪", "险", "陨", "除", "院", "陡", "陕", "限", "降", "陌", "陋", "陈", "陆", "际", "附", "阿", "阻", "阶", "阵", "阴", "阳", "防", "阱", "队", "阔", "阐", "阎", "阅", "阁", "阀", "闽", "闻", "闺", "闹", "闸", "闷", "间", "闲", "闰", "闯", "问", "闭", "闪", "门", "长", "镶", "镰", "镣", "镜", "镐", "镊", "镇", "镀", "锻", "锹", "锰", "锯", "键", "锭", "锨", "锦", "锥", "锤", "锣", "锡", "锚", "错", "锐", "锌", "锋", "锉", "锈", "锅", "锄", "锁", "销", "链", "铺", "铸", "银", "铲", "铭", "铣", "铡", "铝", "铜", "铛", "铐", "铆", "铅", "铃", "铁", "钾", "钻", "钳", "钱", "钮", "钩", "钧", "钦", "钥", "钢", "钠", "钟", "钞", "钝", "钙", "钓", "钉", "针", "鉴", "金", "量", "野", "重", "里", "释", "采", "醒", "醋", "醉", "醇", "酿", "酸", "酷", "酵", "酱", "酬", "酪", "酥", "酣", "酝", "酗", "酒", "配", "酌", "鄙", "都", "郭", "部", "郑", "郎", "郊", "郁", "邻", "邮", "邪", "邦", "那", "邢", "邓", "邑", "邀", "避", "遵", "遮", "遭", "遥", "遣", "遗", "道", "遏", "遍", "遇", "遂", "逾", "逼", "逻", "逸", "逮", "逢", "造", "速", "逞", "逝", "逛", "通", "逗", "途", "递", "逐", "透", "逊", "选", "逆", "逃", "适", "送", "退", "追", "迹", "迷", "述", "迫", "迟", "连", "违", "远", "进", "这", "还", "返", "近", "运", "迎", "迈", "过", "迅", "迄", "迂", "迁", "达", "辽", "边", "辱", "辰", "辫", "辩", "辨", "辣", "辟", "辞", "辜", "辛", "辙", "辖", "辕", "输", "辑", "辐", "辉", "辈", "辆", "辅", "较", "轿", "载", "轻", "轴", "轰", "软", "轮", "转", "轩", "轨", "轧", "车", "躺", "躲", "躯", "躬", "身", "躏", "躁", "蹲", "蹭", "蹬", "蹦", "蹋", "蹈", "蹄", "蹂", "踱", "踪", "踩", "踢", "踏", "踊", "跺", "跷", "践", "跳", "路", "跪", "跨", "跟", "距", "跛", "跑", "跌", "跋", "跃", "趾", "趴", "足", "趣", "趟", "趋", "越", "超", "趁", "起", "赶", "赵", "赴", "走", "赫", "赦", "赤", "赢", "赡", "赠", "赞", "赛", "赚", "赘", "赖", "赔", "赐", "赏", "赎", "赌", "赋", "赊", "资", "赃", "赂", "赁", "贿", "贾", "贼", "贺", "费", "贸", "贷", "贵", "贴", "贱", "贰", "贯", "贮", "购", "贬", "贫", "贪", "贩", "质", "货", "账", "败", "贤", "责", "财", "贡", "负", "贞", "贝", "貌", "豺", "豹", "豫", "豪", "象", "豌", "豆", "豁", "谷", "谴", "谱", "谭", "谬", "谨", "谦", "谤", "谣", "谢", "谜", "谚", "谓", "谒", "谐", "谎", "谍", "谋", "谊", "谈", "谆", "谅", "调", "谁", "课", "诽", "读", "诺", "诸", "请", "诵", "说", "诲", "诱", "误", "语", "诬", "诫", "详", "该", "询", "诡", "诞", "话", "诚", "诗", "试", "译", "词", "诊", "诉", "诈", "识", "诅", "评", "证", "诀", "访", "设", "讽", "讼", "论", "讹", "许", "讶", "讳", "讲", "记", "讯", "议", "训", "让", "讨", "讥", "认", "订", "计", "譬", "警", "誓", "誊", "誉", "言", "触", "解", "角", "觉", "览", "视", "觅", "规", "观", "见", "覆", "要", "西", "襟", "褪", "褥", "褒", "褐", "褂", "裹", "裸", "裳", "裤", "裙", "裕", "裆", "装", "裂", "裁", "袱", "袭", "被", "袜", "袖", "袒", "袍", "袋", "袄", "袁", "衷", "衰", "衬", "衫", "衩", "表", "补", "衣", "衡", "衙", "街", "衔", "衍", "行", "衅", "血", "蠢", "蠕", "蟹", "蟥", "蟋", "蟆", "蟀", "螺", "螟", "融", "螃", "蝶", "蝴", "蝠", "蝙", "蝗", "蝎", "蝌", "蝉", "蝇", "蜻", "蜡", "蜜", "蜘", "蜗", "蜕", "蜓", "蜒", "蜈", "蜂", "蜀", "蛾", "蛹", "蛮", "蛤", "蛛", "蛙", "蛔", "蛋", "蛉", "蛇", "蛆", "蛀", "蚯", "蚪", "蚤", "蚣", "蚜", "蚕", "蚓", "蚌", "蚊", "蚂", "蚁", "蚀", "虾", "虽", "虹", "虱", "虫", "虚", "虑", "虐", "虏", "虎", "蘸", "蘑", "藻", "藤", "藕", "藐", "藏", "薯", "薪", "薛", "薇", "薄", "蕾", "蕴", "蕊", "蕉", "蔽", "蔼", "蔬", "蔫", "蔚", "蔗", "蔓", "蔑", "蓬", "蓝", "蓖", "蓉", "蓄", "蒿", "蒸", "蒲", "蒜", "蒙", "蒋", "蒂", "葵", "葱", "葬", "葫", "董", "葡", "葛", "著", "落", "萨", "萧", "营", "萤", "萝", "萎", "萍", "萌", "萄", "菲", "菱", "菩", "菠", "菜", "菌", "菊", "菇", "莽", "莺", "莹", "获", "莲", "莱", "莫", "莉", "荸", "荷", "药", "荧", "荤", "荣", "荡", "荠", "荞", "荚", "荔", "荒", "荐", "草", "荆", "茸", "茶", "茵", "茴", "茬", "茫", "茧", "茎", "茉", "茅", "茄", "范", "茂", "茁", "苹", "英", "苫", "苦", "若", "苟", "苞", "苛", "苗", "苔", "苏", "苍", "苇", "芽", "芹", "芳", "花", "芯", "芭", "芬", "芦", "芥", "芝", "芜", "芙", "芒", "芍", "芋", "节", "艾", "艺", "艳", "色", "艰", "良", "艘", "艇", "船", "舷", "舶", "舵", "舱", "舰", "般", "航", "舟", "舞", "舔", "舒", "舍", "舌", "舆", "舅", "舀", "臼", "致", "至", "臭", "自", "臣", "臊", "臂", "臀", "膳", "膨", "膝", "膜", "膛", "膘", "膏", "膊", "膀", "腿", "腾", "腻", "腺", "腹", "腰", "腮", "腥", "腕", "腔", "腐", "腌", "腋", "腊", "脾", "脸", "脱", "脯", "脚", "脖", "脓", "脑", "脐", "脏", "脊", "脉", "脆", "脂", "能", "胸", "胶", "胳", "胰", "胯", "胧", "胡", "胞", "胜", "胚", "胖", "胎", "背", "胆", "胃", "胁", "胀", "肿", "肾", "肺", "肴", "育", "肯", "肮", "肪", "肩", "肥", "肤", "肢", "股", "肠", "肝", "肛", "肚", "肘", "肖", "肌", "肋", "肉", "肆", "肄", "肃", "聪", "聚", "聘", "联", "职", "聋", "聊", "聂", "耿", "耽", "耻", "耸", "耳", "耙", "耘", "耗", "耕", "耐", "耍", "而", "者", "考", "老", "耀", "翼", "翻", "翰", "翩", "翠", "翘", "翔", "翎", "翅", "翁", "羽", "羹", "群", "羡", "羞", "羔", "美", "羊", "署", "置", "罪", "罩", "罢", "罚", "罗", "罕", "网", "罐", "缺", "缸", "缴", "缰", "缭", "缩", "缨", "缤", "缠", "缝", "缚", "缘", "编", "缕", "缔", "缓", "缎", "缆", "缅", "缀", "绿", "绽", "综", "绸", "绷", "绵", "维", "绳", "绰", "续", "绪", "绩", "继", "绣", "绢", "统", "绞", "绝", "络", "给", "绘", "绕", "结", "绒", "绑", "经", "绎", "绍", "绊", "终", "织", "细", "绅", "组", "练", "线", "纽", "纺", "纹", "纸", "纷", "纵", "纳", "纲", "纱", "纯", "纬", "纫", "纪", "级", "约", "纤", "红", "纠", "繁", "絮", "累", "紫", "紧", "索", "素", "紊", "系", "糯", "糠", "糟", "糜", "糙", "糖", "糕", "糊", "精", "粹", "粱", "粮", "粪", "粥", "粤", "粟", "粘", "粗", "粒", "粉", "籽", "类", "米", "籍", "簿", "簸", "簇", "篷", "篱", "篮", "篡", "篙", "篓", "篇", "箱", "箭", "箫", "箩", "管", "算", "箕", "箍", "简", "签", "筹", "筷", "筝", "筛", "策", "答", "筒", "筑", "筐", "筏", "筋", "等", "笼", "第", "笨", "符", "笤", "笛", "笙", "笔", "笑", "笋", "笆", "竿", "竹", "端", "竭", "童", "竣", "章", "竟", "竞", "站", "竖", "立", "窿", "窥", "窟", "窝", "窜", "窘", "窗", "窖", "窒", "窑", "窍", "窄", "窃", "突", "穿", "空", "穷", "究", "穴", "穗", "穆", "稿", "稽", "稼", "稻", "稳", "稠", "稚", "税", "稍", "程", "稀", "秽", "移", "秸", "称", "积", "秫", "秩", "秧", "秦", "秤", "租", "秘", "秕", "秒", "科", "种", "秋", "秉", "秆", "秃", "私", "秀", "禾", "禽", "离", "福", "禁", "禀", "祸", "祷", "祭", "票", "祥", "祠", "祟", "神", "祝", "祖", "祈", "社", "礼", "示", "礁", "磷", "磨", "磕", "磅", "磁", "碾", "碴", "碳", "碱", "碰", "碧", "碟", "碘", "碗", "碑", "碎", "碍", "碌", "碉", "硼", "确", "硬", "硫", "硝", "硕", "硅", "础", "砾", "砸", "破", "砰", "砚", "砖", "研", "砍", "砌", "砂", "码", "矿", "矾", "石", "矮", "短", "矫", "矩", "知", "矢", "矛", "矗", "瞻", "瞳", "瞭", "瞬", "瞪", "瞧", "瞒", "瞎", "瞄", "睹", "睬", "睦", "督", "睡", "睛", "睁", "着", "眼", "眷", "眶", "眯", "眨", "眠", "真", "看", "眉", "省", "盾", "盼", "盹", "相", "直", "盲", "盯", "目", "盟", "盛", "盘", "盗", "盖", "盔", "盒", "监", "盐", "盏", "益", "盈", "盆", "盅", "皿", "皱", "皮", "皇", "皆", "的", "皂", "百", "白", "登", "癣", "癞", "癌", "瘾", "瘸", "瘫", "瘪", "瘩", "瘦", "瘤", "瘟", "痹", "痴", "痰", "痪", "痢", "痛", "痘", "痕", "痒", "痊", "症", "病", "疾", "疼", "疹", "疲", "疯", "疮", "疫", "疤", "疟", "疚", "疙", "疗", "疑", "疏", "疆", "畸", "畴", "番", "畦", "略", "畜", "留", "畔", "畏", "界", "畅", "画", "甸", "男", "电", "申", "甲", "由", "田", "甫", "甩", "用", "甥", "生", "甜", "甚", "甘", "瓷", "瓶", "瓮", "瓦", "瓤", "瓣", "瓢", "瓜", "璧", "璃", "瑰", "瑟", "瑞", "琼", "琴", "琳", "琢", "琐", "琉", "理", "琅", "球", "班", "珠", "珍", "珊", "玻", "玷", "玲", "现", "环", "玫", "玩", "玛", "玖", "王", "玉", "率", "玄", "猿", "猾", "猴", "献", "猬", "猫", "猪", "猩", "猜", "猛", "猖", "猎", "狼", "狸", "狱", "狰", "狮", "狭", "独", "狡", "狠", "狞", "狗", "狐", "狈", "狂", "犹", "状", "犯", "犬", "犁", "犀", "牺", "特", "牵", "牲", "物", "牧", "牢", "牡", "牛", "牙", "牍", "牌", "版", "片", "爽", "爹", "爸", "爷", "父", "爵", "爱", "爬", "爪", "爆", "燥", "燕", "燎", "燃", "熬", "熟", "熙", "熔", "熏", "熊", "熄", "煮", "照", "煤", "煞", "煎", "煌", "然", "焰", "焦", "焚", "焙", "焕", "焊", "烹", "热", "烫", "烧", "烦", "烤", "烟", "烛", "烙", "烘", "烈", "烂", "烁", "炼", "点", "炸", "炮", "炭", "炬", "炫", "炕", "炒", "炎", "炊", "炉", "灿", "灾", "灼", "灸", "灶", "灵", "灰", "灯", "灭", "火", "灌", "瀑", "濒", "激", "澳", "澡", "澜", "澎", "澈", "澄", "潮", "潭", "潦", "潜", "潘", "漾", "漱", "漫", "漩", "漠", "演", "漓", "漏", "漆", "漂", "滴", "滩", "滨", "滥", "滤", "满", "滞", "滚", "滔", "滓", "滑", "滋", "溺", "溶", "溯", "溪", "溢", "溜", "源", "溉", "溅", "溃", "湿", "湾", "湘", "湖", "湃", "渺", "游", "渴", "港", "温", "渤", "渣", "渡", "渠", "渗", "渔", "渐", "渊", "清", "添", "淹", "混", "淳", "深", "淮", "淫", "淤", "淡", "淘", "淑", "淌", "淋", "淆", "淀", "涵", "液", "涯", "涮", "涩", "涨", "涧", "润", "涤", "涣", "涡", "涝", "涛", "涕", "涎", "涌", "涉", "消", "涂", "浸", "海", "浴", "浮", "浪", "浩", "浦", "浙", "浓", "浑", "济", "测", "浊", "浇", "浆", "浅", "流", "派", "洽", "洼", "活", "洲", "洪", "津", "洞", "洛", "洗", "洒", "洋", "洁", "泽", "泼", "泻", "泵", "泳", "泰", "泪", "注", "泥", "泣", "波", "泡", "泞", "泛", "法", "泌", "泊", "泉", "泄", "沿", "沾", "沽", "沼", "治", "油", "沸", "河", "沮", "沫", "沪", "沧", "沦", "沥", "没", "沟", "沛", "沙", "沐", "沉", "沈", "沃", "汽", "汹", "汰", "汪", "汤", "污", "池", "江", "汞", "汛", "汗", "汉", "汇", "求", "汁", "永", "水", "氯", "氮", "氨", "氧", "氢", "氛", "气", "氓", "民", "氏", "毯", "毫", "毡", "毛", "毙", "毕", "比", "毒", "每", "母", "毅", "毁", "殿", "殷", "段", "殴", "殖", "残", "殊", "殉", "殃", "歼", "死", "歹", "歪", "歧", "武", "步", "此", "正", "止", "歌", "歉", "歇", "款", "欺", "欲", "欧", "欣", "欢", "次", "欠", "檬", "檩", "檐", "檀", "橱", "橡", "橙", "橘", "橄", "樱", "横", "模", "樟", "樊", "槽", "槐", "榴", "榨", "榜", "榛", "榕", "榔", "榆", "榄", "概", "楼", "楷", "楣", "楞", "楚", "楔", "椿", "椰", "椭", "椒", "椎", "植", "椅", "棺", "棵", "棱", "森", "棠", "棚", "棘", "棕", "棒", "棍", "棋", "棉", "检", "梳", "械", "梯", "梭", "梨", "梧", "梦", "梢", "梗", "梆", "梅", "梁", "桶", "桩", "桨", "桦", "桥", "档", "桑", "桐", "桌", "案", "框", "桅", "桃", "桂", "栽", "格", "根", "核", "样", "株", "校", "栗", "栖", "栓", "树", "栏", "栋", "栈", "标", "栅", "柿", "柴", "柳", "柱", "柬", "查", "柠", "柜", "柔", "染", "柒", "柑", "某", "柏", "柄", "枷", "架", "枯", "枫", "枪", "枣", "枢", "枝", "果", "枚", "林", "枕", "析", "枉", "构", "极", "板", "松", "杰", "杯", "杭", "杨", "来", "条", "杠", "束", "杜", "杖", "村", "材", "杏", "李", "杉", "杈", "杆", "权", "杂", "杀", "朽", "机", "朵", "朴", "朱", "术", "本", "末", "未", "木", "朦", "期", "朝", "望", "朗", "服", "朋", "有", "月", "最", "替", "曾", "曼", "曹", "更", "曲", "曙", "暴", "暮", "暗", "暖", "暑", "暇", "暂", "晾", "智", "晶", "晴", "晰", "景", "普", "晨", "晦", "晤", "晚", "晕", "晓", "晒", "晌", "晋", "晃", "显", "昼", "昵", "是", "昭", "昨", "昧", "春", "映", "星", "昙", "昔", "易", "昏", "明", "昌", "昆", "昂", "旺", "旷", "时", "旱", "旭", "旬", "早", "旨", "旧", "旦", "日", "既", "无", "旗", "族", "旋", "旅", "旁", "施", "方", "新", "斯", "断", "斩", "斧", "斥", "斤", "斟", "斜", "料", "斗", "斑", "斋", "文", "敷", "整", "敲", "数", "敬", "敦", "散", "敢", "敞", "敛", "教", "救", "敏", "敌", "效", "故", "政", "放", "攻", "改", "收", "支", "攘", "攒", "攀", "擦", "擒", "擎", "操", "擅", "擂", "撼", "撵", "撰", "撮", "播", "撬", "撩", "撤", "撞", "撕", "撒", "撑", "撇", "摹", "摸", "摩", "摧", "摘", "摔", "摊", "摇", "摆", "摄", "携", "搭", "搬", "搪", "搞", "搜", "搔", "搓", "搏", "搅", "搂", "搁", "搀", "揽", "援", "揭", "揪", "揩", "揣", "握", "揖", "插", "提", "描", "揍", "揉", "掺", "掸", "掷", "掰", "措", "掩", "推", "控", "接", "探", "掠", "掘", "掖", "排", "掐", "掏", "掌", "掉", "授", "掂", "掀", "捻", "捺", "捷", "捶", "据", "捧", "捣", "换", "捡", "损", "捞", "捕", "捐", "捏", "捎", "捍", "捌", "捉", "捆", "捅", "捂", "挽", "挺", "振", "挫", "挪", "挨", "挥", "挤", "挣", "挡", "挠", "挟", "挚", "挖", "挑", "挎", "按", "指", "挂", "持", "拿", "拾", "拼", "拷", "拴", "拳", "拱", "拯", "拭", "括", "择", "拨", "拧", "拦", "拥", "拣", "拢", "拟", "拜", "招", "拙", "拘", "拗", "拖", "拔", "拓", "拒", "拐", "拍", "拌", "拉", "拇", "拆", "担", "拄", "拂", "抽", "押", "抹", "抵", "抱", "抬", "披", "报", "护", "抢", "抡", "抠", "抛", "抚", "折", "抗", "抖", "投", "抓", "抒", "抑", "把", "抄", "技", "承", "找", "扼", "批", "扶", "扳", "扰", "扯", "扮", "扭", "扬", "扫", "扩", "执", "扣", "扛", "托", "扔", "打", "扒", "扑", "扎", "才", "手", "扇", "扁", "所", "房", "户", "戴", "戳", "截", "戚", "战", "或", "戒", "我", "成", "戏", "戈", "懦", "懒", "懊", "懈", "懂", "憾", "憨", "憔", "憎", "憋", "慷", "慰", "慨", "慧", "慢", "慕", "慎", "慌", "慈", "愿", "愧", "愤", "感", "愚", "愕", "意", "愉", "愈", "愁", "惹", "惶", "想", "惰", "惯", "惭", "惫", "惩", "惨", "惧", "惦", "惠", "惜", "惕", "惑", "惋", "惊", "情", "悼", "悴", "悲", "悯", "悬", "您", "悦", "患", "悠", "悟", "悔", "悍", "悉", "悄", "恼", "恶", "恳", "恰", "息", "恭", "恬", "恩", "恨", "恤", "恢", "恕", "恒", "恐", "恍", "恋", "恃", "总", "怯", "怪", "怨", "性", "急", "怠", "思", "怜", "怖", "怕", "怔", "怒", "怎", "态", "怀", "忿", "忽", "念", "忱", "快", "忧", "忠", "忙", "忘", "志", "忍", "忌", "忆", "必", "心", "徽", "德", "微", "循", "御", "徙", "徘", "得", "徒", "徐", "律", "徊", "很", "待", "径", "征", "往", "彼", "彻", "役", "影", "彰", "彭", "彬", "彪", "彩", "彤", "形", "录", "当", "归", "强", "弹", "弱", "弯", "弧", "弦", "弥", "张", "弟", "弛", "引", "弓", "式", "弊", "弄", "弃", "异", "开", "建", "廷", "延", "廓", "廊", "廉", "庸", "康", "庶", "庵", "庭", "座", "度", "废", "庞", "府", "庙", "店", "底", "应", "库", "庐", "序", "床", "庇", "庆", "庄", "广", "幽", "幼", "幻", "幸", "并", "年", "平", "干", "幢", "幕", "幔", "幌", "幅", "帽", "常", "帮", "席", "带", "帝", "帜", "帚", "帘", "帖", "帕", "帐", "希", "师", "帆", "帅", "布", "市", "币", "巾", "巷", "巴", "已", "己", "差", "巫", "巩", "巨", "巧", "左", "工", "巢", "巡", "州", "川", "巍", "嵌", "崭", "崩", "崖", "崔", "崎", "崇", "峻", "峰", "峭", "峦", "峡", "岸", "岳", "岭", "岩", "岛", "岗", "岖", "岔", "岂", "岁", "屿", "屹", "山", "屯", "履", "屡", "屠", "属", "展", "屑", "屏", "屎", "屋", "届", "屉", "屈", "居", "层", "屁", "局", "尿", "尾", "尽", "尼", "尺", "尸", "就", "尤", "尝", "尚", "尘", "尖", "尔", "少", "小", "尊", "尉", "将", "射", "封", "寿", "导", "寻", "寺", "对", "寸", "寨", "寥", "寡", "察", "寞", "寝", "寓", "寒", "富", "寇", "密", "寄", "寂", "宿", "宾", "宽", "容", "家", "宵", "宴", "害", "宰", "宫", "宪", "宦", "室", "宣", "客", "审", "宠", "实", "宝", "宜", "宛", "定", "宙", "官", "宗", "宏", "完", "宋", "安", "守", "宇", "宅", "它", "宁", "孽", "孵", "孩", "学", "孤", "季", "孟", "孝", "孙", "存", "字", "孕", "孔", "子", "嬉", "嫩", "嫡", "嫌", "嫉", "嫂", "嫁", "媳", "媚", "媒", "婿", "婶", "婴", "婚", "婉", "婆", "娶", "娱", "娩", "娜", "娘", "娇", "娄", "娃", "威", "姿", "姻", "姨", "姥", "姜", "姚", "委", "姓", "姑", "姐", "始", "姊", "姆", "妻", "妹", "妨", "妥", "妙", "妖", "妓", "妒", "妈", "妇", "妆", "妄", "如", "好", "她", "奸", "奶", "奴", "女", "奥", "奢", "奠", "套", "奖", "奕", "奔", "契", "奏", "奋", "奉", "奈", "奇", "奄", "夺", "夹", "夸", "夷", "头", "失", "夯", "央", "夭", "夫", "太", "天", "大", "够", "夜", "多", "外", "夕", "夏", "复", "备", "处", "壹", "壶", "壳", "声", "壮", "士", "壤", "壕", "壁", "墩", "墨", "增", "墙", "墓", "墅", "境", "填", "塞", "塘", "塔", "塑", "塌", "堵", "堰", "堪", "堤", "堡", "堕", "堆", "堂", "基", "培", "埠", "域", "城", "埋", "埃", "埂", "垮", "垫", "垦", "垢", "垛", "垒", "型", "垄", "垃", "垂", "坷", "坯", "坪", "坦", "坤", "坡", "坠", "坟", "坞", "坝", "坛", "坚", "块", "坑", "坐", "坏", "坎", "坊", "均", "址", "圾", "场", "地", "在", "圣", "土", "圈", "圆", "圃", "图", "国", "固", "围", "囱", "困", "园", "囤", "团", "因", "回", "四", "囚", "囊", "嚼", "嚷", "嚣", "嚎", "噪", "噩", "器", "嘿", "嘹", "嘶", "嘴", "嘲", "嘱", "嘉", "嘁", "嘀", "嗽", "嗦", "嗤", "嗡", "嗜", "嗓", "嗅", "喻", "喷", "喳", "喧", "喝", "喜", "喘", "喊", "喉", "喇", "善", "喂", "啼", "啸", "啰", "啦", "啥", "啤", "啡", "啊", "商", "啄", "啃", "唾", "唱", "唯", "售", "唬", "唧", "唤", "唠", "唐", "唉", "唇", "唆", "唁", "哼", "哺", "哲", "哮", "哭", "哪", "哩", "哨", "哥", "哟", "哗", "哑", "哎", "响", "哈", "哆", "哄", "品", "哀", "咽", "咸", "咳", "咱", "咬", "咪", "咨", "咧", "咙", "咖", "咕", "咒", "咐", "咏", "和", "咆", "命", "呼", "呻", "呵", "味", "周", "呢", "呜", "呛", "员", "呕", "呐", "告", "呈", "呆", "呀", "吼", "吻", "吹", "吸", "吵", "吴", "吱", "启", "吮", "吭", "听", "含", "吩", "吨", "吧", "否", "吠", "吟", "吞", "吝", "君", "吗", "吕", "吓", "向", "吐", "吏", "后", "名", "同", "吊", "吉", "合", "吆", "各", "吃", "吁", "叽", "叼", "叹", "司", "号", "叶", "右", "史", "台", "可", "叮", "叭", "召", "叫", "只", "叨", "另", "句", "古", "口", "叠", "叛", "叙", "变", "受", "取", "叔", "发", "反", "双", "友", "及", "叉", "又", "参", "叁", "县", "去", "厨", "厦", "厢", "原", "厚", "厘", "厕", "厌", "压", "厉", "历", "厅", "厂", "卿", "卸", "卷", "卵", "却", "即", "危", "印", "卫", "卧", "卦", "卤", "卢", "卡", "占", "卜", "博", "南", "卖", "单", "卓", "卒", "卑", "协", "华", "半", "午", "升", "千", "十", "匿", "匾", "医", "区", "匹", "匪", "匣", "匠", "匙", "北", "化", "匕", "匈", "匆", "包", "匀", "勿", "勾", "勺", "勤", "募", "勘", "勒", "勋", "勉", "勇", "勃", "势", "劳", "劲", "励", "劫", "努", "助", "动", "劣", "务", "加", "功", "办", "劝", "力", "劈", "剿", "割", "副", "剪", "剩", "剧", "剥", "剖", "剔", "剑", "前", "削", "剃", "剂", "刽", "刻", "刺", "刹", "券", "刷", "制", "到", "刮", "别", "利", "刨", "判", "删", "初", "创", "刚", "则", "刘", "列", "划", "刑", "刊", "切", "分", "刃", "刁", "刀", "凿", "函", "击", "出", "凹", "凸", "凶", "凳", "凰", "凯", "凭", "凫", "凤", "凡", "几", "凝", "凛", "凑", "减", "凌", "凉", "准", "凄", "净", "冻", "冷", "冶", "况", "决", "冲", "冰", "冯", "冬", "冤", "冠", "农", "军", "写", "冗", "冕", "冒", "再", "册", "冈", "内", "冀", "兽", "兼", "养", "典", "具", "其", "兵", "兴", "关", "共", "兰", "六", "公", "八", "全", "入", "兢", "兜", "党", "兔", "兑", "免", "克", "光", "先", "兆", "充", "兄", "元", "允", "儿", "儡", "儒", "僻", "僵", "僧", "僚", "像", "傻", "傲", "催", "储", "傍", "傅", "傀", "偿", "偷", "偶", "健", "停", "做", "偏", "偎", "假", "倾", "值", "债", "倦", "倡", "借", "倚", "候", "倘", "倔", "倒", "倍", "俺", "俱", "俯", "修", "俭", "俩", "信", "保", "俘", "俗", "俐", "俏", "俊", "俄", "促", "便", "侵", "侯", "侮", "侨", "侧", "侦", "侥", "侣", "侠", "依", "供", "侍", "例", "侈", "侄", "使", "佳", "佩", "佣", "你", "作", "佛", "余", "何", "体", "佑", "住", "低", "位", "但", "佃", "似", "伺", "伸", "伶", "伴", "估", "伯", "伪", "伦", "伤", "传", "伟", "伞", "会", "伙", "优", "众", "休", "伐", "伏", "伍", "伊", "企", "仿", "份", "任", "价", "件", "仲", "仰", "们", "仪", "以", "令", "代", "仙", "付", "仗", "他", "仔", "仓", "仑", "从", "仍", "介", "今", "仇", "仆", "仅", "仁", "什", "亿", "人", "亲", "亮", "亭", "京", "享", "亩", "产", "亦", "亥", "交", "亡", "些", "亚", "井", "五", "互", "云", "亏", "于", "二", "事", "争", "予", "了", "乾", "乳", "乱", "买", "书", "乡", "习", "也", "乞", "九", "乙", "乘", "乖", "乔", "乓", "乒", "乐", "乏", "乎", "乍", "乌", "之", "义", "么", "久", "乃", "举", "丽", "主", "为", "丹", "丸", "临", "串", "丰", "中", "个", "丧", "严", "两", "丢", "丝", "东", "丛", "业", "丙", "丘", "世", "且", "专", "丑", "丐", "与", "不", "下", "上", "三", "丈", "万", "七", "丁"}
var textColors = []string{
	"#fde98e",
	"#60c1ff",
	"#fcb08e",
	"#fb88ff",
	"#b4fed4",
	"#cbfaa9",
}
var thumbTextColors = []string{
	"#006600",
	"#005db9",
	"#aa002a",
	"#875400",
	"#6e3700",
	"#660033",
}
var textShadowColor = "#101010"

// GetCaptchaDefaultChars is a type
/**
 * @Description: 获取字符
 * @return map[int]string
 */
func GetCaptchaDefaultChars() *[]string {
	return &chars
}

// GetCaptchaDefaultConfig is a type
/**
 * @Description: 获取默认配置
 * @return CaptchaConfig
 */
func GetCaptchaDefaultConfig() *Config {
	return &Config{
		rangTextLen:      		RangeVal{6, 7},
		rangCheckTextLen: 		RangeVal{2, 4},
		rangTexAnglePos: 		[]RangeVal{
			{20, 35},
			{35, 45},
			{45, 60},
			{290, 305},
			{305, 325},
			{325, 330},
		},
		rangFontSize:       	RangeVal{30, 38},
		fontDPI:            	72,
		rangCheckFontSize:  	RangeVal{24, 30},
		imageFontDistort:   	DistortNone,
		imageFontAlpha:     	1,
		rangFontColors:     	getDefaultTextColors(),
		showTextShadow:    		true,
		textShadowColor:    	getDefaultTextShadowColor(),
		textShadowPoint:     	Point{-1, -1},
		rangThumbFontColors:    getDefaultThumbTextColors(),
		fontHinting: 			font.HintingNone,
		imageSize:          	Size{300, 240},
		imageQuality: 			999,
		thumbnailSize:      	Size{150, 40},
		rangThumbBgColors:  	getDefaultThumbTextColors(),
		thumbFontDistort:   	DistortLevel3,
		thumbBgDistort:     	DistortLevel4,
		thumbBgCirclesNum:  	24,
		thumbBgSlimLineNum: 	2,

		rangFont: 				assets.DefaultBinFontList(),
		rangBackground: 		assets.DefaultBinImageList(),
	}
}

/**
 * @Description: 获取默认文本颜色
 * @return []string
 */
func getDefaultTextColors() []string {
	return textColors
}


/**
 * @Description: 获取默认阴影文本颜色
 * @return string
 */
func getDefaultTextShadowColor() string {
	return textShadowColor
}

/**
 * @Description: 获取默认缩略图文本颜色
 * @return []string
 */
func getDefaultThumbTextColors() []string {
	return thumbTextColors
}