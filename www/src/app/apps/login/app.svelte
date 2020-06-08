<script>
  import Form from 'ui/form/form.svelte';
  import Inline from 'ui/inputs/inline.svelte';
  import Submit from 'ui/submit/submit.svelte';

  import { buildRequiredValidator } from 'ui/validators/validators';

  import { LoginApp } from '.';
</script>

<Form
  action="/api/login/"
  store={LoginApp.store}
  token={window._CSRF_TOKEN}
  on:success={() => LoginApp.openAdmin()}
  on:error={(err) => alert(err.message)}>

  <Inline label="login"
    name="login" store={LoginApp.store}
    validators={[buildRequiredValidator('Enter login')]} />

  <Inline label="password"
    name="password" store={LoginApp.store} type="password"
    validators={[buildRequiredValidator('Enter password')]} />

  <Submit store={LoginApp.store} />
</Form>
