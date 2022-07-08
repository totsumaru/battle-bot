package template

import "github.com/techstart35/battle-bot/discord/shared"

// ソロギミックのテンプレートです
var SoloTemplates = []string{
	"💥｜**%s** は間違えて自爆ボタンを押してしまった。ﾊﾞｰﾝ。",
	"💥｜**%s** はバナナの皮で滑って気絶。",
	"💥｜**%s** は迷子で行方不明に。",
	"💥｜**%s** は腐った肉を食べて腹痛。戦闘不能。",
	"💥｜**%s** は豆の食べ過ぎで破裂。おまめぇ。",
	"💥｜**%s** はMIRAKO.の可愛さで失神。",
	"💥｜**%s** は豆腐の角に頭をぶつけて即死。",
	"💥｜**%s** はタンスに小指をぶつけてショック死。",
	"💥｜**%s** は鳥のフンが頭に落ちてやる気を失う。脱落。",
	"💥｜**%s** は寝坊で試合に間に合わなかった。",
	"💥｜**%s** は盗んだバイクで走り出したが事故。",
	"💥｜**%s** は夏の暑さで溶けてしまった。",
	"💥｜**%s** は冬の寒さで凍ってしまった。",
	"💥｜**%s** はモンハンしすぎて夜ふか死。",
	"💥｜**%s** は嫁から鬼電、即帰宅。",
	"💥｜**%s** は秘孔つかれてあべ死。",
	"💥｜**%s** はカツラが飛んで追いかけて退場。",
	"💥｜**%s** はMIRAKOにお仕置きされてうれ死",
	"💥｜**%s** は八門遁甲を開門するも、対戦相手がいなかった。",
	"💥｜**%s** は九尾を抜かれて瀕死状態に。",
}

// バトルギミックのテンプレートです
var BattleTemplates = []string{
	"⚔️｜**%s** は **%s** をぼこ殴りにして倒した",
	"⚔️｜**%s** は **%s** をブロッコリーで殴った",
	"⚔️｜**%s** は **%s** を食べた",
	"⚔️｜**%s** は **%s** にタイキックをかました",
	"⚔️｜**%s** は **%s** を三輪車で轢いた",
	"⚔️｜**%s** は **%s** に千年殺しをお見舞いした。アーメン。",
	"⚔️｜**%s** は **%s** を落とし穴に落とした",
	"⚔️｜**%s** は **%s** に即効性の毒を盛った。さようなら。",
	"⚔️｜**%s** は **%s** を全力の膝カックンで倒した",
	"⚔️｜**%s** はラリアットで **%s**を倒した",
	"⚔️｜**%s** は豆鉄砲で **%s** を撃ち抜いた",
	"⚔️｜**%s** は **%s** をきゅうりで殴った",
	"⚔️｜**%s** は写輪眼を開眼。 **%s** を幻術に。",
	"⚔️｜**%s** は北斗百烈拳で **%s** を倒した。ほあたぁ!!",
	"⚔️｜南斗水鳥拳伝承者の **%s** は奥義 飛翔白麗で **%s** を倒した",
}

// ソロギミックのテンプレートをランダムに取得します
func GetRandomSoloTmpl() string {
	return SoloTemplates[shared.RandInt(1, len(SoloTemplates)+1)-1]
}

// バトルギミックのテンプレートをランダムに取得します
func GetRandomBattleTmpl() string {
	return BattleTemplates[shared.RandInt(1, len(BattleTemplates)+1)-1]
}
