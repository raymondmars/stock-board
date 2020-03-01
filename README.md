### 雪球股价通知 

这是一个股票价格 Notifier，它会定时读取雪球上的股票价格，并在超过你设置的价格时，通知你。   
目前支持A股，港股，美股的股票通知；其通知方式可以自行修改，但是内置了一个第三方的微信通知，你可以在 [这里](http://sc.ftqq.com/?c=wechat&a=bind) 注册一个账号，获得其通知Token，然后在程序启动时作为环境变量传入即可。  

### 如何运行：  
推荐采用 docker 的方式部署并运行，可以参考如下的 docker-compose.yml 文件： 
```
  version: '3.5'
  x-logging:
    &default-logging
    driver: "json-file"
    options:
      max-size: "100m"
      max-file: "10"

  services:
    xueqiu-notifier:
      build: .
      container_name: xueqiu-notifier
      image: xueqiu-notifier:latest
      restart: always
      environment:
        CONFIG_TARGETS: "01810,23|000166,5.9|000725,6.5"
        INFORM_URL: "YOUR TOKEN..."

      logging: *default-logging    
```   

CONFIG_TARGETS  
  为要通知的股票列表，多个列表用 | 分割；每个股票配置数据由两部分组成，第一部分为：股票代码，第二部分为：设定的目标价格，两部分之间用 , 分割。    
INFORM_URL  
  目前为一个内置的第三方微信通知工具的Token URL   

  




