const number = [121, -121, 10];

/**
 * @param {number} x
 * @return {boolean}
 */
var isPalindrome = function (x) {
  let reverse = x.toString().split("").reverse().join("");
  return reverse == x;
};

number.forEach((element) => {
  isPalindrome(element);
});
