export type Form = {
  action: string;
  fields: FormValue[];
  isTouched: boolean;
  isValid: boolean;
}

export type FormValue<T = string> = {
  value: T;
  isValid: boolean;
  error: string;
}
