golib/var_mtx ライブラリ
===

`interface{}`の値の粒度で排他制御を提供するモジュール群です。  
動的に生成される値での細かな粒度の排他制御を簡単に実現するためのライブラリになります。

* [var_mtx.VarMutex](./VarMutex.md)
	* 値を粒度とした単純な排他ロックを提供します。
* [var_mtx.VarRWMutex](./VarRWMutex.md)
	* 値を粒度としたRead/Writeロック機構を提供します。
