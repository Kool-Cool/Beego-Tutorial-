# Beego-Tutorial-
Step by Step guide for Beego Framework 


## Installing [Beego](https://beego.wiki/docs/intro/introduction/)

- ### S1] Check for `bee` version
  in terminal
 ```
 bee version
  ```

if result is like continue 
```
 [D]  init global config instance failed. If you do not use this, just ignore it.  open conf/app.conf: The system cannot find the path specified.
______
| ___ \
| |_/ /  ___   ___
| ___ \ / _ \ / _ \
| |_/ /|  __/|  __/
\____/  \___| \___| v2.1.0
```

else continue with step S2

- ### S2] installing beego
  
```
go install github.com/beego/bee/v2@latest

bee version 
```

## Initializing Beego 

- S1]
```
bee new myapp
```

- S2] 
```
cd myapp
```
- S3]
```
go get myapp
```

- S4]
```
bee run
```

- S5] go to [http://:8080](http://localhost:8080/)
  
Result should be 

![](https://github.com/Kool-Cool/Beego-Tutorial-/blob/main/images/Screenshot%202024-05-05%20111630.png)
