package service

import (
	"github.com/kataras/iris/v12"
	"github.com/letheliu/hhjc-devops/common/constants"
	"github.com/letheliu/hhjc-devops/common/encrypt"
	"github.com/letheliu/hhjc-devops/entity/dto/result"
	"github.com/letheliu/hhjc-devops/entity/dto/user"
	"github.com/letheliu/hhjc-devops/entity/vo"
	dao2 "github.com/letheliu/hhjc-devops/user/dao"
)

type UserService struct {
	userDao dao2.UserDao
}

/*
*

	用户登录处理
*/
func (userService *UserService) Login(ctx iris.Context) (result.ResultDto, *user.UserDto) {
	var (
		err     error
		userVo  = new(vo.LoginUserVo)
		userDto *user.UserDto
	)
	if err = ctx.ReadJSON(&userVo); err != nil {
		return result.Error("解析入参失败"), nil
	}

	userVo.Passwd = encrypt.Md5(userVo.Passwd)

	userDto, err = userService.userDao.GetUser(*userVo)
	//fmt.Print("userDto", userDto)
	if err != nil || len(userDto.UserId) < 1 {
		return result.Error("用户名密码错误"), nil
	}

	return result.SuccessData(userDto), userDto
}

/*
*

	用户登录处理
*/
func (userService *UserService) ChangePwd(ctx iris.Context) result.ResultDto {
	var userInfo *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	var (
		err       error
		newUserVo = new(vo.ChangeUserPwdVo)
	)
	if err = ctx.ReadJSON(&newUserVo); err != nil {
		return result.Error("解析入参失败")
	}
	newUserVo.UserId = userInfo.UserId
	userVo1 := vo.LoginUserVo{UserId: userInfo.UserId, Passwd: encrypt.Md5(newUserVo.OldPwd)}
	userVo2, err := userService.userDao.GetUser(userVo1)

	if err != nil || len(userVo2.UserId) < 1 {
		return result.Error("原始密码错误")
	}

	var userDto = user.UserDto{UserId: userInfo.UserId, Passwd: encrypt.Md5(newUserVo.NewPwd)}

	err = userService.userDao.UpdateUser(userDto)
	if err != nil {
		return result.Error("修改密码失败")
	}

	return result.Success()
}

/*
*

	用户登录处理
*/
func (userService *UserService) GetUserInfo(ctx iris.Context) result.ResultDto {
	var userInfo *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	var (
		err     error
		userDto *user.UserDto
	)
	userVo := vo.LoginUserVo{UserId: userInfo.UserId}
	userDto, err = userService.userDao.GetUser(userVo)

	if err != nil {
		return result.Error("用户不存在")
	}

	return result.SuccessData(userDto)
}
