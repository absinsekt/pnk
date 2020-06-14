import { Writable, writable } from 'svelte/store';
import { FormValue } from 'ui/form/types';
import App from './app.svelte';

export type LoginFormStore = {
  isTouched: boolean;
  isValid: boolean;
  fields: {
    login: FormValue<string>,
    password: FormValue<string>
  }
};

export class LoginApp {
  public static app: LoginApp;
  public static store: Writable<LoginFormStore>;

  public static init(target: HTMLElement) {
    LoginApp.store = writable({
      isTouched: false,
      isValid: true,
      fields: {
        login: { value: '', isValid: true, error: null },
        password: { value: '', isValid: true, error: null },
      }
    });

    LoginApp.app = new App({target});
  }

  public static openAdmin() {
    window.location.href = "/squat/";
  }
}
