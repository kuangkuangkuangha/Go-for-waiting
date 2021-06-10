### 几个概念

#### 静态网站和动态网站

静态网站：不支持数据交互的网站，服务端不会处理客户端的数据（静态页面，后缀为.html   htm）

动态网站：支持数据交互的网站，你提交的数据服务端是会处理的，比如登入注册输个用户名密码啥的



实现动态网站的技术

| 实现技术 | 网站后缀 |
| -------- | :------: |
| ASP      |   .asp   |
| PHP      |   .php   |
| .NET     |  .aspx   |
| Java     |   .jsp   |





**Inurl:*.php**





#### 端口和端口号

![Snipaste_2021-05-31_20-21-37](/Users/zhangkuang/Desktop/Snipaste_2021-05-31_20-21-37.png)

端口范围0--65535，1024以下的端口都留给系统。80端口一般留给web服务（网站的服务器，比如apache，iis）用，21端口一般留给ftp用，25端口留给邮件服务器用

从哪个端口出去不用管，要看从哪个端口进去 

一般360杀毒软件，迅雷占用80端口，要到远程下载软件包



如果apache占用80端口，你就得请求80端口

如果占用81端口你就得请求81端口

apache还可以在很多端口监听



如何查询端口号有么有被占用？

`sudo lsof -i:端口号`



#### BS和CS架构

B：brower  浏览器

S：server.   服务器

通过浏览器去访问服务器，比如百度，360啥的



C：client 客户端

S：server 服务端

通过一个客户端软件去访问服务器

比如：QQ  ， miniproject。。。



bs架构和cs架构的区别：bs架构要求有操作系统和浏览器就行，与操作平台无关，对客户端的计算机电脑配置要求较低；cs架构只运用与局域网中，要求有相同的操作系统，对计算机电脑配置要求较高



#### 前台和后台

前台是给浏览者看的

后台是给管理员操作的，后台用来操作前台的数据。



### 安装Apache

语言的运行需要环境，Apache为php运行提供了环境，（IIS）也可以为PHP运行提供环境，IIS是微软公司开发的，为asp和aspx提供环境。



mac下apache的地址

`/etc/apache2`

Mac自带Apache，可以直接启用

`sudo apachectl start`

在浏览器中输入localhost，local表示本地，host表示主机，localhost表示了ip地址127.0.0.1



### PHP安装

### mysql安装

mysql是cs架构的，就是如果你要访问服务器就必须通过客户端



### LAMP和WAMP

XAMPP（原来叫LAMPP）
(Apache~Web服务器端软件+MySQL+PHP+PERL~解释器)是一款功能强大的建站集成软件包

LAMP是一个很好的组合

L：Linux（开源）

A：apache（开源）

M：mysql（开源的数据库）

P：php（开源）



### 更改虚拟目录

因为php的运行需要**Apache**的支持，所以PHP的目录要告知apache它在哪里。再哪告知呢？在apache的配置文件中,http.conf是Apache的配置

httpd.conf配置php虚拟目录的位置,就是将网站的要显示的文件（html）的地方

mac下的是

`/Library/WebServer/Documents/`

=Macintosh/资源库/WebServer/Documents



站点：存放与网站有关的所有素材的文件夹的地方

虚拟目录：文件夹+权限

测试：

在浏览器地址栏中输入http://localhost/test.php 发现没有出现指定文件

原因是还没添加权限



### 更改首页

在apache的配置文件中搜索“DirectionaryIndex”，这个指令是设置网站首页的

在DirectionaryIndex中按顺序寻找,没有index.html就找test.php,以此类推

![Snipaste_2021-06-02_23-51-51](/Users/zhangkuang/Desktop/Snipaste_2021-06-02_23-51-51.png)



### 更改监听端口号

在配置文件中找到listen修改端口

![Snipaste_2021-06-03_00-01-50](/Users/zhangkuang/Desktop/Snipaste_2021-06-03_00-01-50.png)



### DNS

domain name server。域名解析系统，将域名解析成ip地址

当我们在浏览器地址栏中输去www.baidu.com是，浏览器会先去DNS中找到www.baidu.com对应的ip地址，然后通过这个ip地址去访问服务器

![Snipaste_2021-06-03_00-23-33](/Users/zhangkuang/Desktop/Snipaste_2021-06-03_00-23-33.png)

客户端输入网址，首先会请求**最近的DNS服务器**，将域名解析成ip地址，DNS服务器有很多，电信，移动的机房里

**最近的DNS解析机器**就是你的本机

hosts的地址在`/etc/hosts`



可以通过修改hosts文件来定义自己的DNS，例如在文件中将127.0.0.1后的localhost改为www.baidu.com

