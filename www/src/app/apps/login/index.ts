import { Writable, writable } from 'svelte/store';
import { FormValue } from 'pnk/form/types';
import App from './app.svelte';

export type LoginFormStore = {
  action: string;
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
    LoginApp.resetStore();
    LoginApp.app = new App({target});
  }

  public static resetStore() {
    LoginApp.store = writable({
      action: '/api/auth/',
      isTouched: false,
      isValid: true,
      fields: {
        login: { value: '', isValid: true, error: null },
        password: { value: '', isValid: true, error: null },
      }
    });
  }

  public static openAdmin() {
    window.location.href = "/squat/";
  }
}
