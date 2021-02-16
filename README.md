# 1.静态文件处理
## 1.字体
> fyne默认字体不支持中文，故而采用自己下载的阿里巴巴普惠字体进行绑定替换
```shell
# 1.进入source目录下，修改生成的字体文件
# 2.执行
go run generator.go
```

# 2.Android打包
```shell
# 在main目录下(name不能有中文，特殊字符)
fyne package -os android -appID com.kokutas.app -name kokutas -appVersion 2.0 -icon source/ali.png
```