然后在浏览器地址栏中输入www.baidu.com访问的将是127.0.0.1

![Snipaste_2021-06-03_00-40-58](/Users/zhangkuang/Desktop/Snipaste_2021-06-03_00-40-58.png)





在互联网上，唯一标识一台计算机的是ip地址。但是ip地址不方便记忆，我们通过一个域名对应一个ip地址

关于网站域名是多ip或动态ip，比如Google，有些是一个IP多个网站域名

比如QQ，QQ并不是只有一个服务器，全国各地有好多QQ的服务器，QQ用户多，消息流量大，服务器承受的压力也大

其中有一个主机专门用来判断，看哪些服务器负担重，哪些服务器正闲着，然后将用户的请求分发给比较轻松的服务器来减轻负担



集群：小饭店原来只有一个厨师，切菜洗菜备料炒菜全干。后来客人多了，厨房一个厨师忙不过来，又请了个厨师，两个厨师都能炒一样的菜，这两个厨师的关系是集群。为了让厨师专心炒菜，把菜做到极致，又请了个配菜师负责切菜，备菜，备料，厨师和配菜师的关系是分布式，一个配菜师也忙不过来了，又请了个配菜师，两个配菜师关系是集群





![Snipaste_2021-06-03_00-12-57](/Users/zhangkuang/Desktop/Snipaste_2021-06-03_00-12-57.png)







### 虚拟主机的配置

现实中买一台服务器一万块，找人托管一万块，有些公司大量做网站来推销自己的产品，比如在百度搜索go语言培训，会返回很多词条，里面就有好多公司的网站，在前排的网站就会有更大的概率被点击，但做一个网站就买一个服务器显然不太现实，所以我们一般用一个主机

![IMG_20210604_160115](/Users/zhangkuang/Desktop/IMG_20210604_160115.jpg)



一个网站可以看作是一个文件夹，所以我们可以在一个主机上建多个文件夹

![Snipaste_2021-06-03_14-39-55](/Users/zhangkuang/Desktop/Snipaste_2021-06-03_14-39-55.png)



一个电脑上装一个Apache，一个Apache支持多个网站，从浏览者的角度来看，每个网站都是一个独立的主机，

但实际上都在一个Apache上



开启虚拟主机

在apache的配置文件httpd.conf中，找到virtual hosts(虚拟主机)， 去掉#，然后找到include后面的文件并修改，这个是虚拟主机的配置文件

![Snipaste_2021-06-03_14-48-25](/Users/zhangkuang/Desktop/Snipaste_2021-06-03_14-48-25.png)







![Snipaste_2021-06-05_13-22-20](/Users/zhangkuang/Desktop/Snipaste_2021-06-05_13-22-20.png)



apache的配置文件夹是conf，apache的配置文件是conf下的httpd.conf

hosts文件是电脑系统的，virtual host 是apache的

php的配置文件是php.aii





### Apache

Apache，指的是(Apache软件基金会)下的一个项目——Apache HTTP Server Project；

Nginx同样也是一款开源的***HTTP服务器软件***（当然它也可以作为邮件代理服务器、通用的TCP代理服务器）。

***HTTP服务器***本质上也是一种应用程序——它通常运行在服务器之上，绑定服务器的IP地址并监听某一个tcp端口来接收并处理HTTP请求，这样客户端（一般来说是IE, Firefox，Chrome这样的浏览器）就能够通过HTTP协议来获取服务器上的网页（HTML格式）、文档（PDF格式）、音频（MP4格式）、视频（MOV格式）等等资源。下图描述的就是这一过程：

![img](https://pic1.zhimg.com/50/904696074e077934e601f175913f42fd_hd.jpg?source=1940ef5c)

不仅仅是Apache HTTP Server和Nginx，绝大多数编程语言所包含的类库中也都实现了简单的HTTP服务器方便开发者使用：

- [HttpServer (Java HTTP Server )]
- [Python SimpleHTTPServer]

使用这些类库能够非常容易的运行一个HTTP服务器，它们都能够通过绑定IP地址并监听tcp端口来提供HTTP服务。

Apache Tomcat 则是(Apache基金会下)的另外一个项目，与Apache HTTP Server相比，Tomcat能够**动态**的生成资源并返回到客户端。Apache HTTP Server和Nginx都能够将某一个文本文件的内容通过HTTP协议返回到客户端，但是这个文本文件的内容是固定的——也就是说无论何时、任何人访问它得到的内容都是完全相同的，这样的资源我们称之为**静态**资源。动态资源则与之相反，在不同的时间、不同的客户端访问得到的内容是不同的，例如：

- 包含显示当前时间的页面
- 显示当前IP地址的页面