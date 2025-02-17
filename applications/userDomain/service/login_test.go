package service

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/TremblingV5/DouTok/applications/userDomain/misc"
)

func TestCheckPassword(t *testing.T) {
	Init()

	curr := fmt.Sprint(time.Now().Unix())

	userId, err, errNo := NewWriteNewUserService(context.Background()).WriteNewUser(curr, curr)
	if err != nil {
		log.Panicln(err)
	} else {
		log.Println(userId, errNo)
	}

	notExistedUserId, err, errNo := NewCheckPasswordService(context.Background()).CheckPassword("*****", "789456")
	if notExistedUserId == 0 && errNo == &misc.UserNameErr {
		log.Println("查询不存在的用户名返回正常")
	} else {
		log.Panicln("查询不存在的用户未报错")
	}

	passwordWrongUserId, err, errNo := NewCheckPasswordService(context.Background()).CheckPassword(curr, "-----")
	if passwordWrongUserId == 0 && errNo == &misc.PasswordErr {
		log.Println("密码错误情况返回正常")
	} else {
		log.Panicln("密码错误的用户未报错")
	}

	userIdSearched, err, errNo := NewCheckPasswordService(context.Background()).CheckPassword(curr, curr)
	if err == nil && errNo == &misc.Success && userIdSearched == userId {
		log.Println(userId, errNo, "登陆功能正常")
	} else {
		log.Panicln("登陆功能异常")
	}
}
