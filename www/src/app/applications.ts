import { LoginApp } from './apps/login';


export function initApps() {
  // init containers
  const containerLogin = document.getElementById('app-login');

  // check containers & init applications
  if (containerLogin !== null) {
    LoginApp.init(containerLogin);
  }
}
