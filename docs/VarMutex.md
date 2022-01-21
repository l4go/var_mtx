type VarMutex
===

interface{}で比較できる値を粒度とした排他ロックを提供するライブラリです。

### import
```
import "github.com/l4go/var_mtx"
```
vendoringして使うことを推奨します。

## 利用サンプル

[example](../examples/ex_var_mtx/ex_var_mtx.go)

## メソッド概略

### func NewVarMutex() \*VarMutex

\*VarMutex型を生成します。

### func (km \*VarMutex) Lock(key interface{})

指定したkeyのロックを取得します。  
ロックが取得できない場合、取得できるまでブロックします。

### func (km \*VarMutex) Unlock(key interface{})

指定したkeyのロックを開放します。

