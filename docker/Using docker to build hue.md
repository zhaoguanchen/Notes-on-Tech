#大标题



git clone https://github.com/cloudera/hue.git
cd hue
make apps
build/env/bin/hue runserver
挂载外部  
docker run -it -p 8888:8888 -v /Users/yiche/Documents/docker:/hue1 gethue/hue:latest bash
/Users/yiche/Documents/docker


##而标题
hue数据库连接：
新建数据库---在hue.ini指定连接---初始化数据库
进入目录：/home/hadoop/app/hue-3.9.0-cdh5.5.4/build/env
 bin/hue syncdb
 bin/hue migrate
Xiugai
