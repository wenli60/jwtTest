/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	_ "git.code.oa.com/trpc-go/trpc-config-tconf"
	_ "git.code.oa.com/trpc-go/trpc-database/gorm"
	_ "git.code.oa.com/trpc-go/trpc-filter/debuglog"
	_ "git.code.oa.com/trpc-go/trpc-filter/recovery"
	_ "git.code.oa.com/trpc-go/trpc-filter/validation"
	"git.code.oa.com/trpc-go/trpc-go"
	_ "git.code.oa.com/trpc-go/trpc-log-atta"
	_ "git.code.oa.com/trpc-go/trpc-metrics-m007"
	_ "git.code.oa.com/trpc-go/trpc-naming-polaris"
	_ "git.code.oa.com/trpc-go/trpc-opentracing-tjg"
	pb "git.code.oa.com/trpcprotocol/tyoffice/tyoffice_notification"
	"git.woa.com/trpcprotocol/test/helloworld/internal/api"
	"git.woa.com/trpcprotocol/test/helloworld/internal/models"
)

func main() {
	srv := trpc.NewServer()
	err := models.Setup()
	if err != nil {
		panic(err)
	}
	pb.RegisterNotificationService(srv, api.New()) // 注册接口服务

	if err := srv.Serve(); err != nil {
		panic(err)
	}
}
