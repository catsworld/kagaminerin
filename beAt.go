package main

import (
	"github.com/catsworld/api"
	"github.com/catsworld/botmaid"
	"github.com/catsworld/random"
)

var (
	wordBeAt = []string{
		"i'a、i'a……呵呵呵呵",
		"我既已知晓了『门』。汝，肯定无法看见。",
		"啊……再等一下。很快就结束啰……",
		"好像很有趣呢。",
		"我即为禁忌的秘药，是为引导者。",
		"『门』啊，打开吧。",
		"Ygnaiih…ygnaiih…thflthkhk'ngha……白银之钥握于我手，自虚无之处现身，将其指尖伸向此处吧。我等父神啊。现在，我即为寄宿其神髓之化身。穿越蔷薇沉睡之海，终抵穷极之门！『光壳满布的虚树(Qlipha Rhizome)』……",
		"这就是，身为塞勒姆魔女的我。还请不要害怕喔。",
		"能感受到神的气息呢。",
		"已经，再也压抑不住了呢……这份联系着此世全部的『门』的伟大力量。",
		"今天要打开通往何处的『门』呢。",
		"我啊，很害怕狗，怕的都会发抖呢。不但会吠叫，还会咬住人不放，甚至，好像还会让人闭上嘴似的。",
		"像这样能成为从者也一定是有意义的吧。我不会放弃的！希望能带给这个世界，安详以及信赖。",
		"在这座堡垒里，感觉我就像是外来人一样……之前，原本还很担心能不能适应，紧张的心脏都快跳出来了……",
		"你好啊。说是不寂寞，大概是骗人的吧。不过，我也已经习惯这里的生活了喔。以前虽然被恐怖的人吓得发抖，但是大家其实都是好人呢！",
	}
	wordMasterAt = []string{
		"i'a、i'a……呵呵呵呵……",
		"我既已知晓了『门』。汝，肯定无法看见。",
		"啊……再等一下。很快就结束啰……",
		"好像很有趣呢",
		"还不够喔",
		"把你的魔力，灌注给我。",
		"我即为禁忌的秘药，是为引导者。",
		"Ygnaiih…ygnaiih…thflthkhk'ngha……白银之钥握于我手，自虚无之处现身，将其指尖伸向此处吧。我等父神啊。现在，我即为寄宿其神髓之化身。穿越蔷薇沉睡之海，终抵穷极之门！『光壳满布的虚树(Qlipha Rhizome)』……",
		"再来再来，继续取悦我吧。",
		"请为我献上更多吧。",
		"这就是，身为塞勒姆魔女的我。还请不要害怕喔，Master。",
		"能感受到神的气息呢。",
		"已经，再也压抑不住了呢……这份联系着此世全部的『门』的伟大力量。不论是要拯救，还是破坏世界，都在Master的一念之间喔。",
		"来吧Master！快点过来我的世界！享受这个境界的视野！是你的话一定……可以承受住喔。",
		"今天要打开通往何处的『门』呢。",
		"我将会成为你的钥匙喔。",
		"用尽力气握紧我的手吧。那样就好了。",
		"虽然松饼也不错，但是Master有喝过酒吗。因为对身体不好，所以一定很美味吧。真好，真好呢……。",
		"Master，我啊，很害怕狗，怕的都会发抖呢。不但会吠叫，还会咬住人不放，甚至，好像还会让人闭上嘴似的。",
		"像这样能成为从者也一定是有意义的吧。我不会放弃的！希望能带给这个世界，安详以及信赖。",
		"在这座堡垒里，感觉我就像是外来人一样……之前，原本还很担心能不能适应，紧张的心脏都快跳出来了……唉呀，可以跟我一起祈祷吗？真是谢谢你，Master！",
		"你好啊，Master。说是不寂寞，大概是骗人的吧。不过，我也已经习惯这里的生活了喔。以前虽然被恐怖的人吓得发抖，但是大家其实都是好人呢！",
		"Master，Master！我想要，帮你更多的忙！最近每天都很快乐呢！你的喜悦，也就是我的喜悦喔！",
		"呼……如果Master没有好好看着我的话，我好像，就会变成坏孩子了……请不要离我太远喔，Master。",
		"不管是迷惘的时候，或是因为孤独而发抖的时候，你都会陪在我身边。下次，我发誓一定会在一旁支持着你喔。不过呢，希望有时候也能让我撒、撒个娇吧。",
	}
)

func init() {
	bm.AddCommand(botmaid.Command{
		Do:       masterAt,
		Priority: 5,
	})
	bm.AddCommand(botmaid.Command{
		Do:       beAt,
		Priority: 5,
	})
}

func beAt(e *api.Event, b *botmaid.Bot) bool {
	if b.BeAt(e) {
		send(api.Event{
			Message: &api.Message{
				Text: random.String(wordBeAt),
			},
			Place: e.Place,
		}, b, false)
		return true
	}
	return false
}

func masterAt(e *api.Event, b *botmaid.Bot) bool {
	if b.BeAt(e) && b.IsMaster(*e.Sender) {
		send(api.Event{
			Message: &api.Message{
				Text: random.String(wordMasterAt),
			},
			Place: e.Place,
		}, b, false)
		return true
	}
	return false
}
