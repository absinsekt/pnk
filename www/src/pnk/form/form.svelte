<style src="./form.styl"></style>

<script>
  import { onMount, onDestroy, createEventDispatcher } from 'svelte';
  import { updateForm, sendForm } from './form';
  import { isUnset } from 'pnk/core/objects';

  const dispatch = createEventDispatcher();

  export let store;
  if (isUnset(store)) throw new Error("Required param (store) is missing!");

  export let action = '.';
  export let data;
  export let token;

  let thisForm;
  const thisForm_unsubscribe = store.subscribe(m => thisForm = m);

  async function submit(e) {
    updateForm(store, 'isTouched', true);

    if (thisForm.isValid) {
      const result = await sendForm(thisForm, token, data);
      dispatch(result.status, result);
    }
  }

  onMount(() => updateForm(store, 'action', action));

  onDestroy(thisForm_unsubscribe);
</script>

<div class="pnk-wgt">
  <form class="pnk-form"
    on:submit|preventDefault|stopPropagation={submit}>

    <slot />
  </form>
</div>
