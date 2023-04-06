## Nginx 简介

### 一、Nginx概述

#### 1.1 概述

Nginx（“engine x”）是一个高性能的 HTTP /反向代理的服务器及电子邮件（IMAP/POP3)代理服务器。官方测试nginx能够支撑5万并发，并且cpu，内存等资源消耗却非常低，运行非常稳定。最重要的是开源，免费，可商用的。

Nginx还支持热部署，几乎可以做到7 * 24 小时不间断运行，即时运行数个月也不需要重启，还能够在不间断服务的情况下对软件进行升级维护。

#### 1.2 Nginx应用场景

- 虚拟主机：一台服务器虚拟出多个网站。

- 静态资源服务：提供http资源访问服务。

- 反向代理，负载均衡。当网站的访问量达到一定程度后，单台服务器不能满足用户的请求时，需要yo哪个多台服务器集群可以使用nginx做反向代理。并且多台服务器可以平均分担负载，不会因为某台服务器负载高宕机而某台服务器闲置的情况。

#### 1.3 正向代理

**正向代理**：一般的访问流程是客户端直接向目标服务器发送请求并获取内容，使用正向代理后，客户端通过配置或其他方式改为向代理服务器发送请求，并指定目标服务器（原始服务器），然后由代理服务器和原始服务器通信，转交请求并获得的内容，再返回给客户端。正向代理隐藏了真实的客户端，为客户端收发请求，使真实客户端对服务器不可见；

![1](images/1.png)

#### 1.4 反向代理

**反向代理：**正好相反。对于客户端来说，反向代理就好像目标服务器。并且客户端不需要进行任何设置。客户端向反向代理发送请求，接着反向代理判断请求走向何处，并将请求转交给客户端，使得这些内容就好像它自己的一样，一次客户端并会并会不感知到反向代理后面的服务，因此不需要客户端做任何设置，只需要把反向代理服务器当成真正的服务器就好了。

![2](images/2.png)

#### 1.5 负载均衡

负载均衡建立在现有网络结构之上，它提供一种链家有效透明的方法扩展网络设备和服务器的宽带、增加吞吐量，加强网络数据处理能力，提高网络的灵活性和可用性。

![3](images/3.png)

#### 1.6 动静分离

为了加快网站的解析速度，可以把动态页面和静态页面由不同的服务器来解析，加快解析速度，降低原来单个服务器的压力。一般来说，都需要将动态资源和静态资源分开，由于Nginx的高并发和静态资源缓存等特性，经常将静态资源部署在Nginx上。如果请求的是静态资源，直接到静态资源目录获取资源，如果是童泰资源的请求，则利用反向代理的原理，把请求转发给对应后台应用去处理，从而实现动静分离。

### 二、Nginx安装

#### 2.1 [进入官网下载](http://nginx.org)

#### 2.2 安装相关依赖

##### 2.2.1 第一步

1、 安装pcre

```bash
wget http://downloads.sourceforge.net/project/pcre/pcre/8.37/pcre-8.37.tar.gz
```

2、解压文件 tar -zxvf 路径

3、pcre主目录执行命令 :

```bash
./configure
```

​     可能遇到的情况：没有c++支持

![4](images/4.png)

安装c++支持：

```bash
yum install -y gcc gcc-c++
```

4、完成后、回到pcre目录下执行 

```bash
make && make install
```

5、查看版本 :

```bash
pcre-config --version
```

##### 2.2.2 第二步，安装其他依赖

**zlib openssl**

```bash
yum -y install make zlib zlib-devel gcc-c++ libtool openssl openssl-devel
```

#### 2.3 安装nginx

1. 解压nginx-xx.tar.gz包

   ```bash
   tar -zxvf nginx-xxx.tar.gz
   ```

2. 进入解压目录，执行./configure

   ```bash
   ./configure
   ```

3. make&&make install

   ```bash
   make && make install
   ```

#### 2.3 开放端口

```bash
#查看开放的端口号
firewall-cmd --list-all
#设置开放的端口号
firewall-cmd --add-service=http –permanent
sudo firewall-cmd --add-port=80/tcp --permanent
#重启防火墙
firewall-cmd –reload
```

#### 2.4 访问 

![5](images/5.png)

### 三、nginx常用命令和配置文件

#### 3.1 常用命令

