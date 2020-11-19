# 基于chain33开发java合约接口
## 1.账户接口
com/fuzamei/chain33/Account.java

## 2.区块链接口
com/fuzamei/chain33/Blockchain.java

## 3.本地数据库接口
com/fuzamei/chain33/LocalDB.java

## 5.世界状态数据库接口
com/fuzamei/chain33/StateDB.java

## 4.参考合约实现
可以参考合约guess的实现，代码见
com/fuzamei/chain33/guess

###1.Tx合约交易接口为
public static void tx(String[] args)

###2.Query查询接口为
public static String[] query(String[] args)

