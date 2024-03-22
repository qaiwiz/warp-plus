/**
 * Calculates the factorial of a given number.
 *
 * @param {number} n - The number to calculate the factorial of.
 * @return {number} The factorial of the given number.
 */
function factorial(n) {
  // If the input is not a positive integer, throw an error.
  if (n < 0 || !Number.isInteger(n)) {
    throw new Error('Input must be a positive integer');
  }

  // If the input is 0 or 1, return 1 as the factorial.
  if (n <= 1) {
    return 1;
  }

  // Calculate the factorial by recursively multiplying the input
  // with the factorial of the input minus one.
  return n * factorial(n - 1);
}