```bash
#查看版本
在/usr/local/nginx/sbin 目录下执行 ./nginx -v

#启动nginx
在/usr/local/nginx/sbin 目录下执行 ./nginx

#关闭nginx
在/usr/local/nginx/sbin 目录下执行 ./nginx -s stop

#重加载nginx
在/usr/local/nginx/sbin 目录下执行 ./nginx -s reload
```

#### 3.2 配置文件详细讲解

```bash
#配置文件位置
位置：/usr/local/nginx/conf/nginx.conf
```

```bash
######Nginx配置文件nginx.conf中文详解#####

#定义Nginx运行的用户和用户组
user www www;

#nginx进程数，建议设置为等于CPU总核心数。
worker_processes 8;
 
#全局错误日志定义类型，[ debug | info | notice | warn | error | crit ]
error_log /usr/local/nginx/logs/error.log info;

#进程pid文件
pid /usr/local/nginx/logs/nginx.pid;

#指定进程可以打开的最大描述符：数目
#工作模式与连接数上限
#这个指令是指当一个nginx进程打开的最多文件描述符数目，理论值应该是最多打开文件数（ulimit -n）与nginx进程数相除，但是nginx分配请求并不是那么均匀，所以最好与ulimit -n 的值保持一致。
#现在在linux 2.6内核下开启文件打开数为65535，worker_rlimit_nofile就相应应该填写65535。
#这是因为nginx调度时分配请求到进程并不是那么的均衡，所以假如填写10240，总并发量达到3-4万时就有进程可能超过10240了，这时会返回502错误。
worker_rlimit_nofile 65535;


events
{
    #参考事件模型，use [ kqueue | rtsig | epoll | /dev/poll | select | poll ]; epoll模型
    #是Linux 2.6以上版本内核中的高性能网络I/O模型，linux建议epoll，如果跑在FreeBSD上面，就用kqueue模型。
    #补充说明：
    #与apache相类，nginx针对不同的操作系统，有不同的事件模型
    #A）标准事件模型
    #Select、poll属于标准事件模型，如果当前系统不存在更有效的方法，nginx会选择select或poll
    #B）高效事件模型
    #Kqueue：使用于FreeBSD 4.1+, OpenBSD 2.9+, NetBSD 2.0 和 MacOS X.使用双处理器的MacOS X系统使用kqueue可能会造成内核崩溃。
    #Epoll：使用于Linux内核2.6版本及以后的系统。
    #/dev/poll：使用于Solaris 7 11/99+，HP/UX 11.22+ (eventport)，IRIX 6.5.15+ 和 Tru64 UNIX 5.1A+。
    #Eventport：使用于Solaris 10。 为了防止出现内核崩溃的问题， 有必要安装安全补丁。
    use epoll;

    #单个进程最大连接数（最大连接数=连接数*进程数）
    #根据硬件调整，和前面工作进程配合起来用，尽量大，但是别把cpu跑到100%就行。每个进程允许的最多连接数，理论上每台nginx服务器的最大连接数为。
    worker_connections 65535;

    #keepalive超时时间。
    keepalive_timeout 60;

    #客户端请求头部的缓冲区大小。这个可以根据你的系统分页大小来设置，一般一个请求头的大小不会超过1k，不过由于一般系统分页都要大于1k，所以这里设置为分页大小。
    #分页大小可以用命令getconf PAGESIZE 取得。
    #[root@web001 ~]# getconf PAGESIZE
    #4096
    #但也有client_header_buffer_size超过4k的情况，但是client_header_buffer_size该值必须设置为“系统分页大小”的整倍数。
    client_header_buffer_size 4k;

    #这个将为打开文件指定缓存，默认是没有启用的，max指定缓存数量，建议和打开文件数一致，inactive是指经过多长时间文件没被请求后删除缓存。
    open_file_cache max=65535 inactive=60s;

    #这个是指多长时间检查一次缓存的有效信息。
    #语法:open_file_cache_valid time 默认值:open_file_cache_valid 60 使用字段:http, server, location 这个指令指定了何时需要检查open_file_cache中缓存项目的有效信息.
    open_file_cache_valid 80s;

    #open_file_cache指令中的inactive参数时间内文件的最少使用次数，如果超过这个数字，文件描述符一直是在缓存中打开的，如上例，如果有一个文件在inactive时间内一次没被使用，它将被移除。
    #语法:open_file_cache_min_uses number 默认值:open_file_cache_min_uses 1 使用字段:http, server, location  这个指令指定了在open_file_cache指令无效的参数中一定的时间范围内可以使用的最小文件数,如果使用更大的值,文件描述符在cache中总是打开状态.
    open_file_cache_min_uses 1;
    
    #语法:open_file_cache_errors on | off 默认值:open_file_cache_errors off 使用字段:http, server, location 这个指令指定是否在搜索一个文件是记录cache错误.
    open_file_cache_errors on;
}
 
 
 
#设定http服务器，利用它的反向代理功能提供负载均衡支持
http
{
    #文件扩展名与文件类型映射表
    include mime.types;

    #默认文件类型
    default_type application/octet-stream;

    #默认编码
    #charset utf-8;

    #服务器名字的hash表大小
    #保存服务器名字的hash表是由指令server_names_hash_max_size 和server_names_hash_bucket_size所控制的。参数hash bucket size总是等于hash表的大小，并且是一路处理器缓存大小的倍数。在减少了在内存中的存取次数后，使在处理器中加速查找hash表键值成为可能。如果hash bucket size等于一路处理器缓存的大小，那么在查找键的时候，最坏的情况下在内存中查找的次数为2。第一次是确定存储单元的地址，第二次是在存储单元中查找键 值。因此，如果Nginx给出需要增大hash max size 或 hash bucket size的提示，那么首要的是增大前一个参数的大小.
    server_names_hash_bucket_size 128;

    #客户端请求头部的缓冲区大小。这个可以根据你的系统分页大小来设置，一般一个请求的头部大小不会超过1k，不过由于一般系统分页都要大于1k，所以这里设置为分页大小。分页大小可以用命令getconf PAGESIZE取得。
    client_header_buffer_size 32k;

    #客户请求头缓冲大小。nginx默认会用client_header_buffer_size这个buffer来读取header值，如果header过大，它会使用large_client_header_buffers来读取。
    large_client_header_buffers 4 64k;

    #设定通过nginx上传文件的大小
    client_max_body_size 8m;

    #开启高效文件传输模式，sendfile指令指定nginx是否调用sendfile函数来输出文件，对于普通应用设为 on，如果用来进行下载等应用磁盘IO重负载应用，可设置为off，以平衡磁盘与网络I/O处理速度，降低系统的负载。注意：如果图片显示不正常把这个改成off。
    #sendfile指令指定 nginx 是否调用sendfile 函数（zero copy 方式）来输出文件，对于普通应用，必须设为on。如果用来进行下载等应用磁盘IO重负载应用，可设置为off，以平衡磁盘与网络IO处理速度，降低系统uptime。
    sendfile on;

    #开启目录列表访问，合适下载服务器，默认关闭。
    autoindex on;

    #此选项允许或禁止使用socke的TCP_CORK的选项，此选项仅在使用sendfile的时候使用
    tcp_nopush on;
     
    tcp_nodelay on;

    #长连接超时时间，单位是秒
    keepalive_timeout 120;

    #FastCGI相关参数是为了改善网站的性能：减少资源占用，提高访问速度。下面参数看字面意思都能理解。
    fastcgi_connect_timeout 300;
    fastcgi_send_timeout 300;
    fastcgi_read_timeout 300;
    fastcgi_buffer_size 64k;
    fastcgi_buffers 4 64k;
    fastcgi_busy_buffers_size 128k;
    fastcgi_temp_file_write_size 128k;

    #gzip模块设置
    gzip on; #开启gzip压缩输出
    gzip_min_length 1k;    #最小压缩文件大小
    gzip_buffers 4 16k;    #压缩缓冲区
    gzip_http_version 1.0;    #压缩版本（默认1.1，前端如果是squid2.5请使用1.0）
    gzip_comp_level 2;    #压缩等级
    gzip_types text/plain application/x-javascript text/css application/xml;    #压缩类型，默认就已经包含textml，所以下面就不用再写了，写上去也不会有问题，但是会有一个warn。
    gzip_vary on;

    #开启限制IP连接数的时候需要使用
    #limit_zone crawler $binary_remote_addr 10m;



    #负载均衡配置
    upstream jh.w3cschool.cn {
     
        #upstream的负载均衡，weight是权重，可以根据机器配置定义权重。weigth参数表示权值，权值越高被分配到的几率越大。
        server 192.168.80.121:80 weight=3;
        server 192.168.80.122:80 weight=2;
        server 192.168.80.123:80 weight=3;

        #nginx的upstream目前支持4种方式的分配
        #1、轮询（默认）
        #每个请求按时间顺序逐一分配到不同的后端服务器，如果后端服务器down掉，能自动剔除。
        #2、weight
        #指定轮询几率，weight和访问比率成正比，用于后端服务器性能不均的情况。
        #例如：
        #upstream bakend {
        #    server 192.168.0.14 weight=10;
        #    server 192.168.0.15 weight=10;
        #}
        #2、ip_hash
        #每个请求按访问ip的hash结果分配，这样每个访客固定访问一个后端服务器，可以解决session的问题。
        #例如：
        #upstream bakend {
        #    ip_hash;
        #    server 192.168.0.14:88;
        #    server 192.168.0.15:80;
        #}
        #3、fair（第三方）
        #按后端服务器的响应时间来分配请求，响应时间短的优先分配。
        #upstream backend {
        #    server server1;
        #    server server2;
        #    fair;
        #}
        #4、url_hash（第三方）
        #按访问url的hash结果来分配请求，使每个url定向到同一个后端服务器，后端服务器为缓存时比较有效。
        #例：在upstream中加入hash语句，server语句中不能写入weight等其他的参数，hash_method是使用的hash算法
        #upstream backend {
        #    server squid1:3128;
        #    server squid2:3128;
        #    hash $request_uri;
        #    hash_method crc32;
        #}

        #tips:
        #upstream bakend{#定义负载均衡设备的Ip及设备状态}{
        #    ip_hash;
        #    server 127.0.0.1:9090 down;
        #    server 127.0.0.1:8080 weight=2;
        #    server 127.0.0.1:6060;
        #    server 127.0.0.1:7070 backup;
        #}
        #在需要使用负载均衡的server中增加 proxy_pass http://bakend/;

        #每个设备的状态设置为:
        #1.down表示单前的server暂时不参与负载
        #2.weight为weight越大，负载的权重就越大。
        #3.max_fails：允许请求失败的次数默认为1.当超过最大次数时，返回proxy_next_upstream模块定义的错误
        #4.fail_timeout:max_fails次失败后，暂停的时间。
        #5.backup： 其它所有的非backup机器down或者忙的时候，请求backup机器。所以这台机器压力会最轻。

        #nginx支持同时设置多组的负载均衡，用来给不用的server来使用。
        #client_body_in_file_only设置为On 可以讲client post过来的数据记录到文件中用来做debug
        #client_body_temp_path设置记录文件的目录 可以设置最多3层目录
        #location对URL进行匹配.可以进行重定向或者进行新的代理 负载均衡
    }
     
     
     
    #虚拟主机的配置
    server
    {
        #监听端口
        listen 80;

        #域名可以有多个，用空格隔开
        server_name www.w3cschool.cn w3cschool.cn;
        index index.html index.htm index.php;
        root /data/www/w3cschool;

        #对******进行负载均衡
        location ~ .*.(php|php5)?$
        {
            fastcgi_pass 127.0.0.1:9000;
            fastcgi_index index.php;
            include fastcgi.conf;
        }
         
        #图片缓存时间设置
        location ~ .*.(gif|jpg|jpeg|png|bmp|swf)$
        {
            expires 10d;
        }
         
        #JS和CSS缓存时间设置
        location ~ .*.(js|css)?$
        {
            expires 1h;
        }
         
        #日志格式设定
        #$remote_addr与$http_x_forwarded_for用以记录客户端的ip地址；
        #$remote_user：用来记录客户端用户名称；
        #$time_local： 用来记录访问时间与时区；
        #$request： 用来记录请求的url与http协议；
        #$status： 用来记录请求状态；成功是200，
        #$body_bytes_sent ：记录发送给客户端文件主体内容大小；
        #$http_referer：用来记录从那个页面链接访问过来的；
        #$http_user_agent：记录客户浏览器的相关信息；
        #通常web服务器放在反向代理的后面，这样就不能获取到客户的IP地址了，通过$remote_add拿到的IP地址是反向代理服务器的iP地址。反向代理服务器在转发请求的http头信息中，可以增加x_forwarded_for信息，用以记录原有客户端的IP地址和原来客户端的请求的服务器地址。
        log_format access '$remote_addr - $remote_user [$time_local] "$request" '
        '$status $body_bytes_sent "$http_referer" '
        '"$http_user_agent" $http_x_forwarded_for';
         
        #定义本虚拟主机的访问日志
        access_log  /usr/local/nginx/logs/host.access.log  main;
        access_log  /usr/local/nginx/logs/host.access.404.log  log404;
         
        #对 "/" 启用反向代理
        location / {
            proxy_pass http://127.0.0.1:88;
            proxy_redirect off;
            proxy_set_header X-Real-IP $remote_addr;
             
            #后端的Web服务器可以通过X-Forwarded-For获取用户真实IP
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
             
            #以下是一些反向代理的配置，可选。
            proxy_set_header Host $host;

            #允许客户端请求的最大单文件字节数
            client_max_body_size 10m;

            #缓冲区代理缓冲用户端请求的最大字节数，
            #如果把它设置为比较大的数值，例如256k，那么，无论使用firefox还是IE浏览器，来提交任意小于256k的图片，都很正常。如果注释该指令，使用默认的client_body_buffer_size设置，也就是操作系统页面大小的两倍，8k或者16k，问题就出现了。
            #无论使用firefox4.0还是IE8.0，提交一个比较大，200k左右的图片，都返回500 Internal Server Error错误
            client_body_buffer_size 128k;

            #表示使nginx阻止HTTP应答代码为400或者更高的应答。
            proxy_intercept_errors on;

            #后端服务器连接的超时时间_发起握手等候响应超时时间
            #nginx跟后端服务器连接超时时间(代理连接超时)
            proxy_connect_timeout 90;

            #后端服务器数据回传时间(代理发送超时)
            #后端服务器数据回传时间_就是在规定时间之内后端服务器必须传完所有的数据
            proxy_send_timeout 90;

            #连接成功后，后端服务器响应时间(代理接收超时)
            #连接成功后_等候后端服务器响应时间_其实已经进入后端的排队之中等候处理（也可以说是后端服务器处理请求的时间）
            proxy_read_timeout 90;

            #设置代理服务器（nginx）保存用户头信息的缓冲区大小
            #设置从被代理服务器读取的第一部分应答的缓冲区大小，通常情况下这部分应答中包含一个小的应答头，默认情况下这个值的大小为指令proxy_buffers中指定的一个缓冲区的大小，不过可以将其设置为更小
            proxy_buffer_size 4k;

            #proxy_buffers缓冲区，网页平均在32k以下的设置
            #设置用于读取应答（来自被代理服务器）的缓冲区数目和大小，默认情况也为分页大小，根据操作系统的不同可能是4k或者8k
            proxy_buffers 4 32k;

            #高负荷下缓冲大小（proxy_buffers*2）
            proxy_busy_buffers_size 64k;

            #设置在写入proxy_temp_path时数据的大小，预防一个工作进程在传递文件时阻塞太长
            #设定缓存文件夹大小，大于这个值，将从upstream服务器传
            proxy_temp_file_write_size 64k;
        }
         
         
        #设定查看Nginx状态的地址
        location /NginxStatus {
            stub_status on;
            access_log on;
            auth_basic "NginxStatus";
            auth_basic_user_file confpasswd;
            #htpasswd文件的内容可以用apache提供的htpasswd工具来产生。
        }
         
        #本地动静分离反向代理配置
        #所有jsp的页面均交由tomcat或resin处理
        location ~ .(jsp|jspx|do)?$ {
            proxy_set_header Host $host;
            proxy_set_header X-Real-IP $remote_addr;
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
            proxy_pass http://127.0.0.1:8080;
        }
         
        #所有静态文件由nginx直接读取不经过tomcat或resin
        location ~ .*.(htm|html|gif|jpg|jpeg|png|bmp|swf|ioc|rar|zip|txt|flv|mid|doc|ppt|
        pdf|xls|mp3|wma)$
        {
            expires 15d; 
        }
         
        location ~ .*.(js|css)?$
        {
            expires 1h;
        }
    }
}
######Nginx配置文件nginx.conf中文详解#####
```

