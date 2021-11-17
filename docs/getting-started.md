# 概览
本文是BridgX的快速入门指南，典型的用户包括三个步骤：<br>
1. 添加云厂商账户；<br>
2. 在集群管理模块创建需要的集群模板；<br>
3. 在扩缩容任务里建立扩缩容任务；<br>

# 前置条件
1、已经有云厂商账户，并且获得了AccessKey和AccessKey Secret信息，如果没有，请提前申请；<br>
[阿里云申请入口链接](https://help.aliyun.com/document_detail/53045.html)<br>

2、已经获取了网址信息，以及用户名和密码信息；<br>
如果是初次部署，可在浏览器中打开 http://localhost:8888 登录系统,采用的认证方式是默认配置可用以下帐号登录<br>
默认用户名:root<br>
默认密码为:admin<br>

# 第一步：添加云厂商账户

1. 进入云厂商账户，点击添加云账户
![image](https://user-images.githubusercontent.com/94337797/142158688-a3a17da1-a068-4396-81fb-cf1f1270f184.png)

2. 填写云厂商的ak和sk信息，并保存；
![image](https://user-images.githubusercontent.com/94337797/142158808-19166f17-9ed6-4f5e-9ffe-65f698bbe7ed.png)


# 第二步：创建机型模板

1. 集群管理->创建集群
![image](https://user-images.githubusercontent.com/94337797/142158959-889069f7-1620-4b27-9764-2c6224b1ce72.png)

2. 进入云厂商配置页面，进行配置：
![image](https://user-images.githubusercontent.com/94337797/142159081-c6024be5-94bb-405d-8596-4e7a95aa8f26.png)


3. 点击下一步进入网络配置页面：
![image](https://user-images.githubusercontent.com/94337797/142159133-f6b14355-2c72-4061-9b75-74dd3742d210.png)


4. 网络配置完成后，点击下一步，进行机器规格配置
![image](https://user-images.githubusercontent.com/94337797/142159200-605ca273-80cf-463f-b84e-b7e8bf649409.png)


5. 机器规格配置完成后，点击下一步，进行存储磁盘配置，配置完成后，点击提交，创建集群成功
![image](https://user-images.githubusercontent.com/94337797/142159248-30ddad6e-cc32-4da8-8d75-d275fdef5684.png)


6. 点击提交，则页面跳转到集群列表，有最新的创建的集群信息，则表示创建成功
![image](https://user-images.githubusercontent.com/94337797/142159291-2424cfc2-4f01-4367-924b-f03b50a868d8.png)


# 第三步：创建快速扩缩容任务

1. 扩缩容任务->创建任务；
![image](https://user-images.githubusercontent.com/94337797/142159354-c10a839c-ed0f-41bd-989c-c5a6de31975f.png)


2. 填写扩缩容任务配置，填写完后点击提交按钮；
![image](https://user-images.githubusercontent.com/94337797/142159394-5a4a738c-c44e-4a06-8bf0-96a590f4cfbe.png)


3. 提交成功后，页面跳转到任务列表，显示刚才创建的任务的执行情况
![image](https://user-images.githubusercontent.com/94337797/142159426-f6f501ac-c32f-44c0-b331-ecd1929f7cd2.png)


4. 如果需要查看本次任务创建的机器的IP，可点击执行明细查看详细信息；
![image](https://user-images.githubusercontent.com/94337797/142159460-162a5eb0-6b6d-45ae-8be6-d1464405eb3a.png)


