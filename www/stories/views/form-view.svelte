<script>
  import Form from 'ui/form/form.svelte';
  import Input from 'ui/input/input.svelte';
  import Submit from 'ui/submit/submit.svelte';

  import { buildNameValidator } from 'ui/validators/validators';

  import { form } from './store';
</script>

<Form store={form}>
  <Input label="Name"
    name="name" store={form} validators={[
      buildNameValidator('Введите фамилию и имя'),
    ]} />

  <Submit store={form} />
</Form>
