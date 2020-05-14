<script>
  import Form from 'app/ui-kit/form/form.svelte';
  import Input from 'app/ui-kit/input/input.svelte';
  import Submit from 'app/ui-kit/submit/submit.svelte';

  import { buildRequiredValidator } from 'app/ui-kit/validators/validators';

  import { LoginApp } from '.';
</script>

<Form
  action="/api/login/"
  store={LoginApp.store}
  token={window._CSRF_TOKEN}
  on:success={() => LoginApp.openAdmin()}
  on:error={(err) => alert(err.message)}>

  <Input label="login"
    name="login" store={LoginApp.store}
    validators={[buildRequiredValidator('Enter login')]} />

  <Input label="password"
    name="password" store={LoginApp.store} type="password"
    validators={[buildRequiredValidator('Enter password')]} />

  <Submit store={LoginApp.store} />
</Form>
