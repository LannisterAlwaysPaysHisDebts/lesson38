# 用户注册/登录 micro service demo

## 运行流程
make 

## 搭建流程
1. 按照rpc接口要求编写`user.proto`,protoc生成代码 
   ```
   protoc -I ./ --go_out=./ --micro_out=./ ./proto/user/*.proto;
   ```

2. 设计数据表,编写基础服务domain，编写对应的单元测试
    基本结构为：
    - model: user表结构
    - repository: user data操作
    - service: 具体业务逻辑
    
    mysql docker启动并测试:
    ```
    docker run -p 3306:3306 -v $PWD/data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456  -d  --name micro-mysql cap1573/mysql:5.6
    ```

3. 编写对外暴露接口handler

4. 编写main函数：
    1. 初始化micro服务;
    2. 建立数据库连接;
    3. 初始化userService与handler，handler与rpc接口绑定;

5. 编写test,通过全局test;

6. 编译代码;

7. 编写Dockerfile, 测试docker搭建

8. 编写makefile;