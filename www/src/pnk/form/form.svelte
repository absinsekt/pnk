<style src="./form.styl"></style>

<script>
  import { onDestroy, createEventDispatcher } from 'svelte';
  import { updateForm, sendForm } from './form';
  import { isUnset } from 'pnk/core/objects';

  const dispatch = createEventDispatcher();

  // store
  export let store;
  if (isUnset(store)) throw new Error("Required param (store) is missing!");

  // data
  export let data;
  // token
  export let token;

  let thisForm;

  async function submit(e) {
    updateForm(store, 'isTouched', true);

    if (thisForm.isValid) {
      const result = await sendForm(thisForm, token, data);
      dispatch(result.status, result);
    }
  }

  onDestroy(store.subscribe((m) => thisForm = m));
</script>

<div class="pnk-wgt">
  <form class="pnk-form"
    on:submit|preventDefault|stopPropagation={submit}>

    <slot />
  </form>
</div>
