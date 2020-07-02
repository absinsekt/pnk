<script>
  import Form from 'pnk/form/form.svelte';
  import Inline from 'pnk/inputs/inline.svelte';
  import Submit from 'pnk/buttons/submit.svelte';

  import { required } from 'pnk/validators/validators';

  import { LoginApp } from '.';
</script>

<Form
  store={LoginApp.store}
  token={window._CSRF_TOKEN}
  on:success={() => LoginApp.openAdmin()}
  on:error={(err) => alert(err.message)}>

  <Inline label="login"
    name="login" store={LoginApp.store}
    validators={[required('Enter login')]} />

  <Inline label="password"
    name="password" store={LoginApp.store} type="password"
    validators={[required('Enter password')]} />

  <Submit store={LoginApp.store} />
</Form>
