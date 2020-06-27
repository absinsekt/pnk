import { isUnset } from "pnk/core/objects";

export function validName(errorMessage: string) {
  return function(value: string): string {
    const re = /[a-zа-я]{2,}/i;
    return value.match(re) === null ? errorMessage : null;
  }
}

export function validEmail(errorMessage: string) {
  return function(value: string): string {
    const re = /^[a-z0-9_.+-]+@[a-z0-9-]+\.[a-z]{2,5}$/i;
    return value.match(re) === null ? errorMessage : null;
  }
}

export function validPhone(errorMessage: string) {
  return function(value: string): string {
    const re = /^[0-9]{10}$/i;
    return value.match(re) === null ? errorMessage : null;
  }
}

export function required(errorMessage: string) {
  return function(value: string|number|Date): string {
    return isUnset(value) ? errorMessage : null;
  }
}
