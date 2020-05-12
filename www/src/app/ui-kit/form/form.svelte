<script>
  import { onMount, onDestroy } from 'svelte';
  import { updateForm, sendForm } from './form';
  import { isUnset } from 'app/core/objects';

  export let action = '.';
  export let data;
  export let token;

  export let model;
  if (isUnset(model)) throw new Error("Required param (model) is missing!");

  let thisForm;
  const thisForm_unsubscribe = model.subscribe(m => thisForm = m);

  async function submit(e) {
    updateForm(model, 'isTouched', true);

    if (thisForm.isValid) {
      const result = await sendForm(thisForm, token, data);
      debugger;
    }
  }

  onMount(() => updateForm(model, 'action', action));

  onDestroy(thisForm_unsubscribe);
</script>

<form on:submit|preventDefault|stopPropagation={submit}>
  <slot />
</form>
