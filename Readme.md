# ==🧢🧢🧢项目介绍==

项目名为云餐厅，基于gin和gorm框架编写的一个网上点餐项目，本人主要参与其中后端接口、基本逻辑业务和集成第三方技术的编写。

 

<hr>

# 🚀🚀🚀==项目功能==

1、首先在浏览器中访问网址：http://localhost:8080/msite

访问前端页面，初始页面如下，附近商家是根据数据库保存的经纬度进行查询，如果没有传入经纬度，则会有一个默认的经纬度进行查询附近商家：

<img src="https://gitee.com/zsj20020818/cloud-restaurant/raw/master/images/image-20230620143020818.png" alt="img" style="zoom:67%;" />

2、用户可以点击右上角的登录/注册，使用短信和验证码两种方式登录，如果是新用户，则在登录的过程中会将新用户的数据保存到云服务器的数据库中。

（1）短信登录效果如下：

<img src="https://gitee.com/zsj20020818/cloud-restaurant/raw/master/images/img.png" alt="img" style="zoom:67%;" />

收到短信验证码如下：

<img src="https://gitee.com/zsj20020818/cloud-restaurant/raw/master/images/004F299829E4758AA5B3CCD205133407.png" alt="img" style="zoom: 33%;" />

输入正确的短信验证码即可完成登录：

<img src="https://gitee.com/zsj20020818/cloud-restaurant/raw/master/images/img_1.png" alt="img" style="zoom:67%;" />

（2）使用用户名密码+图片验证码登录也是同理，验证码功能结合redis以及base64Captcha库实现：

<img src="https://gitee.com/zsj20020818/cloud-restaurant/raw/master/images/img_2.png" alt="img" style="zoom:67%;" />

登录成功后：

<img src="https://gitee.com/zsj20020818/cloud-restaurant/raw/master/images/img_7.png" alt="img" style="zoom:67%;" />

3、成功登录之后，点击搜索并输入搜索的内容，即可搜索到相关店铺

<img src="https://gitee.com/zsj20020818/cloud-restaurant/raw/master/images/img_3.png" alt="img" style="zoom:67%;" />

4、点进店铺之后，选择想要的食品，并点击“去结算”，即可跳转到支付页面

<img src="https://gitee.com/zsj20020818/cloud-restaurant/raw/master/images/img_8.png" alt="img" style="zoom:67%;" />

5、支付的功能集成了支付宝开放平台的沙箱环境，使用沙箱环境账号成功登陆后即可实现支付功能

![![img](https://gitee.com/zsj20020818/cloud-restaurant/raw/master/images/img_4.png)](D:\Study\GoLearn\GoWork\src\CloudRestaurant\images\img_4.png)

![![img](https://gitee.com/zsj20020818/cloud-restaurant/raw/master/images/img_5.png)](D:\Study\GoLearn\GoWork\src\CloudRestaurant\images\img_5.png)



6、登录过后，用户的信息将存到基于redis的session当中，在用户信息当中，用户可以更新自己的头像信息，用户传输的头像将保存到FastDFS分布式文件系统当中

<img src="https://gitee.com/zsj20020818/cloud-restaurant/raw/master/images/img_6.png" alt="img" style="zoom:67%;" />

# ==💡💡💡项目总结==

1、该项目使用了**gin**和**gorm**框架，在框架的基础上实现web应用的开发；

2、项目使用了**MySQL**数据库来保存各种数据；

3、项目结合**Redis**和**Base64Captcha**来实现用户登录时图片验证码的功能；

4、项目还调用了阿里云的**短信服务**来实现让用户能够接收短信验证码的功能；

5、项目对接了支付宝开放平台的**沙箱环境**，实现了用户支付金额的功能；

6、项目使用**FastDFS**分布式文件系统，来实现用户头像文件的上传功能













