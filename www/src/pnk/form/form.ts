import { safeSet, isSet } from 'pnk/core/objects';
import { Form, FormValue, ApiError, ApiResponse } from './types';

export function updateForm(model: {[key: string]: any}, path: string, value: any) {
  model.update((f) => {
    const result = {...f};
    safeSet(result, path, value);

    const isValid = !Object.values(result.fields).some((field: FormValue<any>) => !field.isValid);
    safeSet(result, 'isValid', isValid);

    return result;
  });
}

export function updateFieldValue(model: {[key: string]: any}, field: string, value: any) {
  updateForm(model, `fields.${field}.value`, value);
}

export function updateFieldError(model: {[key: string]: any}, field: string, error: string) {
  updateForm(model, `fields.${field}.error`, error);
  updateForm(model, `fields.${field}.isValid`, error === null);
}

export function serializeForm(form: Form): {[key: string]: string|number|boolean} {
  const result = {};

  for (let key in form.fields) {
    result[key] = form.fields[key].value;
  }

  return result;
}

export const sendForm = (form, csrfToken, data): Promise<ApiResponse|ApiError> => {
  const headers = {
    'Content-Type': 'application/json',
  };

  const additionalData = data && data.toJSON ? {data: data.toJSON()} : {};
  const payload = {...serializeForm(form), ...additionalData}

  if (isSet(csrfToken)) {
    headers['X-CSRF-Token'] = csrfToken;
  }

  return fetch(form.action, {
    body: JSON.stringify(payload),
    cache: 'no-cache',
    headers,
    method: 'POST',
    mode: 'cors',
  }).then((data) => data.json())
  .then((json) => json as ApiResponse)
  .catch((json) => json as ApiError);
}
