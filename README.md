# ChatGPTLineBot
ChatGPTLineBot developed with golang

[![Run on Google Cloud](https://storage.googleapis.com/cloudrun/button.svg)](https://console.cloud.google.com/cloudshell/editor?shellonly=true&cloudshell_image=gcr.io/cloudrun/button&cloudshell_git_repo=https://github.com/bowwowxx/ChatGPTLineBot.git)



### 1.Set parameters in the configuration file(.env)
```
ChannelSecret=xx  
ChannelAccessToken=xx  
OpenApiKey=xx  
```

### 2.Run ChatGPT LineBot
**docker run** 
```
docker run -itd -p 8090:8080 -v $PWD/.env:/ChatGPT/.env bowwow/chatgptlinebot 
```

or  

**go run** 

```
go run main.go
```

 **Result：**
  ![mole](https://github.com/bowwowxx/ChatGPTLineBot/blob/master/docker-run.jpg)  


### 3.Use ngrok tunnels to localhost

 ```
 ngrok  http 8090
 ```

 **Result：**
  ![mole](https://github.com/bowwowxx/ChatGPTLineBot/blob/master/ngork.jpg)  

### 4.Verify line bot webhook

 **Result：**
  ![mole](https://github.com/bowwowxx/ChatGPTLineBot/blob/master/linebot-webhook.jpg)  

### 5.enjoy it !
 **Result：**
  ![mole](https://github.com/bowwowxx/ChatGPTLineBot/blob/master/pikachu.jpg)  