### 四、反向代理

#### 4.1 单端口反向代理

实现效果：打开浏览器，在浏览器地址栏输入 www.123.com ，跳转到 linux 系统的 tomcat 主页

##### 4.1.1  安装 tomcat 

1. 将Tomcat安装包复制到linux中 

2. 进入 tomcat 的 bin 目录中，./startup.sh 启动 tomcat 服务器

3. 对外开放访问的端口号

   ```bash
   firewall-cmd --add-port=8080/tcp --permanent
   firewall-cmd -reload
   ```

4. 修改本机端口映射

```bash
#1.修改本机 hosts 文件
增加：192.168.0.105 www.123.com
```

5. 修改nginx配置文件

![7](images/7.png)

6. 访问测试

<img src="images/8.png" alt="8" style="zoom:50%;" />

#### 4.2 多端口反向代理

##### 4.2.1 实现效果：

```bash
访问 http://192.168.0.105:9001/edu/ 直接跳转到 127.0.0.1:8080
访问 http://192.168.0.105:9001/vod/ 直接跳转到 127.0.0.1:8081
```

##### 4.2.2 安装两个 Tomcat 分别定义端口为 8001,8002

![9](images/9.png)

##### 4.2.3 创建文件夹和测试页面

![10](images/10.png)

##### 4.2.4 配置 nginx配置文件

