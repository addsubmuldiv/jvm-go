# jvm相关基本知识

## 运行时数据区结构
![img.png](img.png)

## 常量指令说明
`iconst`, `bipush`, `sipush`, `ldc` 四条指令分别对应于不同范围的整数，`iconst`为 -1 ~ 5，`bipush` -127 ~ 128,
`sipush` -32768 ~ 32767，`ldc` -2147483648 ~ 2147483647