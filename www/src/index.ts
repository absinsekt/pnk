import '@babel/polyfill';
import App from './app/app.module.svelte';

export default new App({
    target: document.body,
});
