/**
 * @param {number[]} nums
 * @return {number}
 */
const findNumberOfLIS = function (nums) {
    if (nums.length === 0) {
        return 0;
    }

    let numOfLIS = 1; // 最长子序列的数量
    let lengthOfLIS = 1; // 最长子序列的长度

    const numOfItem = [1]; // 以 item[i] 结尾的最长子序列的数量
    const lengthOfItem = [1]; // 以 item[i] 结尾的最长子序列的长度

    for (let i = 1; i < nums.length; i++) {
        let num = 1;
        let length = 1;

        for (let j = 0; j < i; j++) {
            if (nums[i] > nums[j]) {
                if (lengthOfItem[j] + 1 === length) {
                    num += numOfItem[j];
                } else if (lengthOfItem[j] + 1 > length) {
                    num = numOfItem[j];
                    length = lengthOfItem[j] + 1;
                }
            }
        }

        numOfItem.push(num);
        lengthOfItem.push(length);

        if (length === lengthOfLIS) {
            numOfLIS += num;
        } else if (length > lengthOfLIS) {
            numOfLIS = num;
            lengthOfLIS = length;
        }
    }

    return numOfLIS;
};

module.exports = findNumberOfLIS;
