import { writable } from 'svelte/store';
import { FormValue } from 'ui/form/types';

export type OrderFormStore = {
  isTouched: boolean;
  isValid: boolean;
  fields: {
    name: FormValue<string>,
    email: FormValue<string>,
    phone: FormValue<string>,
    agree: FormValue<boolean>
  };
}

export const form = writable<OrderFormStore>({
  isTouched: false,
  isValid: true,
  fields: {
    name: { value: '', isValid: true, error: null },
    email: { value: '', isValid: true, error: null },
    phone: { value: '', isValid: true, error: null },
    agree: { value: false, isValid: true, error: null },
  }
});
