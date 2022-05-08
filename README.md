# 选课系统说明文档

接口文档：https://documenter.getpostman.com/view/18429077/Uyr8mxmM

项目地址：https://github.com/L2ncE/Course-Selection-System

## 前言💞

花了不少时间完成了这个简单的选课系统，不过也把我掌握的亮点技术基本都使用了一遍，例如OAuth2，Docker，JWT这些。不过比较遗憾的是因为时间问题没有使用到Redis进行课程推荐。

## 实现功能列表🤣

*基础功能要求*

1. 账号注册登录
2. 教务处新增编辑修改课程信息（时间、总人数、学分、课程名字）

3. 学生查看自己所选的课程
4. 老师查看选择自己课程的学生详情

5. 学生选课退课

*加分项*

1. 完善的**接口文档**

2. 项目**说明文档**

3. 规范的**commit记录**

4. 代码格式规范、层次清晰(自我感觉良好)

5. 对学生、老师、管理员进行**权限管理**

6. **JWT**登录

7. 修改用户密码

8. 找回密码，使用密保

9. 入参检验

10. 使用**正则表达式**提取教务在线课程（未使用第三方库）

11. 选课退课时使用**事务**实现数据访问

12. GitHub **OAuth2** 第三方登录 

13. 学生选课对课程时间进行冲突检测

14. 选课学分限制

15. 老师给课程打分

16. 搜索课程（模糊搜索）

17. 部署在云服务器上(使用**DOCKER**)

## 一些功能的具体实现说明😎

这个部分将一些功能的实现进行说明

### OAuth2

由于没有前端，所以这个GitHub第三方并不方便进行测试

为了便于展示，我写了一个十分简单的HTML文件

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>
<a href="https://github.com/login/oauth/authorize?client_id=4f0872202e5c46597131&redirect_uri=http://42.192.155.29:8081/oauth">Github 第三方授权登录</a>
</body>
</html>
```

上面的重定向URL即我部署在服务器的这个接口的URL



点击进去的效果

![QQ图片20220421162622.png](https://s2.loli.net/2022/04/21/quwn3OPVmovC6HU.png)

进行认证之后所有信息后端都可以拿到，但是想要进行下一步操作需要有前端同学一起进行，详细可以看接口文档，这里就不多赘述了

### Docker

这次使用了docker进行项目部署

```dockerfile
#从最新的GOLANG镜像开始
FROM golang:latest

MAINTAINER YXH

#地址
RUN mkdir -p /app/CSA
WORKDIR /app/CSA
COPY . /app/CSA

#改个代理
ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.cn"
RUN go mod download #下载所有依赖

RUN go build main.go

#使用这个端口
EXPOSE 8081

