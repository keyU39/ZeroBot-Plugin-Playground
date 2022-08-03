package daily_news

import (
	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	"github.com/FloatTech/zbputils/web"
	"github.com/tidwall/gjson"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

func init() { // 插件主体
	engine := control.Register("daily_news", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Help: "每日早报\n" +
			"- [1] 启用后会在每天早上发送一份早报",
		PrivateDataFolder: "daily_news",
	})

	// 开启
	engine.OnKeyword(`今日早报`, zero.OnlyGroup).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			data, err := web.RequestDataWith(web.NewDefaultClient(), "http://dwz.2xb.cn/zaob", "GET", "", "")
			if err != nil {
				return
			}
			picURL := gjson.Get(string(data), "imageUrl").String()
			if err != nil {
				ctx.SendChain(message.Text("ERROR:", err))
				return
			}
			ctx.SendChain(message.Image(picURL))
		})
}
