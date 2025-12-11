package service

import (
	"fmt"
	"os"
	"time"

	"zeusro.com/hermes/function/web/translate"
	baseModel "zeusro.com/hermes/model"
)

const P string = `其实，我对你是有一些失望的。当初给你定级高级软件工程师，是高于你面试时的水平的。我是希望进来后，你能够拼一把，快速成长起来的。高级软件工程师这个层级，不是把事情做好就可以的。你需要有体系化思考的能力。你做的事情，他的价值点在哪里？你是否作出了壁垒，形成了核心竞争力？你做的事情，和公司内其他团队的差异化在哪里？你的事情，是否沉淀了一套可复用的物理资料和方法论？为什么是你来做，其他人不能做吗？你需要有自己的判断力，而不是我说什么你就做什么。后续，把你的思考沉淀到日报周报月报里，我希望看到你的思考，而不仅仅是进度。另外，提醒一下，你的产出，和同层级比，是有些单薄的，马上要到年底了，加把劲儿。你看咱们团队的那个谁, 人家去年晋升之前，可以一整年都在项目室打地铺的。成长，一定是伴随着痛苦的，当你最痛苦的时候其实才是你成长最快的时候。加油

公司的文化是团队合作，拥抱变化，我觉得你是时候做出一点改变了。公司没有超人，包括我也一样，不要总想着自己是与众不同的，你需要找好自己的位置。只有认清了自己，你才能有相应的举措改善，否则连你自己都觉得没有依着感。来年3月份还会有一次答辩，多思考思考我说的话，你就会获得成长。没有提升，对你、对公司都没有什么好处，这些差距你需要搞得明明白白。

今年就由你勉强承接一下，我希望你能总结经验教训，让自己的感情软着陆，沉下心来做出一些改变，明年的A就是你的了。你应该知道，在我的手下培养出来的，没有弱兵。给你C，我自己也是感觉脸上无光。我希望你接下来的工作能够证明我的眼光，给我一个交代，也给自己一个交代。`

func PUA(start time.Time, title string) string {
	translator := translate.NewDeepSeekTranslator(os.Getenv("DEEPSEEK_API_KEY"))
	pua := fmt.Sprintf("%s \n仿照这个语句，把“高级软件工程师”换成“%s”，结合岗位的特性和原句的句式，写一段新的话", P, title)
	_, output, err := translator.Do(pua)
	if err != nil {
		response := baseModel.NewErrorAPIResponse(time.Since(start), err.Error())
		return response.Message
	}
	return output
}