![11](images/11.png)

##### 4.2.5 开放端口号

```bash
firewall-cmd --add-port=8001/tcp --permanent

firewall-cmd --add-port=8002/tcp --permanent

firewall-cmd --add-port=9001/tcp --permanent

firewall-cmd --reload
```

![12](images/12.png)

##### 4.2.6 测试

> 1、访问：http://192.168.0.105:9001/edu/a.html
>
> ![13](images/13.png)
>
> 2、访问：http:// 192.168.0.105:9001/vod/a.html
>
> ![14](images/14.png)

#### 4.3 负载均衡

1、准备两台 tomcat 服务器，一台8001，一台8002

2、 在两台 tomcat 里面 webapps 目录中，创建名称是 edu 文件夹，在 edu 文件夹中创建页面a.html，用于测试

3、配置nginx配置文件

```bash
# 在http模块中配置
upstream myserver{
         server  192.168.0.105:8001 weight=1;
         server  192.168.0.105:8002 weight=2;
    }

# 在server模块配置
         listen       80;
         server_name  192.168.0.105;
         location / {
            proxy_pass http://myserver;
            root   html;
            index  index.html index.htm;
        }
```

4、nginx提供了几种分配方式(策略)

```bash
#1.轮询（默认）
每个请求按时间顺序逐一分配到不同的后端服务器，如果后端服务器 down 掉，能自动剔除。
#2.weight
weight 代表权重，默认是1，权重越高被分配的客户端越多。
指定轮询几率，weight和访问比率成正比，用于后端服务器性能不均的情况。例如：

upstream server_pool{
	server 192.168.5.21 weight=10;
	server 192.168.5.22 weight=10;
}
#3. ip_hash
每个请求按访问 ip 的 hash 结果分配，这样每个访客固定一个后端服务器，可以解决session的问题。例如：
upstream server_pool{
	ip_hash
	server 192.168.5.21:80;
	server 192.168.5.22:80;
}
#4.fair(第三方，需要安装第三方模块)
按后端服务器的响应时间来分配请求，响应时间短的优先分配。
upstream server_pool{
	server 192.168.5.21:80;
	server 192.168.5.22:80;
	fair;
}
```

