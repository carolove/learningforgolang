# learning golang by compare with java

# c && c++ debug tools [参考](https://www.jianshu.com/p/bb430259b7af)
## 工具集
- [cpptools](https://marketplace.visualstudio.com/items?itemName=ms-vscode.cpptools)
- [C/C++ Clang](https://marketplace.visualstudio.com/items?itemName=mitaki28.vscode-clang)

## 个人配置
```
"C_Cpp.autocomplete": "Disabled",
"clang.cxxflags": ["-std=c++11"]
```

## 程序配置
- command + p -> > c&&c++: configuration -> add includes
- debug launch -> add absolute path of the exec

## 正则匹配
- () 标记一个子表达式的开始和结束位置。子表达式可以获取供以后使用。要匹配这些字符，请使用 \( 和 \), 
- example: ,(\d) => \t$1

## 查询进程使用的fd相关信息
- lsof -p  9785 | awk '{print $9}'| awk -F"->" '{print $2}' | sort -n | uniq -c | head -n 50
