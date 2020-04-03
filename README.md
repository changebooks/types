# types
类型
==

<pre>
num := alphabet.Num('c')
fmt.Println(num)

sum := alphabet.Sum("abcxyz")
fmt.Println(sum)
</pre>

<pre>
keyExist := dict.KeyExist("abc", map[string]int{"abc": 1, "def": 2})
fmt.Println(keyExist)

keyExist2 := dict.KeyExist("ghi", map[string]int{"abc": 1, "def": 2})
fmt.Println(keyExist2)

valueExist, key := dict.ValueExist(1, map[string]int{"abc": 1, "def": 2})
fmt.Println(valueExist, key)

valueExist2, key2 := dict.ValueExist(3, map[string]int{"abc": 1, "def": 2})
fmt.Println(valueExist2, key2)

keys := dict.Keys(map[string]int{"abc": 1, "def": 2})
fmt.Println(keys)

values := dict.Values(map[string]int{"abc": 1, "def": 2})
fmt.Println(values)

data := dict.Merge(map[string]int{"abc": 1, "def": 2}, map[string]int{"hij": 3, "klm": 4})
fmt.Println(data)
</pre>

<pre>
keyExist := sparse.KeyExist(1, map[int]int{1: 10, 2: 20})
fmt.Println(keyExist)

valueExist, key := sparse.ValueExist(20, map[int]int{1: 10, 2: 20})
fmt.Println(valueExist, key)

keys := sparse.Keys(map[int]int{1: 10, 2: 20})
fmt.Println(keys)

values := sparse.Values(map[int]int{1: 10, 2: 20})
fmt.Println(values)

data := sparse.Merge(map[int]int{1: 10, 2: 20}, map[int]int{3: 30, 4: 40})
fmt.Println(data)
</pre>

<pre>
exist := ints.InArr(1, []int{1, 2, 3})
fmt.Println(exist)

data := ints.Merge([]int{1, 2}, []int{3}, []int{4, 5})
fmt.Println(data)

data2 := ints.Reverse([]int{1, 2, 3, 4, 5})
fmt.Println(data2)
</pre>