export function buildNameValidator(errorMessage: string) {
  return function(value: string): string {
    const re = /[a-zа-я]{2,}\s[a-zа-я]{2,}/i;
    return value.match(re) === null ? errorMessage : null;
  }
}

export function buildEmailValidator(errorMessage: string) {
  return function(value: string): string {
    const re = /^[a-z0-9_.+-]+@[a-z0-9-]+\.[a-z]{2,5}$/i;
    return value.match(re) === null ? errorMessage : null;
  }
}

export function buildPhoneValidator(errorMessage: string) {
  return function(value: string): string {
    const re = /^[0-9]{10}$/i;
    return value.match(re) === null ? errorMessage : null;
  }
}

export function buildRequiredValidator(errorMessage: string) {
  return function(value: string): string {
    return value === '' ? errorMessage : null;
  }
}
