# 二分查找法

**Notice**:

1. 当变量命名为 `begin`, `end` 时, 此乃左闭右开区间;

   `begin, end := 0, len(arr)`, 循环条件为 `begin < end`;

2. 当变量命名为 `first`, `last` 时, 此乃闭区间;

   `first, last := 0, len(arr)-1`, 循环条件为 `first <= last`;

3. `mid` 必须为 `(end-begin)/2 + begin`, 因为 `(begin+end)/2` 中的 `(begin+end)` 存在溢出的可能性;

   > `first`, `last` 同理;

## 在有序数组中寻找特定值的下标

若找到, 则返回对应下标, 否则返回 `-1`

```go
func binarySearch(arr []int, target int) int {
	begin, end := 0, len(arr)

	for begin < end {
		mid := (end-begin)/2 + begin

		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			end = mid
		} else {
			begin = mid + 1
		}
	}

	return -1
}
```

## 在有序数组中寻找第一个大于(或大于等于)特定值的下标

若存在, 则返回对应下标, 否则返回数组的长度 - `len(arr)`

```diff
func binarySearch(arr []int, target int) int {
	begin, end := 0, len(arr)

	for begin < end {
		mid := (end-begin)/2 + begin

-		if arr[mid] == target {
-			return mid
-		} else if arr[mid] > target {
+		// 若寻找的是第一个大于, 则比较符号为 >
+		// 若寻找的是第一个大于等于, 则比较符号为 >=
+		if arr[mid] > target {
			end = mid
		} else {
			begin = mid + 1
		}
	}

-	return -1
+	return end
}
```

## 在有序数组中寻找最后一个小于(或小于等于)特定值的下标

若存在, 则返回对应下标, 否则返回 `-1`

```diff
func binarySearch(arr []int, target int) int {
	begin, end := 0, len(arr)

	for begin < end {
		mid := (end-begin)/2 + begin

+		// 若寻找的是最后一个小于, 则比较符号为 >=
+		// 若寻找的是最后一个小于等于, 则比较符号为 >
+		if arr[mid] >= target {
			end = mid
		} else {
			begin = mid + 1
		}
	}

-	return end
+	return end - 1
}
```