5、测试

![15](images/15.png)

### 五、动静分离

#### 5.1 什么是动静分离

 Nginx 动静分离简单来说就是把动态跟静态请求分开，不能理解成只是单纯的把动态页面和

静态页面物理分离。

![16](images/16.png)

#### 5.2 在linux系统中准备静态资源，用于进行访问

![17](images/17.png)

#### 5.3 nginx配置文件

![18](images/18.png)

#### 5.4 测试

![19](images/19.png)

![20](images/20.png)

### 六、高可用

#### 6.1 Keeplived+Nginx高可用集群（主从模式）

![21](images/21.png)

1、需要两台服务器 192.168.0.105 和 192.168.0.102

![22](images/22.png)

2、在两台服务器上安装nginx和keeplived

> #1.安装keepalived
> 	yum install keepalived -y
> #2.keepalived.conf文件
> 	安装之后，在etc里面生成目录 keepalived，有文件 keepalived.conf
> #3.修改配置文件
>
> 105主服务器
>
> ![23](images/23.png)
>
> 102副服务器
>
> ![24](images/24.png)
>
> #4. 启动keepalived
>
> ![25](images/25.png)
>
> #5.测试
>
> ![26](images/26.png)
>
> ![27](images/27.png)





#### 6.2 Keeplived+Nginx 高可用集群（双主模式）

![28](images/28.png)

修改配置

![graphic](images/graphic.png)

**（2） 配置 LB-02 节点**

![graphic2](images/graphic2.png)