type VarRWMutex
===

interface{}で比較できる値を粒度とした、リードライト(Read/Write)ロックを提供するライブラリです。

### import
```
import "github.com/l4go/var_mtx"
```
vendoringして使うことを推奨します。

## 利用サンプル

[example](../examples/ex_var_mtx2/ex_var_mtx2.go)

## メソッド概略

### func NewVarRWMutex() \*VarRWMutex

\*VarRWMutex型を生成します。

### func (km \*VarRWMutex) Lock(key interface{})

指定したkeyの排他(Write)ロックを取得します。  
ロックが取得できない場合、取得できるまでブロックします。

### func (km \*VarRWMutex) Unlock(key interface{})

指定したkeyの排他(Write)ロックを開放します。

### func (km \*VarRWMutex) RLock(key interface{})

指定したkeyの共有(Read)ロックを取得します。  
ロックが取得できない場合、取得できるまでブロックします。

### func (km \*VarRWMutex) RUnlock(key interface{})

指定したkeyの共有(Read)ロックを開放します。

