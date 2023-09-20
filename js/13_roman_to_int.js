const cases = ["III", "LVIII", "MCMXCIV"];

const roman = {
  I: 1,
  V: 5,
  X: 10,
  L: 50,
  C: 100,
  D: 500,
  M: 1000,
};

/**
 * @param {string} s
 * @return {number}
 */
var romanToInt = function (s) {
  let total = 0;
  for (let i = 0; i < s.length; i++) {
    let current = roman[s[i]];
    let next = roman[s[i + 1]];
    if (current < next) {
      total += next - current;
      i++;
    } else {
      total += current;
    }
  }
  return total;
};

cases.forEach((element) => {
  console.log(romanToInt(element));
});
