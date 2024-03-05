# minhash

golang Minhash实现

## 使用方式

```go
a := "flying fish flew by the space station"
b := "we will not allow you to bring your pet armadillo along"
c := "he figured a few sticks of dynamite were easier than a fishing pole to catch fish"
h := minhash.New(32)
sa := h.Add(a)
sb := h.Add(b)
sc := h.Add(c)
ha := h.Hash(sa)
hb := h.Hash(sb)
hc := h.Hash(sc)
fmt.Println(minhash.Jaccard(ha, hb))
fmt.Println(minhash.Jaccard(ha, hc))
```

数值越大表示两句句子越接近