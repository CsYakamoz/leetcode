package golang

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
