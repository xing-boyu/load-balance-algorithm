The package load-balance-algorithm implements two kinds of Round-Robin algorithm.

- [Smooth weighted round-robin balancing algorithm](https://github.com/phusion/nginx/commit/27e94984486058d73157038f7950a0a36ecc6e35?diff=split#diff-3f2250b728a3f5fe1e2d31cbf63c2268)

example to use it:

```
func SmoothRoundRobinAlg_Next(t *testing.T) {

	smoothRoundRobinAlg := &SmoothRoundRobinAlg{}
	smoothRoundRobinAlg.Add("A", 5)
	smoothRoundRobinAlg.Add("B", 2)
	smoothRoundRobinAlg.Add("C", 3)

	for i:=0; i<10; i++ {
    	fmt.Printf("%s ", roundRobinAlg.Next())
    }
}
```


- [Simple weighted Round-Robin algorithm](http://kb.linuxvirtualserver.org/wiki/Weighted_Round-Robin_Scheduling)

example to use it:

```
func RoundRobinAlg_Next() {

	roundRobinAlg := &RoundRobinAlg{}
	roundRobinAlg.Add("A", 5)
	roundRobinAlg.Add("B", 2)
	roundRobinAlg.Add("C", 3)

	for i:=0; i<10; i++ {
		fmt.Printf("%s ", roundRobinAlg.Next())
	}
}
```