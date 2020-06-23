export type Form = {
  action: string;
  fields: {[key: string]: FormValue<string|number|boolean>};
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
