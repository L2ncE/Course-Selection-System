# 选课系统

### 基础功能要求:

- [x] 账号注册登录
- [x] 教务处新增编辑修改课程信息（时间、总人数、学分、课程名字）
- [ ] 学生查看自己所选的课程
- [ ] 老师查看选择自己课程的学生详情
- [ ] 学生选课退课

- [x] 代码格式规范、层次清晰(自我感觉良好)
- [x] 项目说明文档
- [x] 修改密码
- [x] 对学生、老师、学校等等进行**权限管理**
- [x] JWT登录
- [x] 找回密码，使用密保
- [ ] OAuth2
- [ ] RPC和MQ
- [ ] 高并发
- [ ] 微服务
- [ ] 学生选课对课程时间进行冲突检测
- [ ] 选课学分限制
- [ ] 部署在云服务器上(使用DOCKER)
- [x] 入参检验
- [ ] 搜索课程
- [ ] 推荐课程

### DEADLINE

2022.4.23 949764437@qq.com

###### 进度 2022/4/2

项目的基础框架以及学生和老师的dao层、model层、service层

老师api层的登陆注册已完成

###### 进度 2022/4/6

老师学生的用户基本功能已完成（暂未测试）

学科相关已布局，可以通过接口获得学科ID NAME对照表(已测试)

###### 进度 2022/4/7

为了防止学生老师中有同名的出现，现在以统一验证码（ID）为标准，登录等相关操作也使用统一验证码来完成

已完成老师部分并测试

PS.发现这样并不是最优解，目前方案为姓名可以重复，但是昵称不能重复，以昵称为统一标准

###### 进度 2022/4/7

按照昨天的想法进行更新，目前昵称以及ID都是唯一标准，以ID作为主键

老师相关已完全测试通过

明天将学生以及管理员部分重构后进行下一步

###### 进度 2022/4/9

学生以及管理员部分重构完成

###### 进度 2022/4/11

进行选课，课程相关操作，已完成部分接口

###### 进度 2022/4/12

将课程的相关服务封装成了接口，还未测试，数据库中课程表实现还未进行