#运行的命令
ENTRYPOINT  ["./main"]
```

在服务器终端进行镜像生成

```shell
docker image build -t csa -f /app/CSA/csa.dockerfile .
```

可以在Goland上使用Docker插件，绑定服务器之间进行容器创建

![image-20220421163728246](https://s2.loli.net/2022/04/21/9r8iQxIXtCswyGj.png)

设置对外端口后即可进行使用，并且可以看到日志

![image-20220421163850739](https://s2.loli.net/2022/04/21/Csm3dA9SHQTPtca.png)

## 数据库😉

### 管理员表

```mysql
CREATE TABLE `admin`
(
    `Id`       bigint(20)  NOT NULL AUTO_INCREMENT,
    `Name`     varchar(10) NOT NULL,
    `Password` varchar(20) NOT NULL,
    `Identity` tinyint(4) DEFAULT '0',
    PRIMARY KEY (`Id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4
```

### 课程表

```mysql
CREATE TABLE `course`
(
    `Id`     bigint(20)  NOT NULL AUTO_INCREMENT,
    `Name`   varchar(20) NOT NULL,
    `Credit` float       NOT NULL,
    PRIMARY KEY (`Id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4
```

### 专业表

```mysql
CREATE TABLE `major`
(
    `MajorNum`  bigint(20)  NOT NULL,
    `MajorName` varchar(20) NOT NULL,
    PRIMARY KEY (`MajorNum`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
```

### 学生课程表

```mysql
CREATE TABLE `stu_course`
(
    `StudentNum` bigint(20) NOT NULL,
    `TCourseNum` bigint(20) NOT NULL,
    `Grade`      varchar(10) DEFAULT '暂未上传',
    `Id`         bigint(20) NOT NULL AUTO_INCREMENT,
    `Time`       varchar(20) DEFAULT NULL,
    PRIMARY KEY (`Id`),
    KEY `StudentNum` (`StudentNum`),
    KEY `stu_course_ibfk_2` (`TCourseNum`),
    CONSTRAINT `stu_course_ibfk_1` FOREIGN KEY (`StudentNum`) REFERENCES `student` (`Id`),
    CONSTRAINT `stu_course_ibfk_2` FOREIGN KEY (`TCourseNum`) REFERENCES `t_course` (`Id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4
```

### 学生表

```mysql
CREATE TABLE `student`
(
    `Id`       bigint(20)  NOT NULL AUTO_INCREMENT,
    `Name`     varchar(10) NOT NULL,
    `Password` varchar(20) DEFAULT NULL,
    `Question` varchar(20) DEFAULT '未设置密保',
    `Answer`   varchar(20) DEFAULT NULL,
    `MajorNum` bigint(20)  DEFAULT NULL,
    `Identity` tinyint(4)  DEFAULT '1',
    `NickName` varchar(10) DEFAULT NULL,
    `Credit`   float       DEFAULT '0',
    PRIMARY KEY (`Id`),
    KEY `MajorNum` (`MajorNum`),
    CONSTRAINT `student_ibfk_1` FOREIGN KEY (`MajorNum`) REFERENCES `major` (`MajorNum`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4
```

### 老师课程表

```mysql
CREATE TABLE `t_course`
(
    `Id`         bigint(20) NOT NULL AUTO_INCREMENT,
    `CourseNum`  bigint(20) NOT NULL,
    `TeacherNum` bigint(20) NOT NULL,
    `Time`       varchar(255) DEFAULT NULL,
    `Num`        int(11)      DEFAULT '0',
    `Total`      int(11)      DEFAULT NULL,
    PRIMARY KEY (`Id`),
    KEY `CourseNum` (`CourseNum`),
    KEY `TeacherNum` (`TeacherNum`),
    CONSTRAINT `t_course_ibfk_1` FOREIGN KEY (`CourseNum`) REFERENCES `course` (`Id`),
    CONSTRAINT `t_course_ibfk_2` FOREIGN KEY (`TeacherNum`) REFERENCES `teacher` (`Id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4
```

### 老师表

```mysql
CREATE TABLE `teacher`
(
    `Id`       bigint(20)  NOT NULL AUTO_INCREMENT,
    `Name`     varchar(10) NOT NULL,
    `Password` varchar(20) NOT NULL,
    `Question` varchar(20) DEFAULT '未设置密保',
    `Answer`   varchar(20) DEFAULT NULL,
    `Identity` tinyint(4)  DEFAULT '2',
    `NickName` varchar(10) DEFAULT NULL,
    PRIMARY KEY (`Id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4
```

## 更新日志😁

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

###### 进度 2022/4/13

对课程进行重构并测试，管理员权限检验进行升级，已完美实现。课程表中的增删改查已完美实现

###### 进度 2022/4/15

老师课程相关已完美实现，并测试

###### 进度 2022/4/16

学生课程相关以实现，基础功能实现，多对多暂未实现

###### 进度 2022/4/17

实现多对多，实现老师自己课程的学生详情，使用事务进行选课退课，均已实现。接口文档已完善，删除需要匹配ID

###### 进度 2022/4/18

实现GitHub第三方登录，暂未测试

###### 进度 2022/4/19

测试GitHub第三方登录，实现选课冲突测试，实现选课学分限制

###### 进度 2022/4/20

实现老师打分，实现搜索课程功能，管理员表使用docker打包部署

###### 进度 2022/4/21

已测试完毕，修复所有已知BUG

###### 进度 2022/5/1

新增日志功能

##### 进度 2022/5/8

新增viper配置