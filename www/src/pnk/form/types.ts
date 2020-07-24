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

export const enum Status {
  Success = 'success',
  Error = 'error',
}

export type ApiResponse = {
  status: Status;
}

export type ApiSuccess<T= any> = ApiResponse & {
  data: {
    items: T[];
    count: number;
    offset: number;
  };
}

export type ApiError = ApiResponse & {
  message: string;
}
