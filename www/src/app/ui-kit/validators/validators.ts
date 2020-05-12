export function validName(value: string): string {
  const reName = /[a-zа-я]{2,}\s[a-zа-я]{2,}/i;

  return value.match(reName) === null
    ? 'Enter your first and last names'
    : null;
}

export function validEmail(value: string): string {
  const reEmail = /^[a-z0-9_.+-]+@[a-z0-9-]+\.[a-z]{2,5}$/i;

  return value.match(reEmail) === null
    ? 'Enter valid email please'
    : null;
}

export function validPhone(value: string): string {
  const rePhone = /^[0-9]{10}$/i;

  return value.match(rePhone) === null
    ? 'Enter correct phone number'
    : null;
}
