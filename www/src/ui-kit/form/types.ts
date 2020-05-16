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

export type ApiResponse = {
  status: 'success';
  message: string;
  data: {
    items: [];
    count: number;
    offset: number;
  };
}

export type ApiError = {
  status: 'error';
  message: string;
}
