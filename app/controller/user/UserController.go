package user

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"github.com/letheliu/hhjc-devops/common/cache/factory"
	"github.com/letheliu/hhjc-devops/common/constants"
	"github.com/letheliu/hhjc-devops/common/jwt"
	"github.com/letheliu/hhjc-devops/common/seq"
	"github.com/letheliu/hhjc-devops/entity/dto/result"
	"github.com/letheliu/hhjc-devops/user/service"
)

type UserController struct {
	userService service.UserService
}

func UserControllerRouter(party iris.Party) {
	var (
		adinUser = party.Party("/user")
		aus      = UserController{userService: service.UserService{}}
	)
	//登录
	adinUser.Post("/login", hero.Handler(aus.login))

	//退出登录
	adinUser.Post("/logout", hero.Handler(aus.logout))

	//修改密码
	adinUser.Post("/changePwd", hero.Handler(aus.changePwd))

	//修改密码
	adinUser.Get("/getUserInfo", hero.Handler(aus.getUserInfo))

}

/*
*
登录处理类
*/
func (aus *UserController) login(ctx iris.Context) {
	resultDto, userDto := aus.userService.Login(ctx)

	if userDto != nil {
		userDto.TokenId = seq.Generator()
		token, _ := jwt.G_JWT.GenerateToken(userDto)
		//token 保存至redis
		factory.SetToken(constants.REDIS_ADMIN_FORMAT, userDto.TokenId, token)
		ctx.SetCookieKV(jwt.DEFAULT_TOKEN, token)
	}

	ctx.JSON(resultDto)
}

/*
*
退出登录处理类
*/
func (aus *UserController) logout(ctx iris.Context) {

	ctx.RemoveCookie(jwt.DEFAULT_TOKEN)

	ctx.JSON(result.Success())
}

/*
*
登录处理类
*/
func (aus *UserController) changePwd(ctx iris.Context) {
	resultDto := aus.userService.ChangePwd(ctx)
	ctx.JSON(resultDto)
}

/*
*
登录处理类
*/
func (aus *UserController) getUserInfo(ctx iris.Context) {
	resultDto := aus.userService.GetUserInfo(ctx)
	ctx.JSON(resultDto)
}
