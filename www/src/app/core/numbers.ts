export const pad = (num: number, size: number): string => {
  let result = num + "";

  while (result.length < size) result = "0" + result;
  return result;
};

export const humanize = (num, divider = " ") => {
  const [int, fract] = num.toString().split('.')
  const result = [];

  for (let i = int.length - 3; i > -3; i -= 3) {
    const sl = i > 0 ? int.slice(i, i + 3) : int.slice(0, i + 3);

    result.push(sl);
  }

  return typeof(fract) === 'undefined'
    ? result.reverse().join(divider)
    : [result.reverse().join(divider), fract].join('.');
};

export const toCurrency = (num: number, sign: string, join: string = '', prefixed = false) => {
  const humanized = humanize(num);
  return prefixed
    ? `${sign}${join}${humanized}`
    : `${humanized}${join}${sign}`;
}
