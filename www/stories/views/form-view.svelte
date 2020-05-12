<script>
  import Form from 'app/ui-kit/form/form.svelte';
  import Input from 'app/ui-kit/input/input.svelte';
  import Submit from 'app/ui-kit/submit/submit.svelte';

  import { validName, validEmail, validPhone } from 'app/ui-kit/validators/validators';

  import { form } from './store';
</script>

<Form model={form}>
  <Input label="Name"
    name="name" model={form} validators={[validName]} />

  <Submit model={form} />
</Form>
