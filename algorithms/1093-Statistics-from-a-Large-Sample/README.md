## [1093. Statistics from a Large Sample](https://leetcode.com/problems/statistics-from-a-large-sample/)

### Description

<p>We sampled integers between <code>0</code> and <code>255</code>, and stored the results in an array <code>count</code>:&nbsp; <code>count[k]</code> is the number of integers we sampled equal to <code>k</code>.</p>

<p>Return the minimum, maximum, mean, median, and mode of the sample respectively, as an array of <strong>floating point numbers</strong>.&nbsp; The mode is guaranteed to be unique.</p>

<p><em>(Recall that the median of a sample is:</em></p>

<ul>
	<li><em>The middle element, if the elements of the sample were sorted and the number of elements is odd;</em></li>
	<li><em>The average of the middle two elements, if the elements of the sample were sorted and the number of elements is even.)</em></li>
</ul>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>
<pre><strong>Input:</strong> count = [0,1,3,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]
<strong>Output:</strong> [1.00000,3.00000,2.37500,2.50000,3.00000]
</pre><p><strong>Example 2:</strong></p>
<pre><strong>Input:</strong> count = [0,4,3,2,2,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]
<strong>Output:</strong> [1.00000,4.00000,2.18182,2.00000,1.00000]
</pre>
<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ol>
	<li><code>count.length == 256</code></li>
	<li><code>1 &lt;= sum(count) &lt;= 10^9</code></li>
	<li>The mode of the sample that count represents is unique.</li>
	<li>Answers within <code>10^-5</code> of the true value will be accepted as correct.</li>
</ol>

**Difficulty:** `Medium`

**Tags:** `Math` `Two Pointers`

### Solution One

```go
type FreqRecord struct {
	accFreq int
	freq    int
	val     int
}

func sampleStats(count []int) []float64 {
	record := make([]FreqRecord, 0)

	sum, minimum, maximum, mode, maxFreq, accFreq :=
		0, 256, -1, 0, 0, 0

	for i := 0; i < 256; i++ {
		if count[i] == 0 {
			continue
		}

		sum += i * count[i]
		accFreq += count[i]
		record = append(
			record,
			FreqRecord{accFreq: accFreq, val: i, freq: count[i]},
		)
		minimum = min(minimum, i)
		maximum = max(maximum, i)

		if count[i] > maxFreq {
			mode = i
			maxFreq = count[i]
		}
	}

	return []float64{
		float64(minimum), float64(maximum),
		float64(sum) / float64(accFreq),
		getMedian(record, accFreq), float64(mode),
	}
}

func getMedian(record []FreqRecord, accFreq int) float64 {
	if accFreq%2 == 0 {
		idx := binarySearch(record, accFreq/2)
		if record[idx].accFreq == accFreq/2 {
			return float64(record[idx].val+record[idx+1].val) / 2
		} else {
			return float64(record[idx].val)
		}
	} else {
		idx := binarySearch(record, (accFreq/2)+1)
		return float64(record[idx].val)
	}
}

func binarySearch(record []FreqRecord, target int) int {
	begin, end := 0, len(record)
	for begin < end {
		mid := (end-begin)/2 + begin
		if record[mid].accFreq >= target {
			end = mid
		} else {
			begin = mid + 1
		}
	}

	return end
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func max(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}
```
