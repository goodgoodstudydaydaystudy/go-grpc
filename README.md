# go-grpc


### 笔记：

1、将用户信息以结构体的形式封装传递，便于对用户信息的增删

2、数据库返回的error内容也应该经过拦截器，处理匹配错误信息。

3、期初认为通过sql的error判断具体错误（如密码错或账号错），实际上单独查询密码或者账号，单独核对，就可以简单完成这个功能了。

4、“映射”不仅仅是map的形式。（还不是很理解）

5、“结构体”理解为一个类型，在“结构体”中可以包含各种属性，若某个函数或功能接口属于某个类型，则在实例化的结构体中，还能调取该功能或函数接口。

6、为了避免外部“随意访问”内部功能或属性，通常内部类型会拒绝外部访问。当pkg之间需要相互调用，解决方法是：通过指定对象函数或功能接口返回内部类型（返回实例），供外部调用。

7、‘dao’，data asscess object，数据访问层。一般用于操作db。

8、‘dao’采用interface，包含db内的操作函数，外部访问db就会经过dao的interface再调用到db内的功能。

9、**具体外部如何调用db内的函数：db的操作函数都属于同一个结构体类型，结构体还拥有slqx.db的指针；同时不允许外部随意访问（包括db操作函数）。dao内构建了一个xxxDao的接口，以及一个NewDao()返回xxxDao接口。当外部想调用db函数，首先需要调用NewDao()，获得接口实例，通过实例访问接口内的db函数。**

10、**dao的xxxDao接口完全使用db的操作函数，与此同时db的操作函数属于同一个结构体，所以xxxDao也同等于db内的结构体类型。NewDao()返回xxxDao接口，相当于返回了结构体，估外部也能访问。—— duck type **

---

- [x] 规范命名
- [x] 简单了解client和server的拦截器
- [x] Account Server
  - [x] 在server启动的时候建立数据库连接, 建立失败则启动失败
  - [x] 在server关闭的时候关闭数据库连接
  - [x] 尝试返回自定义错误码
  - [x] server结构体私有化, 提供New函数
- [x] Account Client
  - [x] client结构体私有化
  - [x] 测试自定义错误码能不能用
- [x] 仿照account修改pay的server和client, 先不要做pay的db

- [x] 注册增加“性别”，登录resp返回增加“性别”
  - [x] 简化新增属性的修改工作
  - [x] 封装userInfo的信息，在pb增加userInfo的结构体

- [x] client增加rpc查询接口

- [x] 增加wallet系统
- Wallet Server
  - [x] 增加充值功能
  - [x] db若表中无用户记录，直接插入数据. 使用insert or update的语法
  
- Wallet Client
  - [x] 增加客户端

- [x] account.proto 的Get接口不能直接返回info。

- [x] 在user-logic加入第一个接口: 登录
  - [x] 校验用户名, 密码 
  - [x] 返回UserInfo, 以及这个人的余额
- [x] 在user-logic加入注册接口
- [x] logic_test 测试登录返回余额时，wallet_db的get函数有异常
- [ ] wallet_db创建协程池管理修改order状态的协程
- [ ] 优化main的代码管理
