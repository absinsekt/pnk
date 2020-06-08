import { writable } from 'svelte/store';
import { FormValue } from 'ui/form/types';

export type OrderFormStore = {
  isTouched: boolean;
  isValid: boolean;
  fields: {
    name: FormValue<string>,
    email: FormValue<string>,
    phone: FormValue<string>,
    address: FormValue<string>,
    agree: FormValue<boolean>,
    paymentMethod: FormValue<number>,
  };
}

export const form = writable<OrderFormStore>({
  isTouched: false,
  isValid: true,
  fields: {
    name: { value: '', isValid: true, error: null },
    email: { value: '', isValid: true, error: null },
    phone: { value: '', isValid: true, error: null },
    address: {value: '', isValid: true, error: null },
    paymentMethod: {value: null, isValid: true, error: null },
    agree: { value: false, isValid: true, error: null },
  }
});
