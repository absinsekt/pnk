export const pad = (num: number, size: number): string => {
  let result = num + "";

  while (result.length < size) result = "0" + result;
  return result;
};

export function ID(length: number): string {
  const codes = [
    { offset: 48, length: 10 },
    { offset: 65, length: 26 },
    { offset: 97, length: 26 },
  ];

  let result = '';

  for (let i = 0; i < length; i++) {
    const palette = codes[Math.floor(Math.random() * 3)];

    result += String.fromCharCode(
      palette.offset + Math.floor(Math.random() * palette.length)
    );
  }

  return result;
}

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
