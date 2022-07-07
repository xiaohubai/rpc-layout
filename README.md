# 开发过程

``` go
kratos upgrade                                                               //安装依赖,升级版本
make init
kratos new github.com/xiaohubai/rpc-layout                                   //创建大仓
kratos new app/dict --nomod                                                  //大仓模式 公用go.mod
kratos proto add api/dict/v1/dict.proto                                      //添加proto文件
kratos proto client api/dict/v1/dict.proto                                   //生成client代码
kratos proto server api/dict/v1/dict.proto -t app/dict/internal/service      //生成server代码

```


# wire
```
biz.repo定义接口，usecase实例的代码实现
data：实现的数据库连接，实现biz.repo接口定义
service : 关联到http grpc的调用实例和usercase
serve:service实例注册到grpc http

```


#TODO:
1.消费consume kafka,入es  和warn告警
2.获取字典序 rpc http方式
3.上传excel 和下载excel
4.服务注册、发现，consul+ACL
5.远程配置文件



