/**
 * @param {number[]} nums1
 * @param {number} m
 * @param {number[]} nums2
 * @param {number} n
 * @return {void} Do not return anything, modify nums1 in-place instead.
 */
var merge = function (nums1, m, nums2, n) {
  let temp = nums1.concat(nums2);

  for (let i = 0; i < temp.length; i++) {
    if (temp[i] === 0) {
      temp.splice(i, 1);
      i--;
    }
  }
  nums1 = temp.sort((a, b) => a - b);
  console.log(nums1);
};

merge([1, 2, 3, 0, 0, 0], 3, [2, 5, 6], 3);
merge([0], 1, [1], 1);
